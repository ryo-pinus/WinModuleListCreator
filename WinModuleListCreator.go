package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("USAGE:WinModuleListCreator DirectoryPath")
		return
	}

	// 指定したディレクトリに含まれるモジュールの情報をCSV出力する
	w := csv.NewWriter(os.Stdout)
	for _, info := range CreateWinModuleInformationList(os.Args[1]) {
		rec := []string{info.name, info.hash, info.buildTime.String(), info.linkerVersion}
		w.Write(rec)
	}
	w.Flush()
}

func CreateWinModuleInformationList(baseDirectory string) []*WinModuleInformation {
	var list []*WinModuleInformation
	var tagetExt = map[string]int{".exe": 1, ".dll": 2}

	filepath.Walk(baseDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if tagetExt[strings.ToLower(filepath.Ext(path))] != 0 {
				list = append(list, NewWinModuleInformation(path, info.Size()))
			}
		}
		return err
	})
	return list
}

type WinModuleInformation struct {
	name          string
	size          int64
	hash          string
	buildTime     *time.Time
	linkerVersion string
}

func NewWinModuleInformation(filePath string, fileSize int64) *WinModuleInformation {
	var buildTime, linkerVersion = ReadModuleInformation(filePath)
	obj := &WinModuleInformation{
		name:          filePath,
		size:          fileSize,
		hash:          CreateFileHash(filePath),
		buildTime:     buildTime,
		linkerVersion: linkerVersion,
	}
	return obj
}

// モジュールの情報を読み取る
func ReadModuleInformation(filePath string) (*time.Time, string) {
	var buildTime *time.Time
	var linkerVersion string

	file, err := os.Open(filePath)
	if err != nil {
		return buildTime, linkerVersion
	}
	defer file.Close()

	const OFFSET_OF_E_LFANEW int64 = 60
	const OFFSET_OF_TIME_DATE_STAMP int64 = 8
	const OFFSET_OF_LINKER_VERSION int64 = 26
	var offset int64

	// オフセットを取得
	offset = OFFSET_OF_E_LFANEW
	_, err = file.Seek(offset, 0)
	if err != nil {
		return buildTime, linkerVersion
	}
	e_lfanewBytes := make([]byte, 4)
	_, err = file.Read(e_lfanewBytes)
	if err != nil {
		return buildTime, linkerVersion
	}
	e_lfanew := binary.LittleEndian.Uint32(e_lfanewBytes)

	// ビルド日時を取得
	offset = int64(e_lfanew) + OFFSET_OF_TIME_DATE_STAMP
	_, err = file.Seek(offset, 0)
	if err != nil {
		return buildTime, linkerVersion
	}
	buildDateTimeBytes := make([]byte, 4)
	_, err = file.Read(buildDateTimeBytes)
	if err != nil {
		return buildTime, linkerVersion
	}
	utime := time.Unix(int64(binary.LittleEndian.Uint32(buildDateTimeBytes)), 0)
	buildTime = &utime

	// リンカのバージョン番号を取得
	offset = int64(e_lfanew) + OFFSET_OF_LINKER_VERSION
	_, err = file.Seek(offset, 0)
	if err != nil {
		return buildTime, linkerVersion
	}
	versionBytes := make([]byte, 2)
	_, err = file.Read(versionBytes)
	if err != nil {
		return buildTime, linkerVersion
	}
	linkerVersion = fmt.Sprintf("%d.%d", versionBytes[0], versionBytes[1])

	return buildTime, linkerVersion
}

//　ファイルのハッシュ値を作成する
func CreateFileHash(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	return fmt.Sprintf("%X", h.Sum(nil))
}

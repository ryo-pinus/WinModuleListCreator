package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestReadModuleInformation(t *testing.T) {
	path := "data_TestReadModuleInformation.dll"
	err := ioutil.WriteFile(path, []byte{
		0x4D, 0x5A, 0x90, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00,
		0xB8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		0x0E, 0x1F, 0xBA, 0x0E, 0x00, 0xB4, 0x09, 0xCD, 0x21, 0xB8, 0x01, 0x4C, 0xCD, 0x21, 0x54, 0x68,
		0x69, 0x73, 0x20, 0x70, 0x72, 0x6F, 0x67, 0x72, 0x61, 0x6D, 0x20, 0x63, 0x61, 0x6E, 0x6E, 0x6F,
		0x74, 0x20, 0x62, 0x65, 0x20, 0x72, 0x75, 0x6E, 0x20, 0x69, 0x6E, 0x20, 0x44, 0x4F, 0x53, 0x20,
		0x6D, 0x6F, 0x64, 0x65, 0x2E, 0x0D, 0x0D, 0x0A, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x50, 0x45, 0x00, 0x00, 0x4C, 0x01, 0x03, 0x00, 0x7F, 0x26, 0x4E, 0x59, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xE0, 0x00, 0x22, 0x00, 0x0B, 0x01, 0x30, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7E, 0x27, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00,
		0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x80, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x60, 0x85,
		0x00, 0x00, 0x10, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x10, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}, os.ModePerm)
	if err != nil {
		t.Fatalf("ioutil.WriteFile failed.")
	}
	defer os.Remove(path)

	buildTime, linkerVersion := ReadModuleInformation(path)
	if buildTime == nil || *buildTime != time.Date(2017, 6, 24, 17, 44, 47, 0, time.Local) {
		t.Fatalf("ReadModuleInformation() failed. buildTime:[%s]", buildTime)
	}
	if linkerVersion != "48.0" {
		t.Fatalf("ReadModuleInformation() failed. linkerVersion:[%s]", linkerVersion)
	}
}

func TestCreateFileHash(t *testing.T) {
	path := "data_TestCreateFileHash.dat"
	err := ioutil.WriteFile(path, []byte("abc"), os.ModePerm)
	if err != nil {
		t.Fatalf("ioutil.WriteFile failed.")
	}
	defer os.Remove(path)

	hash := CreateFileHash(path)
	if hash != "BA7816BF8F01CFEA414140DE5DAE2223B00361A396177A9CB410FF61F20015AD" {
		t.Fatalf("CreateFileHash() failed. [%s]", hash)
	}
}

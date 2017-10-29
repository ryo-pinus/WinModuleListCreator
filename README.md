# WinModuleListCreator

## 概要

`WinModuleListCreator`は、Windowsのモジュール(exe及びdll)の情報をリスト化するツールとなります。

## ビルド方法
~~~
go build
~~~
## 実行方法
 ~~~
WinModuleListCreator ディレクトリパス
~~~

 * 指定したディレクトリパスに含まれるモジュールの情報をCSV形式で標準出力に出力します。
 * 出力される情報は下記の通りとなります

    * モジュールのパス
    * モジュールのハッシュ値
    * モジュールのビルド日時
    * モジュールのリンカバージョン

## ライセンス
 MIT License
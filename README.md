# a-tour-of-go

このリポジトリは、プログラミング言語 Go のチュートリアルサービス "[A Tour of Go](https://go-tour-jp.appspot.com/)" を利用し学習を行った記録です。

この学習は、"A Tour of Go" で用意されているブラウザ上のエディタではなく、macOS に Go 環境を用意して実施しました。

## 環境構築
- OS: macOS High Sierra
- Editor: Atom
- Package Manager: Homebrew

### Go のインストール
```bash
$ brew install go
```

### go-delve (debugger) のインストール
[delve/install.md at master · go-delve/delve](https://github.com/go-delve/delve/blob/master/Documentation/installation/osx/install.md)
```bash
$ go get -u github.com/go-delve/delve/cmd/dlv
```
Homebrew でもインストールできるようです。(今回は上記 go get にてインストールしました)

### 利用する Atom パッケージ
- go-plus
- atom-runner

## コンパイルと実行
```bash
$ go run filename
```
Atom 上で実行する場合、Ctrl+R で atom-runner が起動します。

# a-tour-of-go

このリポジトリは、プログラミング言語 Go のチュートリアルサービス "[A Tour of Go](https://go-tour-jp.appspot.com/)" を利用し学習を行った記録です。"A Tour of Go" はブラウザ上にエディタおよび実行環境を用意していますが、あえて macOS に Go 環境を用意して学習を行いました。

## 環境構築
- OS: macOS High Sierra
- Editor: Atom
- Package Manager: Homebrew

### Go のインストール
```bash
$ brew install go
```

### go-delve (debugger) のインストール
```bash
$ go get -u github.com/go-delve/delve/cmd/dlv
```
[delve/install.md at master · go-delve/delve](https://github.com/go-delve/delve/blob/master/Documentation/installation/osx/install.md)

Homebrew でもインストールできるようです。(今回は上記 go get にてインストールしました)

### 利用する Atom パッケージ
- go-plus
- atom-runner

## 備忘録

### コンパイルと実行
```bash
$ go run filename
```
Atom 上で実行する場合、Ctrl+R で Atom-Runner が起動します。  
中断は Atom Runner のタブをクリックした後 Ctrl+Shift+C です。

### gotour

> この Go Tour はスタンドアロンのプログラムとして動かすこともできます。  
> (https://go-tour-jp.appspot.com/welcome/3)

`go tool tour` では動作しなかったので、`go get` でインストールしました。
インストールされた `gotour` は $GOPATH/bin にあります。
環境変数 GOPATH が設定されていない場合は `go env GOPATH` で確認できます。
デフォルトではおそらく ~/go でしょう。

```bash
$ go get github.com/atotto/go-tour-jp/gotour

$ go env GOPATH
/Users/username/go

$ ~/go/bin/gotour
```

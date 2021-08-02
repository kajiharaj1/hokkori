# ベースとなるDockerイメージ指定(作成時のLatest)
FROM golang:1.14.6

# コンテナ内に作業ディレクトリを作成
RUN mkdir -p /go/src/hokkori
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/hokkori
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ./ /go/src/hokkori

# モジュールインポート
RUN cd /go/src/hokkori

# godocインストール
RUN go get golang.org/x/tools/cmd/godoc
RUN go get github.com/gin-gonic/gin
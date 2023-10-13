# ベースとなるDockerイメージを指定
FROM golang:1.19

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコンテナにコピー
COPY go.mod .
COPY go.sum .

# 必要な依存関係をダウンロード
RUN go mod download

# その他のファイルやディレクトリをコンテナにコピー
COPY . .

# Goのアプリケーションをビルド
RUN go build -o server server.go

# コンテナが起動するときに実行されるコマンドを設定
CMD [ "./server" ]
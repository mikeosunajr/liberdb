all: liberdb

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

liberdb: cmd/liberdb/*.go
	go build -ldflags="-s -w" github.com/mikeosunajr/liberdb/cmd/liberdb
	upx liberdb


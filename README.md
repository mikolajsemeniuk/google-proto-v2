https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip
unzip protoc-3.19.4-osx-x86_64.zip

https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go-grpc_out=. *.proto
```

https://stackoverflow.com/questions/47704968/protoc-command-not-found-linux
```sh
PROTOC_ZIP=protoc-3.15.8-osx-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.15.8/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

```sh
sudo protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. \
    --go-grpc_opt=paths=source_relative \
    greet/*.proto
```
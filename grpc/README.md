## gRPC 使用

### protoc 安装

```
https://github.com/google/protobuf/releases下载包
```

### 安装protobuf库文件

```
go get github.com/golang/protobuf/protoc-gen-go
```

### 安装grpc文件

```
go get google.golang.org/grpc
```

### 构建proto并编译为 *.pb.go
```
protoc -I proto proto/hello.proto --go_out=plugins=grpc:proto

// protoc -I 固定语法
// proto proto/hello.proto  proto我们新建的文件夹 文件名hello.proto
// -go_out=plugins=grpc: 固定语法
// proto 生成的*.pb.go放的位置
```
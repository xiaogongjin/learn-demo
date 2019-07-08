package grpc

import (
	pb "./proto" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)
const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)
func main() {
	log.Println(">>>>>>>>")
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := pb.NewHelloClient(conn)
	// 调用方法
	reqBody := new(pb.HelloRequest)

	reqBody.Name = "gRPC"

	r, err := c.SayHello(context.Background(), reqBody)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.Message)
}

//run
//2019/07/08 15:39:37 >>>>>>>>
//2019/07/08 15:39:37 Hello gRPC.
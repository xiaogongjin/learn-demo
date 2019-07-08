package main
import (
	"net"
	pb "./proto" // 引入编译生成的包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)
const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)
// 定义helloService并实现约定的接口
type HelloService struct{}
//实现proto中定义的方法SayHello
func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + in.Name + "."
	return resp, nil
}
func main() {

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService{})
	log.Println("Listen on " + Address)
	s.Serve(listen)
}
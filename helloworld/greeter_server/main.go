package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/t67y110v/mesh/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	sleepDuration = flag.Duration("sleep duration ", 0, "fake waitiong duriung request handling")
	port          = flag.Int("port", 8080, "the server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received :%v", in.GetName())
	time.Sleep(*sleepDuration)
	log.Printf("Returned: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	log.Println("greeter  SERVER UP")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {

		log.Fatalf("failed to listent %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/t67y110v/mesh/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultName = "world"

var (
	addr = flag.String("addr", "localhost:50051", " the addres to connect to ")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	log.Println("greeter client UP")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("didnt connect %v", err)

	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("couldnt greet %v", err)

	}

	log.Printf("Greeting %s", r.GetMessage())

}

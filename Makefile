run_dns :
	go run .\dns_server\main.go

run_resolver :	
	go run .\resolver\main.go

run_client:
	go run .\helloworld\greeter_client\main.go

run_server: 
	go run .\helloworld\greeter_server\main.go


generate:
	protoc ./helloworld/helloworld/helloworld.proto --go_out=./helloworld/helloworld --go-grpc_out=./helloworld/helloworldls
	
package main

import (
	"context"
	"log"

	"gihub.com/wesleywillians/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

func main()  {
	// conexão
	connection, error := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if error != nil{
		log.Fatalf("Could not connect: %v", error)
	}
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	Hello(error, client)
	
}
// função remota
func Hello(error error, client pb.HelloServiceClient)  {
	resquest := &pb.HelloResquest{
		Name: "thiago",
	}

	response, error := client.Hello(context.Background(), resquest)
	if error != nil {
		log.Fatalf("error during the execution %v", error)
	}

	log.Println(response.Msg)
}
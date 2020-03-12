package main

import (
	pb "github.com/maybgit/microservice/example/proto"
	"github.com/maybgit/microservice/micro"
)

func main() {
	micro := &micro.MicService{}
	micro.NewServer()
	pb.RegisterUserServiceServer(micro.GrpcServer, &pb.UserService{})
	micro.Start()
}

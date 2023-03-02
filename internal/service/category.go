package main

import (
	"fmt"

	pb "github.com/rafatravassos/grpc-test/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func main() {
	fmt.Println("Hello")
}

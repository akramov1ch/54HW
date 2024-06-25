package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "54HW/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBookServiceServer
	books sync.Map
}

func (s *server) GetBookInfo(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	if val, ok := s.books.Load(req.GetTitle()); ok {
		book := val.(*pb.Book)
		return &pb.BookResponse{
			Title:     book.GetTitle(),
			Author:    book.GetAuthor(),
			Summary:   book.GetSummary(),
			Timestamp: time.Now().Format(time.RFC3339),
		}, nil
	}
	return nil, fmt.Errorf("book not found")
}

func (s *server) AddBook(ctx context.Context, book *pb.Book) (*pb.BookResponse, error) {
	s.books.Store(book.GetTitle(), book)
	return &pb.BookResponse{
		Title:     book.GetTitle(),
		Author:    book.GetAuthor(),
		Summary:   book.GetSummary(),
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

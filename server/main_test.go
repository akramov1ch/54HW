package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "54HW/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddBook(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(bufDialer))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewBookServiceClient(conn)

	book := &pb.Book{
		Title:   "Go Programming",
		Author:  "John Doe",
		Summary: "An introduction to Go programming language.",
	}

	resp, err := client.AddBook(ctx, book)
	if err != nil {
		t.Fatalf("AddBook failed: %v", err)
	}
	if !proto.Equal(resp, &pb.BookResponse{
		Title:     "Go Programming",
		Author:    "John Doe",
		Summary:   "An introduction to Go programming language.",
		Timestamp: resp.Timestamp,
	}) {
		t.Fatalf("AddBook response mismatch: got %v, want %v", resp, book)
	}
}

func TestGetBookInfo(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(bufDialer))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewBookServiceClient(conn)

	book := &pb.Book{
		Title:   "Go Programming",
		Author:  "John Doe",
		Summary: "An introduction to Go programming language.",
	}

	_, err = client.AddBook(ctx, book)
	if err != nil {
		t.Fatalf("AddBook failed: %v", err)
	}

	req := &pb.BookRequest{Title: "Go Programming"}
	resp, err := client.GetBookInfo(ctx, req)
	if err != nil {
		t.Fatalf("GetBookInfo failed: %v", err)
	}
	if resp.GetTitle() != book.GetTitle() || resp.GetAuthor() != book.GetAuthor() || resp.GetSummary() != book.GetSummary() {
		t.Fatalf("GetBookInfo response mismatch: got %v, want %v", resp, book)
	}
}

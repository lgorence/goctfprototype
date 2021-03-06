package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/lgorence/goctfprototype/pb"
	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Failed to run termcli: %v", err)
	}
}

func run() error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost:1234", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	proxy := pb.NewTermproxyServiceClient(conn)

	var callOpts []grpc.CallOption
	srv, err := proxy.OpenTerminal(context.Background(), callOpts...)
	if err != nil {
		return err
	}

	go func() {
		for {
			message, err := srv.Recv()
			if err == io.EOF {
				return
			}
			_, err = os.Stdout.Write(message.Contents)
			if err != nil {
				panic(err)
			}
		}
	}()

	stdinBuffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(stdinBuffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		err = srv.Send(&pb.TerminalBytes{
			Contents: stdinBuffer[:n],
		})
		if err != nil {
			return err
		}
	}

	return srv.CloseSend()
}

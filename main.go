package main

import (
	"ddia/server"
	"fmt"
)

func main() {
	s := server.New(&server.Config{
		Port: ":6379",
	})
	s.Start()
	fmt.Println("hello world")
}

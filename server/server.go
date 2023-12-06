package server

import (
	"ddia/logger"
	"fmt"
	"net"
)

const (
	network = "tcp"
	reply   = "+Ok\r\n"
)

type Config struct {
	Port string
}

type Server struct {
	*Config
}

func New(c *Config) *Server {
	return &Server{
		Config: c,
	}
}

func (s *Server) Start() {
	listner, err := net.Listen(network, s.Port)
	if err != nil {
		logger.Error("error openning")
	}
	defer func() {
		_ = listner.Close()
	}()

	for {
		conn, err := listner.Accept()
		if err != nil {
			logger.Error(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	buf := make([]byte, 1024)
	logger.Info("New Connection from ", conn.RemoteAddr().String())
	_, err := conn.Read(buf)
	if err != nil {
		logger.Error("Error when reading")
	}
	conn.Write([]byte(reply))
}

func main() {
	fmt.Println("hello world")
}

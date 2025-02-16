package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection", ws.RemoteAddr())

	// TODO use mutex
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) broadcast(msg []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msg); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error!", err)
			continue
		}
		msg := buf[:n]

		fmt.Println(string(msg))
		ws.Write([]byte("Thanks!"))
		s.broadcast(msg)
	}
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("sub request")

	for {
		payload := "aaa"
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func Serve() {
	server := NewServer()
	fmt.Println("Running!")
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/feed", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(":8080", nil)
}

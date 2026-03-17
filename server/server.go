package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server struct {
	clients   map[net.Conn]string
	broadcast chan string
	protocol  string
	address   string
	listener  net.Listener
	mu        sync.Mutex
}

func NewServer(protocol, address string) *Server {
	s := &Server{
		protocol:  protocol,
		address:   address,
		clients:   make(map[net.Conn]string),
		broadcast: make(chan string),
	}
	var err error
	s.listener, err = net.Listen(protocol, address)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fmt.Println("Chat server started on", address)
	return s
}


func (s *Server) Start() {
	defer s.close()
	go s.handleBroadcast()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		// Al connectar, assignem el RemoteAddr com a Nick per defecte
		s.mu.Lock()
		s.clients[conn] = conn.RemoteAddr().String()
		s.mu.Unlock()

		go s.handleClient(conn)

	}

}

func (s *Server) close() {
	if s.listener != nil {
		s.listener.Close()
	}
}


func (s *Server) handleClient(conn net.Conn) {
	//Definim una funció per a que s'executi quan s'acabi aquesta
	defer func() {
		s.mu.Lock()
		nick := s.clients[conn]
		delete(s.clients, conn)
		s.mu.Unlock()
		s.broadcast <- fmt.Sprintf("--- %s ha sortit del xat ---", nick)
		conn.Close()
	}()

	addr := conn.RemoteAddr().String()
	fmt.Printf("New client: %s\n", addr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()

		// 1. Gestió de la comanda /QUIT
		if strings.ToUpper(msg) == "/QUIT" {
			return 
		}

		// 2. Gestió de la comanda /NICK <nou_nickname>
		if strings.HasPrefix(msg, "/NICK ") {
			newNick := strings.TrimSpace(strings.TrimPrefix(msg, "/NICK "))
			s.handleNickChange(conn, newNick)
			continue
		}

		// 3. Missatge normal (Broadcast)
		s.mu.Lock()
		currentNick := s.clients[conn]
		s.mu.Unlock()
		
		s.broadcast <- fmt.Sprintf("[%s]: %s", currentNick, msg)
	}
}

func (s *Server) handleNickChange(conn net.Conn, newNick string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Comprovar si el nou nick ja existeix (ha de ser únic)
	for _, existingNick := range s.clients {
		if existingNick == newNick {
			fmt.Fprintln(conn, "ERROR: Aquest Nick ja està en ús.")
			return
		}
	}

	oldNick := s.clients[conn]
	s.clients[conn] = newNick
	fmt.Fprintln(conn, "Nick canviat correctament a:", newNick)
	s.broadcast <- fmt.Sprintf("--- %s ara es diu %s ---", oldNick, newNick)
}

func (s *Server) handleBroadcast() {
	for {
		msg := <-s.broadcast
		s.mu.Lock()
		for client := range s.clients {
			fmt.Fprintln(client, msg) // Send message to each client
		}
		s.mu.Unlock()
	}
}

func main() {
	s := NewServer("tcp", ":8080")
	if s != nil {
		s.Start()
	}
}

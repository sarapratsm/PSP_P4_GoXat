package server

import(
	"fmt"
	"sync"
	"client"
)

type Server struct{
	clients make(map[*Client]bool)
	address string
	protocol string
	mut sync.mutex
}


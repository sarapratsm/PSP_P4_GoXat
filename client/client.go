package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn     net.Conn
	protocol string
	address  string
}

func NewClient(protocol, address string) *Client {
	return &Client{
		protocol: protocol,
		address:  address,
	}
}

func (c *Client) Connect() {
	conn, err := net.Dial(c.protocol, c.address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	// guardem la connexió dins el client
	c.conn = conn

	go c.serverListener()
	fmt.Println("--- T'has connectat al xat ---")
	
	// entrem en bucle de lectura del teclat
	c.Read()
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) Read() {
	inputScanner := bufio.NewScanner(os.Stdin)
	
	for inputScanner.Scan() {
		text := inputScanner.Text()
		
		// Enviar el missatge al servidor
		fmt.Fprintln(c.conn, text)

		// Si l'usuari vol marxar...
		if strings.ToUpper(text) == "/QUIT" {
			fmt.Println("Desconnectant...")
			return
		}
	}
}

func (c *Client) serverListener() {
	// Aquesta funció corre en segon pla (goroutine)
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		// Imprimim el que arriba del servidor
		fmt.Printf("\r%s\n> ", scanner.Text()) 
	}
	fmt.Println("\nConexió tancada pel servidor.")
	os.Exit(0)
}

func main() {
	// Per provar-ho, creem el client i connectem
	c := NewClient("tcp", "localhost:4356")
	c.Connect()
}
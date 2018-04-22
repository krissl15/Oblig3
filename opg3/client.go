package main

import (
	"fmt"
	"bufio"
	"os"
	"net"
)

//For å teste serverfunksjonalitet kan udp og tcp skrives inn avhengig av hvilken server du skal koble til.
func main() {
	fmt.Println("Would you like to connect to the server through UDP or TCP?")
	nettverk := bufio.NewScanner(os.Stdin)
	for nettverk.Scan() {
		connection := strings.ToLower(nettverk.Text())
		p := make([]byte, 2048)
		conn, err := net.Dial(connection, "127.0.0.1:17")
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		fmt.Fprintf(conn, "Hello Server.")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}
}

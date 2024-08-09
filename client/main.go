package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableKeepAlives = true
	c := &http.Client{Transport: t}

	resp, err := c.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	bd, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bd))
}

//

// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"net"
// )

// func main() {
// 	// Define the server address and port (HTTPS typically uses port 443)
// 	serverAddr := "longconnection.fly.dev:443"

// 	// Resolve the TCP address
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
// 	if err != nil {
// 		fmt.Println("Error resolving address:", err)
// 		return
// 	}

// 	// Establish a TCP connection
// 	conn, err := net.DialTCP("tcp", nil, tcpAddr)
// 	if err != nil {
// 		fmt.Println("Error establishing TCP connection:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Configure TLS with the server name
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: false,                    // Set to true if you don't want to verify the server's certificate (not recommended in production)
// 		ServerName:         "longconnection.fly.dev", // The server name must match the certificate
// 	}

// 	// Upgrade the TCP connection to a TLS connection
// 	tlsConn := tls.Client(conn, tlsConfig)

// 	// Perform the TLS handshake
// 	err = tlsConn.Handshake()
// 	if err != nil {
// 		fmt.Println("Error during TLS handshake:", err)
// 		return
// 	}

// 	// Now you can send an HTTPS request
// 	request := "GET /longconnection HTTP/1.1\r\nHost: longconnection.fly.dev\r\n\r\n"
// 	_, err = tlsConn.Write([]byte(request))
// 	if err != nil {
// 		fmt.Println("Error writing to TLS connection:", err)
// 		return
// 	}

// 	// Read the HTTPS response
// 	buffer := make([]byte, 4096)
// 	n, err := tlsConn.Read(buffer)
// 	if err != nil {
// 		fmt.Println("Error reading from TLS connection:", err)
// 		return
// 	}

// 	// Print the response
// 	fmt.Println("HTTPS response:")
// 	fmt.Println(string(buffer[:n]))
// }

package httpClient

import (
	"bufio"
	"fmt"
	"net"
)

// PrintBody sends a generic GET request to the server at url and returns prints
// the body of the response to std Out
func PrintBody(url string) error {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return err
	}

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

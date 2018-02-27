package httpClient

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// PrintBody sends a generic GET request to the server at url and prints
// the body of the response to std Out
func PrintBody(url string) error {
	u, err := parseURL(url)
	if err != nil {
		return err
	}

	header := "GET " + u.path + " HTTP/1.1\r\n"
	header += "Host: " + u.hostname + "\r\n"
	header += "User-Agent: Noofbizzle\r\n"
	header += "Accept: text/html\r\n"
	header += "Accept-Language: en-us\r\n"
	header += "Accept-Encoding: gzip,deflate\r\n"
	header += "Accept-Charset: ISO-8859-1,utf-8\r\n\r\n"

	conn, err := net.Dial("tcp", "["+u.hostname+"]:"+u.port)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(header))
	if err != nil {
		return err
	}

	resp := bufio.NewReader(conn)
	respHeader := make(map[string]string)
	// get the header
	for {
		line, err := resp.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "\r\n" {
			break //header ends with an empty line
		}
		if strings.HasPrefix(line, "HTTP/") {
			respHeader["first"] = line
			continue
		}
		nv := strings.Split(line, ": ")
		respHeader[nv[0]] = strings.TrimSuffix(nv[1], "\r\n")
	}

	n, err := strconv.Atoi(respHeader["Content-Length"])
	if err != nil {
		return err
	}

	buf := make([]byte, n)
	_, err = resp.Read(buf)
	if err != nil {
		return err
	}

	fmt.Printf("%s", buf)

	return nil
}

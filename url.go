package httpClient

import (
	"errors"
	"strings"
)

type protocol uint

const (
	protocolhttp protocol = iota
	protocolhttps
	protocolUnsupported
)

type url struct {
	scheme   protocol
	hostname string
	path     string
	port     string
	values   map[string]string
}

func parseURL(in string) (url, error) {
	u := url{values: make(map[string]string)}
	if strings.HasPrefix(in, "http://") {
		u.scheme = protocolhttp
		in = in[7:]
	} else if strings.HasPrefix(in, "https://") {
		u.scheme = protocolhttps
		in = in[8:]
	} else {
		u.scheme = protocolUnsupported
		return u, errors.New("unsupported http protocol")
	}

	i := 0
	for i < len(in) {
		if in[i] == '/' {
			break
		}
		u.hostname += string(in[i])
		i++
	}

	for i < len(in) {
		if in[i] == ':' {
			u.port = in[i+1:]
			break
		}
		u.path += string(in[i])
		i++
	}

	if u.port == "" {
		u.port = "80"
	}

	return u, nil
}

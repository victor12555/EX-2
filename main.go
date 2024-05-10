package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

const (
	url = "https://distopia-a1e2.savi2w.workers.dev/"
)

func main() {
	tlsCustomConfig := tls.Config{
		MaxVersion: tls.VersionTLS12,
		MinVersion: tls.VersionTLS12,
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				con, err := tls.Dial(network, addr, &tlsCustomConfig)
				return con, err
			},
		},
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("erro criando requisição:", err)
		return
	}

	request.Header.Set("User-Agent", "victor")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("erro criando requisição:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
}

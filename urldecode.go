package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawUrl := scanner.Text()
		decodedUrl, err := url.QueryUnescape(strings.TrimSpace(rawUrl))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar URL: %s\n", err)
			continue
		}
		fmt.Println(decodedUrl)
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler entrada: %s\n", scanner.Err())
		os.Exit(1)
	}
}

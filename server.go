package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var dialer = &net.Dialer{Timeout: 10 * time.Second}

func runServer(config *tls.Config, cfg *conf) {
	l, err := tls.Listen("tcp", cfg.Listen, config)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		go handleConn(conn, cfg)
	}
}

func handleConn(conn net.Conn, cfg *conf) {
	dist, err := dialer.Dial("tcp", cfg.Backend)
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	pipeAndClose(conn, dist)
}

func pipeAndClose(c1, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		io.Copy(c1, c2)
		wg.Done()
	}()

	go func() {
		io.Copy(c2, c1)
		wg.Done()
	}()

	wg.Wait()
}

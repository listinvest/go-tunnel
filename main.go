package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

var dialer = &net.Dialer{Timeout: 10 * time.Second}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	var fn string

	flag.StringVar(&fn, "c", "config.json", "Config file")
	flag.Parse()

	cfg, err := loadConfig(fn)
	if err != nil {
		log.Fatal(err)
	}

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	config := &tls.Config{
		GetCertificate: certManager.GetCertificate,
	}

	createServer(config, cfg)
}

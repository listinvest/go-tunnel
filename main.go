package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/crypto/acme/autocert"
)

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

	runServer(config, cfg)
}

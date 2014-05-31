package main

import (
	"log"
	"os"

	"github.com/getlantern/httpconn"
)

func main() {
	proxy := httpconn.NewProxy(0)
	err := proxy.ListenAndServe(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to listen and serve: %s", err)
	}
}

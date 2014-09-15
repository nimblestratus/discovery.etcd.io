package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/coreos/go-systemd/activation"
	_ "github.com/coreos/discovery.etcd.io/http"

	"os"
)

var default_url = "https://discovery.etcd.io"
var addr = flag.String("addr", "", "web service address")
var url =  flag.String("url", default_url , "url prefix for discovery, defaults to " + default_url)



func main() {

	log.SetFlags(0)
	flag.Parse()

	os.Setenv("DISCOVERY_URL",*url)


	if *addr != "" {
		http.ListenAndServe(*addr, nil)
	}

	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		log.Print(len(listeners))
		panic("Unexpected number of socket activation fds:")
	}

	http.Serve(listeners[0], nil)
}

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/spf13/pflag"
)

type Flags struct {
	Listen      string
	UpstreamURL string
}

func main() {
	f := Flags{}
	pflag.StringVar(
		&f.Listen, "listen", "",
		"hostname:port to listen on. e.g. 127.0.0.1:8080",
	)
	pflag.StringVar(
		&f.UpstreamURL, "upstream", "",
		"upstream url to proxy to. e.g. https://example.com",
	)

	pflag.Parse()

	if f.Listen == "" || f.UpstreamURL == "" {
		log.Fatal("--listen and --upstream must be set")
	}

	u, err := url.Parse(f.UpstreamURL)
	if err != nil {
		log.Fatalf("invalid --upstream: %s", err.Error())
	}

	log.Fatal(http.ListenAndServe(
		f.Listen,
		httputil.NewSingleHostReverseProxy(u),
	))
}

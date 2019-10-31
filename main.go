package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const infoMsg = `CORS Proxy

Access-Control-Allow-Origin: *

http://localhost:%s -> %s
`


func main() {
	port := flag.String("port", "8001", "Port")
	flag.Parse()
	if flag.Arg(0) == "" {
		fmt.Fprintf(os.Stderr, "Specify URL to proxy. eg: corsproxy http://example.com\n")
		os.Exit(1)
	}
	proxyURL, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	proxyURL.Path = ""
	proxyURL.RawQuery = ""
	proxyURL.Fragment = ""

	reverseProxy := httputil.NewSingleHostReverseProxy(proxyURL)
	reverseProxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Add("Access-Control-Allow-Origin", "*")
		return nil
	}
	fmt.Println(fmt.Sprintf(infoMsg, *port, proxyURL.String()))
	http.ListenAndServe(":" + *port, reverseProxy)
}
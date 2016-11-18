package main

import (
        "flag"
        "log"
        "net/http"
        "net/http/httputil"
)

func main() {
        var addr string
        var ocspHost string
        flag.StringVar(&ocspHost, "ocsphost", "", "OCSP server to proxy requests to")
        flag.StringVar(&addr, "http", ":8080", "HTTP host:port to listen to")
        flag.Parse()
        if ocspHost == "" {
                log.Fatal("need ocsphost parameter")
        }
        rp := &httputil.ReverseProxy{
                Director: func(req *http.Request) {
                        req.URL.Scheme = "http"
                        req.URL.Host = ocspHost
                        req.Host = req.URL.Host
                },
                Transport: http.DefaultTransport,
        }
        http.ListenAndServe(addr, rp)
}

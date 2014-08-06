package main

import (
    "flag"
    "fmt"
    "net/http"
    "path/filepath"
)

var host = flag.String("host", "0.0.0.0", "Host for the server to listen on.")
var port = flag.String("port", "8080", "Port for the server to listen on.")
var path = flag.String("path", ".", "Path of the directory for the static " +
    "server to serve files from.")

func main() {
    flag.Parse()

    absPath, _ := filepath.Abs(*path)
    fmt.Println("Serving static files from directory", absPath, "on", *host + ":" + *port)

    panic(http.ListenAndServe(*host + ":" + *port, http.FileServer(http.Dir(*path))))
}

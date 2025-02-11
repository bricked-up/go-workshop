package router

import (
	"fmt"
	"log"
	"net/http"
)


func Init(port string) {
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(res, "Hello from my simple server!");
    });

    fmt.Println("Listening to port: ", port);
    log.Fatal(http.ListenAndServe(port, nil));
}


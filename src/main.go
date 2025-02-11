package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":3000";


func main() {
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(res, "Hello from my simple server!");
    });

    fmt.Println("Listening to port: ", PORT);
    log.Fatal(http.ListenAndServe(PORT, nil));
}

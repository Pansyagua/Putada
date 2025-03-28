package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hej från min Go-app! Nu fungerar det!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server startar på port 8080...")
    http.ListenAndServe(":8080", nil)
}

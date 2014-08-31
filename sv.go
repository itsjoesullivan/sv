package main

import (
  "fmt"
  "net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello, world")
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

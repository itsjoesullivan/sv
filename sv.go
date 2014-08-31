package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  filename := r.URL.Path[1:]
  body, err := ioutil.ReadFile(filename)
  if err != nil {
  }
  fmt.Fprintf(w, string(body[:]))
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
  "time"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  var startTime = time.Now()
  filename := r.URL.Path[1:]
  if r.URL.Path == "/" {
    filename = "index.html"
  }
  var status = "200";
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    status = "404";
    fmt.Fprintf(w, "error.")
  } else {
    fmt.Fprintf(w, string(body[:]))
  }
  var duration = string(time.Now().Sub(startTime).Nanoseconds() / 1000)
  log.Printf(r.Method + " " + filename + " " + status + " " + duration + "ms")
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type User struct {
  Id        int `json:"id"`
  Name      string `json:"name"`
  LastName   string `json:"lastName"`
}

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("GET /hello/{name}", helloWeb)

  log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
  //name := r.PathValue("name")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // Create basic user
  user := &User {
    Id: 1,
    Name: "Valentin",
    LastName: "Palacios",
  }

  fmt.Println(user)

  json.NewEncoder(w).Encode(user)
}

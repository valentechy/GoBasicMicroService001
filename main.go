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

type UsersList []User

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("GET /hello/{name}", helloWeb)

  fmt.Println("Server running on port 8181")
  log.Fatal(http.ListenAndServe(":8181", mux))
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
  //name := r.PathValue("name")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // Create basic user
  user1 := User {
    Id: 1,
    Name: "Valentin",
    LastName: "Palacios",
  }

  user2 := User {
    Id: 2,
    Name: "Juan",
    LastName: "Perez",
  }

  userList := []User{user1, user2}

  json.NewEncoder(w).Encode(&userList)
}

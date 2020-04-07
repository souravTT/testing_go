package controller

import (
  "net/http"
  "fmt"
)

var Group = struct {
  Index         func(w http.ResponseWriter, r *http.Request)
  Show          func(w http.ResponseWriter, r *http.Request)
  New           func(w http.ResponseWriter, r *http.Request)
  Create        func(w http.ResponseWriter, r *http.Request)
  Edit          func(w http.ResponseWriter, r *http.Request)
  Update        func(w http.ResponseWriter, r *http.Request)
  Delete        func(w http.ResponseWriter, r *http.Request)
  AssignRoles   func(w http.ResponseWriter, r *http.Request)
}{
  Index: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Tutorial")
  },
  Show: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Show Tutorial")
  },
  New: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis New Tutorial")
  },
  Create: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Create Tutorial")
  },
  Edit: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Edit Tutorial")
  },
  Update: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Update Tutorial")
  },
  Delete: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis Delete Tutorial")
  },
  AssignRoles: func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Go Redis AssignRoles Tutorial")
  },
}


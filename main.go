package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
  "github.com/souravTT/testing_go/controller"
)

func groupPaths(router *mux.Router){
	router.HandleFunc("/groups", controller.Group.Index).Methods(http.MethodGet)
	router.HandleFunc("/group/{id}", controller.Group.Show).Methods(http.MethodGet)
	router.HandleFunc("/group", controller.Group.Create).Methods(http.MethodPost)
	router.HandleFunc("/group/{id}", controller.Group.Update).Methods(http.MethodPut)
	router.HandleFunc("/group/{id}", controller.Group.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/group/{id}/assign_roles", controller.Group.AssignRoles).Methods(http.MethodPost)
}

func main(){
  Server()
}

func handleRequests() {
  router := mux.NewRouter().StrictSlash(true)

  groupPaths(router)

  host := "127.0.0.1"
  port := 3000

  address := ""
  if host != ""{
    address += host
  }
  address += ":"
  if port != 0 {
    address += strconv.Itoa(port)
  }
  _ = http.ListenAndServe(address, router)
}

func Server() {
  handleRequests()
}

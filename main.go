package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Name string
	Mail string
	Pass string
}

var user User

func GetForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user.Name = r.FormValue("username")
	user.Mail = r.FormValue("user-email")
	user.Pass = r.FormValue("user-pass")
	fmt.Printf("Name: %v, Login: %v, Pass: %v \n", user.Name, user.Mail, user.Pass)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", GetForm).Methods("GET")
	router.HandleFunc("/post", PostUser).Methods("POST")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Server error!")
	}
}

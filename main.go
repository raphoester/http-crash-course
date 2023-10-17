package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	ctr := &Controller{
		prefix: "hello",
	}

	router.HandleFunc("/register", ctr.Register)
	router.HandleFunc("/my-account/{username}", ctr.Account)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}

type Controller struct {
	prefix string
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "failed decoding input json", http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	usersDatabase = append(usersDatabase, user)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var usersDatabase []User

func (c *Controller) Account(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	var user *User
	for _, u := range usersDatabase {
		if u.Username == username {
			user = &u
			break
		}
	}

	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

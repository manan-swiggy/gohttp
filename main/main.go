package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type User struct {
	User string `json:"user"`
	Password string `json:"pass"`
}

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/update/user", UserHandler)
	http.HandleFunc("/get/user", GetHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "editting user name")
	u := User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}

	file, _ := json.Marshal(u)
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "editting user password")
}
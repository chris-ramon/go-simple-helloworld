package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name     string
	LastName string
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.Name, u.LastName)
}

func Index(w http.ResponseWriter, r *http.Request) {
	user1 := User{LastName: "apellido"}
	user1.SetName("el nombre")
	fullname := user1.FullName()
	log.Printf("FullName: %s", fullname)
	indexHTML, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Printf("failed to read html file", err)
		return
	}
	w.Write([]byte(string(indexHTML)))
}

func main() {
	http.HandleFunc("/", Index)
	log.Printf("running on port :8080")
	http.ListenAndServe(":8080", nil)
}

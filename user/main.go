package main

import (
	"go_web_cinema/pkg/render"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port for port assignment
const Port = "8082"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler).Methods("GET")
	r.HandleFunc("/user", userPatchHandler).Methods("PATCH")
	log.Printf("Starting on port %s", Port)
	log.Fatal(http.ListenAndServe(":"+Port, r))
}

type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	IsPaid bool   `json:"is_paid"`
	Pwd    string `json:"-"`
	Token  string `json:"token"`
}

type UserStorage []*User

var UU = UserStorage{
	&User{1, "bob@mail.ru", "Bob", true, "god", "1"},
	&User{2, "alice@mail.ru", "Alice", false, "secret", "2"},
}

func (uu UserStorage) GetByToken(token string) *User {
	for _, u := range uu {
		if u.Token == token {
			return u
		}
	}
	return nil
}

func (uu UserStorage) GetByID(id int) *User {
	for _, u := range uu {
		if u.ID == id {
			return u
		}
	}
	return nil
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")
	usr := UU.GetByToken(token)
	if usr == nil {
		render.RenderJSON
	}
}

func userPatchHandler(w http.ResponseWriter, r *http.Request) {

}

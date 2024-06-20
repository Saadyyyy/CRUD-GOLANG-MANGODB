package main

import (
	"fmt"
	"golang-mongo/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/create", uc.CreateUser)
	r.DELETE("/user/delete/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8812", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		fmt.Println(err)
	}
	return s
}

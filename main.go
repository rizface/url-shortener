package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load(".env")
}

func shorter() string {
	alfa := "abcdefghijklmnopqrstuvwxyz"
	var short []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		random := rand.Intn(27)
		short = append(short,alfa[random])
	}
	return string(short)
}

func main() {

	r := mux.NewRouter()
	var store = make(map[string]interface{})

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmp,err := template.ParseGlob("./view/*.gohtml")
		fmt.Println("Ini error")
		if err != nil {
			panic(err)
		}
		tmp.ExecuteTemplate(writer,"index.gohtml",nil)
	}).Methods(http.MethodGet)

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		url := request.PostFormValue("link")
		short := shorter()
		store[short] = url
		writer.Write([]byte(request.Host + "/" + short))
	}).Methods(http.MethodPost)


	r.HandleFunc("/{shortUrl}", func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		short := params["shortUrl"]
		url := store[short]
		if url == nil {
			http.Redirect(writer,request,"/",http.StatusSeeOther)
		} else {
			http.Redirect(writer,request,url.(string),http.StatusSeeOther)
		}
	}).Methods(http.MethodGet)

	server := http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}


	/*
	localhost:8080/ddipxqqpzr
	localhost:8080/jsqthqxqiv
	*/

}

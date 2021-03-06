package main

import (
	"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
	//"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

//var schemaDecoder = schema.NewDecoder()
//var sessionStore = sessions.NewCookieStore([]byte("your-secret-stuff-here"))

var templates map[string]*template.Template

func init() {
	loadTemplates()
}

func httpserver() {

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	router.HandleFunc("/", IndexRoute).Methods("GET")
	router.HandleFunc("/about", AboutRoute).Methods("GET")
	router.HandleFunc("/contact", ContactRoute).Methods("GET")
	router.HandleFunc("/signin", SigninRoute).Methods("GET")
	router.HandleFunc("/signup", SignupRoute).Methods("GET")

	if err := http.ListenAndServe(":9050", router); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//IndexRoute ... todo
func IndexRoute(res http.ResponseWriter, req *http.Request) {

	if err := templates["index"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//AboutRoute ... todo
func AboutRoute(res http.ResponseWriter, req *http.Request) {

	if err := templates["about"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//ContactRoute ... todo
func ContactRoute(res http.ResponseWriter, req *http.Request) {

	if err := templates["contact"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//SigninRoute ... todo
func SigninRoute(res http.ResponseWriter, req *http.Request) {
	if err := templates["signin"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//SignupRoute ... todo
func SignupRoute(res http.ResponseWriter, req *http.Request) {
	if err := templates["signup"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func loadTemplates() {
	var baseTemplate = "templates/layout/_base.html"
	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/index.html"))
	templates["about"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/about.html"))
	templates["contact"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/contact.html"))
	templates["signin"] = template.Must(template.ParseFiles(baseTemplate, "templates/account/signin.html"))
	templates["signup"] = template.Must(template.ParseFiles(baseTemplate, "templates/account/signup.html"))
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func ParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%v", string(b))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	//var r1 map[string]string
	a, _ := strconv.Atoi(r.Header["A"][0])
	b, _ := strconv.Atoi(r.Header["B"][0])
	res := strconv.Itoa(a + b)
	//r1["a+b"] = res
	fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Get("a+b")
	w.Header().Set("A+b", res)
	w.Header()["a+b"] = []string{res}
	fmt.Fprintf(w, "I got message:\n%v", res)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/bad", BadHandler).Methods("GET")
	router.HandleFunc("/name/{PARAM}", ParamHandler).Methods("GET")
	router.HandleFunc("/data", DataHandler).Methods("POST")
	router.HandleFunc("/headers", ArticlesHandler).Methods("POST")
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

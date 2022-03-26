
package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/pkg/handlers"
)

func main(){
	router:=mux.NewRouter()

	router.HandleFunc("/books",handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}",handlers.GetBook).Methods(http.MethodGet)
	// "/books"
	router.HandleFunc("/books/add/one",handlers.AddBooks).Methods(http.MethodGet)
	// "/books/{id}"
	router.HandleFunc("/books/u/{id}",handlers.UpdateBook).Methods(http.MethodGet)
	// "/books/{id}"
	router.HandleFunc("/books/dated/{id}",handlers.DeleteBook).Methods(http.MethodGet)


	log.Println("API is running!")
	http.ListenAndServe(":4000",router)
}
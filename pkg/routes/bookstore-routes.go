package routes

import (
	"github.com/K-Tech19/go-bookStore/pkg/controllers" // similar too "./controllers" if you are using Node.js. This is considered an absolute path (correct way in doing things)
	"github.com/gorilla/mux"
)


var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}
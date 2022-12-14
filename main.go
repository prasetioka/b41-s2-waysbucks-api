package main

import (
	"waysbucks-api/database"
	"waysbucks-api/pkg/mysql"
	"waysbucks-api/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// import "godotenv" here ...
	"github.com/joho/godotenv"
)

func main() {

	// Init godotenv here ...
	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Initialization "uploads" folder to public here ...
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
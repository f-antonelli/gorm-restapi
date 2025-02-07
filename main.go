package main

import (
	"net/http"

	"github.com/f-antonelli/gorm-restapi/db"
	"github.com/f-antonelli/gorm-restapi/models"
	"github.com/f-antonelli/gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	//	USERS
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")

	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")

	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")

	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//	TASKS
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")

	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")

	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

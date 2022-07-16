package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhtam3010/student/pkg/controllers"
	"github.com/minhtam3010/student/pkg/routes/middleware"
)

var RegisterStudentRoutes = func(routes *mux.Router) {
	routes.Use(middleware.LoggingMiddleware)

	routes.HandleFunc("/student/", controllers.GetAllStudent).Methods("GET")
	routes.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	routes.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	routes.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	routes.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
}

package controllers

import "github.com/surabiattack/fullstack/api/middlewares"

func (s *Server) InitializeRoutes() {
	//home
	s.Router.HandleFunc("/", middlewares.SetMiddlerwareJSON(s.Home)).Methods("GET")

	//login
	s.Router.HandleFunc("/login", middlewares.SetMiddlerwareJSON(s.Login)).Methods("POST")

	// Users Route
	s.Router.HandleFunc("/users", middlewares.SetMiddlerwareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlerwareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlerwareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlerwareJSON(middlewares.SetMiddlerwareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlerwareAuthentication(s.DeleteUser)).Methods("DELETE")

	// TODO POST
}
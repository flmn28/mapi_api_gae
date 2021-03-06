package main

import (
	"google.golang.org/appengine"
	"handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/locations/:id", handler.GetLocation)
	e.GET("/locations", handler.GetAllLocations)
	e.POST("/locations", handler.PostLocation)
	e.PUT("/locations/:id", handler.PutLocation)
	e.DELETE("/locations/:id", handler.DeleteLocation)

	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.PostUser)
	e.PUT("/users/:id", handler.PutUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.POST("/login", handler.Login)
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handler.Restricted)

	http.Handle("/", e)
	appengine.Main()
}

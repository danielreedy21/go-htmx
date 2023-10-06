package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type example struct {
	name string
	age  int16
}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/other", func(c echo.Context) error {
		return c.String(http.StatusOK, "Other route")
	})

	// here is going to be an htmx thing
	e.Logger.Fatal(e.Start(":1323"))
}

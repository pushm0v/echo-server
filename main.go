package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Host: c.Request().Host,
			Path: c.Path(),
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}

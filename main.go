package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", listEnvs())
	e.GET("/:name", listEnvs())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func listEnvs() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, os.Environ())
	}
}

func getEnv() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, os.Getenv(c.Param("name")))
	}
}

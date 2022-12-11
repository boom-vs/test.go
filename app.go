package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"tools/internal/app/endpoint"
	"tools/internal/app/mw"
	"tools/internal/app/servise"
)

type App struct {
	e    *endpoint.Endpoint
	s    *servise.Servise
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = servise.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RolCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil

}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

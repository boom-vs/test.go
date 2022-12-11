package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Servise interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Servise
}

func New(s Servise) *Endpoint {
	return &Endpoint{
		s: s,
	}

}
func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft() //time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	s := fmt.Sprintf("Days left, %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

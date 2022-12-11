package servise

import (
	"time"
)

type Servise struct {
}

func New() *Servise {
	return &Servise{}

}
func (s *Servise) DaysLeft() int64 {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	return int64(dur.Hours()) / 24
}

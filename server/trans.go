package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Trains(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")

}

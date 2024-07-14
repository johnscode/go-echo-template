package main

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	// Set custom JSONSerializer
	e.JSONSerializer = &echo.DefaultJSONSerializer{}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "go-echo-template")
	})

	userEndpoints := e.Group("/users")
	userEndpoints.GET("", userList)
	userEndpoints.GET("/:Id", getUser)

	e.Logger.Fatal(e.Start(":3000"))
}

func userList(c echo.Context) error {
	return c.JSON(http.StatusOK, UserList)
}

func getUser(c echo.Context) error {
	id := c.Param("Id")
	if u, ok := UserList[id]; ok {
		return c.JSON(http.StatusOK, u)
	}
	return c.JSON(http.StatusNotFound, fmt.Sprintf("No user found with Id, %s", id))
}

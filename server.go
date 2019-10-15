package main

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"bytes"
	"os/exec"
)

type (
	body struct {
		Url string `json:"url" form:"url" query:"url"`
	}
)

func main() {
	e := echo.New()

	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("reqBody: %v", string(reqBody[:]))
	}))
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hey! This is pdf`s builder service")
	})
	e.POST("/pdf", func(c echo.Context) error {
		b := &body{}
		if err := c.Bind(b); err != nil {
			return err
		}
		fmt.Println("url: ", b.Url)
		url := b.Url
		if len(url) == 0 {
			url = "google.com"
		}
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		if len(requestID) == 0 {
			requestID = "test"
		}
		path := fmt.Sprintf("static/pdf/%s.pdf", requestID)
		app := fmt.Sprintf("docker run idocking/wkhtmltopdf wkhtmltopdf %s - > %s", url, path)
		fmt.Println("exec: ", app)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd := exec.Command("/bin/sh", "-c", app)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())

		return c.File(path)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

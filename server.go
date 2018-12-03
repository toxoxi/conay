package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/yorori/golang-echo-react-redux/server/qiita"
)

func main() {
	loadErr := godotenv.Load("env/develop.env")
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}
	// 各種アクセストークン取得
	qiitaToken := os.Getenv("QIITA_TOKEN")

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/qiita/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		fmt.Println("id:", id)
		articles, err := qiita.GetUserArticles(qiitaToken, id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(articles)
		return c.JSON(http.StatusOK, articles)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

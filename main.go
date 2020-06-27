package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	/// logs configs
	myLog, err := os.OpenFile("logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("no se pudo abrir el archivo de logs")
	}

	defer myLog.Close()

	logConfig := middleware.LoggerConfig{
		Output: myLog,
	}

	e.Use(middleware.LoggerWithConfig(logConfig))

	/// CORS configs
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://some-allowed-url.com"},
	}))

	/// cors test
	e.GET("/data", LoadData)

	/// String http response
	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "hola mundo")
	})

	/// Static files http response
	e.Static("/static", "static")

	e.GET("/imagen/:name", func(c echo.Context) error {

		/// Mostrar archivos en browser
		if c.Param("name") == "imagen1" {
			return c.File("imgs/imagen1.jpg")
		}
		if c.Param("name") == "imagen2" {
			return c.File("imgs/imagen2.jpg")
		}

		/// Para descargar archivo
		if c.Param("name") == "att" {
			return c.Attachment("imgs/imagen2.jpg", "nameInDownload.jpg")
		}
		return c.HTML(http.StatusNotFound, "<h2>No encontrado</h2>")
	})

	/// html response
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>HTML CONTENT</h1>")
	})

	/// no content response
	e.GET("/no-content", func(e echo.Context) error {
		return e.NoContent(http.StatusOK)
	})

	/// json response
	e.GET("/person", func(e echo.Context) error {
		p := Person{
			FirstName: "Marcos",
			LastName:  "Carlomagno",
			Age:       25,
		}
		return e.JSON(http.StatusOK, p)
	})

	e.File("/", "static")

	e.Start(":8080")
}

// Person cool interface
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

/// LoadData loads data for cors testing
func LoadData(c echo.Context) error {
	return c.String(http.StatusOK, "data loaded!")
}

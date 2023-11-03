package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/api"
	"github.com/linweiyuan/chatgpt-arkose-token-api/browser"
	"github.com/gin-contrib/cors"
)

func init() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

//goland:noinspection GoUnhandledErrorResult
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// Create a CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // You can specify the allowed origins here

	// Apply the CORS middleware with the configuration
	router.Use(cors.New(config))


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arkose.html", nil)
	})

	router.GET("/token", api.GetArkoseToken)
	router.GET("/bx", api.GetBX)

	router.Run(fmt.Sprintf(":%d", browser.Port))
}

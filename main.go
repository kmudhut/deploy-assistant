package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"time"
)

func main() {

	var err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	var tokenFromEnv = os.Getenv("TOKEN")

	var router = gin.Default()

	router.POST("/pull", func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == tokenFromEnv {
			var reqBody map[string]string
			if err := c.ShouldBindJSON(&reqBody); err != nil {
				c.JSON(400, gin.H{"error": "Tory były złe, a JSON też był zły"})
				return
			}

			cmd := exec.Command("git",
				"-C", reqBody["appName"],
				"pull",
				"origin",
				reqBody["branch"],
			)

			out, err := cmd.CombinedOutput()
			if err != nil {
				c.String(400, "git pull się nie udał:", string(out)+" "+err.Error())
			} else {
				c.String(200, "Udało się: "+string(out))
			}

		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] " + "Niepoprawny token:" + token)
			c.String(401, "Niepoprawny token.")
			return
		}
	})
	router.Run()
}

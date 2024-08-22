package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading environment variable")
	}

	apikey := os.Getenv("API_KEY")
	port := os.Getenv("PORT")
	ctx := context.Background()

	model := createModel(apikey, ctx)

	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/talk", func(c *gin.Context) {
			q := c.Query("q")
			var msg Message
			c.BindJSON(&msg)
			switch q {
			case "good":
				resp, err := generateContentGood(model, msg)
				assertGenerationError(err)
				c.JSON(http.StatusOK, gin.H{
					"message": resp.Candidates[0].Content.Parts[0],
				})

			case "bad":
				resp, err := generateContentBad(model, msg)
				assertGenerationError(err)
				c.JSON(http.StatusOK, gin.H{
					"message": resp.Candidates[0].Content.Parts[0],
				})

			}
		})
	}

	log.Printf("Starting Server at port %s", port)
	r.Run(port)

}

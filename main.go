package main

import (
	  "fmt"
	  "os"
	  "github.com/gin-gonic/gin"
	  "github.com/joho/godotenv"
	  congnitoClient "github.com/jeanroths/OAuth_2.0_P-D_local/cognitoClient"
	  "errors"
	  "net/http"
)

func CreateUser(c *gin.Context, cognito congnitoClient.CognitoInterface) error {
    var user congnitoClient.User
    if err := c.ShouldBindJSON(&user); err != nil {
      return errors.New("invalid json")
    }
    err := cognito.SignUp(&user)
    if err != nil {
      return errors.New("could not create use")
    }
    return nil
  }

func main() {
    err := godotenv.Load()
    if err != nil {
      panic(err)
    }
    cognitoClient := congnitoClient.NewCognitoClient(os.Getenv("COGNITO_CLIENT_ID"))
    r := gin.Default()

	r.POST("user", func(context *gin.Context) {
        err := CreateUser(context, cognitoClient)
        if err != nil {
            context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        context.JSON(http.StatusCreated, gin.H{"message": "user created"})
    })


    fmt.Println("Server is running on port 8080")
    err = r.Run(":8080")
    if err != nil {
      panic(err)
    }
  }

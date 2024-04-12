package controller

import (
	"fmt"

	resterr "github.com/Flavio-coutinho/CRUD-GoLang/src/configuration/rest_err"
	"github.com/Flavio-coutinho/CRUD-GoLang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

		var userRequest request.UserRequest

		if err := c.ShouldBindJSON(&userRequest); err != nil {
				restErr := resterr.NewBadRequestError(

					fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error))

					c.JSON(restErr.Code, restErr)
					return
		}
		fmt.Println(userRequest)
}
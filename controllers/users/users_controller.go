package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shubh199815/bookstore_users-api/domain/users"
	// "encoding/json"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/shubh199815/bookstore_users-api/services"
	"github.com/shubh199815/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err !=nil {
	// 	fmt.Println(err.Error())
	// 	// Handle JSON error
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to the caller
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO Handle User creation
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id shpuld be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

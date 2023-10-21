package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runnersApp/models"
	"runnersApp/services"

	"github.com/gin-gonic/gin"
)



type RunnersController struct {
	runnersService *services.RunnersService
	usersService   *services.UsersService
}

func NewRunnersController(runnersService *services.RunnersService, userService *services.UsersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
		usersService: userService,
	}
}

func (rh RunnersController) CreateRunner(ctx *gin.Context) {

	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ROLE_ADMIN})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Error while unmarshaling  + create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh RunnersController) UpdateRunner(ctx *gin.Context) {

	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ROLE_ADMIN})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading + update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Error while unmarshaling  + update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr = rh.runnersService.UpdateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusOK)

}

func (rh RunnersController) DeleteRunner(ctx *gin.Context) {


	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ROLE_ADMIN})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	runnerId := ctx.Param("id")
	responseErr = rh.runnersService.DeleteRunner(runnerId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)

}
func (rh RunnersController) GetRunner(ctx *gin.Context) {


	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ROLE_ADMIN, ROLE_RUNNER})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	runnerId := ctx.Param("id")
	response, responseErr := rh.runnersService.GetRunner(runnerId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
func (rh RunnersController) GetRunnersBatch(ctx *gin.Context) {

	accessToken := ctx.Request.Header.Get("Token")
	auth, responseErr := rh.usersService.AuthorizeUser(accessToken, []string{ROLE_ADMIN, ROLE_RUNNER})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")
	response, responseErr := rh.runnersService.GetRunnersBatch(country, year)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

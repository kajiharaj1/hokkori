package routes

import (
	"net/http"
	"src/config"

	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")

	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconfirmation")

	if password != passwordConf {
		println("Error: password and passwordConf not match")
		ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
	}

	db := config.DummyDB()
	if err := db.SaveUser(username, email, password); err != nil {
		println("Error: " + err.Error())
	} else {
		println("username: " + username)
		println("emailaddress: " + email)
		println("password: " + password)
		println("passwordConf: " + passwordConf)
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogin(ctx *gin.Context) {
	println("post/login")

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: " + err.Error())
	} else {
		println("username: " + username)
		println("password: " + password)
		user.Authenticate()
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
}

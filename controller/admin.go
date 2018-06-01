package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/1377195627/goblog/core"
)

type Admin struct {
}

func (a *Admin) Login(ctx *gin.Context) {
	type Data struct {
		Password string `json:"password" form:"password" binding:"required"`
	}

	data := Data{}

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(200, gin.H{
			"code":    100,
			"message": "password is null",
		})
		return
	}

	session := sessions.Default(ctx)

	md5Data :=md5.Sum([]byte(data.Password))
	sha1Data :=sha1.Sum([]byte(md5Data[:]))
	passwd := hex.EncodeToString(sha1Data[:])

	if passwd == core.Conf.Password {
		session.Set("login",true)
		session.Save()

		ctx.JSON(200, gin.H{
			"code": 0,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    100,
			"message": "password errors",
		})
	}
}

func (a *Admin) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)

	login := session.Get("login")
	if login != nil {
		session.Set("login", nil)
		session.Save()
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    100,
			"message": "no login",
		})
	}
}
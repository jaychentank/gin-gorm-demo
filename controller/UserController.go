package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = RandomString(10)
	}

	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(422, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	//返回响应
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

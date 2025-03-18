package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/go_gateway/dto"
	"github.com/lining4069/go_gateway/middleware"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}       // 登录控制器
	group.POST("/login", adminLogin.AdminLogin) // 登录接口<绑定>控制器处理函数
}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

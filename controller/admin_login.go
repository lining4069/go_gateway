package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/go_gateway/dto"
	"github.com/lining4069/go_gateway/middleware"
)

type AdminLoginController struct {
}

// AdminLoginRegister 登录功能子路由注册
func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}       // 登录控制器
	group.POST("/login", adminLogin.AdminLogin) // 登录接口<绑定>控制器处理函数
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	out := &dto.AdminLoginOutput{Token: params.UserName}
	middleware.ResponseSuccess(c, out)
}

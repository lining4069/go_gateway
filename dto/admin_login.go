package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/go_gateway/public"
)

type AdminLoginInput struct {
	UserName string `json:"name" form:"name" comment:"姓名" example:"姓名" validate:"required is_valid_username"`
	Password string `json:"passwd" form:"passwd" comment:"密码" example:"123456" validate:"required"`
}

// BindValidParam 检验参数从请求数据中解析为结构体实例是否成功
func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wangys891019/blog-service/global"
	"github.com/wangys891019/blog-service/pkg/convert"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	// 如果没有传递,默认page为1
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	// 没有传递,为1
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}
	// 超过全局的配置,使用全局的配置
	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

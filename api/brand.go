package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/brand/models"
	"github.com/provider-go/pkg/ioput"
)

func CreateBrand(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	brandName := ioput.ParamToString(json["brandName"])
	brandPic := ioput.ParamToString(json["brandPic"])
	brief := ioput.ParamToString(json["brief"])
	content := ioput.ParamToString(json["content"])
	firstChar := ioput.ParamToString(json["firstChar"])
	err := models.CreateBrand(brandName, brandPic, brief, content, firstChar)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteBrand(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	err := models.DeleteBrand(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func ListBrand(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := ioput.ParamToInt(json["pageSize"])
	pageNum := ioput.ParamToInt(json["pageNum"])
	list, total, err := models.ListBrand(pageSize, pageNum)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		ioput.ReturnSuccessResponse(ctx, res)
	}
}

func ViewBrand(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	row, err := models.ViewBrand(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, row)
	}
}

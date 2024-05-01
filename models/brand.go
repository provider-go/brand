package models

import (
	"github.com/provider-go/brand/global"
	"time"
)

type Brand struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	BrandName  string    `json:"brandName" gorm:"column:brand_name;type:varchar(50);not null;default:'';comment:品牌名称"`
	BrandPic   string    `json:"brandPic" gorm:"column:brand_pic;type:varchar(200);not null;default:'';comment:图片路径"`
	Brief      string    `json:"brief" gorm:"column:brief;type:varchar(255);not null;default:'';comment:简要描述"`
	Content    string    `json:"content" gorm:"column:content;type:text;not null;comment:内容"`
	FirstChar  string    `json:"firstChar" gorm:"column:first_char;type:char(1);default:null;comment:品牌首字母"`
	Status     int       `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:状态：0：正常；1：隐藏"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateBrand(brandName, brandPic, brief, content, firstChar string) error {
	return global.DB.Table("brands").Select("brand_name", "brand_pic", "brief", "content", "first_char").
		Create(&Brand{BrandName: brandName, BrandPic: brandPic, Brief: brief, Content: content, FirstChar: firstChar}).Error
}

func DeleteBrand(id int32) error {
	return global.DB.Table("brands").Where("id = ?", id).Delete(&Brand{}).Error
}

func ListBrand(pageSize, pageNum int) ([]*Brand, int64, error) {
	var rows []*Brand
	//计算列表数量
	var count int64
	global.DB.Table("brands").Count(&count)

	if err := global.DB.Table("brands").Order("id desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewBrand(id int32) (*Brand, error) {
	row := new(Brand)
	if err := global.DB.Table("brands").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

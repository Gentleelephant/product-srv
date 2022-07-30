package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int32          `json:"id" gorm:"primary_key"`
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

// Category product category
type Category struct {
	BaseModel
	Name             string      `json:"name" gorm:"type:varchar(64);not null"`
	ParentCategoryId int32       `json:"parent_category_id" gorm:"type:int(11);not null"`
	ParentCategory   *Category   `json:"parent_category"`
	SubCategories    []*Category `json:"sub_categories" gorm:"foreignKey:ParentCategoryId;references:ID"`
}

type Brand struct {
	BaseModel
	Name string `json:"name" gorm:"type:varchar(64);not null"`
	Logo string `json:"logo" gorm:"type:varchar(255);not null;default:''"`
}

type Advertise struct {
	BaseModel
	Index int32  `json:"index" gorm:"type:int(11);not null;default:1"`
	Image string `json:"image" gorm:"type:varchar(255);not null"`
	Url   string `json:"url" gorm:"type:varchar(255);not null"`
	Sort  int32  `json:"sort" gorm:"type:int(11);not null;default:1"`
}

// Product product
type Product struct {
	BaseModel
	CategoryId int32    `json:"category_id" gorm:"type:int(11);not null"`
	Category   Category `json:"category"`

	BrandId int32 `json:"brand_id" gorm:"type:int(11);not null"`
	Brand   Brand `json:"brand"`

	Selling  bool `json:"selling" gorm:"type:tinyint(1);not null;default:0"`
	ShipFree bool `json:"ship_free" gorm:"type:tinyint(1);not null;default:0"`
	// 是否热卖
	IsPop bool `json:"is_pop" gorm:"type:tinyint(1);not null;default:0"`
	// 是否新品
	IsNew bool `json:"is_new" gorm:"type:tinyint(1);not null;default:0"`

	Name string `json:"name" gorm:"type:varchar(64);not null"`
	SN   string `json:"sn" gorm:"type:varchar(64);not null"`
	// 被收藏
	FavNum int32 `json:"fav_num" gorm:"type:int(11);not null;default:0"`
	// 被购买
	SoldNum int32 `json:"sold_num" gorm:"type:int(11);not null;default:0"`
	// price
	Price float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	// real price
	RealPrice float64 `json:"real_price" gorm:"type:decimal(10,2);not null"`

	// short desc
	ShortDesc string `json:"short_desc" gorm:"type:varchar(255);not null"`
	// images
	Images []string `json:"images" gorm:"type:varchar(1024);not null"`
	// cover image
	CoverImage string `json:"cover_image" gorm:"type:varchar(255);not null"`
}

type MyList []string

func (l *MyList) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *MyList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(b, l)
}

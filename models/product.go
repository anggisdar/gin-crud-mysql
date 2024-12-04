package models

type Product struct {
	Id          int    `gorm:"primaryKey" json:"ìd"`
	NameProduct string `gorm:"type:varchar(300)" json:"name_product"`
	Description string `gorm:"type:text" json:"description"`
}

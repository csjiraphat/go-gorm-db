// models/Item.go
package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model         // gorm จะสร้าง ID, CreatedAt, UpdatedAt, DeletedAt ให้เอง
	Name       string  // ชื่อสินค้า
	Price      float64 // ราคาสินค้า
}

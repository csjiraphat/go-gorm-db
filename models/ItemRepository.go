// models/ItemRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// สร้าง struct ชื่อ ItemRepository ที่มีฟิลด์ชื่อ Db เป็น pointer ของ gorm.DB
type ItemRepository struct {
	Db *gorm.DB
}

// ทำหน้าที่สร้าง Instance ของ ItemRepository และส่งคืนกลับไป ให้เรียกใช้งานได้
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{Db: db}
}

// ทำหน้าที่ดึงข้อมูล Item ทั้งหมดจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งาน
// (r *ItemRepository) คือการระบุว่าเป็น method ของ ItemRepository
// r คือตัวแปรที่เป็น pointer ของ ItemRepository
// ฟังก์ชัน GetItems รับพารามิเตอร์เป็น c *gin.Context และมีการส่งกลับค่าเป็น JSON กลับไปให้ผู้ใช้งานผ่าน c.JSON(200, items) ซึ่ง c คือตัวแปรที่เป็น pointer ของ gin.Context
// ในกรณีนี้เราไม่ต้องใช้ return เพราะเราใช้ c.JSON แทน
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *ItemRepository) GetItems(c *gin.Context) {
	var items []Item
	// ดึงข้อมูล Item ทั้งหมดจากฐานข้อมูล และเก็บลงในตัวแปร items
	r.Db.Find(&items)  // SELECT * FROM items
	c.JSON(200, items) // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ทำหน้าที่เพิ่มข้อมูล Item ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, newItem)
func (r *ItemRepository) PostItem(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)  // รับค่า JSON จากผู้ใช้งาน และแปลงเป็น struct ของ Item
	r.Db.Create(&newItem) // INSERT INTO items (name, price) VALUES (newItem.Name, newItem.Price)
	c.JSON(200, newItem)  // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันค้นหา Item จากฐานข้อมูล โดยใช้ id ที่รับเข้ามาเป็นเงื่อนไขในการค้นหา
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *ItemRepository) GetItem(c *gin.Context) {
	id := c.Param("id")   // รับค่า id จากผู้ใช้งาน
	var item Item         // สร้างตัวแปร item เพื่อเก็บข้อมูลที่ค้นหาได้
	r.Db.First(&item, id) // SELECT * FROM items WHERE id = id
	c.JSON(200, item)     // ส่งข้อมูลกลับไปให้ผู้ใช้งาน
}

// ฟังก์ชันอัพเดทข้อมูล Item ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, item)
func (r *ItemRepository) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	r.Db.First(&item, id) // SELECT * FROM items WHERE id = id
	c.BindJSON(&item)
	r.Db.Save(&item) // UPDATE items SET name = item.Name, price = item.Price WHERE id = id
	c.JSON(200, item)
}

// ฟังก์ชันลบข้อมูล Item ออกจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, gin.H{"id" + id: "is deleted"})
// gin.H ทำหน้าที่สร้าง map ของ key และ value และส่งกลับไปให้ผู้ใช้งาน
// H คือตัวย่อของ Header ใน HTTP
func (r *ItemRepository) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	r.Db.Delete(&item, id) // DELETE FROM items WHERE id = id
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}

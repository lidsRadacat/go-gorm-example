package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Age         int
	Name        string
	Description string
	CityID      int
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	defaultTableNameHandler()

	fmt.Println("has table : ", db.HasTable(&User{})) // false
	fmt.Println("has table : ", db.HasTable("users")) // false

	createTable(db)

	fmt.Println("has table : ", db.HasTable(&User{})) // true

	addIndex(db)

	deleteTable(db)

	fmt.Println("has table : ", db.HasTable(&User{})) // false
}

func createTable(db *gorm.DB) {
	db.CreateTable(&User{})
}

func deleteTable(db *gorm.DB) {
	db.DropTable(&User{})
}

func addIndex(db *gorm.DB) {
	db.Model(&User{}).AddIndex("idx_user_name", "name")
}

func defaultTableNameHandler() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
}

// TableName这个优先级比 DefaultTableNameHandler 这个高
func (User) TableName() string {
	return "profiles"
}

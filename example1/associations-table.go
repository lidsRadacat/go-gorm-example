package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User1 struct {
	gorm.Model
	Profile   Profile
	ProfileID uint
}

type Profile struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.LogMode(true)

	db.CreateTable(&Profile{})
	db.CreateTable(&User1{})
	profile := Profile{Name: "aaa"}
	user := User1{Profile: profile}
	db.Create(&user)

	findUser(user.ProfileID, db)
}

func findUser(id uint, db *gorm.DB) {
	user := User1{}
	user.ProfileID = id
	profile := Profile{}
	db.Model(&user).Related(&profile)
	fmt.Println(user)
	fmt.Println(profile)
}

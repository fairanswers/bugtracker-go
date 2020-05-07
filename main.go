package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Status struct {
	gorm.Model
	Name string
}

type Severity struct {
	gorm.Model
	Name string
}

type Issue struct {
	gorm.Model
	Title    string
	Desc     string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Title string
	Desc  string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Status{}, &Severity{}, &Issue{}, &Comment{})

	// Create
	user := User{Name: "User1", Email: "user@example.com"}
	db.Create(&user)
	fmt.Println(user)
	// // Read
	// var product Product
	// var user User
	// db.First(&user, 1)                               // find product with id 1
	// db.First(&user, "email = ?", "user@example.com") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)

	fmt.Println("SFSG")
}

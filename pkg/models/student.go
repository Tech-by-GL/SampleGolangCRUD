package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/minhtam3010/student/pkg/config"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Year  string `json:"year"`
	Major string `json:"major"`
}

var Students []Student

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Students)
}

func (s *Student) CreateStudent() (*Student, gorm.Errors) {
	db.NewRecord(s)

	err := db.Create(&s)
	if err != nil {
		log.Println(err)
	}

	return s, err.GetErrors()
}

func GetAllStudent() ([]Student, gorm.Errors) {
	var Students []Student
	err := db.Find(&Students)
	if err != nil {
		log.Println(err)
	}

	return Students, err.GetErrors()
}

func GetStudentById(ID int64) (*Student, gorm.Errors) {
	var student Student
	err := db.Where("ID=?", ID).Find(&student)
	if err != nil {
		log.Println(err)
	}
	return &student, err.GetErrors()
}

func DeleteStudent(ID int64) (Student, gorm.Errors) {
	var student Student
	err := db.Where("ID=?", ID).Delete(student)
	if err != nil {
		log.Println(err)
	}
	return student, err.GetErrors()
}

func UpdateStudent(studentDetails Student) gorm.Errors {
	err := db.Save(&studentDetails)
	if err != nil {
		log.Println(err)
	}
	return err.GetErrors()
}

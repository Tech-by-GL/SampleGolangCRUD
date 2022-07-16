package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/student/pkg/models"
	"github.com/minhtam3010/student/pkg/utils"
)

func WriteResponse(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetAllStudent()
	WriteResponse(w, []byte(err.Error()))
	if res, err := json.Marshal(students); err != nil {
		WriteResponse(w, []byte(err.Error()))
	} else {

		WriteResponse(w, res)
	}
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	studentDetails, err := models.GetStudentById(ID)
	WriteResponse(w, []byte(err.Error()))

	if res, err := json.Marshal(studentDetails); err != nil {
		WriteResponse(w, []byte(err.Error()))
	} else {

		WriteResponse(w, res)
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	createStudent := &models.Student{}
	utils.ParseBody(r, createStudent)
	s, err := createStudent.CreateStudent()
	WriteResponse(w, []byte(err.Error()))
	if res, err := json.Marshal(s); err != nil {
		WriteResponse(w, []byte(err.Error()))
	} else {

		WriteResponse(w, res)
	}
}

func Update(studentDetails *models.Student, updateStudent *models.Student) {
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.Year != "" {
		studentDetails.Year = updateStudent.Year
	}
	if updateStudent.Major != "" {
		studentDetails.Major = updateStudent.Major
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	updateStudent := &models.Student{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	studentId := vars["studentId"]

	var ID int64
	var err error

	if ID, err = strconv.ParseInt(studentId, 0, 0); err != nil {
		fmt.Println("error while parsing")
	}

	studentDetails, err := models.GetStudentById(ID)
	WriteResponse(w, []byte(err.Error()))
	Update(studentDetails, updateStudent)

	if res, err := json.Marshal(studentDetails); err != nil {
		WriteResponse(w, []byte(err.Error()))
	} else {

		WriteResponse(w, res)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	fmt.Println(ID)
	student, err := models.DeleteStudent(ID)
	WriteResponse(w, []byte(err.Error()))
	if res, err := json.Marshal(student); err != nil {
		WriteResponse(w, []byte(err.Error()))
	} else {

		WriteResponse(w, res)
	}
}

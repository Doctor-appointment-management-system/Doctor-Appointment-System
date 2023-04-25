package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Patients struct {
	ID                      int
	Name                    string
	Age                     int
	Gender                  string
	Address                 string
	City                    string
	Phone                   string
	Disease                 string
	Selected_specialisation string
	Patient_history         string
}

type Database interface {
	AddPatient(p *Patients) error
	GetPatient(p *Patients) error
	UpdatePatient(p *Patients) error
	DeletePatient(p *Patients) error
}

type HTTPHandler struct {
	db Database
}

type MySQLDatabase struct {
	db *sql.DB
}

func NewMySQLDatabase(connectionString string) (*MySQLDatabase, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLDatabase{db}, nil
}

// Add Patient to database-  CREATE OPERATION

func (m *MySQLDatabase) AddPatient(p *Patients) error {
	sql_query := fmt.Sprintf(`INSERT INTO patient(Name,Age,Gender,Address,City,Phone,Disease,Selected_Specialisation,Patient_history) VALUES('%s',%d,'%s','%s','%s','%s','%s','%s','%s')`, p.Name, p.Age, p.Gender, p.Address, p.City, p.Phone, p.Disease, p.Selected_specialisation, p.Patient_history)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) AddPatient(c *gin.Context) {
	var patient Patients
	if err := c.BindJSON(&patient); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.db.AddPatient(&patient); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, patient)
}

// Get the Patient details from the database - READ OPERATION

func (m *MySQLDatabase) GetPatient(p *Patients) error {
	sql_query := fmt.Sprintf(`SELECT * FROM Patient WHERE Phone='%s'`, p.Phone)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) GetPatient(c *gin.Context) {
	var patient Patients
	err := c.BindJSON(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.db.AddPatient(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, patient)
}

// Update the patient details in the database - UPDATE OPERATION

func (m *MySQLDatabase) UpdatePatient(p *Patients) error {
	update_query := fmt.Sprintf("UPDATE Patients SET Name='%s',Age=%d,Gender='%s',Address='%s',City='%s',Phone='%s', Diseases='%s',Selected_specialisation='%s',Patient_history='%s', WHERE Id=%d", p.Name, p.Age, p.Gender, p.Address, p.City, p.Phone, p.Disease, p.Selected_specialisation, p.Patient_history, p.ID)
	fmt.Println(update_query)
	_, err := m.db.Exec(update_query)
	return err
}

func (h *HTTPHandler) UpdatePatient(c *gin.Context) {
	var patient Patients
	err := c.BindJSON(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	err = h.db.UpdatePatient(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, patient)
}

// Delete the patient from Database - DELETE OPERATION

func (m *MySQLDatabase) DeletePatient(p *Patients) error {
	sql_query := fmt.Sprintf("DELETE FROM Patient WHERE Phone='%s'", p.Phone)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) DeletePatient(c *gin.Context) {
	var patient Patients
	err := c.BindJSON(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.db.DeletePatient(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, patient)
}
func main() {
	db, err := NewMySQLDatabase("root:india@123@tcp(localhost:3306)/das_new")
	if err != nil {
		log.Fatal(err)
	}
	defer db.db.Close()

	handler := &HTTPHandler{db}

	router := gin.Default()
	router.POST("patient/add_patients", handler.AddPatient)
	router.GET("patient/get_patient", handler.GetPatient)
	router.PUT("patient/update_patient", handler.UpdatePatient)
	router.DELETE("patient/delete_patient", handler.DeletePatient)
}

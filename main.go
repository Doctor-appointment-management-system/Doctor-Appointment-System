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

type Doctor struct {
	ID                               int
	Name                             string
	Gender                           string
	Address                          string
	City                             string
	Phone                            string
	Specialisation                   string
	Opening_time                     string
	Closing_time                     string
	Availability_time                string
	Availability                     string
	Available_for_home_visit         string
	Available_for_online_consultancy string
	Fees                             int
}

type dbase interface {
	AddPatient(p *Patients) error
	GetPatient(p *Patients) error
	UpdatePatient(p *Patients) error
	DeletePatient(p *Patients) error
	AddDoctor(p *Doctor) error
	GetDoctor(p *Doctor) error
	UpdateDoctort(p *Doctor) error
	DeleteDoctor(p *Doctor) error
}

type HTTPHandler struct {
	db dbase
}

type MySQLdbase struct {
	db *sql.DB
}

func NewMySQLdbase(connectionString string) (*MySQLdbase, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLdbase{db}, nil
}

// Add Doctor to dbase-  CREATE OPERATION
func (m *MySQLdbase) AddDoctor(d *Doctor) error {
	sql_query := fmt.Sprintf(`INSERT INTO Doctor (Name,Gender,Address,City,Phone,Specialisation,Opening_time,Closing_time,Availability_time,Availability,Available_for_home_visit,Available_for_online_consultancy,Fees) VALUES ( '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%d)`, d.Name, d.Gender, d.Address, d.City, d.Phone, d.Specialisation, d.Opening_time, d.Closing_time, d.Availability_time, d.Availability, d.Available_for_home_visit, d.Available_for_online_consultancy, d.Fees)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) AddDoctor(c *gin.Context) {
	var doctor Doctor
	if err := c.BindJSON(&doctor); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.db.AddDoctor(&doctor); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, doctor)
}

// Add Patient to dbase-  CREATE OPERATION

func (m *MySQLdbase) AddPatient(p *Patients) error {
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

func (m *MySQLdbase) GetDoctor(p *Doctor) error {
	sql_query := fmt.Sprintf(`SELECT * FROM Doctor WHERE Phone='%s'`, p.Phone)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) GetDoctor(c *gin.Context) {
	var doctor Doctor
	err := c.BindJSON(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.db.GetDoctor(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, doctor)
}

// Get the Patient details from the dbase - READ OPERATION

func (m *MySQLdbase) GetPatient(p *Patients) error {
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

func (m *MySQLdbase) UpdateDoctort(d *Doctor) error {
	update_query := fmt.Sprintf("UPDATE Doctor SET Address='%s',City='%s',Phone='%s', Opening_time ='%s',Closing_time='%s',Fees %d, WHERE Id=%d", d.Address, d.City, d.Phone, d.Opening_time, d.Closing_time, d.Fees, d.ID)
	fmt.Println(update_query)
	_, err := m.db.Exec(update_query)
	return err
}

func (h *HTTPHandler) UpdateDoctort(c *gin.Context) {
	var doctor Doctor
	err := c.BindJSON(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	err = h.db.UpdateDoctort(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, doctor)
}

// Update the patient details in the dbase - UPDATE OPERATION

func (m *MySQLdbase) UpdatePatient(p *Patients) error {
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

func (m *MySQLdbase) DeleteDoctor(d *Doctor) error {
	sql_query := fmt.Sprintf("DELETE FROM Doctor WHERE ID= %d", d.ID)
	_, err := m.db.Exec(sql_query)
	return err
}

func (h *HTTPHandler) DeleteDoctor(c *gin.Context) {
	var doctor Doctor
	err := c.BindJSON(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.db.DeleteDoctor(&doctor)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, doctor)
}

// Delete the patient from dbase - DELETE OPERATION

func (m *MySQLdbase) DeletePatient(p *Patients) error {
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

func Err(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func dbCreation() {

	//connecting to mysql

	db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/")
	Err(err)
	defer db.Close()

	// database creation

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS das_new")
	Err(err)
}
func db_connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das_new")

	if err != nil {
		return nil, err
	}
	return db, nil
}

func sql_tabel_creation() {
	db, err := db_connection()
	Err(err)
	// sql table creation

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Patient(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Age INT,Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Disease VARCHAR(25),Selected_Specialisation VARCHAR(20),Patient_history VARCHAR(250), PRIMARY KEY (ID) );")

	Err(err)
}

func main() {
	dbCreation()
	sql_tabel_creation()

	db, err := NewMySQLdbase("root:india@123@tcp(localhost:3306)/das_new")
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
	router.POST("doctor/add_doctor", handler.AddDoctor)
	router.GET("doctor/get_doctor", handler.GetDoctor)
	router.PUT("doctor/update_doctor", handler.UpdateDoctort)
	router.DELETE("doctor/delete_doctor", handler.DeleteDoctor)
	router.Run(":8080")
}

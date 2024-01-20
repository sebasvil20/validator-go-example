package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/sebasvil20/validator-go-example/src/models"
)

func main() {
	// open file invalid.json and unmarshal it to variable students
	/*jsonContent, err := os.ReadFile("./src/json_mocks/invalid.json")
	if err != nil {
		panic(err)
	}*/

	// open file valid.json and unmarshal it to variable students
	jsonContent, err := os.ReadFile("./src/json_mocks/valid.json")
	if err != nil {
		panic(err)
	}

	var student models.Student
	err = json.Unmarshal(jsonContent, &student)
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	validate.RegisterValidation("is-phone", validatePhoneNumber)
	err = validate.Struct(student)
	if err != nil {
		log.Println(err.Error())
	}
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumberPattern := `^\d{3}-\d{3}-\d{4}$`
	regx := regexp.MustCompile(phoneNumberPattern)
	return regx.MatchString(fl.Field().String())
}

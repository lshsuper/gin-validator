package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
)

func phoneStrategy(fl validator.FieldLevel) bool {

	if phone, ok := fl.Field().Interface().(string); ok {
		reg := `^(0|86|17951)?(13[0-9]|15[012356789]|16[6]|19[89]]|17[01345678]|18[0-9]|14[579])[0-9]{8}$`

		if matched, _ := regexp.MatchString(reg, phone); matched {
			return true
		}

	}
	return false
}

func mustStrategy(fl validator.FieldLevel) bool {
	if con, ok := fl.Field().Interface().(string); ok {
		if len(strings.Trim(con, " ")) > 0 {
			return true
		}
	}
	return false
}

func gtStrategy(fl validator.FieldLevel) bool {

	if n, ok := fl.Field().Interface().(int); ok {
		gtNum, _ := strconv.Atoi(fl.Param())
		return n > gtNum
	}

	if n, ok := fl.Field().Interface().(float64); ok {
		gtNum, _ := strconv.ParseFloat(fl.Param(), 10)
		return n > gtNum
	}
	return false
}

func ltStrategy(fl validator.FieldLevel) bool {

	if n, ok := fl.Field().Interface().(int); ok {
		gtNum, _ := strconv.Atoi(fl.Param())
		return n < gtNum
	}

	if n, ok := fl.Field().Interface().(float64); ok {
		gtNum, _ := strconv.ParseFloat(fl.Param(), 10)
		return n < gtNum
	}
	return false
}

func eqStrategy(fl validator.FieldLevel) bool {

	if n, ok := fl.Field().Interface().(int); ok {
		gtNum, _ := strconv.Atoi(fl.Param())
		return n == gtNum
	}

	if n, ok := fl.Field().Interface().(float64); ok {
		gtNum, _ := strconv.ParseFloat(fl.Param(), 10)
		return n == gtNum
	}
	return false
}

func gteStrategy(fl validator.FieldLevel) bool {

	if n, ok := fl.Field().Interface().(int); ok {
		gtNum, _ := strconv.Atoi(fl.Param())
		return n >= gtNum
	}

	if n, ok := fl.Field().Interface().(float64); ok {
		gtNum, _ := strconv.ParseFloat(fl.Param(), 10)
		return n >= gtNum
	}
	return false
}

func lteStrategy(fl validator.FieldLevel) bool {

	if n, ok := fl.Field().Interface().(int); ok {
		gtNum, _ := strconv.Atoi(fl.Param())
		return n <= gtNum
	}

	if n, ok := fl.Field().Interface().(float64); ok {
		gtNum, _ := strconv.ParseFloat(fl.Param(), 10)
		return n <= gtNum
	}
	return false
}
func mailStrategy(fl validator.FieldLevel) bool {
	if mail, ok := fl.Field().Interface().(string); ok {
		reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		if matched, _ := regexp.MatchString(reg, mail); matched {
			return true
		}
	}
	return false
}

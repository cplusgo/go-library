package go_library

import (
	"reflect"
	"fmt"
)

func IsStruct(data interface{}) bool  {
	dataType := reflect.TypeOf(data)
	fmt.Println(dataType.Kind())
	fmt.Println(reflect.Struct)
	return dataType.Kind() == reflect.Struct
}

func GetStructName(data interface{}) string  {
	dataType := reflect.TypeOf(data)
	return dataType.Name()
}

func HasMethod(data interface{}, methodName string) bool  {
	isExist := false
	dataType := reflect.TypeOf(data)
	numMethod := dataType.NumMethod()
	fmt.Println(numMethod)
	for i:=0; i<numMethod; i++ {
		method := dataType.Method(i)
		fmt.Println(method.Name)
		if method.Name == methodName {
			isExist = true
			break
		}
	}
	return isExist
}

func HasAttr(data interface{}, attrName string) bool {
	isExist := false
	dataType := reflect.TypeOf(data)
	numAttr := dataType.NumField()
	for i:=0; i<numAttr; i++ {
		field := dataType.Field(i)
		if field.Name == attrName {
			isExist = true
			break
		}
	}
	return isExist
}
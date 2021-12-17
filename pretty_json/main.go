package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	// 자료형 정의
	type MyType struct {
		Name     string
		Projects []string
	}
	// 값 설정
	value := MyType{Name: "mina", Projects: []string{"GopherJS", "ReactJS"}}
	// prettyjson 모듈 호출, 'var prettyjson = require("prettyjson")' 자바스크립트 코드와 같음
	prettyJSON := js.Global.Call("require", "prettyJSON")
	//prettyjson.rendor(value); 자바스크립트 코드와 같음
	result := prettyJSON.Call("rendor", value)

	/*
		code...
	*/
}

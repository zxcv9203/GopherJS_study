package main

import "syscall/js"

func main() {
	//console은 *js.Object형 변수
	console := js.Global.Get("console")
	// *js.Object에는 console의 메서드를 호출하는 Call 메서드가 있습니다.
	console.Call("log", "Hello GopherJS!")
}

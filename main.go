package main

import (
	"syscall/js"
)

var mGlobal jsGlobal

const modName = "urlv4"
const fnHello = "Hello"

func registerMain() {
	// Register Callback. Make sure modName is Set first
	// TODO how to make mutiple libraries ? Only one object can be efined why ? The O 2021.02.24
	// TeeJeeObject := make(map[string]interface{})
	// js.Global().Set("TeeJeeGo", js.ValueOf(TeeJeeObject))

	objModulemodName := make(map[string]interface{})
	js.Global().Set(modName, js.ValueOf(objModulemodName))
	js.Global().Get(modName).Set(fnHello, js.FuncOf(testHelloFn))
}

func main() {
	// Create a channel -> long-running pogram
	c := make(chan struct{}, 0)

	// Create reference to webpage via sysall js wasm model
	// Will be different in case multiple webpages are used
	// Reused in different get en set struct functions
	// Create empty hashGrid f type XmlDivGrid // The Grid will be ued in divgrid.go
	mGlobal = jsGlobal{
		Doc: js.Global().Get("document"),
		// DivGridMap: make(map[string]DivGrid),
	}

	// Info loading WASM file
	s3("main.go", modName, " init  is done")

	// Register functions
	registerMain()

	// divgridInitTest()

	<-c
}

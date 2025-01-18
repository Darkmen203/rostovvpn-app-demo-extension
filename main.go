package main

import (
	_ "github.com/Darkmen203/rostovvpn-app-demo-extension/hiddify_extension"

	"github.com/Darkmen203/rostovvpn-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}

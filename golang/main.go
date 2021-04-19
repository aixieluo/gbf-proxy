package main

import "gbf-proxy/golang/cmd"

var version = "latest"

func main() {
	cmd.Version = version
	cmd.Execute()
}

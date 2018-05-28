package main

import (
	"fmt"
	"os/exec"
)

func main() {
	f, err := exec.Command("/bin/bash", "-c", "cat /proc/535/status").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}

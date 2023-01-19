package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 执行一个cmd，获取输出结果
	args := os.Args
	if len(args) > 1 {
		filename := strings.Join(args[1:], "")
		cmd := exec.Command("libcamera-jpeg", "-o", filename)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(out))
	} else {
		cmd := exec.Command("libcamera-jpeg", "-o", "save.jpg")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(out))
	}
}

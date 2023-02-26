//通过runtime包获取所在的操作系统类型

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main(){
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n",path)
}

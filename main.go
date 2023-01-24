package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "youtube":
		query := fmt.Sprintf("https://www.youtube.com/results?search_query=%v", os.Args[2])
		openBrowser(query)
	}

}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], strings.Replace(url, "&", "^&", -1))...)
	return cmd.Start() == nil
}

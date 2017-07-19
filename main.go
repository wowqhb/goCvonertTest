package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

var (
	exe = "D:/idocv/converter/word2html.exe"
)

func main() {
	_i := flag.Int("t", 1, "times")
	flag.Parse()
	i := *(_i)
	for j := 0; j < i; j++ {
		go func(j int) {
			bs, err := exec.Command(exe, fmt.Sprintf("D:/src/%d.docx", j), fmt.Sprintf("D:/tmp/%d.html", j)).Output()
			fmt.Println(fmt.Sprint(j, "/", i, " - ", "result:", err, "; ", "time:", string(bs)))
		}(j)
	}

	time.Sleep(24 * time.Hour)
}

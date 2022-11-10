package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, dir := range pathSplit {
		fpath := filepath.Join(dir, file)
		finfo, err := os.Stat(fpath)

		if err == nil {
			mode := finfo.Mode()

			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fpath)
					return
				}
			}
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func SaveData(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

func randomInt() int {
	return rand.Intn(10000)
}

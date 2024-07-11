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

	err = fp.Sync()
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
}

func LogAppend(fp *os.File, line string) error {
	buf := []byte(line)
	buf = append(buf, '\n')
	_, err := fp.Write(buf)
	if err != nil {
		return err
	}
	return fp.Sync()
}

func randomInt() int {
	return rand.Intn(10000)
}

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Hello World")
	SaveData("test", []byte{1, 2, 3})
}

func SaveData(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Intn(100))
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	//fsync (flush the data into the disk)
	err = fp.Sync()
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

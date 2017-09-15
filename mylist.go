package envi

import "os"

func Mylist() []string {
	myenv := os.Environ()
	return myenv
}

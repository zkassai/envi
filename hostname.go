package envi

import "os"

func Hostname() string {
	myhostname,_:=os.Hostname()
	return myhostname
}

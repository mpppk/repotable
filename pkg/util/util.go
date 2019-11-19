// Package util provides some utilities
package util

func PanicIfErrExist(err error) {
	if err != nil {
		panic(err)
	}
}

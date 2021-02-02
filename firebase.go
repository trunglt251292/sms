package sms

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// LogText
func LogText(text string) {
	fmt.Println(aurora.Green("*** Text : " + text))
}

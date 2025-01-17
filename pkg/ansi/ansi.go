package ansi

import (
	"regexp"
)

const (
	ansiRegexStr = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
)

var (
	re = regexp.MustCompile(ansiRegexStr)
)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

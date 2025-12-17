package botgolang

import (
	"net/url"
	"regexp"
)

const loggngTokenRepl = "token=$1.**********.$3"

var loggngTokenRegexp = regexp.MustCompile(`token=([0-9]{3})\.([0-9]{10})\.([0-9]{10})`)

func loggingHideUrlToken(u *url.URL) string {

	return loggngTokenRegexp.ReplaceAllString(u.String(), loggngTokenRepl)
}

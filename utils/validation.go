package utils

import (
	"log"
	"regexp"
	"strings"
)

func IsOnion(identifier string) bool {
	// TODO: At some point we will want to support i2p
	log.Println("NEW VALIDATOR LOGIC USED")
	if len(identifier) >= 22 && strings.HasSuffix(identifier, ".onion") {
		isv2url, _ := regexp.MatchString(`(^|\.)[a-z2-7]{16}\.onion$`, identifier)
		isv3url, _ := regexp.MatchString(`(^|\.)[a-z2-7]{56}\.onion$`, identifier)
		return isv2url || isv3url
	}
	return false
}

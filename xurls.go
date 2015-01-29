/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package xurls

import (
	"regexp"
)

func reverseJoin(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[len(a)-1])
	for i := len(a)-2; i >= 0; i-- {
		s := a[i]
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func reverseCombineOr(regexes []string) string {
	return `(` + reverseJoin(regexes, `|`) + `)`
}

var (
	letters = "a-zA-Z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF"
	iriChar = letters + `0-9`
	ipv4Addr = `((25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[0-9]))`
	ipv6Addr = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	ipAddr = `(` + ipv4Addr + `|` + ipv6Addr + `)`
	iri = `[` + iriChar + `]([` + iriChar + `\-]{0,61}[` + iriChar + `]){0,1}`
	gtld = `(?i)` + reverseCombineOr(TLDs) + `(?-i)`
	hostName = `(` + iri + `\.)+` + gtld
	domainName = `(` + hostName + `|` + ipAddr + `)`
	webUrl = `((https?:\/\/(([a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(\%[a-fA-F0-9]{2})){1,64}(\:([a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(\%[a-fA-F0-9]{2})){1,25})?\@)?)?(` + domainName + `)(\:\d{1,5})?)(\/(([` + iriChar + `\;\/\?\:\@\&\=\#\~\-\.\+\!\*\'\(\)\,\_])|(\%[a-fA-F0-9]{2}))*)?(\b|$)`
	emailAddr = `(mailto:)?[a-zA-Z0-9\.\_\%\-\+]{1,256}\@` + domainName
	all = `(` + webUrl + `|` + emailAddr + `)`

	WebUrl = regexp.MustCompile(webUrl)
	EmailAddr = regexp.MustCompile(emailAddr)
	All = regexp.MustCompile(all)
)
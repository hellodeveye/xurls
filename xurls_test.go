/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package xurls

import (
	"testing"
)

func TestWebURL(t *testing.T) {
	for _, c := range [...]struct {
		in   string
		want string
	}{
		{``, ``},
		{`foo`, ``},
		{`foo.a`, ``},
		{`foo.random`, ``},
		{`foo.com`, `foo.com`},
		{`foo.com bar.com`, `foo.com`},
		{`foo.onion`, `foo.onion`},
		{`foo.i2p`, `foo.i2p`},
		{`中国.中国`, `中国.中国`},
		{`foo.com/`, `foo.com/`},
		{`1.1.1.1`, `1.1.1.1`},
		{`121.1.1.1`, `121.1.1.1`},
		{`255.1.1.1`, `255.1.1.1`},
		{`300.1.1.1`, ``},
		{`1.1.1`, ``},
		{`1.1..1`, ``},
		{`1080:0:0:0:8:800:200C:4171`, `1080:0:0:0:8:800:200C:4171`},
		{`3ffe:2a00:100:7031::1`, `3ffe:2a00:100:7031::1`},
		{`1080::8:800:200c:417a`, `1080::8:800:200c:417a`},
		{`1:1`, ``},
		{`:2:`, ``},
		{`1:2:3`, ``},
		{`test.foo.com`, `test.foo.com`},
		{`test.foo.com/path`, `test.foo.com/path`},
		{`test.foo.com/path/more/`, `test.foo.com/path/more/`},
		{`TEST.FOO.COM/PATH`, `TEST.FOO.COM/PATH`},
		{`foo.com/a.,:;-+_()?@&=$~!*%'"a`, `foo.com/a.,:;-+_()?@&=$~!*%'"a`},
		//{`foo.com/path_(more)`, `foo.com/path_(more)`},
		{`http://foo.com`, `http://foo.com`},
		{` http://foo.com `, `http://foo.com`},
		{`,http://foo.com,`, `http://foo.com`},
		{`(http://foo.com)`, `http://foo.com`},
		{`<http://foo.com>`, `http://foo.com`},
		{`"http://foo.com"`, `http://foo.com`},
		{`http://foo.com`, `http://foo.com`},
		{`http://test.foo.com/`, `http://test.foo.com/`},
		{`http://foo.com/path`, `http://foo.com/path`},
		{`http://foo.com:8080/path`, `http://foo.com:8080/path`},
		{`http://1.1.1.1/path`, `http://1.1.1.1/path`},
		{`www.foo.com`, `www.foo.com`},
		{` foo.com/bar `, `foo.com/bar`},
		{`<foo.com/bar>`, `foo.com/bar`},
		{`,foo.com/bar,`, `foo.com/bar`},
		{`,foo.com/bar,more`, `foo.com/bar,more`},
		{`(foo.com/bar`, `foo.com/bar`},
		{`(foo.com/bar)more`, `foo.com/bar)more`},
		{`"foo.com/bar"`, `foo.com/bar`},
		{`"foo.com/bar"more`, `foo.com/bar"more`},
	} {
		got := WebURL.FindString(c.in)
		if got != c.want {
			t.Errorf(`WebURL.FindString("%s") got "%s", want "%s"`, c.in, got, c.want)
		}
	}
}

func TestEmail(t *testing.T) {
	for _, c := range [...]struct {
		in   string
		want string
	}{
		{``, ``},
		{`foo`, ``},
		{`foo@bar`, ``},
		{`foo@bar.a`, ``},
		{`foo@bar.com`, `foo@bar.com`},
		{`foo@bar.com bar@bar.com`, `foo@bar.com`},
		{`foo@bar.onion`, `foo@bar.onion`},
		{`foo@中国.中国`, `foo@中国.中国`},
		{`mailto:foo@bar.com`, `foo@bar.com`},
		{`foo@test.bar.com`, `foo@test.bar.com`},
		{`FOO@TEST.BAR.COM`, `FOO@TEST.BAR.COM`},
		{`foo@bar.com/path`, `foo@bar.com`},
		{`foo+test@bar.com`, `foo+test@bar.com`},
		{`foo+._%-@bar.com`, `foo+._%-@bar.com`},
	} {
		got := Email.FindString(c.in)
		if got != c.want {
			t.Errorf(`EmailAddr.FindString("%s") got "%s", want "%s"`, c.in, got, c.want)
		}
	}
}

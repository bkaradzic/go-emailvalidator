/*
 * Copyright 2014 Branimir Karadzic. All rights reserved.
 * https://github.com/bkaradzic/go-emailvalidator/blob/master/LICENSE
 */

package emailvalidator

import (
	"testing"
)

func validEmail(t *testing.T, email string) {

	result := IsValid(email)

	if !result {
		t.Error("Should be valid ", email)
	}
}

func invalidEmail(t *testing.T, email string) {

	result := IsValid(email)

	if result {
		t.Error("Should be invalid ", email)
	}
}

func disposableEmail(t *testing.T, disposable bool, email string) {

	result := IsDisposable(email)

	if disposable {
		if !result {
			t.Error("Should be disposable ", email)
		}
	} else {
		if result {
			t.Error("Should be non-disposable ", email)
		}
	}
}

func TestEmail(t *testing.T) {

	validEmail(t, `dclo@us.ibm.com`)
	validEmail(t, `abc\@def@example.com`)
	validEmail(t, `abc\\@example.com`)
	validEmail(t, `Fred\ Bloggs@example.com`)
	validEmail(t, `Joe.\\Blow@example.com`)
	validEmail(t, `"Abc@def"@example.com`)
	validEmail(t, `"Fred Bloggs"@example.com`)
	validEmail(t, `customer/department=shipping@example.com`)
	validEmail(t, `$A12345@example.com`)
	validEmail(t, `!def!xyz%abc@example.com`)
	validEmail(t, `_somename@example.com`)
	validEmail(t, `user+mailbox@example.com`)
	validEmail(t, `peter.piper@example.com`)
	validEmail(t, `Doug\ \"Ace\"\ Lovell@example.com`)
	validEmail(t, `"Doug \"Ace\" L."@example.com`)
	validEmail(t, `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@letters-in-local.org`)
	validEmail(t, `01234567890@numbers-in-local.net`)
	validEmail(t, `&'*+-./=?^_{}~@other-valid-characters-in-local.net`)
	validEmail(t, `mixed-1234-in-{+^}-local@sld.net`)
	validEmail(t, `a@single-character-in-local.org`)
	validEmail(t, `"quoted"@sld.com`)
	validEmail(t, `"\e\s\c\a\p\e\d"@sld.com`)
	validEmail(t, `"quoted-at-sign@sld.org"@sld.com`)
	validEmail(t, `"escaped\"quote"@sld.com`)
	validEmail(t, `"back\slash"@sld.com`)
	validEmail(t, `single-character-in-sld@x.org`)
	validEmail(t, `local@dash-in-sld.com`)
	validEmail(t, `letters-in-sld@123.com`)
	validEmail(t, `uncommon-tld@sld.museum`)
	validEmail(t, `uncommon-tld@sld.travel`)
	validEmail(t, `uncommon-tld@sld.mobi`)
	validEmail(t, `country-code-tld@sld.uk`)
	validEmail(t, `country-code-tld@sld.rw`)
	validEmail(t, `local@sld.newTLD`)
	validEmail(t, `punycode-numbers-in-tld@sld.xn--3e0b707e`)
	validEmail(t, `local@sub.domains.com`)
	validEmail(t, `ichiro.yamamoto@xn--wgv71a.jp`)
	validEmail(t, `山本一呂@日本.jp`)
	validEmail(t, `山本一呂@xn--wgv71a.jp`)
//	validEmail(t, `bracketed-IP-instead-of-domain@[127.0.0.1]`) // not supported yet...

	invalidEmail(t, `abc@def@example.com`)
	invalidEmail(t, `abc\\@def@example.com`)
	invalidEmail(t, `abc\@example.com`)
	invalidEmail(t, `@example.com`)
	invalidEmail(t, `doug@`)
	invalidEmail(t, `"qu@example.com`)
	invalidEmail(t, `ote"@example.com`)
	invalidEmail(t, `.dot@example.com`)
	invalidEmail(t, `dot.@example.com`)
	invalidEmail(t, `two..dot@example.com`)
	invalidEmail(t, `"Doug "Ace" L."@example.com`)
	invalidEmail(t, `Doug\ \"Ace\"\ L\.@example.com`)
	invalidEmail(t, `hello world@example.com`)
	invalidEmail(t, `gatsby@f.sc.ot.t.f.i.tzg.era.l.d.`)

	disposableEmail(t, true, `test@mailInator.com`)
	disposableEmail(t, true, `test@zA.com`)
	disposableEmail(t, false, `test@0za.com`)
}


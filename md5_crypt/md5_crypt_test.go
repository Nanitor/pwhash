// (C) Copyright 2012, Jeramey Crawford <jeramey@antihe.ro>. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package md5_crypt

import "testing"

type testData struct {
	salt, key, result string
}

func TestCrypt(t *testing.T) {
	data := []testData{
		{
			"$1$$", "abcdefghijk",
			"$1$$pL/BYSxMXs.jVuSV1lynn1",
		},
		{
			"$1$an overlong salt$", "abcdfgh",
			"$1$an overl$ZYftmJDIw8sG5s4gG6r.70",
		},
		{
			"$1$12345678$", "Lorem ipsum dolor sit amet",
			"$1$12345678$Suzx8CrBlkNJwVHHHv5tZ.",
		},
		{
			"$1$deadbeef$", "password",
			"$1$deadbeef$Q7g0UO4hRC0mgQUQ/qkjZ0",
		},
		{
			"$1$$", "missing salt",
			"$1$$Lv61fbMiEGprscPkdE9Iw/",
		},
		{
			"$1$holy-moly-batman$", "1234567",
			"$1$holy-mol$WKomB0dWknSxdW/e8WYHG0",
		},
		{
			"$1$asdfjkl;$", "A really long password. Longer " +
				"than a password has any right to be" +
				". Hey bub, don't mess with this password.",
			"$1$asdfjkl;$DUqPhKwbK4smV0aEMyDdx/",
		},
	}

	for i, d := range data {
		hash := Crypt(d.key, d.salt)
		if hash != d.result {
			t.Errorf("Test %d failed\nExpected: %s\n     Saw: %s",
				i, d.result, hash)
		}
	}
}

func TestVerify(t *testing.T) {
	data := []string{
		"password",
		"12345",
		"That's amazing! I've got the same combination on my luggage!",
		"And change the combination on my luggage!",
		"         random  spa  c    ing.",
		"94ajflkvjzpe8u3&*j1k513KLJ&*()",
	}
	for i, d := range data {
		hash := Crypt(d, "")
		if !Verify(d, hash) {
			t.Errorf("Test %d failed: %s", i, d)
		}
	}
}

func TestGenerateSalt(t *testing.T) {
	salt := GenerateSalt(0)
	if len(salt) != len(MagicPrefix)+1 {
		t.Errorf("Expected len 1, saw len %d", len(salt))
	}

	for i := 1; i <= 8; i++ {
		salt = GenerateSalt(i)
		if len(salt) != len(MagicPrefix)+i {
			t.Errorf("Expected len %d, saw len %d", i, len(salt))
		}
	}

	salt = GenerateSalt(9)
	if len(salt) != len(MagicPrefix)+8 {
		t.Errorf("Expected len 8, saw len %d", len(salt))
	}
}

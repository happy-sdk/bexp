package vars

import "testing"

func TestNewValue(t *testing.T) {
	var tests = []struct {
		val  interface{}
		want string
	}{
		{nil, ""},
		{"", ""},
	}

	for _, tt := range tests {
		got := NewValue(tt.val).String()
		if got != tt.want {
			t.Errorf("want: %s got %s", tt.want, got)
		}
	}
}

func TestValueFromString(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want string
	}{
		{"STRING", "some-string", "some-string"},
		{"STRING", "some-string with space ", "some-string with space"},
		{"STRING", " some-string with space", "some-string with space"},
		{"STRING", "1234567", "1234567"},
	}
	for _, tt := range tests {
		if got := NewValue(tt.val); got.String() != tt.want {
			t.Errorf("ValueFromString() = %q, want %q", got.String(), tt.want)
		}
		if rv := NewValue(tt.val); string(rv.Rune()) != tt.want {
			t.Errorf("Value.Rune() = %q, want %q", string(rv.Rune()), tt.want)
		}
	}
}

func TestValueParseInt64(t *testing.T) {
	val := Value("200")
	iout, erri1 := val.AsInt()
	if iout != 200 {
		t.Errorf("Value(11).AsInt() = %d, err(%v) want 200", iout, erri1)
	}

	val2 := Value("x")
	iout2, erri2 := val2.AsInt()
	if iout2 != 0 || erri2 == nil {
		t.Errorf("Value(11).AsInt() = %d, err(%v) want 0 and err", iout2, erri2)
	}
}

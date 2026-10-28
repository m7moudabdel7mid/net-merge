package engine

import "testing"

func TestDisplayName(t *testing.T) {
	for _, c := range []struct{ name, label, want string }{
		{"en0", "Wi-Fi", "Wi-Fi (en0)"},
		{"en5", "", "en5"},
		{"Ethernet 2", "Ethernet 2", "Ethernet 2"},
	} {
		if got := (SlotStats{Name: c.name, Label: c.label}).DisplayName(); got != c.want {
			t.Errorf("DisplayName(%q, %q) = %q, want %q", c.name, c.label, got, c.want)
		}
	}
}

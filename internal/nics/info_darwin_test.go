package nics

import "testing"

func TestParseServices(t *testing.T) {
	out := `An asterisk (*) denotes that a network service is disabled.
(1) LAN
(Hardware Port: USB 10/100/1000 LAN, Device: en5)

(2) Wi-Fi
(Hardware Port: Wi-Fi, Device: en0)

(*) Wi-Fi copy
(Hardware Port: Wi-Fi, Device: en0)

(3) Thunderbolt Bridge
(Hardware Port: Thunderbolt Bridge, Device: bridge0)
`
	got := parseServices(out)
	for dev, want := range map[string]string{"en5": "LAN", "en0": "Wi-Fi", "bridge0": "Thunderbolt Bridge"} {
		if got[dev] != want {
			t.Errorf("service for %s = %q, want %q", dev, got[dev], want)
		}
	}
}

func TestParsePorts(t *testing.T) {
	got := parsePorts("Hardware Port: Wi-Fi\nDevice: en0\nEthernet Address: aa\n\nHardware Port: USB 10/100/1000 LAN\nDevice: en5\n")
	if got["en0"] != "Wi-Fi" || got["en5"] != "USB 10/100/1000 LAN" {
		t.Fatalf("parsePorts = %v", got)
	}
}

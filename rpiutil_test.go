package rpiutil_test

import (
	"rpiutil"
	"testing"
)

func baseTest(t *testing.T, data string, revExpected int) {
	rev := rpiutil.GetPCBRevisionFrom(data)
	if rev != revExpected {
		t.Fatalf("GetRevisionFrom(): expected %d, got %d", revExpected, rev)
	}
}

func TestMyPi(t *testing.T) {
	revExpected := 1
	data := `processor	: 0
model name	: ARMv6-compatible processor rev 7 (v6l)
BogoMIPS	: 2.00
Features	: swp half thumb fastmult vfp edsp java tls
CPU implementer	: 0x41
CPU architecture	: 7
CPU variant	: 0x0
CPU part	: 0xb76
CPU revision	: 7

Hardware	: BCM2708
Revision	: 0002
Serial	: 00000000fea2fbf5`

	baseTest(t, data, revExpected)
}

func Test0002(t *testing.T) {
	revExpected := 1
	data := `foofoofoofoo
Hardware	: BCM2708
Revision	: 0002
barbarbarbar`

	baseTest(t, data, revExpected)
}

func Test1000002(t *testing.T) {
	revExpected := 1
	data := `foofoofoofoo
Hardware	: BCM2708
Revision	: 1000002
barbarbarbar`

	baseTest(t, data, revExpected)
}

func Test0003(t *testing.T) {
	revExpected := 1
	data := `foofoofoofoo
Hardware	: BCM2708
Revision	: 0003
barbarbarbar`

	baseTest(t, data, revExpected)
}

func Test1000003(t *testing.T) {
	revExpected := 1
	data := `foofoofoofoo
Hardware	: BCM2708
Revision	: 1000003
barbarbarbar`

	baseTest(t, data, revExpected)
}

func Test1234(t *testing.T) {
	revExpected := 2
	data := `foofoofoofoo
Hardware	: BCM2708
Revision	: 1234
barbarbarbar`

	baseTest(t, data, revExpected)
}

func TestWrongHardware(t *testing.T) {
	revExpected := -1
	data := `foofoofoofoo
Hardware	: WRONGHARDWARE
Revision	: 0001
barbarbarbar`

	baseTest(t, data, revExpected)
}

func TestEmptyString(t *testing.T) {
	revExpected := -1
	data := ``

	baseTest(t, data, revExpected)
}

func TestBadField(t *testing.T) {
	revExpected := -1
	data := `foofoofoofoo
Hardware
Revision	: 0001
barbarbarbar`

	baseTest(t, data, revExpected)
}

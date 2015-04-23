package tcx

import "testing"

var tcx *Tcx

func TestParse(t *testing.T) {
	var err error
	tcx, err = ParseFile("testdata/test1.tcx")

	if err != nil {
		t.Error("Error parsing TCX file: ", err)
	}
}

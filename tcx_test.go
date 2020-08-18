package tcx

import (
	"fmt"

	"testing"
)

var tcx *Tcx

func TestParse(t *testing.T) {
	var err error
	tcx, err = ParseFile("testdata/test1.tcx")

	if err != nil {
		t.Error("Error parsing TCX file: ", err)
	}

	fmt.Println(tcx.Activities[0].TotalDuration())
	fmt.Println(tcx.Activities[0].AverageHeartbeat())
	fmt.Println(tcx.Activities[0].AveragePace())
}

func TestFitbit_DistanceIsParsed(t *testing.T) {
	var err error
	tcx, err = ParseFile("testdata/fitbit-test.tcx")

	if err != nil {
		t.Error("Error parsing TCX file: ", err)
	}

	fmt.Println(tcx.Activities[0].TotalDuration())
	fmt.Println(tcx.Activities[0].AverageHeartbeat())
	fmt.Println(tcx.Activities[0].AveragePace())

	if tcx.Activities[0].Laps[0].Track[2].DistanceInMeters == 0 {
		t.Fail()
	}

}

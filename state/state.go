package state

import "fmt"

var HsState = [4]string{
	"[V---\\[kylling rev korn HS]\\ \\___________________________/ ________________________/---Ø]",
	"[V---\\____________________\\ \\__[HS rev korn og kylling]__/ ________________________/---Ø]",
	"[V---\\_____________________\\ \\___________________________/ [HS rev korn og kylling]/---Ø]"}

var Status = HsState[0]

func MoveTo(MoveTo string) string {

	if MoveTo == "Restart" || MoveTo == "HS og kylling til Vest" {
		Status = HsState[0]
		fmt.Println(HsState[0] + "Init State")
		return HsState[0]
	}

	if MoveTo == "HS rev og kylling til båt" && (Status == HsState[0] || Status == HsState[2]) {
		Status = HsState[1]
		fmt.Println(HsState[1])
		return HsState[1]
	}

	if MoveTo == "HS rev og kylling til Øst" && (Status == HsState[1] || Status == HsState[3]) {
		Status = HsState[2]
		fmt.Println(HsState[2])
		return HsState[2]
	}

	return ""
}

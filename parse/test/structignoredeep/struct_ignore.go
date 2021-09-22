package buildtags

import (
	"fmt"
)

// genny ignore struct
type _t_ struct {
	Number int
}

func _t_Print(t _t_) {
	if t.Number > 0 {
		fmt.Println(t)
	}
}

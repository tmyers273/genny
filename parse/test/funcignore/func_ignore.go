package buildtags

import (
	"fmt"
)

type _t_ struct {
	Number int
}

func _t_Print(t _t_) {
	fmt.Println(t)
}

// genny ignore
func _t_Ignored(t _t_) {
	fmt.Println("should be ignored")
}
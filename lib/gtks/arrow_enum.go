package gtks

// ENUM_NAME = GTK_ARROW_UP, 0
// ENUM_NAME = GTK_ARROW_DOWN, 1
// ENUM_NAME = GTK_ARROW_LEFT, 2
// ENUM_NAME = GTK_ARROW_RIGHT, 3
// ENUM_NAME = GTK_ARROW_NONE, 4
type ArrowType int

const (
	ARROW_UP ArrowType = iota
	ARROW_DOWN
	ARROW_LEFT
	ARROW_RIGHT
	ARROW_NONE
)

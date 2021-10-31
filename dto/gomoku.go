package dto

type Client struct {
	Action string
	User   string
	Gomoku GomokuData
}

type GomokuData struct {
	X       int
	Y       int
	IsBlack bool
	IsWhite bool
}

type GomokuBoard struct {
	Rows    []bool
	Columns []bool
}

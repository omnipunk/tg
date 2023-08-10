package tx

import (
	"errors"
)

var (
	ScreenNotExistErr   = errors.New("screen does not exist")
	SessionNotExistErr  = errors.New("session does not exist")
	KeyboardNotExistErr = errors.New("keyboard does not exist")
)

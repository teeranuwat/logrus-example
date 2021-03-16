package logsrus

type ErrorString struct {
	s string
}

// NewError for Create Customer Error message
func NewError(text string) error {
	return &ErrorString{text}
}

func (e *ErrorString) Error() string {
	return e.s
}

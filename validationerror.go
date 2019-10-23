package govalidare

// ValidationError The structure for validation errors
type ValidationError struct {
	field string
	value interface{}
	rule  string
}

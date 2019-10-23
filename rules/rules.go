package rules

const (
	// REQUIRED rule for a value
	REQUIRED string = "required"
)

// CallRule Calls the specified
func CallRule(name string, value interface{}) bool {
	switch name {
	case REQUIRED:
		return required(value)
	}
	return true
}

func required(value interface{}) bool {
	if value == nil {
		return false
	}

	return true
}

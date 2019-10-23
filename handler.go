package govalidare

// Handler handles the validations for the given struct
type Handler struct {
	bind             Bind
	rules            map[string][]Rule
	validationErrors []ValidationError
	valid            bool
}

// New Creates a new validation handler
func New(b Bind) *Handler {
	return &Handler{bind: b}
}

// AddRuleForField Adds a rule for a given field name
func (h *Handler) AddRuleForField(f string, r Rule) {
	h.rules[f] = append(h.rules[f], r)
}

// Validate Validates the binded struct
func (h *Handler) Validate() bool {

}

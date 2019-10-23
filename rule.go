package govalidare

import (
	"govalidare/rules"
)

var defaultRules = [...]string{
	rules.REQUIRED,
}

type ruleFunc func(interface{}) bool
type callbackFunc func(bool)

// A Rule to validate
type Rule struct {
	rule         string
	call         ruleFunc
	compareValue interface{}
	callback     callbackFunc
}

// CallValidationRule Calls the specified validation rule
func (r *Rule) CallValidationRule(value interface{}) bool {
	if r.isDefaultRule() {
		return rules.CallRule(r.rule, value)
	}

	return r.call(value)
}

func (r *Rule) isDefaultRule() bool {
	for _, d := range defaultRules {
		if d == r.rule {
			return true
		}
	}
	return false
}

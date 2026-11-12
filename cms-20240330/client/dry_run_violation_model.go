// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunViolation interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *DryRunViolation
	GetField() *string
	SetMessage(v string) *DryRunViolation
	GetMessage() *string
	SetSuggestion(v string) *DryRunViolation
	GetSuggestion() *string
}

type DryRunViolation struct {
	// The path of the violated field, such as schedule.intervalSecs.
	//
	// example:
	//
	// schedule.intervalSecs
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The violation description.
	//
	// example:
	//
	// schedule interval must be a multiple of 15 seconds
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The remediation suggestion.
	//
	// example:
	//
	// must satisfy: 15-86400 seconds and a multiple of 15
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s DryRunViolation) String() string {
	return dara.Prettify(s)
}

func (s DryRunViolation) GoString() string {
	return s.String()
}

func (s *DryRunViolation) GetField() *string {
	return s.Field
}

func (s *DryRunViolation) GetMessage() *string {
	return s.Message
}

func (s *DryRunViolation) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DryRunViolation) SetField(v string) *DryRunViolation {
	s.Field = &v
	return s
}

func (s *DryRunViolation) SetMessage(v string) *DryRunViolation {
	s.Message = &v
	return s
}

func (s *DryRunViolation) SetSuggestion(v string) *DryRunViolation {
	s.Suggestion = &v
	return s
}

func (s *DryRunViolation) Validate() error {
	return dara.Validate(s)
}

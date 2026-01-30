// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidationOptions interface {
	dara.Model
	String() string
	GoString() string
	SetSkipVerifyAIChatCompletion(v bool) *ValidationOptions
	GetSkipVerifyAIChatCompletion() *bool
}

type ValidationOptions struct {
	SkipVerifyAIChatCompletion *bool `json:"skipVerifyAIChatCompletion,omitempty" xml:"skipVerifyAIChatCompletion,omitempty"`
}

func (s ValidationOptions) String() string {
	return dara.Prettify(s)
}

func (s ValidationOptions) GoString() string {
	return s.String()
}

func (s *ValidationOptions) GetSkipVerifyAIChatCompletion() *bool {
	return s.SkipVerifyAIChatCompletion
}

func (s *ValidationOptions) SetSkipVerifyAIChatCompletion(v bool) *ValidationOptions {
	s.SkipVerifyAIChatCompletion = &v
	return s
}

func (s *ValidationOptions) Validate() error {
	return dara.Validate(s)
}

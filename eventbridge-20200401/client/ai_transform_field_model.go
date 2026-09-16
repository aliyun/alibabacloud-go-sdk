// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTransformField interface {
	dara.Model
	String() string
	GoString() string
	SetForm(v string) *AiTransformField
	GetForm() *string
	SetValue(v string) *AiTransformField
	GetValue() *string
}

type AiTransformField struct {
	// The value form. Currently uses JSONPATH.
	//
	// example:
	//
	// JSONPATH
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The JSONPath expression.
	//
	// example:
	//
	// $.data.message
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AiTransformField) String() string {
	return dara.Prettify(s)
}

func (s AiTransformField) GoString() string {
	return s.String()
}

func (s *AiTransformField) GetForm() *string {
	return s.Form
}

func (s *AiTransformField) GetValue() *string {
	return s.Value
}

func (s *AiTransformField) SetForm(v string) *AiTransformField {
	s.Form = &v
	return s
}

func (s *AiTransformField) SetValue(v string) *AiTransformField {
	s.Value = &v
	return s
}

func (s *AiTransformField) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMatcher interface {
	dara.Model
	String() string
	GoString() string
	SetLabel(v string) *Matcher
	GetLabel() *string
	SetOperator(v string) *Matcher
	GetOperator() *string
	SetValue(v string) *Matcher
	GetValue() *string
}

type Matcher struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Matcher) String() string {
	return dara.Prettify(s)
}

func (s Matcher) GoString() string {
	return s.String()
}

func (s *Matcher) GetLabel() *string {
	return s.Label
}

func (s *Matcher) GetOperator() *string {
	return s.Operator
}

func (s *Matcher) GetValue() *string {
	return s.Value
}

func (s *Matcher) SetLabel(v string) *Matcher {
	s.Label = &v
	return s
}

func (s *Matcher) SetOperator(v string) *Matcher {
	s.Operator = &v
	return s
}

func (s *Matcher) SetValue(v string) *Matcher {
	s.Value = &v
	return s
}

func (s *Matcher) Validate() error {
	return dara.Validate(s)
}

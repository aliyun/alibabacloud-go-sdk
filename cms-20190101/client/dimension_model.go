// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDimension interface {
	dara.Model
	String() string
	GoString() string
	SetLabel(v string) *Dimension
	GetLabel() *string
	SetValue(v string) *Dimension
	GetValue() *string
}

type Dimension struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Dimension) String() string {
	return dara.Prettify(s)
}

func (s Dimension) GoString() string {
	return s.String()
}

func (s *Dimension) GetLabel() *string {
	return s.Label
}

func (s *Dimension) GetValue() *string {
	return s.Value
}

func (s *Dimension) SetLabel(v string) *Dimension {
	s.Label = &v
	return s
}

func (s *Dimension) SetValue(v string) *Dimension {
	s.Value = &v
	return s
}

func (s *Dimension) Validate() error {
	return dara.Validate(s)
}

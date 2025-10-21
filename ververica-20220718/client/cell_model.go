// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCell interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *Cell
	GetValue() *string
}

type Cell struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Cell) String() string {
	return dara.Prettify(s)
}

func (s Cell) GoString() string {
	return s.String()
}

func (s *Cell) GetValue() *string {
	return s.Value
}

func (s *Cell) SetValue(v string) *Cell {
	s.Value = &v
	return s
}

func (s *Cell) Validate() error {
	return dara.Validate(s)
}

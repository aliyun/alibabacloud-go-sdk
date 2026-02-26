// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDivisionQuery interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v string) *DivisionQuery
	GetDivisionCode() *string
}

type DivisionQuery struct {
	// divisionCode
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
}

func (s DivisionQuery) String() string {
	return dara.Prettify(s)
}

func (s DivisionQuery) GoString() string {
	return s.String()
}

func (s *DivisionQuery) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *DivisionQuery) SetDivisionCode(v string) *DivisionQuery {
	s.DivisionCode = &v
	return s
}

func (s *DivisionQuery) Validate() error {
	return dara.Validate(s)
}

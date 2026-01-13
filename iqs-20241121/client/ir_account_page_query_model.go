// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIrAccountPageQuery interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *IrAccountPageQuery
	GetName() *string
}

type IrAccountPageQuery struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s IrAccountPageQuery) String() string {
	return dara.Prettify(s)
}

func (s IrAccountPageQuery) GoString() string {
	return s.String()
}

func (s *IrAccountPageQuery) GetName() *string {
	return s.Name
}

func (s *IrAccountPageQuery) SetName(v string) *IrAccountPageQuery {
	s.Name = &v
	return s
}

func (s *IrAccountPageQuery) Validate() error {
	return dara.Validate(s)
}

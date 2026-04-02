// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUuidFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *UuidFilter
	GetEq() *string
	SetIn(v []*string) *UuidFilter
	GetIn() []*string
}

type UuidFilter struct {
	Eq *string   `json:"eq,omitempty" xml:"eq,omitempty"`
	In []*string `json:"in,omitempty" xml:"in,omitempty" type:"Repeated"`
}

func (s UuidFilter) String() string {
	return dara.Prettify(s)
}

func (s UuidFilter) GoString() string {
	return s.String()
}

func (s *UuidFilter) GetEq() *string {
	return s.Eq
}

func (s *UuidFilter) GetIn() []*string {
	return s.In
}

func (s *UuidFilter) SetEq(v string) *UuidFilter {
	s.Eq = &v
	return s
}

func (s *UuidFilter) SetIn(v []*string) *UuidFilter {
	s.In = v
	return s
}

func (s *UuidFilter) Validate() error {
	return dara.Validate(s)
}

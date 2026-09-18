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
	// The exact match condition for the alert rule UUID. Only the alert rule whose UUID equals the specified value is returned.
	//
	// example:
	//
	// a1b2c3d4-e5f6-7890-abcd-ef1234567890
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
	// The set match condition for alert rule UUIDs. All alert rules whose UUIDs are in the specified list are returned.
	//
	// example:
	//
	// ["a1b2c3d4-e5f6-7890-abcd-ef1234567890","b2c3d4e5-f6a7-8901-bcde-f12345678901"]
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

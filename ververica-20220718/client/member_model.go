// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMember interface {
	dara.Model
	String() string
	GoString() string
	SetMember(v string) *Member
	GetMember() *string
	SetRole(v string) *Member
	GetRole() *string
}

type Member struct {
	// This parameter is required.
	//
	// example:
	//
	// user: 181319557522****
	Member *string `json:"member,omitempty" xml:"member,omitempty"`
	// example:
	//
	// VIEWER
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s Member) String() string {
	return dara.Prettify(s)
}

func (s Member) GoString() string {
	return s.String()
}

func (s *Member) GetMember() *string {
	return s.Member
}

func (s *Member) GetRole() *string {
	return s.Role
}

func (s *Member) SetMember(v string) *Member {
	s.Member = &v
	return s
}

func (s *Member) SetRole(v string) *Member {
	s.Role = &v
	return s
}

func (s *Member) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMember interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *Member
	GetMemberId() *string
	SetMemberType(v string) *Member
	GetMemberType() *string
}

type Member struct {
	// example:
	//
	// ken_12345
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// WorkspaceAdmin
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s Member) String() string {
	return dara.Prettify(s)
}

func (s Member) GoString() string {
	return s.String()
}

func (s *Member) GetMemberId() *string {
	return s.MemberId
}

func (s *Member) GetMemberType() *string {
	return s.MemberType
}

func (s *Member) SetMemberId(v string) *Member {
	s.MemberId = &v
	return s
}

func (s *Member) SetMemberType(v string) *Member {
	s.MemberType = &v
	return s
}

func (s *Member) Validate() error {
	return dara.Validate(s)
}

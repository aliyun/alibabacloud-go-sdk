// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRdMemberListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberListShrink(v string) *AddRdMemberListShrinkRequest
	GetMemberListShrink() *string
}

type AddRdMemberListShrinkRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
}

func (s AddRdMemberListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRdMemberListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddRdMemberListShrinkRequest) GetMemberListShrink() *string {
	return s.MemberListShrink
}

func (s *AddRdMemberListShrinkRequest) SetMemberListShrink(v string) *AddRdMemberListShrinkRequest {
	s.MemberListShrink = &v
	return s
}

func (s *AddRdMemberListShrinkRequest) Validate() error {
	return dara.Validate(s)
}

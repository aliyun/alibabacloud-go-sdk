// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRdMemberListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberListShrink(v string) *DeleteRdMemberListShrinkRequest
	GetMemberListShrink() *string
}

type DeleteRdMemberListShrinkRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
}

func (s DeleteRdMemberListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRdMemberListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListShrinkRequest) GetMemberListShrink() *string {
	return s.MemberListShrink
}

func (s *DeleteRdMemberListShrinkRequest) SetMemberListShrink(v string) *DeleteRdMemberListShrinkRequest {
	s.MemberListShrink = &v
	return s
}

func (s *DeleteRdMemberListShrinkRequest) Validate() error {
	return dara.Validate(s)
}

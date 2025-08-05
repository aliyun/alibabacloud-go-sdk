// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUids(v []*int64) *DeleteInstanceMembersRequest
	GetMemberUids() []*int64
}

type DeleteInstanceMembersRequest struct {
	// The UIDs of the members.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234123412341234
	MemberUids []*int64 `json:"MemberUids,omitempty" xml:"MemberUids,omitempty" type:"Repeated"`
}

func (s DeleteInstanceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersRequest) GetMemberUids() []*int64 {
	return s.MemberUids
}

func (s *DeleteInstanceMembersRequest) SetMemberUids(v []*int64) *DeleteInstanceMembersRequest {
	s.MemberUids = v
	return s
}

func (s *DeleteInstanceMembersRequest) Validate() error {
	return dara.Validate(s)
}

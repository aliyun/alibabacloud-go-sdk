// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *AddGroupMemberRequest
	GetGroupId() *string
	SetMemberId(v string) *AddGroupMemberRequest
	GetMemberId() *string
	SetMemberType(v string) *AddGroupMemberRequest
	GetMemberType() *string
}

type AddGroupMemberRequest struct {
	// The ID of the destination group to which the member is added.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The member ID. If member_type is set to user, set this parameter to a user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e4***1b1
	MemberId *string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// The type of the member. Set the value to user. When you create a group, you can directly add the group to a parent group.
	//
	// 	- user
	//
	// Note: A group can be added to only one group. A user can be added to multiple groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s AddGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddGroupMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *AddGroupMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *AddGroupMemberRequest) SetGroupId(v string) *AddGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *AddGroupMemberRequest) SetMemberId(v string) *AddGroupMemberRequest {
	s.MemberId = &v
	return s
}

func (s *AddGroupMemberRequest) SetMemberType(v string) *AddGroupMemberRequest {
	s.MemberType = &v
	return s
}

func (s *AddGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RemoveGroupMemberRequest
	GetGroupId() *string
	SetMemberId(v string) *RemoveGroupMemberRequest
	GetMemberId() *string
	SetMemberType(v string) *RemoveGroupMemberRequest
	GetMemberType() *string
}

type RemoveGroupMemberRequest struct {
	// The ID of the group from which you want to remove a member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The ID of the member. If member_type is set to user, set this parameter to the ID of the corresponding user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e4***1b1
	MemberId *string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// The type of the member that you want to remove from the group. Only common users can be removed. If you want to remove all members from a group, you can directly delete the group. Valid value:
	//
	// 	- user
	//
	// Note: A group can be a member of only one group. It cannot be a member of multiple groups. A user can be a member of multiple groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s RemoveGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveGroupMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *RemoveGroupMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *RemoveGroupMemberRequest) SetGroupId(v string) *RemoveGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetMemberId(v string) *RemoveGroupMemberRequest {
	s.MemberId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetMemberType(v string) *RemoveGroupMemberRequest {
	s.MemberType = &v
	return s
}

func (s *RemoveGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

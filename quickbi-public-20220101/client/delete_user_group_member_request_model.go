// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupId(v string) *DeleteUserGroupMemberRequest
	GetUserGroupId() *string
	SetUserId(v string) *DeleteUserGroupMemberRequest
	GetUserId() *string
}

type DeleteUserGroupMemberRequest struct {
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e537****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user ID of the Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2fe4fbd8****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DeleteUserGroupMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserGroupMemberRequest) SetUserGroupId(v string) *DeleteUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteUserGroupMemberRequest) SetUserId(v string) *DeleteUserGroupMemberRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

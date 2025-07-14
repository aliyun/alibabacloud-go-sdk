// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupIds(v string) *DeleteUserGroupMembersRequest
	GetUserGroupIds() *string
	SetUserId(v string) *DeleteUserGroupMembersRequest
	GetUserId() *string
}

type DeleteUserGroupMembersRequest struct {
	// The ID of the user group(s) to exit.
	//
	// - Supports batch parameters, separate IDs with a comma (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The UserID of the user to be removed from the user group. Note that this UserID refers to the Quick BI UserID, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 204627493484****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *DeleteUserGroupMembersRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserGroupMembersRequest) SetUserGroupIds(v string) *DeleteUserGroupMembersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *DeleteUserGroupMembersRequest) SetUserId(v string) *DeleteUserGroupMembersRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}

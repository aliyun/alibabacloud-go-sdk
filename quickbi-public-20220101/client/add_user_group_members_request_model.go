// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupIds(v string) *AddUserGroupMembersRequest
	GetUserGroupIds() *string
	SetUserId(v string) *AddUserGroupMembersRequest
	GetUserId() *string
}

type AddUserGroupMembersRequest struct {
	// The IDs of the user groups. Separate the IDs with commas (,). Example: aGroupId,bGroupId,cGroupIds
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d5fb19b-****-****-99da-1248fc27ca51
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The user ID of the Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5****37a5
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *AddUserGroupMembersRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddUserGroupMembersRequest) SetUserGroupIds(v string) *AddUserGroupMembersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserId(v string) *AddUserGroupMembersRequest {
	s.UserId = &v
	return s
}

func (s *AddUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}

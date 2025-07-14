// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupId(v string) *AddUserGroupMemberRequest
	GetUserGroupId() *string
	SetUserIdList(v string) *AddUserGroupMemberRequest
	GetUserIdList() *string
}

type AddUserGroupMemberRequest struct {
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 555c4cd****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the Quick BI user. Separate multiple IDs with commas (,). Example: abc,efg. You can enter a maximum of 1000 entries.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e537a5****,3dadsu****
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s AddUserGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMemberRequest) GetUserIdList() *string {
	return s.UserIdList
}

func (s *AddUserGroupMemberRequest) SetUserGroupId(v string) *AddUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMemberRequest) SetUserIdList(v string) *AddUserGroupMemberRequest {
	s.UserIdList = &v
	return s
}

func (s *AddUserGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

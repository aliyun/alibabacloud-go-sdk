// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentUserGroupId(v string) *CreateUserGroupRequest
	GetParentUserGroupId() *string
	SetUserGroupDescription(v string) *CreateUserGroupRequest
	GetUserGroupDescription() *string
	SetUserGroupId(v string) *CreateUserGroupRequest
	GetUserGroupId() *string
	SetUserGroupName(v string) *CreateUserGroupRequest
	GetUserGroupName() *string
}

type CreateUserGroupRequest struct {
	// The ID of the parent user group. You can add new user groups to this group:
	//
	// 	- If you enter the ID of a parent user group, the new user group is added to the user group with the ID.
	//
	// 	- If you enter -1, the new user group is added to the root directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// The description of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// User group description
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The unique ID of the user group.
	//
	// 	- If you specify the UserGroupId parameter, the system automatically generates the UserGroupId parameter. If you specify the UserGroupId parameter, the user ID is used as the user group ID. You must ensure that the user ID is unique within the organization.
	//
	// 	- Format verification: Maximum length 64, cannot be -1,
	//
	// example:
	//
	// pop0001
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the RAM user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou Financial Report
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetParentUserGroupId() *string {
	return s.ParentUserGroupId
}

func (s *CreateUserGroupRequest) GetUserGroupDescription() *string {
	return s.UserGroupDescription
}

func (s *CreateUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreateUserGroupRequest) SetParentUserGroupId(v string) *CreateUserGroupRequest {
	s.ParentUserGroupId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupDescription(v string) *CreateUserGroupRequest {
	s.UserGroupDescription = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupId(v string) *CreateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}

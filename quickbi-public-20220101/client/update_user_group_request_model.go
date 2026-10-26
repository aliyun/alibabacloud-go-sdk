// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupDescription(v string) *UpdateUserGroupRequest
	GetUserGroupDescription() *string
	SetUserGroupId(v string) *UpdateUserGroupRequest
	GetUserGroupId() *string
	SetUserGroupName(v string) *UpdateUserGroupRequest
	GetUserGroupName() *string
}

type UpdateUserGroupRequest struct {
	// The description of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// Description
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// pop0001
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) GetUserGroupDescription() *string {
	return s.UserGroupDescription
}

func (s *UpdateUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *UpdateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *UpdateUserGroupRequest) SetUserGroupDescription(v string) *UpdateUserGroupRequest {
	s.UserGroupDescription = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v string) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupName(v string) *UpdateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *UpdateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}

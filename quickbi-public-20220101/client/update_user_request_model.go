// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminUser(v bool) *UpdateUserRequest
	GetAdminUser() *bool
	SetAuthAdminUser(v bool) *UpdateUserRequest
	GetAuthAdminUser() *bool
	SetIsDeleted(v bool) *UpdateUserRequest
	GetIsDeleted() *bool
	SetNickName(v string) *UpdateUserRequest
	GetNickName() *string
	SetRoleIds(v string) *UpdateUserRequest
	GetRoleIds() *string
	SetUserId(v string) *UpdateUserRequest
	GetUserId() *string
	SetUserType(v int32) *UpdateUserRequest
	GetUserType() *int32
}

type UpdateUserRequest struct {
	// Indicates whether the organization administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Indicate whether the RAM user is a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// User status:
	//
	// 	- **false**: Active
	//
	//  	- **true**: Inactive
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The nickname of the account.
	//
	// 	- Format check: The value can be up to 50 characters in length.
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// Xiao Zhang
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The IDs of the preset or custom organization roles bound to the user, separated by English commas \\",\\", with a maximum of 3. The value range is as follows: - Organization Administrator (preset role): 111111111 - Permission Administrator (preset role): 111111112 - Regular User (preset role): 111111113
	//
	// example:
	//
	// 111111111,456
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The ID of the user to be updated. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of user who is a member of the organization. Valid values:
	//
	// 	- 1 : developer
	//
	// 	- 2 : visitors
	//
	// 	- 3 : Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *UpdateUserRequest) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *UpdateUserRequest) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *UpdateUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *UpdateUserRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *UpdateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserRequest) GetUserType() *int32 {
	return s.UserType
}

func (s *UpdateUserRequest) SetAdminUser(v bool) *UpdateUserRequest {
	s.AdminUser = &v
	return s
}

func (s *UpdateUserRequest) SetAuthAdminUser(v bool) *UpdateUserRequest {
	s.AuthAdminUser = &v
	return s
}

func (s *UpdateUserRequest) SetIsDeleted(v bool) *UpdateUserRequest {
	s.IsDeleted = &v
	return s
}

func (s *UpdateUserRequest) SetNickName(v string) *UpdateUserRequest {
	s.NickName = &v
	return s
}

func (s *UpdateUserRequest) SetRoleIds(v string) *UpdateUserRequest {
	s.RoleIds = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUserType(v int32) *UpdateUserRequest {
	s.UserType = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}

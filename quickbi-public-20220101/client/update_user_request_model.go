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
	SetCopilotModules(v string) *UpdateUserRequest
	GetCopilotModules() *string
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
	// Whether to assign the organization administrator role to the user. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	// 	Notice:
	//
	// This parameter is deprecated and is ignored if RoleIds is also specified.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Whether to assign the permission administrator role to the user. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	// 	Notice:
	//
	// This parameter is deprecated and is ignored if RoleIds is also specified.
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// example:
	//
	// [
	//
	//     {
	//
	//         "moduleType": "smartQAskNum",
	//
	//         "status": 1
	//
	//     },
	//
	//     {
	//
	//         "moduleType": "smartQDevNum",
	//
	//         "status": 0
	//
	//     }
	//
	// ]
	CopilotModules *string `json:"CopilotModules,omitempty" xml:"CopilotModules,omitempty"`
	// The user status:
	//
	// - **`false`**: active
	//
	// - **`true`**: inactive
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The nickname of the user.
	//
	// - The nickname can be up to 50 characters in length.
	//
	// - The nickname can contain Chinese characters, letters, digits, and the following special characters: `_ \\ / | () ] [`
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The IDs of the built-in or custom organization roles to assign to the user. Specify up to three comma-separated role IDs.
	//
	// - organization administrator (built-in role): 111111111
	//
	// - permission administrator (built-in role): 111111112
	//
	// - standard user (built-in role): 111111113
	//
	// example:
	//
	// 111111111,456
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The ID of the Quick BI user to update. This is not an Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user type of the organization member. Valid values:
	//
	// - `1`: developer
	//
	// - `2`: viewer
	//
	// - `3`: analyst
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

func (s *UpdateUserRequest) GetCopilotModules() *string {
	return s.CopilotModules
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

func (s *UpdateUserRequest) SetCopilotModules(v string) *UpdateUserRequest {
	s.CopilotModules = &v
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

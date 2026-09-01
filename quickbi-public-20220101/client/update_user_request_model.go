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
	// Specifies whether to assign the organization administrator role. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// <notice>This parameter is deprecated. When RoleIds is specified, this parameter does not take effect.</notice>
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Specifies whether to assign the organization permission management administrator role. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// <notice>This parameter has expired and is not recommended. When RoleIds is specified, this parameter does not take effect.</notice>
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The intelligent module quota modification information.
	//
	// Pass the parameter as a JSON array. Each array element contains the following fields:
	//
	// moduleType -- The intelligent module.
	//
	// - smartQAskNum -- Smart Q questions.
	//
	// - smartQDevNum -- Smart Q building.
	//
	// - qreport -- Smart Q reports.
	//
	// - qExploreNum -- Smart Q exploration edition.
	//
	// status -- Specifies whether to enable the module.
	//
	// - 0 -- Revoke authorization.
	//
	// - 1 -- Grant authorization.
	//
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
	// The user status. Valid values:
	//
	// 	- **false**: Activated.
	//
	// 	- **true**: Deactivated.
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The nickname.
	//
	// - Format check: The maximum length is 50 characters.
	//
	// - Special format check: Chinese characters, English characters, digits, _ \\ / | () ] [
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The IDs of preset or custom organization roles to attach to the user, separated by commas (,). A maximum of three role IDs are supported. Valid values:
	//
	// - Organization administrator (preset role): 111111111
	//
	// - Permission management administrator (preset role): 111111112
	//
	// - Common user (preset role): 111111113
	//
	// example:
	//
	// 111111111,456
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The ID of the user to update. This user ID is the Quick BI UserID, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user type of the organization member. Valid values:
	//
	// - 1: Developer.
	//
	// - 2: Visitor.
	//
	// - 3: Analyst.
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

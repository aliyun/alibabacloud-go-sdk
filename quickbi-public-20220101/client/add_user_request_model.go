// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AddUserRequest
	GetAccountId() *string
	SetAccountName(v string) *AddUserRequest
	GetAccountName() *string
	SetAdminUser(v bool) *AddUserRequest
	GetAdminUser() *bool
	SetAuthAdminUser(v bool) *AddUserRequest
	GetAuthAdminUser() *bool
	SetCopilotModules(v string) *AddUserRequest
	GetCopilotModules() *string
	SetNickName(v string) *AddUserRequest
	GetNickName() *string
	SetRoleIds(v string) *AddUserRequest
	GetRoleIds() *string
	SetUserType(v int32) *AddUserRequest
	GetUserType() *int32
}

type AddUserRequest struct {
	// The ID of the Alibaba Cloud account.	Warning: The `AccountId` parameter will be required in Quick BI versions released after December 31, 2024. We recommend that you update your API calls to include this parameter before then.
	//
	// example:
	//
	// 191476xxxxx23754
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	//
	// The name of the Alibaba Cloud account.
	//
	// - For a sub-account, use the format `master account:sub-account`. Example: `master_test@aliyun.com:subaccount`.
	//
	// - The maximum length is 50 characters.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Deprecated
	//
	// Specifies whether to assign the organization administrator role. Valid values:
	//
	// - true
	//
	// - false
	//
	// 	Notice:
	//
	// This parameter is deprecated. It is ignored if `RoleIds` is specified.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Deprecated
	//
	// Specifies whether to assign the permission administrator role. Valid values:
	//
	// - true
	//
	// - false
	//
	// 	Notice:
	//
	// This parameter is deprecated. It is ignored if `RoleIds` is specified.
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The Copilot modules to enable for the user. To enable multiple modules, specify their codes separated by a comma (,).
	//
	// - `qreport`: Q Report
	//
	// - `qExploreNum`: Q Explore
	//
	// - `smartQAskNum`: Q\\&A with Data
	//
	// - `smartQDevNum`: Q-assisted Building
	//
	// example:
	//
	// qreport,qExploreNum
	CopilotModules *string `json:"CopilotModules,omitempty" xml:"CopilotModules,omitempty"`
	// The user\\"s nickname.
	//
	// - The maximum length is 50 characters.
	//
	// - The nickname can contain Chinese characters, letters, digits, and the following special characters: `_ \\ / | () []`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张三
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The IDs of the predefined or custom organization roles to assign. You can specify up to three role IDs, separated by commas (,). Valid values for predefined roles:
	//
	// - `111111111`: organization administrator
	//
	// - `111111112`: permission administrator
	//
	// - `111111113`: regular user
	//
	// example:
	//
	// 111111111,456
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The type of the organization member. Valid values:
	//
	// - 1: developer
	//
	// - 2: viewer
	//
	// - 3: analyst
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AddUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserRequest) GoString() string {
	return s.String()
}

func (s *AddUserRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *AddUserRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *AddUserRequest) GetAdminUser() *bool {
	return s.AdminUser
}

func (s *AddUserRequest) GetAuthAdminUser() *bool {
	return s.AuthAdminUser
}

func (s *AddUserRequest) GetCopilotModules() *string {
	return s.CopilotModules
}

func (s *AddUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *AddUserRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *AddUserRequest) GetUserType() *int32 {
	return s.UserType
}

func (s *AddUserRequest) SetAccountId(v string) *AddUserRequest {
	s.AccountId = &v
	return s
}

func (s *AddUserRequest) SetAccountName(v string) *AddUserRequest {
	s.AccountName = &v
	return s
}

func (s *AddUserRequest) SetAdminUser(v bool) *AddUserRequest {
	s.AdminUser = &v
	return s
}

func (s *AddUserRequest) SetAuthAdminUser(v bool) *AddUserRequest {
	s.AuthAdminUser = &v
	return s
}

func (s *AddUserRequest) SetCopilotModules(v string) *AddUserRequest {
	s.CopilotModules = &v
	return s
}

func (s *AddUserRequest) SetNickName(v string) *AddUserRequest {
	s.NickName = &v
	return s
}

func (s *AddUserRequest) SetRoleIds(v string) *AddUserRequest {
	s.RoleIds = &v
	return s
}

func (s *AddUserRequest) SetUserType(v int32) *AddUserRequest {
	s.UserType = &v
	return s
}

func (s *AddUserRequest) Validate() error {
	return dara.Validate(s)
}

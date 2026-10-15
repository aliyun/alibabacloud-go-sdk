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
	// The Alibaba Cloud account ID.
	//
	// 	Warning: For Quick BI versions released after December 31, 2024, AccountId is a required parameter. Update your API calls before this date.
	//
	// example:
	//
	// 191476xxxxx23754
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	//
	// The Alibaba Cloud account name.
	//
	// - Note: For a RAM user, the format is "primary account:RAM user". Example: master_test@aliyun.com:subaccount
	//
	// - Format check: The maximum length is 50 characters.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Deprecated
	//
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
	// Deprecated
	//
	// Specifies whether to assign the organization permission management administrator role. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// <notice>This parameter has expired and is no longer recommended. When RoleIds is specified, this parameter does not take effect.</notice>
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The intelligent module quotas of the user. Separate multiple modules with commas (,).
	//
	// - qreport -- Small Q Report
	//
	// - qExploreNum -- Small Q Explorer
	//
	// - smartQAskNum -- Small Q Ask
	//
	// - smartQDevNum -- Small Q Builder
	//
	// example:
	//
	// qreport,qExploreNum
	CopilotModules *string `json:"CopilotModules,omitempty" xml:"CopilotModules,omitempty"`
	// The nickname of the Alibaba Cloud account.
	//
	// - Format check: The maximum length is 50 characters.
	//
	// - Special character validation: Chinese characters, English characters, digits, _ \\ / | () ] [
	//
	// This parameter is required.
	//
	// example:
	//
	// John
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
	// The user type of the organization member. Valid values:
	//
	// - 1: Developer.
	//
	// - 2: Visitor.
	//
	// - 3: Analyst.
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

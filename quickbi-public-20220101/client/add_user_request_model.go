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
	SetNickName(v string) *AddUserRequest
	GetNickName() *string
	SetRoleIds(v string) *AddUserRequest
	GetRoleIds() *string
	SetUserType(v int32) *AddUserRequest
	GetUserType() *int32
}

type AddUserRequest struct {
	// Aliyun account ID.
	//
	// 	Warning: For versions of Quick BI released after December 31, 2024, AccountId will be a required parameter. Please modify your API before this date.
	//
	// <props="china">Published only on the China site
	//
	// example:
	//
	// 191476xxxxx23754
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	//
	// Aliyun account name.
	//
	// - Note: If it is a sub-account, the format should be \\"primary account: sub-account\\". For example: master_test@aliyun.com:subaccount
	//
	// - Format check: Maximum length of 50 characters.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Deprecated
	//
	// Whether to assign the organization administrator role. Value range:
	//
	// - true: Yes
	//
	// - false: No
	//
	// <notice>This parameter is deprecated and not recommended for use. It is invalid when RoleIds is provided.</notice>
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
	// Whether to assign the organization permission administrator role. Value range:
	//
	// - true: Yes
	//
	// - false: No
	//
	// <notice>This parameter is deprecated and not recommended for use. It is invalid when RoleIds is provided.</notice>
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// Aliyun account nickname.
	//
	// - Format check: Maximum length of 50 characters.
	//
	// - Special format validation: Chinese and English characters, numbers, _ \\ / | () ] [
	//
	// This parameter is required.
	//
	// example:
	//
	// ddd
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// Preset or custom organization role IDs bound to the user, separated by commas, with a maximum of 3. Value range:
	//
	// - Organization Administrator (preset role): 111111111
	//
	// - Permission Administrator (preset role): 111111112
	//
	// - Regular User (preset role): 111111113
	//
	// example:
	//
	// 111111111,456
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The user type of the organization member. Value range:
	//
	// - 1: Developer
	//
	// - 2: Visitor
	//
	// - 3: Analyst
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

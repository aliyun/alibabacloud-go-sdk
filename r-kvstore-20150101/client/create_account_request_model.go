// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetAccountPrivilege(v string) *CreateAccountRequest
	GetAccountPrivilege() *string
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetInstanceId(v string) *CreateAccountRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetParameters(v string) *CreateAccountRequest
	GetParameters() *string
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateAccountRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *CreateAccountRequest
	GetSourceBiz() *string
}

type CreateAccountRequest struct {
	// The description of the account.
	//
	// 	- Must start with a Chinese character or an English letter. Cannot start with `http://` or `https://`.
	//
	// 	- Can contain Chinese characters, English letters, digits, underscores (_), and hyphens (-).
	//
	// 	- Must be 2 to 256 characters in length.
	//
	// example:
	//
	// testaccount
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name. The name must meet the following requirements:
	//
	// 	- Starts with a lowercase letter and contains only lowercase letters, digits, or underscores (_).
	//
	// 	- Contains up to 100 characters.
	//
	// 	- Cannot be a <props="china">[Redis reserved account name](https://www.alibabacloud.com/help/en/redis/user-guide/create-and-manage-database-accounts#section-u3q-817-om3)<props="intl">[Redis reserved account name](https://www.alibabacloud.com/help/zh/redis/user-guide/create-and-manage-database-accounts#section-u3q-817-om3).
	//
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, special characters, and digits. The following special characters are supported: `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// uWonno21****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The permissions of the account. Valid values:
	//
	// 	- **RoleReadOnly**: read-only permissions.
	//
	// 	- **RoleReadWrite**: read and write permissions. This is the default value.
	//
	// example:
	//
	// RoleReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The account type. Set the value to **Normal*	- (standard account).
	//
	// example:
	//
	// Normal
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The account parameters to modify in JSON format. The new values overwrite the original values.
	//
	// example:
	//
	// {"access-db-id":"1","cu-limit":"10"}
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is used only for internal maintenance. You do not need to specify this parameter.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAccountRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetInstanceId(v string) *CreateAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetParameters(v string) *CreateAccountRequest {
	s.Parameters = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetSecurityToken(v string) *CreateAccountRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAccountRequest) SetSourceBiz(v string) *CreateAccountRequest {
	s.SourceBiz = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}

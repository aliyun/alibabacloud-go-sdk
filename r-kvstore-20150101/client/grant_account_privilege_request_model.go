// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GrantAccountPrivilegeRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest
	GetAccountPrivilege() *string
	SetInstanceId(v string) *GrantAccountPrivilegeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GrantAccountPrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantAccountPrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GrantAccountPrivilegeRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *GrantAccountPrivilegeRequest
	GetSourceBiz() *string
}

type GrantAccountPrivilegeRequest struct {
	// The name of the account. You can call the [DescribeAccounts](~~DescribeAccounts~~) operation to obtain the name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions of the account. Default value: RoleReadWrite. Valid values:
	//
	// 	- RoleReadOnly: The account has the read-only permissions.
	//
	// 	- RoleReadWrite: The account has the read and write permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// RoleReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The ID of the instance to which the account belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s GrantAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GrantAccountPrivilegeRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *GrantAccountPrivilegeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantAccountPrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantAccountPrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantAccountPrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantAccountPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantAccountPrivilegeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GrantAccountPrivilegeRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetInstanceId(v string) *GrantAccountPrivilegeRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetSecurityToken(v string) *GrantAccountPrivilegeRequest {
	s.SecurityToken = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetSourceBiz(v string) *GrantAccountPrivilegeRequest {
	s.SourceBiz = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}

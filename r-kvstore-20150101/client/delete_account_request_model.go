// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAccountRequest
	GetAccountName() *string
	SetInstanceId(v string) *DeleteAccountRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAccountRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteAccountRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *DeleteAccountRequest
	GetSourceBiz() *string
}

type DeleteAccountRequest struct {
	// The username of the account. You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/473816.html) operation to query the username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
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

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAccountRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAccountRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetInstanceId(v string) *DeleteAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetSecurityToken(v string) *DeleteAccountRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAccountRequest) SetSourceBiz(v string) *DeleteAccountRequest {
	s.SourceBiz = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}

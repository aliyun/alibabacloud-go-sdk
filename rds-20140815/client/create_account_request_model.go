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
	SetAccountType(v string) *CreateAccountRequest
	GetAccountType() *string
	SetCheckPolicy(v bool) *CreateAccountRequest
	GetCheckPolicy() *bool
	SetDBInstanceId(v string) *CreateAccountRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
}

type CreateAccountRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	CheckPolicy     *bool   `json:"CheckPolicy,omitempty" xml:"CheckPolicy,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *CreateAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAccountRequest) GetCheckPolicy() *bool {
	return s.CheckPolicy
}

func (s *CreateAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetCheckPolicy(v bool) *CreateAccountRequest {
	s.CheckPolicy = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
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

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}

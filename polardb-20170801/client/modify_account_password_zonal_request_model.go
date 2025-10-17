// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPasswordZonalRequest
	GetAccountName() *string
	SetClientToken(v string) *ModifyAccountPasswordZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyAccountPasswordZonalRequest
	GetDBClusterId() *string
	SetNewAccountPassword(v string) *ModifyAccountPasswordZonalRequest
	GetNewAccountPassword() *string
	SetOwnerAccount(v string) *ModifyAccountPasswordZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountPasswordZonalRequest
	GetOwnerId() *int64
	SetPasswordType(v string) *ModifyAccountPasswordZonalRequest
	GetPasswordType() *string
	SetResourceOwnerAccount(v string) *ModifyAccountPasswordZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountPasswordZonalRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountPasswordZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// Tair
	PasswordType         *string `json:"PasswordType,omitempty" xml:"PasswordType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountPasswordZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPasswordZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAccountPasswordZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPasswordZonalRequest) GetNewAccountPassword() *string {
	return s.NewAccountPassword
}

func (s *ModifyAccountPasswordZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountPasswordZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountPasswordZonalRequest) GetPasswordType() *string {
	return s.PasswordType
}

func (s *ModifyAccountPasswordZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountPasswordZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountPasswordZonalRequest) SetAccountName(v string) *ModifyAccountPasswordZonalRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetClientToken(v string) *ModifyAccountPasswordZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetDBClusterId(v string) *ModifyAccountPasswordZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordZonalRequest {
	s.NewAccountPassword = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetOwnerAccount(v string) *ModifyAccountPasswordZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetOwnerId(v int64) *ModifyAccountPasswordZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetPasswordType(v string) *ModifyAccountPasswordZonalRequest {
	s.PasswordType = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetResourceOwnerAccount(v string) *ModifyAccountPasswordZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) SetResourceOwnerId(v int64) *ModifyAccountPasswordZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountPasswordZonalRequest) Validate() error {
	return dara.Validate(s)
}

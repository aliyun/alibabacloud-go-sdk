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
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests, case-sensitive, and up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the edge cluster. This operation applies only to zonal edge clusters. You can obtain a valid ID by calling the DescribeDBClustersZonal operation. Standard PolarDB cluster IDs are not supported. For standard clusters, use the ModifyAccountPassword operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The new password of the account. The password must meet the following requirements:
	//
	// 	- Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Is 8 to 32 characters in length.
	//
	// 	- Special characters include `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the password.
	//
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

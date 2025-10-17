// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountZonalRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountZonalRequest
	GetAccountPassword() *string
	SetClientToken(v string) *ResetAccountZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ResetAccountZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ResetAccountZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetAccountZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResetAccountZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetAccountZonalRequest
	GetResourceOwnerId() *int64
}

type ResetAccountZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// Test1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f54a343f1a42***********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetAccountZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountZonalRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountZonalRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResetAccountZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetAccountZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetAccountZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetAccountZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetAccountZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetAccountZonalRequest) SetAccountName(v string) *ResetAccountZonalRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountZonalRequest) SetAccountPassword(v string) *ResetAccountZonalRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountZonalRequest) SetClientToken(v string) *ResetAccountZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetAccountZonalRequest) SetDBClusterId(v string) *ResetAccountZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountZonalRequest) SetOwnerAccount(v string) *ResetAccountZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountZonalRequest) SetOwnerId(v int64) *ResetAccountZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountZonalRequest) SetResourceOwnerAccount(v string) *ResetAccountZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountZonalRequest) SetResourceOwnerId(v int64) *ResetAccountZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountZonalRequest) Validate() error {
	return dara.Validate(s)
}

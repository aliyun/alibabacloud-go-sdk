// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *RevokeAccountPrivilegeZonalRequest
	GetAccountName() *string
	SetClientToken(v string) *RevokeAccountPrivilegeZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *RevokeAccountPrivilegeZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *RevokeAccountPrivilegeZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *RevokeAccountPrivilegeZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeAccountPrivilegeZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeAccountPrivilegeZonalRequest
	GetResourceOwnerId() *int64
}

type RevokeAccountPrivilegeZonalRequest struct {
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
	// testdb
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeAccountPrivilegeZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeZonalRequest) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *RevokeAccountPrivilegeZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeAccountPrivilegeZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RevokeAccountPrivilegeZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *RevokeAccountPrivilegeZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeAccountPrivilegeZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeAccountPrivilegeZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeAccountPrivilegeZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeAccountPrivilegeZonalRequest) SetAccountName(v string) *RevokeAccountPrivilegeZonalRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetClientToken(v string) *RevokeAccountPrivilegeZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetDBClusterId(v string) *RevokeAccountPrivilegeZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetDBName(v string) *RevokeAccountPrivilegeZonalRequest {
	s.DBName = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetOwnerAccount(v string) *RevokeAccountPrivilegeZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetOwnerId(v int64) *RevokeAccountPrivilegeZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) SetResourceOwnerId(v int64) *RevokeAccountPrivilegeZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalRequest) Validate() error {
	return dara.Validate(s)
}

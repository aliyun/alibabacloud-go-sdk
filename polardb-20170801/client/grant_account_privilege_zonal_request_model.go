// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GrantAccountPrivilegeZonalRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *GrantAccountPrivilegeZonalRequest
	GetAccountPrivilege() *string
	SetClientToken(v string) *GrantAccountPrivilegeZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *GrantAccountPrivilegeZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *GrantAccountPrivilegeZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *GrantAccountPrivilegeZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantAccountPrivilegeZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GrantAccountPrivilegeZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantAccountPrivilegeZonalRequest
	GetResourceOwnerId() *int64
}

type GrantAccountPrivilegeZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ReadWrite,ReadOnly
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb_1,testdb_2
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantAccountPrivilegeZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeZonalRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GrantAccountPrivilegeZonalRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *GrantAccountPrivilegeZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GrantAccountPrivilegeZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GrantAccountPrivilegeZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *GrantAccountPrivilegeZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantAccountPrivilegeZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantAccountPrivilegeZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantAccountPrivilegeZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantAccountPrivilegeZonalRequest) SetAccountName(v string) *GrantAccountPrivilegeZonalRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeZonalRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetClientToken(v string) *GrantAccountPrivilegeZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetDBClusterId(v string) *GrantAccountPrivilegeZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetDBName(v string) *GrantAccountPrivilegeZonalRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetOwnerId(v int64) *GrantAccountPrivilegeZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeZonalRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *RevokeAccountPrivilegeRequest
	GetAccountName() *string
	SetDBInstanceId(v string) *RevokeAccountPrivilegeRequest
	GetDBInstanceId() *string
	SetDBName(v string) *RevokeAccountPrivilegeRequest
	GetDBName() *string
	SetOwnerAccount(v string) *RevokeAccountPrivilegeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeAccountPrivilegeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeAccountPrivilegeRequest
	GetResourceOwnerId() *int64
}

type RevokeAccountPrivilegeRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeAccountPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *RevokeAccountPrivilegeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RevokeAccountPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *RevokeAccountPrivilegeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeAccountPrivilegeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeAccountPrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeAccountPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeAccountPrivilegeRequest) SetAccountName(v string) *RevokeAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBInstanceId(v string) *RevokeAccountPrivilegeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBName(v string) *RevokeAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}

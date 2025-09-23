// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortDBClusterMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AbortDBClusterMigrationRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AbortDBClusterMigrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AbortDBClusterMigrationRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *AbortDBClusterMigrationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AbortDBClusterMigrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AbortDBClusterMigrationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *AbortDBClusterMigrationRequest
	GetSecurityToken() *string
}

type AbortDBClusterMigrationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-bp150t3l1td9863q9
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AbortDBClusterMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *AbortDBClusterMigrationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AbortDBClusterMigrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AbortDBClusterMigrationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AbortDBClusterMigrationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AbortDBClusterMigrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AbortDBClusterMigrationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AbortDBClusterMigrationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AbortDBClusterMigrationRequest) SetDBClusterId(v string) *AbortDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetOwnerAccount(v string) *AbortDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetOwnerId(v int64) *AbortDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetResourceGroupId(v string) *AbortDBClusterMigrationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *AbortDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetResourceOwnerId(v int64) *AbortDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) SetSecurityToken(v string) *AbortDBClusterMigrationRequest {
	s.SecurityToken = &v
	return s
}

func (s *AbortDBClusterMigrationRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDBClusterMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ContinueDBClusterMigrationRequest
	GetDBClusterId() *string
	SetForceSwitch(v string) *ContinueDBClusterMigrationRequest
	GetForceSwitch() *string
	SetOwnerAccount(v string) *ContinueDBClusterMigrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ContinueDBClusterMigrationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ContinueDBClusterMigrationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ContinueDBClusterMigrationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ContinueDBClusterMigrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ContinueDBClusterMigrationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ContinueDBClusterMigrationRequest
	GetSecurityToken() *string
}

type ContinueDBClusterMigrationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-k2ju1lnl5i4ohv501
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// false
	ForceSwitch  *string `json:"ForceSwitch,omitempty" xml:"ForceSwitch,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ContinueDBClusterMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *ContinueDBClusterMigrationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ContinueDBClusterMigrationRequest) GetForceSwitch() *string {
	return s.ForceSwitch
}

func (s *ContinueDBClusterMigrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ContinueDBClusterMigrationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ContinueDBClusterMigrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContinueDBClusterMigrationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ContinueDBClusterMigrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ContinueDBClusterMigrationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ContinueDBClusterMigrationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ContinueDBClusterMigrationRequest) SetDBClusterId(v string) *ContinueDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetForceSwitch(v string) *ContinueDBClusterMigrationRequest {
	s.ForceSwitch = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetOwnerAccount(v string) *ContinueDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetOwnerId(v int64) *ContinueDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetRegionId(v string) *ContinueDBClusterMigrationRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetResourceGroupId(v string) *ContinueDBClusterMigrationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *ContinueDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetResourceOwnerId(v int64) *ContinueDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) SetSecurityToken(v string) *ContinueDBClusterMigrationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ContinueDBClusterMigrationRequest) Validate() error {
	return dara.Validate(s)
}

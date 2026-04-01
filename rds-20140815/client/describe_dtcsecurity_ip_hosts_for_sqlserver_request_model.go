// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTCSecurityIpHostsForSQLServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest
	GetSecurityToken() *string
}

type DescribeDTCSecurityIpHostsForSQLServerRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeDBInstanceAttribute operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDTCSecurityIpHostsForSQLServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTCSecurityIpHostsForSQLServerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetDBInstanceId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetOwnerAccount(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetOwnerId(v int64) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetRegionId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetResourceGroupId(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetResourceOwnerAccount(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetResourceOwnerId(v int64) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) SetSecurityToken(v string) *DescribeDTCSecurityIpHostsForSQLServerRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerRequest) Validate() error {
	return dara.Validate(s)
}

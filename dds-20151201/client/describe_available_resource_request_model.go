// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceClass(v string) *DescribeAvailableResourceRequest
	GetDBInstanceClass() *string
	SetDbType(v string) *DescribeAvailableResourceRequest
	GetDbType() *string
	SetEngineVersion(v string) *DescribeAvailableResourceRequest
	GetEngineVersion() *string
	SetInstanceChargeType(v string) *DescribeAvailableResourceRequest
	GetInstanceChargeType() *string
	SetOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetReplicationFactor(v string) *DescribeAvailableResourceRequest
	GetReplicationFactor() *string
	SetResourceGroupId(v string) *DescribeAvailableResourceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest
	GetResourceOwnerId() *int64
	SetStorageType(v string) *DescribeAvailableResourceRequest
	GetStorageType() *string
	SetZoneId(v string) *DescribeAvailableResourceRequest
	GetZoneId() *string
}

type DescribeAvailableResourceRequest struct {
	// The instance type of the instance.
	//
	// example:
	//
	// dds.mongo.standard
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **normal**: replica set instance
	//
	// 	- **sharding**: sharded cluster instance
	//
	// example:
	//
	// sharding
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The major engine version of the instance.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid*	- (default): subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the latest available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of nodes, only applicable to replica sets.
	//
	// example:
	//
	// 3
	ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- local_ssd: local SSD
	//
	// 	- cloud_essd1: PL1 enhanced SSD (ESSD)
	//
	// 	- cloud_essd2: PL2 ESSD
	//
	// 	- cloud_essd3: PL3 ESSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// This parameter is empty by default, which indicates all types of storage resources are queried.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// local_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The ID of the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the available zones.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeAvailableResourceRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeAvailableResourceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableResourceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeAvailableResourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableResourceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceRequest) SetDBInstanceClass(v string) *DescribeAvailableResourceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDbType(v string) *DescribeAvailableResourceRequest {
	s.DbType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngineVersion(v string) *DescribeAvailableResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetReplicationFactor(v string) *DescribeAvailableResourceRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceGroupId(v string) *DescribeAvailableResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetStorageType(v string) *DescribeAvailableResourceRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}

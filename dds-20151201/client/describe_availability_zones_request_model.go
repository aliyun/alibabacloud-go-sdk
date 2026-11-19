// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailabilityZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeAvailabilityZonesRequest
	GetAcceptLanguage() *string
	SetDBInstanceClass(v string) *DescribeAvailabilityZonesRequest
	GetDBInstanceClass() *string
	SetDbType(v string) *DescribeAvailabilityZonesRequest
	GetDbType() *string
	SetEngineVersion(v string) *DescribeAvailabilityZonesRequest
	GetEngineVersion() *string
	SetExcludeSecondaryZoneId(v string) *DescribeAvailabilityZonesRequest
	GetExcludeSecondaryZoneId() *string
	SetExcludeZoneId(v string) *DescribeAvailabilityZonesRequest
	GetExcludeZoneId() *string
	SetInstanceChargeType(v string) *DescribeAvailabilityZonesRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeAvailabilityZonesRequest
	GetInstanceType() *string
	SetMongoType(v string) *DescribeAvailabilityZonesRequest
	GetMongoType() *string
	SetOwnerAccount(v string) *DescribeAvailabilityZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailabilityZonesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailabilityZonesRequest
	GetRegionId() *string
	SetReplicationFactor(v string) *DescribeAvailabilityZonesRequest
	GetReplicationFactor() *string
	SetResourceGroupId(v string) *DescribeAvailabilityZonesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailabilityZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailabilityZonesRequest
	GetResourceOwnerId() *int64
	SetStorageSupport(v string) *DescribeAvailabilityZonesRequest
	GetStorageSupport() *string
	SetStorageType(v string) *DescribeAvailabilityZonesRequest
	GetStorageType() *string
	SetZoneId(v string) *DescribeAvailabilityZonesRequest
	GetZoneId() *string
}

type DescribeAvailabilityZonesRequest struct {
	// The language of the **RegionName*	- and **ZoneName*	- response parameters. Valid values:
	//
	// - **zh**: Chinese. Default value.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The instance type.
	//
	// example:
	//
	// dds.mongo.standard
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The database type of the instance. Valid values:
	//
	// - **normal**: replica set instance.
	//
	// - **sharding**: sharded cluster instance.
	//
	// example:
	//
	// normal
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The secondary zone ID to be excluded from the query results in the current region. This parameter is used together with ExcludeZoneId to filter multiple zones.
	//
	// example:
	//
	// cn-shanghai-b
	ExcludeSecondaryZoneId *string `json:"ExcludeSecondaryZoneId,omitempty" xml:"ExcludeSecondaryZoneId,omitempty"`
	// The zone ID to be excluded from the query results in the current region.
	//
	// example:
	//
	// cn-shanghai-g
	ExcludeZoneId *string `json:"ExcludeZoneId,omitempty" xml:"ExcludeZoneId,omitempty"`
	// The billing method. Valid values:
	//
	// - **PrePaid**: subscription. Default value.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type. Valid values:
	//
	// - **sharding**: sharded cluster instance.
	//
	// - **replicate**: replica set instance or standalone instance.
	//
	// example:
	//
	// replicate
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The MongoDB edition, such as the high-availability edition or the test edition (dbfs).
	//
	// example:
	//
	// dbfs
	MongoType    *string `json:"MongoType,omitempty" xml:"MongoType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of nodes.
	//
	// > This parameter is supported only for replica set instances.
	//
	// example:
	//
	// 3
	ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The resource group ID. For more information about resource groups, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmx2m4rqu7pry
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud**: displays only the zones that support cloud disk instances.
	//
	// - **local**: displays only the zones that support local disk instances.
	//
	// - **default*	- or empty: displays the zones that support both local disk and cloud disk instances.
	//
	// example:
	//
	// local
	StorageSupport *string `json:"StorageSupport,omitempty" xml:"StorageSupport,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud_essd1**: PL1 ESSD.
	//
	// - **cloud_essd2**: PL2 ESSD.
	//
	// - **cloud_essd3**: PL3 ESSD.
	//
	// - **local_ssd**: local SSD.
	//
	// >- Instances of version 4.4 and later support only cloud disks. If you do not specify this parameter, **cloud_essd1*	- is selected by default.
	//
	// >- Instances of version 4.2 and earlier support only local disks. If you do not specify this parameter, **local_ssd*	- is selected by default.
	//
	// example:
	//
	// local_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailabilityZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailabilityZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailabilityZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeAvailabilityZonesRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeAvailabilityZonesRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeAvailabilityZonesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailabilityZonesRequest) GetExcludeSecondaryZoneId() *string {
	return s.ExcludeSecondaryZoneId
}

func (s *DescribeAvailabilityZonesRequest) GetExcludeZoneId() *string {
	return s.ExcludeZoneId
}

func (s *DescribeAvailabilityZonesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailabilityZonesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAvailabilityZonesRequest) GetMongoType() *string {
	return s.MongoType
}

func (s *DescribeAvailabilityZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailabilityZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailabilityZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailabilityZonesRequest) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeAvailabilityZonesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAvailabilityZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailabilityZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailabilityZonesRequest) GetStorageSupport() *string {
	return s.StorageSupport
}

func (s *DescribeAvailabilityZonesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAvailabilityZonesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailabilityZonesRequest) SetAcceptLanguage(v string) *DescribeAvailabilityZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetDBInstanceClass(v string) *DescribeAvailabilityZonesRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetDbType(v string) *DescribeAvailabilityZonesRequest {
	s.DbType = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetEngineVersion(v string) *DescribeAvailabilityZonesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetExcludeSecondaryZoneId(v string) *DescribeAvailabilityZonesRequest {
	s.ExcludeSecondaryZoneId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetExcludeZoneId(v string) *DescribeAvailabilityZonesRequest {
	s.ExcludeZoneId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetInstanceChargeType(v string) *DescribeAvailabilityZonesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetInstanceType(v string) *DescribeAvailabilityZonesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetMongoType(v string) *DescribeAvailabilityZonesRequest {
	s.MongoType = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetOwnerAccount(v string) *DescribeAvailabilityZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetOwnerId(v int64) *DescribeAvailabilityZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetRegionId(v string) *DescribeAvailabilityZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetReplicationFactor(v string) *DescribeAvailabilityZonesRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetResourceGroupId(v string) *DescribeAvailabilityZonesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetResourceOwnerAccount(v string) *DescribeAvailabilityZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetResourceOwnerId(v int64) *DescribeAvailabilityZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetStorageSupport(v string) *DescribeAvailabilityZonesRequest {
	s.StorageSupport = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetStorageType(v string) *DescribeAvailabilityZonesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) SetZoneId(v string) *DescribeAvailabilityZonesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailabilityZonesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeDBInstancesRequest
	GetChargeType() *string
	SetConnectionDomain(v string) *DescribeDBInstancesRequest
	GetConnectionDomain() *string
	SetDBInstanceClass(v string) *DescribeDBInstancesRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *DescribeDBInstancesRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *DescribeDBInstancesRequest
	GetDBInstanceId() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesRequest
	GetDBInstanceStatus() *string
	SetDBInstanceType(v string) *DescribeDBInstancesRequest
	GetDBInstanceType() *string
	SetDBNodeType(v string) *DescribeDBInstancesRequest
	GetDBNodeType() *string
	SetEngine(v string) *DescribeDBInstancesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBInstancesRequest
	GetEngineVersion() *string
	SetExpireTime(v string) *DescribeDBInstancesRequest
	GetExpireTime() *string
	SetExpired(v string) *DescribeDBInstancesRequest
	GetExpired() *string
	SetNetworkType(v string) *DescribeDBInstancesRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeDBInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesRequest
	GetRegionId() *string
	SetReplicationFactor(v string) *DescribeDBInstancesRequest
	GetReplicationFactor() *string
	SetResourceGroupId(v string) *DescribeDBInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest
	GetTag() []*DescribeDBInstancesRequestTag
	SetVSwitchId(v string) *DescribeDBInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeDBInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstancesRequest
	GetZoneId() *string
}

type DescribeDBInstancesRequest struct {
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The endpoint of the node. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the endpoint of the node.
	//
	// example:
	//
	// dds-bp1ea17b41abecf43****.mongodb.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The instance type. For more information about valid values, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// example:
	//
	// dds.mongo.mid
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance. The name must meet the following requirements:
	//
	// 	- The name must start with a letter.
	//
	// 	- It can contain digits, letters, underscores (_), and hyphens (-).
	//
	// 	- It must be 2 to 256 characters in length.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp199659b178****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the instance. For more information about valid values, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **sharding**: sharded cluster instance
	//
	// 	- **replicate**: replica set or standalone instance
	//
	// example:
	//
	// sharding
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The type of the node in the instance. This parameter is used to filter standard or test instance.
	//
	// 1.  Valid value for a standalone or DBFS instance.
	//
	// 2.  Valid value for a standard instance that comes in the replica set or sharded cluster architecture: standard
	//
	// 3.  Valid value when all instances are displayed: default
	//
	// example:
	//
	// default
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The database engine of the instance. Set the value to **MongoDB**.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// 	- **6.0**
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// 	- **3.4**
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2019-12-26T16:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Specifies whether the instance has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value of this parameter must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of nodes in the replica set instance. Valid values:
	//
	// 	- **3**
	//
	// 	- **5**
	//
	// 	- **7**
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
	// The tags of the instance.
	Tag []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID of the instance.
	//
	// example:
	//
	// vsw-bp1vj604nj5a9zz74****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesRequest) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeDBInstancesRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesRequest) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesRequest) GetDBNodeType() *string {
	return s.DBNodeType
}

func (s *DescribeDBInstancesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstancesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesRequest) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeDBInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesRequest) GetTag() []*DescribeDBInstancesRequestTag {
	return s.Tag
}

func (s *DescribeDBInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesRequest) SetChargeType(v string) *DescribeDBInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetConnectionDomain(v string) *DescribeDBInstancesRequest {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceClass(v string) *DescribeDBInstancesRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceId(v string) *DescribeDBInstancesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceType(v string) *DescribeDBInstancesRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBNodeType(v string) *DescribeDBInstancesRequest {
	s.DBNodeType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngine(v string) *DescribeDBInstancesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngineVersion(v string) *DescribeDBInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetExpireTime(v string) *DescribeDBInstancesRequest {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetExpired(v string) *DescribeDBInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetNetworkType(v string) *DescribeDBInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetReplicationFactor(v string) *DescribeDBInstancesRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesRequest) SetVSwitchId(v string) *DescribeDBInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetVpcId(v string) *DescribeDBInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetZoneId(v string) *DescribeDBInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesRequestTag struct {
	// The tag key of the instance. Valid values of N: **1*	- to **20**.
	//
	// 	- The key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- It can be up to 64 characters in length.
	//
	// 	- It cannot be an empty string.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values of N: **1*	- to **20**.
	//
	// 	- The value cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- The value can be up to 128 characters in length.
	//
	// 	- It can be an empty string.
	//
	// example:
	//
	// apitest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}

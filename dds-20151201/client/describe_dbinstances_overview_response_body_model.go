// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstances(v []*DescribeDBInstancesOverviewResponseBodyDBInstances) *DescribeDBInstancesOverviewResponseBody
	GetDBInstances() []*DescribeDBInstancesOverviewResponseBodyDBInstances
	SetRequestId(v string) *DescribeDBInstancesOverviewResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeDBInstancesOverviewResponseBody
	GetTotalCount() *string
}

type DescribeDBInstancesOverviewResponseBody struct {
	// The information of instances.
	DBInstances []*DescribeDBInstancesOverviewResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 52820D2B-B2DD-59F0-BDF2-83EC19C6F1CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances in the query results.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstancesOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponseBody) GetDBInstances() []*DescribeDBInstancesOverviewResponseBodyDBInstances {
	return s.DBInstances
}

func (s *DescribeDBInstancesOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesOverviewResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeDBInstancesOverviewResponseBody) SetDBInstances(v []*DescribeDBInstancesOverviewResponseBodyDBInstances) *DescribeDBInstancesOverviewResponseBody {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBody) SetRequestId(v string) *DescribeDBInstancesOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBody) SetTotalCount(v string) *DescribeDBInstancesOverviewResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesOverviewResponseBodyDBInstances struct {
	// The read and write throughput consumed by the instance.
	//
	// >
	//
	// 	- This parameter is returned when the instance is a serverless instance.
	//
	// 	- Serverless instances are available only in the China site (aliyun.com).
	//
	// example:
	//
	// 100
	CapacityUnit *string `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
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
	// The time when the instance was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-05T03:18:53Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance type. The instance type varies based on the instance architecture. For more information about instance types supported by different instance architectures, see the following references:
	//
	// 	- [Standalone instance types](https://help.aliyun.com/document_detail/311407.html)
	//
	// 	- [Replica set instance types](https://help.aliyun.com/document_detail/311410.html)
	//
	// 	- [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html)
	//
	// example:
	//
	// dds.mongo.mid
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test db
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the instance. For more information about valid values, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage space of the instance. Unit: GB.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **sharding**: sharded cluster instance
	//
	// 	- **replicate**: replica set or standalone instance
	//
	// example:
	//
	// replicate
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The time when the instance data was destroyed. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-10T16:00:00Z
	DestroyTime *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	// The database engine of the instance. Set the value to **MongoDB**.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 4.2
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-02-05T16:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The kind code of the instance. Valid values:
	//
	// 	- **0**: physical machine
	//
	// 	- **1**: Elastic Compute Service (ECS) instance
	//
	// 	- **2**: Docker cluster
	//
	// 	- **18**: Kubernetes cluster
	//
	// example:
	//
	// 0
	KindCode *string `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// The date when the last downgrade operation was performed.
	//
	// example:
	//
	// 2021-05-08
	LastDowngradeTime *string `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	// Indicates whether the instance is locked. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before it is rolled back.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked after the storage space is exhausted.
	//
	// 	- **Released**: The instance is released. After an instance is released, the instance cannot be unlocked. You can only restore the backup data of the instance to a new instance. This process requires a long period of time.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The details of the mongos nodes.
	//
	// >  This parameter is returned when the instance is a sharded cluster instance.
	MongosList []*DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**
	//
	// 	- **VPC**
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of nodes in the instance.
	//
	// >  This parameter is returned when the instance is a replica set instance.
	//
	// example:
	//
	// 3
	ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm22cdcgc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The details of the shard nodes.
	//
	// >  This parameter is returned when the instance is a sharded cluster instance.
	ShardList []*DescribeDBInstancesOverviewResponseBodyDBInstancesShardList `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Repeated"`
	// The tags to add to the instance.
	Tags []*DescribeDBInstancesOverviewResponseBodyDBInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether password-free access over VPC is enabled. Valid values:
	//
	// 	- **Open**: Password-free access over VPC is enabled.
	//
	// 	- **Close**: Password-free access over VPC is disabled.
	//
	// example:
	//
	// Open
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetKindCode() *string {
	return s.KindCode
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetLastDowngradeTime() *string {
	return s.LastDowngradeTime
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetMongosList() []*DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList {
	return s.MongosList
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetShardList() []*DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	return s.ShardList
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetTags() []*DescribeDBInstancesOverviewResponseBodyDBInstancesTags {
	return s.Tags
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetCapacityUnit(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetChargeType(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetCreationTime(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceClass(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceDescription(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceStatus(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceStorage(v int32) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDBInstanceType(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetDestroyTime(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.DestroyTime = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetEngine(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetEngineVersion(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetExpireTime(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetKindCode(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetLastDowngradeTime(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.LastDowngradeTime = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetLockMode(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetMongosList(v []*DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.MongosList = v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetNetworkType(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetRegionId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetReplicationFactor(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetResourceGroupId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetShardList(v []*DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ShardList = v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetTags(v []*DescribeDBInstancesOverviewResponseBodyDBInstancesTags) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetVpcAuthMode(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) SetZoneId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList struct {
	// The instance type of the mongos node.
	//
	// example:
	//
	// dds.mongos.standard
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The description of the mongos node.
	//
	// example:
	//
	// mongos node describe.
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the mongos node.
	//
	// example:
	//
	// s-bp10e3b0d02f****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) SetNodeClass(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) SetNodeDescription(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) SetNodeId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesOverviewResponseBodyDBInstancesShardList struct {
	// The instance type of the shard node.
	//
	// example:
	//
	// dds.shard.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The description of the shard node.
	//
	// example:
	//
	// testshard
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the shard node.
	//
	// example:
	//
	// d-bp1cac6f2083****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage space of the shard node. Unit: GB.
	//
	// example:
	//
	// 10
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The number of read-only nodes in the shard node. Valid values: **0*	- to **5**.
	//
	// example:
	//
	// 2
	ReadonlyReplicas *int32 `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) SetNodeClass(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) SetNodeDescription(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) SetNodeId(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) SetNodeStorage(v int32) *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) SetReadonlyReplicas(v int32) *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList {
	s.ReadonlyReplicas = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesShardList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesOverviewResponseBodyDBInstancesTags struct {
	// The tag key. Valid values of N: **1*	- to **20**.
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
	// The tag value. Valid values of N: **1*	- to **20**.
	//
	// 	- The value cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- The value can be up to 128 characters in length.
	//
	// 	- The value can be an empty string.
	//
	// example:
	//
	// apitest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewResponseBodyDBInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesTags) SetKey(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesTags) SetValue(v string) *DescribeDBInstancesOverviewResponseBodyDBInstancesTags {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesOverviewResponseBodyDBInstancesTags) Validate() error {
	return dara.Validate(s)
}

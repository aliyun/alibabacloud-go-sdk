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
	// A list of instance details.
	DBInstances []*DescribeDBInstancesOverviewResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 52820D2B-B2DD-59F0-BDF2-83EC19C6F1CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances returned.
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
	if s.DBInstances != nil {
		for _, item := range s.DBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesOverviewResponseBodyDBInstances struct {
	// The read/write throughput of the instance.
	//
	// > - This parameter is returned only for Serverless instances.
	//
	// >
	//
	// > - Serverless instances are available only on the China site (aliyun.com).
	//
	// example:
	//
	// 100
	CapacityUnit *string `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid**: subscription
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the instance was created. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is in UTC.
	//
	// example:
	//
	// 2022-01-05T03:18:53Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance type. For more information about the instance types available for different instance architectures, see:
	//
	// - [Standalone instance types](https://help.aliyun.com/document_detail/311407.html)
	//
	// - [Replica set instance types](https://help.aliyun.com/document_detail/311410.html)
	//
	// - [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html)
	//
	// <props="china">
	//
	// > This parameter is not required for Serverless instances.
	//
	// example:
	//
	// dds.mongo.mid
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage capacity of the instance, in GB.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance architecture. Valid values:
	//
	// - **sharding**: sharded cluster instance
	//
	// - **replicate**: replica set or standalone instance
	//
	// <props="china">
	//
	// - **serverless**: Serverless instance
	//
	// example:
	//
	// replicate
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The time when the data of the instance was destroyed. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2021-12-10T16:00:00Z
	DestroyTime *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	// The database engine. The value is **MongoDB**.
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
	// The expiration time of the subscription instance. The time is in the *yyyy-MM-dd*T*HH:mm*Z format. The time is in UTC.
	//
	// example:
	//
	// 2022-02-05T16:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance edition. Valid values:
	//
	// - **0**: physical server
	//
	// - **1**: ECS
	//
	// - **2**: DOCKER
	//
	// - **18**: instance on the new Kubernetes-based architecture
	//
	// example:
	//
	// 0
	KindCode *string `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// The time when the instance was last downgraded.
	//
	// example:
	//
	// 2021-05-08
	LastDowngradeTime *string `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	// The lock mode of the instance.
	//
	// - **Unlock**: The instance is not locked.
	//
	// - **ManualLock**: The instance is manually locked.
	//
	// - **LockByExpiration**: The instance is automatically locked after it expires.
	//
	// - **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// - **LockByDiskQuota**: The instance is automatically locked after its storage space is exhausted.
	//
	// - **Released**: The instance is released. You cannot unlock a released instance. You can only restore the data of the instance to a new instance. The restoration may take a long time.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The details of the mongos nodes.
	//
	// > This parameter is returned only for sharded cluster instances.
	MongosList []*DescribeDBInstancesOverviewResponseBodyDBInstancesMongosList `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// - **Classic**: classic network
	//
	// - **VPC**: virtual private cloud (VPC)
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
	// > This parameter is returned only for replica set instances.
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
	// > This parameter is returned only for sharded cluster instances.
	ShardList []*DescribeDBInstancesOverviewResponseBodyDBInstancesShardList `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Repeated"`
	// The tags of the instance.
	Tags []*DescribeDBInstancesOverviewResponseBodyDBInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether password-free access over a VPC is enabled. Valid values:
	//
	// - **Open**: enabled
	//
	// - **Close**: disabled
	//
	// example:
	//
	// Open
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	// The zone of the instance.
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
	if s.MongosList != nil {
		for _, item := range s.MongosList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ShardList != nil {
		for _, item := range s.ShardList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// Test mongos node
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
	// Test shard node
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the shard node.
	//
	// example:
	//
	// d-bp1cac6f2083****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage capacity of the shard node, in GB.
	//
	// example:
	//
	// 10
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The number of read-only nodes in the shard. Valid values: **0*	- to **5**.
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
	// The tag key.
	//
	// - The key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// - The key can be up to 64 characters in length.
	//
	// - The key cannot be an empty string.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// - The value cannot start with `aliyun`, `acs`:, `http://`, or `https://`.
	//
	// - The value can be up to 128 characters in length.
	//
	// - The value can be an empty string.
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

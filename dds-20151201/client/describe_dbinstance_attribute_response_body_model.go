// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstances(v *DescribeDBInstanceAttributeResponseBodyDBInstances) *DescribeDBInstanceAttributeResponseBody
	GetDBInstances() *DescribeDBInstanceAttributeResponseBodyDBInstances
	SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The instance details.
	DBInstances *DescribeDBInstanceAttributeResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A935A8EE-A6CC-53DE-98BA-20ABAA7E632B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDBInstances() *DescribeDBInstanceAttributeResponseBodyDBInstances {
	return s.DBInstances
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstances(v *DescribeDBInstanceAttributeResponseBodyDBInstances) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstances struct {
	DBInstance []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstances) GetDBInstance() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstances) SetDBInstance(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) *DescribeDBInstanceAttributeResponseBodyDBInstances {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance struct {
	// Indicates whether performance burst is enabled for the ESSD AutoPL disk.
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The read and write throughput consumed by the instance.
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
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The details of the ConfigServer node.
	//
	// >  This parameter is returned if the instance is a sharded cluster instance.
	ConfigserverList *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList `json:"ConfigserverList,omitempty" xml:"ConfigserverList,omitempty" type:"Struct"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-02T07:43:59Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The minor version of the current database in the instance.
	//
	// example:
	//
	// 5.0.5-20220721143518_0
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// dds.mongo.mid
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp11483712c1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the orders generated for the instance. Valid values:
	//
	// 	- **all_completed**: All orders are being produced or complete.
	//
	// 	- **order_unpaid**: The instance has unpaid orders.
	//
	// 	- **order_wait_for_produce**: Orders are being delivered for production.
	//
	// >  The order production process includes the following steps: place an order, pay for an order, deliver an order for production, produce an order, and complete the production.
	//
	// 	- If an order is in the **order_wait_for_produce*	- state for a long time, an error occurs when the order is being delivered for production. The system will automatically retry.
	//
	// 	- The instance status changes only when the order is in the producing and complete state, such as changing configurations and running.
	//
	// example:
	//
	// all_completed
	DBInstanceOrderStatus *string `json:"DBInstanceOrderStatus,omitempty" xml:"DBInstanceOrderStatus,omitempty"`
	// Indicates whether release protection is enabled for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DBInstanceReleaseProtection *bool `json:"DBInstanceReleaseProtection,omitempty" xml:"DBInstanceReleaseProtection,omitempty"`
	// The status of the instance. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 10
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **replicate**: replica set instance
	//
	// 	- **sharding**: sharded cluster instance
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
	DestroyTime          *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	DisasterRecoveryInfo *string `json:"DisasterRecoveryInfo,omitempty" xml:"DisasterRecoveryInfo,omitempty"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The Key Management Service (KMS) key used for disk encryption.
	//
	// example:
	//
	// 07609cc3-3109-408f-a35e-c548e776da0b
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine of the instance.
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
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the subscription instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// >  This parameter is returned if the instance is a subscription instance.
	//
	// example:
	//
	// 2022-02-05T16:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the secondary zone 2 of the instance. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I
	//
	// 	- **cn-hongkong-b**: Hongkong Zone B
	//
	// 	- **cn-hongkong-c**: Hongkong Zone C
	//
	// 	- **cn-hongkong-d**: Hongkong Zone D
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C
	//
	// 	- **ap-southeast-1a**: Singapore Zone A
	//
	// 	- **ap-southeast-1b**: Singapore Zone B
	//
	// 	- **ap-southeast-1c**: Singapore Zone C
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C
	//
	// 	- **eu-central-1a**: Frankfurt Zone A
	//
	// 	- **eu-central-1b**: Frankfurt Zone B
	//
	// 	- **eu-central-1c**: Frankfurt Zone C
	//
	// >
	//
	// 	- This parameter is returned if the instance is a replica set or sharded cluster instance that runs MongoDB 4.4 or 5.0 and uses multi-zone deployment.
	//
	// 	- This parameter is returned only if you use the China site (aliyun.com).
	//
	// example:
	//
	// cn-hangzhou-h
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
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
	// 1
	KindCode *string `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// The date when the last downgrade operation was performed on the instance.
	//
	// example:
	//
	// 2022-01-08
	LastDowngradeTime *string `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	// The lock status of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before the instance is rolled back.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked after the storage space is exhausted.
	//
	// 	- **Released**: The instance is released.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The end time of the maintenance window. The time follows the ISO 8601 standard in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 03:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time follows the ISO 8601 standard in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The maximum number of connections to the instance.
	//
	// example:
	//
	// 500
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum IOPS of the instance.
	//
	// example:
	//
	// 1000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The maximum MBPS of the instance.
	//
	// example:
	//
	// 350
	MaxMBPS *int32 `json:"MaxMBPS,omitempty" xml:"MaxMBPS,omitempty"`
	// The details of the mongos node.
	//
	// >  This parameter is returned if the instance is a sharded cluster instance.
	MongosList *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Struct"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**: classic network
	//
	// 	- **VPC**: VPC
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The access protocol type of the instance. Valid values:
	//
	// 	- **mongodb**
	//
	// 	- **dynamodb**
	//
	// >  This parameter is returned if the instance is a sharded cluster instance.
	//
	// example:
	//
	// mongodb
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The provisioned performance of the ESSD AutoPL disk.
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of read-only nodes in the instance.
	//
	// example:
	//
	// 1
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The logical ID of the replica set instance.
	//
	// >  ApsaraDB for MongoDB does not support new instances of this type. This parameter applies only to previous-version replica set instances.
	//
	// example:
	//
	// bls-m****
	ReplacateId *string `json:"ReplacateId,omitempty" xml:"ReplacateId,omitempty"`
	// The name of the replica set instance.
	//
	// >  This parameter is returned if the instance is a replica set instance.
	//
	// example:
	//
	// mgset-10ace****
	ReplicaSetName *string `json:"ReplicaSetName,omitempty" xml:"ReplicaSetName,omitempty"`
	// The information of the replica set instance.
	//
	// >  This parameter is returned if the instance is a replica set instance.
	ReplicaSets *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Struct"`
	// The number of nodes in the instance.
	//
	// >  This parameter is returned if the instance is a replica set instance.
	//
	// example:
	//
	// 3
	ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// >  This parameter is returned only if you use the China site (aliyun.com).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the secondary zone 1 of the instance. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I
	//
	// 	- **cn-hongkong-b**: Hongkong Zone B
	//
	// 	- **cn-hongkong-c**: Hongkong Zone C
	//
	// 	- **cn-hongkong-d**: Hongkong Zone D
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C
	//
	// 	- **ap-southeast-1a**: Singapore Zone A
	//
	// 	- **ap-southeast-1b**: Singapore Zone B
	//
	// 	- **ap-southeast-1c**: Singapore Zone C
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C
	//
	// 	- **eu-central-1a**: Frankfurt Zone A
	//
	// 	- **eu-central-1b**: Frankfurt Zone B
	//
	// 	- **eu-central-1c**: Frankfurt Zone C
	//
	// >
	//
	// 	- This parameter is returned if the instance is a replica set or sharded cluster instance that runs MongoDB 4.4 or 5.0 and uses multi-zone deployment.
	//
	// 	- This parameter is returned only if you use the China site (aliyun.com).
	//
	// example:
	//
	// cn-hangzhou-i
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The details of the shard node.
	//
	// >  This parameter is returned if the instance is a sharded cluster instance.
	ShardList *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Struct"`
	// The storage engine of the instance.
	//
	// example:
	//
	// WiredTiger
	StorageEngine *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// **cloud_essd1**: ESSD PL1 **cloud_essd2**: ESSD PL2 **cloud_essd3**: ESSD PL3 **local_ssd**: local SSD **cloud_essd_dbfs_s**: DBFS disk
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The progress of data synchronization in percentage. When you are changing the configurations of the instance, you must synchronize the data of the instance. You can obtain the data synchronization progress based on the value returned for this parameter.
	//
	// example:
	//
	// 0.8
	SyncPercent *string `json:"SyncPercent,omitempty" xml:"SyncPercent,omitempty"`
	// The details of the instance tags.
	Tags *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// Indicates whether the cluster backup mode is enabled. Valid values:
	//
	// 	- **true**: The cluster backup mode is enabled.
	//
	// 	- **false**: The cluster backup mode is disabled.
	//
	// example:
	//
	// true
	UseClusterBackup *bool `json:"UseClusterBackup,omitempty" xml:"UseClusterBackup,omitempty"`
	// The instance ID.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// dds-bp11483712c1****
	VPCCloudInstanceIds *string `json:"VPCCloudInstanceIds,omitempty" xml:"VPCCloudInstanceIds,omitempty"`
	// The VPC ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// vsw-bp1oo2a7isyrb8igf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Indicates whether password-free access within the VPC is enabled. Valid values:
	//
	// 	- **Open**: Password-free access within the VPC is enabled.
	//
	// 	- **Close**: Password-free access within the VPC is disabled, and you must use a password for access.
	//
	// 	- **NotSupport**: Password-free access within the VPC is not supported.
	//
	// example:
	//
	// Open
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	// The ID of the zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetConfigserverList() *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList {
	return s.ConfigserverList
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceOrderStatus() *string {
	return s.DBInstanceOrderStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceReleaseProtection() *bool {
	return s.DBInstanceReleaseProtection
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetDisasterRecoveryInfo() *string {
	return s.DisasterRecoveryInfo
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetKindCode() *string {
	return s.KindCode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetLastDowngradeTime() *string {
	return s.LastDowngradeTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMaxMBPS() *int32 {
	return s.MaxMBPS
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetMongosList() *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList {
	return s.MongosList
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetReadonlyReplicas() *string {
	return s.ReadonlyReplicas
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetReplacateId() *string {
	return s.ReplacateId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetReplicaSetName() *string {
	return s.ReplicaSetName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetReplicaSets() *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets {
	return s.ReplicaSets
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetShardList() *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList {
	return s.ShardList
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetStorageEngine() *string {
	return s.StorageEngine
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetSyncPercent() *string {
	return s.SyncPercent
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetTags() *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags {
	return s.Tags
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetUseClusterBackup() *bool {
	return s.UseClusterBackup
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetVPCCloudInstanceIds() *string {
	return s.VPCCloudInstanceIds
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetBurstingEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetCapacityUnit(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetConfigserverList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ConfigserverList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceOrderStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceOrderStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceReleaseProtection(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceReleaseProtection = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDestroyTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetDisasterRecoveryInfo(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.DisasterRecoveryInfo = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEncrypted(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.Encrypted = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEncryptionKey(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetHiddenZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.HiddenZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetKindCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetLastDowngradeTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.LastDowngradeTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMaxMBPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MaxMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetMongosList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.MongosList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetProtocolType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetProvisionedIops(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReadonlyReplicas(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReadonlyReplicas = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplacateId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplacateId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicaSetName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicaSetName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicaSets(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicaSets = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetReplicationFactor(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ReplicationFactor = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetSecondaryZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetShardList(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ShardList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetStorageEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.StorageEngine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetSyncPercent(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.SyncPercent = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetTags(v *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetUseClusterBackup(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.UseClusterBackup = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVPCCloudInstanceIds(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VPCCloudInstanceIds = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetVpcAuthMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList struct {
	ConfigserverAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute `json:"ConfigserverAttribute,omitempty" xml:"ConfigserverAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) GetConfigserverAttribute() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	return s.ConfigserverAttribute
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) SetConfigserverAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList {
	s.ConfigserverAttribute = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute struct {
	// The endpoint of the Configserver node.
	//
	// example:
	//
	// dds-bp18b0934e7053e4-cs****.mongodb.rds.aliyuncs.com
	ConnectString *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	// The minor version of the current MongoDB kernel.
	//
	// example:
	//
	// mongodb_20230613_4.0.25
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	// The lock status of the Configserver node. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maximum number of connections to the Configserver node.
	//
	// example:
	//
	// 1000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum IOPS of the Configserver node.
	//
	// example:
	//
	// 1000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The instance type of the Configserver node.
	//
	// example:
	//
	// dds.cs.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The name of the Configserver node.
	//
	// example:
	//
	// testConfigserver
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the Configserver node.
	//
	// example:
	//
	// dds-bp11483712c1****-cs
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage capacity of the Configserver node. Unit: GB.
	//
	// example:
	//
	// 20
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The port number that is used to connect to the Configserver node.
	//
	// example:
	//
	// 3717
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the Configserver node. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetConnectString() *string {
	return s.ConnectString
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetConnectString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.ConnectString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetNodeStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList struct {
	MongosAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute `json:"MongosAttribute,omitempty" xml:"MongosAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) GetMongosAttribute() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	return s.MongosAttribute
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) SetMongosAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList {
	s.MongosAttribute = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute struct {
	// The endpoint of the mongos node.
	//
	// example:
	//
	// s-bp1d8c262a15****.mongodb.rds.aliyuncs.com
	ConnectSting *string `json:"ConnectSting,omitempty" xml:"ConnectSting,omitempty"`
	// The endpoint of the mongos node.
	//
	// example:
	//
	// s-bp1d8c262a15****.mongodb.rds.aliyuncs.com
	ConnectString *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	// The minor version of the current MongoDB kernel.
	//
	// example:
	//
	// mongodb_20220518_4.0.21
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	// The lock status of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maximum number of connections to the mongos node.
	//
	// example:
	//
	// 1000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum IOPS of the mongos node.
	//
	// example:
	//
	// 800
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The instance type of the mongos node.
	//
	// example:
	//
	// dds.mongos.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The name of the mongos node.
	//
	// example:
	//
	// mongos1
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the mongos node.
	//
	// example:
	//
	// s-bp1d8c262a15****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The port number that is used to connect to the mongos node.
	//
	// example:
	//
	// 3717
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the mongos node. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VPC ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// vsw-bp1vj604nj5a9zz74****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the mongos node.
	//
	// example:
	//
	// s-bp1d8c262a158****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetConnectSting() *string {
	return s.ConnectSting
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetConnectString() *string {
	return s.ConnectString
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetConnectSting(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.ConnectSting = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetConnectString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.ConnectString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) SetVpcCloudInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets struct {
	ReplicaSet []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) GetReplicaSet() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	return s.ReplicaSet
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) SetReplicaSet(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets {
	s.ReplicaSet = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet struct {
	// The endpoint of the node.
	//
	// example:
	//
	// dds-bp11483712c1****.mongodb.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The port number that is used to connect to the node.
	//
	// example:
	//
	// 3717
	ConnectionPort *string `json:"ConnectionPort,omitempty" xml:"ConnectionPort,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**: classic network
	//
	// 	- **VPC**: VPC
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- **Primary**
	//
	// 	- **Secondary**
	//
	// example:
	//
	// Primary
	ReplicaSetRole *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	// The instance ID.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// dds-bp11483712c1****
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	// The VPC ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is VPC.
	//
	// example:
	//
	// vpc-bp1jk5vwkcri27qme****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance.
	//
	// >  This parameter is returned if the network type of the instance is virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1jk5vwkcri27qme****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetConnectionPort() *string {
	return s.ConnectionPort
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetReplicaSetRole() *string {
	return s.ReplicaSetRole
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetVPCCloudInstanceId() *string {
	return s.VPCCloudInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetConnectionDomain(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetConnectionPort(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ConnectionPort = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetReplicaSetRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.ReplicaSetRole = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVPCCloudInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList struct {
	ShardAttribute []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute `json:"ShardAttribute,omitempty" xml:"ShardAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) GetShardAttribute() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	return s.ShardAttribute
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) SetShardAttribute(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList {
	s.ShardAttribute = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute struct {
	// The endpoint of the shard node.
	//
	// example:
	//
	// d-bp1af0680a9c6d3****.mongodb.rds.aliyuncs.com:****
	ConnectString *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	// The minor version of the current MongoDB kernel.
	//
	// example:
	//
	// mongodb_20230613_4.0.25
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	// The lock status of the shard node. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maximum number of connections to the shard node.
	//
	// example:
	//
	// 8000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum MBPS of the shard node.
	//
	// example:
	//
	// 350
	MaxDiskMbps *string `json:"MaxDiskMbps,omitempty" xml:"MaxDiskMbps,omitempty"`
	// The maximum IOPS of the shard node.
	//
	// example:
	//
	// 8000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The instance type of the shard node.
	//
	// example:
	//
	// dds.shard.mid
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The name of the shard node.
	//
	// example:
	//
	// testshard
	NodeDescription *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	// The ID of the shard node.
	//
	// example:
	//
	// d-bp16e09d9c5d****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage capacity of the shard node. Unit: GB.
	//
	// example:
	//
	// 10
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The port number that is used to connect to the shard node.
	//
	// example:
	//
	// 3717
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The number of read-only nodes in the shard node. Valid values: **0*	- to **5**. The value must be an integer.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ReplicaSetName   *string `json:"ReplicaSetName,omitempty" xml:"ReplicaSetName,omitempty"`
	// The status of the shard node. For more information, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetConnectString() *string {
	return s.ConnectString
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetMaxDiskMbps() *string {
	return s.MaxDiskMbps
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeDescription() *string {
	return s.NodeDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetReplicaSetName() *string {
	return s.ReplicaSetName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetConnectString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ConnectString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetMaxDiskMbps(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.MaxDiskMbps = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetNodeStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.NodeStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetPort(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetReadonlyReplicas(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ReadonlyReplicas = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetReplicaSetName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.ReplicaSetName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags struct {
	Tag []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) GetTag() []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag {
	return s.Tag
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) SetTag(v []*DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// api
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}

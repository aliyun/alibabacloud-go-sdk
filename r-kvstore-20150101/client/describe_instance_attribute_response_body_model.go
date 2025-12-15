// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstanceAttributeResponseBodyInstances) *DescribeInstanceAttributeResponseBody
	GetInstances() *DescribeInstanceAttributeResponseBodyInstances
	SetRequestId(v string) *DescribeInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAttributeResponseBody struct {
	// Details about the instances.
	Instances *DescribeInstanceAttributeResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CA40C261-EB72-4EDA-AC57-958722162595
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) GetInstances() *DescribeInstanceAttributeResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAttributeResponseBody) SetInstances(v *DescribeInstanceAttributeResponseBodyInstances) *DescribeInstanceAttributeResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstances struct {
	DBInstanceAttribute []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstances) GetDBInstanceAttribute() []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	return s.DBInstanceAttribute
}

func (s *DescribeInstanceAttributeResponseBodyInstances) SetDBInstanceAttribute(v []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) *DescribeInstanceAttributeResponseBodyInstances {
	s.DBInstanceAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstances) Validate() error {
	if s.DBInstanceAttribute != nil {
		for _, item := range s.DBInstanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute struct {
	// The architecture of the instance. Valid values:
	//
	// 	- **cluster**: cluster architecture
	//
	// 	- **standard**: standard architecture
	//
	// 	- **rwsplit**: read/write splitting architecture
	//
	// example:
	//
	// standard
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The retention period of audit logs. Unit: day. A value of 0 indicates that the audit log feature is disabled. For information about how to enable the feature, see [Enable the audit log feature](https://help.aliyun.com/document_detail/102015.html).
	//
	// example:
	//
	// 15
	AuditLogRetention *string `json:"AuditLogRetention,omitempty" xml:"AuditLogRetention,omitempty"`
	// Indicates whether a secondary zone is automatically allocated.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoSecondaryZone *bool `json:"AutoSecondaryZone,omitempty" xml:"AutoSecondaryZone,omitempty"`
	// The availability metric of the current month.
	//
	// example:
	//
	// 100%
	AvailabilityValue *string `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	// The earliest point in time to which data can be restored. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// >
	//
	// 	- This parameter is returned only when the data flashback feature is enabled for the instance. For more information, see [Restore data to a point in time by using the data flashback feature](https://help.aliyun.com/document_detail/148479.html).
	//
	// 	- When you call the [RestoreInstance](https://help.aliyun.com/document_detail/473824.html) operation to implement data flashback, you can obtain the earliest point in time for data flashback from the return value of this parameter and set the **RestoreTime*	- parameter to this point in time.
	//
	// example:
	//
	// 2021-07-06T05:49:55Z
	BackupLogStartTime *string `json:"BackupLogStartTime,omitempty" xml:"BackupLogStartTime,omitempty"`
	// The bandwidth of the instance. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The storage capacity of the instance. Unit: MB.
	//
	// example:
	//
	// 1024
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
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
	// This parameter is returned only when the instance is in a cloud box.
	//
	// example:
	//
	// cloudbox
	CloudType     *string `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	ComputingType *string `json:"ComputingType,omitempty" xml:"ComputingType,omitempty"`
	// The parameter configurations of the instance in the JSON format. For more information, see [Parameter descriptions](https://help.aliyun.com/document_detail/43885.html). You can use the [DescribeAuditLogConfig](https://help.aliyun.com/document_detail/473830.html) operation to query audit log configurations.
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// r-bp1d72gwl41z7f****.redis.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The maximum number of connections supported by the instance.
	//
	// example:
	//
	// 10000
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-06T10:42:03Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the subscription expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-04-06T10:42:03Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance. The return value is **Redis**.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Valid values: **2.8**, **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the distributed instance to which the instance belongs.
	//
	// >  This parameter is returned only when the Tair (Redis OSS-compatible) instance is a child instance of a distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// Indicates whether your Alibaba Cloud account has pending orders for renewal and configuration change. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	HasRenewChangeOrder *string `json:"HasRenewChangeOrder,omitempty" xml:"HasRenewChangeOrder,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/107984.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1d72gwl41z7f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the release protection feature is enabled for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	InstanceReleaseProtection *bool `json:"InstanceReleaseProtection,omitempty" xml:"InstanceReleaseProtection,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **Normal**: The instance is normal.
	//
	// 	- **Creating**: The instance is being created.
	//
	// 	- **Changing**: The configurations of the instance are being changed.
	//
	// 	- **Inactive**: The instance is disabled.
	//
	// 	- **Flushing**: The instance is being released.
	//
	// 	- **Released**: The instance is released.
	//
	// 	- **Transforming**: The billing method of the instance is being changed.
	//
	// 	- **Unavailable**: The instance is unavailable.
	//
	// 	- **Error**: The instance failed to be created.
	//
	// 	- **Migrating**: The instance is being migrated.
	//
	// 	- **BackupRecovering**: The instance is being restored from a backup.
	//
	// 	- **MinorVersionUpgrading**: The minor version of the instance is being updated.
	//
	// 	- **NetworkModifying**: The network type of the instance is being changed.
	//
	// 	- **SSLModifying**: The SSL configurations of the instance are being changed.
	//
	// 	- **MajorVersionUpgrading**: The major version of the instance is being upgraded. The instance remains accessible during the upgrade.
	//
	// >  For more information about instance states, see [Instance states and impacts](https://help.aliyun.com/document_detail/200740.html).
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **Tair**
	//
	// 	- **Redis**
	//
	// 	- **Memcache**
	//
	// example:
	//
	// Redis
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the order for instance configuration change has reached the final state. Valid values:
	//
	// 	- **true**: The configuration change has been completed or has not been performed.
	//
	// 	- **false**: The configurations of the instance are being changed.
	//
	// example:
	//
	// true
	IsOrderCompleted *bool `json:"IsOrderCompleted,omitempty" xml:"IsOrderCompleted,omitempty"`
	// Indicates whether the instance is managed by ApsaraDB RDS. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsRds *bool `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	// Indicates whether the transparent data encryption (TDE) feature is supported for the instance. Valid values:
	//
	// 	- **true**: This feature is supported only for DRAM-based classic instances.
	//
	// 	- **false**: This feature is not supported.
	//
	// example:
	//
	// true
	IsSupportTDE *bool `json:"IsSupportTDE,omitempty" xml:"IsSupportTDE,omitempty"`
	// The end time of the maintenance window. The time is in the *HH:mmZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in the *HH:mmZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **CLASSIC**
	//
	// 	- **VPC**
	//
	// example:
	//
	// CLASSIC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The node type. Valid values:
	//
	// 	- **double**: The instance contains a master node and a replica node.
	//
	// 	- **single**: The instance is a standalone instance.
	//
	// example:
	//
	// double
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The plan type. Valid values:
	//
	// 	- **standard**: standard plan.
	//
	// 	- **customized**: custom plan. This plan type is phased out.
	//
	// example:
	//
	// standard
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The service port of the instance.
	//
	// example:
	//
	// 6379
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the instance.
	//
	// >  This parameter is not returned when the instance is deployed in the classic network.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The expected maximum queries per second (QPS).
	//
	// example:
	//
	// 100000
	QPS *int64 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	// The number of read replicas. This parameter is available only for read/write splitting instances that use cloud disks.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// If the instance is a cluster instance that uses cloud disks, this parameter indicates the actual instance type of individual shards in the instance. The InstanceClass parameter indicates the virtual instance type.
	//
	// >  To query fees for instances of the instance type, you can specify the instance type that is returned by this parameter in the [DescribePrice](https://help.aliyun.com/document_detail/473807.html) operation.
	//
	// example:
	//
	// tair.rdb.with.proxy.1g
	RealInstanceClass *string `json:"RealInstanceClass,omitempty" xml:"RealInstanceClass,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone.
	//
	// >  The **ReplicaCount*	- and **SlaveReplicaCount*	- parameters are applicable only to cloud-native instances. If the instance is a cluster instance, the preceding parameters indicate the number of replica nodes **per node*	- in the primary and secondary zones of the instance.
	//
	// example:
	//
	// 1
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The ID of the replica node.
	//
	// example:
	//
	// bls-awxxxxxxxxxxxxx
	ReplicaId *string `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	// The architecture of the replica node. Valid values:
	//
	// 	- **master-slave**: the standard master-replica architecture.
	//
	// 	- **cluster**: the cluster architecture, which includes the read/write splitting instances and cluster instances.
	//
	// example:
	//
	// master-slave
	ReplicationMode *string `json:"ReplicationMode,omitempty" xml:"ReplicationMode,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the secondary zone.
	//
	// >  This parameter is returned only if the instance has a secondary zone ID.
	//
	// example:
	//
	// cn-hongkong-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The number of shards. This parameter is available only for instances that are purchased on the China site (aliyun.com).
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The number of read replicas in the secondary zone. This parameter is returned only after read/write splitting is enabled for the instance across multiple zones.
	//
	// example:
	//
	// 2
	SlaveReadOnlyCount *int64 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replica nodes in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// The storage capacity of the cloud disk.
	//
	// example:
	//
	// 50
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type.
	//
	// example:
	//
	// essd_pl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Details about the tags.
	Tags *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Indicates whether password authentication is enabled. Valid values:
	//
	// 	- **Open**: Password authentication is enabled.
	//
	// 	- **Close**: Password authentication is disabled and [password-free access](https://help.aliyun.com/document_detail/85168.html) is enabled.
	//
	// example:
	//
	// Open
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	// The ID of the instance in the VPC.
	//
	// example:
	//
	// r-bp1d72gwl41z7f****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hongkong-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The deployment type of the instance. Valid values:
	//
	// 	- **singlezone**: The instance is deployed in a single zone.
	//
	// 	- **doublezone**: The instance is deployed in two zones of the same region.
	//
	// example:
	//
	// singlezone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetAuditLogRetention() *string {
	return s.AuditLogRetention
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetAutoSecondaryZone() *bool {
	return s.AutoSecondaryZone
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetAvailabilityValue() *string {
	return s.AvailabilityValue
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetBackupLogStartTime() *string {
	return s.BackupLogStartTime
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetCloudType() *string {
	return s.CloudType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetComputingType() *string {
	return s.ComputingType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetConfig() *string {
	return s.Config
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetConnections() *int64 {
	return s.Connections
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetEngine() *string {
	return s.Engine
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetHasRenewChangeOrder() *string {
	return s.HasRenewChangeOrder
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceReleaseProtection() *bool {
	return s.InstanceReleaseProtection
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetIsOrderCompleted() *bool {
	return s.IsOrderCompleted
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetIsRds() *bool {
	return s.IsRds
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetIsSupportTDE() *bool {
	return s.IsSupportTDE
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetPort() *int64 {
	return s.Port
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetQPS() *int64 {
	return s.QPS
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetReadOnlyCount() *int32 {
	return s.ReadOnlyCount
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetRealInstanceClass() *string {
	return s.RealInstanceClass
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetReplicaId() *string {
	return s.ReplicaId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetReplicationMode() *string {
	return s.ReplicationMode
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetSlaveReadOnlyCount() *int64 {
	return s.SlaveReadOnlyCount
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetStorage() *string {
	return s.Storage
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetTags() *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags {
	return s.Tags
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetArchitectureType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetAuditLogRetention(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.AuditLogRetention = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetAutoSecondaryZone(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.AutoSecondaryZone = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetBackupLogStartTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.BackupLogStartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetBandwidth(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCapacity(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Capacity = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetChargeType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCloudType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.CloudType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetComputingType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ComputingType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConfig(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Config = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConnectionDomain(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConnections(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Connections = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCreateTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEndTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEngine(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEngineVersion(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetGlobalInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetHasRenewChangeOrder(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.HasRenewChangeOrder = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceClass(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceName(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceReleaseProtection(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceReleaseProtection = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsOrderCompleted(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsOrderCompleted = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsRds(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsRds = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsSupportTDE(v bool) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsSupportTDE = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetNetworkType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetNodeType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.NodeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPackageType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.PackageType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPort(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPrivateIp(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetQPS(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.QPS = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReadOnlyCount(v int32) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReadOnlyCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetRealInstanceClass(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.RealInstanceClass = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicaCount(v int32) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicaCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicaId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicaId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicationMode(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicationMode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetResourceGroupId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetSecondaryZoneId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetSecurityIPList(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetShardCount(v int32) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ShardCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetSlaveReadOnlyCount(v int64) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetSlaveReplicaCount(v int32) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.SlaveReplicaCount = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetStorage(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Storage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetStorageType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetTags(v *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVSwitchId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcAuthMode(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcAuthMode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcCloudInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetZoneId(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetZoneType(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ZoneType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags struct {
	Tag []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) GetTag() []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	return s.Tag
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) SetTag(v []*DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags {
	s.Tag = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// tagkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetKey(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetValue(v string) *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) Validate() error {
	return dara.Validate(s)
}

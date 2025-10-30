// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstanceAttributeResponseBodyItems) *DescribeDBInstanceAttributeResponseBody
	GetItems() *DescribeDBInstanceAttributeResponseBodyItems
	SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The queried instance.
	Items *DescribeDBInstanceAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 5E6EDEB8-D73E-5F2D-B948-86C8AEB05A68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) GetItems() *DescribeDBInstanceAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceAttributeResponseBody) SetItems(v *DescribeDBInstanceAttributeResponseBodyItems) *DescribeDBInstanceAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItems) GetDBInstanceAttribute() []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	return s.DBInstanceAttribute
}

func (s *DescribeDBInstanceAttributeResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *DescribeDBInstanceAttributeResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItems) Validate() error {
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

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute struct {
	// Queries the current instance availability status, in percentage (%).
	//
	// > This parameter is only applicable to instances in the storage reserved mode.
	//
	// example:
	//
	// 100.0%
	AvailabilityValue *string `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	CacheStorageSize  *string `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	// Access mode, with the following values:
	//
	// - **Performance**: Standard access mode.
	//
	// - **Safty**: High-security access mode.
	//
	// - **LVS**: LVS link mode.
	//
	// example:
	//
	// LVS
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// Instance connection address.
	//
	// example:
	//
	// gp-bp13ue79qk8y1****-master.gpdb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Minor version number of the kernel.
	//
	// example:
	//
	// mm.v6.3.10.1-202207141918
	CoreVersion *string `json:"CoreVersion,omitempty" xml:"CoreVersion,omitempty"`
	// Number of CPU cores for the compute node, unit: Core.
	//
	// example:
	//
	// 2
	CpuCores *int32 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// Number of CPU cores per node.
	//
	// > This parameter is only applicable to instances in the storage reserved mode.
	//
	// example:
	//
	// 0
	CpuCoresPerNode *int32 `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty"`
	// Instance creation time.
	//
	// example:
	//
	// 2022-08-11T09:16:26Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Instance series, with the following values:
	//
	// - **Basic**: Basic Edition.
	//
	// - **HighAvailability**: High Availability Edition.
	//
	// example:
	//
	// HighAvailability
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// Instance specification.
	//
	// > This parameter is only applicable to reserved storage mode instances.
	//
	// example:
	//
	// gpdb.group.segsdx1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// Instance family, with the following values:
	//
	// - **s**: Shared type.
	//
	// - **x**: General type.
	//
	// - **d**: Dedicated package.
	//
	// - **h**: Dedicated physical machine.
	//
	// example:
	//
	// x
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// Number of CPU cores.
	//
	// example:
	//
	// 2
	DBInstanceCpuCores *int32 `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty"`
	// Instance description.
	//
	// example:
	//
	// gp-bp13ue79qk8y1****
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// Maximum BPS (disk throughput) of the compute group, in Mbps.
	//
	// > This parameter is only applicable to reserved storage mode instances.
	//
	// example:
	//
	// 0
	DBInstanceDiskMBPS *int64 `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty"`
	// Number of compute groups.
	//
	// > This parameter is only applicable to reserved storage mode instances.
	//
	// example:
	//
	// 0
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// gp-bp13ue79qk8y1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Memory of the compute node.
	//
	// > The unit for storage-reserved mode is MB; for Serverless and storage-elastic modes, it is GB.
	//
	// example:
	//
	// 16
	DBInstanceMemory *int64 `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	// Instance resource type, with the following values:
	//
	// - **Serverless**: Serverless mode.
	//
	// - **StorageElastic**: Storage elastic mode.
	//
	// - **Classic**: Storage reserved mode.
	//
	// example:
	//
	// StorageElastic
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// This parameter is deprecated and will not return any value.
	//
	// example:
	//
	// null
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// Instance status. For more details, see the supplementary explanation of the DBInstanceStatus parameter.
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// Maximum storage space of a single replica, in GB.
	//
	// example:
	//
	// 50
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The deployment mode.
	//
	// example:
	//
	// single
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// Encryption key.
	//
	// > This parameter is returned only for instances with disk encryption enabled.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-************
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// Encryption type, with the following value:
	//
	// - **CloudDisk**: Cloud disk encryption.
	//
	// > This parameter is returned only for instances with cloud disk encryption.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// Database engine.
	//
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Database version.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Instance expiration time (in UTC).
	//
	// > The expiration time for pay-as-you-go instances is `2999-09-08T16:00:00Z`.
	//
	// example:
	//
	// 2999-09-08T16:00:00Z
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GraphEngineStatus *string `json:"GraphEngineStatus,omitempty" xml:"GraphEngineStatus,omitempty"`
	// Compute group machine type, with the following values:
	//
	// - **0**: SSD
	//
	// - **1**: HDD
	//
	// > This parameter applies only to storage-reserved mode instances.
	//
	// example:
	//
	// 0
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// Idle release waiting time. Unit: seconds.
	//
	// > This parameter is returned only for instances in the Serverless automatic scheduling mode.
	//
	// example:
	//
	// 600
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// Instance network type, with the following values:
	//
	// - **Classic**: Classic network.
	//
	// - **VPC**: VPC network.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceSpec        *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// Lock mode, with the following values:
	//
	// - **Unlock**: Normal.
	//
	// - **ManualLock**: Manually triggered lock.
	//
	// - **LockByExpiration**: Automatically locked when the instance expires.
	//
	// - **LockByRestoration**: Automatically locked before the instance rolls back.
	//
	// - **LockByDiskQuota**: Automatically locked when the instance space is full.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// This parameter is deprecated and will not return any value.
	//
	// example:
	//
	// null
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// Maintenance end time.
	//
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// Maintenance start time.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The specifications of AI coordinator node resources of the instance. If the coordinator nodes of the instance are not AI nodes, null is returned.
	//
	// example:
	//
	// ADB.AIMedium.2
	MasterAISpec *string `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	// Master resources.
	//
	// example:
	//
	// 4
	MasterCU *int32 `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	// Number of Master nodes.
	//
	// example:
	//
	// 1
	MasterNodeNum *int32 `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	// Maximum number of concurrent connections for the instance.
	//
	// > This parameter is only applicable to reserved storage mode instances.
	//
	// example:
	//
	// 500
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// Memory size per replica, see the **MemoryUnit*	- parameter for the unit.
	//
	// > This parameter is only applicable to instances in the storage reserved mode.
	//
	// example:
	//
	// 0
	MemoryPerNode *int32 `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty"`
	// Memory size of the compute node.
	//
	// > The unit is MB for the storage reserved mode; GB for Serverless and storage elastic modes.
	//
	// example:
	//
	// 16
	MemorySize *int64 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// Memory unit.
	//
	// > This parameter is only applicable to reserved storage mode instances.
	//
	// example:
	//
	// GB
	MemoryUnit *string `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty"`
	// Minor version of the kernel.
	//
	// example:
	//
	// 6.3.10.1-202207141918
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// Billing type, with the following values:
	//
	// - **Postpaid**: Pay-as-you-go.
	//
	// - **Prepaid**: Subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Instance port number.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The service type.
	//
	// example:
	//
	// standard
	ProdType *string `json:"ProdType,omitempty" xml:"ProdType,omitempty"`
	// This parameter has been deprecated and will not return a value.
	//
	// example:
	//
	// null
	ReadDelayTime *string `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// ID of the resource group where the instance is located.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Instance running time.
	//
	// example:
	//
	// 4 days 22:58:55
	RunningTime *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// This parameter is deprecated and will not return any value.
	//
	// example:
	//
	// null
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// Performance Level (PL), currently only **PL1*	- is supported.
	//
	// example:
	//
	// PL1
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// Number of Segment nodes.
	//
	// > This parameter applies only to instances in the storage elastic mode and Serverless manual scheduling mode.
	//
	// example:
	//
	// 4
	SegNodeNum *int32 `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The specifications of AI compute node resources of the instance. If the compute nodes of the instance are not AI nodes, null is returned.
	//
	// example:
	//
	// ADB.AIMedium.2
	SegmentAISpec *string `json:"SegmentAISpec,omitempty" xml:"SegmentAISpec,omitempty"`
	// Number of compute groups.
	//
	// > This parameter applies only to storage-reserved mode instances.
	//
	// example:
	//
	// 0
	SegmentCounts *int32 `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty"`
	// The mode of the Serverless instance, with the following values:
	//
	// - **Manual**: Manual scheduling.
	//
	// - **Auto**: Automatic scheduling.
	//
	// > This parameter is returned only for Serverless mode instances.
	//
	// example:
	//
	// Auto
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// Compute resource threshold. Unit: ACU.
	//
	// > This parameter is returned only for instances in the Serverless automatic scheduling mode.
	//
	// example:
	//
	// 32
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The secondary zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The time when the instance started running.
	//
	// example:
	//
	// 2022-08-11T09:26:43Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Storage size per replica, see the **StorageUnit*	- parameter for units.
	//
	// > This parameter applies only to storage-reserved mode instances.
	//
	// example:
	//
	// 0
	StoragePerNode *int32 `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty"`
	// Storage space size, unit: GB.
	//
	// example:
	//
	// 50
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Storage type, with the following values:
	//
	// - **cloud_essd**: ESSD cloud disk.
	//
	// - **cloud_efficiency**: Efficient cloud disk.
	//
	// > This parameter is only applicable to instances in the storage elastic mode.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Storage unit, with the following values:
	//
	// - **GB SSD**
	//
	// - **TB SSD**
	//
	// - **GB HDD**
	//
	// > This parameter is only applicable to instances in the storage reserved mode.
	//
	// example:
	//
	// GB SSD
	StorageUnit *string `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty"`
	// Indicates whether backup recovery is supported, with the following values:
	//
	// - **true**: Backup recovery is supported.
	//
	// - **false**: Backup recovery is not supported.
	//
	// example:
	//
	// true
	SupportRestore *bool `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	// Tag key-value pairs.
	Tags *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// vSwitch ID.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Indicates whether vector engine optimization is enabled. The values are as follows:
	//
	// - **enabled**: Indicates that vector engine optimization is enabled.
	//
	// - **disabled**: Indicates that vector engine optimization is disabled.
	//
	// example:
	//
	// enabled
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetAvailabilityValue() *string {
	return s.AvailabilityValue
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCacheStorageSize() *string {
	return s.CacheStorageSize
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCoreVersion() *string {
	return s.CoreVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCpuCoresPerNode() *int32 {
	return s.CpuCoresPerNode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceCategory() *string {
	return s.DBInstanceCategory
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceClassType() *string {
	return s.DBInstanceClassType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceCpuCores() *int32 {
	return s.DBInstanceCpuCores
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceDiskMBPS() *int64 {
	return s.DBInstanceDiskMBPS
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceGroupCount() *string {
	return s.DBInstanceGroupCount
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceMemory() *int64 {
	return s.DBInstanceMemory
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceMode() *string {
	return s.DBInstanceMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDeployMode() *string {
	return s.DeployMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetGraphEngineStatus() *string {
	return s.GraphEngineStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetHostType() *string {
	return s.HostType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIdleTime() *int32 {
	return s.IdleTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMasterAISpec() *string {
	return s.MasterAISpec
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMasterCU() *int32 {
	return s.MasterCU
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMasterNodeNum() *int32 {
	return s.MasterNodeNum
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMemoryPerNode() *int32 {
	return s.MemoryPerNode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMemorySize() *int64 {
	return s.MemorySize
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMemoryUnit() *string {
	return s.MemoryUnit
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetProdType() *string {
	return s.ProdType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetReadDelayTime() *string {
	return s.ReadDelayTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetRunningTime() *string {
	return s.RunningTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSegDiskPerformanceLevel() *string {
	return s.SegDiskPerformanceLevel
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSegNodeNum() *int32 {
	return s.SegNodeNum
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSegmentAISpec() *string {
	return s.SegmentAISpec
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSegmentCounts() *int32 {
	return s.SegmentCounts
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetServerlessMode() *string {
	return s.ServerlessMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetServerlessResource() *int32 {
	return s.ServerlessResource
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStoragePerNode() *int32 {
	return s.StoragePerNode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetStorageUnit() *string {
	return s.StorageUnit
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSupportRestore() *bool {
	return s.SupportRestore
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTags() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags {
	return s.Tags
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVectorConfigurationStatus() *string {
	return s.VectorConfigurationStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCacheStorageSize(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CacheStorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCoreVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CoreVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCpuCoresPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CpuCoresPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCategory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCategory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCpuCores(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDiskMBPS(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDiskMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceGroupCount(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDeployMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DeployMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetGraphEngineStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.GraphEngineStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetHostType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIdleTime(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IdleTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceSpec(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterAISpec(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterAISpec = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterCU(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterCU = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryPerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemorySize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMemoryUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MemoryUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetProdType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ProdType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetRunningTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.RunningTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegDiskPerformanceLevel(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegNodeNum(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegmentAISpec(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegmentAISpec = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSegmentCounts(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SegmentCounts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessResource(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessResource = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStandbyZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStartTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStoragePerNode(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StoragePerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetStorageUnit(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.StorageUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSupportRestore(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportRestore = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVectorConfigurationStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VectorConfigurationStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) GetTag() []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	return s.Tag
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags) Validate() error {
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

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag struct {
	// Tag key.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTagsTag) Validate() error {
	return dara.Validate(s)
}

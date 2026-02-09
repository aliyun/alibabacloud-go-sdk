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
	AvailabilityValue         *string                                                              `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	CacheStorageSize          *string                                                              `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	ConnectionMode            *string                                                              `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	ConnectionString          *string                                                              `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CoreVersion               *string                                                              `json:"CoreVersion,omitempty" xml:"CoreVersion,omitempty"`
	CpuCores                  *int32                                                               `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CpuCoresPerNode           *int32                                                               `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty"`
	CreationTime              *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceCategory        *string                                                              `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	DBInstanceClass           *string                                                              `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceClassType       *string                                                              `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceCpuCores        *int32                                                               `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty"`
	DBInstanceDescription     *string                                                              `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceDiskMBPS        *int64                                                               `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty"`
	DBInstanceGroupCount      *string                                                              `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	DBInstanceId              *string                                                              `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMemory          *int64                                                               `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	DBInstanceMode            *string                                                              `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	DBInstanceNetType         *string                                                              `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStatus          *string                                                              `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage         *int64                                                               `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DeployMode                *string                                                              `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	EncryptionKey             *string                                                              `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType            *string                                                              `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	Engine                    *string                                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion             *string                                                              `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime                *string                                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GraphEngineStatus         *string                                                              `json:"GraphEngineStatus,omitempty" xml:"GraphEngineStatus,omitempty"`
	HostType                  *string                                                              `json:"HostType,omitempty" xml:"HostType,omitempty"`
	IdleTime                  *int32                                                               `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	InstanceNetworkType       *string                                                              `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstanceSpec              *string                                                              `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	LockMode                  *string                                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime           *string                                                              `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime         *string                                                              `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MasterAISpec              *string                                                              `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	MasterCU                  *int32                                                               `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	MasterNodeNum             *int32                                                               `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	MaxConnections            *int32                                                               `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MemoryPerNode             *int32                                                               `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty"`
	MemorySize                *int64                                                               `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	MemoryUnit                *string                                                              `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty"`
	MinorVersion              *string                                                              `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	PayType                   *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                      *string                                                              `json:"Port,omitempty" xml:"Port,omitempty"`
	ProdType                  *string                                                              `json:"ProdType,omitempty" xml:"ProdType,omitempty"`
	ReadDelayTime             *string                                                              `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	RegionId                  *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId           *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RunningTime               *string                                                              `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SecurityIPList            *string                                                              `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SegDiskPerformanceLevel   *string                                                              `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	SegNodeNum                *int32                                                               `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	SegmentAISpec             *string                                                              `json:"SegmentAISpec,omitempty" xml:"SegmentAISpec,omitempty"`
	SegmentCounts             *int32                                                               `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty"`
	ServerlessMode            *string                                                              `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	ServerlessResource        *int32                                                               `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	StandbyZoneId             *string                                                              `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	StartTime                 *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StoragePerNode            *int32                                                               `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty"`
	StorageSize               *int64                                                               `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType               *string                                                              `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUnit               *string                                                              `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty"`
	SupportRestore            *bool                                                                `json:"SupportRestore,omitempty" xml:"SupportRestore,omitempty"`
	Tags                      *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                 *string                                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VectorConfigurationStatus *string                                                              `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	VpcId                     *string                                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                    *string                                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

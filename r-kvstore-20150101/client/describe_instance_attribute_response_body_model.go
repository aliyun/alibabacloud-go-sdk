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
	ArchitectureType          *string                                                                `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AuditLogRetention         *string                                                                `json:"AuditLogRetention,omitempty" xml:"AuditLogRetention,omitempty"`
	AutoSecondaryZone         *bool                                                                  `json:"AutoSecondaryZone,omitempty" xml:"AutoSecondaryZone,omitempty"`
	AvailabilityValue         *string                                                                `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	BackupLogStartTime        *string                                                                `json:"BackupLogStartTime,omitempty" xml:"BackupLogStartTime,omitempty"`
	Bandwidth                 *int64                                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity                  *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType                *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CloudType                 *string                                                                `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	ComputingType             *string                                                                `json:"ComputingType,omitempty" xml:"ComputingType,omitempty"`
	Config                    *string                                                                `json:"Config,omitempty" xml:"Config,omitempty"`
	ConnectionDomain          *string                                                                `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	Connections               *int64                                                                 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime                *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime                   *string                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Engine                    *string                                                                `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion             *string                                                                `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	GlobalInstanceId          *string                                                                `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	HasRenewChangeOrder       *string                                                                `json:"HasRenewChangeOrder,omitempty" xml:"HasRenewChangeOrder,omitempty"`
	InstanceClass             *string                                                                `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId                *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName              *string                                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceReleaseProtection *bool                                                                  `json:"InstanceReleaseProtection,omitempty" xml:"InstanceReleaseProtection,omitempty"`
	InstanceStatus            *string                                                                `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType              *string                                                                `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsOrderCompleted          *bool                                                                  `json:"IsOrderCompleted,omitempty" xml:"IsOrderCompleted,omitempty"`
	IsRds                     *bool                                                                  `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	IsSupportTDE              *bool                                                                  `json:"IsSupportTDE,omitempty" xml:"IsSupportTDE,omitempty"`
	MaintainEndTime           *string                                                                `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime         *string                                                                `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkType               *string                                                                `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeType                  *string                                                                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	PackageType               *string                                                                `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Port                      *int64                                                                 `json:"Port,omitempty" xml:"Port,omitempty"`
	PrivateIp                 *string                                                                `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	QPS                       *int64                                                                 `json:"QPS,omitempty" xml:"QPS,omitempty"`
	ReadOnlyCount             *int32                                                                 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	RealInstanceClass         *string                                                                `json:"RealInstanceClass,omitempty" xml:"RealInstanceClass,omitempty"`
	RegionId                  *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaCount              *int32                                                                 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ReplicaId                 *string                                                                `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	ReplicationMode           *string                                                                `json:"ReplicationMode,omitempty" xml:"ReplicationMode,omitempty"`
	ResourceGroupId           *string                                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryZoneId           *string                                                                `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityIPList            *string                                                                `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	ShardCount                *int32                                                                 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	SlaveReadOnlyCount        *int64                                                                 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	SlaveReplicaCount         *int32                                                                 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	Storage                   *string                                                                `json:"Storage,omitempty" xml:"Storage,omitempty"`
	StorageType               *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                      *DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                 *string                                                                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcAuthMode               *string                                                                `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	VpcCloudInstanceId        *string                                                                `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                     *string                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                    *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneType                  *string                                                                `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

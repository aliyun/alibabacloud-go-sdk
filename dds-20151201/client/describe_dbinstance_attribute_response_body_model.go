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
	if s.DBInstances != nil {
		if err := s.DBInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DBInstance != nil {
		for _, item := range s.DBInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance struct {
	BurstingEnabled             *bool                                                                         `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	CapacityUnit                *string                                                                       `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	ChargeType                  *string                                                                       `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ColdDataEnabled             *bool                                                                         `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	ConfigserverList            *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverList `json:"ConfigserverList,omitempty" xml:"ConfigserverList,omitempty" type:"Struct"`
	CreationTime                *string                                                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CurrentKernelVersion        *string                                                                       `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	DBInstanceClass             *string                                                                       `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription       *string                                                                       `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId                *string                                                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceOrderStatus       *string                                                                       `json:"DBInstanceOrderStatus,omitempty" xml:"DBInstanceOrderStatus,omitempty"`
	DBInstanceReleaseProtection *bool                                                                         `json:"DBInstanceReleaseProtection,omitempty" xml:"DBInstanceReleaseProtection,omitempty"`
	DBInstanceStatus            *string                                                                       `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage           *int32                                                                        `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceType              *string                                                                       `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DestroyTime                 *string                                                                       `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	DisasterRecoveryInfo        *string                                                                       `json:"DisasterRecoveryInfo,omitempty" xml:"DisasterRecoveryInfo,omitempty"`
	Encrypted                   *bool                                                                         `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	EncryptionKey               *string                                                                       `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	Engine                      *string                                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion               *string                                                                       `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime                  *string                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HiddenZoneId                *string                                                                       `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	KindCode                    *string                                                                       `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	LastDowngradeTime           *string                                                                       `json:"LastDowngradeTime,omitempty" xml:"LastDowngradeTime,omitempty"`
	LockMode                    *string                                                                       `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaintainEndTime             *string                                                                       `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime           *string                                                                       `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaxConnections              *int32                                                                        `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOPS                     *int32                                                                        `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	MaxMBPS                     *int32                                                                        `json:"MaxMBPS,omitempty" xml:"MaxMBPS,omitempty"`
	MongosList                  *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosList       `json:"MongosList,omitempty" xml:"MongosList,omitempty" type:"Struct"`
	NetworkType                 *string                                                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ProtocolType                *string                                                                       `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ProvisionedIops             *int64                                                                        `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	ReadonlyReplicas            *string                                                                       `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	RegionId                    *string                                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplacateId                 *string                                                                       `json:"ReplacateId,omitempty" xml:"ReplacateId,omitempty"`
	ReplicaSetName              *string                                                                       `json:"ReplicaSetName,omitempty" xml:"ReplicaSetName,omitempty"`
	ReplicaSets                 *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSets      `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Struct"`
	ReplicationFactor           *string                                                                       `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceGroupId             *string                                                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// mdb.shard.2x.xlarge.d
	SearchNodeClass *string `json:"SearchNodeClass,omitempty" xml:"SearchNodeClass,omitempty"`
	// example:
	//
	// 2
	SearchNodeCount *int32 `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
	// example:
	//
	// 20
	SearchNodeStorage   *int32                                                                 `json:"SearchNodeStorage,omitempty" xml:"SearchNodeStorage,omitempty"`
	SecondaryZoneId     *string                                                                `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	ShardList           *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardList `json:"ShardList,omitempty" xml:"ShardList,omitempty" type:"Struct"`
	StorageEngine       *string                                                                `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	StorageType         *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SyncPercent         *string                                                                `json:"SyncPercent,omitempty" xml:"SyncPercent,omitempty"`
	Tags                *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UseClusterBackup    *bool                                                                  `json:"UseClusterBackup,omitempty" xml:"UseClusterBackup,omitempty"`
	VPCCloudInstanceIds *string                                                                `json:"VPCCloudInstanceIds,omitempty" xml:"VPCCloudInstanceIds,omitempty"`
	VPCId               *string                                                                `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId           *string                                                                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcAuthMode         *string                                                                `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
	ZoneId              *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetSearchNodeClass() *string {
	return s.SearchNodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetSearchNodeCount() *int32 {
	return s.SearchNodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) GetSearchNodeStorage() *int32 {
	return s.SearchNodeStorage
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetColdDataEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.ColdDataEnabled = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetSearchNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.SearchNodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetSearchNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.SearchNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) SetSearchNodeStorage(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance {
	s.SearchNodeStorage = &v
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
	if s.ConfigserverList != nil {
		if err := s.ConfigserverList.Validate(); err != nil {
			return err
		}
	}
	if s.MongosList != nil {
		if err := s.MongosList.Validate(); err != nil {
			return err
		}
	}
	if s.ReplicaSets != nil {
		if err := s.ReplicaSets.Validate(); err != nil {
			return err
		}
	}
	if s.ShardList != nil {
		if err := s.ShardList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.ConfigserverAttribute != nil {
		for _, item := range s.ConfigserverAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceConfigserverListConfigserverAttribute struct {
	ConnectString        *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaxConnections       *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOPS              *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	NodeClass            *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription      *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeStorage          *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	if s.MongosAttribute != nil {
		for _, item := range s.MongosAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceMongosListMongosAttribute struct {
	ConnectSting         *string `json:"ConnectSting,omitempty" xml:"ConnectSting,omitempty"`
	ConnectString        *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaxConnections       *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOPS              *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	NodeClass            *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription      *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcCloudInstanceId   *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
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
	if s.ReplicaSet != nil {
		for _, item := range s.ReplicaSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceReplicaSetsReplicaSet struct {
	ConnectionDomain   *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	ConnectionPort     *string `json:"ConnectionPort,omitempty" xml:"ConnectionPort,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ReplicaSetRole     *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	VPCId              *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
	if s.ShardAttribute != nil {
		for _, item := range s.ShardAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceShardListShardAttribute struct {
	ConnectString        *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	CurrentKernelVersion *string `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaxConnections       *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxDiskMbps          *string `json:"MaxDiskMbps,omitempty" xml:"MaxDiskMbps,omitempty"`
	MaxIOPS              *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	NodeClass            *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	NodeDescription      *string `json:"NodeDescription,omitempty" xml:"NodeDescription,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeStorage          *int32  `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ReplicaSetName       *string `json:"ReplicaSetName,omitempty" xml:"ReplicaSetName,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

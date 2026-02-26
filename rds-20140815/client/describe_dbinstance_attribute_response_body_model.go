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
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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
	AccountMaxQuantity             *int32                                                                                `json:"AccountMaxQuantity,omitempty" xml:"AccountMaxQuantity,omitempty"`
	AdvancedFeatures               *string                                                                               `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	AutoUpgradeMinorVersion        *string                                                                               `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	AvailabilityValue              *string                                                                               `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	BabelfishConfig                *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig       `json:"BabelfishConfig,omitempty" xml:"BabelfishConfig,omitempty" type:"Struct"`
	BlueGreenDeploymentName        *string                                                                               `json:"BlueGreenDeploymentName,omitempty" xml:"BlueGreenDeploymentName,omitempty"`
	BlueInstanceName               *string                                                                               `json:"BlueInstanceName,omitempty" xml:"BlueInstanceName,omitempty"`
	BpeEnabled                     *string                                                                               `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	BurstingEnabled                *bool                                                                                 `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	CanTempUpgrade                 *bool                                                                                 `json:"CanTempUpgrade,omitempty" xml:"CanTempUpgrade,omitempty"`
	Category                       *string                                                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	ColdDataEnabled                *bool                                                                                 `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	Collation                      *string                                                                               `json:"Collation,omitempty" xml:"Collation,omitempty"`
	CompressionMode                *string                                                                               `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	CompressionRatio               *string                                                                               `json:"CompressionRatio,omitempty" xml:"CompressionRatio,omitempty"`
	ComputeBurstEnabled            *bool                                                                                 `json:"ComputeBurstEnabled,omitempty" xml:"ComputeBurstEnabled,omitempty"`
	ConnectionMode                 *string                                                                               `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	ConnectionString               *string                                                                               `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ConsoleVersion                 *string                                                                               `json:"ConsoleVersion,omitempty" xml:"ConsoleVersion,omitempty"`
	CreationTime                   *string                                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CurrentKernelVersion           *string                                                                               `json:"CurrentKernelVersion,omitempty" xml:"CurrentKernelVersion,omitempty"`
	DBClusterNodes                 *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes        `json:"DBClusterNodes,omitempty" xml:"DBClusterNodes,omitempty" type:"Struct"`
	DBInstanceCPU                  *string                                                                               `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	DBInstanceClass                *string                                                                               `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceClassType            *string                                                                               `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	DBInstanceDescription          *string                                                                               `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceDiskUsed             *string                                                                               `json:"DBInstanceDiskUsed,omitempty" xml:"DBInstanceDiskUsed,omitempty"`
	DBInstanceId                   *string                                                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMemory               *int64                                                                                `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	DBInstanceNetType              *string                                                                               `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStatus               *string                                                                               `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStorage              *int32                                                                                `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType          *string                                                                               `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DBInstanceType                 *string                                                                               `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	DBMaxQuantity                  *int32                                                                                `json:"DBMaxQuantity,omitempty" xml:"DBMaxQuantity,omitempty"`
	DedicatedHostGroupId           *string                                                                               `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DeletionProtection             *bool                                                                                 `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DisasterRecoveryInfo           *string                                                                               `json:"DisasterRecoveryInfo,omitempty" xml:"DisasterRecoveryInfo,omitempty"`
	DisasterRecoveryInstances      *string                                                                               `json:"DisasterRecoveryInstances,omitempty" xml:"DisasterRecoveryInstances,omitempty"`
	Engine                         *string                                                                               `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion                  *string                                                                               `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime                     *string                                                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Extra                          *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra                 `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	GeneralGroupName               *string                                                                               `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	GreenInstanceName              *string                                                                               `json:"GreenInstanceName,omitempty" xml:"GreenInstanceName,omitempty"`
	GuardDBInstanceId              *string                                                                               `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	IPType                         *string                                                                               `json:"IPType,omitempty" xml:"IPType,omitempty"`
	IncrementSourceDBInstanceId    *string                                                                               `json:"IncrementSourceDBInstanceId,omitempty" xml:"IncrementSourceDBInstanceId,omitempty"`
	InstanceNetworkType            *string                                                                               `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	InstructionSetArch             *string                                                                               `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	IoAccelerationEnabled          *string                                                                               `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	IsAnalyticIns                  *bool                                                                                 `json:"IsAnalyticIns,omitempty" xml:"IsAnalyticIns,omitempty"`
	IsAnalyticReadOnlyIns          *bool                                                                                 `json:"IsAnalyticReadOnlyIns,omitempty" xml:"IsAnalyticReadOnlyIns,omitempty"`
	LatestKernelVersion            *string                                                                               `json:"LatestKernelVersion,omitempty" xml:"LatestKernelVersion,omitempty"`
	LockMode                       *string                                                                               `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                     *string                                                                               `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainTime                   *string                                                                               `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	MasterInstanceId               *string                                                                               `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	MasterZone                     *string                                                                               `json:"MasterZone,omitempty" xml:"MasterZone,omitempty"`
	MaxConnections                 *int32                                                                                `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOMBPS                      *int32                                                                                `json:"MaxIOMBPS,omitempty" xml:"MaxIOMBPS,omitempty"`
	MaxIOPS                        *int32                                                                                `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	MultipleTempUpgrade            *bool                                                                                 `json:"MultipleTempUpgrade,omitempty" xml:"MultipleTempUpgrade,omitempty"`
	OptimizedWritesInfo            *string                                                                               `json:"OptimizedWritesInfo,omitempty" xml:"OptimizedWritesInfo,omitempty"`
	PGBouncerEnabled               *string                                                                               `json:"PGBouncerEnabled,omitempty" xml:"PGBouncerEnabled,omitempty"`
	PayType                        *string                                                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                           *string                                                                               `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType                      *int32                                                                                `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	ReadOnlyDBInstanceIds          *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	ReadOnlyStatus                 *string                                                                               `json:"ReadOnlyStatus,omitempty" xml:"ReadOnlyStatus,omitempty"`
	ReadonlyInstanceSQLDelayedTime *string                                                                               `json:"ReadonlyInstanceSQLDelayedTime,omitempty" xml:"ReadonlyInstanceSQLDelayedTime,omitempty"`
	RegionId                       *string                                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId                *string                                                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIPList                 *string                                                                               `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SecurityIPMode                 *string                                                                               `json:"SecurityIPMode,omitempty" xml:"SecurityIPMode,omitempty"`
	ServerlessConfig               *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig      `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	SlaveZones                     *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones            `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Struct"`
	SuperPermissionMode            *string                                                                               `json:"SuperPermissionMode,omitempty" xml:"SuperPermissionMode,omitempty"`
	SupportCompression             *bool                                                                                 `json:"SupportCompression,omitempty" xml:"SupportCompression,omitempty"`
	TempDBInstanceId               *string                                                                               `json:"TempDBInstanceId,omitempty" xml:"TempDBInstanceId,omitempty"`
	TempUpgradeTimeEnd             *string                                                                               `json:"TempUpgradeTimeEnd,omitempty" xml:"TempUpgradeTimeEnd,omitempty"`
	TempUpgradeTimeStart           *string                                                                               `json:"TempUpgradeTimeStart,omitempty" xml:"TempUpgradeTimeStart,omitempty"`
	TimeZone                       *string                                                                               `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	Tips                           *string                                                                               `json:"Tips,omitempty" xml:"Tips,omitempty"`
	TipsLevel                      *int32                                                                                `json:"TipsLevel,omitempty" xml:"TipsLevel,omitempty"`
	VSwitchId                      *string                                                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// ON
	VectorSupportStatus *string `json:"VectorSupportStatus,omitempty" xml:"VectorSupportStatus,omitempty"`
	VpcCloudInstanceId  *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId              *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	KindCode            *string `json:"kindCode,omitempty" xml:"kindCode,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetAccountMaxQuantity() *int32 {
	return s.AccountMaxQuantity
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetAdvancedFeatures() *string {
	return s.AdvancedFeatures
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetAvailabilityValue() *string {
	return s.AvailabilityValue
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetBabelfishConfig() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig {
	return s.BabelfishConfig
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetBlueGreenDeploymentName() *string {
	return s.BlueGreenDeploymentName
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetBlueInstanceName() *string {
	return s.BlueInstanceName
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCanTempUpgrade() *bool {
	return s.CanTempUpgrade
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCollation() *string {
	return s.Collation
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCompressionMode() *string {
	return s.CompressionMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCompressionRatio() *string {
	return s.CompressionRatio
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetComputeBurstEnabled() *bool {
	return s.ComputeBurstEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetConsoleVersion() *string {
	return s.ConsoleVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetCurrentKernelVersion() *string {
	return s.CurrentKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBClusterNodes() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes {
	return s.DBClusterNodes
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceCPU() *string {
	return s.DBInstanceCPU
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceClassType() *string {
	return s.DBInstanceClassType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceDiskUsed() *string {
	return s.DBInstanceDiskUsed
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceMemory() *int64 {
	return s.DBInstanceMemory
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDBMaxQuantity() *int32 {
	return s.DBMaxQuantity
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDisasterRecoveryInfo() *string {
	return s.DisasterRecoveryInfo
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetDisasterRecoveryInstances() *string {
	return s.DisasterRecoveryInstances
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetExtra() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra {
	return s.Extra
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetGreenInstanceName() *string {
	return s.GreenInstanceName
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIncrementSourceDBInstanceId() *string {
	return s.IncrementSourceDBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetInstructionSetArch() *string {
	return s.InstructionSetArch
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIsAnalyticIns() *bool {
	return s.IsAnalyticIns
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetIsAnalyticReadOnlyIns() *bool {
	return s.IsAnalyticReadOnlyIns
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetLatestKernelVersion() *string {
	return s.LatestKernelVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMasterZone() *string {
	return s.MasterZone
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaxIOMBPS() *int32 {
	return s.MaxIOMBPS
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetMultipleTempUpgrade() *bool {
	return s.MultipleTempUpgrade
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetOptimizedWritesInfo() *string {
	return s.OptimizedWritesInfo
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetPGBouncerEnabled() *string {
	return s.PGBouncerEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetProxyType() *int32 {
	return s.ProxyType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetReadOnlyDBInstanceIds() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds {
	return s.ReadOnlyDBInstanceIds
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetReadOnlyStatus() *string {
	return s.ReadOnlyStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetReadonlyInstanceSQLDelayedTime() *string {
	return s.ReadonlyInstanceSQLDelayedTime
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSecurityIPMode() *string {
	return s.SecurityIPMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetServerlessConfig() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig {
	return s.ServerlessConfig
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSlaveZones() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones {
	return s.SlaveZones
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSuperPermissionMode() *string {
	return s.SuperPermissionMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetSupportCompression() *bool {
	return s.SupportCompression
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTempDBInstanceId() *string {
	return s.TempDBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTempUpgradeTimeEnd() *string {
	return s.TempUpgradeTimeEnd
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTempUpgradeTimeStart() *string {
	return s.TempUpgradeTimeStart
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTips() *string {
	return s.Tips
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetTipsLevel() *int32 {
	return s.TipsLevel
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVectorSupportStatus() *string {
	return s.VectorSupportStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) GetKindCode() *string {
	return s.KindCode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAccountMaxQuantity(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AccountMaxQuantity = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAdvancedFeatures(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AdvancedFeatures = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAutoUpgradeMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetBabelfishConfig(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.BabelfishConfig = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetBlueGreenDeploymentName(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.BlueGreenDeploymentName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetBlueInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.BlueInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetBpeEnabled(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.BpeEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetBurstingEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCanTempUpgrade(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CanTempUpgrade = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCategory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Category = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetColdDataEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ColdDataEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCollation(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Collation = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCompressionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CompressionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCompressionRatio(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CompressionRatio = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetComputeBurstEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ComputeBurstEnabled = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetConsoleVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ConsoleVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetCurrentKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.CurrentKernelVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBClusterNodes(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBClusterNodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceCPU(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCPU = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceDiskUsed(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDiskUsed = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorageType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBInstanceType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDBMaxQuantity(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DBMaxQuantity = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDedicatedHostGroupId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDeletionProtection(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDisasterRecoveryInfo(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DisasterRecoveryInfo = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetDisasterRecoveryInstances(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.DisasterRecoveryInstances = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetExtra(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Extra = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetGeneralGroupName(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.GeneralGroupName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetGreenInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.GreenInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetGuardDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.GuardDBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIPType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIncrementSourceDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IncrementSourceDBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetInstructionSetArch(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.InstructionSetArch = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIoAccelerationEnabled(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIsAnalyticIns(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IsAnalyticIns = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetIsAnalyticReadOnlyIns(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.IsAnalyticReadOnlyIns = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetLatestKernelVersion(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.LatestKernelVersion = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaintainTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMasterZone(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MasterZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxIOMBPS(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxIOMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMaxIOPS(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetMultipleTempUpgrade(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.MultipleTempUpgrade = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetOptimizedWritesInfo(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.OptimizedWritesInfo = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetPGBouncerEnabled(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.PGBouncerEnabled = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetProxyType(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ProxyType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadOnlyDBInstanceIds(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadOnlyStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadOnlyStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetReadonlyInstanceSQLDelayedTime(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ReadonlyInstanceSQLDelayedTime = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSecurityIPMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetServerlessConfig(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.ServerlessConfig = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSlaveZones(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SlaveZones = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSuperPermissionMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SuperPermissionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetSupportCompression(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.SupportCompression = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTempDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.TempDBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTempUpgradeTimeEnd(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.TempUpgradeTimeEnd = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTempUpgradeTimeStart(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.TempUpgradeTimeStart = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTimeZone(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.TimeZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTips(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.Tips = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetTipsLevel(v int32) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.TipsLevel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVectorSupportStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VectorSupportStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetVpcCloudInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.VpcCloudInstanceId = &v
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

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) SetKindCode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) Validate() error {
	if s.BabelfishConfig != nil {
		if err := s.BabelfishConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DBClusterNodes != nil {
		if err := s.DBClusterNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	if s.ReadOnlyDBInstanceIds != nil {
		if err := s.ReadOnlyDBInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlaveZones != nil {
		if err := s.SlaveZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig struct {
	BabelfishEnabled *string `json:"BabelfishEnabled,omitempty" xml:"BabelfishEnabled,omitempty"`
	MigrationMode    *string `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) GetBabelfishEnabled() *string {
	return s.BabelfishEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) GetMigrationMode() *string {
	return s.MigrationMode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) SetBabelfishEnabled(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig {
	s.BabelfishEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) SetMigrationMode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig {
	s.MigrationMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeBabelfishConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes struct {
	DBClusterNode []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode `json:"DBClusterNode,omitempty" xml:"DBClusterNode,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) GetDBClusterNode() []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	return s.DBClusterNode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) SetDBClusterNode(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes {
	s.DBClusterNode = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodes) Validate() error {
	if s.DBClusterNode != nil {
		for _, item := range s.DBClusterNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode struct {
	ClassCode            *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ClassType            *string `json:"ClassType,omitempty" xml:"ClassType,omitempty"`
	Cpu                  *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DisasterRecoveryNode *bool   `json:"DisasterRecoveryNode,omitempty" xml:"DisasterRecoveryNode,omitempty"`
	Memory               *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeRegionId         *string `json:"NodeRegionId,omitempty" xml:"NodeRegionId,omitempty"`
	NodeRole             *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	NodeZoneId           *string `json:"NodeZoneId,omitempty" xml:"NodeZoneId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetClassType() *string {
	return s.ClassType
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetDisasterRecoveryNode() *bool {
	return s.DisasterRecoveryNode
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetMemory() *string {
	return s.Memory
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetNodeRegionId() *string {
	return s.NodeRegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetNodeRole() *string {
	return s.NodeRole
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetNodeZoneId() *string {
	return s.NodeZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetClassCode(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.ClassCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetClassType(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.ClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetCpu(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.Cpu = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetDisasterRecoveryNode(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.DisasterRecoveryNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetMemory(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.Memory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetNodeId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetNodeRegionId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.NodeRegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetNodeRole(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.NodeRole = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetNodeZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.NodeZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeDBClusterNodesDBClusterNode) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra struct {
	AccountSecurityPolicy *string                                                                            `json:"AccountSecurityPolicy,omitempty" xml:"AccountSecurityPolicy,omitempty"`
	DBInstanceIds         *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Struct"`
	RecoveryModel         *string                                                                            `json:"RecoveryModel,omitempty" xml:"RecoveryModel,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) GetAccountSecurityPolicy() *string {
	return s.AccountSecurityPolicy
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) GetDBInstanceIds() *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds {
	return s.DBInstanceIds
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) GetRecoveryModel() *string {
	return s.RecoveryModel
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) SetAccountSecurityPolicy(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra {
	s.AccountSecurityPolicy = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) SetDBInstanceIds(v *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) SetRecoveryModel(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra {
	s.RecoveryModel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtra) Validate() error {
	if s.DBInstanceIds != nil {
		if err := s.DBInstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds struct {
	DBInstanceId []*string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) GetDBInstanceId() []*string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) SetDBInstanceId(v []*string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds {
	s.DBInstanceId = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeExtraDBInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) GetReadOnlyDBInstanceId() []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIds) Validate() error {
	if s.ReadOnlyDBInstanceId != nil {
		for _, item := range s.ReadOnlyDBInstanceId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeReadOnlyDBInstanceIdsReadOnlyDBInstanceId) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig struct {
	AutoPause   *bool    `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	ScaleMax    *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin    *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	SwitchForce *bool    `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) GetSwitchForce() *bool {
	return s.SwitchForce
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) SetAutoPause(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig {
	s.AutoPause = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) SetScaleMax(v float64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) SetScaleMin(v float64) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) SetSwitchForce(v bool) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig {
	s.SwitchForce = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeServerlessConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones struct {
	SlaveZone []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone `json:"SlaveZone,omitempty" xml:"SlaveZone,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) GetSlaveZone() []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone {
	return s.SlaveZone
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) SetSlaveZone(v []*DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones {
	s.SlaveZone = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZones) Validate() error {
	if s.SlaveZone != nil {
		for _, item := range s.SlaveZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttributeSlaveZonesSlaveZone) Validate() error {
	return dara.Validate(s)
}

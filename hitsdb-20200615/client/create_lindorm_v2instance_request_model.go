// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchId(v string) *CreateLindormV2InstanceRequest
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *CreateLindormV2InstanceRequest
	GetArbiterZoneId() *string
	SetArchVersion(v string) *CreateLindormV2InstanceRequest
	GetArchVersion() *string
	SetAutoRenewDuration(v string) *CreateLindormV2InstanceRequest
	GetAutoRenewDuration() *string
	SetAutoRenewal(v bool) *CreateLindormV2InstanceRequest
	GetAutoRenewal() *bool
	SetCapacityStorageSize(v int32) *CreateLindormV2InstanceRequest
	GetCapacityStorageSize() *int32
	SetCloudStorageSize(v int32) *CreateLindormV2InstanceRequest
	GetCloudStorageSize() *int32
	SetCloudStorageType(v string) *CreateLindormV2InstanceRequest
	GetCloudStorageType() *string
	SetClusterMode(v string) *CreateLindormV2InstanceRequest
	GetClusterMode() *string
	SetClusterPattern(v string) *CreateLindormV2InstanceRequest
	GetClusterPattern() *string
	SetDuration(v int32) *CreateLindormV2InstanceRequest
	GetDuration() *int32
	SetEnableCapacityStorage(v bool) *CreateLindormV2InstanceRequest
	GetEnableCapacityStorage() *bool
	SetEngineList(v []*CreateLindormV2InstanceRequestEngineList) *CreateLindormV2InstanceRequest
	GetEngineList() []*CreateLindormV2InstanceRequestEngineList
	SetInstanceAlias(v string) *CreateLindormV2InstanceRequest
	GetInstanceAlias() *string
	SetOwnerAccount(v string) *CreateLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLindormV2InstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateLindormV2InstanceRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateLindormV2InstanceRequest
	GetPricingCycle() *string
	SetPrimaryVSwitchId(v string) *CreateLindormV2InstanceRequest
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *CreateLindormV2InstanceRequest
	GetPrimaryZoneId() *string
	SetRegionId(v string) *CreateLindormV2InstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLindormV2InstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateLindormV2InstanceRequest
	GetSecurityToken() *string
	SetStandbyVSwitchId(v string) *CreateLindormV2InstanceRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *CreateLindormV2InstanceRequest
	GetStandbyZoneId() *string
	SetVPCId(v string) *CreateLindormV2InstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateLindormV2InstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateLindormV2InstanceRequest
	GetZoneId() *string
}

type CreateLindormV2InstanceRequest struct {
	// The ID of the vSwitch in the arbiter zone for a multi-zone instance. The vSwitch must be in the zone specified by ArbiterZoneId. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the arbiter zone for a multi-zone instance. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The deployment architecture. Valid values:
	//
	// - **1.0**: single zone.
	//
	// - **2.0**: multi-zone Basic Edition.
	//
	// - **3.0**: multi-zone High-availability Edition.
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// Valid values: **1*	- to **12**.
	//
	// > This parameter is valid only when you set **AutoRenewal*	- to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Enable auto-renewal.
	//
	// - **false**: Disable auto-renewal.
	//
	// The default value is false.
	//
	// > This parameter is valid only when you set **PayType*	- to **PREPAY**.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The size of the storage-optimized storage. Unit: GB.
	//
	// example:
	//
	// 10000
	CapacityStorageSize *int32 `json:"CapacityStorageSize,omitempty" xml:"CapacityStorageSize,omitempty"`
	// The size of the cloud storage. Unit: GB.
	//
	// example:
	//
	// 320
	CloudStorageSize *int32 `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	// The storage class. This parameter is not required if you select **Big Data*	- or **Local SSD**.
	//
	// - **PerformanceStorage**: performance cloud storage
	//
	// - **StandardStorage**: standard cloud storage
	//
	// example:
	//
	// PerformanceStorage
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	// The instance mode. This parameter is not required.
	//
	// - **BASIC**: general-purpose mode
	//
	// example:
	//
	// BASIC
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The cluster type. Valid value:
	//
	// - **basic**: production
	//
	// example:
	//
	// basic
	ClusterPattern *string `json:"ClusterPattern,omitempty" xml:"ClusterPattern,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// - If you set PricingCycle to **Month**, the valid values are **1*	- to **9**.
	//
	// - If you set PricingCycle to **Year**, the valid values are **1*	- to **3**.
	//
	// > This parameter is required only when you set PayType to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable storage-optimized storage.
	//
	// example:
	//
	// false
	EnableCapacityStorage *bool `json:"EnableCapacityStorage,omitempty" xml:"EnableCapacityStorage,omitempty"`
	// The list of engine information.
	//
	// This parameter is required.
	EngineList []*CreateLindormV2InstanceRequestEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// lindorm-test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period of the instance. Valid values:
	//
	// - **Month**: The unit is month.
	//
	// - **Year**: The unit is year.
	//
	// > This parameter is required only when you set PayType to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the vSwitch in the primary zone for a multi-zone instance. The vSwitch must be in the zone specified by PrimaryZoneId. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The ID of the primary zone for a multi-zone instance. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The ID of the region where you want to create the instance. To query the available regions, call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the vSwitch in the secondary zone for a multi-zone instance. The vSwitch must be in the zone specified by StandbyZoneId. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the secondary zone for a multi-zone instance. **This parameter is required if you want to create a multi-zone instance.**
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The ID of the virtual private cloud (VPC) for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9ydz3vg93s1ozsd****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequest) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *CreateLindormV2InstanceRequest) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *CreateLindormV2InstanceRequest) GetArchVersion() *string {
	return s.ArchVersion
}

func (s *CreateLindormV2InstanceRequest) GetAutoRenewDuration() *string {
	return s.AutoRenewDuration
}

func (s *CreateLindormV2InstanceRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *CreateLindormV2InstanceRequest) GetCapacityStorageSize() *int32 {
	return s.CapacityStorageSize
}

func (s *CreateLindormV2InstanceRequest) GetCloudStorageSize() *int32 {
	return s.CloudStorageSize
}

func (s *CreateLindormV2InstanceRequest) GetCloudStorageType() *string {
	return s.CloudStorageType
}

func (s *CreateLindormV2InstanceRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateLindormV2InstanceRequest) GetClusterPattern() *string {
	return s.ClusterPattern
}

func (s *CreateLindormV2InstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateLindormV2InstanceRequest) GetEnableCapacityStorage() *bool {
	return s.EnableCapacityStorage
}

func (s *CreateLindormV2InstanceRequest) GetEngineList() []*CreateLindormV2InstanceRequestEngineList {
	return s.EngineList
}

func (s *CreateLindormV2InstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *CreateLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLindormV2InstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateLindormV2InstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateLindormV2InstanceRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *CreateLindormV2InstanceRequest) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *CreateLindormV2InstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLindormV2InstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateLindormV2InstanceRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CreateLindormV2InstanceRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CreateLindormV2InstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateLindormV2InstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLindormV2InstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLindormV2InstanceRequest) SetArbiterVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetArbiterZoneId(v string) *CreateLindormV2InstanceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetArchVersion(v string) *CreateLindormV2InstanceRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetAutoRenewDuration(v string) *CreateLindormV2InstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetAutoRenewal(v bool) *CreateLindormV2InstanceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCapacityStorageSize(v int32) *CreateLindormV2InstanceRequest {
	s.CapacityStorageSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCloudStorageSize(v int32) *CreateLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCloudStorageType(v string) *CreateLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetClusterMode(v string) *CreateLindormV2InstanceRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetClusterPattern(v string) *CreateLindormV2InstanceRequest {
	s.ClusterPattern = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetDuration(v int32) *CreateLindormV2InstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetEnableCapacityStorage(v bool) *CreateLindormV2InstanceRequest {
	s.EnableCapacityStorage = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetEngineList(v []*CreateLindormV2InstanceRequestEngineList) *CreateLindormV2InstanceRequest {
	s.EngineList = v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetInstanceAlias(v string) *CreateLindormV2InstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetOwnerAccount(v string) *CreateLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetOwnerId(v int64) *CreateLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPayType(v string) *CreateLindormV2InstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPricingCycle(v string) *CreateLindormV2InstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPrimaryVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPrimaryZoneId(v string) *CreateLindormV2InstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetRegionId(v string) *CreateLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceGroupId(v string) *CreateLindormV2InstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *CreateLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceOwnerId(v int64) *CreateLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetSecurityToken(v string) *CreateLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetStandbyVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetStandbyZoneId(v string) *CreateLindormV2InstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetVPCId(v string) *CreateLindormV2InstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetZoneId(v string) *CreateLindormV2InstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) Validate() error {
	if s.EngineList != nil {
		for _, item := range s.EngineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLindormV2InstanceRequestEngineList struct {
	// The engine type. Valid values:
	//
	// - **TABLE**: LindormTable.
	//
	// - **TSDB**: LindormTSDB.
	//
	// - **LSEARCH**: search engine.
	//
	// - **LTS**: LTS engine.
	//
	// - **LVECTOR**: vector engine.
	//
	// - **LCOLUMN**: column store engine.
	//
	// - **LAI**: AI engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The list of engine nodes.
	NodeGroupList []*CreateLindormV2InstanceRequestEngineListNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
}

func (s CreateLindormV2InstanceRequestEngineList) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormV2InstanceRequestEngineList) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequestEngineList) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateLindormV2InstanceRequestEngineList) GetNodeGroupList() []*CreateLindormV2InstanceRequestEngineListNodeGroupList {
	return s.NodeGroupList
}

func (s *CreateLindormV2InstanceRequestEngineList) SetEngineType(v string) *CreateLindormV2InstanceRequestEngineList {
	s.EngineType = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineList) SetNodeGroupList(v []*CreateLindormV2InstanceRequestEngineListNodeGroupList) *CreateLindormV2InstanceRequestEngineList {
	s.NodeGroupList = v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineList) Validate() error {
	if s.NodeGroupList != nil {
		for _, item := range s.NodeGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLindormV2InstanceRequestEngineListNodeGroupList struct {
	// The number of nodes in the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The disk size of a single node. The default unit is GB.
	//
	// example:
	//
	// 100
	NodeDiskSize *int32 `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	// The disk type of the node. This parameter is not required. **This parameter is available only for specific scenarios. To use this parameter, you must be added to the whitelist.**
	//
	// example:
	//
	// cloud_essd
	NodeDiskType *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	// The node specifications.
	//
	// If you select performance cloud storage or standard cloud storage, set this parameter to one of the following values:
	//
	// - **lindorm.c.2xlarge**: 8 cores, 16 GB.
	//
	// - **lindorm.g.2xlarge**: 8 cores, 32 GB.
	//
	// - **lindorm.c.4xlarge**: 16 cores, 32 GB.
	//
	// - **lindorm.g.4xlarge**: 16 cores, 64 GB.
	//
	// - **lindorm.c.8xlarge**: 32 cores, 64 GB.
	//
	// - **lindorm.g.8xlarge**: 32 cores, 128 GB.
	//
	// - **lindorm.r.2xlarge**: 8 cores, 64 GB.
	//
	// - **lindorm.r.4xlarge**: 16 cores, 128 GB.
	//
	// - **lindorm.r.8xlarge**: 32 cores, 256 GB.
	//
	// If you select the local SSD type, set this parameter to one of the following values:
	//
	// - **lindorm.i4.xlarge**: 4 cores, 32 GB (I4).
	//
	// - **lindorm.i4.2xlarge**: 8 cores, 64 GB (I4).
	//
	// - **lindorm.i4.4xlarge**: 16 cores, 128 GB (I4).
	//
	// - **lindorm.i4.8xlarge**: 32 cores, 256 GB (I4).
	//
	// - **lindorm.i3.xlarge**: 4 cores, 32 GB (I3).
	//
	// - **lindorm.i3.2xlarge**: 8 cores, 64 GB (I3).
	//
	// - **lindorm.i3.4xlarge**: 16 cores, 128 GB (I3).
	//
	// - **lindorm.i3.8xlarge**: 32 cores, 256 GB (I3).
	//
	// - **lindorm.i2.xlarge**: 4 cores, 32 GB (I2).
	//
	// - **lindorm.i2.2xlarge**: 8 cores, 64 GB (I2).
	//
	// - **lindorm.i2.4xlarge**: 16 cores, 128 GB (I2).
	//
	// - **lindorm.i2.8xlarge**: 32 cores, 256 GB (I2).
	//
	// If you select the big data type, set this parameter to one of the following values:
	//
	// - **lindorm.sd3c.3xlarge**: 14 cores, 56 GB (D3C PRO).
	//
	// - **lindorm.sd3c.7xlarge**: 28 cores, 112 GB (D3C PRO).
	//
	// - **lindorm.sd3c.14xlarge**: 56 cores, 224 GB (D3C PRO).
	//
	// - **lindorm.d2c.6xlarge**: 24 cores, 88 GB (D2C).
	//
	// - **lindorm.d2c.12xlarge**: 48 cores, 176 GB (D2C).
	//
	// - **lindorm.d2c.24xlarge**: 96 cores, 352 GB (D2C).
	//
	// - **lindorm.d2s.5xlarge**: 20 cores, 88 GB (D2S).
	//
	// - **lindorm.d2s.10xlarge**: 40 cores, 176 GB (D2S).
	//
	// - **lindorm.d1.2xlarge**: 8 cores, 32 GB (D1NE).
	//
	// - **lindorm.d1.4xlarge**: 16 cores, 64 GB (D1NE).
	//
	// - **lindorm.d1.6xlarge**: 24 cores, 96 GB (D1NE).
	//
	// This parameter is required.
	//
	// example:
	//
	// lindorm.g.2xlarge
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	// The name of the node group. **This parameter is required.**
	//
	// example:
	//
	// group_name_01
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateLindormV2InstanceRequestEngineListNodeGroupList) String() string {
	return dara.Prettify(s)
}

func (s CreateLindormV2InstanceRequestEngineListNodeGroupList) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskSize() *int32 {
	return s.NodeDiskSize
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskType() *string {
	return s.NodeDiskType
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeCount(v int32) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskSize(v int32) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskType(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeSpec(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetResourceGroupName(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) Validate() error {
	return dara.Validate(s)
}

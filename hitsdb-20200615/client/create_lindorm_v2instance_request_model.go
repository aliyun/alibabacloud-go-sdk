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
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// 10000
	CapacityStorageSize *int32 `json:"CapacityStorageSize,omitempty" xml:"CapacityStorageSize,omitempty"`
	// example:
	//
	// 320
	CloudStorageSize *int32 `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	// example:
	//
	// PerformanceStorage
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	// example:
	//
	// BASIC
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// example:
	//
	// basic
	ClusterPattern *string `json:"ClusterPattern,omitempty" xml:"ClusterPattern,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	EnableCapacityStorage *bool `json:"EnableCapacityStorage,omitempty" xml:"EnableCapacityStorage,omitempty"`
	// This parameter is required.
	EngineList []*CreateLindormV2InstanceRequestEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// example:
	//
	// lindorm-test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9ydz3vg93s1ozsd****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	EngineType    *string                                                  `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 7
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// 100
	NodeDiskSize *int32 `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	// example:
	//
	// cloud_essd
	NodeDiskType *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lindorm.g.2xlarge
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
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

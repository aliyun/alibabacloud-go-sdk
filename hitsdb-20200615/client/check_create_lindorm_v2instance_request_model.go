// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateLindormV2InstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchId(v string) *CheckCreateLindormV2InstanceRequest
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *CheckCreateLindormV2InstanceRequest
	GetArbiterZoneId() *string
	SetArchVersion(v string) *CheckCreateLindormV2InstanceRequest
	GetArchVersion() *string
	SetCloudStorageSize(v int32) *CheckCreateLindormV2InstanceRequest
	GetCloudStorageSize() *int32
	SetCloudStorageType(v string) *CheckCreateLindormV2InstanceRequest
	GetCloudStorageType() *string
	SetClusterPattern(v string) *CheckCreateLindormV2InstanceRequest
	GetClusterPattern() *string
	SetEngineList(v []*CheckCreateLindormV2InstanceRequestEngineList) *CheckCreateLindormV2InstanceRequest
	GetEngineList() []*CheckCreateLindormV2InstanceRequestEngineList
	SetInstanceAlias(v string) *CheckCreateLindormV2InstanceRequest
	GetInstanceAlias() *string
	SetOwnerAccount(v string) *CheckCreateLindormV2InstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCreateLindormV2InstanceRequest
	GetOwnerId() *int64
	SetPrimaryVSwitchId(v string) *CheckCreateLindormV2InstanceRequest
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *CheckCreateLindormV2InstanceRequest
	GetPrimaryZoneId() *string
	SetRegionId(v string) *CheckCreateLindormV2InstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckCreateLindormV2InstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckCreateLindormV2InstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCreateLindormV2InstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CheckCreateLindormV2InstanceRequest
	GetSecurityToken() *string
	SetStandbyVSwitchId(v string) *CheckCreateLindormV2InstanceRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *CheckCreateLindormV2InstanceRequest
	GetStandbyZoneId() *string
	SetVPCId(v string) *CheckCreateLindormV2InstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CheckCreateLindormV2InstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CheckCreateLindormV2InstanceRequest
	GetZoneId() *string
}

type CheckCreateLindormV2InstanceRequest struct {
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	ArbiterZoneId    *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// This parameter is required.
	ArchVersion      *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	CloudStorageSize *int32  `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	ClusterPattern   *string `json:"ClusterPattern,omitempty" xml:"ClusterPattern,omitempty"`
	// This parameter is required.
	EngineList       []*CheckCreateLindormV2InstanceRequestEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	InstanceAlias    *string                                          `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	OwnerAccount     *string                                          `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64                                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PrimaryVSwitchId *string                                          `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	PrimaryZoneId    *string                                          `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	StandbyZoneId        *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// This parameter is required.
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckCreateLindormV2InstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckCreateLindormV2InstanceRequest) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *CheckCreateLindormV2InstanceRequest) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *CheckCreateLindormV2InstanceRequest) GetArchVersion() *string {
	return s.ArchVersion
}

func (s *CheckCreateLindormV2InstanceRequest) GetCloudStorageSize() *int32 {
	return s.CloudStorageSize
}

func (s *CheckCreateLindormV2InstanceRequest) GetCloudStorageType() *string {
	return s.CloudStorageType
}

func (s *CheckCreateLindormV2InstanceRequest) GetClusterPattern() *string {
	return s.ClusterPattern
}

func (s *CheckCreateLindormV2InstanceRequest) GetEngineList() []*CheckCreateLindormV2InstanceRequestEngineList {
	return s.EngineList
}

func (s *CheckCreateLindormV2InstanceRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *CheckCreateLindormV2InstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCreateLindormV2InstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCreateLindormV2InstanceRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *CheckCreateLindormV2InstanceRequest) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *CheckCreateLindormV2InstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCreateLindormV2InstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckCreateLindormV2InstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCreateLindormV2InstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCreateLindormV2InstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckCreateLindormV2InstanceRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CheckCreateLindormV2InstanceRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CheckCreateLindormV2InstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CheckCreateLindormV2InstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CheckCreateLindormV2InstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckCreateLindormV2InstanceRequest) SetArbiterVSwitchId(v string) *CheckCreateLindormV2InstanceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetArbiterZoneId(v string) *CheckCreateLindormV2InstanceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetArchVersion(v string) *CheckCreateLindormV2InstanceRequest {
	s.ArchVersion = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetCloudStorageSize(v int32) *CheckCreateLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetCloudStorageType(v string) *CheckCreateLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetClusterPattern(v string) *CheckCreateLindormV2InstanceRequest {
	s.ClusterPattern = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetEngineList(v []*CheckCreateLindormV2InstanceRequestEngineList) *CheckCreateLindormV2InstanceRequest {
	s.EngineList = v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetInstanceAlias(v string) *CheckCreateLindormV2InstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetOwnerAccount(v string) *CheckCreateLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetOwnerId(v int64) *CheckCreateLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetPrimaryVSwitchId(v string) *CheckCreateLindormV2InstanceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetPrimaryZoneId(v string) *CheckCreateLindormV2InstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetRegionId(v string) *CheckCreateLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetResourceGroupId(v string) *CheckCreateLindormV2InstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *CheckCreateLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetResourceOwnerId(v int64) *CheckCreateLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetSecurityToken(v string) *CheckCreateLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetStandbyVSwitchId(v string) *CheckCreateLindormV2InstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetStandbyZoneId(v string) *CheckCreateLindormV2InstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetVPCId(v string) *CheckCreateLindormV2InstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetVSwitchId(v string) *CheckCreateLindormV2InstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) SetZoneId(v string) *CheckCreateLindormV2InstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequest) Validate() error {
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

type CheckCreateLindormV2InstanceRequestEngineList struct {
	// This parameter is required.
	EngineType    *string                                                       `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	NodeGroupList []*CheckCreateLindormV2InstanceRequestEngineListNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
}

func (s CheckCreateLindormV2InstanceRequestEngineList) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateLindormV2InstanceRequestEngineList) GoString() string {
	return s.String()
}

func (s *CheckCreateLindormV2InstanceRequestEngineList) GetEngineType() *string {
	return s.EngineType
}

func (s *CheckCreateLindormV2InstanceRequestEngineList) GetNodeGroupList() []*CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	return s.NodeGroupList
}

func (s *CheckCreateLindormV2InstanceRequestEngineList) SetEngineType(v string) *CheckCreateLindormV2InstanceRequestEngineList {
	s.EngineType = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineList) SetNodeGroupList(v []*CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) *CheckCreateLindormV2InstanceRequestEngineList {
	s.NodeGroupList = v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineList) Validate() error {
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

type CheckCreateLindormV2InstanceRequestEngineListNodeGroupList struct {
	// This parameter is required.
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeDiskSize *int32  `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	// This parameter is required.
	NodeSpec          *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GoString() string {
	return s.String()
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskSize() *int32 {
	return s.NodeDiskSize
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeDiskType() *string {
	return s.NodeDiskType
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeCount(v int32) *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskSize(v int32) *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskType(v string) *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeSpec(v string) *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) SetResourceGroupName(v string) *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

func (s *CheckCreateLindormV2InstanceRequestEngineListNodeGroupList) Validate() error {
	return dara.Validate(s)
}

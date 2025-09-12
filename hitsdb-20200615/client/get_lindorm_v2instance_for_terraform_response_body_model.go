// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceForTerraformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GetLindormV2InstanceForTerraformResponseBody
	GetAliUid() *int64
	SetArbiterVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetArbiterZoneId() *string
	SetAutoRenew(v bool) *GetLindormV2InstanceForTerraformResponseBody
	GetAutoRenew() *bool
	SetCloudStorageSize(v int64) *GetLindormV2InstanceForTerraformResponseBody
	GetCloudStorageSize() *int64
	SetColdStorage(v int32) *GetLindormV2InstanceForTerraformResponseBody
	GetColdStorage() *int32
	SetCreateMilliseconds(v int64) *GetLindormV2InstanceForTerraformResponseBody
	GetCreateMilliseconds() *int64
	SetDeletionProtection(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetDeletionProtection() *string
	SetDiskCategory(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetDiskCategory() *string
	SetDiskThreshold(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetDiskThreshold() *string
	SetDiskUsage(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetDiskUsage() *string
	SetEnableCompute(v bool) *GetLindormV2InstanceForTerraformResponseBody
	GetEnableCompute() *bool
	SetEngineList(v []*GetLindormV2InstanceForTerraformResponseBodyEngineList) *GetLindormV2InstanceForTerraformResponseBody
	GetEngineList() []*GetLindormV2InstanceForTerraformResponseBodyEngineList
	SetExpiredMilliseconds(v int64) *GetLindormV2InstanceForTerraformResponseBody
	GetExpiredMilliseconds() *int64
	SetInitialRootPassword(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetInitialRootPassword() *string
	SetInstanceAlias(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetInstanceAlias() *string
	SetInstanceId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetInstanceStatus() *string
	SetInstanceType(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetInstanceType() *string
	SetMaintainEndTime(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetMaintainStartTime() *string
	SetNetworkType(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetNetworkType() *string
	SetPayType(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetPayType() *string
	SetPrimaryVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetPrimaryZoneId() *string
	SetRegionId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetResourceGroupId() *string
	SetServiceType(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetServiceType() *string
	SetStandbyVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetStandbyZoneId() *string
	SetStorageUsage(v *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) *GetLindormV2InstanceForTerraformResponseBody
	GetStorageUsage() *GetLindormV2InstanceForTerraformResponseBodyStorageUsage
	SetVpcId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetVswitchId() *string
	SetWhiteIpList(v []*GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) *GetLindormV2InstanceForTerraformResponseBody
	GetWhiteIpList() []*GetLindormV2InstanceForTerraformResponseBodyWhiteIpList
	SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceForTerraformResponseBody
	GetZoneEngineInfoMap() map[string]interface{}
	SetZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody
	GetZoneId() *string
}

type GetLindormV2InstanceForTerraformResponseBody struct {
	AliUid              *int64                                                     `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ArbiterVSwitchId    *string                                                    `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	ArbiterZoneId       *string                                                    `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	AutoRenew           *bool                                                      `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CloudStorageSize    *int64                                                     `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	ColdStorage         *int32                                                     `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CreateMilliseconds  *int64                                                     `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	DeletionProtection  *string                                                    `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DiskCategory        *string                                                    `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskThreshold       *string                                                    `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	DiskUsage           *string                                                    `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	EnableCompute       *bool                                                      `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	EngineList          []*GetLindormV2InstanceForTerraformResponseBodyEngineList  `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	ExpiredMilliseconds *int64                                                     `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	InitialRootPassword *string                                                    `json:"InitialRootPassword,omitempty" xml:"InitialRootPassword,omitempty"`
	InstanceAlias       *string                                                    `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceId          *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string                                                    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaintainEndTime     *string                                                    `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime   *string                                                    `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkType         *string                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrimaryVSwitchId    *string                                                    `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	PrimaryZoneId       *string                                                    `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	RegionId            *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId     *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceType         *string                                                    `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StandbyVSwitchId    *string                                                    `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	StandbyZoneId       *string                                                    `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	StorageUsage        *GetLindormV2InstanceForTerraformResponseBodyStorageUsage  `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
	VpcId               *string                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string                                                    `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WhiteIpList         []*GetLindormV2InstanceForTerraformResponseBodyWhiteIpList `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	ZoneEngineInfoMap   map[string]interface{}                                     `json:"ZoneEngineInfoMap,omitempty" xml:"ZoneEngineInfoMap,omitempty"`
	ZoneId              *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetCloudStorageSize() *int64 {
	return s.CloudStorageSize
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetCreateMilliseconds() *int64 {
	return s.CreateMilliseconds
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetDiskThreshold() *string {
	return s.DiskThreshold
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetDiskUsage() *string {
	return s.DiskUsage
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetEnableCompute() *bool {
	return s.EnableCompute
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetEngineList() []*GetLindormV2InstanceForTerraformResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetExpiredMilliseconds() *int64 {
	return s.ExpiredMilliseconds
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetInitialRootPassword() *string {
	return s.InitialRootPassword
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetStorageUsage() *GetLindormV2InstanceForTerraformResponseBodyStorageUsage {
	return s.StorageUsage
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetWhiteIpList() []*GetLindormV2InstanceForTerraformResponseBodyWhiteIpList {
	return s.WhiteIpList
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetZoneEngineInfoMap() map[string]interface{} {
	return s.ZoneEngineInfoMap
}

func (s *GetLindormV2InstanceForTerraformResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetAliUid(v int64) *GetLindormV2InstanceForTerraformResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetArbiterVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetArbiterZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetAutoRenew(v bool) *GetLindormV2InstanceForTerraformResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetCloudStorageSize(v int64) *GetLindormV2InstanceForTerraformResponseBody {
	s.CloudStorageSize = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetColdStorage(v int32) *GetLindormV2InstanceForTerraformResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetCreateMilliseconds(v int64) *GetLindormV2InstanceForTerraformResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetDeletionProtection(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetDiskCategory(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetDiskThreshold(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetDiskUsage(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetEnableCompute(v bool) *GetLindormV2InstanceForTerraformResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetEngineList(v []*GetLindormV2InstanceForTerraformResponseBodyEngineList) *GetLindormV2InstanceForTerraformResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetExpiredMilliseconds(v int64) *GetLindormV2InstanceForTerraformResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetInitialRootPassword(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.InitialRootPassword = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetInstanceAlias(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetInstanceId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetInstanceStatus(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetInstanceType(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetMaintainEndTime(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetMaintainStartTime(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetNetworkType(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetPayType(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetPrimaryVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetPrimaryZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetRegionId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetRequestId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetResourceGroupId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetServiceType(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetStandbyVSwitchId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetStandbyZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetStorageUsage(v *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) *GetLindormV2InstanceForTerraformResponseBody {
	s.StorageUsage = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetVpcId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetVswitchId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetWhiteIpList(v []*GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) *GetLindormV2InstanceForTerraformResponseBody {
	s.WhiteIpList = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceForTerraformResponseBody {
	s.ZoneEngineInfoMap = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) SetZoneId(v string) *GetLindormV2InstanceForTerraformResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceForTerraformResponseBodyEngineList struct {
	ConnectAddressList []*GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList `json:"ConnectAddressList,omitempty" xml:"ConnectAddressList,omitempty" type:"Repeated"`
	Engine             *string                                                                     `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsLastVersion      *bool                                                                       `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	LatestVersion      *string                                                                     `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	NodeGroup          []*GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup          `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Repeated"`
	Version            *string                                                                     `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetConnectAddressList() []*GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList {
	return s.ConnectAddressList
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetEngine() *string {
	return s.Engine
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetIsLastVersion() *bool {
	return s.IsLastVersion
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetNodeGroup() []*GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	return s.NodeGroup
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) GetVersion() *string {
	return s.Version
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetConnectAddressList(v []*GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.ConnectAddressList = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetEngine(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetLatestVersion(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetNodeGroup(v []*GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.NodeGroup = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) SetVersion(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineList {
	s.Version = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Port    *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) GetAddress() *string {
	return s.Address
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) GetPort() *string {
	return s.Port
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) GetType() *string {
	return s.Type
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) SetAddress(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList {
	s.Address = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) SetPort(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList {
	s.Port = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) SetType(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList {
	s.Type = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListConnectAddressList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup struct {
	Category              *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CpuCoreCount          *int32  `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	EnableAttachLocalDisk *bool   `json:"EnableAttachLocalDisk,omitempty" xml:"EnableAttachLocalDisk,omitempty"`
	LocalDiskCapacity     *int64  `json:"LocalDiskCapacity,omitempty" xml:"LocalDiskCapacity,omitempty"`
	LocalDiskCategory     *string `json:"LocalDiskCategory,omitempty" xml:"LocalDiskCategory,omitempty"`
	MemorySizeGiB         *int32  `json:"MemorySizeGiB,omitempty" xml:"MemorySizeGiB,omitempty"`
	NodeSpec              *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	Quantity              *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupName     *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SpecId                *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetCategory() *string {
	return s.Category
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetEnableAttachLocalDisk() *bool {
	return s.EnableAttachLocalDisk
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetLocalDiskCapacity() *int64 {
	return s.LocalDiskCapacity
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetLocalDiskCategory() *string {
	return s.LocalDiskCategory
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetMemorySizeGiB() *int32 {
	return s.MemorySizeGiB
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetSpecId() *string {
	return s.SpecId
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) GetStatus() *string {
	return s.Status
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetCategory(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.Category = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetCpuCoreCount(v int32) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.CpuCoreCount = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetEnableAttachLocalDisk(v bool) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.EnableAttachLocalDisk = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetLocalDiskCapacity(v int64) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.LocalDiskCapacity = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetLocalDiskCategory(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.LocalDiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetMemorySizeGiB(v int32) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.MemorySizeGiB = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetNodeSpec(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.NodeSpec = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetQuantity(v int32) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.Quantity = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetResourceGroupName(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetSpecId(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.SpecId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) SetStatus(v string) *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup {
	s.Status = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyEngineListNodeGroup) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceForTerraformResponseBodyStorageUsage struct {
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	EngineUsage            map[string]interface{}   `json:"EngineUsage,omitempty" xml:"EngineUsage,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBodyStorageUsage) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) GetCapacityByDiskCategory() []map[string]interface{} {
	return s.CapacityByDiskCategory
}

func (s *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) GetEngineUsage() map[string]interface{} {
	return s.EngineUsage
}

func (s *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2InstanceForTerraformResponseBodyStorageUsage {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) SetEngineUsage(v map[string]interface{}) *GetLindormV2InstanceForTerraformResponseBodyStorageUsage {
	s.EngineUsage = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyStorageUsage) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceForTerraformResponseBodyWhiteIpList struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) GetIpList() *string {
	return s.IpList
}

func (s *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) SetGroupName(v string) *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList {
	s.GroupName = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) SetIpList(v string) *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList {
	s.IpList = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponseBodyWhiteIpList) Validate() error {
	return dara.Validate(s)
}

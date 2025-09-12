// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GetLindormV2InstanceResponseBody
	GetAliUid() *int64
	SetArbiterVSwitchId(v string) *GetLindormV2InstanceResponseBody
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *GetLindormV2InstanceResponseBody
	GetArbiterZoneId() *string
	SetAutoRenew(v bool) *GetLindormV2InstanceResponseBody
	GetAutoRenew() *bool
	SetCloudStorageSize(v int64) *GetLindormV2InstanceResponseBody
	GetCloudStorageSize() *int64
	SetColdStorage(v int32) *GetLindormV2InstanceResponseBody
	GetColdStorage() *int32
	SetCreateMilliseconds(v int64) *GetLindormV2InstanceResponseBody
	GetCreateMilliseconds() *int64
	SetDeletionProtection(v string) *GetLindormV2InstanceResponseBody
	GetDeletionProtection() *string
	SetDiskCategory(v string) *GetLindormV2InstanceResponseBody
	GetDiskCategory() *string
	SetDiskThreshold(v string) *GetLindormV2InstanceResponseBody
	GetDiskThreshold() *string
	SetDiskUsage(v string) *GetLindormV2InstanceResponseBody
	GetDiskUsage() *string
	SetEnableCompute(v bool) *GetLindormV2InstanceResponseBody
	GetEnableCompute() *bool
	SetEngineList(v []*GetLindormV2InstanceResponseBodyEngineList) *GetLindormV2InstanceResponseBody
	GetEngineList() []*GetLindormV2InstanceResponseBodyEngineList
	SetExpiredMilliseconds(v int64) *GetLindormV2InstanceResponseBody
	GetExpiredMilliseconds() *int64
	SetInitialRootPassword(v string) *GetLindormV2InstanceResponseBody
	GetInitialRootPassword() *string
	SetInstanceAlias(v string) *GetLindormV2InstanceResponseBody
	GetInstanceAlias() *string
	SetInstanceId(v string) *GetLindormV2InstanceResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *GetLindormV2InstanceResponseBody
	GetInstanceStatus() *string
	SetInstanceType(v string) *GetLindormV2InstanceResponseBody
	GetInstanceType() *string
	SetMaintainEndTime(v string) *GetLindormV2InstanceResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *GetLindormV2InstanceResponseBody
	GetMaintainStartTime() *string
	SetNetworkType(v string) *GetLindormV2InstanceResponseBody
	GetNetworkType() *string
	SetPayType(v string) *GetLindormV2InstanceResponseBody
	GetPayType() *string
	SetPrimaryVSwitchId(v string) *GetLindormV2InstanceResponseBody
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *GetLindormV2InstanceResponseBody
	GetPrimaryZoneId() *string
	SetRegionId(v string) *GetLindormV2InstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetLindormV2InstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLindormV2InstanceResponseBody
	GetResourceGroupId() *string
	SetServiceType(v string) *GetLindormV2InstanceResponseBody
	GetServiceType() *string
	SetStandbyVSwitchId(v string) *GetLindormV2InstanceResponseBody
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *GetLindormV2InstanceResponseBody
	GetStandbyZoneId() *string
	SetStorageUsage(v *GetLindormV2InstanceResponseBodyStorageUsage) *GetLindormV2InstanceResponseBody
	GetStorageUsage() *GetLindormV2InstanceResponseBodyStorageUsage
	SetVpcId(v string) *GetLindormV2InstanceResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *GetLindormV2InstanceResponseBody
	GetVswitchId() *string
	SetWhiteIpList(v []*GetLindormV2InstanceResponseBodyWhiteIpList) *GetLindormV2InstanceResponseBody
	GetWhiteIpList() []*GetLindormV2InstanceResponseBodyWhiteIpList
	SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceResponseBody
	GetZoneEngineInfoMap() map[string]interface{}
	SetZoneId(v string) *GetLindormV2InstanceResponseBody
	GetZoneId() *string
}

type GetLindormV2InstanceResponseBody struct {
	AliUid              *int64                                         `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ArbiterVSwitchId    *string                                        `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	ArbiterZoneId       *string                                        `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	AutoRenew           *bool                                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CloudStorageSize    *int64                                         `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	ColdStorage         *int32                                         `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CreateMilliseconds  *int64                                         `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	DeletionProtection  *string                                        `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DiskCategory        *string                                        `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskThreshold       *string                                        `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	DiskUsage           *string                                        `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	EnableCompute       *bool                                          `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	EngineList          []*GetLindormV2InstanceResponseBodyEngineList  `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	ExpiredMilliseconds *int64                                         `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	InitialRootPassword *string                                        `json:"InitialRootPassword,omitempty" xml:"InitialRootPassword,omitempty"`
	InstanceAlias       *string                                        `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceId          *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string                                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaintainEndTime     *string                                        `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime   *string                                        `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkType         *string                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string                                        `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrimaryVSwitchId    *string                                        `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	PrimaryZoneId       *string                                        `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	RegionId            *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId     *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceType         *string                                        `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StandbyVSwitchId    *string                                        `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	StandbyZoneId       *string                                        `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	StorageUsage        *GetLindormV2InstanceResponseBodyStorageUsage  `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
	VpcId               *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WhiteIpList         []*GetLindormV2InstanceResponseBodyWhiteIpList `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	ZoneEngineInfoMap   map[string]interface{}                         `json:"ZoneEngineInfoMap,omitempty" xml:"ZoneEngineInfoMap,omitempty"`
	ZoneId              *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetLindormV2InstanceResponseBody) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *GetLindormV2InstanceResponseBody) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *GetLindormV2InstanceResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetLindormV2InstanceResponseBody) GetCloudStorageSize() *int64 {
	return s.CloudStorageSize
}

func (s *GetLindormV2InstanceResponseBody) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *GetLindormV2InstanceResponseBody) GetCreateMilliseconds() *int64 {
	return s.CreateMilliseconds
}

func (s *GetLindormV2InstanceResponseBody) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *GetLindormV2InstanceResponseBody) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *GetLindormV2InstanceResponseBody) GetDiskThreshold() *string {
	return s.DiskThreshold
}

func (s *GetLindormV2InstanceResponseBody) GetDiskUsage() *string {
	return s.DiskUsage
}

func (s *GetLindormV2InstanceResponseBody) GetEnableCompute() *bool {
	return s.EnableCompute
}

func (s *GetLindormV2InstanceResponseBody) GetEngineList() []*GetLindormV2InstanceResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormV2InstanceResponseBody) GetExpiredMilliseconds() *int64 {
	return s.ExpiredMilliseconds
}

func (s *GetLindormV2InstanceResponseBody) GetInitialRootPassword() *string {
	return s.InitialRootPassword
}

func (s *GetLindormV2InstanceResponseBody) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetLindormV2InstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetLindormV2InstanceResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetLindormV2InstanceResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *GetLindormV2InstanceResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *GetLindormV2InstanceResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLindormV2InstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetLindormV2InstanceResponseBody) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *GetLindormV2InstanceResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *GetLindormV2InstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2InstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormV2InstanceResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormV2InstanceResponseBody) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *GetLindormV2InstanceResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *GetLindormV2InstanceResponseBody) GetStorageUsage() *GetLindormV2InstanceResponseBodyStorageUsage {
	return s.StorageUsage
}

func (s *GetLindormV2InstanceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLindormV2InstanceResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetLindormV2InstanceResponseBody) GetWhiteIpList() []*GetLindormV2InstanceResponseBodyWhiteIpList {
	return s.WhiteIpList
}

func (s *GetLindormV2InstanceResponseBody) GetZoneEngineInfoMap() map[string]interface{} {
	return s.ZoneEngineInfoMap
}

func (s *GetLindormV2InstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLindormV2InstanceResponseBody) SetAliUid(v int64) *GetLindormV2InstanceResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetArbiterVSwitchId(v string) *GetLindormV2InstanceResponseBody {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetArbiterZoneId(v string) *GetLindormV2InstanceResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetAutoRenew(v bool) *GetLindormV2InstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetCloudStorageSize(v int64) *GetLindormV2InstanceResponseBody {
	s.CloudStorageSize = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetColdStorage(v int32) *GetLindormV2InstanceResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetCreateMilliseconds(v int64) *GetLindormV2InstanceResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDeletionProtection(v string) *GetLindormV2InstanceResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskCategory(v string) *GetLindormV2InstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskThreshold(v string) *GetLindormV2InstanceResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskUsage(v string) *GetLindormV2InstanceResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetEnableCompute(v bool) *GetLindormV2InstanceResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetEngineList(v []*GetLindormV2InstanceResponseBodyEngineList) *GetLindormV2InstanceResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetExpiredMilliseconds(v int64) *GetLindormV2InstanceResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInitialRootPassword(v string) *GetLindormV2InstanceResponseBody {
	s.InitialRootPassword = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceAlias(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceId(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceStatus(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceType(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetMaintainEndTime(v string) *GetLindormV2InstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetMaintainStartTime(v string) *GetLindormV2InstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetNetworkType(v string) *GetLindormV2InstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetPayType(v string) *GetLindormV2InstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetPrimaryVSwitchId(v string) *GetLindormV2InstanceResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetPrimaryZoneId(v string) *GetLindormV2InstanceResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetRegionId(v string) *GetLindormV2InstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetRequestId(v string) *GetLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetResourceGroupId(v string) *GetLindormV2InstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetServiceType(v string) *GetLindormV2InstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetStandbyVSwitchId(v string) *GetLindormV2InstanceResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetStandbyZoneId(v string) *GetLindormV2InstanceResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetStorageUsage(v *GetLindormV2InstanceResponseBodyStorageUsage) *GetLindormV2InstanceResponseBody {
	s.StorageUsage = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetVpcId(v string) *GetLindormV2InstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetVswitchId(v string) *GetLindormV2InstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetWhiteIpList(v []*GetLindormV2InstanceResponseBodyWhiteIpList) *GetLindormV2InstanceResponseBody {
	s.WhiteIpList = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceResponseBody {
	s.ZoneEngineInfoMap = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetZoneId(v string) *GetLindormV2InstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceResponseBodyEngineList struct {
	ConnectAddressList []*GetLindormV2InstanceResponseBodyEngineListConnectAddressList `json:"ConnectAddressList,omitempty" xml:"ConnectAddressList,omitempty" type:"Repeated"`
	EnableBackup       *string                                                         `json:"EnableBackup,omitempty" xml:"EnableBackup,omitempty"`
	EnableCDC          *string                                                         `json:"EnableCDC,omitempty" xml:"EnableCDC,omitempty"`
	Engine             *string                                                         `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsLastVersion      *bool                                                           `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	LatestVersion      *string                                                         `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	NodeGroup          []*GetLindormV2InstanceResponseBodyEngineListNodeGroup          `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Repeated"`
	Version            *string                                                         `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetConnectAddressList() []*GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	return s.ConnectAddressList
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetEnableBackup() *string {
	return s.EnableBackup
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetEnableCDC() *string {
	return s.EnableCDC
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetEngine() *string {
	return s.Engine
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetIsLastVersion() *bool {
	return s.IsLastVersion
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetNodeGroup() []*GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	return s.NodeGroup
}

func (s *GetLindormV2InstanceResponseBodyEngineList) GetVersion() *string {
	return s.Version
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetConnectAddressList(v []*GetLindormV2InstanceResponseBodyEngineListConnectAddressList) *GetLindormV2InstanceResponseBodyEngineList {
	s.ConnectAddressList = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetEnableBackup(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.EnableBackup = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetEnableCDC(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.EnableCDC = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetEngine(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormV2InstanceResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetLatestVersion(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetNodeGroup(v []*GetLindormV2InstanceResponseBodyEngineListNodeGroup) *GetLindormV2InstanceResponseBodyEngineList {
	s.NodeGroup = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetVersion(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.Version = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceResponseBodyEngineListConnectAddressList struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Port    *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyEngineListConnectAddressList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineListConnectAddressList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) GetAddress() *string {
	return s.Address
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) GetPort() *string {
	return s.Port
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) GetType() *string {
	return s.Type
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetAddress(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Address = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetPort(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Port = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetType(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Type = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceResponseBodyEngineListNodeGroup struct {
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

func (s GetLindormV2InstanceResponseBodyEngineListNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineListNodeGroup) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetCategory() *string {
	return s.Category
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetEnableAttachLocalDisk() *bool {
	return s.EnableAttachLocalDisk
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetLocalDiskCapacity() *int64 {
	return s.LocalDiskCapacity
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetLocalDiskCategory() *string {
	return s.LocalDiskCategory
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetMemorySizeGiB() *int32 {
	return s.MemorySizeGiB
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetSpecId() *string {
	return s.SpecId
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) GetStatus() *string {
	return s.Status
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetCategory(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Category = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetCpuCoreCount(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.CpuCoreCount = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetEnableAttachLocalDisk(v bool) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.EnableAttachLocalDisk = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetLocalDiskCapacity(v int64) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.LocalDiskCapacity = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetLocalDiskCategory(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.LocalDiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetMemorySizeGiB(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.MemorySizeGiB = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetNodeSpec(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.NodeSpec = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetQuantity(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Quantity = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetResourceGroupName(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetSpecId(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.SpecId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetStatus(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Status = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceResponseBodyStorageUsage struct {
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	EngineUsage            map[string]interface{}   `json:"EngineUsage,omitempty" xml:"EngineUsage,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyStorageUsage) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) GetCapacityByDiskCategory() []map[string]interface{} {
	return s.CapacityByDiskCategory
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) GetEngineUsage() map[string]interface{} {
	return s.EngineUsage
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2InstanceResponseBodyStorageUsage {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) SetEngineUsage(v map[string]interface{}) *GetLindormV2InstanceResponseBodyStorageUsage {
	s.EngineUsage = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceResponseBodyWhiteIpList struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyWhiteIpList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyWhiteIpList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) GetIpList() *string {
	return s.IpList
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) SetGroupName(v string) *GetLindormV2InstanceResponseBodyWhiteIpList {
	s.GroupName = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) SetIpList(v string) *GetLindormV2InstanceResponseBodyWhiteIpList {
	s.IpList = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) Validate() error {
	return dara.Validate(s)
}

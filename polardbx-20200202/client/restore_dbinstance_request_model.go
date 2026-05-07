// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RestoreDBInstanceRequest
	GetAutoRenew() *bool
	SetBackupSetId(v string) *RestoreDBInstanceRequest
	GetBackupSetId() *string
	SetBackupSetRegion(v string) *RestoreDBInstanceRequest
	GetBackupSetRegion() *string
	SetCNNodeCount(v string) *RestoreDBInstanceRequest
	GetCNNodeCount() *string
	SetClientToken(v string) *RestoreDBInstanceRequest
	GetClientToken() *string
	SetCloneInstanceName(v string) *RestoreDBInstanceRequest
	GetCloneInstanceName() *string
	SetCnClass(v string) *RestoreDBInstanceRequest
	GetCnClass() *string
	SetDBNodeClass(v string) *RestoreDBInstanceRequest
	GetDBNodeClass() *string
	SetDBNodeCount(v int32) *RestoreDBInstanceRequest
	GetDBNodeCount() *int32
	SetDNNodeCount(v string) *RestoreDBInstanceRequest
	GetDNNodeCount() *string
	SetDnClass(v string) *RestoreDBInstanceRequest
	GetDnClass() *string
	SetEngineVersion(v string) *RestoreDBInstanceRequest
	GetEngineVersion() *string
	SetGdnRole(v string) *RestoreDBInstanceRequest
	GetGdnRole() *string
	SetNetworkType(v string) *RestoreDBInstanceRequest
	GetNetworkType() *string
	SetPayType(v string) *RestoreDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *RestoreDBInstanceRequest
	GetPeriod() *string
	SetPrimaryZone(v string) *RestoreDBInstanceRequest
	GetPrimaryZone() *string
	SetRecoveryTypeCode(v string) *RestoreDBInstanceRequest
	GetRecoveryTypeCode() *string
	SetRegionId(v string) *RestoreDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RestoreDBInstanceRequest
	GetResourceGroupId() *string
	SetRestoreTime(v string) *RestoreDBInstanceRequest
	GetRestoreTime() *string
	SetSecondaryZone(v string) *RestoreDBInstanceRequest
	GetSecondaryZone() *string
	SetSeries(v string) *RestoreDBInstanceRequest
	GetSeries() *string
	SetSourceInstanceRegion(v string) *RestoreDBInstanceRequest
	GetSourceInstanceRegion() *string
	SetStorageType(v string) *RestoreDBInstanceRequest
	GetStorageType() *string
	SetTertiaryZone(v string) *RestoreDBInstanceRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *RestoreDBInstanceRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *RestoreDBInstanceRequest
	GetUsedTime() *int32
	SetVPCId(v string) *RestoreDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *RestoreDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RestoreDBInstanceRequest
	GetZoneId() *string
}

type RestoreDBInstanceRequest struct {
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	BackupSetRegion *string `json:"BackupSetRegion,omitempty" xml:"BackupSetRegion,omitempty"`
	// example:
	//
	// 2
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-***
	CloneInstanceName *string `json:"CloneInstanceName,omitempty" xml:"CloneInstanceName,omitempty"`
	// example:
	//
	// polarx.x4.medium.2e
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount *int32 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// example:
	//
	// 2
	DNNodeCount *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// example:
	//
	// mysql.n4.medium.25
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// standby
	GdnRole *string `json:"GdnRole,omitempty" xml:"GdnRole,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Clone
	RecoveryTypeCode *string `json:"RecoveryTypeCode,omitempty" xml:"RecoveryTypeCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2024-10-14T00:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	SourceInstanceRegion *string `json:"SourceInstanceRegion,omitempty" xml:"SourceInstanceRegion,omitempty"`
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RestoreDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RestoreDBInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *RestoreDBInstanceRequest) GetBackupSetRegion() *string {
	return s.BackupSetRegion
}

func (s *RestoreDBInstanceRequest) GetCNNodeCount() *string {
	return s.CNNodeCount
}

func (s *RestoreDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreDBInstanceRequest) GetCloneInstanceName() *string {
	return s.CloneInstanceName
}

func (s *RestoreDBInstanceRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *RestoreDBInstanceRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *RestoreDBInstanceRequest) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *RestoreDBInstanceRequest) GetDNNodeCount() *string {
	return s.DNNodeCount
}

func (s *RestoreDBInstanceRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *RestoreDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *RestoreDBInstanceRequest) GetGdnRole() *string {
	return s.GdnRole
}

func (s *RestoreDBInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *RestoreDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *RestoreDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *RestoreDBInstanceRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *RestoreDBInstanceRequest) GetRecoveryTypeCode() *string {
	return s.RecoveryTypeCode
}

func (s *RestoreDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestoreDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreDBInstanceRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *RestoreDBInstanceRequest) GetSeries() *string {
	return s.Series
}

func (s *RestoreDBInstanceRequest) GetSourceInstanceRegion() *string {
	return s.SourceInstanceRegion
}

func (s *RestoreDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *RestoreDBInstanceRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *RestoreDBInstanceRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *RestoreDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *RestoreDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *RestoreDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RestoreDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RestoreDBInstanceRequest) SetAutoRenew(v bool) *RestoreDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetBackupSetId(v string) *RestoreDBInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetBackupSetRegion(v string) *RestoreDBInstanceRequest {
	s.BackupSetRegion = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetCNNodeCount(v string) *RestoreDBInstanceRequest {
	s.CNNodeCount = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetClientToken(v string) *RestoreDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetCloneInstanceName(v string) *RestoreDBInstanceRequest {
	s.CloneInstanceName = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetCnClass(v string) *RestoreDBInstanceRequest {
	s.CnClass = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetDBNodeClass(v string) *RestoreDBInstanceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetDBNodeCount(v int32) *RestoreDBInstanceRequest {
	s.DBNodeCount = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetDNNodeCount(v string) *RestoreDBInstanceRequest {
	s.DNNodeCount = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetDnClass(v string) *RestoreDBInstanceRequest {
	s.DnClass = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetEngineVersion(v string) *RestoreDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetGdnRole(v string) *RestoreDBInstanceRequest {
	s.GdnRole = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetNetworkType(v string) *RestoreDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetPayType(v string) *RestoreDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetPeriod(v string) *RestoreDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetPrimaryZone(v string) *RestoreDBInstanceRequest {
	s.PrimaryZone = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetRecoveryTypeCode(v string) *RestoreDBInstanceRequest {
	s.RecoveryTypeCode = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetRegionId(v string) *RestoreDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetResourceGroupId(v string) *RestoreDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetRestoreTime(v string) *RestoreDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetSecondaryZone(v string) *RestoreDBInstanceRequest {
	s.SecondaryZone = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetSeries(v string) *RestoreDBInstanceRequest {
	s.Series = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetSourceInstanceRegion(v string) *RestoreDBInstanceRequest {
	s.SourceInstanceRegion = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetStorageType(v string) *RestoreDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetTertiaryZone(v string) *RestoreDBInstanceRequest {
	s.TertiaryZone = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetTopologyType(v string) *RestoreDBInstanceRequest {
	s.TopologyType = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetUsedTime(v int32) *RestoreDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetVPCId(v string) *RestoreDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetVSwitchId(v string) *RestoreDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *RestoreDBInstanceRequest) SetZoneId(v string) *RestoreDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *RestoreDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

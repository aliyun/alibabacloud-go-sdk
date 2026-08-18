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
	// Specifies whether to enable auto-renewal. Default value: true.
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The backup set ID.
	//
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region where the backup set resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	BackupSetRegion *string `json:"BackupSetRegion,omitempty" xml:"BackupSetRegion,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	CNNodeCount *string `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// The client token used to ensure the idempotence of the request. Use a different value for each request.
	//
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-***
	CloneInstanceName *string `json:"CloneInstanceName,omitempty" xml:"CloneInstanceName,omitempty"`
	// The compute node specifications. Valid values:
	//
	// - polarx.x4.medium.2e: 2 cores, 8 GB
	//
	// - polarx.x4.large.2e: 4 cores, 16 GB
	//
	// - polarx.x8.large.2e: 4 cores, 32 GB
	//
	// - polarx.x4.xlarge.2e: 8 cores, 32 GB
	//
	// - polarx.x8.xlarge.2e: 8 cores, 64 GB
	//
	// - polarx.x4.2xlarge.2e: 16 cores, 64 GB
	//
	// - polarx.x8.2xlarge.2e: 16 cores, 128 GB
	//
	// - polarx.x4.4xlarge.2e: 32 cores, 128 GB
	//
	// - polarx.x8.4xlarge.2e: 32 cores, 256 GB
	//
	// - polarx.st.8xlarge.2e: 60 cores, 470 GB
	//
	// - polarx.st.12xlarge.2e: 90 cores, 720 GB
	//
	// example:
	//
	// polarx.x4.medium.2e
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// The node specifications. Valid values:
	//
	// - polarx.x4.medium.2e: 2 cores, 8 GB
	//
	// - polarx.x4.large.2e: 4 cores, 16 GB
	//
	// - polarx.x8.large.2e: 4 cores, 32 GB
	//
	// - polarx.x4.xlarge.2e: 8 cores, 32 GB
	//
	// - polarx.x8.xlarge.2e: 8 cores, 64 GB
	//
	// - polarx.x4.2xlarge.2e: 16 cores, 64 GB
	//
	// - polarx.x8.2xlarge.2e: 16 cores, 128 GB
	//
	// - polarx.x4.4xlarge.2e: 32 cores, 128 GB
	//
	// - polarx.x8.4xlarge.2e: 32 cores, 256 GB
	//
	// - polarx.st.8xlarge.2e: 60 cores, 470 GB
	//
	// - polarx.st.12xlarge.2e: 90 cores, 720 GB
	//
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of instance nodes. The minimum value is 2.
	//
	// example:
	//
	// 2
	DBNodeCount *int32 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The number of storage nodes.
	//
	// example:
	//
	// 2
	DNNodeCount *string `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// The storage node specifications. Valid values:
	//
	// - mysql.n4.medium.25: 2 cores, 8 GB
	//
	// - mysql.n4.large.25: 4 cores, 16 GB
	//
	// - mysql.x8.large.25: 4 cores, 32 GB
	//
	// - mysql.n4.xlarge.25: 8 cores, 32 GB
	//
	// - mysql.x8.xlarge.25: 8 cores, 64 GB
	//
	// - mysql.n4.2xlarge.25: 16 cores, 64 GB
	//
	// - mysql.x8.2xlarge.25: 16 cores, 128 GB
	//
	// - mysql.x4.4xlarge.25: 32 cores, 128 GB
	//
	// - mysql.x8.4xlarge.25: 32 cores, 256 GB
	//
	// - mysql.st.8xlarge.25: 60 cores, 470 GB
	//
	// - mysql.st.12xlarge.25: 90 cores, 720 GB
	//
	// example:
	//
	// mysql.n4.medium.25
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// The MySQL DPI engine version. Valid values: 5.7 and 8.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The GDN role.
	//
	// example:
	//
	// standby
	GdnRole *string `json:"GdnRole,omitempty" xml:"GdnRole,omitempty"`
	// The network type. Only VPC is supported.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance.
	//
	// - PREPAY: subscription
	//
	// - POSTPAY: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle. Valid values for subscription: Year and Month. Default value for pay-as-you-go: Hour.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The recovery type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Clone
	RecoveryTypeCode *string `json:"RecoveryTypeCode,omitempty" xml:"RecoveryTypeCode,omitempty"`
	// The region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter can be left empty. This parameter is not supported.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The point in time to which you want to restore the instance. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2024-10-14T00:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// The instance series. Valid values:
	//
	// - enterprise: Enterprise Edition.
	//
	// - standard: Standard Edition.
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The region where the source instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	SourceInstanceRegion *string `json:"SourceInstanceRegion,omitempty" xml:"SourceInstanceRegion,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone for Three-zone deployment.
	//
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The topology type. Valid values:
	//
	// - 3azones: three-zone deployment.
	//
	// - 1azone: single-zone deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The subscription duration. Specify the number of months or years.
	//
	// > If Period is set to Year, valid values of this parameter are 1, 2, and 3.
	//
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
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone of the instance.
	//
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *CreateClusterRequest
	GetAutoRenew() *string
	SetBackupId(v string) *CreateClusterRequest
	GetBackupId() *string
	SetClientToken(v string) *CreateClusterRequest
	GetClientToken() *string
	SetCloudType(v string) *CreateClusterRequest
	GetCloudType() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetColdStorageSize(v string) *CreateClusterRequest
	GetColdStorageSize() *string
	SetCoreDiskQuantity(v string) *CreateClusterRequest
	GetCoreDiskQuantity() *string
	SetCoreDiskSize(v string) *CreateClusterRequest
	GetCoreDiskSize() *string
	SetCoreDiskType(v string) *CreateClusterRequest
	GetCoreDiskType() *string
	SetCoreInstanceQuantity(v string) *CreateClusterRequest
	GetCoreInstanceQuantity() *string
	SetCoreInstanceType(v string) *CreateClusterRequest
	GetCoreInstanceType() *string
	SetDbInstanceConnType(v string) *CreateClusterRequest
	GetDbInstanceConnType() *string
	SetDbInstanceType(v string) *CreateClusterRequest
	GetDbInstanceType() *string
	SetDbType(v string) *CreateClusterRequest
	GetDbType() *string
	SetDepMode(v string) *CreateClusterRequest
	GetDepMode() *string
	SetDuration(v string) *CreateClusterRequest
	GetDuration() *string
	SetEngine(v string) *CreateClusterRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateClusterRequest
	GetEngineVersion() *string
	SetIsColdStorage(v string) *CreateClusterRequest
	GetIsColdStorage() *string
	SetMasterInstanceType(v string) *CreateClusterRequest
	GetMasterInstanceType() *string
	SetNetType(v string) *CreateClusterRequest
	GetNetType() *string
	SetPayType(v string) *CreateClusterRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateClusterRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateClusterRequest
	GetRegionId() *string
	SetRestoreTime(v string) *CreateClusterRequest
	GetRestoreTime() *string
	SetSecurityIPList(v string) *CreateClusterRequest
	GetSecurityIPList() *string
	SetSrcDBInstanceId(v string) *CreateClusterRequest
	GetSrcDBInstanceId() *string
	SetVSwitchId(v string) *CreateClusterRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateClusterRequest
	GetZoneId() *string
}

type CreateClusterRequest struct {
	AutoRenew   *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BackupId    *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CloudType *string `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	// This parameter is required.
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ColdStorageSize *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// This parameter is required.
	CoreDiskQuantity *string `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	// This parameter is required.
	CoreDiskSize *string `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// This parameter is required.
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// This parameter is required.
	CoreInstanceQuantity *string `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	// This parameter is required.
	CoreInstanceType   *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	DbInstanceConnType *string `json:"DbInstanceConnType,omitempty" xml:"DbInstanceConnType,omitempty"`
	DbInstanceType     *string `json:"DbInstanceType,omitempty" xml:"DbInstanceType,omitempty"`
	DbType             *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DepMode            *string `json:"DepMode,omitempty" xml:"DepMode,omitempty"`
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IsColdStorage *string `json:"IsColdStorage,omitempty" xml:"IsColdStorage,omitempty"`
	// This parameter is required.
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// This parameter is required.
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType      *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// This parameter is required.
	SecurityIPList  *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateClusterRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateClusterRequest) GetCloudType() *string {
	return s.CloudType
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetColdStorageSize() *string {
	return s.ColdStorageSize
}

func (s *CreateClusterRequest) GetCoreDiskQuantity() *string {
	return s.CoreDiskQuantity
}

func (s *CreateClusterRequest) GetCoreDiskSize() *string {
	return s.CoreDiskSize
}

func (s *CreateClusterRequest) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *CreateClusterRequest) GetCoreInstanceQuantity() *string {
	return s.CoreInstanceQuantity
}

func (s *CreateClusterRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *CreateClusterRequest) GetDbInstanceConnType() *string {
	return s.DbInstanceConnType
}

func (s *CreateClusterRequest) GetDbInstanceType() *string {
	return s.DbInstanceType
}

func (s *CreateClusterRequest) GetDbType() *string {
	return s.DbType
}

func (s *CreateClusterRequest) GetDepMode() *string {
	return s.DepMode
}

func (s *CreateClusterRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateClusterRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateClusterRequest) GetIsColdStorage() *string {
	return s.IsColdStorage
}

func (s *CreateClusterRequest) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *CreateClusterRequest) GetNetType() *string {
	return s.NetType
}

func (s *CreateClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateClusterRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateClusterRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateClusterRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateClusterRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequest) SetAutoRenew(v string) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetBackupId(v string) *CreateClusterRequest {
	s.BackupId = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetCloudType(v string) *CreateClusterRequest {
	s.CloudType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetColdStorageSize(v string) *CreateClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskQuantity(v string) *CreateClusterRequest {
	s.CoreDiskQuantity = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskSize(v string) *CreateClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreDiskType(v string) *CreateClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceQuantity(v string) *CreateClusterRequest {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceType(v string) *CreateClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDbInstanceConnType(v string) *CreateClusterRequest {
	s.DbInstanceConnType = &v
	return s
}

func (s *CreateClusterRequest) SetDbInstanceType(v string) *CreateClusterRequest {
	s.DbInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDbType(v string) *CreateClusterRequest {
	s.DbType = &v
	return s
}

func (s *CreateClusterRequest) SetDepMode(v string) *CreateClusterRequest {
	s.DepMode = &v
	return s
}

func (s *CreateClusterRequest) SetDuration(v string) *CreateClusterRequest {
	s.Duration = &v
	return s
}

func (s *CreateClusterRequest) SetEngine(v string) *CreateClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateClusterRequest) SetEngineVersion(v string) *CreateClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateClusterRequest) SetIsColdStorage(v string) *CreateClusterRequest {
	s.IsColdStorage = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceType(v string) *CreateClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetNetType(v string) *CreateClusterRequest {
	s.NetType = &v
	return s
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetPricingCycle(v string) *CreateClusterRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetRestoreTime(v string) *CreateClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityIPList(v string) *CreateClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateClusterRequest) SetSrcDBInstanceId(v string) *CreateClusterRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}

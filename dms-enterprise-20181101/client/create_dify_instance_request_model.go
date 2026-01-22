// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdbpgInstanceMode(v string) *CreateDifyInstanceRequest
	GetAdbpgInstanceMode() *string
	SetAutoRenew(v bool) *CreateDifyInstanceRequest
	GetAutoRenew() *bool
	SetBackupVSwitchId(v string) *CreateDifyInstanceRequest
	GetBackupVSwitchId() *string
	SetClientToken(v string) *CreateDifyInstanceRequest
	GetClientToken() *string
	SetDataRegion(v string) *CreateDifyInstanceRequest
	GetDataRegion() *string
	SetDatabaseOption(v string) *CreateDifyInstanceRequest
	GetDatabaseOption() *string
	SetDbEngineType(v string) *CreateDifyInstanceRequest
	GetDbEngineType() *string
	SetDbEngineVersion(v string) *CreateDifyInstanceRequest
	GetDbEngineVersion() *string
	SetDbInstanceAccount(v string) *CreateDifyInstanceRequest
	GetDbInstanceAccount() *string
	SetDbInstanceCategory(v string) *CreateDifyInstanceRequest
	GetDbInstanceCategory() *string
	SetDbInstanceClass(v string) *CreateDifyInstanceRequest
	GetDbInstanceClass() *string
	SetDbInstancePassword(v string) *CreateDifyInstanceRequest
	GetDbInstancePassword() *string
	SetDbResourceId(v int32) *CreateDifyInstanceRequest
	GetDbResourceId() *int32
	SetDbStorageSize(v string) *CreateDifyInstanceRequest
	GetDbStorageSize() *string
	SetDbStorageType(v string) *CreateDifyInstanceRequest
	GetDbStorageType() *string
	SetDryRun(v bool) *CreateDifyInstanceRequest
	GetDryRun() *bool
	SetEdition(v string) *CreateDifyInstanceRequest
	GetEdition() *string
	SetEnableExtraEndpoint(v bool) *CreateDifyInstanceRequest
	GetEnableExtraEndpoint() *bool
	SetGpuNodeSpec(v string) *CreateDifyInstanceRequest
	GetGpuNodeSpec() *string
	SetKvStoreAccount(v string) *CreateDifyInstanceRequest
	GetKvStoreAccount() *string
	SetKvStoreEngineVersion(v string) *CreateDifyInstanceRequest
	GetKvStoreEngineVersion() *string
	SetKvStoreInstanceClass(v string) *CreateDifyInstanceRequest
	GetKvStoreInstanceClass() *string
	SetKvStoreNodeType(v string) *CreateDifyInstanceRequest
	GetKvStoreNodeType() *string
	SetKvStoreOption(v string) *CreateDifyInstanceRequest
	GetKvStoreOption() *string
	SetKvStorePassword(v string) *CreateDifyInstanceRequest
	GetKvStorePassword() *string
	SetKvStoreResourceId(v int32) *CreateDifyInstanceRequest
	GetKvStoreResourceId() *int32
	SetKvStoreType(v string) *CreateDifyInstanceRequest
	GetKvStoreType() *string
	SetMajorVersion(v string) *CreateDifyInstanceRequest
	GetMajorVersion() *string
	SetModelId(v string) *CreateDifyInstanceRequest
	GetModelId() *string
	SetModelOption(v string) *CreateDifyInstanceRequest
	GetModelOption() *string
	SetNatGatewayOption(v string) *CreateDifyInstanceRequest
	GetNatGatewayOption() *string
	SetOnlyIntranet(v bool) *CreateDifyInstanceRequest
	GetOnlyIntranet() *bool
	SetOssPath(v string) *CreateDifyInstanceRequest
	GetOssPath() *string
	SetOssResourceId(v int32) *CreateDifyInstanceRequest
	GetOssResourceId() *int32
	SetPayPeriod(v int32) *CreateDifyInstanceRequest
	GetPayPeriod() *int32
	SetPayPeriodType(v string) *CreateDifyInstanceRequest
	GetPayPeriodType() *string
	SetPayType(v string) *CreateDifyInstanceRequest
	GetPayType() *string
	SetReplicas(v int32) *CreateDifyInstanceRequest
	GetReplicas() *int32
	SetResourceQuota(v string) *CreateDifyInstanceRequest
	GetResourceQuota() *string
	SetSecurityGroupId(v string) *CreateDifyInstanceRequest
	GetSecurityGroupId() *string
	SetSegDiskPerformanceLevel(v string) *CreateDifyInstanceRequest
	GetSegDiskPerformanceLevel() *string
	SetSegNodeNum(v int32) *CreateDifyInstanceRequest
	GetSegNodeNum() *int32
	SetStorageType(v string) *CreateDifyInstanceRequest
	GetStorageType() *string
	SetVSwitchId(v string) *CreateDifyInstanceRequest
	GetVSwitchId() *string
	SetVectordbAccount(v string) *CreateDifyInstanceRequest
	GetVectordbAccount() *string
	SetVectordbCategory(v string) *CreateDifyInstanceRequest
	GetVectordbCategory() *string
	SetVectordbEngineVersion(v string) *CreateDifyInstanceRequest
	GetVectordbEngineVersion() *string
	SetVectordbInstanceSpec(v string) *CreateDifyInstanceRequest
	GetVectordbInstanceSpec() *string
	SetVectordbOption(v string) *CreateDifyInstanceRequest
	GetVectordbOption() *string
	SetVectordbPassword(v string) *CreateDifyInstanceRequest
	GetVectordbPassword() *string
	SetVectordbResourceId(v int32) *CreateDifyInstanceRequest
	GetVectordbResourceId() *int32
	SetVectordbStorageSize(v string) *CreateDifyInstanceRequest
	GetVectordbStorageSize() *string
	SetVectordbStorageType(v string) *CreateDifyInstanceRequest
	GetVectordbStorageType() *string
	SetVectordbType(v string) *CreateDifyInstanceRequest
	GetVectordbType() *string
	SetVpcId(v string) *CreateDifyInstanceRequest
	GetVpcId() *string
	SetWorkspaceDescription(v string) *CreateDifyInstanceRequest
	GetWorkspaceDescription() *string
	SetWorkspaceId(v string) *CreateDifyInstanceRequest
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *CreateDifyInstanceRequest
	GetWorkspaceName() *string
	SetWorkspaceOption(v string) *CreateDifyInstanceRequest
	GetWorkspaceOption() *string
	SetZoneId(v string) *CreateDifyInstanceRequest
	GetZoneId() *string
}

type CreateDifyInstanceRequest struct {
	AdbpgInstanceMode *string `json:"AdbpgInstanceMode,omitempty" xml:"AdbpgInstanceMode,omitempty"`
	AutoRenew         *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BackupVSwitchId   *string `json:"BackupVSwitchId,omitempty" xml:"BackupVSwitchId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DataRegion         *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	DatabaseOption     *string `json:"DatabaseOption,omitempty" xml:"DatabaseOption,omitempty"`
	DbEngineType       *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	DbEngineVersion    *string `json:"DbEngineVersion,omitempty" xml:"DbEngineVersion,omitempty"`
	DbInstanceAccount  *string `json:"DbInstanceAccount,omitempty" xml:"DbInstanceAccount,omitempty"`
	DbInstanceCategory *string `json:"DbInstanceCategory,omitempty" xml:"DbInstanceCategory,omitempty"`
	DbInstanceClass    *string `json:"DbInstanceClass,omitempty" xml:"DbInstanceClass,omitempty"`
	DbInstancePassword *string `json:"DbInstancePassword,omitempty" xml:"DbInstancePassword,omitempty"`
	DbResourceId       *int32  `json:"DbResourceId,omitempty" xml:"DbResourceId,omitempty"`
	DbStorageSize      *string `json:"DbStorageSize,omitempty" xml:"DbStorageSize,omitempty"`
	DbStorageType      *string `json:"DbStorageType,omitempty" xml:"DbStorageType,omitempty"`
	DryRun             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// Community
	Edition              *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EnableExtraEndpoint  *bool   `json:"EnableExtraEndpoint,omitempty" xml:"EnableExtraEndpoint,omitempty"`
	GpuNodeSpec          *string `json:"GpuNodeSpec,omitempty" xml:"GpuNodeSpec,omitempty"`
	KvStoreAccount       *string `json:"KvStoreAccount,omitempty" xml:"KvStoreAccount,omitempty"`
	KvStoreEngineVersion *string `json:"KvStoreEngineVersion,omitempty" xml:"KvStoreEngineVersion,omitempty"`
	KvStoreInstanceClass *string `json:"KvStoreInstanceClass,omitempty" xml:"KvStoreInstanceClass,omitempty"`
	KvStoreNodeType      *string `json:"KvStoreNodeType,omitempty" xml:"KvStoreNodeType,omitempty"`
	KvStoreOption        *string `json:"KvStoreOption,omitempty" xml:"KvStoreOption,omitempty"`
	KvStorePassword      *string `json:"KvStorePassword,omitempty" xml:"KvStorePassword,omitempty"`
	KvStoreResourceId    *int32  `json:"KvStoreResourceId,omitempty" xml:"KvStoreResourceId,omitempty"`
	KvStoreType          *string `json:"KvStoreType,omitempty" xml:"KvStoreType,omitempty"`
	MajorVersion         *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	ModelId              *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// Disable
	ModelOption      *string `json:"ModelOption,omitempty" xml:"ModelOption,omitempty"`
	NatGatewayOption *string `json:"NatGatewayOption,omitempty" xml:"NatGatewayOption,omitempty"`
	OnlyIntranet     *bool   `json:"OnlyIntranet,omitempty" xml:"OnlyIntranet,omitempty"`
	OssPath          *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	OssResourceId    *int32  `json:"OssResourceId,omitempty" xml:"OssResourceId,omitempty"`
	PayPeriod        *int32  `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	PayPeriodType    *string `json:"PayPeriodType,omitempty" xml:"PayPeriodType,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Replicas         *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// This parameter is required.
	ResourceQuota *string `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty"`
	// This parameter is required.
	SecurityGroupId         *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	SegNodeNum              *int32  `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	StorageType             *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// This parameter is required.
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VectordbAccount       *string `json:"VectordbAccount,omitempty" xml:"VectordbAccount,omitempty"`
	VectordbCategory      *string `json:"VectordbCategory,omitempty" xml:"VectordbCategory,omitempty"`
	VectordbEngineVersion *string `json:"VectordbEngineVersion,omitempty" xml:"VectordbEngineVersion,omitempty"`
	VectordbInstanceSpec  *string `json:"VectordbInstanceSpec,omitempty" xml:"VectordbInstanceSpec,omitempty"`
	VectordbOption        *string `json:"VectordbOption,omitempty" xml:"VectordbOption,omitempty"`
	VectordbPassword      *string `json:"VectordbPassword,omitempty" xml:"VectordbPassword,omitempty"`
	VectordbResourceId    *int32  `json:"VectordbResourceId,omitempty" xml:"VectordbResourceId,omitempty"`
	VectordbStorageSize   *string `json:"VectordbStorageSize,omitempty" xml:"VectordbStorageSize,omitempty"`
	VectordbStorageType   *string `json:"VectordbStorageType,omitempty" xml:"VectordbStorageType,omitempty"`
	VectordbType          *string `json:"VectordbType,omitempty" xml:"VectordbType,omitempty"`
	// This parameter is required.
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkspaceDescription *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	WorkspaceId          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName        *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	WorkspaceOption      *string `json:"WorkspaceOption,omitempty" xml:"WorkspaceOption,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDifyInstanceRequest) GetAdbpgInstanceMode() *string {
	return s.AdbpgInstanceMode
}

func (s *CreateDifyInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDifyInstanceRequest) GetBackupVSwitchId() *string {
	return s.BackupVSwitchId
}

func (s *CreateDifyInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDifyInstanceRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDifyInstanceRequest) GetDatabaseOption() *string {
	return s.DatabaseOption
}

func (s *CreateDifyInstanceRequest) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *CreateDifyInstanceRequest) GetDbEngineVersion() *string {
	return s.DbEngineVersion
}

func (s *CreateDifyInstanceRequest) GetDbInstanceAccount() *string {
	return s.DbInstanceAccount
}

func (s *CreateDifyInstanceRequest) GetDbInstanceCategory() *string {
	return s.DbInstanceCategory
}

func (s *CreateDifyInstanceRequest) GetDbInstanceClass() *string {
	return s.DbInstanceClass
}

func (s *CreateDifyInstanceRequest) GetDbInstancePassword() *string {
	return s.DbInstancePassword
}

func (s *CreateDifyInstanceRequest) GetDbResourceId() *int32 {
	return s.DbResourceId
}

func (s *CreateDifyInstanceRequest) GetDbStorageSize() *string {
	return s.DbStorageSize
}

func (s *CreateDifyInstanceRequest) GetDbStorageType() *string {
	return s.DbStorageType
}

func (s *CreateDifyInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDifyInstanceRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateDifyInstanceRequest) GetEnableExtraEndpoint() *bool {
	return s.EnableExtraEndpoint
}

func (s *CreateDifyInstanceRequest) GetGpuNodeSpec() *string {
	return s.GpuNodeSpec
}

func (s *CreateDifyInstanceRequest) GetKvStoreAccount() *string {
	return s.KvStoreAccount
}

func (s *CreateDifyInstanceRequest) GetKvStoreEngineVersion() *string {
	return s.KvStoreEngineVersion
}

func (s *CreateDifyInstanceRequest) GetKvStoreInstanceClass() *string {
	return s.KvStoreInstanceClass
}

func (s *CreateDifyInstanceRequest) GetKvStoreNodeType() *string {
	return s.KvStoreNodeType
}

func (s *CreateDifyInstanceRequest) GetKvStoreOption() *string {
	return s.KvStoreOption
}

func (s *CreateDifyInstanceRequest) GetKvStorePassword() *string {
	return s.KvStorePassword
}

func (s *CreateDifyInstanceRequest) GetKvStoreResourceId() *int32 {
	return s.KvStoreResourceId
}

func (s *CreateDifyInstanceRequest) GetKvStoreType() *string {
	return s.KvStoreType
}

func (s *CreateDifyInstanceRequest) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *CreateDifyInstanceRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateDifyInstanceRequest) GetModelOption() *string {
	return s.ModelOption
}

func (s *CreateDifyInstanceRequest) GetNatGatewayOption() *string {
	return s.NatGatewayOption
}

func (s *CreateDifyInstanceRequest) GetOnlyIntranet() *bool {
	return s.OnlyIntranet
}

func (s *CreateDifyInstanceRequest) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateDifyInstanceRequest) GetOssResourceId() *int32 {
	return s.OssResourceId
}

func (s *CreateDifyInstanceRequest) GetPayPeriod() *int32 {
	return s.PayPeriod
}

func (s *CreateDifyInstanceRequest) GetPayPeriodType() *string {
	return s.PayPeriodType
}

func (s *CreateDifyInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDifyInstanceRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateDifyInstanceRequest) GetResourceQuota() *string {
	return s.ResourceQuota
}

func (s *CreateDifyInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateDifyInstanceRequest) GetSegDiskPerformanceLevel() *string {
	return s.SegDiskPerformanceLevel
}

func (s *CreateDifyInstanceRequest) GetSegNodeNum() *int32 {
	return s.SegNodeNum
}

func (s *CreateDifyInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDifyInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDifyInstanceRequest) GetVectordbAccount() *string {
	return s.VectordbAccount
}

func (s *CreateDifyInstanceRequest) GetVectordbCategory() *string {
	return s.VectordbCategory
}

func (s *CreateDifyInstanceRequest) GetVectordbEngineVersion() *string {
	return s.VectordbEngineVersion
}

func (s *CreateDifyInstanceRequest) GetVectordbInstanceSpec() *string {
	return s.VectordbInstanceSpec
}

func (s *CreateDifyInstanceRequest) GetVectordbOption() *string {
	return s.VectordbOption
}

func (s *CreateDifyInstanceRequest) GetVectordbPassword() *string {
	return s.VectordbPassword
}

func (s *CreateDifyInstanceRequest) GetVectordbResourceId() *int32 {
	return s.VectordbResourceId
}

func (s *CreateDifyInstanceRequest) GetVectordbStorageSize() *string {
	return s.VectordbStorageSize
}

func (s *CreateDifyInstanceRequest) GetVectordbStorageType() *string {
	return s.VectordbStorageType
}

func (s *CreateDifyInstanceRequest) GetVectordbType() *string {
	return s.VectordbType
}

func (s *CreateDifyInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDifyInstanceRequest) GetWorkspaceDescription() *string {
	return s.WorkspaceDescription
}

func (s *CreateDifyInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDifyInstanceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateDifyInstanceRequest) GetWorkspaceOption() *string {
	return s.WorkspaceOption
}

func (s *CreateDifyInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDifyInstanceRequest) SetAdbpgInstanceMode(v string) *CreateDifyInstanceRequest {
	s.AdbpgInstanceMode = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetAutoRenew(v bool) *CreateDifyInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetBackupVSwitchId(v string) *CreateDifyInstanceRequest {
	s.BackupVSwitchId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetClientToken(v string) *CreateDifyInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDataRegion(v string) *CreateDifyInstanceRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDatabaseOption(v string) *CreateDifyInstanceRequest {
	s.DatabaseOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbEngineType(v string) *CreateDifyInstanceRequest {
	s.DbEngineType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbEngineVersion(v string) *CreateDifyInstanceRequest {
	s.DbEngineVersion = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbInstanceAccount(v string) *CreateDifyInstanceRequest {
	s.DbInstanceAccount = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbInstanceCategory(v string) *CreateDifyInstanceRequest {
	s.DbInstanceCategory = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbInstanceClass(v string) *CreateDifyInstanceRequest {
	s.DbInstanceClass = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbInstancePassword(v string) *CreateDifyInstanceRequest {
	s.DbInstancePassword = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbResourceId(v int32) *CreateDifyInstanceRequest {
	s.DbResourceId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbStorageSize(v string) *CreateDifyInstanceRequest {
	s.DbStorageSize = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDbStorageType(v string) *CreateDifyInstanceRequest {
	s.DbStorageType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetDryRun(v bool) *CreateDifyInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetEdition(v string) *CreateDifyInstanceRequest {
	s.Edition = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetEnableExtraEndpoint(v bool) *CreateDifyInstanceRequest {
	s.EnableExtraEndpoint = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetGpuNodeSpec(v string) *CreateDifyInstanceRequest {
	s.GpuNodeSpec = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreAccount(v string) *CreateDifyInstanceRequest {
	s.KvStoreAccount = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreEngineVersion(v string) *CreateDifyInstanceRequest {
	s.KvStoreEngineVersion = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreInstanceClass(v string) *CreateDifyInstanceRequest {
	s.KvStoreInstanceClass = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreNodeType(v string) *CreateDifyInstanceRequest {
	s.KvStoreNodeType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreOption(v string) *CreateDifyInstanceRequest {
	s.KvStoreOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStorePassword(v string) *CreateDifyInstanceRequest {
	s.KvStorePassword = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreResourceId(v int32) *CreateDifyInstanceRequest {
	s.KvStoreResourceId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetKvStoreType(v string) *CreateDifyInstanceRequest {
	s.KvStoreType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetMajorVersion(v string) *CreateDifyInstanceRequest {
	s.MajorVersion = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetModelId(v string) *CreateDifyInstanceRequest {
	s.ModelId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetModelOption(v string) *CreateDifyInstanceRequest {
	s.ModelOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetNatGatewayOption(v string) *CreateDifyInstanceRequest {
	s.NatGatewayOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetOnlyIntranet(v bool) *CreateDifyInstanceRequest {
	s.OnlyIntranet = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetOssPath(v string) *CreateDifyInstanceRequest {
	s.OssPath = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetOssResourceId(v int32) *CreateDifyInstanceRequest {
	s.OssResourceId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetPayPeriod(v int32) *CreateDifyInstanceRequest {
	s.PayPeriod = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetPayPeriodType(v string) *CreateDifyInstanceRequest {
	s.PayPeriodType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetPayType(v string) *CreateDifyInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetReplicas(v int32) *CreateDifyInstanceRequest {
	s.Replicas = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetResourceQuota(v string) *CreateDifyInstanceRequest {
	s.ResourceQuota = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetSecurityGroupId(v string) *CreateDifyInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetSegDiskPerformanceLevel(v string) *CreateDifyInstanceRequest {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetSegNodeNum(v int32) *CreateDifyInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetStorageType(v string) *CreateDifyInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVSwitchId(v string) *CreateDifyInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbAccount(v string) *CreateDifyInstanceRequest {
	s.VectordbAccount = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbCategory(v string) *CreateDifyInstanceRequest {
	s.VectordbCategory = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbEngineVersion(v string) *CreateDifyInstanceRequest {
	s.VectordbEngineVersion = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbInstanceSpec(v string) *CreateDifyInstanceRequest {
	s.VectordbInstanceSpec = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbOption(v string) *CreateDifyInstanceRequest {
	s.VectordbOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbPassword(v string) *CreateDifyInstanceRequest {
	s.VectordbPassword = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbResourceId(v int32) *CreateDifyInstanceRequest {
	s.VectordbResourceId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbStorageSize(v string) *CreateDifyInstanceRequest {
	s.VectordbStorageSize = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbStorageType(v string) *CreateDifyInstanceRequest {
	s.VectordbStorageType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVectordbType(v string) *CreateDifyInstanceRequest {
	s.VectordbType = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetVpcId(v string) *CreateDifyInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetWorkspaceDescription(v string) *CreateDifyInstanceRequest {
	s.WorkspaceDescription = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetWorkspaceId(v string) *CreateDifyInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetWorkspaceName(v string) *CreateDifyInstanceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetWorkspaceOption(v string) *CreateDifyInstanceRequest {
	s.WorkspaceOption = &v
	return s
}

func (s *CreateDifyInstanceRequest) SetZoneId(v string) *CreateDifyInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}

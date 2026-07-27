// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticDbClusterDescription(v string) *CreateDBClusterRequest
	GetAgenticDbClusterDescription() *string
	SetAgenticDbClusterId(v string) *CreateDBClusterRequest
	GetAgenticDbClusterId() *string
	SetAgenticDbType(v string) *CreateDBClusterRequest
	GetAgenticDbType() *string
	SetAllowShutDown(v string) *CreateDBClusterRequest
	GetAllowShutDown() *string
	SetArchitecture(v string) *CreateDBClusterRequest
	GetArchitecture() *string
	SetAutoRenew(v bool) *CreateDBClusterRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreateDBClusterRequest
	GetAutoUseCoupon() *bool
	SetBackupRetentionPolicyOnClusterDeletion(v string) *CreateDBClusterRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetBurstingEnabled(v string) *CreateDBClusterRequest
	GetBurstingEnabled() *string
	SetClientToken(v string) *CreateDBClusterRequest
	GetClientToken() *string
	SetCloneDataPoint(v string) *CreateDBClusterRequest
	GetCloneDataPoint() *string
	SetCloudProvider(v string) *CreateDBClusterRequest
	GetCloudProvider() *string
	SetClusterNetworkType(v string) *CreateDBClusterRequest
	GetClusterNetworkType() *string
	SetCreationCategory(v string) *CreateDBClusterRequest
	GetCreationCategory() *string
	SetCreationOption(v string) *CreateDBClusterRequest
	GetCreationOption() *string
	SetDBClusterDescription(v string) *CreateDBClusterRequest
	GetDBClusterDescription() *string
	SetDBMinorVersion(v string) *CreateDBClusterRequest
	GetDBMinorVersion() *string
	SetDBNodeClass(v string) *CreateDBClusterRequest
	GetDBNodeClass() *string
	SetDBNodeNum(v int32) *CreateDBClusterRequest
	GetDBNodeNum() *int32
	SetDBType(v string) *CreateDBClusterRequest
	GetDBType() *string
	SetDBVersion(v string) *CreateDBClusterRequest
	GetDBVersion() *string
	SetDefaultTimeZone(v string) *CreateDBClusterRequest
	GetDefaultTimeZone() *string
	SetEnsRegionId(v string) *CreateDBClusterRequest
	GetEnsRegionId() *string
	SetGDNId(v string) *CreateDBClusterRequest
	GetGDNId() *string
	SetHotStandbyCluster(v string) *CreateDBClusterRequest
	GetHotStandbyCluster() *string
	SetLoosePolarLogBin(v string) *CreateDBClusterRequest
	GetLoosePolarLogBin() *string
	SetLooseXEngine(v string) *CreateDBClusterRequest
	GetLooseXEngine() *string
	SetLooseXEngineUseMemoryPct(v string) *CreateDBClusterRequest
	GetLooseXEngineUseMemoryPct() *string
	SetLowerCaseTableNames(v string) *CreateDBClusterRequest
	GetLowerCaseTableNames() *string
	SetOwnerAccount(v string) *CreateDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBClusterRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *CreateDBClusterRequest
	GetParameterGroupId() *string
	SetPayType(v string) *CreateDBClusterRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBClusterRequest
	GetPeriod() *string
	SetPromotionCode(v string) *CreateDBClusterRequest
	GetPromotionCode() *string
	SetProvisionedIops(v int64) *CreateDBClusterRequest
	GetProvisionedIops() *int64
	SetProxyClass(v string) *CreateDBClusterRequest
	GetProxyClass() *string
	SetProxyType(v string) *CreateDBClusterRequest
	GetProxyType() *string
	SetRegionId(v string) *CreateDBClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBClusterRequest
	GetResourceOwnerId() *int64
	SetScaleMax(v string) *CreateDBClusterRequest
	GetScaleMax() *string
	SetScaleMin(v string) *CreateDBClusterRequest
	GetScaleMin() *string
	SetScaleRoNumMax(v string) *CreateDBClusterRequest
	GetScaleRoNumMax() *string
	SetScaleRoNumMin(v string) *CreateDBClusterRequest
	GetScaleRoNumMin() *string
	SetSecurityIPList(v string) *CreateDBClusterRequest
	GetSecurityIPList() *string
	SetServerlessType(v string) *CreateDBClusterRequest
	GetServerlessType() *string
	SetSourceResourceId(v string) *CreateDBClusterRequest
	GetSourceResourceId() *string
	SetSourceUid(v int64) *CreateDBClusterRequest
	GetSourceUid() *int64
	SetStandbyAZ(v string) *CreateDBClusterRequest
	GetStandbyAZ() *string
	SetStorageAutoScale(v string) *CreateDBClusterRequest
	GetStorageAutoScale() *string
	SetStorageEncryption(v bool) *CreateDBClusterRequest
	GetStorageEncryption() *bool
	SetStorageEncryptionKey(v string) *CreateDBClusterRequest
	GetStorageEncryptionKey() *string
	SetStoragePayType(v string) *CreateDBClusterRequest
	GetStoragePayType() *string
	SetStorageSpace(v int64) *CreateDBClusterRequest
	GetStorageSpace() *int64
	SetStorageType(v string) *CreateDBClusterRequest
	GetStorageType() *string
	SetStorageUpperBound(v int64) *CreateDBClusterRequest
	GetStorageUpperBound() *int64
	SetStrictConsistency(v string) *CreateDBClusterRequest
	GetStrictConsistency() *string
	SetTDEStatus(v bool) *CreateDBClusterRequest
	GetTDEStatus() *bool
	SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest
	GetTag() []*CreateDBClusterRequestTag
	SetTargetMinorVersion(v string) *CreateDBClusterRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v string) *CreateDBClusterRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDBClusterRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBClusterRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBClusterRequest
	GetZoneId() *string
}

type CreateDBClusterRequest struct {
	// The Agentic cluster description.
	//
	// example:
	//
	// pagc-******
	AgenticDbClusterDescription *string `json:"AgenticDbClusterDescription,omitempty" xml:"AgenticDbClusterDescription,omitempty"`
	// The Agentic cluster ID.
	//
	// example:
	//
	// pagc-******
	AgenticDbClusterId *string `json:"AgenticDbClusterId,omitempty" xml:"AgenticDbClusterId,omitempty"`
	// The Agentic database engine type. Valid values:
	//
	// example:
	//
	// dedicated
	AgenticDbType *string `json:"AgenticDbType,omitempty" xml:"AgenticDbType,omitempty"`
	// Specifies whether to enable No-activity Suspension. Valid values:
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// The CPU architecture. Valid values:
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The data retention policy for backups when the cluster is deleted. Valid values:
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// Specifies whether to enable I/O performance burst for the ESSD AutoPL cloud disk. Valid values:
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value of this parameter is generated by the client and is unique among different requests. The value is case-sensitive and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The point in time at which data is cloned. Valid values:
	//
	// example:
	//
	// LATEST
	CloneDataPoint *string `json:"CloneDataPoint,omitempty" xml:"CloneDataPoint,omitempty"`
	// The cloud service provider to which the instance belongs.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The network type of the cluster. Currently, only Virtual Private Cloud (VPC) is supported. Set the value to **VPC**.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The method used to create the cluster. Valid values:
	//
	// example:
	//
	// Normal
	CreationOption *string `json:"CreationOption,omitempty" xml:"CreationOption,omitempty"`
	// The cluster name. The cluster name must meet the following requirements:
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The minor version of the database engine. Valid values:
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// The node specifications. For details, see the following topics:
	//
	// example:
	//
	// polar.mysql.x4.medium
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes for Standard Edition and Enterprise Edition. Valid values:
	//
	// example:
	//
	// 1
	DBNodeNum *int32 `json:"DBNodeNum,omitempty" xml:"DBNodeNum,omitempty"`
	// The database engine type. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The time zone of the cluster (UTC). The value can be set to any time frame within the range of **-12:00 to +13:00**, such as **00:00**. Default value: **SYSTEM**. The default time zone is the same as the time zone of the region.
	//
	// example:
	//
	// SYSTEM
	DefaultTimeZone *string `json:"DefaultTimeZone,omitempty" xml:"DefaultTimeZone,omitempty"`
	// The ENS node ID required when creating an ENS database.
	//
	// example:
	//
	// vn-hanoi-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The global database network (GDN) ID.
	//
	// example:
	//
	// gdn-***********
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// Specifies whether the hot standby cluster is enabled. Valid values:
	//
	// example:
	//
	// ON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// Specifies whether to enable the Binlog feature. Valid values:
	//
	// example:
	//
	// ON
	LoosePolarLogBin *string `json:"LoosePolarLogBin,omitempty" xml:"LoosePolarLogBin,omitempty"`
	// Specifies whether to enable the X-Engine storage engine. Valid values:
	//
	// example:
	//
	// ON
	LooseXEngine *string `json:"LooseXEngine,omitempty" xml:"LooseXEngine,omitempty"`
	// The percentage of memory allocated to the X-Engine storage engine. Valid values: integers from 10 to 90.
	//
	// example:
	//
	// 50
	LooseXEngineUseMemoryPct *string `json:"LooseXEngineUseMemoryPct,omitempty" xml:"LooseXEngineUseMemoryPct,omitempty"`
	// Specifies whether table names are case-sensitive. Valid values:
	//
	// example:
	//
	// 1
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitempty" xml:"LowerCaseTableNames,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The billing method. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required when PayType is set to **Prepaid**. Specifies whether the subscription cluster uses a yearly or monthly billing cycle.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code. If not specified, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// <p id="p_wyg_t4a_glm" props="china" icmsditafragmentmagic=1>The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}.</p>
	//
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The specification of the database proxy for Standard Edition. Valid values:
	//
	// example:
	//
	// polar.maxscale.g2.medium.c
	ProxyClass *string `json:"ProxyClass,omitempty" xml:"ProxyClass,omitempty"`
	// The database proxy type. Valid values:
	//
	// example:
	//
	// Exclusive
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum scaling limit per node. Valid values: 1 PCU to 32 PCUs.
	//
	// example:
	//
	// 3
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum scaling limit per node. Valid values: 1 PCU to 31 PCUs.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The maximum number of read-only nodes for scaling. Valid values: 0 to 15.
	//
	// example:
	//
	// 4
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// The minimum number of read-only nodes for scaling. Valid values: 0 to 15.
	//
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// The IP whitelist of the PolarDB cluster.
	//
	// example:
	//
	// 10.***.***.***
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The serverless type. The value is fixed as **AgileServerless**.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The instance ID of the source ApsaraDB RDS instance or the source PolarDB cluster. This parameter is required only when **CreationOption*	- is set to **MigrationFromRDS**, **CloneFromRDS**, **CloneFromPolarDB**, or **RecoverFromRecyclebin**.
	//
	// example:
	//
	// rm-*************
	SourceResourceId *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	// The UID of the account that owns the source backup set in cross-account backup restoration scenarios.
	//
	// example:
	//
	// 1022xxxxxxxx
	SourceUid *int64 `json:"SourceUid,omitempty" xml:"SourceUid,omitempty"`
	// The zone of the hot standby cluster.
	//
	// example:
	//
	// cn-hangzhou-g
	StandbyAZ *string `json:"StandbyAZ,omitempty" xml:"StandbyAZ,omitempty"`
	// Specifies whether to enable automatic storage scaling for the Standard Edition cluster. Valid values:
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// Specifies whether to enable cloud disk encryption. Valid values:
	StorageEncryption *bool `json:"StorageEncryption,omitempty" xml:"StorageEncryption,omitempty"`
	// The key ID of the custom Key Management Service (KMS) key for cloud disk encryption in the same region as the instance. Specifying this parameter automatically enables cloud disk encryption, which cannot be disabled after being enabled. Leave this parameter empty to use the default service key for cloud disk encryption.
	//
	// example:
	//
	// 1022xxxxxxxx
	StorageEncryptionKey *string `json:"StorageEncryptionKey,omitempty" xml:"StorageEncryptionKey,omitempty"`
	// The billing type for storage. Valid values:
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage space for the space-based billing (subscription) plan. Unit: GB.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type for Enterprise Edition. Valid values:
	//
	// example:
	//
	// PSL4
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The upper limit of automatic storage scaling for the Standard Edition cluster. Unit: GB.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// Specifies whether to enable multi-zone strong data consistency for the cluster. Valid values:
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// Specifies whether to enable Transparent Data Encryption (TDE). Valid values:
	//
	// example:
	//
	// true
	TDEStatus *bool `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	// The list of tags.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The target minor engine version.
	//
	// example:
	//
	// 8.0.1.1.54
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// This parameter is required when PayType is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) GetAgenticDbClusterDescription() *string {
	return s.AgenticDbClusterDescription
}

func (s *CreateDBClusterRequest) GetAgenticDbClusterId() *string {
	return s.AgenticDbClusterId
}

func (s *CreateDBClusterRequest) GetAgenticDbType() *string {
	return s.AgenticDbType
}

func (s *CreateDBClusterRequest) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *CreateDBClusterRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateDBClusterRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBClusterRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateDBClusterRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *CreateDBClusterRequest) GetBurstingEnabled() *string {
	return s.BurstingEnabled
}

func (s *CreateDBClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBClusterRequest) GetCloneDataPoint() *string {
	return s.CloneDataPoint
}

func (s *CreateDBClusterRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *CreateDBClusterRequest) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *CreateDBClusterRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *CreateDBClusterRequest) GetCreationOption() *string {
	return s.CreationOption
}

func (s *CreateDBClusterRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *CreateDBClusterRequest) GetDBMinorVersion() *string {
	return s.DBMinorVersion
}

func (s *CreateDBClusterRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateDBClusterRequest) GetDBNodeNum() *int32 {
	return s.DBNodeNum
}

func (s *CreateDBClusterRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreateDBClusterRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *CreateDBClusterRequest) GetDefaultTimeZone() *string {
	return s.DefaultTimeZone
}

func (s *CreateDBClusterRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateDBClusterRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *CreateDBClusterRequest) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *CreateDBClusterRequest) GetLoosePolarLogBin() *string {
	return s.LoosePolarLogBin
}

func (s *CreateDBClusterRequest) GetLooseXEngine() *string {
	return s.LooseXEngine
}

func (s *CreateDBClusterRequest) GetLooseXEngineUseMemoryPct() *string {
	return s.LooseXEngineUseMemoryPct
}

func (s *CreateDBClusterRequest) GetLowerCaseTableNames() *string {
	return s.LowerCaseTableNames
}

func (s *CreateDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBClusterRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *CreateDBClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBClusterRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBClusterRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateDBClusterRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateDBClusterRequest) GetProxyClass() *string {
	return s.ProxyClass
}

func (s *CreateDBClusterRequest) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateDBClusterRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateDBClusterRequest) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *CreateDBClusterRequest) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *CreateDBClusterRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBClusterRequest) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *CreateDBClusterRequest) GetSourceResourceId() *string {
	return s.SourceResourceId
}

func (s *CreateDBClusterRequest) GetSourceUid() *int64 {
	return s.SourceUid
}

func (s *CreateDBClusterRequest) GetStandbyAZ() *string {
	return s.StandbyAZ
}

func (s *CreateDBClusterRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *CreateDBClusterRequest) GetStorageEncryption() *bool {
	return s.StorageEncryption
}

func (s *CreateDBClusterRequest) GetStorageEncryptionKey() *string {
	return s.StorageEncryptionKey
}

func (s *CreateDBClusterRequest) GetStoragePayType() *string {
	return s.StoragePayType
}

func (s *CreateDBClusterRequest) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *CreateDBClusterRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBClusterRequest) GetStorageUpperBound() *int64 {
	return s.StorageUpperBound
}

func (s *CreateDBClusterRequest) GetStrictConsistency() *string {
	return s.StrictConsistency
}

func (s *CreateDBClusterRequest) GetTDEStatus() *bool {
	return s.TDEStatus
}

func (s *CreateDBClusterRequest) GetTag() []*CreateDBClusterRequestTag {
	return s.Tag
}

func (s *CreateDBClusterRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *CreateDBClusterRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBClusterRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBClusterRequest) SetAgenticDbClusterDescription(v string) *CreateDBClusterRequest {
	s.AgenticDbClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetAgenticDbClusterId(v string) *CreateDBClusterRequest {
	s.AgenticDbClusterId = &v
	return s
}

func (s *CreateDBClusterRequest) SetAgenticDbType(v string) *CreateDBClusterRequest {
	s.AgenticDbType = &v
	return s
}

func (s *CreateDBClusterRequest) SetAllowShutDown(v string) *CreateDBClusterRequest {
	s.AllowShutDown = &v
	return s
}

func (s *CreateDBClusterRequest) SetArchitecture(v string) *CreateDBClusterRequest {
	s.Architecture = &v
	return s
}

func (s *CreateDBClusterRequest) SetAutoRenew(v bool) *CreateDBClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBClusterRequest) SetAutoUseCoupon(v bool) *CreateDBClusterRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateDBClusterRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *CreateDBClusterRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *CreateDBClusterRequest) SetBurstingEnabled(v string) *CreateDBClusterRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetCloneDataPoint(v string) *CreateDBClusterRequest {
	s.CloneDataPoint = &v
	return s
}

func (s *CreateDBClusterRequest) SetCloudProvider(v string) *CreateDBClusterRequest {
	s.CloudProvider = &v
	return s
}

func (s *CreateDBClusterRequest) SetClusterNetworkType(v string) *CreateDBClusterRequest {
	s.ClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationCategory(v string) *CreateDBClusterRequest {
	s.CreationCategory = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationOption(v string) *CreateDBClusterRequest {
	s.CreationOption = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBMinorVersion(v string) *CreateDBClusterRequest {
	s.DBMinorVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeClass(v string) *CreateDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeNum(v int32) *CreateDBClusterRequest {
	s.DBNodeNum = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBType(v string) *CreateDBClusterRequest {
	s.DBType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBVersion(v string) *CreateDBClusterRequest {
	s.DBVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDefaultTimeZone(v string) *CreateDBClusterRequest {
	s.DefaultTimeZone = &v
	return s
}

func (s *CreateDBClusterRequest) SetEnsRegionId(v string) *CreateDBClusterRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetGDNId(v string) *CreateDBClusterRequest {
	s.GDNId = &v
	return s
}

func (s *CreateDBClusterRequest) SetHotStandbyCluster(v string) *CreateDBClusterRequest {
	s.HotStandbyCluster = &v
	return s
}

func (s *CreateDBClusterRequest) SetLoosePolarLogBin(v string) *CreateDBClusterRequest {
	s.LoosePolarLogBin = &v
	return s
}

func (s *CreateDBClusterRequest) SetLooseXEngine(v string) *CreateDBClusterRequest {
	s.LooseXEngine = &v
	return s
}

func (s *CreateDBClusterRequest) SetLooseXEngineUseMemoryPct(v string) *CreateDBClusterRequest {
	s.LooseXEngineUseMemoryPct = &v
	return s
}

func (s *CreateDBClusterRequest) SetLowerCaseTableNames(v string) *CreateDBClusterRequest {
	s.LowerCaseTableNames = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetParameterGroupId(v string) *CreateDBClusterRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetPromotionCode(v string) *CreateDBClusterRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateDBClusterRequest) SetProvisionedIops(v int64) *CreateDBClusterRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateDBClusterRequest) SetProxyClass(v string) *CreateDBClusterRequest {
	s.ProxyClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetProxyType(v string) *CreateDBClusterRequest {
	s.ProxyType = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleMax(v string) *CreateDBClusterRequest {
	s.ScaleMax = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleMin(v string) *CreateDBClusterRequest {
	s.ScaleMin = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleRoNumMax(v string) *CreateDBClusterRequest {
	s.ScaleRoNumMax = &v
	return s
}

func (s *CreateDBClusterRequest) SetScaleRoNumMin(v string) *CreateDBClusterRequest {
	s.ScaleRoNumMin = &v
	return s
}

func (s *CreateDBClusterRequest) SetSecurityIPList(v string) *CreateDBClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBClusterRequest) SetServerlessType(v string) *CreateDBClusterRequest {
	s.ServerlessType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceResourceId(v string) *CreateDBClusterRequest {
	s.SourceResourceId = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceUid(v int64) *CreateDBClusterRequest {
	s.SourceUid = &v
	return s
}

func (s *CreateDBClusterRequest) SetStandbyAZ(v string) *CreateDBClusterRequest {
	s.StandbyAZ = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageAutoScale(v string) *CreateDBClusterRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageEncryption(v bool) *CreateDBClusterRequest {
	s.StorageEncryption = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageEncryptionKey(v string) *CreateDBClusterRequest {
	s.StorageEncryptionKey = &v
	return s
}

func (s *CreateDBClusterRequest) SetStoragePayType(v string) *CreateDBClusterRequest {
	s.StoragePayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageSpace(v int64) *CreateDBClusterRequest {
	s.StorageSpace = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageType(v string) *CreateDBClusterRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageUpperBound(v int64) *CreateDBClusterRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *CreateDBClusterRequest) SetStrictConsistency(v string) *CreateDBClusterRequest {
	s.StrictConsistency = &v
	return s
}

func (s *CreateDBClusterRequest) SetTDEStatus(v bool) *CreateDBClusterRequest {
	s.TDEStatus = &v
	return s
}

func (s *CreateDBClusterRequest) SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDBClusterRequest) SetTargetMinorVersion(v string) *CreateDBClusterRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBClusterRequestTag struct {
	// The tag key. To add multiple tags to the cluster at a time, click **Add*	- to add tag keys.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. To add multiple tags to the cluster at a time, click **Add*	- to add tag values.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBClusterRequestTag) SetKey(v string) *CreateDBClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBClusterRequestTag) SetValue(v string) *CreateDBClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBClusterRequestTag) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowShutDown(v string) *CreateDBClusterRequest
	GetAllowShutDown() *string
	SetArchitecture(v string) *CreateDBClusterRequest
	GetArchitecture() *string
	SetAutoRenew(v bool) *CreateDBClusterRequest
	GetAutoRenew() *bool
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
	// Whether to enable idle pause. Values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled (default)
	//
	// > Only supported by Serverless clusters.
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// CPU architecture. Available values include:
	//
	// - X86
	//
	// - ARM
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Whether to enable auto-renewal, with available values as follows:
	//
	// - **true**: Auto-renew.
	//
	// - **false**: Do not auto-renew.
	//
	// The default is **false**.
	//
	// > This parameter takes effect only when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Backup retention policy upon cluster deletion, with valid values as follows:
	//
	// 	- **ALL**: Permanently retain all backups.
	//
	// 	- **LATEST**: Permanently retain the latest backup (automatically backed up before deletion).
	//
	// 	- **NONE**: Do not retain backup sets upon cluster deletion.
	//
	// By default, the value is set to **NONE**, indicating no backup sets are retained upon cluster deletion.
	//
	// > This parameter applies only when **DBType*	- is **MySQL**.
	//
	// > Serverless clusters do not support this parameter.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Used to ensure idempotency of the request. Generated by the client, ensuring uniqueness across different requests, case-sensitive, and not exceeding 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The point in time to clone data, with the following options:
	//
	// - **LATEST**: Data from the latest time point.
	//
	// - **BackupID**: Historical backup set ID, please enter the specific backup set ID.
	//
	// - **Timestamp**: Historical time point, please enter the specific time in the format `YYYY-MM-DDThh:mm:ssZ` (UTC time).
	//
	// The default value is **LATEST**.
	//
	// > If **CreationOption*	- is **CloneFromRDS**, this parameter can only be set to **LATEST**.
	//
	// example:
	//
	// LATEST
	CloneDataPoint *string `json:"CloneDataPoint,omitempty" xml:"CloneDataPoint,omitempty"`
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// Cluster network type, currently only VPC is supported, with a fixed value of **VPC**.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// Product series, with valid values as follows:
	//
	// 	- **Normal**: Cluster Edition (default)
	//
	// 	- **Basic**: Single-node
	//
	// 	- **ArchiveNormal**: High Compression Engine (X-Engine)
	//
	// 	- **NormalMultimaster**: Multi-master Cluster Edition
	//
	// 	- **SENormal**: Standard Edition
	//
	// > 	- **MySQL*	- **5.6**, **5.7**, **8.0**, **PostgreSQL*	- **14**, and **Oracle Syntax Compatible 2.0*	- support **Basic**.
	//
	// > 	- **MySQL*	- **8.0*	- supports **ArchiveNormal*	- and **NormalMultimaster**.
	//
	// > 	- **MySQL*	- **5.6**, **5.7**, **8.0**, and **PostgreSQL*	- **14*	- support **SENormal**.
	//
	// For more information about product series, see [Product Series](https://help.aliyun.com/document_detail/183258.html).
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// Creation method, with the following values supported:
	//
	// 	- **Normal**: Creates a brand new PolarDB cluster. For console operations, refer to the following documents:
	//
	//     	- [Create a PolarDB MySQL Edition Database Cluster](https://help.aliyun.com/document_detail/58769.html)
	//
	//     	- [Create a PolarDB PostgreSQL Edition Database Cluster](https://help.aliyun.com/document_detail/118063.html)
	//
	//     	- [Create a PolarDB PostgreSQL Edition (Oracle Compatible) Database Cluster](https://help.aliyun.com/document_detail/118182.html)
	//
	// 	- **CloneFromPolarDB**: Clones data from an existing PolarDB cluster to a new PolarDB cluster. For console operations, refer to the following documents:
	//
	//     	- [Clone a PolarDB MySQL Edition Cluster](https://help.aliyun.com/document_detail/87966.html)
	//
	//     	- [Clone a PolarDB PostgreSQL Edition Cluster](https://help.aliyun.com/document_detail/118108.html)
	//
	//     	- [Clone a PolarDB PostgreSQL Edition (Oracle Compatible) Cluster](https://help.aliyun.com/document_detail/118221.html)
	//
	// 	- **RecoverFromRecyclebin**: Recovers data from a released PolarDB cluster to a new PolarDB cluster. For console operations, refer to the following documents:
	//
	//     	- [Restore a Released PolarDB MySQL Edition Cluster](https://help.aliyun.com/document_detail/164880.html)
	//
	//     	- [Restore a Released PolarDB PostgreSQL Edition Cluster](https://help.aliyun.com/document_detail/432844.html)
	//
	//     	- [Restore a Released PolarDB PostgreSQL Edition (Oracle Compatible) Cluster](https://help.aliyun.com/document_detail/424632.html)
	//
	// 	- **CloneFromRDS**: Clones data from an existing RDS instance to a new PolarDB cluster. Console operation guide is available at [One-click Clone from RDS MySQL to PolarDB MySQL Edition](https://help.aliyun.com/document_detail/121812.html).
	//
	// 	- **MigrationFromRDS**: Migrates data from an existing RDS instance to a new PolarDB cluster. The created PolarDB cluster operates in read-only mode with Binlog enabled by default. Console operation guide is at [One-click Upgrade from RDS MySQL to PolarDB MySQL Edition](https://help.aliyun.com/document_detail/121582.html).
	//
	// 	- **CreateGdnStandby**: Creates a standby cluster. Console operation guide can be found at [Add Standby Cluster](https://help.aliyun.com/document_detail/160381.html).
	//
	// 	- **UpgradeFromPolarDB**: Upgrades and migrates from PolarDB. Console operation guide is detailed in [Major Version Upgrade](https://help.aliyun.com/document_detail/459712.html).
	//
	// The default value is **Normal**.
	//
	// > When **DBType*	- is **MySQL*	- and **DBVersion*	- is **8.0**, this parameter can also take the value **CreateGdnStandby**.
	//
	// example:
	//
	// Normal
	CreationOption *string `json:"CreationOption,omitempty" xml:"CreationOption,omitempty"`
	// Cluster name, which must meet the following requirements:
	//
	// 	- Cannot start with `http://` or `https://`.
	//
	// 	- Length should be between 2 and 256 characters.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// Database engine minor version number. Valid values include:
	//
	// - **8.0.2**
	//
	// - **8.0.1**
	//
	// > This parameter takes effect only when **DBType*	- is **MySQL*	- and **DBVersion*	- is **8.0**.
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// Node specifications. For details, refer to the following documents:
	//
	// - PolarDB MySQL Edition: [Compute Node Specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// - PolarDB PostgreSQL Edition (Oracle Compatible): [Compute Node Specifications](https://help.aliyun.com/document_detail/207921.html).
	//
	// - PolarDB PostgreSQL Edition: [Compute Node Specifications](https://help.aliyun.com/document_detail/209380.html).
	//
	// > - For a Serverless cluster in PolarDB MySQL, enter **polar.mysql.sl.small**.
	//
	// <props="china">> - For a Serverless cluster in both PolarDB PostgreSQL (Oracle Compatible) and PolarDB PostgreSQL, enter **polar.pg.sl.small.c**.
	//
	// example:
	//
	// polar.mysql.x4.medium
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes. This parameter is supported for Standard Edition clusters. Valid values:
	//
	// 	- **1*	- (default): only one primary node.
	//
	// 	- **2**: one read-only node and one primary node.
	//
	// >
	//
	// 	- By default, an Enterprise Edition cluster has two nodes and a Standard Edition cluster has one node.
	//
	// 	- This parameter is supported only for PolarDB for MySQL clusters.
	//
	// example:
	//
	// 1
	DBNodeNum *int32 `json:"DBNodeNum,omitempty" xml:"DBNodeNum,omitempty"`
	// Database engine type, with available values as follows:
	//
	// - **MySQL**
	//
	// - **PostgreSQL**
	//
	// - **Oracle**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// Database engine version number.
	//
	// 	- For MySQL, the version numbers are as follows:
	//
	//     	- **5.6**
	//
	//     	- **5.7**
	//
	//     	- **8.0**
	//
	// 	- For PostgreSQL, the version numbers are as follows:
	//
	//     	- **11**
	//
	//     	- **14**
	//
	//     	- **15**
	//
	//     <props="china">
	//
	//
	//
	//       > When creating a Serverless cluster in PolarDB PostgreSQL, only version **14*	- is supported.
	//
	//
	//
	//
	//
	// 	- For Oracle, the version numbers are as follows:
	//
	//     	- **11**
	//
	//     	- **14**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Cluster timezone (UTC), with selectable values ranging from **-12:00*	- to **+13:00*	- at whole-hour intervals, e.g., **00:00**. The default value is **SYSTEM**, which matches the Region\\"s timezone.
	//
	// > This parameter applies only when **DBType*	- is **MySQL**.
	//
	// example:
	//
	// SYSTEM
	DefaultTimeZone *string `json:"DefaultTimeZone,omitempty" xml:"DefaultTimeZone,omitempty"`
	// example:
	//
	// vn-hanoi-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Global Database Network (GDN) ID.
	//
	// > This parameter is required when **CreationOption*	- is **CreateGdnStandby**.
	//
	// example:
	//
	// gdn-***********
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// Specifies whether to enable the hot standby storage cluster feature. Valid values:
	//
	// 	- **ON*	- (default): enables the hot standby storage cluster feature.
	//
	// 	- **OFF**: disables the hot standby storage cluster feature.
	//
	// 	- **STANDBY**: enables the hot standby storage cluster feature for Standard Edition clusters.
	//
	// >  The default value for Standard Edition clusters is **STANDBY**.
	//
	// example:
	//
	// ON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// Enable Binlog feature, valid values are as follows:
	//
	// - **ON**: Cluster enables the Binlog feature. - **OFF**: Cluster disables the Binlog feature. > This parameter takes effect only when the **DBType*	- parameter is set to **MySQL**.
	//
	// example:
	//
	// ON
	LoosePolarLogBin *string `json:"LoosePolarLogBin,omitempty" xml:"LoosePolarLogBin,omitempty"`
	// Enable the X-Engine storage engine feature, with valid values as follows:
	//
	// - **ON**: The cluster enables the X-Engine engine.
	//
	// - **OFF**: The cluster disables the X-Engine engine.
	//
	// > This parameter is effective only when **CreationOption*	- is not **CreateGdnStandby**, **DBType*	- is **MySQL**, and **DBVersion*	- is **8.0**. The memory specification of nodes that enable the X-Engine engine must be at least 8 GB.
	//
	// example:
	//
	// ON
	LooseXEngine *string `json:"LooseXEngine,omitempty" xml:"LooseXEngine,omitempty"`
	// Set the ratio for enabling the X-Engine storage engine, with a range of integers from 10 to 90.
	//
	// > This parameter takes effect only when **LooseXEngine*	- is **ON**.
	//
	// example:
	//
	// 50
	LooseXEngineUseMemoryPct *string `json:"LooseXEngineUseMemoryPct,omitempty" xml:"LooseXEngineUseMemoryPct,omitempty"`
	// Whether table names are case-sensitive, with valid values as follows:
	//
	// 	- **1**: Case-insensitive
	//
	// 	- **0**: Case-sensitive
	//
	// The default value is **1**.
	//
	// > This parameter applies only when **DBType*	- is **MySQL**.
	//
	// example:
	//
	// 1
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitempty" xml:"LowerCaseTableNames,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Parameter template ID.
	//
	// > You can view the list of parameter templates in the target region, including the parameter template ID, by calling the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) interface.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// Payment type, with available values as follows:
	//
	// - **Postpaid**: Pay-as-you-go.
	//
	// - **Prepaid**: Subscription (monthly or yearly).
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// If the payment type is **Prepaid**, this parameter is required. It specifies whether the prepaid cluster is on a monthly or yearly basis.
	//
	// - **Year**: Yearly subscription.
	//
	// - **Month**: Monthly subscription.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// <p id="p_wyg_t4a_glm">The provisioned read and write IOPS for ESSD AutoPL cloud disks. Possible values: 0 to min{50,000, 1000*capacity-Baseline Performance}.</p>
	//
	// <p id="p_6de_jxy_k2g">Baseline Performance = min{1,800+50*capacity, 50000}.</p>
	//
	// <note id="note_7kj_j0o_rgs">This parameter is supported only when StorageType is ESSDAUTOPL.</note>
	//
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// Standard edition database proxy specifications. Values are as follows:
	//
	// - **polar.maxscale.g2.medium.c**: 2 cores.
	//
	// - **polar.maxscale.g2.large.c**: 4 cores.
	//
	// - **polar.maxscale.g2.xlarge.c**: 8 cores.
	//
	// - **polar.maxscale.g2.2xlarge.c**: 16 cores.
	//
	// - **polar.maxscale.g2.3xlarge.c**: 24 cores.
	//
	// - **polar.maxscale.g2.4xlarge.c**: 32 cores.
	//
	// - **polar.maxscale.g2.8xlarge.c**: 64 cores.
	//
	// example:
	//
	// polar.maxscale.g2.medium.c
	ProxyClass *string `json:"ProxyClass,omitempty" xml:"ProxyClass,omitempty"`
	// Database proxy type, with values including:
	//
	// - **EXCLUSIVE**: Enterprise Exclusive Edition
	//
	// - **GENERAL**: Enterprise General Purpose Edition
	//
	// > The proxy type must match the type of the cluster\\"s node specifications, i.e.,
	//
	// >- If the node specification is general, the proxy type should be Enterprise General Purpose Edition;
	//
	// >- If the node specification is dedicated, the proxy type should be Enterprise Exclusive Edition.
	//
	// example:
	//
	// Exclusive
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// Region ID.
	//
	// > You can view available regions through the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Maximum scaling limit for a single node. The value range is: 1 PCU~32 PCU.
	//
	// > Only supported by Serverless clusters.
	//
	// example:
	//
	// 3
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// Minimum scaling limit for a single node. The value range is: 1 PCU~31 PCU.
	//
	// > Only supported by Serverless clusters.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// Maximum scaling limit for the number of read-only nodes. The value range is: 0~15.
	//
	// > Only supported by Serverless clusters.
	//
	// example:
	//
	// 4
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// Minimum scaling limit for the number of read-only nodes. The value range is: 0~15.
	//
	// > Only supported by Serverless clusters.
	//
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// PolarDB cluster whitelist IP address.
	//
	// > Supports configuring multiple whitelist IP addresses, with English commas separating multiple IP addresses.
	//
	// example:
	//
	// 10.***.***.***
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// Serverless type. The current value is fixed to **AgileServerless*	- (sensitive mode).
	//
	// > This parameter is only supported by Serverless clusters.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Source RDS instance ID or source PolarDB cluster ID. This parameter is mandatory only when **CreationOption*	- is set to **MigrationFromRDS**, **CloneFromRDS**, **CloneFromPolarDB**, or **RecoverFromRecyclebin**.
	//
	// 	- If **CreationOption*	- is **MigrationFromRDS*	- or **CloneFromRDS**, you need to input the source RDS instance ID. The source RDS instance version must be RDS MySQL 5.6, 5.7, or 8.0 High Availability edition.
	//
	// 	- If **CreationOption*	- is **CloneFromPolarDB**, you need to input the source PolarDB cluster ID. The DBType of the cloned cluster will default to match the source cluster. For example, if the source cluster is MySQL 8.0, the cloned cluster must also have **DBType*	- set to **MySQL*	- and **DBVersion*	- to **8.0**.
	//
	// 	- If **CreationOption*	- is **RecoverFromRecyclebin**, you need to input the released source PolarDB cluster ID. The DBType of the cluster being recovered from the recycle bin must match the source cluster. For example, if the source cluster was MySQL 8.0, the recovered cluster must also have **DBType*	- set to **MySQL*	- and **DBVersion*	- to **8.0**.
	//
	// example:
	//
	// rm-*************
	SourceResourceId *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	// example:
	//
	// 1022xxxxxxxx
	SourceUid *int64 `json:"SourceUid,omitempty" xml:"SourceUid,omitempty"`
	// The availability zone where the hot standby cluster is stored. Applicable to the standard edition 3AZ scenario.
	//
	// > This parameter takes effect only when multi-zone data strong consistency is enabled.
	//
	// example:
	//
	// cn-hangzhou-g
	StandbyAZ *string `json:"StandbyAZ,omitempty" xml:"StandbyAZ,omitempty"`
	// Whether to enable automatic storage expansion for standard edition clusters, with valid values as follows:
	//
	// - Enable: Enables automatic storage expansion.
	//
	// - Disable: Disables automatic storage expansion.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// Specifies whether to enable disk encryption. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter takes effect only when **DBType*	- is set to **MySQL**.
	//
	// >  This parameter takes effect only when **StorageType*	- is set to one of the Standard Edition storage types.
	StorageEncryption *bool `json:"StorageEncryption,omitempty" xml:"StorageEncryption,omitempty"`
	// The ID of the custom key that is used for disk encryption in the region in which the instance resides. If this parameter is specified, disk encryption is automatically enabled and cannot be disabled afterwards. If you want to use the default service key for disk encryption, leave this parameter empty.
	//
	// You can obtain the ID of the key in the KMS console or create a key.
	//
	// >  This parameter takes effect only when **DBType*	- is set to **MySQL**.
	//
	// >  This parameter takes effect only when **StorageType*	- is set to one of the Standard Edition storage types.
	StorageEncryptionKey *string `json:"StorageEncryptionKey,omitempty" xml:"StorageEncryptionKey,omitempty"`
	// The storage billing type, with valid values as follows:
	//
	// - Postpaid: Pay-as-you-go (hourly).
	//
	// - Prepaid: Pay-per-use based on space (subscription).
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage that is billed based on the subscription billing method. Unit: GB.
	//
	// >
	//
	// 	- Valid values for the subscription storage capacity of a PolarDB for MySQL Standard Edition cluster: 20 to 32000.
	//
	// 	- Valid values for the subscription storage capacity of a Standard Edition cluster that uses the ESSD AUTOPL storage type: 40 to 64000, in increments of 10.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// Enterprise edition storage types include:
	//
	// - **PSL5**
	//
	// - **PSL4**
	//
	// Standard edition storage types include:
	//
	// - **ESSDPL0**
	//
	// - **ESSDPL1**
	//
	// - **ESSDPL2**
	//
	// - **ESSDPL3**
	//
	// - **ESSDAUTOPL**
	//
	// example:
	//
	// PSL4
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Set the upper limit for automatic storage expansion of standard edition clusters, in GB.
	//
	// > The maximum value is 32000.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// Whether the cluster has enabled strong data consistency across multiple zones. Values are as follows:
	//
	// - **ON**: Indicates strong data consistency across multiple zones is enabled, applicable to the standard edition 3AZ scenario.
	//
	// - **OFF**: Indicates strong data consistency across multiple zones is not enabled.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// Enables TDE encryption. Valid values are as follows:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled (default).
	//
	// > 	- This parameter takes effect only when **DBType*	- is **PostgreSQL*	- or **Oracle**.
	//
	// > 	- You can call the [ModifyDBClusterTDE](https://help.aliyun.com/document_detail/167982.html) interface to enable TDE encryption for a PolarDB MySQL cluster.
	//
	// > 	- Once the TDE feature is enabled, it cannot be disabled.
	//
	// example:
	//
	// true
	TDEStatus *bool `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	// List of tags.
	Tag                []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TargetMinorVersion *string                      `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// If the payment type is **Prepaid**, this parameter is required.
	//
	// - When **Period*	- is **Month**, **UsedTime*	- should be an integer within `[1-9]`.
	//
	// - When **Period*	- is **Year**, **UsedTime*	- should be an integer within `[1-3]`.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// Virtual switch ID.
	//
	// > If VPCId has been selected, VSwitchId is mandatory.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Availability Zone ID.
	//
	// > You can view the available zones through the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) interface.
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

func (s *CreateDBClusterRequest) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *CreateDBClusterRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateDBClusterRequest) GetAutoRenew() *bool {
	return s.AutoRenew
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
	// Tag key. If you need to add multiple tags to the target cluster at once, click **Add*	- to add a tag key.
	//
	// > Up to 20 pairs of tags can be added each time, where `Tag.N.Key` corresponds to `Tag.N.Value`.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value. If you need to add multiple tags to the target cluster at once, click **Add*	- to add tag values.
	//
	// > Up to 20 pairs of tags can be added each time, where `Tag.N.Value` corresponds to `Tag.N.Key`.
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

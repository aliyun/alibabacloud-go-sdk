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
	// Specifies whether to enable pause on inactivity. Valid values:
	//
	// - **true**: enables pause on inactivity.
	//
	// - **false*	- (default): disables pause on inactivity.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// The CPU architecture. Valid values:
	//
	// - X86
	//
	// - ARM
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: disables auto-renewal.
	//
	// Default value: **false**.
	//
	// > This parameter takes effect only when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// - true (default): Automatically uses a coupon.
	//
	// - false: does not use a coupon.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The backup retention policy to apply when the cluster is deleted. Valid values:
	//
	// - **ALL**: retains all backup sets.
	//
	// - **LATEST**: retains only the last backup set. An automatic backup is performed before the cluster is deleted.
	//
	// - **NONE**: does not retain backup sets.
	//
	// Default value: **NONE**.
	//
	// > - This parameter is valid only if **DBType*	- is set to **MySQL**.
	//
	// >
	//
	// > - Serverless clusters do not support this parameter.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// Specifies whether to enable the performance burst feature for the ESSD AutoPL cloud disk. Valid values:
	//
	// - **true**: enables the performance burst feature.
	//
	// - **false*	- (default): disables the performance burst feature.
	//
	// > This parameter is supported only when **StorageType*	- is set to ESSDAUTOPL.
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// A client-generated token that ensures the idempotence of the request. This token must be unique across all requests and is case-sensitive. It can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The point in time for the clone. Valid values:
	//
	// - **LATEST**: The latest point in time.
	//
	// - **BackupID**: The ID of a historical backup set.
	//
	// - **Timestamp**: A specific point in time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// Default value: **LATEST**.
	//
	// > If you set **CreationOption*	- to **CloneFromRDS**, you can set this parameter only to **LATEST**.
	//
	// example:
	//
	// LATEST
	CloneDataPoint *string `json:"CloneDataPoint,omitempty" xml:"CloneDataPoint,omitempty"`
	// The cloud service provider of the instance.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The network type of the cluster. Only **VPC*	- is supported.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// - **Normal**: Cluster Edition (default)
	//
	// - **Basic**: Single-node Edition
	//
	// - **ArchiveNormal**: X-Engine Edition
	//
	// - **NormalMultimaster**: Multi-master Cluster Edition
	//
	// - **SENormal**: Standard Edition
	//
	// > 	- The **Basic*	- edition is supported for PolarDB for MySQL **5.6**, **5.7**, and **8.0**; PolarDB for PostgreSQL **14**; and PolarDB for PostgreSQL (compatible with Oracle) **2.0**.
	//
	// >
	//
	// > 	- The **ArchiveNormal*	- and **NormalMultimaster*	- editions are supported for PolarDB for MySQL **8.0**.
	//
	// >
	//
	// > 	- The **SENormal*	- edition is supported for PolarDB for MySQL **5.6**, **5.7**, and **8.0*	- and PolarDB for PostgreSQL **14**.
	//
	// For more information about product editions, see [Editions](https://help.aliyun.com/document_detail/183258.html).
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The method to create the cluster. Valid values:
	//
	// - **Normal**: Creates a new PolarDB cluster. For more information, see the following topics:
	//
	//   - [Create a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/58769.html)
	//
	//   - [Create a PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/118063.html)
	//
	//   - [Create a PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/118182.html)
	//
	// - **CloneFromPolarDB**: Clones data from an existing PolarDB cluster. For more information, see the following topics:
	//
	//   - [Clone a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/87966.html)
	//
	//   - [Clone a PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/118108.html)
	//
	//   - [Clone a PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/118221.html)
	//
	// - **RecoverFromRecyclebin**: Restores a PolarDB cluster from the recycle bin. For more information, see the following topics:
	//
	//   - [Restore a released PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/164880.html)
	//
	//   - [Restore a released PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/432844.html)
	//
	//   - [Restore a released PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/424632.html)
	//
	// - **CloneFromRDS**: Clones data from an existing ApsaraDB RDS instance to a new PolarDB cluster. For more information, see [One-click cloning from ApsaraDB RDS for MySQL to PolarDB for MySQL](https://help.aliyun.com/document_detail/121812.html).
	//
	// - **MigrationFromRDS**: Migrates data from an existing ApsaraDB RDS instance. The created PolarDB cluster is in read-only mode and has binary logging enabled by default. For more information, see [One-click upgrade from ApsaraDB RDS for MySQL to PolarDB for MySQL](https://help.aliyun.com/document_detail/121582.html).
	//
	// - **CreateGdnStandby**: Creates a secondary cluster in a Global Database Network (GDN). For more information, see [Add a secondary cluster](https://help.aliyun.com/document_detail/160381.html).
	//
	// - **UpgradeFromPolarDB**: Upgrades the major version of a PolarDB cluster. For more information, see [Perform a major version upgrade](https://help.aliyun.com/document_detail/459712.html).
	//
	// Default value: **Normal**.
	//
	// > If **DBType*	- is set to **MySQL*	- and **DBVersion*	- is set to **8.0**, you can set this parameter to **CreateGdnStandby**.
	//
	// example:
	//
	// Normal
	CreationOption *string `json:"CreationOption,omitempty" xml:"CreationOption,omitempty"`
	// The description of the cluster. The description must meet the following requirements:
	//
	// - It cannot start with `http://` or `https://`.
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The minor version of the database engine. Valid values:
	//
	// - **8.0.2**
	//
	// - **8.0.1**
	//
	// > This parameter is valid only if **DBType*	- is set to **MySQL*	- and **DBVersion*	- is set to **8.0**.
	//
	// example:
	//
	// 8.0.1
	DBMinorVersion *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	// The node specification. For more information, see the following topics:
	//
	// - PolarDB for MySQL: [Compute node specifications](https://help.aliyun.com/document_detail/102542.html)
	//
	// - PolarDB for PostgreSQL (compatible with Oracle): [Compute node specifications](https://help.aliyun.com/document_detail/207921.html)
	//
	// - PolarDB for PostgreSQL: [Compute node specifications](https://help.aliyun.com/document_detail/209380.html)
	//
	// > 	- To create a PolarDB for MySQL Cluster Edition serverless cluster, set this parameter to **polar.mysql.sl.small**.
	//
	// >
	//
	// > 	- To create a PolarDB for MySQL Standard Edition serverless cluster, set this parameter to **polar.mysql.sl.small.c**.
	//
	// >
	//
	// > 	- To create a PolarDB for PostgreSQL Cluster Edition serverless cluster, set this parameter to **polar.pg.sl.small**.
	//
	// >
	//
	// > 	- To create a PolarDB for PostgreSQL Standard Edition serverless cluster, set this parameter to **polar.pg.sl.small.c**.
	//
	// >
	//
	// > 	- To create a PolarDB for PostgreSQL (compatible with Oracle) serverless cluster, set this parameter to **polar.o.sl.small**.
	//
	// example:
	//
	// polar.mysql.x4.medium
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes for a Standard Edition or Enterprise Edition cluster. Valid values:
	//
	// - Standard Edition: 1 to 8. A cluster of this edition includes one read/write node and up to seven read-only nodes.
	//
	// - Enterprise Edition: 1 to 16. A cluster of this edition includes one read/write node and up to 15 read-only nodes.
	//
	// > 	- By default, an Enterprise Edition cluster has two nodes and a Standard Edition cluster has one node.
	//
	// >
	//
	// > 	- This parameter is supported only for PolarDB for MySQL.
	//
	// >
	//
	// > 	- You cannot change the number of nodes in a Multi-master Cluster Edition cluster.
	//
	// example:
	//
	// 1
	DBNodeNum *int32 `json:"DBNodeNum,omitempty" xml:"DBNodeNum,omitempty"`
	// The database engine. Valid values:
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
	// The version of the database engine.
	//
	// - Valid values for MySQL:
	//
	//   - **5.6**
	//
	//   - **5.7**
	//
	//   - **8.0**
	//
	// - Valid values for PostgreSQL:
	//
	//   - **11**
	//
	//   - **14**
	//
	//   - **15**<props="china">
	//
	// > If you create a serverless cluster for PolarDB for PostgreSQL, you must set this parameter to `14`.
	//
	// \\	- Valid values for Oracle:
	//
	// \\	- **11**
	//
	// \\	- **14**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Cluster time zone (UTC). The value can be any full-hour offset from **-12:00 to +13:00**, such as **00:00**. The default value is **SYSTEM**, which uses the region\\"s time zone.
	//
	// > This parameter takes effect only when **DBType*	- is **MySQL**.
	//
	// example:
	//
	// SYSTEM
	DefaultTimeZone *string `json:"DefaultTimeZone,omitempty" xml:"DefaultTimeZone,omitempty"`
	// The ID of the Edge Node Service (ENS) node. This parameter is required if you want to create an ENS database instance.
	//
	// example:
	//
	// vn-hanoi-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the Global Database Network (GDN).
	//
	// > This parameter is required if **CreationOption*	- is set to **CreateGdnStandby**.
	//
	// example:
	//
	// gdn-***********
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// Specifies whether to enable the hot standby cluster feature. Valid values:
	//
	// - **ON*	- (default): enables a hot standby storage cluster.
	//
	// - **OFF**: disables the hot standby cluster feature.
	//
	// - **STANDBY**: enables a hot standby cluster.
	//
	// - **EQUAL**: enables hot standby for both storage and computing resources.
	//
	// - **3AZ**: enables multi-AZ strong consistency.
	//
	// > The value **STANDBY*	- is valid only for PolarDB for PostgreSQL.
	//
	// example:
	//
	// ON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// Specifies whether to enable binary logging. Valid values:
	//
	// - **ON**: enables binary logging.
	//
	// - **OFF**: disables binary logging.
	//
	// > This parameter is valid only if **DBType*	- is set to **MySQL**.
	//
	// example:
	//
	// ON
	LoosePolarLogBin *string `json:"LoosePolarLogBin,omitempty" xml:"LoosePolarLogBin,omitempty"`
	// Specifies whether to enable the X-Engine storage engine. Valid values:
	//
	// - **ON**: enables the X-Engine storage engine.
	//
	// - **OFF**: disables the X-Engine storage engine.
	//
	// > This parameter is valid only if the **CreationOption*	- parameter is not set to **CreateGdnStandby**, **DBType*	- is set to **MySQL**, and **DBVersion*	- is set to **8.0**. To enable the X-Engine storage engine, the node must have at least 8 GB of memory.
	//
	// example:
	//
	// ON
	LooseXEngine *string `json:"LooseXEngine,omitempty" xml:"LooseXEngine,omitempty"`
	// The percentage of memory allocated to the X-Engine storage engine. Valid values: integers from 10 to 90.
	//
	// > This parameter is valid only if **LooseXEngine*	- is set to **ON**.
	//
	// example:
	//
	// 50
	LooseXEngineUseMemoryPct *string `json:"LooseXEngineUseMemoryPct,omitempty" xml:"LooseXEngineUseMemoryPct,omitempty"`
	// The time zone of the cluster. The value must be a UTC offset in the `±HH:mm` format. Valid values: from **-12:00*	- to **+13:00*	- on the hour. For example, **00:00**. The default value **SYSTEM*	- indicates that the cluster uses the time zone of its region.
	//
	// - **1**: Case-insensitive
	//
	// - **0**: Case-sensitive
	//
	// The default value is **1**.
	//
	// > This parameter is valid only if **DBType*	- is set to **MySQL**.
	//
	// example:
	//
	// 1
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitempty" xml:"LowerCaseTableNames,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// > You can call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to query the parameter templates in a specific region, including the IDs of the parameter templates.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. This parameter is required if you set the **PayType*	- parameter to **Prepaid**. Valid values:
	//
	// - **Year**: The subscription duration is measured in years.
	//
	// - **Month**: The subscription duration is measured in months.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The promotion code. If you do not specify this parameter, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// <props="china">
	//
	// The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	//
	//
	// <props="china">
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	//
	//
	// <props="china">
	//
	// > This parameter is supported only when **StorageType*	- is set to ESSDAUTOPL.
	//
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The specification of the database proxy for a Standard Edition cluster. Valid values:
	//
	// - **polar.maxscale.g2.medium.c**: 2 cores
	//
	// - **polar.maxscale.g2.large.c**: 4 cores
	//
	// - **polar.maxscale.g2.xlarge.c**: 8 cores
	//
	// - **polar.maxscale.g2.2xlarge.c**: 16 cores
	//
	// - **polar.maxscale.g2.3xlarge.c**: 24 cores
	//
	// - **polar.maxscale.g2.4xlarge.c**: 32 cores
	//
	// - **polar.maxscale.g2.8xlarge.c**: 64 cores
	//
	// example:
	//
	// polar.maxscale.g2.medium.c
	ProxyClass *string `json:"ProxyClass,omitempty" xml:"ProxyClass,omitempty"`
	// The type of the database proxy. Valid values:
	//
	// - **EXCLUSIVE**: Enterprise Dedicated
	//
	// - **GENERAL**: Enterprise General-purpose
	//
	// > The proxy type must be consistent with the type that corresponds to the node specification of the cluster:
	//
	// >
	//
	// > - If the node specification is general-purpose, the proxy type must be Enterprise General-purpose.
	//
	// >
	//
	// > - If the node specification is dedicated, the proxy type must be Enterprise Dedicated.
	//
	// example:
	//
	// Exclusive
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of PCUs for a single-node serverless cluster to scale up to. Valid values: 1 to 32.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// 3
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of PolarDB compute units (PCUs) for a single-node serverless cluster to scale down to. Valid values: 1 to 31.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The maximum number of read-only nodes that the serverless cluster scales up to. Valid values: 0 to 15.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// 4
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// The minimum number of read-only nodes that the serverless cluster scales down to. Valid values: 0 to 15.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// The IP whitelist of the PolarDB cluster.
	//
	// > You can specify multiple IP addresses in the IP whitelist. Separate the IP addresses with commas (,).
	//
	// example:
	//
	// 10.***.***.***
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The type of the serverless cluster. Set the value to **AgileServerless**.
	//
	// > This parameter is supported only for serverless clusters.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The ID of the source ApsaraDB RDS instance or source PolarDB cluster. This parameter is required only if **CreationOption*	- is set to **MigrationFromRDS**, **CloneFromRDS**, **CloneFromPolarDB**, or **RecoverFromRecyclebin**.
	//
	// - If **CreationOption*	- is set to **MigrationFromRDS*	- or **CloneFromRDS**, specify the ID of the source ApsaraDB RDS instance. The source ApsaraDB RDS instance must be ApsaraDB RDS for MySQL 5.6, 5.7, or 8.0 High-availability Edition.
	//
	// - If **CreationOption*	- is set to **CloneFromPolarDB**, specify the ID of the source PolarDB cluster. The new cluster must use the same database engine as the source cluster. For example, if the source cluster runs MySQL 8.0, you must set **DBType*	- to **MySQL*	- and **DBVersion*	- to **8.0*	- for the new cluster.
	//
	// - If **CreationOption*	- is set to **RecoverFromRecyclebin**, specify the ID of the released source PolarDB cluster. The restored cluster must use the same database engine as the source cluster. For example, if the source cluster runs MySQL 8.0, you must set **DBType*	- to **MySQL*	- and **DBVersion*	- to **8.0*	- for the restored cluster.
	//
	// example:
	//
	// rm-*************
	SourceResourceId *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	// The UID of the source backup set owner in cross-account backup and restoration scenarios.
	//
	// example:
	//
	// 1022xxxxxxxx
	SourceUid *int64 `json:"SourceUid,omitempty" xml:"SourceUid,omitempty"`
	// The zone for the hot standby cluster.
	//
	// > This parameter is valid only when the hot standby cluster feature or multi-AZ strong consistency is enabled.
	//
	// example:
	//
	// cn-hangzhou-g
	StandbyAZ *string `json:"StandbyAZ,omitempty" xml:"StandbyAZ,omitempty"`
	// Specifies whether to enable automatic storage scaling for a Standard Edition cluster. Valid values:
	//
	// - Enable: enables automatic storage scaling.
	//
	// - Disable: disables automatic storage scaling.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// Specifies whether to enable cloud disk encryption. Valid values:
	//
	// - **true**: enables cloud disk encryption.
	//
	// - **false*	- (default): disables cloud disk encryption.
	//
	// > This parameter is valid only if **DBType*	- is set to **MySQL**.
	//
	// > This parameter is valid only if **StorageType*	- is set to a Standard Edition storage type.
	StorageEncryption *bool `json:"StorageEncryption,omitempty" xml:"StorageEncryption,omitempty"`
	// The ID of a custom key from Key Management Service (KMS) for cloud disk encryption. The key must be in the same region as the cluster. If you specify this parameter, cloud disk encryption is automatically enabled and cannot be disabled. If this parameter is empty, the default service key is used.
	//
	// You can view the key ID or create a new key in the Key Management Service (KMS) console.
	//
	// > This parameter is valid only if **DBType*	- is set to **MySQL**.
	//
	// > This parameter is valid only if **StorageType*	- is set to a Standard Edition storage type.
	//
	// example:
	//
	// 1022xxxxxxxx
	StorageEncryptionKey *string `json:"StorageEncryptionKey,omitempty" xml:"StorageEncryptionKey,omitempty"`
	// The billing method for storage. Valid values:
	//
	// - Postpaid: pay-by-capacity (a pay-as-you-go method).
	//
	// - Prepaid: pay-by-space (a subscription method).
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage space for a pay-by-space (subscription) cluster. Unit: GB.
	//
	// > - Valid values for a PolarDB for MySQL Enterprise Edition cluster: 10 to 50000.
	//
	// >
	//
	// > - Valid values for a PolarDB for MySQL Standard Edition cluster: 20 to 64000.
	//
	// >
	//
	// > - If the storage type of a Standard Edition cluster is ESSD AutoPL, the storage space must be a multiple of 10 between 40 and 64000.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// Valid values for Enterprise Edition:
	//
	// - **PSL5**
	//
	// - **PSL4**
	//
	// Valid values for Standard Edition:
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
	// The maximum storage capacity for a Standard Edition cluster when automatic storage scaling is enabled. Unit: GB.
	//
	// > The maximum value is 32000.
	//
	// example:
	//
	// 800
	StorageUpperBound *int64 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// Specifies whether to enable multi-AZ strong consistency for the cluster. Valid values:
	//
	// - **ON**: enables multi-AZ strong consistency. This feature is applicable to Standard Edition clusters that are deployed across three zones.
	//
	// - **OFF**: disables multi-AZ strong consistency.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// Specifies whether to enable transparent data encryption (TDE). Valid values:
	//
	// - **true**: enables TDE.
	//
	// - **false*	- (default): disables TDE.
	//
	// > 	- This parameter is valid only when **DBType*	- is set to **PostgreSQL*	- or **Oracle**.
	//
	// >
	//
	// > 	- You can call the [ModifyDBClusterTDE](https://help.aliyun.com/document_detail/167982.html) operation to enable TDE for a PolarDB for MySQL cluster.
	//
	// >
	//
	// > 	- TDE cannot be disabled after it is enabled.
	//
	// example:
	//
	// true
	TDEStatus *bool `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	// The tags to add to the cluster.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The target minor engine version.
	//
	// example:
	//
	// 8.0.1.1.54
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The subscription duration. This parameter is required if you set the **PayType*	- parameter to **Prepaid**.
	//
	// - If **Period*	- is set to **Month**, **UsedTime*	- must be an integer from `[1-9]`.
	//
	// - If **Period*	- is set to **Year**, **UsedTime*	- must be an integer from `[1-3]`.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the VSwitch.
	//
	// > If you specify the VPCId parameter, you must also specify this parameter.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available zones.
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
	// The key of the tag.
	//
	// > You can add up to 20 tags at a time. The Nth tag is a key-value pair, where `Tag.N.Key` is the key and `Tag.N.Value` is the value.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > You can add up to 20 tags at a time. The Nth tag is a key-value pair, where `Tag.N.Key` is the key and `Tag.N.Value` is the value.
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

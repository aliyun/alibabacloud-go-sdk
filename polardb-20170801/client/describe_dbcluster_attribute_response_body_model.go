// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiCreatingTime(v string) *DescribeDBClusterAttributeResponseBody
	GetAiCreatingTime() *string
	SetAiType(v string) *DescribeDBClusterAttributeResponseBody
	GetAiType() *string
	SetArchitecture(v string) *DescribeDBClusterAttributeResponseBody
	GetArchitecture() *string
	SetAutoUpgradeMinorVersion(v string) *DescribeDBClusterAttributeResponseBody
	GetAutoUpgradeMinorVersion() *string
	SetBlktagTotal(v int64) *DescribeDBClusterAttributeResponseBody
	GetBlktagTotal() *int64
	SetBlktagUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetBlktagUsed() *int64
	SetBranch(v *DescribeDBClusterAttributeResponseBodyBranch) *DescribeDBClusterAttributeResponseBody
	GetBranch() *DescribeDBClusterAttributeResponseBodyBranch
	SetBurstingEnabled(v string) *DescribeDBClusterAttributeResponseBody
	GetBurstingEnabled() *string
	SetCategory(v string) *DescribeDBClusterAttributeResponseBody
	GetCategory() *string
	SetColumnTable(v string) *DescribeDBClusterAttributeResponseBody
	GetColumnTable() *string
	SetCompressStorageMode(v string) *DescribeDBClusterAttributeResponseBody
	GetCompressStorageMode() *string
	SetCompressStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetCompressStorageUsed() *int64
	SetCreationTime(v string) *DescribeDBClusterAttributeResponseBody
	GetCreationTime() *string
	SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBody
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBody
	GetDBClusterId() *string
	SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBody
	GetDBClusterNetworkType() *string
	SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBody
	GetDBClusterStatus() *string
	SetDBNodes(v []*DescribeDBClusterAttributeResponseBodyDBNodes) *DescribeDBClusterAttributeResponseBody
	GetDBNodes() []*DescribeDBClusterAttributeResponseBodyDBNodes
	SetDBType(v string) *DescribeDBClusterAttributeResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterAttributeResponseBody
	GetDBVersion() *string
	SetDBVersionStatus(v string) *DescribeDBClusterAttributeResponseBody
	GetDBVersionStatus() *string
	SetDataLevel1BackupChainSize(v int64) *DescribeDBClusterAttributeResponseBody
	GetDataLevel1BackupChainSize() *int64
	SetDataSyncMode(v string) *DescribeDBClusterAttributeResponseBody
	GetDataSyncMode() *string
	SetDeletionLock(v int32) *DescribeDBClusterAttributeResponseBody
	GetDeletionLock() *int32
	SetEngine(v string) *DescribeDBClusterAttributeResponseBody
	GetEngine() *string
	SetExpireTime(v string) *DescribeDBClusterAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v string) *DescribeDBClusterAttributeResponseBody
	GetExpired() *string
	SetHasCompleteStandbyRes(v bool) *DescribeDBClusterAttributeResponseBody
	GetHasCompleteStandbyRes() *bool
	SetHotStandbyCluster(v string) *DescribeDBClusterAttributeResponseBody
	GetHotStandbyCluster() *string
	SetImciAutoIndex(v string) *DescribeDBClusterAttributeResponseBody
	GetImciAutoIndex() *string
	SetImperceptibleSwitch(v string) *DescribeDBClusterAttributeResponseBody
	GetImperceptibleSwitch() *string
	SetInodeTotal(v int64) *DescribeDBClusterAttributeResponseBody
	GetInodeTotal() *int64
	SetInodeUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetInodeUsed() *int64
	SetIsLatestVersion(v bool) *DescribeDBClusterAttributeResponseBody
	GetIsLatestVersion() *bool
	SetIsProxyLatestVersion(v bool) *DescribeDBClusterAttributeResponseBody
	GetIsProxyLatestVersion() *bool
	SetLockMode(v string) *DescribeDBClusterAttributeResponseBody
	GetLockMode() *string
	SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBody
	GetMaintainTime() *string
	SetOrca(v string) *DescribeDBClusterAttributeResponseBody
	GetOrca() *string
	SetPayType(v string) *DescribeDBClusterAttributeResponseBody
	GetPayType() *string
	SetProvisionedIops(v string) *DescribeDBClusterAttributeResponseBody
	GetProvisionedIops() *string
	SetProxyCpuCores(v string) *DescribeDBClusterAttributeResponseBody
	GetProxyCpuCores() *string
	SetProxyServerlessType(v string) *DescribeDBClusterAttributeResponseBody
	GetProxyServerlessType() *string
	SetProxyStandardCpuCores(v string) *DescribeDBClusterAttributeResponseBody
	GetProxyStandardCpuCores() *string
	SetProxyStatus(v string) *DescribeDBClusterAttributeResponseBody
	GetProxyStatus() *string
	SetProxyType(v string) *DescribeDBClusterAttributeResponseBody
	GetProxyType() *string
	SetRegionId(v string) *DescribeDBClusterAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDBClusterAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBody
	GetResourceGroupId() *string
	SetRestoreDataPoint(v string) *DescribeDBClusterAttributeResponseBody
	GetRestoreDataPoint() *string
	SetRestoreType(v string) *DescribeDBClusterAttributeResponseBody
	GetRestoreType() *string
	SetRowCompression(v string) *DescribeDBClusterAttributeResponseBody
	GetRowCompression() *string
	SetSQLSize(v int64) *DescribeDBClusterAttributeResponseBody
	GetSQLSize() *int64
	SetSearchClusterStatus(v string) *DescribeDBClusterAttributeResponseBody
	GetSearchClusterStatus() *string
	SetSearchCompressStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetSearchCompressStorageUsed() *int64
	SetSearchStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetSearchStorageUsed() *int64
	SetServerlessType(v string) *DescribeDBClusterAttributeResponseBody
	GetServerlessType() *string
	SetSourceDBCluster(v string) *DescribeDBClusterAttributeResponseBody
	GetSourceDBCluster() *string
	SetSourceRegionId(v string) *DescribeDBClusterAttributeResponseBody
	GetSourceRegionId() *string
	SetStandbyHAMode(v string) *DescribeDBClusterAttributeResponseBody
	GetStandbyHAMode() *string
	SetStorageMax(v int64) *DescribeDBClusterAttributeResponseBody
	GetStorageMax() *int64
	SetStoragePayType(v string) *DescribeDBClusterAttributeResponseBody
	GetStoragePayType() *string
	SetStorageSpace(v int64) *DescribeDBClusterAttributeResponseBody
	GetStorageSpace() *int64
	SetStorageType(v string) *DescribeDBClusterAttributeResponseBody
	GetStorageType() *string
	SetStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetStorageUsed() *int64
	SetStrictConsistency(v string) *DescribeDBClusterAttributeResponseBody
	GetStrictConsistency() *string
	SetSubCategory(v string) *DescribeDBClusterAttributeResponseBody
	GetSubCategory() *string
	SetSupportInstantSwitchWithImci(v string) *DescribeDBClusterAttributeResponseBody
	GetSupportInstantSwitchWithImci() *string
	SetTags(v []*DescribeDBClusterAttributeResponseBodyTags) *DescribeDBClusterAttributeResponseBody
	GetTags() []*DescribeDBClusterAttributeResponseBodyTags
	SetVPCId(v string) *DescribeDBClusterAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBody
	GetVSwitchId() *string
	SetZoneIds(v string) *DescribeDBClusterAttributeResponseBody
	GetZoneIds() *string
}

type DescribeDBClusterAttributeResponseBody struct {
	// The start time of the free trial for the AI feature.
	//
	// example:
	//
	// 2024-03-13T01:20:28Z
	AiCreatingTime *string `json:"AiCreatingTime,omitempty" xml:"AiCreatingTime,omitempty"`
	// The AI node type. Valid values:
	//
	// - **SearchNode**: a search node.
	//
	// - **DLNode**: an AI node.
	//
	// example:
	//
	// DLNode
	AiType *string `json:"AiType,omitempty" xml:"AiType,omitempty"`
	// The CPU architecture. Valid values:
	//
	// - **X86**
	//
	// - **ARM**
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The method for minor version upgrades.
	//
	// - Auto: automatic upgrade
	//
	// - Manual: manual upgrade
	//
	// example:
	//
	// Manual
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// The maximum number of blktags for the file system.
	//
	// example:
	//
	// 7,864,320
	BlktagTotal *int64 `json:"BlktagTotal,omitempty" xml:"BlktagTotal,omitempty"`
	// The number of used blktags.
	//
	// example:
	//
	// 5,242,880
	BlktagUsed *int64                                        `json:"BlktagUsed,omitempty" xml:"BlktagUsed,omitempty"`
	Branch     *DescribeDBClusterAttributeResponseBodyBranch `json:"Branch,omitempty" xml:"Branch,omitempty" type:"Struct"`
	// Indicates whether the performance burst feature is enabled for the ESSD AutoPL disk. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [product edition](https://help.aliyun.com/document_detail/183258.html). Valid values:
	//
	// - **Normal**: Cluster Edition
	//
	// - **Basic**: Single Node Edition
	//
	// - **Archive**: X-Engine
	//
	// - **NormalMultimaster**: Multi-master Cluster Edition
	//
	// - **SENormal**: Standard Edition
	//
	// > 	- The single node edition is not supported for PolarDB for PostgreSQL that runs PostgreSQL 11.
	//
	// >
	//
	// > 	- The Standard Edition is supported on PolarDB for MySQL that runs MySQL 8.0 or 5.7 and on PolarDB for PostgreSQL that runs PostgreSQL 14.
	//
	// >
	//
	// > 	- PolarDB for MySQL that runs MySQL 8.0 supports X-Engine and the Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether column-oriented tables are enabled.
	//
	// example:
	//
	// OFF
	ColumnTable *string `json:"ColumnTable,omitempty" xml:"ColumnTable,omitempty"`
	// Indicates whether storage compression is enabled. Valid values:
	//
	// - ON: enabled
	//
	// - OFF: disabled
	//
	// example:
	//
	// ON
	CompressStorageMode *string `json:"CompressStorageMode,omitempty" xml:"CompressStorageMode,omitempty"`
	// The size of the compressed storage data.
	//
	// > This parameter is returned only when the storage compression feature is enabled for the cluster.
	//
	// example:
	//
	// 15529410560
	CompressStorageUsed *int64 `json:"CompressStorageUsed,omitempty" xml:"CompressStorageUsed,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The cluster status. For more information about the valid values, see [Cluster states](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The details of the nodes.
	DBNodes []*DescribeDBClusterAttributeResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The status of the minor engine version. Valid values:
	//
	// - **Stable**: The current version is stable.
	//
	// - **Old**: The current version is old. Upgrade the cluster to the latest version.
	//
	// - **HighRisk**: The current version has critical defects. Immediately upgrade the cluster to the latest version.
	//
	// - **Beta**: The current version is a beta version.
	//
	// > 	- For more information about how to upgrade the minor engine version, see [Upgrade versions](https://help.aliyun.com/document_detail/158572.html).
	//
	// >
	//
	// > 	- This parameter is returned only when the **DBType*	- parameter is set to **MySQL**.
	//
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// The total size of level-1 backups (snapshots). Unit: bytes.
	//
	// example:
	//
	// 74448896
	DataLevel1BackupChainSize *int64 `json:"DataLevel1BackupChainSize,omitempty" xml:"DataLevel1BackupChainSize,omitempty"`
	// The data replication mode. Valid values:
	//
	// - **AsyncSync**: asynchronous
	//
	// - **SemiSync**: semi-synchronous
	//
	// example:
	//
	// AsyncSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// The lock state of the cluster for deletion. Valid values:
	//
	// - **0**: The cluster is not locked and can be deleted.
	//
	// - **1**: The cluster is locked and cannot be deleted.
	//
	// example:
	//
	// 0
	DeletionLock *int32 `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	// The cluster engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The expiration time of the cluster.
	//
	// > This parameter is returned only for subscription clusters. An empty value is returned for pay-as-you-go clusters.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired.
	//
	// > This parameter is returned only for subscription clusters.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates whether resources for the new primary database are provisioned after a cross-zone failover. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	HasCompleteStandbyRes *bool `json:"HasCompleteStandbyRes,omitempty" xml:"HasCompleteStandbyRes,omitempty"`
	// Indicates whether the hot standby storage cluster (and standby compute nodes) is enabled. Valid values:
	//
	// - **StandbyClusterON**: The hot standby storage cluster or both the hot standby storage cluster and standby compute nodes are enabled.
	//
	// - **StandbyClusterOFF**: The hot standby storage cluster or both the hot standby storage cluster and standby compute nodes are disabled.
	//
	// example:
	//
	// StandbyClusterON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// The automatic IMCI-based query acceleration feature. Valid values:
	//
	// - `ON`: enabled.
	//
	// - `OFF`: disabled.
	//
	// example:
	//
	// OFF
	ImciAutoIndex *string `json:"ImciAutoIndex,omitempty" xml:"ImciAutoIndex,omitempty"`
	// The imperceptible switchover feature. Valid values:
	//
	// - `true`: enabled.
	//
	// - `false`: disabled.
	//
	// example:
	//
	// true
	ImperceptibleSwitch *string `json:"ImperceptibleSwitch,omitempty" xml:"ImperceptibleSwitch,omitempty"`
	// The maximum number of inodes for the file system.
	//
	// example:
	//
	// 6,291,456
	InodeTotal *int64 `json:"InodeTotal,omitempty" xml:"InodeTotal,omitempty"`
	// The number of used inodes.
	//
	// example:
	//
	// 4,194,304
	InodeUsed *int64 `json:"InodeUsed,omitempty" xml:"InodeUsed,omitempty"`
	// Indicates whether the kernel is the latest version. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the database proxy is the latest version. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	IsProxyLatestVersion *bool `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// The lock mode. Valid values:
	//
	// - **Unlock**: The cluster is not locked.
	//
	// - **ManualLock**: The cluster is manually locked.
	//
	// - **LockByExpiration**: The cluster is automatically locked after it expires.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance window of the cluster. The time is in the `HH:mmZ-HH:mmZ` format and is in UTC. For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 00:00 to 01:00 (UTC+8).
	//
	// example:
	//
	// 18:00Z-19:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// ON
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// <props="china">
	//
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
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
	// > This parameter is supported only when StorageType is set to ESSDAUTOPL.
	//
	// example:
	//
	// 2500
	ProvisionedIops *string `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of CPU cores of the database proxy.
	//
	// example:
	//
	// 4
	ProxyCpuCores *string `json:"ProxyCpuCores,omitempty" xml:"ProxyCpuCores,omitempty"`
	// The Serverless type of the database proxy. Valid values:
	//
	// - AgileServerless: agile, which indicates a Serverless cluster.
	//
	// - SteadyServerless: steady, which indicates a cluster with defined specifications (a subscription or pay-as-you-go cluster).
	//
	// example:
	//
	// SteadyServerless
	ProxyServerlessType *string `json:"ProxyServerlessType,omitempty" xml:"ProxyServerlessType,omitempty"`
	// The number of CPU cores of the database proxy with standard specifications.
	//
	// example:
	//
	// 2
	ProxyStandardCpuCores *string `json:"ProxyStandardCpuCores,omitempty" xml:"ProxyStandardCpuCores,omitempty"`
	// The status of the database proxy. Valid values:
	//
	// - **Creating**
	//
	// - **Running**
	//
	// - **Deleting**: The proxy is being released.
	//
	// - **Rebooting**
	//
	// - **DBNodeCreating**: A node is being added.
	//
	// - **DBNodeDeleting**: A node is being deleted.
	//
	// - **ClassChanging**: The node specifications are being changed.
	//
	// - **NetAddressCreating**: A network connection is being created.
	//
	// - **NetAddressDeleting**: A network connection is being deleted.
	//
	// - **NetAddressModifying**: A network connection is being modified.
	//
	// - **Deleted**: The proxy is released.
	//
	// example:
	//
	// Running
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The type of the database proxy. Valid values:
	//
	// - **Exclusive**: Dedicated Enterprise Edition
	//
	// - **General*	- : Standard Enterprise Edition
	//
	// example:
	//
	// Exclusive
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 074467EF-86B9-4C23-ACBF-E9B81A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-***************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// - If RestoreType is set to **RestoreByTime*	- or **RestoreByTimeOss**, this parameter indicates the point in time for the restoration.
	//
	// - If RestoreType is set to **RestoreByBackupSet*	- or **RestoreByBackupSetOss**, this parameter indicates the ID of the backup set that is used for the restoration.
	//
	// > This parameter is supported only for clusters that are restored from a backup set or a point in time after June 1, 2024.
	//
	// example:
	//
	// 2179639137
	RestoreDataPoint *string `json:"RestoreDataPoint,omitempty" xml:"RestoreDataPoint,omitempty"`
	// The method that is used to restore the cluster. Valid values:
	//
	// - **RestoreByTime**: The cluster is restored to a point in time from a level-1 backup.
	//
	// - **RestoreByBackupSet**: The cluster is restored from a level-1 backup set.
	//
	// - **RestoreByTimeOss**: The cluster is restored to a point in time from a level-2 backup.
	//
	// - **RestoreByBackupSetOss**: The cluster is restored from a level-2 backup set.
	//
	// - **CloneFromSourceCluster**: The cluster is cloned from a source cluster.
	//
	// > This parameter is supported only for clusters that are restored from a backup set or a point in time after June 1, 2024.
	//
	// example:
	//
	// RestoreByTime
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The row compression settings.
	//
	// example:
	//
	// OFF
	RowCompression *string `json:"RowCompression,omitempty" xml:"RowCompression,omitempty"`
	// The storage usage for SQL statements. Unit: bytes. A value of -1 indicates that no data is available.
	//
	// example:
	//
	// 0
	SQLSize *int64 `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
	// The running state of the search node.
	//
	// example:
	//
	// Running
	SearchClusterStatus *string `json:"SearchClusterStatus,omitempty" xml:"SearchClusterStatus,omitempty"`
	// The size of the compressed storage data of the search node.
	//
	// > This parameter is returned only when the storage compression feature is enabled for the cluster.
	//
	// example:
	//
	// 15529410560
	SearchCompressStorageUsed *int64 `json:"SearchCompressStorageUsed,omitempty" xml:"SearchCompressStorageUsed,omitempty"`
	// The storage usage of the search node.
	//
	// example:
	//
	// 3012558848
	SearchStorageUsed *int64 `json:"SearchStorageUsed,omitempty" xml:"SearchStorageUsed,omitempty"`
	// The Serverless type of the cluster. Valid values:
	//
	// - AgileServerless: agile, which indicates a Serverless cluster.
	//
	// - SteadyServerless: steady, which indicates a cluster with defined specifications for which the Serverless feature is enabled.
	//
	// > This parameter is supported only for Serverless clusters or clusters with defined specifications for which the Serverless feature is enabled.
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The source cluster ID.
	//
	// > This parameter is supported only for clusters that are restored from a backup set or a point in time after June 1, 2024.
	//
	// example:
	//
	// pc-pz51ziv48317b2880
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The region ID of the source cluster.
	//
	// > This parameter is returned only when the source cluster ID exists.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The cross-zone disaster recovery mode. Valid values:
	//
	// - **ON**: The cross-zone disaster recovery mode is enabled.
	//
	// - **OFF**: The cross-zone disaster recovery mode is disabled.
	//
	// - **0**: The customer drill mode.
	//
	// example:
	//
	// OFF
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// The maximum storage capacity of the current cluster specifications. Unit: bytes.
	//
	// example:
	//
	// 10995116277760
	StorageMax *int64 `json:"StorageMax,omitempty" xml:"StorageMax,omitempty"`
	// The billing method for storage. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage space of the subscription cluster. Unit: bytes.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage class. The value is fixed as **HighPerformance**.
	//
	// example:
	//
	// HighPerformance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The used storage space. Unit: bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether strong consistency is enabled for the multi-zone data. Valid values:
	//
	// - **ON**: Strong consistency is enabled for the multi-zone data. This applies to Standard Edition clusters that are deployed across three zones.
	//
	// - **OFF**: Strong consistency is not enabled for the multi-zone data.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// The specification type of the compute node. Valid values:
	//
	// - **Exclusive**: Dedicated
	//
	// - **General**: General-purpose
	//
	// > This parameter is returned only for PolarDB for MySQL Cluster Edition clusters.
	//
	// example:
	//
	// Exclusive
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	// Indicates whether the instant switchover feature that is compatible with IMCI is supported.
	//
	// example:
	//
	// ON
	SupportInstantSwitchWithImci *string `json:"SupportInstantSwitchWithImci,omitempty" xml:"SupportInstantSwitchWithImci,omitempty"`
	// The details of the tags.
	Tags []*DescribeDBClusterAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The virtual switch ID.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i,cn-hangzhou-g
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) GetAiCreatingTime() *string {
	return s.AiCreatingTime
}

func (s *DescribeDBClusterAttributeResponseBody) GetAiType() *string {
	return s.AiType
}

func (s *DescribeDBClusterAttributeResponseBody) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeDBClusterAttributeResponseBody) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *DescribeDBClusterAttributeResponseBody) GetBlktagTotal() *int64 {
	return s.BlktagTotal
}

func (s *DescribeDBClusterAttributeResponseBody) GetBlktagUsed() *int64 {
	return s.BlktagUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetBranch() *DescribeDBClusterAttributeResponseBodyBranch {
	return s.Branch
}

func (s *DescribeDBClusterAttributeResponseBody) GetBurstingEnabled() *string {
	return s.BurstingEnabled
}

func (s *DescribeDBClusterAttributeResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClusterAttributeResponseBody) GetColumnTable() *string {
	return s.ColumnTable
}

func (s *DescribeDBClusterAttributeResponseBody) GetCompressStorageMode() *string {
	return s.CompressStorageMode
}

func (s *DescribeDBClusterAttributeResponseBody) GetCompressStorageUsed() *int64 {
	return s.CompressStorageUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBNodes() []*DescribeDBClusterAttributeResponseBodyDBNodes {
	return s.DBNodes
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBVersionStatus() *string {
	return s.DBVersionStatus
}

func (s *DescribeDBClusterAttributeResponseBody) GetDataLevel1BackupChainSize() *int64 {
	return s.DataLevel1BackupChainSize
}

func (s *DescribeDBClusterAttributeResponseBody) GetDataSyncMode() *string {
	return s.DataSyncMode
}

func (s *DescribeDBClusterAttributeResponseBody) GetDeletionLock() *int32 {
	return s.DeletionLock
}

func (s *DescribeDBClusterAttributeResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterAttributeResponseBody) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClusterAttributeResponseBody) GetHasCompleteStandbyRes() *bool {
	return s.HasCompleteStandbyRes
}

func (s *DescribeDBClusterAttributeResponseBody) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *DescribeDBClusterAttributeResponseBody) GetImciAutoIndex() *string {
	return s.ImciAutoIndex
}

func (s *DescribeDBClusterAttributeResponseBody) GetImperceptibleSwitch() *string {
	return s.ImperceptibleSwitch
}

func (s *DescribeDBClusterAttributeResponseBody) GetInodeTotal() *int64 {
	return s.InodeTotal
}

func (s *DescribeDBClusterAttributeResponseBody) GetInodeUsed() *int64 {
	return s.InodeUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeDBClusterAttributeResponseBody) GetIsProxyLatestVersion() *bool {
	return s.IsProxyLatestVersion
}

func (s *DescribeDBClusterAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClusterAttributeResponseBody) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDBClusterAttributeResponseBody) GetOrca() *string {
	return s.Orca
}

func (s *DescribeDBClusterAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterAttributeResponseBody) GetProvisionedIops() *string {
	return s.ProvisionedIops
}

func (s *DescribeDBClusterAttributeResponseBody) GetProxyCpuCores() *string {
	return s.ProxyCpuCores
}

func (s *DescribeDBClusterAttributeResponseBody) GetProxyServerlessType() *string {
	return s.ProxyServerlessType
}

func (s *DescribeDBClusterAttributeResponseBody) GetProxyStandardCpuCores() *string {
	return s.ProxyStandardCpuCores
}

func (s *DescribeDBClusterAttributeResponseBody) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeDBClusterAttributeResponseBody) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeDBClusterAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClusterAttributeResponseBody) GetRestoreDataPoint() *string {
	return s.RestoreDataPoint
}

func (s *DescribeDBClusterAttributeResponseBody) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeDBClusterAttributeResponseBody) GetRowCompression() *string {
	return s.RowCompression
}

func (s *DescribeDBClusterAttributeResponseBody) GetSQLSize() *int64 {
	return s.SQLSize
}

func (s *DescribeDBClusterAttributeResponseBody) GetSearchClusterStatus() *string {
	return s.SearchClusterStatus
}

func (s *DescribeDBClusterAttributeResponseBody) GetSearchCompressStorageUsed() *int64 {
	return s.SearchCompressStorageUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetSearchStorageUsed() *int64 {
	return s.SearchStorageUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClusterAttributeResponseBody) GetSourceDBCluster() *string {
	return s.SourceDBCluster
}

func (s *DescribeDBClusterAttributeResponseBody) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DescribeDBClusterAttributeResponseBody) GetStandbyHAMode() *string {
	return s.StandbyHAMode
}

func (s *DescribeDBClusterAttributeResponseBody) GetStorageMax() *int64 {
	return s.StorageMax
}

func (s *DescribeDBClusterAttributeResponseBody) GetStoragePayType() *string {
	return s.StoragePayType
}

func (s *DescribeDBClusterAttributeResponseBody) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *DescribeDBClusterAttributeResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClusterAttributeResponseBody) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBClusterAttributeResponseBody) GetStrictConsistency() *string {
	return s.StrictConsistency
}

func (s *DescribeDBClusterAttributeResponseBody) GetSubCategory() *string {
	return s.SubCategory
}

func (s *DescribeDBClusterAttributeResponseBody) GetSupportInstantSwitchWithImci() *string {
	return s.SupportInstantSwitchWithImci
}

func (s *DescribeDBClusterAttributeResponseBody) GetTags() []*DescribeDBClusterAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeDBClusterAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterAttributeResponseBody) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *DescribeDBClusterAttributeResponseBody) SetAiCreatingTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.AiCreatingTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetAiType(v string) *DescribeDBClusterAttributeResponseBody {
	s.AiType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetArchitecture(v string) *DescribeDBClusterAttributeResponseBody {
	s.Architecture = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetAutoUpgradeMinorVersion(v string) *DescribeDBClusterAttributeResponseBody {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetBlktagTotal(v int64) *DescribeDBClusterAttributeResponseBody {
	s.BlktagTotal = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetBlktagUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.BlktagUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetBranch(v *DescribeDBClusterAttributeResponseBodyBranch) *DescribeDBClusterAttributeResponseBody {
	s.Branch = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetBurstingEnabled(v string) *DescribeDBClusterAttributeResponseBody {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCategory(v string) *DescribeDBClusterAttributeResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetColumnTable(v string) *DescribeDBClusterAttributeResponseBody {
	s.ColumnTable = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCompressStorageMode(v string) *DescribeDBClusterAttributeResponseBody {
	s.CompressStorageMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCompressStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.CompressStorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBNodes(v []*DescribeDBClusterAttributeResponseBodyDBNodes) *DescribeDBClusterAttributeResponseBody {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBType(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBVersionStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDataLevel1BackupChainSize(v int64) *DescribeDBClusterAttributeResponseBody {
	s.DataLevel1BackupChainSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDataSyncMode(v string) *DescribeDBClusterAttributeResponseBody {
	s.DataSyncMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDeletionLock(v int32) *DescribeDBClusterAttributeResponseBody {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetEngine(v string) *DescribeDBClusterAttributeResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetExpired(v string) *DescribeDBClusterAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetHasCompleteStandbyRes(v bool) *DescribeDBClusterAttributeResponseBody {
	s.HasCompleteStandbyRes = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetHotStandbyCluster(v string) *DescribeDBClusterAttributeResponseBody {
	s.HotStandbyCluster = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetImciAutoIndex(v string) *DescribeDBClusterAttributeResponseBody {
	s.ImciAutoIndex = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetImperceptibleSwitch(v string) *DescribeDBClusterAttributeResponseBody {
	s.ImperceptibleSwitch = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetInodeTotal(v int64) *DescribeDBClusterAttributeResponseBody {
	s.InodeTotal = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetInodeUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.InodeUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetIsLatestVersion(v bool) *DescribeDBClusterAttributeResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetIsProxyLatestVersion(v bool) *DescribeDBClusterAttributeResponseBody {
	s.IsProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetLockMode(v string) *DescribeDBClusterAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetOrca(v string) *DescribeDBClusterAttributeResponseBody {
	s.Orca = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetPayType(v string) *DescribeDBClusterAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProvisionedIops(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProxyCpuCores(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProxyCpuCores = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProxyServerlessType(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProxyServerlessType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProxyStandardCpuCores(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProxyStandardCpuCores = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProxyStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetProxyType(v string) *DescribeDBClusterAttributeResponseBody {
	s.ProxyType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRegionId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRestoreDataPoint(v string) *DescribeDBClusterAttributeResponseBody {
	s.RestoreDataPoint = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRestoreType(v string) *DescribeDBClusterAttributeResponseBody {
	s.RestoreType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRowCompression(v string) *DescribeDBClusterAttributeResponseBody {
	s.RowCompression = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSQLSize(v int64) *DescribeDBClusterAttributeResponseBody {
	s.SQLSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSearchClusterStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.SearchClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSearchCompressStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.SearchCompressStorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSearchStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.SearchStorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetServerlessType(v string) *DescribeDBClusterAttributeResponseBody {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSourceDBCluster(v string) *DescribeDBClusterAttributeResponseBody {
	s.SourceDBCluster = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSourceRegionId(v string) *DescribeDBClusterAttributeResponseBody {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStandbyHAMode(v string) *DescribeDBClusterAttributeResponseBody {
	s.StandbyHAMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageMax(v int64) *DescribeDBClusterAttributeResponseBody {
	s.StorageMax = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStoragePayType(v string) *DescribeDBClusterAttributeResponseBody {
	s.StoragePayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageSpace(v int64) *DescribeDBClusterAttributeResponseBody {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageType(v string) *DescribeDBClusterAttributeResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStrictConsistency(v string) *DescribeDBClusterAttributeResponseBody {
	s.StrictConsistency = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSubCategory(v string) *DescribeDBClusterAttributeResponseBody {
	s.SubCategory = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSupportInstantSwitchWithImci(v string) *DescribeDBClusterAttributeResponseBody {
	s.SupportInstantSwitchWithImci = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetTags(v []*DescribeDBClusterAttributeResponseBodyTags) *DescribeDBClusterAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetVPCId(v string) *DescribeDBClusterAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetZoneIds(v string) *DescribeDBClusterAttributeResponseBody {
	s.ZoneIds = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) Validate() error {
	if s.Branch != nil {
		if err := s.Branch.Validate(); err != nil {
			return err
		}
	}
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyBranch struct {
	BranchLsn     *string                                                    `json:"BranchLsn,omitempty" xml:"BranchLsn,omitempty"`
	BranchTime    *string                                                    `json:"BranchTime,omitempty" xml:"BranchTime,omitempty"`
	ChildBranch   []*DescribeDBClusterAttributeResponseBodyBranchChildBranch `json:"ChildBranch,omitempty" xml:"ChildBranch,omitempty" type:"Repeated"`
	ParentInsName *string                                                    `json:"ParentInsName,omitempty" xml:"ParentInsName,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyBranch) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyBranch) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) GetBranchLsn() *string {
	return s.BranchLsn
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) GetBranchTime() *string {
	return s.BranchTime
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) GetChildBranch() []*DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	return s.ChildBranch
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) GetParentInsName() *string {
	return s.ParentInsName
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) SetBranchLsn(v string) *DescribeDBClusterAttributeResponseBodyBranch {
	s.BranchLsn = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) SetBranchTime(v string) *DescribeDBClusterAttributeResponseBodyBranch {
	s.BranchTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) SetChildBranch(v []*DescribeDBClusterAttributeResponseBodyBranchChildBranch) *DescribeDBClusterAttributeResponseBodyBranch {
	s.ChildBranch = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) SetParentInsName(v string) *DescribeDBClusterAttributeResponseBodyBranch {
	s.ParentInsName = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranch) Validate() error {
	if s.ChildBranch != nil {
		for _, item := range s.ChildBranch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyBranchChildBranch struct {
	BranchLsn            *string `json:"BranchLsn,omitempty" xml:"BranchLsn,omitempty"`
	BranchTime           *string `json:"BranchTime,omitempty" xml:"BranchTime,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	HasChild             *bool   `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	InsName              *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyBranchChildBranch) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyBranchChildBranch) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) GetBranchLsn() *string {
	return s.BranchLsn
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) GetBranchTime() *string {
	return s.BranchTime
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) GetHasChild() *bool {
	return s.HasChild
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) GetInsName() *string {
	return s.InsName
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) SetBranchLsn(v string) *DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	s.BranchLsn = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) SetBranchTime(v string) *DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	s.BranchTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) SetHasChild(v bool) *DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	s.HasChild = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) SetInsName(v string) *DescribeDBClusterAttributeResponseBodyBranchChildBranch {
	s.InsName = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyBranchChildBranch) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyDBNodes struct {
	// The number of CPU cores that are added for the elastic scaling feature within seconds.
	//
	// example:
	//
	// 6
	AddedCpuCores *string `json:"AddedCpuCores,omitempty" xml:"AddedCpuCores,omitempty"`
	// The number of CPU cores of the node.
	//
	// example:
	//
	// 2
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The time when the node was created.
	//
	// example:
	//
	// 2020-03-23T21:35:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The CXL remote memory configuration.
	//
	// example:
	//
	// 3072
	DBNodeCXLRemoteMemory *string `json:"DBNodeCXLRemoteMemory,omitempty" xml:"DBNodeCXLRemoteMemory,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The node description.
	//
	// example:
	//
	// test
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The role of the node. Valid values:
	//
	// - **Writer**: the primary node.
	//
	// - **Reader**: a read-only node.
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The status of the node. Valid values:
	//
	// - **Creating**
	//
	// - **Running**
	//
	// - **Deleting**
	//
	// - **Rebooting**
	//
	// - **DBNodeCreating**: A node is being added.
	//
	// - **DBNodeDeleting**: A node is being deleted.
	//
	// - **ClassChanging**: The node specifications are being changed.
	//
	// - **NetAddressCreating**: A network connection is being created.
	//
	// - **NetAddressDeleting**: A network connection is being deleted.
	//
	// - **NetAddressModifying**: A network connection is being modified.
	//
	// - **MinorVersionUpgrading**: The minor version is being upgraded.
	//
	// - **Maintaining**: The instance is being maintained.
	//
	// - **Switching**: A switchover is in progress.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The failover priority. Each node has a failover priority. A larger value indicates a higher priority for the node to be promoted to the primary node during a failover. Valid values: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// Indicates whether the hot replica feature is enabled. Valid values:
	//
	// - **ON**: enabled
	//
	// - **OFF**: disabled
	//
	// example:
	//
	// ON
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	// Indicates whether the In-Memory Column Index (IMCI) feature is enabled. Valid values:
	//
	// - **ON**: enabled
	//
	// - **OFF**: disabled
	//
	// example:
	//
	// ON
	ImciSwitch *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	// The ID of the primary node in a Multi-master Cluster Edition cluster.
	//
	// example:
	//
	// pi-bp18z52akld3*****
	MasterId *string `json:"MasterId,omitempty" xml:"MasterId,omitempty"`
	// The maximum number of concurrent connections to the cluster.
	//
	// example:
	//
	// 8000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum input/output operations per second (IOPS).
	//
	// example:
	//
	// 32000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The memory size of the node. Unit: MB.
	//
	// example:
	//
	// 8192
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The name of the hot replica compute node that corresponds to the node in an architecture where both hot standby storage and hot standby compute are enabled.
	//
	// example:
	//
	// pi-bp18z52mirror*****
	MirrorInsName *string `json:"MirrorInsName,omitempty" xml:"MirrorInsName,omitempty"`
	// The local secondary node in a multi-master cluster.
	//
	// example:
	//
	// pi-****************
	MultiMasterLocalStandby *string `json:"MultiMasterLocalStandby,omitempty" xml:"MultiMasterLocalStandby,omitempty"`
	// The primary node in a multi-master cluster.
	//
	// example:
	//
	// pi-****************
	MultiMasterPrimaryNode *string `json:"MultiMasterPrimaryNode,omitempty" xml:"MultiMasterPrimaryNode,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// off
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// The size of the remote memory. Unit: MB.
	//
	// example:
	//
	// 3072
	RemoteMemorySize *string `json:"RemoteMemorySize,omitempty" xml:"RemoteMemorySize,omitempty"`
	// Indicates whether global consistency (high-performance mode) is enabled for the node. Valid values:
	//
	// - **ON**: enabled
	//
	// - **OFF**: disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
	// The routing weight. Valid values: 1 to 100. Default value: 1.
	//
	// example:
	//
	// 1
	ServerWeight *string `json:"ServerWeight,omitempty" xml:"ServerWeight,omitempty"`
	// The Serverless type of the node. Valid values:
	//
	// - AgileServerless: agile, which indicates a Serverless node.
	//
	// - SteadyServerless: steady, which indicates that Serverless capabilities are added to a node with defined specifications.
	//
	// > 	- This parameter is supported only for Serverless clusters or clusters with defined specifications for which the Serverless feature is enabled. For more information, see [Serverless](https://help.aliyun.com/document_detail/452274.html).
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Indicates whether the node is in the primary zone or a secondary zone. This parameter is mainly used for peer resources. Valid values:
	//
	// - Primary: the primary zone
	//
	// - Standby: a secondary zone
	//
	// example:
	//
	// Primary
	SubCluster *string `json:"SubCluster,omitempty" xml:"SubCluster,omitempty"`
	// The description of the cluster subgroup.
	//
	// example:
	//
	// test Description
	SubGroupDescription *string `json:"SubGroupDescription,omitempty" xml:"SubGroupDescription,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetAddedCpuCores() *string {
	return s.AddedCpuCores
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeCXLRemoteMemory() *string {
	return s.DBNodeCXLRemoteMemory
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeDescription() *string {
	return s.DBNodeDescription
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetFailoverPriority() *int32 {
	return s.FailoverPriority
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetHotReplicaMode() *string {
	return s.HotReplicaMode
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetImciSwitch() *string {
	return s.ImciSwitch
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMasterId() *string {
	return s.MasterId
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMirrorInsName() *string {
	return s.MirrorInsName
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMultiMasterLocalStandby() *string {
	return s.MultiMasterLocalStandby
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetMultiMasterPrimaryNode() *string {
	return s.MultiMasterPrimaryNode
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetOrca() *string {
	return s.Orca
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetRemoteMemorySize() *string {
	return s.RemoteMemorySize
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetSccMode() *string {
	return s.SccMode
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetServerWeight() *string {
	return s.ServerWeight
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetSubCluster() *string {
	return s.SubCluster
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetSubGroupDescription() *string {
	return s.SubGroupDescription
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetAddedCpuCores(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.AddedCpuCores = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetCpuCores(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeCXLRemoteMemory(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeCXLRemoteMemory = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeDescription(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeId(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeRole(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeStatus(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetFailoverPriority(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.FailoverPriority = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetHotReplicaMode(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.HotReplicaMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetImciSwitch(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.ImciSwitch = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMasterId(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MasterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMaxConnections(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMaxIOPS(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMemorySize(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMirrorInsName(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MirrorInsName = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMultiMasterLocalStandby(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MultiMasterLocalStandby = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMultiMasterPrimaryNode(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MultiMasterPrimaryNode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetOrca(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.Orca = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetRemoteMemorySize(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.RemoteMemorySize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetSccMode(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.SccMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetServerWeight(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.ServerWeight = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetServerlessType(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetSubCluster(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.SubCluster = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetSubGroupDescription(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.SubGroupDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// MySQL
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClusterAttributeResponseBodyTags) SetKey(v string) *DescribeDBClusterAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyTags) SetValue(v string) *DescribeDBClusterAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

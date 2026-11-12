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
	SetConnectionResourceQuota(v int64) *DescribeDBClusterAttributeResponseBody
	GetConnectionResourceQuota() *int64
	SetConnectionResourceUsed(v int64) *DescribeDBClusterAttributeResponseBody
	GetConnectionResourceUsed() *int64
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
	SetStorageAutoScale(v string) *DescribeDBClusterAttributeResponseBody
	GetStorageAutoScale() *string
	SetStorageMax(v int64) *DescribeDBClusterAttributeResponseBody
	GetStorageMax() *int64
	SetStoragePayType(v string) *DescribeDBClusterAttributeResponseBody
	GetStoragePayType() *string
	SetStorageSpace(v int64) *DescribeDBClusterAttributeResponseBody
	GetStorageSpace() *int64
	SetStorageType(v string) *DescribeDBClusterAttributeResponseBody
	GetStorageType() *string
	SetStorageUpperBound(v int32) *DescribeDBClusterAttributeResponseBody
	GetStorageUpperBound() *int32
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
	// The start time of the free AI feature.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-13T01:20:28Z
	AiCreatingTime *string `json:"AiCreatingTime,omitempty" xml:"AiCreatingTime,omitempty"`
	// The type of the AI node. Valid values:
	//
	//
	//
	// - **SearchNode**: Search node.
	//
	// - **DLNode**: AI node.
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
	// The minor engine version update method. Valid values:
	//
	// - Auto: Automatic update.
	//
	// - Manual: Manual update.
	//
	// example:
	//
	// Manual
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// The maximum number of blktags in the file system.
	//
	// example:
	//
	// 7864320
	BlktagTotal *int64 `json:"BlktagTotal,omitempty" xml:"BlktagTotal,omitempty"`
	// The current blktag usage.
	//
	// example:
	//
	// 5242880
	BlktagUsed *int64 `json:"BlktagUsed,omitempty" xml:"BlktagUsed,omitempty"`
	// The branch information of the cluster.
	//
	// example:
	//
	// {}
	Branch *DescribeDBClusterAttributeResponseBodyBranch `json:"Branch,omitempty" xml:"Branch,omitempty" type:"Struct"`
	// Indicates whether I/O performance burst is enabled for the ESSD AutoPL cloud disk. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [edition](https://help.aliyun.com/document_detail/183258.html). Valid values:
	//
	// 	- **Normal**: Cluster Edition
	//
	// 	- **Basic**: Single Node Edition
	//
	// 	- **Archive**: X-Engine Edition
	//
	// 	- **NormalMultimaster**: Multi-master Cluster Edition
	//
	// 	- **SENormal**: Standard Edition
	//
	// > 	- PolarDB for PostgreSQL 11 does not support the Single Node Edition.
	//
	// > 	- PolarDB for MySQL 8.0 and 5.7, and PolarDB for PostgreSQL 14 support the Standard Edition.
	//
	// > 	- PolarDB for MySQL 8.0 supports the X-Engine Edition and Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether the column store table is enabled.
	//
	// example:
	//
	// OFF
	ColumnTable *string `json:"ColumnTable,omitempty" xml:"ColumnTable,omitempty"`
	// Indicates whether storage compression is enabled. Valid values:
	//
	// - ON: Enabled.
	//
	// - OFF: Disabled.
	//
	// example:
	//
	// ON
	CompressStorageMode *string `json:"CompressStorageMode,omitempty" xml:"CompressStorageMode,omitempty"`
	// The size of the compressed storage data.
	//
	// >This parameter is returned only when the storage compression feature is enabled for the cluster.
	//
	// example:
	//
	// 15529410560
	CompressStorageUsed *int64 `json:"CompressStorageUsed,omitempty" xml:"CompressStorageUsed,omitempty"`
	// The quota for connection resources (exclusive to AgenticDB).
	//
	// example:
	//
	// 0
	ConnectionResourceQuota *int64 `json:"ConnectionResourceQuota,omitempty" xml:"ConnectionResourceQuota,omitempty"`
	// The number of used connection resources (exclusive to AgenticDB).
	//
	// example:
	//
	// 0
	ConnectionResourceUsed *int64 `json:"ConnectionResourceUsed,omitempty" xml:"ConnectionResourceUsed,omitempty"`
	// The time when the cluster was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the cluster.
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
	// The status of the cluster. For more information about valid values, see [Cluster status table](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The details of the nodes.
	//
	// example:
	//
	// []
	DBNodes []*DescribeDBClusterAttributeResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The type of the database engine.
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
	// The status of the current minor database engine version. Valid values:
	//
	// 	- **Stable**: The current version is stable.
	//
	// 	- **Old**: The current version is outdated. Upgrade to the latest version.
	//
	// 	- **HighRisk**: The current version has critical defects. Upgrade to the latest version immediately.
	//
	// 	- **Beta**: The current version is a beta version.
	//
	// > 	- For more information about how to upgrade the minor database engine version, see [Version upgrade](https://help.aliyun.com/document_detail/158572.html).
	//
	// > 	- This parameter is returned only when the database engine type (**DBType**) is **MySQL**.
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
	// The data replication relationship pattern. Valid values:
	//
	// - **AsyncSync**: Asynchronous.
	//
	// - **SemiSync**: Semi-synchronous.
	//
	// example:
	//
	// AsyncSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// The lock status for cluster deletion. Valid values:
	//
	// 	- **0**: Unlocked. The cluster can be deleted.
	//
	// 	- **1**: Locked. The cluster cannot be deleted.
	//
	// example:
	//
	// 0
	DeletionLock *int32 `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	// The engine of the cluster.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The expiration time of the cluster.
	//
	// > A specific value is returned only for clusters whose billing method is **Prepaid*	- (subscription). For **Postpaid*	- (pay-as-you-go) clusters, an empty value is returned.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired.
	//
	// > This parameter is returned only for clusters whose billing method is **Prepaid*	- (subscription).
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates whether resources for the new primary database are replenished after a cross-zone switch. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	HasCompleteStandbyRes *bool `json:"HasCompleteStandbyRes,omitempty" xml:"HasCompleteStandbyRes,omitempty"`
	// Indicates whether the storage Hot Standby Cluster (and standby compute nodes) is enabled. Valid values:
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
	// - `ON`: Enabled.
	//
	// - `OFF`: Disabled.
	//
	// example:
	//
	// OFF
	ImciAutoIndex *string `json:"ImciAutoIndex,omitempty" xml:"ImciAutoIndex,omitempty"`
	// The seamless switchover feature. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	ImperceptibleSwitch *string `json:"ImperceptibleSwitch,omitempty" xml:"ImperceptibleSwitch,omitempty"`
	// The maximum number of inodes in the file system.
	//
	// example:
	//
	// 6291456
	InodeTotal *int64 `json:"InodeTotal,omitempty" xml:"InodeTotal,omitempty"`
	// The current inode usage.
	//
	// example:
	//
	// 4194304
	InodeUsed *int64 `json:"InodeUsed,omitempty" xml:"InodeUsed,omitempty"`
	// Indicates whether the latest kernel version is used. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the latest version of PolarProxy is used. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsProxyLatestVersion *bool `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// The lock mode. Valid values:
	//
	// - **Unlock**: Unlocked.
	//
	// - **ManualLock**: Manually locked.
	//
	// - **LockByExpiration**: Automatically locked due to cluster expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance window of the cluster. The time is in the `HH:mmZ-HH:mmZ` format and displayed in UTC. For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 00:00 to 01:00 (UTC+8).
	//
	// example:
	//
	// 18:00Z-19:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// ON
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: Pay-as-you-go.
	//
	// - **Prepaid**: Subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// <p id="p_wyg_t4a_glm" props="china" icmsditafragmentmagic=1>The provisioned read and write IOPS for the ESSD AutoPL cloud disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline performance}.</p>
	//
	// <p id="p_6de_jxy_k2g" props="china" icmsditafragmentmagic=1>Baseline performance = min{1,800 + 50 × Capacity, 50,000}.</p>
	//
	// <note id="note_7kj_j0o_rgs" props="china" icmsditafragmentmagic=1>This parameter is supported only when StorageType is set to ESSDAUTOPL.</note>
	//
	// example:
	//
	// 2500
	ProvisionedIops *string `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of CPU cores for PolarProxy.
	//
	// example:
	//
	// 4
	ProxyCpuCores *string `json:"ProxyCpuCores,omitempty" xml:"ProxyCpuCores,omitempty"`
	// The serverless type of the database proxy. Valid values:
	//
	// - AgileServerless: Agile state, which refers to serverless clusters.
	//
	// - SteadyServerless: Steady state, which refers to cluster with defined specifications (clusters whose billing method is subscription or pay-as-you-go).
	//
	// example:
	//
	// SteadyServerless
	ProxyServerlessType *string `json:"ProxyServerlessType,omitempty" xml:"ProxyServerlessType,omitempty"`
	// The number of CPU cores for the standard configuration of the database proxy.
	//
	// example:
	//
	// 2
	ProxyStandardCpuCores *string `json:"ProxyStandardCpuCores,omitempty" xml:"ProxyStandardCpuCores,omitempty"`
	// The status of the database proxy. Valid values:
	//
	// - **Creating**: Creating.
	//
	// - **Running**: Running.
	//
	// - **Deleting**: Releasing.
	//
	// - **Rebooting**: Restarting.
	//
	// - **DBNodeCreating**: Increasing nodes.
	//
	// - **DBNodeDeleting**: Deleting nodes.
	//
	// - **ClassChanging**: Changing node specifications.
	//
	// - **NetAddressCreating**: Creating network connectivity.
	//
	// - **NetAddressDeleting**: Deleting network connectivity.
	//
	// - **NetAddressModifying**: Modifying network connectivity.
	//
	// - **Deleted**: Released.
	//
	// example:
	//
	// Running
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The type of the database proxy. Valid values:
	//
	// - **Exclusive**: Dedicated Enterprise Edition.
	//
	// - **General**: Standard Enterprise Edition.
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
	// The ID of the request.
	//
	// example:
	//
	// 074467EF-86B9-4C23-ACBF-E9B81A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-***************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 	- If RestoreType is set to **RestoreByTime*	- or **RestoreByTimeOss**, this value indicates the point in time to which the cluster is restored.
	//
	// 	- If RestoreType is set to **RestoreByBackupSet*	- or **RestoreByBackupSetOss**, this value indicates the ID of the backup set from which the cluster is restored.
	//
	// <note>This parameter is supported only for clusters restored from backup sets or points in time after June 1, 2024.</note>
	//
	// example:
	//
	// 2179639137
	RestoreDataPoint *string `json:"RestoreDataPoint,omitempty" xml:"RestoreDataPoint,omitempty"`
	// The restoration method of the cluster. Valid values:
	//
	// 	- **RestoreByTime**: Point-in-time restore based on level-1 backups.
	//
	// 	- **RestoreByBackupSet**: Restoration from a backup set based on level-1 backups.
	//
	// 	- **RestoreByTimeOss**: Point-in-time restore based on level-2 backups.
	//
	// 	- **RestoreByBackupSetOss**: Restoration from a backup set based on level-2 backups.
	//
	// 	- **CloneFromSourceCluster**: Cloned from the source cluster.
	//
	// <note>This parameter is supported only for clusters restored from backup sets or points in time after June 1, 2024.</note>
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
	// The storage usage of SQL statements. Unit: bytes. A value of -1 indicates that no data is available.
	//
	// example:
	//
	// 0
	SQLSize *int64 `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
	// The running status of the search node.
	//
	// example:
	//
	// Running
	SearchClusterStatus *string `json:"SearchClusterStatus,omitempty" xml:"SearchClusterStatus,omitempty"`
	// The size of compressed storage data on the search node.
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
	// The serverless type of the cluster. Valid values:
	//
	// - AgileServerless: Agile state, which refers to serverless clusters.
	//
	// - SteadyServerless: Steady state, which refers to cluster with defined specifications that have serverless capabilities enabled.
	//
	// > This parameter is supported only for serverless clusters or cluster with defined specifications that have the serverless feature enabled.
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The ID of the source cluster.
	//
	// <note>This parameter is supported only for clusters restored from backup sets or points in time after June 1, 2024.</note>
	//
	// example:
	//
	// pc-pz51ziv48317b2880
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The region ID of the source cluster.
	//
	// <note>This parameter is returned only when the source cluster ID exists.</note>
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The cross-zone disaster recovery mode. Valid values:
	//
	// - **ON**: Cross-zone disaster recovery is enabled.
	//
	// - **OFF**: Cross-zone disaster recovery is disabled.
	//
	// - **0**: Customer drill mode.
	//
	// example:
	//
	// OFF
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// Indicates whether automatic storage scaling is enabled for the Standard Edition cluster. A value of Enable indicates that the feature is enabled, and a value of Disable indicates that the feature is disabled.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The maximum storage capacity for the current cluster specifications. Unit: bytes.
	//
	// example:
	//
	// 10995116277760
	StorageMax *int64 `json:"StorageMax,omitempty" xml:"StorageMax,omitempty"`
	// The storage billing method. Valid values:
	//
	// - **Postpaid**: Pay-as-you-go based on storage capacity.
	//
	// - **Prepaid**: Subscription based on storage space.
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage space for subscription billing. Unit: bytes.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type. The value is fixed as **HighPerformance**.
	//
	// example:
	//
	// HighPerformance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The upper limit for automatic storage scaling of the Standard Edition cluster, in GB.
	//
	// example:
	//
	// 800
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// The amount of used storage space, in bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether strong data consistency across multiple zones is enabled for the cluster. Valid values:
	//
	// - **ON**: Strong data consistency across multiple zones is enabled. This applies to the Standard Edition in a three-zone deployment scenario.
	//
	// - **OFF**: Strong data consistency across multiple zones is disabled.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// The specification type of the compute node. Valid values:
	//
	// 	- **Exclusive**: Dedicated
	//
	// 	- **General**: General-purpose
	//
	// > This parameter is returned only for PolarDB for MySQL clusters of the Cluster Edition.
	//
	// example:
	//
	// Exclusive
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	// Indicates whether the failover with hot replica feature compatible with IMCI is supported.
	//
	// example:
	//
	// ON
	SupportInstantSwitchWithImci *string `json:"SupportInstantSwitchWithImci,omitempty" xml:"SupportInstantSwitchWithImci,omitempty"`
	// The details of the tags.
	//
	// example:
	//
	// []
	Tags []*DescribeDBClusterAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
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

func (s *DescribeDBClusterAttributeResponseBody) GetConnectionResourceQuota() *int64 {
	return s.ConnectionResourceQuota
}

func (s *DescribeDBClusterAttributeResponseBody) GetConnectionResourceUsed() *int64 {
	return s.ConnectionResourceUsed
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

func (s *DescribeDBClusterAttributeResponseBody) GetStorageAutoScale() *string {
	return s.StorageAutoScale
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

func (s *DescribeDBClusterAttributeResponseBody) GetStorageUpperBound() *int32 {
	return s.StorageUpperBound
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

func (s *DescribeDBClusterAttributeResponseBody) SetConnectionResourceQuota(v int64) *DescribeDBClusterAttributeResponseBody {
	s.ConnectionResourceQuota = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetConnectionResourceUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.ConnectionResourceUsed = &v
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

func (s *DescribeDBClusterAttributeResponseBody) SetStorageAutoScale(v string) *DescribeDBClusterAttributeResponseBody {
	s.StorageAutoScale = &v
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

func (s *DescribeDBClusterAttributeResponseBody) SetStorageUpperBound(v int32) *DescribeDBClusterAttributeResponseBody {
	s.StorageUpperBound = &v
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
	// The log sequence number (LSN) of the current branch.
	//
	// example:
	//
	// 2/2
	BranchLsn *string `json:"BranchLsn,omitempty" xml:"BranchLsn,omitempty"`
	// The branch time of the current branch. The time follows the RFC 3339 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	BranchTime *string `json:"BranchTime,omitempty" xml:"BranchTime,omitempty"`
	// The list of child branches.
	//
	// example:
	//
	// []
	ChildBranch []*DescribeDBClusterAttributeResponseBodyBranchChildBranch `json:"ChildBranch,omitempty" xml:"ChildBranch,omitempty" type:"Repeated"`
	// The name of the parent branch instance.
	//
	// example:
	//
	// pc-2bbb22b2bb22bb222
	ParentInsName *string `json:"ParentInsName,omitempty" xml:"ParentInsName,omitempty"`
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
	// The log sequence number (LSN) of the child branch.
	//
	// example:
	//
	// 2/2
	BranchLsn *string `json:"BranchLsn,omitempty" xml:"BranchLsn,omitempty"`
	// The branch time of the child branch. The time follows the RFC 3339 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	BranchTime *string `json:"BranchTime,omitempty" xml:"BranchTime,omitempty"`
	// The description of the child branch cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// Indicates whether lower-level child branches exist.
	//
	// example:
	//
	// false
	HasChild *bool `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	// The name of the child branch instance.
	//
	// example:
	//
	// pc-2bbb22b2bb22bb222
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
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
	// The number of CPU cores added by automatically scaling local resources.
	//
	// example:
	//
	// 6
	AddedCpuCores *string `json:"AddedCpuCores,omitempty" xml:"AddedCpuCores,omitempty"`
	// The number of CPU cores for the node.
	//
	// example:
	//
	// 2
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The time when the node was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-03-23T21:35:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The Compute Express Link (CXL) remote memory configuration.
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
	// The description of the node.
	//
	// example:
	//
	// test
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The role of the node. Valid values:
	//
	// - **Writer**: Primary node.
	//
	// - **Reader**: Read-only node.
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- **Creating**: The node is being created.
	//
	// 	- **Running**: The node is running.
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// 	- **Rebooting**: The node is restarting.
	//
	// 	- **DBNodeCreating**: A node is being added.
	//
	// 	- **DBNodeDeleting**: A node is being deleted.
	//
	// 	- **ClassChanging**: The node specifications are being changed.
	//
	// 	- **NetAddressCreating**: A network connection is being created.
	//
	// 	- **NetAddressDeleting**: A network connection is being deleted.
	//
	// 	- **NetAddressModifying**: A network connection is being modified.
	//
	// 	- **MinorVersionUpgrading**: The minor version is being upgraded.
	//
	// 	- **Maintaining**: The instance is under maintenance.
	//
	// 	- **Switching**: A switchover is in progress.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// Indicates whether the node uses the exclusive host mode.
	//
	// example:
	//
	// false
	DedicatedHostModel *bool `json:"DedicatedHostModel,omitempty" xml:"DedicatedHostModel,omitempty"`
	// The failover priority. Each node has a failover priority, which determines the probability of the node being elected as the primary node during a failover. A higher value indicates a higher priority.
	//
	// Valid values: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// Indicates whether hot standby is enabled. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
	//
	// example:
	//
	// ON
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	// Indicates whether the In-Memory Column Index (IMCI) feature is enabled. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
	//
	// example:
	//
	// ON
	ImciSwitch *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	// The ID of the primary node in the Multi-master Cluster Edition.
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
	// The maximum number of I/O requests per second, namely IOPS.
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
	// The name of the hot replica corresponding to the node in the architecture with hot standby storage and hot standby compute enabled.
	//
	// example:
	//
	// pi-bp18z52mirror*****
	MirrorInsName *string `json:"MirrorInsName,omitempty" xml:"MirrorInsName,omitempty"`
	// The local secondary node in the multi-master architecture.
	//
	// example:
	//
	// pi-****************
	MultiMasterLocalStandby *string `json:"MultiMasterLocalStandby,omitempty" xml:"MultiMasterLocalStandby,omitempty"`
	// The primary node in the multi-master architecture.
	//
	// example:
	//
	// pi-****************
	MultiMasterPrimaryNode *string `json:"MultiMasterPrimaryNode,omitempty" xml:"MultiMasterPrimaryNode,omitempty"`
	// The Orca feature status. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// The maximum remote memory capacity of the node.
	//
	// example:
	//
	// 0
	RemoteMemoryMax *int32 `json:"RemoteMemoryMax,omitempty" xml:"RemoteMemoryMax,omitempty"`
	// The minimum remote memory capacity of the node.
	//
	// example:
	//
	// 0
	RemoteMemoryMin *int32 `json:"RemoteMemoryMin,omitempty" xml:"RemoteMemoryMin,omitempty"`
	// The recommended remote memory capacity of the node.
	//
	// example:
	//
	// 0
	RemoteMemoryRecommended *int32 `json:"RemoteMemoryRecommended,omitempty" xml:"RemoteMemoryRecommended,omitempty"`
	// The size of the remote memory. Unit: MB.
	//
	// example:
	//
	// 3072
	RemoteMemorySize *string `json:"RemoteMemorySize,omitempty" xml:"RemoteMemorySize,omitempty"`
	// Indicates whether the global consistency (high-performance mode) feature is enabled for the node. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
	// The routing weight.
	//
	// Valid values: 1 to 100. Default value: 1.
	//
	// example:
	//
	// 1
	ServerWeight *string `json:"ServerWeight,omitempty" xml:"ServerWeight,omitempty"`
	// The serverless type of the node. Valid values:
	//
	// - AgileServerless: Agile state, which refers to serverless nodes.
	//
	// - SteadyServerless: Steady state, which refers to nodes in a cluster with defined specifications that have serverless capabilities enabled.
	//
	// > 	- This parameter is supported only for serverless clusters or a cluster with defined specifications that have the serverless feature enabled. For more information, see [Serverless](https://help.aliyun.com/document_detail/452274.html).
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Indicates whether the node is in the primary zone or secondary zone. This is mainly used for resource symmetry.
	//
	// Valid values:
	//
	// - Primary: Primary zone.
	//
	// - Standby: Secondary zone.
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

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetDedicatedHostModel() *bool {
	return s.DedicatedHostModel
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

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetRemoteMemoryMax() *int32 {
	return s.RemoteMemoryMax
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetRemoteMemoryMin() *int32 {
	return s.RemoteMemoryMin
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) GetRemoteMemoryRecommended() *int32 {
	return s.RemoteMemoryRecommended
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

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDedicatedHostModel(v bool) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DedicatedHostModel = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetRemoteMemoryMax(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.RemoteMemoryMax = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetRemoteMemoryMin(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.RemoteMemoryMin = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetRemoteMemoryRecommended(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.RemoteMemoryRecommended = &v
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

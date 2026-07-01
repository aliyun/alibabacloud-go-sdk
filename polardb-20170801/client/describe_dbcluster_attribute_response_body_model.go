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
	// The start time of the free AI trial.
	//
	// example:
	//
	// 2024-03-13T01:20:28Z
	AiCreatingTime *string `json:"AiCreatingTime,omitempty" xml:"AiCreatingTime,omitempty"`
	// The AI node type. Valid values:
	//
	//
	//
	// - **SearchNode**: search node.
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
	// The minor version update method. Valid values:
	//
	// - Auto: automatic update.
	//
	// - Manual: manual update.
	//
	// example:
	//
	// Manual
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// The maximum number of blktags in the file system.
	//
	// example:
	//
	// 7,864,320
	BlktagTotal *int64 `json:"BlktagTotal,omitempty" xml:"BlktagTotal,omitempty"`
	// The current blktag usage.
	//
	// example:
	//
	// 5,242,880
	BlktagUsed *int64                                        `json:"BlktagUsed,omitempty" xml:"BlktagUsed,omitempty"`
	Branch     *DescribeDBClusterAttributeResponseBodyBranch `json:"Branch,omitempty" xml:"Branch,omitempty" type:"Struct"`
	// Indicates whether I/O performance burst is enabled for the ESSD AutoPL cloud disk. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [edition](https://help.aliyun.com/document_detail/183258.html) of the cluster. Valid values:
	//
	// 	- **Normal**: Cluster Edition
	//
	// 	- **Basic**: Single Node Edition
	//
	// 	- **Archive**: PolarDB X-Engine Edition
	//
	// 	- **NormalMultimaster**: Multi-master Cluster Edition
	//
	// 	- **SENormal**: PolarDB for MySQL Standard Edition
	//
	// > 	- PolarDB for PostgreSQL (PostgreSQL 11) does not support Single Node Edition.
	//
	// >	- PolarDB for MySQL 8.0, PolarDB for MySQL 5.7, and PolarDB for PostgreSQL (PostgreSQL 14) support PolarDB for MySQL Standard Edition.
	//
	// >	- PolarDB for MySQL 8.0 supports PolarDB X-Engine Edition and Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether column store tables are enabled.
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
	// The compressed storage data size.
	//
	// >This parameter is returned only when the storage compression feature is enabled for the cluster.
	//
	// example:
	//
	// 15529410560
	CompressStorageUsed     *int64 `json:"CompressStorageUsed,omitempty" xml:"CompressStorageUsed,omitempty"`
	ConnectionResourceQuota *int64 `json:"ConnectionResourceQuota,omitempty" xml:"ConnectionResourceQuota,omitempty"`
	ConnectionResourceUsed  *int64 `json:"ConnectionResourceUsed,omitempty" xml:"ConnectionResourceUsed,omitempty"`
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
	// The cluster status. For more information, see [Cluster status table](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The details of nodes.
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
	// The status of the current minor engine version. Valid values:
	//
	// 	- **Stable**: The current version is stable.
	//
	// 	- **Old**: The current version is outdated. Upgrade to the latest version.
	//
	// 	- **HighRisk**: The current version has critical defects. Upgrade to the latest version immediately.
	//
	// 	- **Beta**: The current version is a beta version.
	//
	// > 	- For information about how to upgrade the minor engine version, see [Version upgrade](https://help.aliyun.com/document_detail/158572.html).
	//
	// > 	- This parameter is returned only when the database engine type (**DBType**) is **MySQL**.
	//
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// The total size of level-1 backups (snapshots), in bytes.
	//
	// example:
	//
	// 74448896
	DataLevel1BackupChainSize *int64 `json:"DataLevel1BackupChainSize,omitempty" xml:"DataLevel1BackupChainSize,omitempty"`
	// The data replication relationship mode. Valid values:
	//
	// - **AsyncSync**: asynchronous
	//
	// - **SemiSync**: semi-synchronous
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
	// The cluster engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The expiration time of the cluster.
	//
	// > A specific value is returned only for clusters whose billing method is **Prepaid*	- (subscription). An empty value is returned for **Postpaid*	- (pay-as-you-go) clusters.
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
	// Indicates whether resources are replenished for the new primary database after a cross-zone failover. Valid values:
	//
	// - **true**: Resources are replenished.
	//
	// - **false**: Resources are not replenished.
	//
	// example:
	//
	// false
	HasCompleteStandbyRes *bool `json:"HasCompleteStandbyRes,omitempty" xml:"HasCompleteStandbyRes,omitempty"`
	// Indicates whether the Hot Standby Cluster (and standby compute nodes) is enabled. Valid values:
	//
	// - **StandbyClusterON**: The Hot Standby Cluster and standby compute nodes are enabled.
	//
	// - **StandbyClusterOFF**: The Hot Standby Cluster and standby compute nodes are disabled.
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
	// The failover with hot replica feature. Valid values:
	//
	// - `true`: enabled.
	//
	// - `false`: disabled.
	//
	// example:
	//
	// true
	ImperceptibleSwitch *string `json:"ImperceptibleSwitch,omitempty" xml:"ImperceptibleSwitch,omitempty"`
	// The maximum number of inodes in the file system.
	//
	// example:
	//
	// 6,291,456
	InodeTotal *int64 `json:"InodeTotal,omitempty" xml:"InodeTotal,omitempty"`
	// The current inode usage.
	//
	// example:
	//
	// 4,194,304
	InodeUsed *int64 `json:"InodeUsed,omitempty" xml:"InodeUsed,omitempty"`
	// Indicates whether the cluster runs the latest Milvus version. Valid values:
	//
	// - **true**: The cluster runs the latest Milvus version.
	//
	// - **false**: The cluster does not run the latest Milvus version.
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the database proxy is the latest version. Valid values:
	//
	// - **true**: The database proxy is the latest version.
	//
	// - **false**: The database proxy is not the latest version.
	//
	// example:
	//
	// false
	IsProxyLatestVersion *bool `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// The lock mode. Valid values:
	//
	// - **Unlock**: not locked.
	//
	// - **ManualLock**: manually locked.
	//
	// - **LockByExpiration**: automatically locked due to cluster expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance window of the cluster, in the `HH:mmZ-HH:mmZ` format (UTC). For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 00:00 to 01:00 (UTC+08:00).
	//
	// example:
	//
	// 18:00Z-19:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
	// <p id="p_wyg_t4a_glm" props="china" icmsditafragmentmagic=1>The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}.</p>
	//
	// <p id="p_6de_jxy_k2g" props="china" icmsditafragmentmagic=1>Baseline performance = min{1,800 + 50 × capacity, 50,000}.</p>
	//
	// <note id="note_7kj_j0o_rgs" props="china" icmsditafragmentmagic=1>This parameter is supported only when StorageType is set to ESSDAUTOPL.</note>
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
	// The serverless type of the database proxy. Valid values:
	//
	// - AgileServerless: agile serverless cluster.
	//
	// - SteadyServerless: steady serverless, which is a cluster with defined specifications (subscription or pay-as-you-go billing).
	//
	// example:
	//
	// SteadyServerless
	ProxyServerlessType *string `json:"ProxyServerlessType,omitempty" xml:"ProxyServerlessType,omitempty"`
	// The number of CPU cores in the standard configuration of the database proxy.
	//
	// example:
	//
	// 2
	ProxyStandardCpuCores *string `json:"ProxyStandardCpuCores,omitempty" xml:"ProxyStandardCpuCores,omitempty"`
	// The status of the database proxy. Valid values:
	//
	// - **Creating**: being created.
	//
	// - **Running**: running.
	//
	// - **Deleting**: being released.
	//
	// - **Rebooting**: being restarted.
	//
	// - **DBNodeCreating**: adding a node.
	//
	// - **DBNodeDeleting**: deleting a node.
	//
	// - **ClassChanging**: changing node specifications.
	//
	// - **NetAddressCreating**: creating network connectivity.
	//
	// - **NetAddressDeleting**: deleting network connectivity.
	//
	// - **NetAddressModifying**: modifying network connectivity.
	//
	// - **Deleted**: released.
	//
	// example:
	//
	// Running
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The type of the database proxy. Valid values:
	//
	// - **Exclusive**: Dedicated Enterprise Edition
	//
	// - **General**: Standard Enterprise Edition
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
	// 	- If RestoreType is **RestoreByTime*	- or **RestoreByTimeOss**, this value indicates the point in time to which the cluster is restored.
	//
	// 	- If RestoreType is **RestoreByBackupSet*	- or **RestoreByBackupSetOss**, this value indicates the backup set ID used for the restoration.
	//
	// <note>This parameter is supported only for clusters restored from a backup set or point in time after June 1, 2024.</note>
	//
	// example:
	//
	// 2179639137
	RestoreDataPoint *string `json:"RestoreDataPoint,omitempty" xml:"RestoreDataPoint,omitempty"`
	// The restoration method of the cluster. Valid values:
	//
	// 	- **RestoreByTime**: point-in-time restore based on a level-1 backup.
	//
	// 	- **RestoreByBackupSet**: restore from a level-1 backup set.
	//
	// 	- **RestoreByTimeOss**: point-in-time restore based on a level-2 backup.
	//
	// 	- **RestoreByBackupSetOss**: restore from a level-2 backup set.
	//
	// 	- **CloneFromSourceCluster**: clone from the source cluster.
	//
	// <note>This parameter is supported only for clusters restored from a backup set or point in time after June 1, 2024.</note>
	//
	// example:
	//
	// RestoreByTime
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The row compression setting.
	//
	// example:
	//
	// OFF
	RowCompression *string `json:"RowCompression,omitempty" xml:"RowCompression,omitempty"`
	// The storage size of SQL statements, in bytes. A value of -1 indicates that no data is available.
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
	// The compressed storage data size of the search node.
	//
	// >This parameter is returned only when the storage compression feature is enabled for the cluster.
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
	// - AgileServerless: agile serverless cluster.
	//
	// - SteadyServerless: steady serverless, which is a cluster with defined specifications that has the serverless feature enabled.
	//
	// > This parameter is supported only for serverless clusters or clusters with defined specifications that have the serverless feature enabled.
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The source cluster ID.
	//
	// <note>This parameter is supported only for clusters restored from a backup set or point in time after June 1, 2024.</note>
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
	// The maximum storage capacity for the current cluster specifications, in bytes.
	//
	// example:
	//
	// 10995116277760
	StorageMax *int64 `json:"StorageMax,omitempty" xml:"StorageMax,omitempty"`
	// The billing method for storage. Valid values:
	//
	// - **Postpaid**: pay-by-capacity (pay-as-you-go).
	//
	// - **Prepaid**: pay-by-space (subscription).
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage space for pay-by-space (subscription) billing. Unit: bytes.
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
	// The storage usage, in bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether multi-zone strong data consistency is enabled for the cluster. Valid values:
	//
	// - **ON**: Multi-zone strong data consistency is enabled. This applies to PolarDB for MySQL Standard Edition with three-zone deployment.
	//
	// - **OFF**: Multi-zone strong data consistency is not enabled.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// The specification type of compute nodes. Valid values:
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
	// Indicates whether failover with hot replica is supported with IMCI compatibility.
	//
	// example:
	//
	// ON
	SupportInstantSwitchWithImci *string `json:"SupportInstantSwitchWithImci,omitempty" xml:"SupportInstantSwitchWithImci,omitempty"`
	// The details of tags.
	Tags []*DescribeDBClusterAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
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
	// The number of CPU cores added by second-level rapid scaling.
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
	// - **Writer**: read/write node.
	//
	// - **Reader**: read-only node.
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The node status. Valid values:
	//
	// 	- **Creating**: being created.
	//
	// 	- **Running**: running.
	//
	// 	- **Deleting**: being deleted.
	//
	// 	- **Rebooting**: being restarted.
	//
	// 	- **DBNodeCreating**: adding a node.
	//
	// 	- **DBNodeDeleting**: deleting a node.
	//
	// 	- **ClassChanging**: changing node specifications.
	//
	// 	- **NetAddressCreating**: creating network connectivity.
	//
	// 	- **NetAddressDeleting**: deleting network connectivity.
	//
	// 	- **NetAddressModifying**: modifying network connectivity.
	//
	// 	- **MinorVersionUpgrading**: performing a minor engine version upgrade.
	//
	// 	- **Maintaining**: instance under maintenance.
	//
	// 	- **Switching**: being switched.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The failover priority. Each node has a failover priority that determines the probability of the node being elected as the primary node during a failover. A higher value indicates a higher priority.
	//
	// Valid values: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// Indicates whether hot standby is enabled. Valid values:
	//
	// - **ON**: enabled.
	//
	// - **OFF**: disabled.
	//
	// example:
	//
	// ON
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	// Indicates whether In-Memory Column Index (IMCI) is enabled. Valid values:
	//
	// - **ON**: enabled.
	//
	// - **OFF**: disabled.
	//
	// example:
	//
	// ON
	ImciSwitch *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	// The primary node ID of the Multi-master Cluster Edition.
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
	// The maximum number of I/O requests per second (IOPS).
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
	// The name of the hot replica compute node that corresponds to the node in the Hot Standby Cluster and compute architecture.
	//
	// example:
	//
	// pi-bp18z52mirror*****
	MirrorInsName *string `json:"MirrorInsName,omitempty" xml:"MirrorInsName,omitempty"`
	// The multi-master local standby node.
	//
	// example:
	//
	// pi-****************
	MultiMasterLocalStandby *string `json:"MultiMasterLocalStandby,omitempty" xml:"MultiMasterLocalStandby,omitempty"`
	// The multi-master primary node.
	//
	// example:
	//
	// pi-****************
	MultiMasterPrimaryNode *string `json:"MultiMasterPrimaryNode,omitempty" xml:"MultiMasterPrimaryNode,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// off
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// The remote memory size. Unit: MB.
	//
	// example:
	//
	// 3072
	RemoteMemorySize *string `json:"RemoteMemorySize,omitempty" xml:"RemoteMemorySize,omitempty"`
	// Indicates whether the global consistency (high-performance mode) feature is enabled for the node. Valid values:
	//
	// - **ON**: The feature is enabled.
	//
	// - **OFF**: The feature is disabled.
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
	// - AgileServerless: agile serverless node.
	//
	// - SteadyServerless: steady serverless node, which is a node with defined specifications that has the serverless capability enabled.
	//
	// > This parameter is supported only for serverless clusters or clusters with defined specifications that have the serverless feature enabled. For more information, see [Serverless](https://help.aliyun.com/document_detail/452274.html).
	//
	// example:
	//
	// SteadyServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Indicates whether the node is in the primary zone or the secondary zone. This parameter is mainly used for resource-equivalent deployments.
	//
	// Valid values:
	//
	// - Primary: primary zone.
	//
	// - Standby: secondary zone.
	//
	// example:
	//
	// Primary
	SubCluster *string `json:"SubCluster,omitempty" xml:"SubCluster,omitempty"`
	// The cluster subgroup description.
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

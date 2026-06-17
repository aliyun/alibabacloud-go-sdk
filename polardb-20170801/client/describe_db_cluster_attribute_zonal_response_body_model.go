// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbClusterAttributeZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiCreatingTime(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetAiCreatingTime() *string
	SetAiType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetAiType() *string
	SetArchitecture(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetArchitecture() *string
	SetAutoUpgradeMinorVersion(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetAutoUpgradeMinorVersion() *string
	SetBlktagTotal(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetBlktagTotal() *int64
	SetBlktagUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetBlktagUsed() *int64
	SetBurstingEnabled(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetBurstingEnabled() *string
	SetCategory(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetCategory() *string
	SetCompressStorageMode(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetCompressStorageMode() *string
	SetCompressStorageUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetCompressStorageUsed() *int64
	SetCreationTime(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetCreationTime() *string
	SetDBClusterClass(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBClusterClass() *string
	SetDBClusterDescription(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBClusterId() *string
	SetDBClusterNetworkType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBClusterNetworkType() *string
	SetDBClusterStatus(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBClusterStatus() *string
	SetDBNodes(v []*DescribeDbClusterAttributeZonalResponseBodyDBNodes) *DescribeDbClusterAttributeZonalResponseBody
	GetDBNodes() []*DescribeDbClusterAttributeZonalResponseBodyDBNodes
	SetDBType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBVersion() *string
	SetDBVersionStatus(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDBVersionStatus() *string
	SetDataLevel1BackupChainSize(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetDataLevel1BackupChainSize() *int64
	SetDataSyncMode(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetDataSyncMode() *string
	SetDeletionLock(v int32) *DescribeDbClusterAttributeZonalResponseBody
	GetDeletionLock() *int32
	SetEngine(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetEngine() *string
	SetExpireTime(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetExpireTime() *string
	SetExpired(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetExpired() *string
	SetHasCompleteStandbyRes(v bool) *DescribeDbClusterAttributeZonalResponseBody
	GetHasCompleteStandbyRes() *bool
	SetHotStandbyCluster(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetHotStandbyCluster() *string
	SetImciAutoIndex(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetImciAutoIndex() *string
	SetImperceptibleSwitch(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetImperceptibleSwitch() *string
	SetInodeTotal(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetInodeTotal() *int64
	SetInodeUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetInodeUsed() *int64
	SetIsLatestVersion(v bool) *DescribeDbClusterAttributeZonalResponseBody
	GetIsLatestVersion() *bool
	SetIsProxyLatestVersion(v bool) *DescribeDbClusterAttributeZonalResponseBody
	GetIsProxyLatestVersion() *bool
	SetLockMode(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetLockMode() *string
	SetMaintainTime(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetMaintainTime() *string
	SetOrca(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetOrca() *string
	SetPayType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetPayType() *string
	SetProvisionedIops(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProvisionedIops() *string
	SetProxyCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProxyCpuCores() *string
	SetProxyServerlessType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProxyServerlessType() *string
	SetProxyStandardCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProxyStandardCpuCores() *string
	SetProxyStatus(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProxyStatus() *string
	SetProxyType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetProxyType() *string
	SetRegionId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetResourceGroupId() *string
	SetRestoreDataPoint(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetRestoreDataPoint() *string
	SetRestoreType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetRestoreType() *string
	SetRowCompression(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetRowCompression() *string
	SetSQLSize(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetSQLSize() *int64
	SetServerlessType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetServerlessType() *string
	SetSourceDBCluster(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetSourceDBCluster() *string
	SetSourceRegionId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetSourceRegionId() *string
	SetStandbyHAMode(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetStandbyHAMode() *string
	SetStorageMax(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetStorageMax() *int64
	SetStoragePayType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetStoragePayType() *string
	SetStorageSpace(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetStorageSpace() *int64
	SetStorageType(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetStorageType() *string
	SetStorageUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody
	GetStorageUsed() *int64
	SetStrictConsistency(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetStrictConsistency() *string
	SetSubCategory(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetSubCategory() *string
	SetSupportInstantSwitchWithImci(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetSupportInstantSwitchWithImci() *string
	SetTags(v []*DescribeDbClusterAttributeZonalResponseBodyTags) *DescribeDbClusterAttributeZonalResponseBody
	GetTags() []*DescribeDbClusterAttributeZonalResponseBodyTags
	SetVPCId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetVSwitchId() *string
	SetZoneIds(v string) *DescribeDbClusterAttributeZonalResponseBody
	GetZoneIds() *string
}

type DescribeDbClusterAttributeZonalResponseBody struct {
	// The start time of the free trial for the AI feature
	//
	// example:
	//
	// 2024-03-13T01:20:28Z
	AiCreatingTime *string `json:"AiCreatingTime,omitempty" xml:"AiCreatingTime,omitempty"`
	// The AI node type. Valid values:
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
	// x86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The upgrade method for the minor version
	//
	// - Auto: Automatic upgrade
	//
	// - Manual: Manual upgrade
	//
	// example:
	//
	// Auto
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
	BlktagUsed *int64 `json:"BlktagUsed,omitempty" xml:"BlktagUsed,omitempty"`
	// Indicates whether the I/O performance burst feature is enabled for the ESSD AutoPL disk. Valid values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [product edition](https://help.aliyun.com/document_detail/183258.html). Valid values:
	//
	// - **Normal**: Cluster Edition
	//
	// - **Basic**: Single Node
	//
	// - **Archive**: X-Engine
	//
	// - **NormalMultimaster**: Multi-master Cluster Edition
	//
	// - **SENormal**: Standard Edition
	//
	// > 	- PolarDB for PostgreSQL clusters of PostgreSQL 11 do not support the single node edition.
	//
	// >
	//
	// > 	- PolarDB for MySQL clusters of 8.0 and 5.7, and PolarDB for PostgreSQL clusters of PostgreSQL 14 support the Standard Edition.
	//
	// >
	//
	// > 	- PolarDB for MySQL clusters of 8.0 support X-Engine and the Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether storage compression is enabled. Valid values:
	//
	// - ON: Enabled
	//
	// - OFF: Disabled
	//
	// example:
	//
	// ON
	CompressStorageMode *string `json:"CompressStorageMode,omitempty" xml:"CompressStorageMode,omitempty"`
	// The size of the compressed data in storage.
	//
	// > This parameter is returned only when storage compression is enabled for the cluster.
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
	// The cluster specifications.
	//
	// example:
	//
	// polar.mysql.x8.medium.c
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
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
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The cluster status. Valid values:
	//
	// - **Creating**: The cluster is being created.
	//
	// - **Running**: The cluster is running.
	//
	// - **Deleting**: The cluster is being released.
	//
	// - **Rebooting**: The cluster is being restarted.
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
	// - **Deleted**: The cluster is released.
	//
	// - **ClassChanged**: Resources are being reclaimed after an upgrade or downgrade.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The details of the node information.
	DBNodes []*DescribeDbClusterAttributeZonalResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
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
	// The status of the current minor version of the database. Valid values:
	//
	// - **Stable**: The current version is stable.
	//
	// - **Old**: The current version is old. We recommend that you upgrade the cluster to the latest version.
	//
	// - **HighRisk**: The current version has critical bugs. Upgrade the cluster to the latest version immediately.
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
	// The data replication relationship mode. Valid values:
	//
	// - **AsyncSync**: Asynchronous
	//
	// - **SemiSync**: Semi-synchronous
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
	// The database engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The expiration time of the cluster.
	//
	// > This parameter is returned only for **Prepaid*	- (subscription) clusters. An empty value is returned for **Postpaid*	- (pay-as-you-go) clusters.
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
	// Indicates whether resources of the new primary database are supplemented after a cross-zone failover. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	HasCompleteStandbyRes *bool `json:"HasCompleteStandbyRes,omitempty" xml:"HasCompleteStandbyRes,omitempty"`
	// Indicates whether the hot standby storage cluster (and standby compute nodes) feature is enabled. Valid values:
	//
	// - **StandbyClusterON**: The hot standby storage cluster feature or both the hot standby storage cluster and standby compute node features are enabled.
	//
	// - **StandbyClusterOFF**: The hot standby storage cluster feature or both the hot standby storage cluster and standby compute node features are disabled.
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
	// Failover with hot replica. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// False
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
	// Indicates whether the cluster uses the latest kernel version. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the database proxy uses the latest version. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
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
	// - **LockByExpiration**: The cluster is automatically locked upon expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance window of the cluster. The time is in the `HH:mmZ-HH:mmZ` format and is displayed in UTC. For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 00:00 to 01:00 (UTC+8).
	//
	// example:
	//
	// 18:00Z-19:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: Enabled
	//
	// - off: Disabled
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
	// <props="china">
	//
	// The read and write IOPS that are preconfigured for the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
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
	// 1000
	ProvisionedIops *string `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of CPU cores for the database proxy.
	//
	// example:
	//
	// 4
	ProxyCpuCores *string `json:"ProxyCpuCores,omitempty" xml:"ProxyCpuCores,omitempty"`
	// The Serverless type of the database proxy. Valid values:
	//
	// - AgileServerless: Agile. This indicates a Serverless cluster.
	//
	// - SteadyServerless: Steady. This indicates a cluster with defined specifications (a subscription or pay-as-you-go cluster).
	//
	// example:
	//
	// SteadyServerless
	ProxyServerlessType *string `json:"ProxyServerlessType,omitempty" xml:"ProxyServerlessType,omitempty"`
	// The number of CPU cores for the database proxy with standard specifications.
	//
	// example:
	//
	// 2
	ProxyStandardCpuCores *string `json:"ProxyStandardCpuCores,omitempty" xml:"ProxyStandardCpuCores,omitempty"`
	// The status of the database proxy. Valid values:
	//
	// - **Creating**: The proxy is being created.
	//
	// - **Running**: The proxy is running.
	//
	// - **Deleting**: The proxy is being released.
	//
	// - **Rebooting**: The proxy is being restarted.
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
	// rg-************
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
	// The method used to restore the cluster. Valid values:
	//
	// - **RestoreByTime**: The cluster is restored from a point in time based on a level-1 backup.
	//
	// - **RestoreByBackupSet**: The cluster is restored from a backup set based on a level-1 backup.
	//
	// - **RestoreByTimeOss**: The cluster is restored from a point in time based on a level-2 backup.
	//
	// - **RestoreByBackupSetOss**: The cluster is restored from a backup set based on a level-2 backup.
	//
	// - **CloneFromSourceCluster**: The cluster is cloned from a source cluster.
	//
	// > This parameter is supported only for clusters that are restored from a backup set or a point in time after June 1, 2024.
	//
	// example:
	//
	// RestoreByTime
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// RowCompression
	//
	// example:
	//
	// OFF
	RowCompression *string `json:"RowCompression,omitempty" xml:"RowCompression,omitempty"`
	// The storage usage for SQL. Unit: bytes. A value of -1 indicates that no data is available.
	//
	// example:
	//
	// 0
	SQLSize *int64 `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
	// The Serverless type. Valid values:
	//
	// - AgileServerless: Agile
	//
	// - SteadyServerless: Steady
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The source cluster ID.
	//
	// > This parameter is supported only for clusters that are restored from a backup set or a point in time after June 1, 2024.
	//
	// example:
	//
	// pc-pz51ziv48317b2880
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The ID of the region where the source cluster resides.
	//
	// > This parameter is returned only if a source cluster ID exists.
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
	// - **0**: Customer drill mode.
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
	// - **Postpaid**: Pay-as-you-go.
	//
	// - **Prepaid**: Subscription.
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage class of the Enterprise Edition. Valid values:
	//
	// - **PSL5**
	//
	// - **PSL4**
	//
	// The storage class of the Standard Edition. Valid values:
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
	// ESSDPL0
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The used storage space. Unit: bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether strong consistency is enabled for the multi-zone cluster. Valid values:
	//
	// - **ON**: Strong consistency is enabled for the multi-zone cluster. This applies to Standard Edition clusters that are deployed in three zones.
	//
	// - **OFF**: Strong consistency is not enabled for the multi-zone cluster.
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
	// Indicates whether the failover with hot replica feature that is compatible with the IMCI feature is supported.
	//
	// example:
	//
	// ON
	SupportInstantSwitchWithImci *string `json:"SupportInstantSwitchWithImci,omitempty" xml:"SupportInstantSwitchWithImci,omitempty"`
	// The details of the tags.
	Tags []*DescribeDbClusterAttributeZonalResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The virtual switch ID.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i,cn-hangzhou-g
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeDbClusterAttributeZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbClusterAttributeZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetAiCreatingTime() *string {
	return s.AiCreatingTime
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetAiType() *string {
	return s.AiType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetBlktagTotal() *int64 {
	return s.BlktagTotal
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetBlktagUsed() *int64 {
	return s.BlktagUsed
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetBurstingEnabled() *string {
	return s.BurstingEnabled
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetCompressStorageMode() *string {
	return s.CompressStorageMode
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetCompressStorageUsed() *int64 {
	return s.CompressStorageUsed
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBNodes() []*DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	return s.DBNodes
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDBVersionStatus() *string {
	return s.DBVersionStatus
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDataLevel1BackupChainSize() *int64 {
	return s.DataLevel1BackupChainSize
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDataSyncMode() *string {
	return s.DataSyncMode
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetDeletionLock() *int32 {
	return s.DeletionLock
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetHasCompleteStandbyRes() *bool {
	return s.HasCompleteStandbyRes
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetImciAutoIndex() *string {
	return s.ImciAutoIndex
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetImperceptibleSwitch() *string {
	return s.ImperceptibleSwitch
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetInodeTotal() *int64 {
	return s.InodeTotal
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetInodeUsed() *int64 {
	return s.InodeUsed
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetIsProxyLatestVersion() *bool {
	return s.IsProxyLatestVersion
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetOrca() *string {
	return s.Orca
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProvisionedIops() *string {
	return s.ProvisionedIops
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProxyCpuCores() *string {
	return s.ProxyCpuCores
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProxyServerlessType() *string {
	return s.ProxyServerlessType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProxyStandardCpuCores() *string {
	return s.ProxyStandardCpuCores
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetRestoreDataPoint() *string {
	return s.RestoreDataPoint
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetRowCompression() *string {
	return s.RowCompression
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetSQLSize() *int64 {
	return s.SQLSize
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetSourceDBCluster() *string {
	return s.SourceDBCluster
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStandbyHAMode() *string {
	return s.StandbyHAMode
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStorageMax() *int64 {
	return s.StorageMax
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStoragePayType() *string {
	return s.StoragePayType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetStrictConsistency() *string {
	return s.StrictConsistency
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetSubCategory() *string {
	return s.SubCategory
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetSupportInstantSwitchWithImci() *string {
	return s.SupportInstantSwitchWithImci
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetTags() []*DescribeDbClusterAttributeZonalResponseBodyTags {
	return s.Tags
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDbClusterAttributeZonalResponseBody) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetAiCreatingTime(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.AiCreatingTime = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetAiType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.AiType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetArchitecture(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.Architecture = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetAutoUpgradeMinorVersion(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetBlktagTotal(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.BlktagTotal = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetBlktagUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.BlktagUsed = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetBurstingEnabled(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetCategory(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetCompressStorageMode(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.CompressStorageMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetCompressStorageUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.CompressStorageUsed = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetCreationTime(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBClusterClass(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBClusterClass = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBClusterDescription(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBClusterId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBClusterNetworkType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBClusterStatus(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBNodes(v []*DescribeDbClusterAttributeZonalResponseBodyDBNodes) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBNodes = v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBVersion(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDBVersionStatus(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDataLevel1BackupChainSize(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.DataLevel1BackupChainSize = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDataSyncMode(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.DataSyncMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetDeletionLock(v int32) *DescribeDbClusterAttributeZonalResponseBody {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetEngine(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetExpireTime(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetExpired(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetHasCompleteStandbyRes(v bool) *DescribeDbClusterAttributeZonalResponseBody {
	s.HasCompleteStandbyRes = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetHotStandbyCluster(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.HotStandbyCluster = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetImciAutoIndex(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ImciAutoIndex = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetImperceptibleSwitch(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ImperceptibleSwitch = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetInodeTotal(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.InodeTotal = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetInodeUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.InodeUsed = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetIsLatestVersion(v bool) *DescribeDbClusterAttributeZonalResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetIsProxyLatestVersion(v bool) *DescribeDbClusterAttributeZonalResponseBody {
	s.IsProxyLatestVersion = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetLockMode(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetMaintainTime(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetOrca(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.Orca = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetPayType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProvisionedIops(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProxyCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProxyCpuCores = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProxyServerlessType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProxyServerlessType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProxyStandardCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProxyStandardCpuCores = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProxyStatus(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetProxyType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ProxyType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetRegionId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetRequestId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetResourceGroupId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetRestoreDataPoint(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.RestoreDataPoint = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetRestoreType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.RestoreType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetRowCompression(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.RowCompression = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetSQLSize(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.SQLSize = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetServerlessType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetSourceDBCluster(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.SourceDBCluster = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetSourceRegionId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStandbyHAMode(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.StandbyHAMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStorageMax(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.StorageMax = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStoragePayType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.StoragePayType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStorageSpace(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStorageType(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStorageUsed(v int64) *DescribeDbClusterAttributeZonalResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetStrictConsistency(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.StrictConsistency = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetSubCategory(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.SubCategory = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetSupportInstantSwitchWithImci(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.SupportInstantSwitchWithImci = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetTags(v []*DescribeDbClusterAttributeZonalResponseBodyTags) *DescribeDbClusterAttributeZonalResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetVPCId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetVSwitchId(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) SetZoneIds(v string) *DescribeDbClusterAttributeZonalResponseBody {
	s.ZoneIds = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBody) Validate() error {
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

type DescribeDbClusterAttributeZonalResponseBodyDBNodes struct {
	// The number of CPU cores for bursting.
	//
	// example:
	//
	// 6
	AddedCpuCores *string `json:"AddedCpuCores,omitempty" xml:"AddedCpuCores,omitempty"`
	// The number of CPU cores in the node.
	//
	// example:
	//
	// 2
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
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
	// The node role. Valid values:
	//
	// - **Writer**: The primary node.
	//
	// - **Reader**: The read-only node.
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The node status. Valid values:
	//
	// - **Creating**: The node is being created.
	//
	// - **Running**: The node is running.
	//
	// - **Deleting**: The node is being deleted.
	//
	// - **Rebooting**: The node is being restarted.
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
	// - **MinorVersionUpgrading**: The minor version is being updated.
	//
	// - **Maintaining**: The instance is under maintenance.
	//
	// - **Switching**: A failover is in progress.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The failover priority. Each node has a failover priority. A higher priority value indicates a higher probability that the node is elected as the primary node during a failover.
	//
	// Valid values: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// Indicates whether hot standby is enabled. Valid values:
	//
	// - **ON**: Enabled
	//
	// - **OFF**: Disabled
	//
	// example:
	//
	// ON
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	// Indicates whether the In-Memory Column Index (IMCI) feature is enabled. Valid values:
	//
	// - **ON**: Enabled
	//
	// - **OFF**: Disabled
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
	// The maximum number of input/output operations per second (IOPS).
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
	// The name of the hot replica compute node that corresponds to the node in an architecture where hot standby storage and hot standby compute are enabled.
	//
	// example:
	//
	// pi-bp18z52mirror*****
	MirrorInsName *string `json:"MirrorInsName,omitempty" xml:"MirrorInsName,omitempty"`
	// MultiMasterLocalStandby
	//
	// example:
	//
	// MultiMasterLocalStandby
	MultiMasterLocalStandby *string `json:"MultiMasterLocalStandby,omitempty" xml:"MultiMasterLocalStandby,omitempty"`
	// MultiMasterPrimaryNode
	//
	// example:
	//
	// MultiMasterPrimaryNode
	MultiMasterPrimaryNode *string `json:"MultiMasterPrimaryNode,omitempty" xml:"MultiMasterPrimaryNode,omitempty"`
	// The Orca feature. Valid values:
	//
	// - on: Enabled
	//
	// - off: Disabled
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
	// Indicates whether the global consistency (high-performance mode) feature is enabled for the node. Valid values:
	//
	// - **ON**: Enabled
	//
	// - **OFF**: Disabled
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
	// The Serverless type. Valid values:
	//
	// - AgileServerless: Agile
	//
	// - SteadyServerless: Steady
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Indicates whether the node is in the primary or secondary zone. This parameter is mainly used for resource peering.
	//
	// Valid values:
	//
	// - Primary: The primary zone
	//
	// - Standby: The secondary zone
	//
	// example:
	//
	// Primary
	SubCluster *string `json:"SubCluster,omitempty" xml:"SubCluster,omitempty"`
	// SubGroupDescription
	//
	// example:
	//
	// SubGroupDescription
	SubGroupDescription *string `json:"SubGroupDescription,omitempty" xml:"SubGroupDescription,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDbClusterAttributeZonalResponseBodyDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbClusterAttributeZonalResponseBodyDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetAddedCpuCores() *string {
	return s.AddedCpuCores
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetDBNodeDescription() *string {
	return s.DBNodeDescription
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetFailoverPriority() *int32 {
	return s.FailoverPriority
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetHotReplicaMode() *string {
	return s.HotReplicaMode
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetImciSwitch() *string {
	return s.ImciSwitch
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMasterId() *string {
	return s.MasterId
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMirrorInsName() *string {
	return s.MirrorInsName
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMultiMasterLocalStandby() *string {
	return s.MultiMasterLocalStandby
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetMultiMasterPrimaryNode() *string {
	return s.MultiMasterPrimaryNode
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetOrca() *string {
	return s.Orca
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetRemoteMemorySize() *string {
	return s.RemoteMemorySize
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetSccMode() *string {
	return s.SccMode
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetServerWeight() *string {
	return s.ServerWeight
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetSubCluster() *string {
	return s.SubCluster
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetSubGroupDescription() *string {
	return s.SubGroupDescription
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetAddedCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.AddedCpuCores = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetCpuCores(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.CpuCores = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetCreationTime(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetDBNodeClass(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetDBNodeDescription(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.DBNodeDescription = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetDBNodeId(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetDBNodeRole(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetDBNodeStatus(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetFailoverPriority(v int32) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.FailoverPriority = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetHotReplicaMode(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.HotReplicaMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetImciSwitch(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.ImciSwitch = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMasterId(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MasterId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMaxConnections(v int32) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMaxIOPS(v int32) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMemorySize(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMirrorInsName(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MirrorInsName = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMultiMasterLocalStandby(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MultiMasterLocalStandby = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetMultiMasterPrimaryNode(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.MultiMasterPrimaryNode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetOrca(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.Orca = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetRemoteMemorySize(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.RemoteMemorySize = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetSccMode(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.SccMode = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetServerWeight(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.ServerWeight = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetServerlessType(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetSubCluster(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.SubCluster = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetSubGroupDescription(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.SubGroupDescription = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) SetZoneId(v string) *DescribeDbClusterAttributeZonalResponseBodyDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDbClusterAttributeZonalResponseBodyTags struct {
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

func (s DescribeDbClusterAttributeZonalResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbClusterAttributeZonalResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDbClusterAttributeZonalResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDbClusterAttributeZonalResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDbClusterAttributeZonalResponseBodyTags) SetKey(v string) *DescribeDbClusterAttributeZonalResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyTags) SetValue(v string) *DescribeDbClusterAttributeZonalResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeDbClusterAttributeZonalResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

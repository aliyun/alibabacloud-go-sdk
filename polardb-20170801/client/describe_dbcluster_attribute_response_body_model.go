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
	SetBurstingEnabled(v string) *DescribeDBClusterAttributeResponseBody
	GetBurstingEnabled() *string
	SetCategory(v string) *DescribeDBClusterAttributeResponseBody
	GetCategory() *string
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
	// Start time for free AI activation
	//
	// example:
	//
	// 2024-03-13T01:20:28Z
	AiCreatingTime *string `json:"AiCreatingTime,omitempty" xml:"AiCreatingTime,omitempty"`
	// Types of AI nodes. Values include:
	//
	// - **SearchNode**: Search node.
	//
	// - **DLNode**: AI node.
	//
	// example:
	//
	// DLNode
	AiType *string `json:"AiType,omitempty" xml:"AiType,omitempty"`
	// CPU architecture. Available options are:
	//
	// - **X86**
	//
	// - **ARM**
	//
	// example:
	//
	// X86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The minor version upgrade method.
	//
	// 	- Auto
	//
	// 	- Manual
	//
	// example:
	//
	// Manual
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// Maximum number of blktags in the file system.
	//
	// example:
	//
	// 7,864,320
	BlktagTotal *int64 `json:"BlktagTotal,omitempty" xml:"BlktagTotal,omitempty"`
	// Current blktag usage.
	//
	// example:
	//
	// 5,242,880
	BlktagUsed *int64 `json:"BlktagUsed,omitempty" xml:"BlktagUsed,omitempty"`
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// [Product Series](https://help.aliyun.com/document_detail/183258.html), with values as follows:
	//
	// 	- **Normal**: Cluster Edition
	//
	// 	- **Basic**: Single Node
	//
	// 	- **Archive**: High Compression Engine (X-Engine)
	//
	// 	- **NormalMultimaster**: Multi-Master Cluster Edition
	//
	// 	- **SENormal**: Standard Edition
	//
	// > 	- PolarDB PostgreSQL version 11 does not support single-node.
	//
	// >	- PolarDB MySQL versions 8.0 and 5.7, and PolarDB PostgreSQL version 14 support the Standard Edition.
	//
	// >	- PolarDB MySQL version 8.0 supports High Compression Engine (X-Engine) and Multi-Master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Whether storage compression is enabled. Values are as follows:
	//
	// - ON: Enabled
	//
	// - OFF: Disabled
	//
	// example:
	//
	// ON
	CompressStorageMode *string `json:"CompressStorageMode,omitempty" xml:"CompressStorageMode,omitempty"`
	// Compressed storage data size.
	//
	// > This parameter is supported only when the cluster\\"s storage compression feature is enabled.
	//
	// example:
	//
	// 15529410560
	CompressStorageUsed *int64 `json:"CompressStorageUsed,omitempty" xml:"CompressStorageUsed,omitempty"`
	// Cluster creation time.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Cluster description.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// Cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Network type of the cluster.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// Cluster status. For the full list of values, refer to [Cluster Status Table](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The information about the nodes.
	DBNodes []*DescribeDBClusterAttributeResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// Database engine type.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// Database engine version.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The status of the minor version. Valid values:
	//
	// 	- **Stable**: The minor version is stable.
	//
	// 	- **Old**: The minor version is outdated. We recommend that you update it to the latest version.
	//
	// 	- **HighRisk**: The minor version has critical defects. We recommend that you immediately update it to the latest version.
	//
	// 	- **Beta**: The minor version is a Beta version.
	//
	// >  For information about how to update the minor version, see [Minor version update](https://help.aliyun.com/document_detail/158572.html).
	//
	// example:
	//
	// Stable
	DBVersionStatus *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	// Total size of Level 1 backups (snapshots), in bytes.
	//
	// example:
	//
	// 74448896
	DataLevel1BackupChainSize *int64 `json:"DataLevel1BackupChainSize,omitempty" xml:"DataLevel1BackupChainSize,omitempty"`
	// Data replication relationship mode. Values are as follows:
	//
	// - **AsyncSync**: Asynchronous
	//
	// - **SemiSync**: Semi-synchronous
	//
	// example:
	//
	// AsyncSync
	DataSyncMode *string `json:"DataSyncMode,omitempty" xml:"DataSyncMode,omitempty"`
	// Lock status for cluster deletion, with values as follows:
	//
	// 	- **0**: Unlocked, cluster can be deleted.
	//
	// 	- **1**: Locked, cluster cannot be deleted.
	//
	// example:
	//
	// 0
	DeletionLock *int32 `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	// Cluster engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Cluster expiration time.
	//
	// > Only clusters with **Prepaid*	- (subscription) payment methods return specific parameter values; **Postpaid*	- (pay-as-you-go) clusters return empty values.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Whether the cluster has expired.
	//
	// > This parameter is only supported for clusters with **Prepaid*	- (Subscription) payment methods.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Whether to replenish resources for the new primary after cross-AZ switch. Values are as follows:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	HasCompleteStandbyRes *bool `json:"HasCompleteStandbyRes,omitempty" xml:"HasCompleteStandbyRes,omitempty"`
	// Whether to enable storage hot backup cluster (and Standby compute nodes). Values are as follows:
	//
	// - **StandbyClusterON**: Enable storage hot backup/Enable storage hot backup and Standby compute nodes.
	//
	// - **StandbyClusterOFF**: Disable storage hot backup/Disable storage hot backup and Standby compute nodes.
	//
	// example:
	//
	// StandbyClusterON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// Indicates whether the automatic IMCI-based query acceleration feature is enabled. Valid values:
	//
	// 	- `ON`: enabled
	//
	// 	- `OFF`: disabled
	//
	// example:
	//
	// OFF
	ImciAutoIndex *string `json:"ImciAutoIndex,omitempty" xml:"ImciAutoIndex,omitempty"`
	// Indicates whether failover with hot replica is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false` (default)
	ImperceptibleSwitch *string `json:"ImperceptibleSwitch,omitempty" xml:"ImperceptibleSwitch,omitempty"`
	// Maximum number of inodes in the file system.
	//
	// example:
	//
	// 6,291,456
	InodeTotal *int64 `json:"InodeTotal,omitempty" xml:"InodeTotal,omitempty"`
	// Current inode usage.
	//
	// example:
	//
	// 4,194,304
	InodeUsed *int64 `json:"InodeUsed,omitempty" xml:"InodeUsed,omitempty"`
	// Indicates whether it is the latest kernel version. Values are as follows:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether it is the latest version of the database proxy, with possible values as follows:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	IsProxyLatestVersion *bool `json:"IsProxyLatestVersion,omitempty" xml:"IsProxyLatestVersion,omitempty"`
	// Lock mode. Possible values are as follows:
	//
	// - **Unlock**: Unlocked.
	//
	// - **ManualLock**: Manually triggered lock.
	//
	// - **LockByExpiration**: Automatic cluster lock upon expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance window for the cluster, formatted as `HH:mmZ-HH:mmZ` (UTC time). For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 0:00 to 1:00 (UTC+08:00).
	//
	// example:
	//
	// 18:00Z-19:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// Orca function with possible values as follows:
	//
	// - **on**: Enabled
	//
	// - **off**: Disabled
	//
	// example:
	//
	// ON
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// Payment type. Possible values are:
	//
	// - **Postpaid**: Pay-As-You-Go
	//
	// - **Prepaid**: Prepaid (Subscription).
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Describes the preconfigured read and write IOPS for ESSD AutoPL cloud disks. Possible values: 0 to min{50,000, 1000*capacity - baseline performance}.<br>Baseline performance = min{1,800 + 50*capacity, 50000}.<br	Note: This parameter is supported only when StorageType is ESSDAUTOPL.
	//
	// example:
	//
	// 2500
	ProvisionedIops *string `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// Number of CPU cores for the database proxy.
	//
	// example:
	//
	// 4
	ProxyCpuCores *string `json:"ProxyCpuCores,omitempty" xml:"ProxyCpuCores,omitempty"`
	// Serverless type for the database proxy. Currently, the value is fixed to AgileServerless.
	//
	// example:
	//
	// AgileServerless
	ProxyServerlessType *string `json:"ProxyServerlessType,omitempty" xml:"ProxyServerlessType,omitempty"`
	// Standard configuration CPU cores for the database proxy.
	//
	// example:
	//
	// 2
	ProxyStandardCpuCores *string `json:"ProxyStandardCpuCores,omitempty" xml:"ProxyStandardCpuCores,omitempty"`
	// Status of the database proxy. Possible values include:
	//
	// - **Creating**: Creating
	//
	// - **Running**: Running
	//
	// - **Deleting**: Releasing
	//
	// - **Rebooting**: Restarting
	//
	// - **DBNodeCreating**: Adding nodes
	//
	// - **DBNodeDeleting**: Deleting nodes
	//
	// - **ClassChanging**: Changing node specifications
	//
	// - **NetAddressCreating**: Creating network connections
	//
	// - **NetAddressDeleting**: Deleting network connections
	//
	// - **NetAddressModifying**: Modifying network connections
	//
	// - **Deleted**: Released
	//
	// example:
	//
	// Running
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// Database proxy types, with the following values:
	//
	// - **Exclusive**: Enterprise Exclusive Edition
	//
	// - **General**: Enterprise General Purpose Edition
	//
	// example:
	//
	// Exclusive
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 074467EF-86B9-4C23-ACBF-E9B81A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-***************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// If RestoreType is **RestoreByTime*	- or **RestoreByTimeOss**, this value represents the recovery time point. If RestoreType is **RestoreByBackupSet*	- or **RestoreByBackupSetOss**, this value indicates the ID of the backup set on which the recovery is based.
	//
	// <note>Only clusters restored from a backup set or time point after June 1, 2024, support this parameter.</note>
	//
	// example:
	//
	// 2179639137
	RestoreDataPoint *string `json:"RestoreDataPoint,omitempty" xml:"RestoreDataPoint,omitempty"`
	// Cluster recovery method, with possible values:
	//
	// 	- **RestoreByTime**: Restore from a time point based on primary backup. 	- **RestoreByBackupSet**: Restore from a backup set based on primary backup. 	- **RestoreByTimeOss**: Restore from a time point based on secondary backup. 	- **RestoreByBackupSetOss**: Restore from a backup set based on secondary backup. 	- **CloneFromSourceCluster**: Clone from the source cluster.
	//
	// <note>This parameter is only supported for clusters restored from a backup set or time point after June 1, 2024.</note>
	//
	// example:
	//
	// RestoreByTime
	RestoreType    *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	RowCompression *string `json:"RowCompression,omitempty" xml:"RowCompression,omitempty"`
	// Storage amount of SQL, in bytes. If the value is -1, it indicates no data.
	//
	// example:
	//
	// 0
	SQLSize *int64 `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
	// Serverless type. Valid values are as follows:
	//
	// - AgileServerless: Agile - SteadyServerless: Stable
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Source cluster ID. <note>Clusters restored from backup sets or specific points in time after June 1, 2024, support this parameter.</note>
	//
	// example:
	//
	// pc-pz51ziv48317b2880
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The region ID of the source cluster.
	//
	// >  This parameter is returned only if the source cluster ID exists.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// Cross-AZ disaster recovery mode. Values are as follows:
	//
	// - **ON**: Enable cross-AZ disaster recovery mode.
	//
	// - **OFF**: Disable cross-AZ disaster recovery mode.
	//
	// - **0**: Customer drill mode.
	//
	// example:
	//
	// OFF
	StandbyHAMode *string `json:"StandbyHAMode,omitempty" xml:"StandbyHAMode,omitempty"`
	// The maximum storage capacity of the current cluster specification, in bytes.
	//
	// example:
	//
	// 10995116277760
	StorageMax *int64 `json:"StorageMax,omitempty" xml:"StorageMax,omitempty"`
	// Storage billing type. Valid values are as follows:
	//
	// - **Postpaid**: Pay-as-you-go (by capacity).
	//
	// - **Prepaid**: Subscription (by space).
	//
	// example:
	//
	// Prepaid
	StoragePayType *string `json:"StoragePayType,omitempty" xml:"StoragePayType,omitempty"`
	// Storage space for pay-by-space (subscription) billing. Unit: Byte.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// Storage type, with a fixed value of **HighPerformance**.
	//
	// example:
	//
	// HighPerformance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Amount of used storage space, in bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether multi-AZ data strong consistency is enabled for the cluster. The value ranges are as follows:
	//
	// - **ON**: Indicates that multi-AZ data strong consistency is enabled, applicable to the Standard 3AZ scenario.
	//
	// - **OFF**: Indicates that multi-AZ data strong consistency is not enabled.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// Specification type of compute nodes, with possible values as follows:
	//
	// 	- **Exclusive**: Dedicated specification
	//
	// 	- **General**: General-purpose specification
	//
	// > This parameter is supported only for PolarDB MySQL Edition with the product series set to Cluster Edition.
	//
	// example:
	//
	// Exclusive
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	// Indicates whether queries based on In-Memory Column Indexes (IMCIs) are supported during and after a failover with hot replica.
	//
	// example:
	//
	// ON
	SupportInstantSwitchWithImci *string `json:"SupportInstantSwitchWithImci,omitempty" xml:"SupportInstantSwitchWithImci,omitempty"`
	// Details of tags.
	Tags []*DescribeDBClusterAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// VPC ID.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// VSwitch ID.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Availability Zone IDs.
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

func (s *DescribeDBClusterAttributeResponseBody) GetBurstingEnabled() *string {
	return s.BurstingEnabled
}

func (s *DescribeDBClusterAttributeResponseBody) GetCategory() *string {
	return s.Category
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

func (s *DescribeDBClusterAttributeResponseBody) SetBurstingEnabled(v string) *DescribeDBClusterAttributeResponseBody {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCategory(v string) *DescribeDBClusterAttributeResponseBody {
	s.Category = &v
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
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyDBNodes struct {
	// Number of CPU cores for second-level elastic scaling.
	//
	// example:
	//
	// 6
	AddedCpuCores *string `json:"AddedCpuCores,omitempty" xml:"AddedCpuCores,omitempty"`
	// Number of CPU cores for the node.
	//
	// example:
	//
	// 2
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// Node creation time.
	//
	// example:
	//
	// 2020-03-23T21:35:43Z
	CreationTime          *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBNodeCXLRemoteMemory *string `json:"DBNodeCXLRemoteMemory,omitempty" xml:"DBNodeCXLRemoteMemory,omitempty"`
	// Node specification.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// test
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// Node ID.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// Node role, with possible values as follows:
	//
	// - **Writer**: Primary node.
	//
	// - **Reader**: Read-only node.
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// Node status, with possible values as follows:
	//
	// 	- **Creating**: Creating
	//
	// 	- **Running**: Running
	//
	// 	- **Deleting**: Deleting
	//
	// 	- **Rebooting**: Rebooting
	//
	// 	- **DBNodeCreating**: Adding node
	//
	// 	- **DBNodeDeleting**: Removing node
	//
	// 	- **ClassChanging**: Modifying node specification
	//
	// 	- **NetAddressCreating**: Creating network connection
	//
	// 	- **NetAddressDeleting**: Deleting network connection
	//
	// 	- **NetAddressModifying**: Modifying network connection
	//
	// 	- **MinorVersionUpgrading**: Upgrading minor version
	//
	// 	- **Maintaining**: Instance maintenance
	//
	// 	- **Switching**: Switching
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// Failover priority. Each node has a failover priority, determining the likelihood of being elected as the primary node during a failover. A higher value indicates a higher priority.
	//
	// Range: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// Whether hot standby is enabled. Possible values are:
	//
	// - **ON**: Enabled
	//
	// - **OFF**: Disabled
	//
	// example:
	//
	// ON
	HotReplicaMode *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	// Whether columnar index is enabled. Possible values are:
	//
	// - **ON**: Enabled
	//
	// - **OFF**: Disabled
	//
	// example:
	//
	// ON
	ImciSwitch *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	// Primary node ID of the multi-master architecture cluster edition.
	//
	// example:
	//
	// pi-bp18z52akld3*****
	MasterId *string `json:"MasterId,omitempty" xml:"MasterId,omitempty"`
	// Maximum concurrent connections of the cluster.
	//
	// example:
	//
	// 8000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// Maximum number of I/O requests, that is, IOPS.
	//
	// example:
	//
	// 32000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// Node memory size, in MB.
	//
	// example:
	//
	// 8192
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The name of the hot standby compute node corresponding to the node when the hot standby storage and compute clusters feature is enabled.
	//
	// example:
	//
	// pi-bp18z52mirror*****
	MirrorInsName           *string `json:"MirrorInsName,omitempty" xml:"MirrorInsName,omitempty"`
	MultiMasterLocalStandby *string `json:"MultiMasterLocalStandby,omitempty" xml:"MultiMasterLocalStandby,omitempty"`
	MultiMasterPrimaryNode  *string `json:"MultiMasterPrimaryNode,omitempty" xml:"MultiMasterPrimaryNode,omitempty"`
	// Orca feature, valid values are:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// off
	Orca *string `json:"Orca,omitempty" xml:"Orca,omitempty"`
	// Remote memory size, in MB.
	//
	// example:
	//
	// 3072
	RemoteMemorySize *string `json:"RemoteMemorySize,omitempty" xml:"RemoteMemorySize,omitempty"`
	// Whether the node has the global consistency (high-performance mode) feature enabled. Possible values are:
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
	// Routing weight.
	//
	// Range: 1~100. Default is 1.
	//
	// example:
	//
	// 1
	ServerWeight *string `json:"ServerWeight,omitempty" xml:"ServerWeight,omitempty"`
	// Serverless type. Possible values include:
	//
	// - **AgileServerless**: Agile
	//
	// - **SteadyServerless**: Steady
	//
	// > This parameter is only supported by Serverless clusters.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// Identifies whether the node is in the primary or standby availability zone, primarily used in resource mirroring scenarios.
	//
	// Values include:
	//
	// - **Primary**: Primary Availability Zone
	//
	// - **Standby**: Standby Availability Zone
	//
	// example:
	//
	// Primary
	SubCluster          *string `json:"SubCluster,omitempty" xml:"SubCluster,omitempty"`
	SubGroupDescription *string `json:"SubGroupDescription,omitempty" xml:"SubGroupDescription,omitempty"`
	// Availability zone ID.
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
	// Tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value.
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

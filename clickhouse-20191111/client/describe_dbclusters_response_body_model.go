// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusters(v *DescribeDBClustersResponseBodyDBClusters) *DescribeDBClustersResponseBody
	GetDBClusters() *DescribeDBClustersResponseBodyDBClusters
	SetPageNumber(v int32) *DescribeDBClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBClustersResponseBody
	GetTotalCount() *int32
}

type DescribeDBClustersResponseBody struct {
	// The clusters.
	DBClusters *DescribeDBClustersResponseBodyDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) GetDBClusters() *DescribeDBClustersResponseBodyDBClusters {
	return s.DBClusters
}

func (s *DescribeDBClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBClustersResponseBody) SetDBClusters(v *DescribeDBClustersResponseBodyDBClusters) *DescribeDBClustersResponseBody {
	s.DBClusters = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) Validate() error {
	if s.DBClusters != nil {
		if err := s.DBClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyDBClusters struct {
	DBCluster []*DescribeDBClustersResponseBodyDBClustersDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClusters) GetDBCluster() []*DescribeDBClustersResponseBodyDBClustersDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClustersResponseBodyDBClusters) SetDBCluster(v []*DescribeDBClustersResponseBodyDBClustersDBCluster) *DescribeDBClustersResponseBodyDBClusters {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClusters) Validate() error {
	if s.DBCluster != nil {
		for _, item := range s.DBCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyDBClustersDBCluster struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The site ID. Valid values:
	//
	// 	- **26842**: the China site (aliyun.com)
	//
	// 	- **26888**: the international site (alibabacloud.com)
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// example:
	//
	// Basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the cluster.
	//
	// example:
	//
	// clickhouse_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The VPC endpoint of the cluster.
	//
	// example:
	//
	// cc-bp1fs5o051c61****.clickhouse.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The version number of the backend management system of ApsaraDB for ClickHouse. Valid values:
	//
	// 	- **v1**
	//
	// 	- **v2**
	//
	// example:
	//
	// v1
	ControlVersion *string `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-10-28T07:24:45Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: -**S4**: 4 CPU cores and 16 GB of memory -**S8**: 8 CPU cores and 32 GB of memory
	//
	//     	- **S16**: 16 CPU cores and 64 GB of memory
	//
	//     	- **S32**: 32 CPU cores and 128 GB of memory
	//
	//     	- **S64**: 64 CPU cores and 256 GB of memory
	//
	//     	- **S104**: 104 CPU cores and 384 GB of memory
	//
	// 	- Valid values when the cluster is of Double-replica Edition: -**C4**: 4 CPU cores and 16 GB of memory -**C8**: 8 CPU cores and 32 GB of memory -**C16**: 16 CPU cores and 64 GB of memory -**C32**: 32 CPU cores and 128 GB of memory -**C64**: 64 CPU cores and 256 GB of memory -**C104**: 104 CPU cores and 384 GB of memory
	//
	// example:
	//
	// C8
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// example:
	//
	// 2
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of each node. Valid values: 100 to 32000. Unit: GB.
	//
	// >  This value is a multiple of 100.
	//
	// example:
	//
	// 100
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The engine version of the cluster.
	//
	// example:
	//
	// 23.8
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// The time when the cluster expired. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2011-05-30T12:11:4Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The extended storage space.
	//
	// example:
	//
	// 100GB
	ExtStorageSize *int32 `json:"ExtStorageSize,omitempty" xml:"ExtStorageSize,omitempty"`
	// The extended storage type. Valid values:
	//
	// 	- **CloudSSD**: standard SSD.
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	ExtStorageType *string `json:"ExtStorageType,omitempty" xml:"ExtStorageType,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// 	- **true**: The cluster has expired.
	//
	// 	- **false**: The cluster has not expired.
	//
	// example:
	//
	// false
	IsExpired *string `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// 	- **LockByRestoration**: The cluster is automatically locked because the cluster is about to be rolled back.
	//
	// 	- **LockByDiskQuota**: The cluster is automatically locked because the disk space is exhausted.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The cause why the cluster was locked.
	//
	// >  If the value of the LockMode parameter is Unlock, an empty string is returned for this parameter.
	//
	// example:
	//
	// DISK_FULL
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: The cluster uses the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: The cluster uses the subscription billing method.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time window during which write operations are stopped for specification changes.
	//
	// example:
	//
	// 2025-02-08T00:00:00+08:00,2025-02-12T00:00:00+08:00
	ScaleOutDisableWriteWindows *string `json:"ScaleOutDisableWriteWindows,omitempty" xml:"ScaleOutDisableWriteWindows,omitempty"`
	// The status of a data migration task.
	ScaleOutStatus *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags.
	Tags *DescribeDBClustersResponseBodyDBClustersDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetControlVersion() *string {
	return s.ControlVersion
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeCount() *int64 {
	return s.DBNodeCount
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeStorage() *int64 {
	return s.DBNodeStorage
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExtStorageSize() *int32 {
	return s.ExtStorageSize
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExtStorageType() *string {
	return s.ExtStorageType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetIsExpired() *string {
	return s.IsExpired
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetScaleOutDisableWriteWindows() *string {
	return s.ScaleOutDisableWriteWindows
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetScaleOutStatus() *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	return s.ScaleOutStatus
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetTags() *DescribeDBClustersResponseBodyDBClustersDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetAliUid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.AliUid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetBid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Bid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetControlVersion(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ControlVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDbVersion(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DbVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageSize(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetIsExpired(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.IsExpired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPort(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetScaleOutDisableWriteWindows(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ScaleOutDisableWriteWindows = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetScaleOutStatus(v *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ScaleOutStatus = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetTags(v *DescribeDBClustersResponseBodyDBClustersDBClusterTags) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcCloudInstanceId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) Validate() error {
	if s.ScaleOutStatus != nil {
		if err := s.ScaleOutStatus.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus struct {
	// The progress of the data migration task in percentage.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress of the data migration task. This value is displayed in the following format: Data volume that has been migrated/Total data volume.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0MB/60469MB
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetProgress(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetRatio(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Ratio = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyDBClustersDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) GetTag() []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) *DescribeDBClustersResponseBodyDBClustersDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) Validate() error {
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

type DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag struct {
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}

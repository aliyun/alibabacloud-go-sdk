// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersWithBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBClustersWithBackupsResponseBodyItems) *DescribeDBClustersWithBackupsResponseBody
	GetItems() *DescribeDBClustersWithBackupsResponseBodyItems
	SetPageNumber(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBClustersWithBackupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBClustersWithBackupsResponseBody struct {
	// The details about the cluster.
	Items *DescribeDBClustersWithBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F8529AA2-522F-4B30-B80B-8F7D39******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetItems() *DescribeDBClustersWithBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetItems(v *DescribeDBClustersWithBackupsResponseBodyItems) *DescribeDBClustersWithBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageNumber(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetRequestId(v string) *DescribeDBClustersWithBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersWithBackupsResponseBodyItems struct {
	DBCluster []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) GetDBCluster() []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) SetDBCluster(v []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) *DescribeDBClustersWithBackupsResponseBodyItems {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersWithBackupsResponseBodyItemsDBCluster struct {
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-05-09T09:33:51Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- Creating: The cluster is being created.
	//
	// 	- Running: The cluster is running.
	//
	// 	- Deleting: The cluster is being released.
	//
	// 	- Rebooting: The cluster is restarting.
	//
	// 	- DBNodeCreating: The node is being added.
	//
	// 	- DBNodeDeleting: The node is being deleted.
	//
	// 	- ClassChanging: The specifications of the node are being changed.
	//
	// 	- NetAddressCreating: The network connection is being created.
	//
	// 	- NetAddressDeleting: The network connection is being deleted.
	//
	// 	- NetAddressModifying: The network connection is being modified.
	//
	// 	- Deleted: The cluster has been released.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The specifications of the node.
	//
	// example:
	//
	// polar.mysql.x4.medium
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The time when the cluster was deleted.
	//
	// example:
	//
	// 2022-05-12T03:25:37Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// Indicates whether the cluster is locked and can be deleted. Valid values:
	//
	// 	- **0**: The cluster is not locked and can be deleted.
	//
	// 	- **1**: The cluster is locked and cannot be deleted.
	//
	// example:
	//
	// 0
	DeletionLock *int32 `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The time when the cluster expires.
	//
	// > A specific value will be returned only for subscription clusters. For pay-as-you-go clusters, an empty string will be returned.
	//
	// example:
	//
	// 2022-09-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired.
	//
	// > A specific value will be returned only for subscription clusters.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates whether the cluster was released. Valid values:
	//
	// 	- 1: released
	//
	// 	- 0: not released
	//
	// example:
	//
	// 1
	IsDeleted *int32 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The state of the cluster lock. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked after the cluster expires.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID of the cluster.
	//
	// example:
	//
	// vpc-******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the instance is located.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDeletionLock() *int32 {
	return s.DeletionLock
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletedTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletedTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletionLock(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetIsDeleted(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) Validate() error {
	return dara.Validate(s)
}

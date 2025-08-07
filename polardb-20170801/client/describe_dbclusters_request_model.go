// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *DescribeDBClustersRequest
	GetConnectionString() *string
	SetDBClusterDescription(v string) *DescribeDBClustersRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeDBClustersRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeDBClustersRequest
	GetDBClusterStatus() *string
	SetDBNodeIds(v string) *DescribeDBClustersRequest
	GetDBNodeIds() *string
	SetDBType(v string) *DescribeDBClustersRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClustersRequest
	GetDBVersion() *string
	SetDescribeType(v string) *DescribeDBClustersRequest
	GetDescribeType() *string
	SetExpired(v bool) *DescribeDBClustersRequest
	GetExpired() *bool
	SetOwnerAccount(v string) *DescribeDBClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClustersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeDBClustersRequest
	GetPayType() *string
	SetRecentCreationInterval(v int32) *DescribeDBClustersRequest
	GetRecentCreationInterval() *int32
	SetRecentExpirationInterval(v int32) *DescribeDBClustersRequest
	GetRecentExpirationInterval() *int32
	SetRegionId(v string) *DescribeDBClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClustersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClustersRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest
	GetTag() []*DescribeDBClustersRequestTag
}

type DescribeDBClustersRequest struct {
	// The endpoint of the cluster.
	//
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The description of the cluster. Fuzzy match is supported.
	//
	// example:
	//
	// pc-****************
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the cluster. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// pc-****************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The state of the cluster that you want to query. For information about valid values, see [Cluster states](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The ID of the node. You can specify multiple node IDs. Separate multiple node IDs with commas (,).
	//
	// example:
	//
	// pi-***************
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// The database engine that the cluster runs. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version of the cluster.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The query mode of the list. The value Simple indicates that the simple mode is used. In this mode, only the basic metadata information of the cluster is returned.
	//
	// > If you do not specify this parameter, the detailed mode is used by default. Detailed information about the cluster is returned.
	//
	// example:
	//
	// Simple
	DescribeType *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	// Specifies whether the cluster has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Filters clusters created in the last N days. Valid values: 0 to 15.
	//
	// example:
	//
	// 7
	RecentCreationInterval *int32 `json:"RecentCreationInterval,omitempty" xml:"RecentCreationInterval,omitempty"`
	// Filters clusters that expire after N days. Valid values: 0 to 15.
	//
	// example:
	//
	// 6
	RecentExpirationInterval *int32 `json:"RecentExpirationInterval,omitempty" xml:"RecentExpirationInterval,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available regions.
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
	// rg-**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the cluster.
	Tag []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClustersRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeDBClustersRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *DescribeDBClustersRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClustersRequest) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeDBClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersRequest) GetRecentCreationInterval() *int32 {
	return s.RecentCreationInterval
}

func (s *DescribeDBClustersRequest) GetRecentExpirationInterval() *int32 {
	return s.RecentExpirationInterval
}

func (s *DescribeDBClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClustersRequest) GetTag() []*DescribeDBClustersRequestTag {
	return s.Tag
}

func (s *DescribeDBClustersRequest) SetConnectionString(v string) *DescribeDBClustersRequest {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBNodeIds(v string) *DescribeDBClustersRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBType(v string) *DescribeDBClustersRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBVersion(v string) *DescribeDBClustersRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDescribeType(v string) *DescribeDBClustersRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClustersRequest) SetExpired(v bool) *DescribeDBClustersRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPayType(v string) *DescribeDBClustersRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRecentCreationInterval(v int32) *DescribeDBClustersRequest {
	s.RecentCreationInterval = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRecentExpirationInterval(v int32) *DescribeDBClustersRequest {
	s.RecentExpirationInterval = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersRequestTag struct {
	// The key of the tag. You can use tags to filter clusters. You can specify up to 20 tags. N specifies the serial number of each tag. The values that you specify for N must be unique and consecutive integers that start from 1. The value of Tag.N.Key is Tag.N.Value.
	//
	// > The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// MySQL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > The tag value can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// 5.6
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersRequestTag) Validate() error {
	return dara.Validate(s)
}

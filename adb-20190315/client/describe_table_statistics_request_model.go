// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTableStatisticsRequest
	GetDBClusterId() *string
	SetKeyword(v string) *DescribeTableStatisticsRequest
	GetKeyword() *string
	SetOrder(v string) *DescribeTableStatisticsRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeTableStatisticsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTableStatisticsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTableStatisticsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTableStatisticsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTableStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTableStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTableStatisticsRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeTableStatisticsRequest
	GetSchemaName() *string
}

type DescribeTableStatisticsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query a list of cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The keyword that is used to query information by table name.
	//
	// example:
	//
	// you_table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The order in which to sort the queried information. Specify this parameter as an ordered JSON array that consists of the `Field` and `Type` fields. Example: `[{ "Field":"TableName", "Type":"Asc" }]`.
	//
	// 	- `Field` specifies the field that is used to sort the queried information. The following fields are supported: `TableName`, ColdDataSize, DataSize, PrimaryKeyIndexSize, RowCount, IndexSize, SchemaName, and PartitionCount.
	//
	// 	- `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     	- **Desc**: descending order.
	//
	//     	- **Asc**: ascending order.
	//
	// example:
	//
	// [ { "Field":"TableName", "Type":"Asc" } ]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
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
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTableStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTableStatisticsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeTableStatisticsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeTableStatisticsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTableStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTableStatisticsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTableStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTableStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTableStatisticsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTableStatisticsRequest) SetDBClusterId(v string) *DescribeTableStatisticsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetKeyword(v string) *DescribeTableStatisticsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOrder(v string) *DescribeTableStatisticsRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageNumber(v int32) *DescribeTableStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageSize(v int32) *DescribeTableStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetRegionId(v string) *DescribeTableStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetSchemaName(v string) *DescribeTableStatisticsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableAccessCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTableAccessCountRequest
	GetDBClusterId() *string
	SetOrder(v string) *DescribeTableAccessCountRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeTableAccessCountRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTableAccessCountRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTableAccessCountRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeTableAccessCountRequest
	GetStartTime() *string
	SetTableName(v string) *DescribeTableAccessCountRequest
	GetTableName() *string
}

type DescribeTableAccessCountRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of all AnalyticDB for MySQL clusters within a specified region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"TableSchema","Type":"Asc"}]`.
	//
	// 	- `Field` indicates the field that is used to sort the retrieved entries. Valid values:
	//
	//     	- `TableSchema`: the name of the database to which the table belongs.
	//
	//     	- `TableName`: the name of the table.
	//
	//     	- `AccessCount`: the number of accesses to the table.
	//
	// 	- `Type` indicates the sorting method. Valid values:
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// >  If this parameter is not specified, query results are sorted in ascending order of the database to which a specific table belongs.
	//
	// example:
	//
	// [{"Field":"TableSchema","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The value must be a positive integer. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The date to query. Specify the time in the *yyyy-MM-dd	- format. The time must be in UTC.
	//
	// >  Only data for the last 30 days can be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-08-30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the specific table.
	//
	// >  If this parameter is not specified, the number of accesses to all tables within the specified cluster for a specified date is returned.
	//
	// example:
	//
	// CUSTOMER
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableAccessCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableAccessCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTableAccessCountRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeTableAccessCountRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTableAccessCountRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableAccessCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableAccessCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTableAccessCountRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableAccessCountRequest) SetDBClusterId(v string) *DescribeTableAccessCountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetOrder(v string) *DescribeTableAccessCountRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageNumber(v int32) *DescribeTableAccessCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageSize(v int32) *DescribeTableAccessCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetRegionId(v string) *DescribeTableAccessCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetStartTime(v string) *DescribeTableAccessCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetTableName(v string) *DescribeTableAccessCountRequest {
	s.TableName = &v
	return s
}

func (s *DescribeTableAccessCountRequest) Validate() error {
	return dara.Validate(s)
}

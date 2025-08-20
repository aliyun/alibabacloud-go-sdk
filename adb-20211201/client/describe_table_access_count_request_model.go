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
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2ze627uzpkh8a8****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"TableSchema","Type":"Asc"}]`. Fields in the request parameter:
	//
	// 	- `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     	- `TableSchema`: the name of the database to which the table belongs.
	//
	//     	- `TableName`: the name of the table.
	//
	//     	- `AccessCount`: the number of accesses to the table.
	//
	// 	- `Type` specifies the sorting order. Valid values:
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// >  If you do not specify this parameter, query results are sorted in ascending order based on the database and the table.
	//
	// example:
	//
	// [{"Field":"TableSchema","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **10*	- (default)
	//
	// 	- **30**
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  Only data within the last 30 days can be queried.
	//
	// example:
	//
	// 2022-09-25T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table.
	//
	// >  If you leave this parameter empty, the number of accesses to all tables in the cluster on a date is returned.
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

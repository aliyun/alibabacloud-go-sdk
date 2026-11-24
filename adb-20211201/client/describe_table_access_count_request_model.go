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
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all clusters in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2ze627uzpkh8a8****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Sorts the query results by a specified field. The value is a JSON string. Example: `[{"Field":"TableSchema","Type":"Asc"}]`. Parameters in the JSON string:
	//
	// - `Field`: the field by which to sort. Valid values:
	//
	//     - `TableSchema`: the name of the database to which the table belongs.
	//
	//     - `TableName`: the name of the table.
	//
	//     - `AccessCount`: the number of times the table is accessed.
	//
	// - `Type`: the sort order. Valid values:
	//
	//     - `Asc`: ascending order.
	//
	//     - `Desc`: descending order.
	//
	// > If this parameter is not specified, the results are sorted in ascending order by the database to which the target table belongs.
	//
	// example:
	//
	// [{"Field":"TableSchema","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **10*	- (default)
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. The time is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// > You can only query data within the last 30 days.
	//
	// example:
	//
	// 2022-09-25T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the target table.
	//
	// > If this parameter is left empty, the access frequency data of all tables in the cluster within the specified date range is returned.
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

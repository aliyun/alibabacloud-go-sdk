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
	SetPageNumber(v int32) *DescribeTableStatisticsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTableStatisticsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTableStatisticsRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeTableStatisticsRequest
	GetSchemaName() *string
}

type DescribeTableStatisticsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The search keyword. Fuzzy search by table name is supported.
	//
	// example:
	//
	// you_table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The sorting rule for a specified field. The value is a JSON array. The array is ordered. Specify the fields to be sorted and the sorting order in sequence:
	//
	// For example:
	//
	// ```
	//
	// [
	//
	//     {
	//
	//         "Field":"Name",
	//
	//         "Type":"Asc"
	//
	//     }
	//
	// ]
	//
	// ```
	//
	// Field specifies the name of the field to be sorted. Only the Name field is supported.
	//
	// Type specifies the sorting order. Desc indicates descending order, and Asc indicates ascending order.
	//
	// Both parameters are case-insensitive.
	//
	// example:
	//
	// [{"Field":"SchemaName","Type":"Desc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The database name.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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

func (s *DescribeTableStatisticsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTableStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableStatisticsRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *DescribeTableStatisticsRequest) SetSchemaName(v string) *DescribeTableStatisticsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsRequest) Validate() error {
	return dara.Validate(s)
}

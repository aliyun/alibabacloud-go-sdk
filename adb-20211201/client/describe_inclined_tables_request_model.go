// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeInclinedTablesRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeInclinedTablesRequest
	GetLang() *string
	SetOrder(v string) *DescribeInclinedTablesRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeInclinedTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInclinedTablesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInclinedTablesRequest
	GetRegionId() *string
	SetTableType(v string) *DescribeInclinedTablesRequest
	GetTableType() *string
}

type DescribeInclinedTablesRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-k2jofo4pi5zhd****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh (default)**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON format.
	//
	// Example:
	//
	//     [
	//
	//         {
	//
	//             "Field":"Name",
	//
	//             "Type":"Asc"
	//
	//         }
	//
	//     ]
	//
	// Field specifies the field by which to sort the query results. Set the value to Name. Type specifies the sorting order. Valid values: Desc and Asc.
	//
	// Field and Type are case-insensitive.
	//
	// example:
	//
	// [      {          "Field":"Name",          "Type":"Asc"      }  ]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1.
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
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the table. Valid values:
	//
	// 	- **FactTable**: the partitioned table.
	//
	// 	- **DimensionTable**: the dimension table.
	//
	// example:
	//
	// FactTable
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeInclinedTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeInclinedTablesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInclinedTablesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeInclinedTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInclinedTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInclinedTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInclinedTablesRequest) GetTableType() *string {
	return s.TableType
}

func (s *DescribeInclinedTablesRequest) SetDBClusterId(v string) *DescribeInclinedTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetLang(v string) *DescribeInclinedTablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOrder(v string) *DescribeInclinedTablesRequest {
	s.Order = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageNumber(v int32) *DescribeInclinedTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageSize(v int32) *DescribeInclinedTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetRegionId(v string) *DescribeInclinedTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetTableType(v string) *DescribeInclinedTablesRequest {
	s.TableType = &v
	return s
}

func (s *DescribeInclinedTablesRequest) Validate() error {
	return dara.Validate(s)
}

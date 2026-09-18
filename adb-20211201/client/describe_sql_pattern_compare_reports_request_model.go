// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSqlPatternCompareReportsRequest
	GetDBClusterId() *string
	SetMaxResults(v int32) *DescribeSqlPatternCompareReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSqlPatternCompareReportsRequest
	GetNextToken() *string
	SetOrder(v string) *DescribeSqlPatternCompareReportsRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeSqlPatternCompareReportsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternCompareReportsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSqlPatternCompareReportsRequest
	GetRegionId() *string
}

type DescribeSqlPatternCompareReportsRequest struct {
	// The ID of the AnalyticDB for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze1234567890****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of rows per page for token-based pagination. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// > - When you use `NextToken` for pagination, keep this parameter unchanged.
	//
	// > - This parameter does not take effect when you use `PageNumber` and `PageSize` for pagination.
	//
	// > - We recommend that you use `PageNumber` and `PageSize` for pagination.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page.
	//
	// > - Do not specify this parameter for the first query. For subsequent queries, pass in the `NextToken` value returned by the previous query.
	//
	// > - Do not use this parameter together with `PageNumber` or `PageSize`.
	//
	// > - Use `PageNumber` and `PageSize` for pagination.
	//
	// example:
	//
	// djE6Mjo1MA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sorts the query results by a specified field. The value is a JSON array string, for example, `[{"Field":"CreatedAt","Type":"Desc"}]`. The array can contain only one object. Fields:
	//
	// - `Field`: the field by which to sort. Valid values:
	//
	//     - `CreatedAt`: the time when the report was created.
	//
	//     - `StartTime`: the start time of time range 1.
	//
	//     - `CompareStartTime`: the start time of time range 2.
	//
	// - `Type`: the sort order. This value is case-insensitive. Valid values:
	//
	//     - `Asc`: ascending order.
	//
	//     - `Desc`: descending order.
	//
	// > If you do not specify this parameter, the results are sorted by `CreatedAt` in descending order by default.
	//
	// example:
	//
	// [{"Field":"CreatedAt","Type":"Desc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1.
	//
	// Default value: 1.
	//
	// > Use this parameter together with `PageSize`. If you specify this parameter, `NextToken` must be empty.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// > Use this parameter together with `PageNumber`. If you specify this parameter, `NextToken` must be empty.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSqlPatternCompareReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSqlPatternCompareReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSqlPatternCompareReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSqlPatternCompareReportsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSqlPatternCompareReportsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternCompareReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternCompareReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSqlPatternCompareReportsRequest) SetDBClusterId(v string) *DescribeSqlPatternCompareReportsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetMaxResults(v int32) *DescribeSqlPatternCompareReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetNextToken(v string) *DescribeSqlPatternCompareReportsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetOrder(v string) *DescribeSqlPatternCompareReportsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetPageNumber(v int32) *DescribeSqlPatternCompareReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetPageSize(v int32) *DescribeSqlPatternCompareReportsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) SetRegionId(v string) *DescribeSqlPatternCompareReportsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsRequest) Validate() error {
	return dara.Validate(s)
}

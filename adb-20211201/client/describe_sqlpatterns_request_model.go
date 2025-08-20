// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPatternsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLPatternsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSQLPatternsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeSQLPatternsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeSQLPatternsRequest
	GetLang() *string
	SetOrder(v string) *DescribeSQLPatternsRequest
	GetOrder() *string
	SetPageNumber(v int32) *DescribeSQLPatternsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLPatternsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSQLPatternsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSQLPatternsRequest
	GetStartTime() *string
	SetUserName(v string) *DescribeSQLPatternsRequest
	GetUserName() *string
}

type DescribeSQLPatternsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vb8de93v9b****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2022-09-07T03:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language. Valid values:
	//
	// 	- **zh*	- (default): simplified Chinese.
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
	// The order by which to sort query results. Specify the parameter value in the JSON format. Example: `[{"Field":"AverageQueryTime","Type":"Asc"}]`.
	//
	// 	- `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     	- `PatternCreationTime`: the earliest commit time of the SQL pattern within the time range to query.
	//
	//     	- `AverageQueryTime`: the average total amount of time consumed by the SQL pattern within the time range to query.
	//
	//     	- `MaxQueryTime`: the maximum total amount of time consumed by the SQL pattern within the time range to query.
	//
	//     	- `AverageExecutionTime`: the average execution duration of the SQL pattern within the time range to query.
	//
	//     	- `MaxExecutionTime`: the maximum execution duration of the SQL pattern within the time range to query.
	//
	//     	- `AveragePeakMemory`: the average peak memory usage of the SQL pattern within the time range to query.
	//
	//     	- `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query.
	//
	//     	- `AverageScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query.
	//
	//     	- `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query.
	//
	//     	- `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//
	//     	- `FailedCount`: the number of failed queries performed in association with the SQL pattern within the time range to query.
	//
	// 	- `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// example:
	//
	// [{"Field":"AverageQueryTime","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 2
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// > 	- Only data within the last 14 days can be queried.
	//
	// > 	- The maximum time range that can be specified is 24 hours.
	//
	// example:
	//
	// 2022-09-06T03:06:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSQLPatternsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLPatternsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLPatternsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeSQLPatternsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSQLPatternsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSQLPatternsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLPatternsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLPatternsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSQLPatternsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLPatternsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSQLPatternsRequest) SetDBClusterId(v string) *DescribeSQLPatternsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetEndTime(v string) *DescribeSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetKeyword(v string) *DescribeSQLPatternsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetLang(v string) *DescribeSQLPatternsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetOrder(v string) *DescribeSQLPatternsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageNumber(v int32) *DescribeSQLPatternsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageSize(v int32) *DescribeSQLPatternsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetRegionId(v string) *DescribeSQLPatternsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetStartTime(v string) *DescribeSQLPatternsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetUserName(v string) *DescribeSQLPatternsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeSQLPatternsRequest) Validate() error {
	return dara.Validate(s)
}

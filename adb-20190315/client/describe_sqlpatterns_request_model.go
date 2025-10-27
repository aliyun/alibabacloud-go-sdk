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
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters in a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-30T00:15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of file titles and error messages. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// [{"Field":"AverageQueryTime","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1.
	//
	// > If you do not specify this parameter, the value **1*	- is used.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// > If you do not specify this parameter, the value **30*	- is used.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// >
	//
	// 	- Only data within the last 14 days can be queried. For example, if the current time is 2021-11-22T12:00:00Z, you can query SQL patterns at a point in time as early as 2021-11-09T12:00:00Z.
	//
	// 	- The maximum time range that can be specified is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-30T00:10:00Z
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

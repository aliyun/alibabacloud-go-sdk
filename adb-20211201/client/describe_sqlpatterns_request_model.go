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
	SetSqlPatternHash(v int64) *DescribeSQLPatternsRequest
	GetSqlPatternHash() *int64
	SetStartTime(v string) *DescribeSQLPatternsRequest
	GetStartTime() *string
	SetUserName(v string) *DescribeSQLPatternsRequest
	GetUserName() *string
}

type DescribeSQLPatternsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of all AnalyticDB for MySQL Data Lakehouse Edition clusters in a specified region, including the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vb8de93v9b****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in UTC in the <i>yyyy-MM-ddTHH:mm:ssZ</i> format.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2022-09-07T03:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language configuration. Valid values:
	//
	// - **zh**: Simplified Chinese (default language).
	//
	// - **en**: English.
	//
	// - **ja**: Japanese.
	//
	// - **zh-tw**: Traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sorting rule for the query results. The format is a JSON character string. Example: `[{"Field":"AverageQueryTime","Type":"Asc"}]`. The JSON character string contains the following parameters:
	//
	// - `Field`: the field by which to sort the results. Valid values:
	//
	//     - `PatternCreationTime`: the earliest commit time of the pattern within the query time range.
	//
	//     - `AverageQueryTime`: the average total time consumed by the pattern within the query time range.
	//
	//     - `MaxQueryTime`: the maximum total time consumed by the pattern within the query time range.
	//
	//     - `AverageExecutionTime`: the average execution duration of the pattern within the query time range.
	//
	//     - `MaxExecutionTime`: the maximum execution duration of the pattern within the query time range.
	//
	//     - `AveragePeakMemory`: the average peak memory usage of the pattern within the query time range.
	//
	//     - `MaxPeakMemory`: the maximum peak memory usage of the pattern within the query time range.
	//
	//     - `AverageScanSize`: the average amount of data read by the pattern within the query time range.
	//
	//     - `MaxScanSize`: the maximum amount of data read by the pattern within the query time range.
	//
	//     - `QueryCount`: the number of executions of the pattern within the query time range.
	//
	//     - `FailedCount`: the number of failed executions of the pattern within the query time range.
	//
	// - `Type`: the sort order. The value is not case-sensitive. Valid values:
	//
	//     - `Asc`: ascending order.
	//
	//     - `Desc`: descending order.
	//
	// example:
	//
	// [{"Field":"AverageQueryTime","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Valid values: integers greater than 0 and up to the maximum integer value.
	//
	// Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **10*	- (default value)
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The hash value of the SQL pattern, which is used to filter a specific SQL pattern.
	SqlPatternHash *int64 `json:"SqlPatternHash,omitempty" xml:"SqlPatternHash,omitempty"`
	// The start time of the query. Specify the time in UTC in the <i>yyyy-MM-ddTHH:mm:ssZ</i> format.
	//
	// >  - You can only query data from the last 14 days.
	//
	// - The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 2022-09-06T03:06:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account that executes the SQL statement.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *DescribeSQLPatternsRequest) GetSqlPatternHash() *int64 {
	return s.SqlPatternHash
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

func (s *DescribeSQLPatternsRequest) SetSqlPatternHash(v int64) *DescribeSQLPatternsRequest {
	s.SqlPatternHash = &v
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

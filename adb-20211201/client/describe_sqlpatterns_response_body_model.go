// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPatternsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSQLPatternsResponseBody
	GetAccessDeniedDetail() *string
	SetPageNumber(v int32) *DescribeSQLPatternsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLPatternsResponseBody
	GetPageSize() *int32
	SetPatternDetails(v []*DescribeSQLPatternsResponseBodyPatternDetails) *DescribeSQLPatternsResponseBody
	GetPatternDetails() []*DescribeSQLPatternsResponseBodyPatternDetails
	SetRequestId(v string) *DescribeSQLPatternsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSQLPatternsResponseBody
	GetTotalCount() *int32
}

type DescribeSQLPatternsResponseBody struct {
	// Details about the access denial. This parameter is returned only if RAM authentication fails.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of SQL patterns.
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F3174013-5B7A-5A47-9FE0-6B5D397BD86A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSQLPatternsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSQLPatternsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLPatternsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLPatternsResponseBody) GetPatternDetails() []*DescribeSQLPatternsResponseBodyPatternDetails {
	return s.PatternDetails
}

func (s *DescribeSQLPatternsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLPatternsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSQLPatternsResponseBody) SetAccessDeniedDetail(v string) *DescribeSQLPatternsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPageNumber(v int32) *DescribeSQLPatternsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPageSize(v int32) *DescribeSQLPatternsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPatternDetails(v []*DescribeSQLPatternsResponseBodyPatternDetails) *DescribeSQLPatternsResponseBody {
	s.PatternDetails = v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetRequestId(v string) *DescribeSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetTotalCount(v int32) *DescribeSQLPatternsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) Validate() error {
	if s.PatternDetails != nil {
		for _, item := range s.PatternDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLPatternsResponseBodyPatternDetails struct {
	// The client IP address used to submit the queries.
	//
	// example:
	//
	// 192.168.xx.xx
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The average execution time of queries matching this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 234.78
	AverageExecutionTime *float64 `json:"AverageExecutionTime,omitempty" xml:"AverageExecutionTime,omitempty"`
	// The average CPU cost for queries that match this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	AverageOperatorCost *float64 `json:"AverageOperatorCost,omitempty" xml:"AverageOperatorCost,omitempty"`
	// The average peak memory usage of queries matching this pattern. Unit: bytes.
	//
	// example:
	//
	// 234.22
	AveragePeakMemory *float64 `json:"AveragePeakMemory,omitempty" xml:"AveragePeakMemory,omitempty"`
	// The average duration of queries matching this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 4
	AverageQueryTime *float64 `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	// The average scan time for queries that match this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	AverageScanCost *float64 `json:"AverageScanCost,omitempty" xml:"AverageScanCost,omitempty"`
	// The average amount of data scanned by queries matching this pattern. Unit: bytes.
	//
	// example:
	//
	// 234149.23
	AverageScanSize *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	// Indicates whether queries that match this pattern can be blocked. Valid values:
	//
	// - **true**: The queries can be blocked.
	//
	// - **false**: The queries cannot be blocked.
	//
	// > Currently, AnalyticDB for MySQL allows you to block only SELECT and INSERT statements.
	//
	// example:
	//
	// true
	Blockable *bool `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	// The number of failed queries that match this pattern.
	//
	// example:
	//
	// 18
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The maximum execution time of a query matching this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 2142
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The maximum CPU cost for a query that matches this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	MaxOperatorCost *float64 `json:"MaxOperatorCost,omitempty" xml:"MaxOperatorCost,omitempty"`
	// The maximum peak memory usage of a query matching this pattern. Unit: bytes.
	//
	// example:
	//
	// 234149
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum duration of a query matching this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 2341
	MaxQueryTime *int64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// The maximum scan time for a query that matches this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	MaxScanCost *float64 `json:"MaxScanCost,omitempty" xml:"MaxScanCost,omitempty"`
	// The maximum amount of data scanned by a query matching this pattern. Unit: bytes.
	//
	// example:
	//
	// 32212254
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The total CPU cost of queries matching this pattern as a percentage of the total CPU cost for all queries. Unit: %.
	//
	// example:
	//
	// 20
	OperatorCostPercentage *float64 `json:"OperatorCostPercentage,omitempty" xml:"OperatorCostPercentage,omitempty"`
	// The total CPU cost for all queries that match this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	OperatorCostSum *float64 `json:"OperatorCostSum,omitempty" xml:"OperatorCostSum,omitempty"`
	// The submission time of the first query that matches this pattern within the specified time range.
	//
	// example:
	//
	// 2022-09-06 05:06:00
	PatternCreationTime *string `json:"PatternCreationTime,omitempty" xml:"PatternCreationTime,omitempty"`
	// The ID of the SQL pattern.
	//
	// example:
	//
	// 5575924945138******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The total peak memory usage of queries matching this pattern as a percentage of the total peak memory usage for all queries. Unit: %.
	//
	// example:
	//
	// 10
	PeakMemoryPercentage *float64 `json:"PeakMemoryPercentage,omitempty" xml:"PeakMemoryPercentage,omitempty"`
	// The sum of the peak memory usage for all queries that match this pattern. Unit: bytes.
	//
	// example:
	//
	// 5
	PeakMemorySum *float64 `json:"PeakMemorySum,omitempty" xml:"PeakMemorySum,omitempty"`
	// The number of executed queries that match this pattern.
	//
	// example:
	//
	// 345
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The total query time of queries matching this pattern as a percentage of the total query time for all queries. Unit: %.
	//
	// example:
	//
	// 10
	QueryTimePercentage *float64 `json:"QueryTimePercentage,omitempty" xml:"QueryTimePercentage,omitempty"`
	// The total query duration for all queries that match this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" xml:"QueryTimeSum,omitempty"`
	// The SQL pattern.
	//
	// example:
	//
	// SELECT 	- FROM KEPLER_META_NODE_STATIC_INFO WHERE elastic_node = ? OR (elastic_node = ? AND enable = ?)
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The total scan cost of queries matching this pattern as a percentage of the total scan cost for all queries. Unit: %.
	//
	// example:
	//
	// 5
	ScanCostPercentage *float64 `json:"ScanCostPercentage,omitempty" xml:"ScanCostPercentage,omitempty"`
	// The total scan cost for all queries that match this pattern. Unit: milliseconds.
	//
	// example:
	//
	// 5
	ScanCostSum *float64 `json:"ScanCostSum,omitempty" xml:"ScanCostSum,omitempty"`
	// The total amount of data scanned by queries matching this pattern as a percentage of the total data scanned by all queries. Unit: %.
	//
	// example:
	//
	// 80
	ScanSizePercentage *float64 `json:"ScanSizePercentage,omitempty" xml:"ScanSizePercentage,omitempty"`
	// The total amount of data scanned by all queries that match this pattern. Unit: bytes.
	//
	// example:
	//
	// 5
	ScanSizeSum *float64 `json:"ScanSizeSum,omitempty" xml:"ScanSizeSum,omitempty"`
	// The tables scanned by the SQL pattern.
	//
	// example:
	//
	// tpch.orders
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The name of the database user who submitted the matching SQL statements.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAccessIp() *string {
	return s.AccessIp
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAverageExecutionTime() *float64 {
	return s.AverageExecutionTime
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAverageOperatorCost() *float64 {
	return s.AverageOperatorCost
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAveragePeakMemory() *float64 {
	return s.AveragePeakMemory
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAverageQueryTime() *float64 {
	return s.AverageQueryTime
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAverageScanCost() *float64 {
	return s.AverageScanCost
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetAverageScanSize() *float64 {
	return s.AverageScanSize
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetBlockable() *bool {
	return s.Blockable
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxExecutionTime() *int64 {
	return s.MaxExecutionTime
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxOperatorCost() *float64 {
	return s.MaxOperatorCost
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxPeakMemory() *int64 {
	return s.MaxPeakMemory
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxQueryTime() *int64 {
	return s.MaxQueryTime
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxScanCost() *float64 {
	return s.MaxScanCost
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetMaxScanSize() *int64 {
	return s.MaxScanSize
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetOperatorCostPercentage() *float64 {
	return s.OperatorCostPercentage
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetOperatorCostSum() *float64 {
	return s.OperatorCostSum
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetPatternCreationTime() *string {
	return s.PatternCreationTime
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetPeakMemoryPercentage() *float64 {
	return s.PeakMemoryPercentage
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetPeakMemorySum() *float64 {
	return s.PeakMemorySum
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetQueryCount() *int64 {
	return s.QueryCount
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetQueryTimePercentage() *float64 {
	return s.QueryTimePercentage
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetQueryTimeSum() *float64 {
	return s.QueryTimeSum
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetSQLPattern() *string {
	return s.SQLPattern
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetScanCostPercentage() *float64 {
	return s.ScanCostPercentage
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetScanCostSum() *float64 {
	return s.ScanCostSum
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetScanSizePercentage() *float64 {
	return s.ScanSizePercentage
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetScanSizeSum() *float64 {
	return s.ScanSizeSum
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetTables() *string {
	return s.Tables
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) GetUser() *string {
	return s.User
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAccessIp(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AccessIp = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageExecutionTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageOperatorCost(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageOperatorCost = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAveragePeakMemory(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AveragePeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageQueryTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageScanCost(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageScanCost = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageScanSize(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetBlockable(v bool) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Blockable = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetFailedCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.FailedCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxExecutionTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxOperatorCost(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxOperatorCost = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxPeakMemory(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxQueryTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxScanCost(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxScanCost = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxScanSize(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetOperatorCostPercentage(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.OperatorCostPercentage = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetOperatorCostSum(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.OperatorCostSum = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternCreationTime(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternCreationTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternId(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPeakMemoryPercentage(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PeakMemoryPercentage = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPeakMemorySum(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PeakMemorySum = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryTimePercentage(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryTimePercentage = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryTimeSum(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryTimeSum = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetSQLPattern(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.SQLPattern = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetScanCostPercentage(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.ScanCostPercentage = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetScanCostSum(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.ScanCostSum = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetScanSizePercentage(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.ScanSizePercentage = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetScanSizeSum(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.ScanSizeSum = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetTables(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Tables = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetUser(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.User = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) Validate() error {
	return dara.Validate(s)
}

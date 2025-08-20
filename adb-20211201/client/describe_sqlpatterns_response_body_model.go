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
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
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
	// The queried SQL patterns.
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F3174013-5B7A-5A47-9FE0-6B5D397BD86B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	return dara.Validate(s)
}

type DescribeSQLPatternsResponseBodyPatternDetails struct {
	// The IP address of the SQL client that commits the SQL pattern.
	//
	// example:
	//
	// 192.168.xx.xx
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The average execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 234.78
	AverageExecutionTime *float64 `json:"AverageExecutionTime,omitempty" xml:"AverageExecutionTime,omitempty"`
	AverageOperatorCost  *float64 `json:"AverageOperatorCost,omitempty" xml:"AverageOperatorCost,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 234.22
	AveragePeakMemory *float64 `json:"AveragePeakMemory,omitempty" xml:"AveragePeakMemory,omitempty"`
	// The average total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 4
	AverageQueryTime *float64 `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	AverageScanCost  *float64 `json:"AverageScanCost,omitempty" xml:"AverageScanCost,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 234149.23
	AverageScanSize *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	// Indicates whether the execution of the SQL pattern can be intercepted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  Only SELECT and INSERT statements can be intercepted.
	//
	// example:
	//
	// true
	Blockable *bool `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	// The number of failed queries executed in association with the SQL pattern within the query time range.
	//
	// example:
	//
	// 18
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The maximum execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 2142
	MaxExecutionTime *int64   `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	MaxOperatorCost  *float64 `json:"MaxOperatorCost,omitempty" xml:"MaxOperatorCost,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 234149
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 2341
	MaxQueryTime *int64   `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	MaxScanCost  *float64 `json:"MaxScanCost,omitempty" xml:"MaxScanCost,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 32212254
	MaxScanSize            *int64   `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	OperatorCostPercentage *float64 `json:"OperatorCostPercentage,omitempty" xml:"OperatorCostPercentage,omitempty"`
	OperatorCostSum        *float64 `json:"OperatorCostSum,omitempty" xml:"OperatorCostSum,omitempty"`
	// The earliest commit time of the SQL pattern within the query time range.
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
	PatternId            *string  `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	PeakMemoryPercentage *float64 `json:"PeakMemoryPercentage,omitempty" xml:"PeakMemoryPercentage,omitempty"`
	PeakMemorySum        *float64 `json:"PeakMemorySum,omitempty" xml:"PeakMemorySum,omitempty"`
	// The number of queries executed in association with the SQL pattern within the query time range.
	//
	// example:
	//
	// 345
	QueryCount          *int64   `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	QueryTimePercentage *float64 `json:"QueryTimePercentage,omitempty" xml:"QueryTimePercentage,omitempty"`
	QueryTimeSum        *float64 `json:"QueryTimeSum,omitempty" xml:"QueryTimeSum,omitempty"`
	// The statement of the SQL pattern.
	//
	// example:
	//
	// SELECT 	- FROM KEPLER_META_NODE_STATIC_INFO WHERE elastic_node = ? OR (elastic_node = ? AND enable = ?)
	SQLPattern         *string  `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	ScanCostPercentage *float64 `json:"ScanCostPercentage,omitempty" xml:"ScanCostPercentage,omitempty"`
	ScanCostSum        *float64 `json:"ScanCostSum,omitempty" xml:"ScanCostSum,omitempty"`
	ScanSizePercentage *float64 `json:"ScanSizePercentage,omitempty" xml:"ScanSizePercentage,omitempty"`
	ScanSizeSum        *float64 `json:"ScanSizeSum,omitempty" xml:"ScanSizeSum,omitempty"`
	// The tables scanned based on the SQL pattern.
	//
	// example:
	//
	// tpch.orders
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The name of the database account that is used to commit the SQL pattern.
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

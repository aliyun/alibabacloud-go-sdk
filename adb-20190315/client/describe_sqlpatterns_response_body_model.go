// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPatternsResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	// The queried SQL patterns.
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6BE0EDD1-0DE6-3EB6-81BF-BFE4F2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSQLPatternsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPatternsResponseBody) GoString() string {
	return s.String()
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
	// The average operation duration.
	//
	// example:
	//
	// 5
	AverageOperatorCost *float64 `json:"AverageOperatorCost,omitempty" xml:"AverageOperatorCost,omitempty"`
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
	// The average scan duration.
	//
	// example:
	//
	// 3
	AverageScanCost *float64 `json:"AverageScanCost,omitempty" xml:"AverageScanCost,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 234149.23
	AverageScanSize *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	// Indicates whether the execution of the SQL pattern can be blocked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > Only SELECT and INSERT statements can be blocked.
	//
	// example:
	//
	// true
	Blockable *bool `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	// The number of failed queries executed in association with the SQL pattern within the query time range.
	//
	// example:
	//
	// 234
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The maximum execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 2142
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The maximum operation duration.
	//
	// example:
	//
	// 10
	MaxOperatorCost *float64 `json:"MaxOperatorCost,omitempty" xml:"MaxOperatorCost,omitempty"`
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
	MaxQueryTime *int64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// The maximum scan duration.
	//
	// example:
	//
	// 7
	MaxScanCost *float64 `json:"MaxScanCost,omitempty" xml:"MaxScanCost,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	//
	// example:
	//
	// 234149
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The operation duration percentage.
	//
	// example:
	//
	// 75
	OperatorCostPercentage *float64 `json:"OperatorCostPercentage,omitempty" xml:"OperatorCostPercentage,omitempty"`
	// The total operation duration.
	//
	// example:
	//
	// 20
	OperatorCostSum *float64 `json:"OperatorCostSum,omitempty" xml:"OperatorCostSum,omitempty"`
	// The earliest commit time of the SQL pattern within the query time range. Unit: milliseconds.
	//
	// example:
	//
	// 2021-11-12 03:06:00
	PatternCreationTime *string `json:"PatternCreationTime,omitempty" xml:"PatternCreationTime,omitempty"`
	// The ID of the SQL pattern.
	//
	// example:
	//
	// 5575924945138******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The peak memory percentage.
	//
	// example:
	//
	// 80
	PeakMemoryPercentage *float64 `json:"PeakMemoryPercentage,omitempty" xml:"PeakMemoryPercentage,omitempty"`
	// The total peak memory.
	//
	// example:
	//
	// 3600
	PeakMemorySum *float64 `json:"PeakMemorySum,omitempty" xml:"PeakMemorySum,omitempty"`
	// The number of queries executed in association with the SQL pattern within the query time range.
	//
	// example:
	//
	// 345
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The queue duration of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 80
	QueryTimePercentage *float64 `json:"QueryTimePercentage,omitempty" xml:"QueryTimePercentage,omitempty"`
	// The total query duration.
	//
	// example:
	//
	// 5
	QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" xml:"QueryTimeSum,omitempty"`
	// The statement of the SQL pattern.
	//
	// example:
	//
	// SELECT 	- FROM KEPLER_META_NODE_STATIC_INFO WHERE elastic_node = ? OR (elastic_node = ? AND enable = ?)
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The scan duration percentage.
	//
	// example:
	//
	// 75
	ScanCostPercentage *float64 `json:"ScanCostPercentage,omitempty" xml:"ScanCostPercentage,omitempty"`
	// The total scan duration.
	//
	// example:
	//
	// 11
	ScanCostSum *float64 `json:"ScanCostSum,omitempty" xml:"ScanCostSum,omitempty"`
	// The amount of time consumed to scan data from a data source in the task. Unit: milliseconds.
	//
	// example:
	//
	// 80
	ScanSizePercentage *float64 `json:"ScanSizePercentage,omitempty" xml:"ScanSizePercentage,omitempty"`
	// Total total scan size.
	//
	// example:
	//
	// 3
	ScanSizeSum *float64 `json:"ScanSizeSum,omitempty" xml:"ScanSizeSum,omitempty"`
	// The tables scanned based on the SQL pattern.
	//
	// example:
	//
	// tpch.orders
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The database username that is used to commit the SQL pattern.
	//
	// example:
	//
	// reporter
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

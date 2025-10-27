// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessIp(v string) *DescribePatternPerformanceResponseBody
	GetAccessIp() *string
	SetEndTime(v string) *DescribePatternPerformanceResponseBody
	GetEndTime() *string
	SetFailedCount(v int64) *DescribePatternPerformanceResponseBody
	GetFailedCount() *int64
	SetPerformances(v []*DescribePatternPerformanceResponseBodyPerformances) *DescribePatternPerformanceResponseBody
	GetPerformances() []*DescribePatternPerformanceResponseBodyPerformances
	SetQueryCount(v int64) *DescribePatternPerformanceResponseBody
	GetQueryCount() *int64
	SetRequestId(v string) *DescribePatternPerformanceResponseBody
	GetRequestId() *string
	SetSQLPattern(v string) *DescribePatternPerformanceResponseBody
	GetSQLPattern() *string
	SetStartTime(v string) *DescribePatternPerformanceResponseBody
	GetStartTime() *string
	SetTables(v string) *DescribePatternPerformanceResponseBody
	GetTables() *string
	SetUser(v string) *DescribePatternPerformanceResponseBody
	GetUser() *string
}

type DescribePatternPerformanceResponseBody struct {
	// The IP address of the SQL client that submits the SQL pattern.
	//
	// example:
	//
	// 172.16.14.*
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-18T18:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of failed SQL patterns within the time range.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The queried performance metrics.
	Performances []*DescribePatternPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The number of queries executed in association with the SQL pattern within the time range.
	//
	// example:
	//
	// 35018
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 210f47011634026610213529******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement that represents the SQL pattern.
	//
	// example:
	//
	// SELECT *nFROM HIVE.`ADB_EXTERNAL_TPCH_10GB`.`External_customer`nLIMIT ?
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-18T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tables that are in association with the SQL pattern.
	//
	// example:
	//
	// tpch_1g.part;tpch_1g.supplier;tpch_1g.lineitem;tpch_1g.partsupp;tpch_1g.orders;tpch_1g.nation
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The database account that is used to execute the SQL statement.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribePatternPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBody) GetAccessIp() *string {
	return s.AccessIp
}

func (s *DescribePatternPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePatternPerformanceResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DescribePatternPerformanceResponseBody) GetPerformances() []*DescribePatternPerformanceResponseBodyPerformances {
	return s.Performances
}

func (s *DescribePatternPerformanceResponseBody) GetQueryCount() *int64 {
	return s.QueryCount
}

func (s *DescribePatternPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePatternPerformanceResponseBody) GetSQLPattern() *string {
	return s.SQLPattern
}

func (s *DescribePatternPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePatternPerformanceResponseBody) GetTables() *string {
	return s.Tables
}

func (s *DescribePatternPerformanceResponseBody) GetUser() *string {
	return s.User
}

func (s *DescribePatternPerformanceResponseBody) SetAccessIp(v string) *DescribePatternPerformanceResponseBody {
	s.AccessIp = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetEndTime(v string) *DescribePatternPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetFailedCount(v int64) *DescribePatternPerformanceResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetPerformances(v []*DescribePatternPerformanceResponseBodyPerformances) *DescribePatternPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetQueryCount(v int64) *DescribePatternPerformanceResponseBody {
	s.QueryCount = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetRequestId(v string) *DescribePatternPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetSQLPattern(v string) *DescribePatternPerformanceResponseBody {
	s.SQLPattern = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetStartTime(v string) *DescribePatternPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetTables(v string) *DescribePatternPerformanceResponseBody {
	s.Tables = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetUser(v string) *DescribePatternPerformanceResponseBody {
	s.User = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) Validate() error {
	if s.Performances != nil {
		for _, item := range s.Performances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePatternPerformanceResponseBodyPerformances struct {
	// The performance metric that was queried. Valid values:
	//
	// 	- **AnalyticDB_PatternQueryCount**: the total number of queries executed in association with the SQL pattern.
	//
	// 	- **AnalyticDB_PatternQueryTime**: the total amount of time consumed by the queries executed in association with the SQL pattern.
	//
	// 	- **AnalyticDB_PatternExecutionTime**: the total execution duration of the queries executed in association with the SQL pattern.
	//
	// 	- **AnalyticDB_PatternPeakMemory**: the peak memory usage of the queries executed in association with the SQL pattern.
	//
	// 	- **AnalyticDB_PatternScanSize**: the amount of data scanned in the queries executed in association with the SQL pattern.
	//
	// example:
	//
	// AnalyticDB_PatternQueryCount
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried performance metrics.
	Series []*DescribePatternPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric. Valid values:
	//
	// 	- When the performance metric is related to the query duration (the `Key` value is `AnalyticDB_PatternQueryTime` or `AnalyticDB_PatternExecutionTime`), **ms*	- is returned.
	//
	// 	- When the performance metric is related to the memory usage (the `Key` value is `AnalyticDB_PatternPeakMemory`), **MB*	- is returned.
	//
	// 	- When the performance metric is related to the amount of data scanned (the `Key` value is `AnalyticDB_PatternScanSize`), **MB*	- is returned.
	//
	// 	- When the performance metric is related to the number of queries (the `Key` value is `AnalyticDB_PatternQueryCount`), this parameter is empty.
	//
	// example:
	//
	// ms
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribePatternPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformances) GetKey() *string {
	return s.Key
}

func (s *DescribePatternPerformanceResponseBodyPerformances) GetSeries() []*DescribePatternPerformanceResponseBodyPerformancesSeries {
	return s.Series
}

func (s *DescribePatternPerformanceResponseBodyPerformances) GetUnit() *string {
	return s.Unit
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetKey(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetSeries(v []*DescribePatternPerformanceResponseBodyPerformancesSeries) *DescribePatternPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetUnit(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePatternPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value. Valid values:
	//
	// 	- When the `Key` parameter is set to `AnalyticDB_PatternQueryCount`, `pattern_query_count` is returned, which indicates the number of executions of the SQL statements in association with the SQL pattern.
	//
	// 	- When the `Key` parameter is set to `AnalyticDB_PatternQueryTime`, the following values are returned:
	//
	//     	- `average_query_time`, which indicates the average total amount of time consumed by the SQL statements in association with the SQL pattern.
	//
	//     	- `max_query_time`, which indicates the maximum total amount of time consumed by the SQL statements in association with the SQL pattern.
	//
	// 	- When the `Key` parameter is set to `AnalyticDB_PatternExecutionTime`, the following values are returned:
	//
	//     	- `average_execution_time`, which indicates the average execution duration of the SQL statements in association with the SQL pattern.
	//
	//     	- `max_execution_time`, which indicates the maximum execution duration of the SQL statements in association with the SQL pattern.
	//
	// 	- When the `Key` parameter is set to `AnalyticDB_PatternPeakMemory`, the following values are returned:
	//
	//     	- `average_peak_memory`, which indicates the average peak memory usage of the SQL statements in association with the SQL pattern.
	//
	//     	- `max_peak_memory`, which indicates the maximum peak memory usage of the SQL statements in association with the SQL pattern.
	//
	// 	- When the `Key` parameter is set `AnalyticDB_PatternScanSize`, the following values are returned:
	//
	//     	- `average_scan_size`, which indicates the average amount of data scanned by the SQL statements in association with the SQL pattern.
	//
	//     	- `max_scan_size`, which indicates the maximum amount of data scanned by the SQL statements in association with the SQL pattern.
	//
	// example:
	//
	// max_query_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The queried performance metrics.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) GetName() *string {
	return s.Name
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) GetValues() []*string {
	return s.Values
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) Validate() error {
	return dara.Validate(s)
}

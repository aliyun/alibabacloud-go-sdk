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
	// The IP address of the client that submits the SQL statements related to the pattern.
	//
	// example:
	//
	// 172.16.14.*
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The end time of the query. The time is in UTC and follows the <i>yyyy-MM-ddTHH:mmZ</i> format.
	//
	// example:
	//
	// 2022-08-22T01:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of failures of the pattern within the query time range.
	//
	// example:
	//
	// 1
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The performance details.
	Performances []*DescribePatternPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The number of executions of the pattern within the query time range.
	//
	// example:
	//
	// 1202
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F21AF487-B8C9-57E0-8E3A-A92BC3611FB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specific statement of the SQL pattern.
	//
	// example:
	//
	// SELECT *nFROM HIVE.`ADB_EXTERNAL_TPCH_10GB`.`External_customer`nLIMIT ?
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The start time of the query. The time is in UTC and follows the <i>yyyy-MM-ddTHH:mmZ</i> format.
	//
	// example:
	//
	// 2022-08-21T02:15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The information about the tables involved in the SQL pattern query.
	//
	// example:
	//
	// tpch_1g.part;tpch_1g.supplier;tpch_1g.lineitem;tpch_1g.partsupp;tpch_1g.orders;tpch_1g.nation
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The database account that executes the SQL statement.
	//
	// example:
	//
	// test_user
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
	// The queried performance metrics. Valid values:
	//
	// - **AnalyticDB_PatternQueryCount**: the total number of queries related to the pattern.
	//
	// - **AnalyticDB_PatternQueryTime**: the total time consumed by queries related to the pattern.
	//
	// - **AnalyticDB_PatternExecutionTime**: the execution duration of queries related to the pattern.
	//
	// - **AnalyticDB_PatternPeakMemory**: the peak memory consumed by queries related to the pattern.
	//
	// - **AnalyticDB_PatternScanSize**: the data volume read by queries related to the pattern.
	//
	// example:
	//
	// AnalyticDB_PatternExecutionTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The details of each performance value under the performance metrics.
	Series []*DescribePatternPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The units corresponding to the performance metrics. Valid values:
	//
	// - **ms**: returned when the performance metric is related to query time, which means the `Key` value is `AnalyticDB_PatternQueryTime` or `AnalyticDB_PatternExecutionTime`.
	//
	// - **MB**: returned when the performance metric is related to peak memory, which means the `Key` value is `AnalyticDB_PatternPeakMemory`.
	//
	// - **MB**: returned when the performance metric is the data volume read, which means the `Key` value is `AnalyticDB_PatternScanSize`.
	//
	// - Empty: returned when the performance metric is the number of queries, which means the `Key` value is `AnalyticDB_PatternQueryCount`.
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
	// The name of the performance value. Valid values:
	//
	// - When the `Key` value is `AnalyticDB_PatternQueryCount`, this parameter returns `pattern_query_count`, which indicates the executions of SQL statements related to the current pattern.
	//
	// - When the `Key` value is `AnalyticDB_PatternQueryTime`, this parameter returns:
	//
	//     - `average_query_time`: the average total time consumed by SQL statements related to the current pattern.
	//
	//     - `max_query_time`: the maximum total time consumed by SQL statements related to the current pattern.
	//
	// - When the `Key` value is `AnalyticDB_PatternExecutionTime`, this parameter returns:
	//
	//     - `average_execution_time`: the average execution duration of SQL statements related to the current pattern.
	//
	//     - `max_execution_time`: the maximum execution duration of SQL statements related to the current pattern.
	//
	// - When the `Key` value is `AnalyticDB_PatternPeakMemory`, this parameter returns:
	//
	//     - `average_peak_memory`: the average peak memory consumed by SQL statements related to the current pattern.
	//
	//     - `max_peak_memory`: the maximum peak memory consumed by SQL statements related to the current pattern.
	//
	// - When the `Key` value is `AnalyticDB_PatternScanSize`, this parameter returns:
	//
	//     - `average_scan_size`: the average data volume read by SQL statements related to the current pattern.
	//
	//     - `max_scan_size`: the maximum data volume read by SQL statements related to the current pattern.
	//
	// example:
	//
	// max_query_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of performance values.
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

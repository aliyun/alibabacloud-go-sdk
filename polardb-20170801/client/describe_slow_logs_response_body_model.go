// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogsResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSlowLogsResponseBody
	GetEndTime() *string
	SetEngine(v string) *DescribeSlowLogsResponseBody
	GetEngine() *string
	SetItems(v *DescribeSlowLogsResponseBodyItems) *DescribeSlowLogsResponseBody
	GetItems() *DescribeSlowLogsResponseBodyItems
	SetPageNumber(v int32) *DescribeSlowLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSlowLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSlowLogsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeSlowLogsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeSlowLogsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSlowLogsResponseBody struct {
	// The ID of cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end date of the query.
	//
	// example:
	//
	// 2021-05-30Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// polardb_mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Details about slow query logs.
	Items *DescribeSlowLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of SQL statements that are returned on the current page.
	//
	// example:
	//
	// 6
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2553A660-E4EB-4AF4-A402-8AFF70A49143
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start date of the query.
	//
	// example:
	//
	// 2021-05-01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSlowLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSlowLogsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSlowLogsResponseBody) GetItems() *DescribeSlowLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSlowLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSlowLogsResponseBody) SetDBClusterId(v string) *DescribeSlowLogsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetEndTime(v string) *DescribeSlowLogsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetEngine(v string) *DescribeSlowLogsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetItems(v *DescribeSlowLogsResponseBodyItems) *DescribeSlowLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetPageNumber(v int32) *DescribeSlowLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetRequestId(v string) *DescribeSlowLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetStartTime(v string) *DescribeSlowLogsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogsResponseBodyItems struct {
	SQLSlowLog []*DescribeSlowLogsResponseBodyItemsSQLSlowLog `json:"SQLSlowLog,omitempty" xml:"SQLSlowLog,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBodyItems) GetSQLSlowLog() []*DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	return s.SQLSlowLog
}

func (s *DescribeSlowLogsResponseBodyItems) SetSQLSlowLog(v []*DescribeSlowLogsResponseBodyItemsSQLSlowLog) *DescribeSlowLogsResponseBodyItems {
	s.SQLSlowLog = v
	return s
}

func (s *DescribeSlowLogsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogsResponseBodyItemsSQLSlowLog struct {
	// The date when the data was generated.
	//
	// example:
	//
	// 2021-05-30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// PolarDB_MySQL
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// pi-***************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The longest execution duration of a specific SQL statement in the query. Unit: seconds.
	//
	// example:
	//
	// 60
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The longest lock duration that was caused by a specific SQL statement in the query. Unit: seconds.
	//
	// example:
	//
	// 1
	MaxLockTime *int64 `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty"`
	// The largest number of rows that were parsed by a specific SQL statement in the query.
	//
	// example:
	//
	// 1
	ParseMaxRowCount *int64 `json:"ParseMaxRowCount,omitempty" xml:"ParseMaxRowCount,omitempty"`
	// The total number of rows that were parsed by all SQL statements in the query.
	//
	// example:
	//
	// 2
	ParseTotalRowCounts *int64 `json:"ParseTotalRowCounts,omitempty" xml:"ParseTotalRowCounts,omitempty"`
	// The largest number of rows that were returned by a specific SQL statement in the query.
	//
	// example:
	//
	// 3
	ReturnMaxRowCount *int64 `json:"ReturnMaxRowCount,omitempty" xml:"ReturnMaxRowCount,omitempty"`
	// The total number of rows that were returned by all SQL statements in the query.
	//
	// example:
	//
	// 1
	ReturnTotalRowCounts *int64 `json:"ReturnTotalRowCounts,omitempty" xml:"ReturnTotalRowCounts,omitempty"`
	// The unique ID of the SQL statement. The ID is used to obtain the slow query logs of the SQL statement.
	//
	// example:
	//
	// U2FsdGVkxxxx
	SQLHASH *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
	// The SQL statement that is executed in the query.
	//
	// example:
	//
	// select id,name from tb_table
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The total number of executions of the SQL statements.
	//
	// example:
	//
	// 2
	TotalExecutionCounts *int64 `json:"TotalExecutionCounts,omitempty" xml:"TotalExecutionCounts,omitempty"`
	// The total duration that was caused by all SQL statements in the query. Unit: seconds.
	//
	// example:
	//
	// 2
	TotalExecutionTimes *int64 `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
	// The total lock duration that was caused by all SQL statements in the query. Unit: seconds.
	//
	// example:
	//
	// 1
	TotalLockTimes *int64 `json:"TotalLockTimes,omitempty" xml:"TotalLockTimes,omitempty"`
}

func (s DescribeSlowLogsResponseBodyItemsSQLSlowLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsResponseBodyItemsSQLSlowLog) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxExecutionTime() *int64 {
	return s.MaxExecutionTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetMaxLockTime() *int64 {
	return s.MaxLockTime
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetParseMaxRowCount() *int64 {
	return s.ParseMaxRowCount
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetParseTotalRowCounts() *int64 {
	return s.ParseTotalRowCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetReturnMaxRowCount() *int64 {
	return s.ReturnMaxRowCount
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetReturnTotalRowCounts() *int64 {
	return s.ReturnTotalRowCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLHASH() *string {
	return s.SQLHASH
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalExecutionCounts() *int64 {
	return s.TotalExecutionCounts
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalExecutionTimes() *int64 {
	return s.TotalExecutionTimes
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) GetTotalLockTimes() *int64 {
	return s.TotalLockTimes
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetCreateTime(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.CreateTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetDBName(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetDBNodeId(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.DBNodeId = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxExecutionTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetMaxLockTime(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.MaxLockTime = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetParseMaxRowCount(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ParseMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetParseTotalRowCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ParseTotalRowCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetReturnMaxRowCount(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ReturnMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetReturnTotalRowCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.ReturnTotalRowCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLHASH(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLHASH = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetSQLText(v string) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalExecutionCounts(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalExecutionCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalExecutionTimes(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) SetTotalLockTimes(v int64) *DescribeSlowLogsResponseBodyItemsSQLSlowLog {
	s.TotalLockTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseBodyItemsSQLSlowLog) Validate() error {
	return dara.Validate(s)
}

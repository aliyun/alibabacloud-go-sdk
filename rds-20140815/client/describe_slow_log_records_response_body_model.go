// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogRecordsResponseBody
	GetDBInstanceId() *string
	SetEngine(v string) *DescribeSlowLogRecordsResponseBody
	GetEngine() *string
	SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody
	GetItems() *DescribeSlowLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSlowLogRecordsResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-uf6wjk5*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// An array that consists of the information about each slow log.
	Items *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of SQL log reports on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4DBB1BB0-E5D8-4D41-B1C9-142364DB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogRecordsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeSlowLogRecordsResponseBody) GetItems() *DescribeSlowLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBInstanceId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord `json:"SQLSlowRecord,omitempty" xml:"SQLSlowRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetSQLSlowRecord() []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	return s.SQLSlowRecord
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLSlowRecord = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) Validate() error {
	if s.SQLSlowRecord != nil {
		for _, item := range s.SQLSlowRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord struct {
	// The name of the application that is connected to the instance.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// example
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The hostname of the client.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// example
	ClientHostName *string `json:"ClientHostName,omitempty" xml:"ClientHostName,omitempty"`
	// The duration during which the SQL statement is processed by the CPU. Unit: milliseconds.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 0.002
	CpuTime *int64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time when the execution of the SQL statement started. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-18T01:40:44Z
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// The name and IP address of the client that is connected to the database.
	//
	// example:
	//
	// xxx[xxx] @  [1xx.xxx.xxx.xx]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The number of rows that are affected by the last SQL statement.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 2
	LastRowsAffectedCount *int64 `json:"LastRowsAffectedCount,omitempty" xml:"LastRowsAffectedCount,omitempty"`
	// The lock duration of the query. Unit: seconds.
	//
	// example:
	//
	// 0
	LockTimes *int64 `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	// The number of logical reads.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 383
	LogicalIORead *int64 `json:"LogicalIORead,omitempty" xml:"LogicalIORead,omitempty"`
	// The number of parsed rows.
	//
	// example:
	//
	// 1
	ParseRowCounts *int64 `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// The number of physical reads.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 200
	PhysicalIORead *int64 `json:"PhysicalIORead,omitempty" xml:"PhysicalIORead,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 2001
	QueryTimeMS *int64 `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// The execution duration of the query. Unit: seconds.
	//
	// example:
	//
	// 2
	QueryTimes *int64 `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 1
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The number of affected rows.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 34
	RowsAffectedCount *int64 `json:"RowsAffectedCount,omitempty" xml:"RowsAffectedCount,omitempty"`
	// The unique ID of the SQL statement.
	//
	// example:
	//
	// U2FsdGVk****
	SQLHash *string `json:"SQLHash,omitempty" xml:"SQLHash,omitempty"`
	// The details of the SQL statement.
	//
	// example:
	//
	// select sleep(2)
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The name of the user.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The number of I/O writes.
	//
	// > This parameter is returned only for instances that run SQL Server.
	//
	// example:
	//
	// 22
	WriteIOCount *int64 `json:"WriteIOCount,omitempty" xml:"WriteIOCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetClientHostName() *string {
	return s.ClientHostName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetCpuTime() *int64 {
	return s.CpuTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetLastRowsAffectedCount() *int64 {
	return s.LastRowsAffectedCount
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetLockTimes() *int64 {
	return s.LockTimes
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetLogicalIORead() *int64 {
	return s.LogicalIORead
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetParseRowCounts() *int64 {
	return s.ParseRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetPhysicalIORead() *int64 {
	return s.PhysicalIORead
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetQueryTimeMS() *int64 {
	return s.QueryTimeMS
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetQueryTimes() *int64 {
	return s.QueryTimes
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetRowsAffectedCount() *int64 {
	return s.RowsAffectedCount
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetSQLHash() *string {
	return s.SQLHash
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetWriteIOCount() *int64 {
	return s.WriteIOCount
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetApplicationName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ApplicationName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetClientHostName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ClientHostName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetCpuTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.CpuTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLastRowsAffectedCount(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LastRowsAffectedCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLogicalIORead(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LogicalIORead = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetPhysicalIORead(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.PhysicalIORead = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetQueryTimeMS(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetRowsAffectedCount(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.RowsAffectedCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLHash(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLHash = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetUserName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.UserName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetWriteIOCount(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.WriteIOCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) Validate() error {
	return dara.Validate(s)
}

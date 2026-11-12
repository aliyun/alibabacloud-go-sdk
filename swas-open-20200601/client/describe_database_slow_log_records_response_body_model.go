// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeDatabaseSlowLogRecordsResponseBody
	GetEngine() *string
	SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsResponseBody
	GetPageSize() *int32
	SetPhysicalIORead(v int64) *DescribeDatabaseSlowLogRecordsResponseBody
	GetPhysicalIORead() *int64
	SetRequestId(v string) *DescribeDatabaseSlowLogRecordsResponseBody
	GetRequestId() *string
	SetSlowLogs(v []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) *DescribeDatabaseSlowLogRecordsResponseBody
	GetSlowLogs() []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs
	SetTotalCount(v int32) *DescribeDatabaseSlowLogRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeDatabaseSlowLogRecordsResponseBody struct {
	// The database engine.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The page number of the returned page.
	//
	// Starts from 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 30 to 100.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 200
	PhysicalIORead *int64 `json:"PhysicalIORead,omitempty" xml:"PhysicalIORead,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The slow query logs.
	SlowLogs []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs `json:"SlowLogs,omitempty" xml:"SlowLogs,omitempty" type:"Repeated"`
	// The total number of slow query logs.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetPhysicalIORead() *int64 {
	return s.PhysicalIORead
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetSlowLogs() []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	return s.SlowLogs
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetEngine(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPhysicalIORead(v int64) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PhysicalIORead = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetSlowLogs(v []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.SlowLogs = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetTotalCount(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) Validate() error {
	if s.SlowLogs != nil {
		for _, item := range s.SlowLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabaseSlowLogRecordsResponseBodySlowLogs struct {
	// The database name.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The start time of the execution. The time is displayed in UTC and is in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// > The time displayed in the console is in the UTC+8 time zone, which is 8 hours later than the time in UTC.
	//
	// example:
	//
	// 2022-09-08T01:40:44Z
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// The IP address of the client that is used to connect to the database.
	//
	// example:
	//
	// xxx[xxx] @ [1xx.xxx.xxx.xx]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The lock duration. Unit: seconds.
	//
	// example:
	//
	// 0
	LockTimes *int64 `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	// The number of parsed rows.
	//
	// example:
	//
	// 1
	ParseRowCounts *int64 `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// The execution duration. Unit: milliseconds.
	//
	// example:
	//
	// 2001
	QueryTimeMS *int64 `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// The execution duration. Unit: seconds.
	//
	// example:
	//
	// 2
	QueryTimes *int64 `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	// The number of returned rows.
	//
	// example:
	//
	// 1
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The details of the SQL statement.
	//
	// example:
	//
	// select sleep(2)
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetLockTimes() *int64 {
	return s.LockTimes
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetParseRowCounts() *int64 {
	return s.ParseRowCounts
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetQueryTimeMS() *int64 {
	return s.QueryTimeMS
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetQueryTimes() *int64 {
	return s.QueryTimes
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetDBName(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.DBName = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetExecutionStartTime(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetHostAddress(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.HostAddress = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetLockTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.LockTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetParseRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimeMS(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetReturnRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetSQLText(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.SQLText = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) Validate() error {
	return dara.Validate(s)
}

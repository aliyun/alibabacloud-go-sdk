// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody
	GetDBClusterId() *string
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
	// Cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Database engine.
	//
	// example:
	//
	// polardb_mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// List of slow log details.
	Items *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records on this page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of SQL statements.
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

func (s *DescribeSlowLogRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
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
	// Database name.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// Node ID.
	//
	// example:
	//
	// pi-*****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// Time when the SQL starts execution. The format is `YYYY-MM-DDThh:mmZ` (UTC time).
	//
	// example:
	//
	// 2021-04-07T03:47Z
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// Client address connecting to the database.
	//
	// example:
	//
	// testdb[testdb] @  [100.**.**.242]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// SQL lock duration in seconds.
	//
	// example:
	//
	// 0
	LockTimes *int64 `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	// Number of rows parsed.
	//
	// example:
	//
	// 0
	ParseRowCounts *int64 `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// Query time. Unit: milliseconds.
	//
	// example:
	//
	// 100
	QueryTimeMS *int64 `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// SQL execution duration, in seconds.
	//
	// example:
	//
	// 20
	QueryTimes *int64 `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	// Number of rows returned.
	//
	// example:
	//
	// 0
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// Unique identifier for the SQL statement in slow log statistics.
	//
	// example:
	//
	// U2FsdGVk****
	SQLHash *string `json:"SQLHash,omitempty" xml:"SQLHash,omitempty"`
	// Query statement.
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetLockTimes() *int64 {
	return s.LockTimes
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetParseRowCounts() *int64 {
	return s.ParseRowCounts
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

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetSQLHash() *string {
	return s.SQLHash
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBNodeId(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBNodeId = &v
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

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ParseRowCounts = &v
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

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLHash(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLHash = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) Validate() error {
	return dara.Validate(s)
}

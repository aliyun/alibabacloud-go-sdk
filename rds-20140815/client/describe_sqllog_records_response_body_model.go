// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSQLLogRecordsResponseBodyItems) *DescribeSQLLogRecordsResponseBody
	GetItems() *DescribeSQLLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeSQLLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSQLLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSQLLogRecordsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeSQLLogRecordsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeSQLLogRecordsResponseBody struct {
	// The details about each SQL audit log entry.
	Items *DescribeSQLLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of SQL audit log entries on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSQLLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBody) GetItems() *DescribeSQLLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSQLLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogRecordsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeSQLLogRecordsResponseBody) SetItems(v *DescribeSQLLogRecordsResponseBodyItems) *DescribeSQLLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetRequestId(v string) *DescribeSQLLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) SetTotalRecordCount(v int64) *DescribeSQLLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLLogRecordsResponseBodyItems struct {
	SQLRecord []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItems) GetSQLRecord() []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	return s.SQLRecord
}

func (s *DescribeSQLLogRecordsResponseBodyItems) SetSQLRecord(v []*DescribeSQLLogRecordsResponseBodyItemsSQLRecord) *DescribeSQLLogRecordsResponseBodyItems {
	s.SQLRecord = v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItems) Validate() error {
	if s.SQLRecord != nil {
		for _, item := range s.SQLRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogRecordsResponseBodyItemsSQLRecord struct {
	// The username of the account that is recorded in the SQL audit log entry.
	//
	// example:
	//
	// accounttest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The database name.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time at which the SQL statement was executed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2011-06-11T15:00:23Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP address of the client that is connected to the instance.
	//
	// example:
	//
	// 192.168.0.121
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The number of SQL audit log entries that are returned.
	//
	// example:
	//
	// 30
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// update test.zxb set id=0 limit 1
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// 1025865428
	ThreadID *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	// The execution duration of the SQL statement. Unit: microseconds.
	//
	// example:
	//
	// 600
	TotalExecutionTimes *int64 `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetThreadID() *string {
	return s.ThreadID
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) GetTotalExecutionTimes() *int64 {
	return s.TotalExecutionTimes
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetAccountName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetDBName(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetExecuteTime(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetHostAddress(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetSQLText(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetThreadID(v string) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.ThreadID = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeSQLLogRecordsResponseBodyItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseBodyItemsSQLRecord) Validate() error {
	return dara.Validate(s)
}

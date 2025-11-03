// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody
	GetItems() *DescribeAuditRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeAuditRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAuditRecordsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAuditRecordsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAuditRecordsResponseBody struct {
	// An array that consists of the information of audit log entries.
	Items *DescribeAuditRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3278BEB8-503B-4E46-8F7E-D26E040C9769
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 40
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAuditRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBody) GetItems() *DescribeAuditRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeAuditRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuditRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAuditRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAuditRecordsResponseBody) SetItems(v *DescribeAuditRecordsResponseBodyItems) *DescribeAuditRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageNumber(v int32) *DescribeAuditRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetPageRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetRequestId(v string) *DescribeAuditRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeAuditRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAuditRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAuditRecordsResponseBodyItems struct {
	SQLRecord []*DescribeAuditRecordsResponseBodyItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" type:"Repeated"`
}

func (s DescribeAuditRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItems) GetSQLRecord() []*DescribeAuditRecordsResponseBodyItemsSQLRecord {
	return s.SQLRecord
}

func (s *DescribeAuditRecordsResponseBodyItems) SetSQLRecord(v []*DescribeAuditRecordsResponseBodyItemsSQLRecord) *DescribeAuditRecordsResponseBodyItems {
	s.SQLRecord = v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItems) Validate() error {
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

type DescribeAuditRecordsResponseBodyItemsSQLRecord struct {
	// The account of the database.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test123
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time when the statement was executed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-11T03:30:27Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP addresses of the client.
	//
	// example:
	//
	// 11.xxx.xxx.xxx
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The number of SQL audit log entries that are returned.
	//
	// example:
	//
	// 2
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The statement that was executed.
	//
	// example:
	//
	// { \\"atype\\" : \\"createCollection\\", \\"param\\" : { \\"ns\\" : \\"123.test1\\" }, \\"result\\": \\"OK\\" }
	Syntax *string `json:"Syntax,omitempty" xml:"Syntax,omitempty"`
	// The name of the collection.
	//
	// example:
	//
	// C1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the thread.
	//
	// example:
	//
	// 140682188297984
	ThreadID *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	// The duration of the statement execution. Unit: microseconds.
	//
	// example:
	//
	// 700
	TotalExecutionTimes *int64 `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty"`
}

func (s DescribeAuditRecordsResponseBodyItemsSQLRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponseBodyItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetSyntax() *string {
	return s.Syntax
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetThreadID() *string {
	return s.ThreadID
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) GetTotalExecutionTimes() *int64 {
	return s.TotalExecutionTimes
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetAccountName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.AccountName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetDBName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetExecuteTime(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetHostAddress(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetSyntax(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.Syntax = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetTableName(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.TableName = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetThreadID(v string) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.ThreadID = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeAuditRecordsResponseBodyItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeAuditRecordsResponseBodyItemsSQLRecord) Validate() error {
	return dara.Validate(s)
}

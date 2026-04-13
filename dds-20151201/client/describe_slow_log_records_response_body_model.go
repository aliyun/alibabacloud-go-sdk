// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	// The database engine.
	//
	// example:
	//
	// MongoDB
	Engine *string                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Items  *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page. The value is a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of slow query log entries returned on the page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8076C4BA-DDBD-529C-BFF4-D8620C3F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	LogRecords []*DescribeSlowLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) GetLogRecords() []*DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	return s.LogRecords
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeSlowLogRecordsResponseBodyItemsLogRecords) *DescribeSlowLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItems) Validate() error {
	if s.LogRecords != nil {
		for _, item := range s.LogRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyItemsLogRecords struct {
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DocsExamined       *int64  `json:"DocsExamined,omitempty" xml:"DocsExamined,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	KeysExamined       *int64  `json:"KeysExamined,omitempty" xml:"KeysExamined,omitempty"`
	QueryTimes         *string `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	TableName          *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetDocsExamined() *int64 {
	return s.DocsExamined
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetKeysExamined() *int64 {
	return s.KeysExamined
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetQueryTimes() *string {
	return s.QueryTimes
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetAccountName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetDocsExamined(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.DocsExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetKeysExamined(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.KeysExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetQueryTimes(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) SetTableName(v string) *DescribeSlowLogRecordsResponseBodyItemsLogRecords {
	s.TableName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}

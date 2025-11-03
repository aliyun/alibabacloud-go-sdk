// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunningLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeRunningLogRecordsResponseBody
	GetEngine() *string
	SetItems(v *DescribeRunningLogRecordsResponseBodyItems) *DescribeRunningLogRecordsResponseBody
	GetItems() *DescribeRunningLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeRunningLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeRunningLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeRunningLogRecordsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeRunningLogRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeRunningLogRecordsResponseBody struct {
	// The database engine.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Details about the operational log entries.
	Items *DescribeRunningLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45D2B592-DEBA-4347-BBF3-47FF6C97DBBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRunningLogRecordsResponseBody) GetItems() *DescribeRunningLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeRunningLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRunningLogRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeRunningLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRunningLogRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeRunningLogRecordsResponseBody) SetEngine(v string) *DescribeRunningLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetItems(v *DescribeRunningLogRecordsResponseBodyItems) *DescribeRunningLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetPageNumber(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetRequestId(v string) *DescribeRunningLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeRunningLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRunningLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeRunningLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeRunningLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItems) GetLogRecords() []*DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	return s.LogRecords
}

func (s *DescribeRunningLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeRunningLogRecordsResponseBodyItemsLogRecords) *DescribeRunningLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItems) Validate() error {
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

type DescribeRunningLogRecordsResponseBodyItemsLogRecords struct {
	// The category of the log entry.
	//
	// example:
	//
	// NETWORK
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The connection information of the log entry.
	//
	// example:
	//
	// conn18xxxxxx
	ConnInfo *string `json:"ConnInfo,omitempty" xml:"ConnInfo,omitempty"`
	// The content of the log entry.
	//
	// example:
	//
	// end connection 11.xxx.xxx.xx:3xxxx (0 connections now open)\\n
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the log entry was generated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-26T12:09:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetCategory() *string {
	return s.Category
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetConnInfo() *string {
	return s.ConnInfo
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetContent() *string {
	return s.Content
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCategory(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Category = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetConnInfo(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.ConnInfo = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}

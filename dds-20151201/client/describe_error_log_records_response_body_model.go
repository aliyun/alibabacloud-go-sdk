// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeErrorLogRecordsResponseBody
	GetEngine() *string
	SetItems(v *DescribeErrorLogRecordsResponseBodyItems) *DescribeErrorLogRecordsResponseBody
	GetItems() *DescribeErrorLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeErrorLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeErrorLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeErrorLogRecordsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeErrorLogRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeErrorLogRecordsResponseBody struct {
	// The database engine.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Details about the log entries returned.
	Items *DescribeErrorLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68BCBEC2-1E66-471F-A1A8-E3C60C0A80B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeErrorLogRecordsResponseBody) GetItems() *DescribeErrorLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeErrorLogRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeErrorLogRecordsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeErrorLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeErrorLogRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeErrorLogRecordsResponseBody) SetEngine(v string) *DescribeErrorLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetItems(v *DescribeErrorLogRecordsResponseBodyItems) *DescribeErrorLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetPageNumber(v int32) *DescribeErrorLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeErrorLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetRequestId(v string) *DescribeErrorLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeErrorLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeErrorLogRecordsResponseBodyItems struct {
	LogRecords []*DescribeErrorLogRecordsResponseBodyItemsLogRecords `json:"LogRecords,omitempty" xml:"LogRecords,omitempty" type:"Repeated"`
}

func (s DescribeErrorLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyItems) GetLogRecords() []*DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	return s.LogRecords
}

func (s *DescribeErrorLogRecordsResponseBodyItems) SetLogRecords(v []*DescribeErrorLogRecordsResponseBodyItemsLogRecords) *DescribeErrorLogRecordsResponseBodyItems {
	s.LogRecords = v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItems) Validate() error {
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

type DescribeErrorLogRecordsResponseBodyItemsLogRecords struct {
	// The category of the log entry. Valid values:
	//
	// 	- NETWORK: network connection log
	//
	// 	- ACCESS: access control log
	//
	// 	- \\-: general log
	//
	// 	- COMMAND: slow query log
	//
	// 	- SHARDING: sharded cluster log
	//
	// 	- STORAGE: storage engine log
	//
	// 	- CONNPOOL: connection pool log
	//
	// 	- ASIO: asynchronous I/O operation log
	//
	// 	- WRITE: slow update log
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
	// xxxxxxxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the log entry was generated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-26T12:09:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the log entry.
	//
	// example:
	//
	// 1111111111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) GetCategory() *string {
	return s.Category
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) GetConnInfo() *string {
	return s.ConnInfo
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) GetContent() *string {
	return s.Content
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) GetId() *int32 {
	return s.Id
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetCategory(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Category = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetConnInfo(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.ConnInfo = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) SetId(v int32) *DescribeErrorLogRecordsResponseBodyItemsLogRecords {
	s.Id = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}

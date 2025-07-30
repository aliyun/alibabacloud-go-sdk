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
	SetInstanceId(v string) *DescribeRunningLogRecordsResponseBody
	GetInstanceId() *string
	SetItems(v *DescribeRunningLogRecordsResponseBodyItems) *DescribeRunningLogRecordsResponseBody
	GetItems() *DescribeRunningLogRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeRunningLogRecordsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeRunningLogRecordsResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeRunningLogRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRunningLogRecordsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRunningLogRecordsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeRunningLogRecordsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeRunningLogRecordsResponseBody struct {
	// The type of the database engine.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Details about the log entries.
	Items *DescribeRunningLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries returned on the current page.
	//
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The maximum number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 093B8579-9264-43A0-ABA9-AA86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2018-12-03T07:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
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

func (s *DescribeRunningLogRecordsResponseBody) GetInstanceId() *string {
	return s.InstanceId
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

func (s *DescribeRunningLogRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRunningLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRunningLogRecordsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRunningLogRecordsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeRunningLogRecordsResponseBody) SetEngine(v string) *DescribeRunningLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetInstanceId(v string) *DescribeRunningLogRecordsResponseBody {
	s.InstanceId = &v
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

func (s *DescribeRunningLogRecordsResponseBody) SetPageSize(v int32) *DescribeRunningLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetRequestId(v string) *DescribeRunningLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetStartTime(v string) *DescribeRunningLogRecordsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeRunningLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeRunningLogRecordsResponseBodyItemsLogRecords struct {
	// The content of the log.
	//
	// example:
	//
	// CONFIG REWRITE executed with success.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the log was generated. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-03T07:07:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// >  If a standard instance is queried, `(null)` is returned.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsResponseBodyItemsLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetContent() *string {
	return s.Content
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetContent(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.Content = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetCreateTime(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetInstanceId(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) SetNodeId(v string) *DescribeRunningLogRecordsResponseBodyItemsLogRecords {
	s.NodeId = &v
	return s
}

func (s *DescribeRunningLogRecordsResponseBodyItemsLogRecords) Validate() error {
	return dara.Validate(s)
}

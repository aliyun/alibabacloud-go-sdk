// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeErrorLogRecordsResponseBody
	GetCode() *int64
	SetData(v *DescribeErrorLogRecordsResponseBodyData) *DescribeErrorLogRecordsResponseBody
	GetData() *DescribeErrorLogRecordsResponseBodyData
	SetMessage(v string) *DescribeErrorLogRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeErrorLogRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeErrorLogRecordsResponseBody
	GetSuccess() *bool
}

type DescribeErrorLogRecordsResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeErrorLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAA17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeErrorLogRecordsResponseBody) GetData() *DescribeErrorLogRecordsResponseBodyData {
	return s.Data
}

func (s *DescribeErrorLogRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeErrorLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeErrorLogRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeErrorLogRecordsResponseBody) SetCode(v int64) *DescribeErrorLogRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetData(v *DescribeErrorLogRecordsResponseBodyData) *DescribeErrorLogRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetMessage(v string) *DescribeErrorLogRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetRequestId(v string) *DescribeErrorLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) SetSuccess(v bool) *DescribeErrorLogRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeErrorLogRecordsResponseBodyData struct {
	// example:
	//
	// 2025-07-23T05:48:43Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	ItemsNumbers *int64                                         `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	Logs         []*DescribeErrorLogRecordsResponseBodyDataLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// example:
	//
	// 2025-07-22T05:48:43Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetItemsNumbers() *int64 {
	return s.ItemsNumbers
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetLogs() []*DescribeErrorLogRecordsResponseBodyDataLogs {
	return s.Logs
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeErrorLogRecordsResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetEndTime(v string) *DescribeErrorLogRecordsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetItemsNumbers(v int64) *DescribeErrorLogRecordsResponseBodyData {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetLogs(v []*DescribeErrorLogRecordsResponseBodyDataLogs) *DescribeErrorLogRecordsResponseBodyData {
	s.Logs = v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetMaxRecordsPerPage(v int32) *DescribeErrorLogRecordsResponseBodyData {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetPageNumbers(v int32) *DescribeErrorLogRecordsResponseBodyData {
	s.PageNumbers = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetStartTime(v string) *DescribeErrorLogRecordsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) SetTotalRecords(v int64) *DescribeErrorLogRecordsResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyData) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeErrorLogRecordsResponseBodyDataLogs struct {
	// example:
	//
	// NETWORK
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// conn18xxxxxx
	ConnInfo *string `json:"ConnInfo,omitempty" xml:"ConnInfo,omitempty"`
	// example:
	//
	// 2025-07-15T15:14:27.175188+08:00 0 [Note] [MY-012468] [InnoDB] Transactions deadlock detected, dumping detailed information.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1731983067000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// d-bp128a003436****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeErrorLogRecordsResponseBodyDataLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponseBodyDataLogs) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) GetCategory() *string {
	return s.Category
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) GetConnInfo() *string {
	return s.ConnInfo
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) GetContent() *string {
	return s.Content
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) SetCategory(v string) *DescribeErrorLogRecordsResponseBodyDataLogs {
	s.Category = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) SetConnInfo(v string) *DescribeErrorLogRecordsResponseBodyDataLogs {
	s.ConnInfo = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) SetContent(v string) *DescribeErrorLogRecordsResponseBodyDataLogs {
	s.Content = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) SetCreateTime(v string) *DescribeErrorLogRecordsResponseBodyDataLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) SetDBInstanceName(v string) *DescribeErrorLogRecordsResponseBodyDataLogs {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeErrorLogRecordsResponseBodyDataLogs) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeErrorLogsResponseBodyItems) *DescribeErrorLogsResponseBody
	GetItems() *DescribeErrorLogsResponseBodyItems
	SetPageNumber(v int32) *DescribeErrorLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeErrorLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeErrorLogsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeErrorLogsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeErrorLogsResponseBody struct {
	// Details about the log entries returned.
	Items *DescribeErrorLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of error logs on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98504E07-BB0E-40FC-B152-E4882615812C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeErrorLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogsResponseBody) GetItems() *DescribeErrorLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeErrorLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeErrorLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeErrorLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeErrorLogsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeErrorLogsResponseBody) SetItems(v *DescribeErrorLogsResponseBodyItems) *DescribeErrorLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeErrorLogsResponseBody) SetPageNumber(v int32) *DescribeErrorLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogsResponseBody) SetPageRecordCount(v int32) *DescribeErrorLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeErrorLogsResponseBody) SetRequestId(v string) *DescribeErrorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeErrorLogsResponseBody) SetTotalRecordCount(v int32) *DescribeErrorLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeErrorLogsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeErrorLogsResponseBodyItems struct {
	ErrorLog []*DescribeErrorLogsResponseBodyItemsErrorLog `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty" type:"Repeated"`
}

func (s DescribeErrorLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogsResponseBodyItems) GetErrorLog() []*DescribeErrorLogsResponseBodyItemsErrorLog {
	return s.ErrorLog
}

func (s *DescribeErrorLogsResponseBodyItems) SetErrorLog(v []*DescribeErrorLogsResponseBodyItemsErrorLog) *DescribeErrorLogsResponseBodyItems {
	s.ErrorLog = v
	return s
}

func (s *DescribeErrorLogsResponseBodyItems) Validate() error {
	if s.ErrorLog != nil {
		for _, item := range s.ErrorLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeErrorLogsResponseBodyItemsErrorLog struct {
	// The time when the error log entry was generated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2011-05-30T12:11:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Database   *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The error log information.
	//
	// example:
	//
	// spid52 DBCC TRACEON 3499, server process ID (SPID) 52. This is an informational message only; no user action is required
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	User      *string `json:"User,omitempty" xml:"User,omitempty"`
	UserIp    *string `json:"UserIp,omitempty" xml:"UserIp,omitempty"`
}

func (s DescribeErrorLogsResponseBodyItemsErrorLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogsResponseBodyItemsErrorLog) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) GetDatabase() *string {
	return s.Database
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) GetUser() *string {
	return s.User
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) GetUserIp() *string {
	return s.UserIp
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) SetCreateTime(v string) *DescribeErrorLogsResponseBodyItemsErrorLog {
	s.CreateTime = &v
	return s
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) SetDatabase(v string) *DescribeErrorLogsResponseBodyItemsErrorLog {
	s.Database = &v
	return s
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) SetErrorInfo(v string) *DescribeErrorLogsResponseBodyItemsErrorLog {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) SetUser(v string) *DescribeErrorLogsResponseBodyItemsErrorLog {
	s.User = &v
	return s
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) SetUserIp(v string) *DescribeErrorLogsResponseBodyItemsErrorLog {
	s.UserIp = &v
	return s
}

func (s *DescribeErrorLogsResponseBodyItemsErrorLog) Validate() error {
	return dara.Validate(s)
}

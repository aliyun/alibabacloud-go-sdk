// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseErrorLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorLogs(v []*DescribeDatabaseErrorLogsResponseBodyErrorLogs) *DescribeDatabaseErrorLogsResponseBody
	GetErrorLogs() []*DescribeDatabaseErrorLogsResponseBodyErrorLogs
	SetPageNumber(v int32) *DescribeDatabaseErrorLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseErrorLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDatabaseErrorLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDatabaseErrorLogsResponseBody
	GetTotalCount() *int32
}

type DescribeDatabaseErrorLogsResponseBody struct {
	// An array of error logs.
	ErrorLogs []*DescribeDatabaseErrorLogsResponseBodyErrorLogs `json:"ErrorLogs,omitempty" xml:"ErrorLogs,omitempty" type:"Repeated"`
	// The page number.
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching error logs.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBody) GetErrorLogs() []*DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	return s.ErrorLogs
}

func (s *DescribeDatabaseErrorLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseErrorLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseErrorLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseErrorLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetErrorLogs(v []*DescribeDatabaseErrorLogsResponseBodyErrorLogs) *DescribeDatabaseErrorLogsResponseBody {
	s.ErrorLogs = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageNumber(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageSize(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetRequestId(v string) *DescribeDatabaseErrorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetTotalCount(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) Validate() error {
	if s.ErrorLogs != nil {
		for _, item := range s.ErrorLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDatabaseErrorLogsResponseBodyErrorLogs struct {
	// The time when the error log was generated. The time is in UTC and formatted as `yyyy-MM-ddTHH:mm:ssZ` according to the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// > The time displayed in the console is 8 hours ahead of the specified UTC time.
	//
	// example:
	//
	// 2022-09-08T12:11:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error information.
	//
	// example:
	//
	// spid52 DBCC TRACEON 3499, server process ID (SPID) 52. This is an informational message only; no user action is required
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetCreateTime(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetErrorInfo(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) Validate() error {
	return dara.Validate(s)
}

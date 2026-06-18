// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *ListTraceTasksRequest
	GetClientIp() *string
	SetDiagnoseId(v string) *ListTraceTasksRequest
	GetDiagnoseId() *string
	SetDomainName(v string) *ListTraceTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *ListTraceTasksRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListTraceTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTraceTasksRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListTraceTasksRequest
	GetStartTime() *string
	SetTaskId(v string) *ListTraceTasksRequest
	GetTaskId() *string
	SetTraceId(v string) *ListTraceTasksRequest
	GetTraceId() *string
}

type ListTraceTasksRequest struct {
	// The client IP address.
	//
	// example:
	//
	// 60.xx.xxx.38
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The diagnose ID.
	//
	// example:
	//
	// f2xxx5
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The domain name to diagnose.
	//
	// example:
	//
	// http://www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Time zone: UTC/GMT.
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: any integer from 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time. Time zone: UTC/GMT.
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The diagnostic task ID.
	//
	// example:
	//
	// xxxxxxxxxx-x-x-xxxxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The diagnostic trace ID.
	//
	// example:
	//
	// 000000xxxxxxxxxxxxxxxxxxxxxx33427e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListTraceTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTraceTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTraceTasksRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListTraceTasksRequest) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *ListTraceTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListTraceTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTraceTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTraceTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTraceTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTraceTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTraceTasksRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *ListTraceTasksRequest) SetClientIp(v string) *ListTraceTasksRequest {
	s.ClientIp = &v
	return s
}

func (s *ListTraceTasksRequest) SetDiagnoseId(v string) *ListTraceTasksRequest {
	s.DiagnoseId = &v
	return s
}

func (s *ListTraceTasksRequest) SetDomainName(v string) *ListTraceTasksRequest {
	s.DomainName = &v
	return s
}

func (s *ListTraceTasksRequest) SetEndTime(v string) *ListTraceTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListTraceTasksRequest) SetPageNumber(v int64) *ListTraceTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTraceTasksRequest) SetPageSize(v int64) *ListTraceTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTraceTasksRequest) SetStartTime(v string) *ListTraceTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListTraceTasksRequest) SetTaskId(v string) *ListTraceTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListTraceTasksRequest) SetTraceId(v string) *ListTraceTasksRequest {
	s.TraceId = &v
	return s
}

func (s *ListTraceTasksRequest) Validate() error {
	return dara.Validate(s)
}

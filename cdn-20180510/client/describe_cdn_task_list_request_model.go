// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *DescribeCdnTaskListRequest
	GetClientIp() *string
	SetDiagnoseId(v string) *DescribeCdnTaskListRequest
	GetDiagnoseId() *string
	SetDomainName(v string) *DescribeCdnTaskListRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnTaskListRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeCdnTaskListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnTaskListRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeCdnTaskListRequest
	GetStartTime() *string
	SetTaskId(v string) *DescribeCdnTaskListRequest
	GetTaskId() *string
	SetTraceId(v string) *DescribeCdnTaskListRequest
	GetTraceId() *string
}

type DescribeCdnTaskListRequest struct {
	// example:
	//
	// 60.xx.xxx.38
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// f2xxx5
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// example:
	//
	// http://www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1644467126
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1644467126
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxxxxxxxxx-x-x-xxxxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 000000xxxxxxxxxxxxxxxxxxxxxx33427e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeCdnTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnTaskListRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeCdnTaskListRequest) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DescribeCdnTaskListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnTaskListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnTaskListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnTaskListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnTaskListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnTaskListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCdnTaskListRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeCdnTaskListRequest) SetClientIp(v string) *DescribeCdnTaskListRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetDiagnoseId(v string) *DescribeCdnTaskListRequest {
	s.DiagnoseId = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetDomainName(v string) *DescribeCdnTaskListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetEndTime(v string) *DescribeCdnTaskListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetPageNumber(v int64) *DescribeCdnTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetPageSize(v int64) *DescribeCdnTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetStartTime(v string) *DescribeCdnTaskListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetTaskId(v string) *DescribeCdnTaskListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCdnTaskListRequest) SetTraceId(v string) *DescribeCdnTaskListRequest {
	s.TraceId = &v
	return s
}

func (s *DescribeCdnTaskListRequest) Validate() error {
	return dara.Validate(s)
}

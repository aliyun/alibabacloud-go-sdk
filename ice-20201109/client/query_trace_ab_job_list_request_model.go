// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *QueryTraceAbJobListRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *QueryTraceAbJobListRequest
	GetCreateTimeStart() *int64
	SetJobId(v string) *QueryTraceAbJobListRequest
	GetJobId() *string
	SetPageNumber(v int64) *QueryTraceAbJobListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *QueryTraceAbJobListRequest
	GetPageSize() *int64
	SetTraceMediaId(v string) *QueryTraceAbJobListRequest
	GetTraceMediaId() *string
}

type QueryTraceAbJobListRequest struct {
	// example:
	//
	// 1627357325
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 1627357322
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ****437bd2b51105d07b12a9****
	TraceMediaId *string `json:"TraceMediaId,omitempty" xml:"TraceMediaId,omitempty"`
}

func (s QueryTraceAbJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobListRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *QueryTraceAbJobListRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *QueryTraceAbJobListRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceAbJobListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryTraceAbJobListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTraceAbJobListRequest) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *QueryTraceAbJobListRequest) SetCreateTimeEnd(v int64) *QueryTraceAbJobListRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryTraceAbJobListRequest) SetCreateTimeStart(v int64) *QueryTraceAbJobListRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryTraceAbJobListRequest) SetJobId(v string) *QueryTraceAbJobListRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbJobListRequest) SetPageNumber(v int64) *QueryTraceAbJobListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTraceAbJobListRequest) SetPageSize(v int64) *QueryTraceAbJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTraceAbJobListRequest) SetTraceMediaId(v string) *QueryTraceAbJobListRequest {
	s.TraceMediaId = &v
	return s
}

func (s *QueryTraceAbJobListRequest) Validate() error {
	return dara.Validate(s)
}

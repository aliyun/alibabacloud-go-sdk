// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *QueryTraceM3u8JobRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *QueryTraceM3u8JobRequest
	GetCreateTimeStart() *int64
	SetJobId(v string) *QueryTraceM3u8JobRequest
	GetJobId() *string
	SetPageNumber(v int64) *QueryTraceM3u8JobRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *QueryTraceM3u8JobRequest
	GetPageSize() *int64
}

type QueryTraceM3u8JobRequest struct {
	// example:
	//
	// 1527441303
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 1527441300
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTraceM3u8JobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *QueryTraceM3u8JobRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *QueryTraceM3u8JobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceM3u8JobRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryTraceM3u8JobRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTraceM3u8JobRequest) SetCreateTimeEnd(v int64) *QueryTraceM3u8JobRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryTraceM3u8JobRequest) SetCreateTimeStart(v int64) *QueryTraceM3u8JobRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryTraceM3u8JobRequest) SetJobId(v string) *QueryTraceM3u8JobRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceM3u8JobRequest) SetPageNumber(v int64) *QueryTraceM3u8JobRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTraceM3u8JobRequest) SetPageSize(v int64) *QueryTraceM3u8JobRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTraceM3u8JobRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *QueryTraceM3u8JobListRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *QueryTraceM3u8JobListRequest
	GetCreateTimeStart() *int64
	SetJobId(v string) *QueryTraceM3u8JobListRequest
	GetJobId() *string
	SetPageNumber(v int64) *QueryTraceM3u8JobListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *QueryTraceM3u8JobListRequest
	GetPageSize() *int64
}

type QueryTraceM3u8JobListRequest struct {
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
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTraceM3u8JobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobListRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobListRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *QueryTraceM3u8JobListRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *QueryTraceM3u8JobListRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceM3u8JobListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryTraceM3u8JobListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTraceM3u8JobListRequest) SetCreateTimeEnd(v int64) *QueryTraceM3u8JobListRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryTraceM3u8JobListRequest) SetCreateTimeStart(v int64) *QueryTraceM3u8JobListRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryTraceM3u8JobListRequest) SetJobId(v string) *QueryTraceM3u8JobListRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceM3u8JobListRequest) SetPageNumber(v int64) *QueryTraceM3u8JobListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTraceM3u8JobListRequest) SetPageSize(v int64) *QueryTraceM3u8JobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTraceM3u8JobListRequest) Validate() error {
	return dara.Validate(s)
}

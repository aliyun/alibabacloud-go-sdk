// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *QueryCopyrightJobRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *QueryCopyrightJobRequest
	GetCreateTimeStart() *int64
	SetJobId(v string) *QueryCopyrightJobRequest
	GetJobId() *string
	SetLevel(v int64) *QueryCopyrightJobRequest
	GetLevel() *int64
	SetPageNumber(v int64) *QueryCopyrightJobRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *QueryCopyrightJobRequest
	GetPageSize() *int64
}

type QueryCopyrightJobRequest struct {
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
	// 2a0697e35a7342859f733a9190c4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 0
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCopyrightJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobRequest) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *QueryCopyrightJobRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *QueryCopyrightJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryCopyrightJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *QueryCopyrightJobRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryCopyrightJobRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryCopyrightJobRequest) SetCreateTimeEnd(v int64) *QueryCopyrightJobRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryCopyrightJobRequest) SetCreateTimeStart(v int64) *QueryCopyrightJobRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryCopyrightJobRequest) SetJobId(v string) *QueryCopyrightJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightJobRequest) SetLevel(v int64) *QueryCopyrightJobRequest {
	s.Level = &v
	return s
}

func (s *QueryCopyrightJobRequest) SetPageNumber(v int64) *QueryCopyrightJobRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCopyrightJobRequest) SetPageSize(v int64) *QueryCopyrightJobRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCopyrightJobRequest) Validate() error {
	return dara.Validate(s)
}

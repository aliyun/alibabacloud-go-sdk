// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *QueryCopyrightJobListRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *QueryCopyrightJobListRequest
	GetCreateTimeStart() *int64
	SetJobId(v string) *QueryCopyrightJobListRequest
	GetJobId() *string
	SetLevel(v int64) *QueryCopyrightJobListRequest
	GetLevel() *int64
	SetPageNumber(v int64) *QueryCopyrightJobListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *QueryCopyrightJobListRequest
	GetPageSize() *int64
}

type QueryCopyrightJobListRequest struct {
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
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 0
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCopyrightJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *QueryCopyrightJobListRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *QueryCopyrightJobListRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryCopyrightJobListRequest) GetLevel() *int64 {
	return s.Level
}

func (s *QueryCopyrightJobListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *QueryCopyrightJobListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryCopyrightJobListRequest) SetCreateTimeEnd(v int64) *QueryCopyrightJobListRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryCopyrightJobListRequest) SetCreateTimeStart(v int64) *QueryCopyrightJobListRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryCopyrightJobListRequest) SetJobId(v string) *QueryCopyrightJobListRequest {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightJobListRequest) SetLevel(v int64) *QueryCopyrightJobListRequest {
	s.Level = &v
	return s
}

func (s *QueryCopyrightJobListRequest) SetPageNumber(v int64) *QueryCopyrightJobListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCopyrightJobListRequest) SetPageSize(v int64) *QueryCopyrightJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCopyrightJobListRequest) Validate() error {
	return dara.Validate(s)
}

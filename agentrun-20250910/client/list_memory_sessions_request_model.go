// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *ListMemorySessionsRequest
	GetFrom() *int64
	SetPageNumber(v int32) *ListMemorySessionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemorySessionsRequest
	GetPageSize() *int32
	SetTo(v int64) *ListMemorySessionsRequest
	GetTo() *int64
}

type ListMemorySessionsRequest struct {
	// example:
	//
	// 1740622996
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1747275768
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s ListMemorySessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsRequest) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListMemorySessionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemorySessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemorySessionsRequest) GetTo() *int64 {
	return s.To
}

func (s *ListMemorySessionsRequest) SetFrom(v int64) *ListMemorySessionsRequest {
	s.From = &v
	return s
}

func (s *ListMemorySessionsRequest) SetPageNumber(v int32) *ListMemorySessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMemorySessionsRequest) SetPageSize(v int32) *ListMemorySessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMemorySessionsRequest) SetTo(v int64) *ListMemorySessionsRequest {
	s.To = &v
	return s
}

func (s *ListMemorySessionsRequest) Validate() error {
	return dara.Validate(s)
}

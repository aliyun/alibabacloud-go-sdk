// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *ListMemoryEventRequest
	GetFrom() *int64
	SetPageNumber(v int32) *ListMemoryEventRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemoryEventRequest
	GetPageSize() *int32
	SetTo(v int64) *ListMemoryEventRequest
	GetTo() *int64
}

type ListMemoryEventRequest struct {
	// example:
	//
	// 1742347023
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1736561650
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s ListMemoryEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryEventRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryEventRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListMemoryEventRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryEventRequest) GetTo() *int64 {
	return s.To
}

func (s *ListMemoryEventRequest) SetFrom(v int64) *ListMemoryEventRequest {
	s.From = &v
	return s
}

func (s *ListMemoryEventRequest) SetPageNumber(v int32) *ListMemoryEventRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryEventRequest) SetPageSize(v int32) *ListMemoryEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListMemoryEventRequest) SetTo(v int64) *ListMemoryEventRequest {
	s.To = &v
	return s
}

func (s *ListMemoryEventRequest) Validate() error {
	return dara.Validate(s)
}

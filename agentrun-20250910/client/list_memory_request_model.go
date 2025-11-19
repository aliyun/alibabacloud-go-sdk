// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListMemoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemoryRequest
	GetPageSize() *int32
	SetPattern(v string) *ListMemoryRequest
	GetPattern() *string
}

type ListMemoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
}

func (s ListMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryRequest) GetPattern() *string {
	return s.Pattern
}

func (s *ListMemoryRequest) SetPageNumber(v int32) *ListMemoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryRequest) SetPageSize(v int32) *ListMemoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListMemoryRequest) SetPattern(v string) *ListMemoryRequest {
	s.Pattern = &v
	return s
}

func (s *ListMemoryRequest) Validate() error {
	return dara.Validate(s)
}

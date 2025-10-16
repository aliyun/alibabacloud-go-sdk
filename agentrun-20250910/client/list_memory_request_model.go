// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamePrefix(v string) *ListMemoryRequest
	GetNamePrefix() *string
	SetPageNumber(v int32) *ListMemoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMemoryRequest
	GetPageSize() *int32
}

type ListMemoryRequest struct {
	// example:
	//
	// test
	NamePrefix *string `json:"namePrefix,omitempty" xml:"namePrefix,omitempty"`
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
}

func (s ListMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *ListMemoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryRequest) SetNamePrefix(v string) *ListMemoryRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListMemoryRequest) SetPageNumber(v int32) *ListMemoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryRequest) SetPageSize(v int32) *ListMemoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListMemoryRequest) Validate() error {
	return dara.Validate(s)
}

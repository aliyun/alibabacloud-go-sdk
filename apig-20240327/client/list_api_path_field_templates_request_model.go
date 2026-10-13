// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPathFieldTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListApiPathFieldTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiPathFieldTemplatesRequest
	GetPageSize() *int32
}

type ListApiPathFieldTemplatesRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListApiPathFieldTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiPathFieldTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListApiPathFieldTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiPathFieldTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiPathFieldTemplatesRequest) SetPageNumber(v int32) *ListApiPathFieldTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApiPathFieldTemplatesRequest) SetPageSize(v int32) *ListApiPathFieldTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiPathFieldTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

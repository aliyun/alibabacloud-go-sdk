// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListJobTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobTemplatesRequest
	GetPageSize() *int32
}

type ListJobTemplatesRequest struct {
	// The page number. Starting value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

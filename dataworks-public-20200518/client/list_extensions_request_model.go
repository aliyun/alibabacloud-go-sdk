// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListExtensionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExtensionsRequest
	GetPageSize() *int32
}

type ListExtensionsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExtensionsRequest) GoString() string {
	return s.String()
}

func (s *ListExtensionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExtensionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExtensionsRequest) SetPageNumber(v int32) *ListExtensionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsRequest) SetPageSize(v int32) *ListExtensionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExtensionsRequest) Validate() error {
	return dara.Validate(s)
}

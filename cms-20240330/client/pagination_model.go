// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPagination interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *Pagination
	GetPageNumber() *int32
	SetPageSize(v int32) *Pagination
	GetPageSize() *int32
}

type Pagination struct {
	// The current page number.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s Pagination) String() string {
	return dara.Prettify(s)
}

func (s Pagination) GoString() string {
	return s.String()
}

func (s *Pagination) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *Pagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *Pagination) SetPageNumber(v int32) *Pagination {
	s.PageNumber = &v
	return s
}

func (s *Pagination) SetPageSize(v int32) *Pagination {
	s.PageSize = &v
	return s
}

func (s *Pagination) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaginationInfo interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *PaginationInfo
	GetLimit() *int32
	SetPage(v int32) *PaginationInfo
	GetPage() *int32
	SetTotal(v int32) *PaginationInfo
	GetTotal() *int32
	SetTotalPages(v int32) *PaginationInfo
	GetTotalPages() *int32
}

type PaginationInfo struct {
	Limit      *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	Page       *int32 `json:"page,omitempty" xml:"page,omitempty"`
	Total      *int32 `json:"total,omitempty" xml:"total,omitempty"`
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s PaginationInfo) String() string {
	return dara.Prettify(s)
}

func (s PaginationInfo) GoString() string {
	return s.String()
}

func (s *PaginationInfo) GetLimit() *int32 {
	return s.Limit
}

func (s *PaginationInfo) GetPage() *int32 {
	return s.Page
}

func (s *PaginationInfo) GetTotal() *int32 {
	return s.Total
}

func (s *PaginationInfo) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *PaginationInfo) SetLimit(v int32) *PaginationInfo {
	s.Limit = &v
	return s
}

func (s *PaginationInfo) SetPage(v int32) *PaginationInfo {
	s.Page = &v
	return s
}

func (s *PaginationInfo) SetTotal(v int32) *PaginationInfo {
	s.Total = &v
	return s
}

func (s *PaginationInfo) SetTotalPages(v int32) *PaginationInfo {
	s.TotalPages = &v
	return s
}

func (s *PaginationInfo) Validate() error {
	return dara.Validate(s)
}

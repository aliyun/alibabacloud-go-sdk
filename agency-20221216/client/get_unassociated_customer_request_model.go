// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnassociatedCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetUnassociatedCustomerRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetUnassociatedCustomerRequest
	GetPageSize() *int32
}

type GetUnassociatedCustomerRequest struct {
	// Pagination, current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Pagination, record number on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetUnassociatedCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerRequest) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetUnassociatedCustomerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUnassociatedCustomerRequest) SetCurrentPage(v int32) *GetUnassociatedCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetUnassociatedCustomerRequest) SetPageSize(v int32) *GetUnassociatedCustomerRequest {
	s.PageSize = &v
	return s
}

func (s *GetUnassociatedCustomerRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFootprintListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFootprintListRequest
	GetCode() *string
	SetCurrentPage(v int64) *GetFootprintListRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *GetFootprintListRequest
	GetPageSize() *int64
	SetProductType(v int64) *GetFootprintListRequest
	GetProductType() *int64
}

type GetFootprintListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The pagination parameter. The number of the page that starts from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetFootprintListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFootprintListRequest) GoString() string {
	return s.String()
}

func (s *GetFootprintListRequest) GetCode() *string {
	return s.Code
}

func (s *GetFootprintListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetFootprintListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetFootprintListRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetFootprintListRequest) SetCode(v string) *GetFootprintListRequest {
	s.Code = &v
	return s
}

func (s *GetFootprintListRequest) SetCurrentPage(v int64) *GetFootprintListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFootprintListRequest) SetPageSize(v int64) *GetFootprintListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFootprintListRequest) SetProductType(v int64) *GetFootprintListRequest {
	s.ProductType = &v
	return s
}

func (s *GetFootprintListRequest) Validate() error {
	return dara.Validate(s)
}

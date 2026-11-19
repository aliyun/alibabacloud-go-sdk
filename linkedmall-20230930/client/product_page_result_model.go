// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetNextCursor(v string) *ProductPageResult
	GetNextCursor() *string
	SetPageNumber(v int32) *ProductPageResult
	GetPageNumber() *int32
	SetPageSize(v int32) *ProductPageResult
	GetPageSize() *int32
	SetProducts(v []*Product) *ProductPageResult
	GetProducts() []*Product
	SetRequestId(v string) *ProductPageResult
	GetRequestId() *string
	SetTotal(v int32) *ProductPageResult
	GetTotal() *int32
}

type ProductPageResult struct {
	// The cursor value for the next page.
	//
	// example:
	//
	// 7091510713
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The product collection.
	Products []*Product `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total count.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ProductPageResult) String() string {
	return dara.Prettify(s)
}

func (s ProductPageResult) GoString() string {
	return s.String()
}

func (s *ProductPageResult) GetNextCursor() *string {
	return s.NextCursor
}

func (s *ProductPageResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ProductPageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ProductPageResult) GetProducts() []*Product {
	return s.Products
}

func (s *ProductPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *ProductPageResult) SetNextCursor(v string) *ProductPageResult {
	s.NextCursor = &v
	return s
}

func (s *ProductPageResult) SetPageNumber(v int32) *ProductPageResult {
	s.PageNumber = &v
	return s
}

func (s *ProductPageResult) SetPageSize(v int32) *ProductPageResult {
	s.PageSize = &v
	return s
}

func (s *ProductPageResult) SetProducts(v []*Product) *ProductPageResult {
	s.Products = v
	return s
}

func (s *ProductPageResult) SetRequestId(v string) *ProductPageResult {
	s.RequestId = &v
	return s
}

func (s *ProductPageResult) SetTotal(v int32) *ProductPageResult {
	s.Total = &v
	return s
}

func (s *ProductPageResult) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

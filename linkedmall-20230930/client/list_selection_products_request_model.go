// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSelectionProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSelectionProductsRequest
	GetPageSize() *int32
	SetPurchaserId(v string) *ListSelectionProductsRequest
	GetPurchaserId() *string
}

type ListSelectionProductsRequest struct {
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
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s ListSelectionProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionProductsRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSelectionProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSelectionProductsRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *ListSelectionProductsRequest) SetPageNumber(v int32) *ListSelectionProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSelectionProductsRequest) SetPageSize(v int32) *ListSelectionProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSelectionProductsRequest) SetPurchaserId(v string) *ListSelectionProductsRequest {
	s.PurchaserId = &v
	return s
}

func (s *ListSelectionProductsRequest) Validate() error {
	return dara.Validate(s)
}

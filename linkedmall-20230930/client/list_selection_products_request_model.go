// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *ListSelectionProductsRequest
	GetCursor() *string
	SetPageNumber(v int32) *ListSelectionProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSelectionProductsRequest
	GetPageSize() *int32
	SetPurchaserId(v string) *ListSelectionProductsRequest
	GetPurchaserId() *string
}

type ListSelectionProductsRequest struct {
	// The cursor.
	//
	// example:
	//
	// 134412333321
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The current page number. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the purchaser.
	//
	// This parameter is required.
	//
	// example:
	//
	// PID22000009
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s ListSelectionProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionProductsRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionProductsRequest) GetCursor() *string {
	return s.Cursor
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

func (s *ListSelectionProductsRequest) SetCursor(v string) *ListSelectionProductsRequest {
	s.Cursor = &v
	return s
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

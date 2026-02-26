// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderPageQuery interface {
	dara.Model
	String() string
	GoString() string
	SetOrderIdList(v []*string) *OrderPageQuery
	GetOrderIdList() []*string
	SetPageNumber(v int32) *OrderPageQuery
	GetPageNumber() *int32
	SetPageSize(v int32) *OrderPageQuery
	GetPageSize() *int32
	SetPurchaseOrderId(v string) *OrderPageQuery
	GetPurchaseOrderId() *string
}

type OrderPageQuery struct {
	OrderIdList []*string `json:"orderIdList,omitempty" xml:"orderIdList,omitempty" type:"Repeated"`
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
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
}

func (s OrderPageQuery) String() string {
	return dara.Prettify(s)
}

func (s OrderPageQuery) GoString() string {
	return s.String()
}

func (s *OrderPageQuery) GetOrderIdList() []*string {
	return s.OrderIdList
}

func (s *OrderPageQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *OrderPageQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OrderPageQuery) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *OrderPageQuery) SetOrderIdList(v []*string) *OrderPageQuery {
	s.OrderIdList = v
	return s
}

func (s *OrderPageQuery) SetPageNumber(v int32) *OrderPageQuery {
	s.PageNumber = &v
	return s
}

func (s *OrderPageQuery) SetPageSize(v int32) *OrderPageQuery {
	s.PageSize = &v
	return s
}

func (s *OrderPageQuery) SetPurchaseOrderId(v string) *OrderPageQuery {
	s.PurchaseOrderId = &v
	return s
}

func (s *OrderPageQuery) Validate() error {
	return dara.Validate(s)
}

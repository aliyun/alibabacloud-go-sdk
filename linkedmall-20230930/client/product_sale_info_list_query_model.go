// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSaleInfoListQuery interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v string) *ProductSaleInfoListQuery
	GetDivisionCode() *string
	SetProductIds(v []*string) *ProductSaleInfoListQuery
	GetProductIds() []*string
	SetPurchaserId(v string) *ProductSaleInfoListQuery
	GetPurchaserId() *string
}

type ProductSaleInfoListQuery struct {
	// Area code (recommended to use a five-level township/street-level address code)
	//
	// example:
	//
	// 330106109
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// Collection of product IDs to query, supporting batch queries of 1 to 10 items
	//
	// This parameter is required.
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
	// Purchaser ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 22000009
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s ProductSaleInfoListQuery) String() string {
	return dara.Prettify(s)
}

func (s ProductSaleInfoListQuery) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoListQuery) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *ProductSaleInfoListQuery) GetProductIds() []*string {
	return s.ProductIds
}

func (s *ProductSaleInfoListQuery) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *ProductSaleInfoListQuery) SetDivisionCode(v string) *ProductSaleInfoListQuery {
	s.DivisionCode = &v
	return s
}

func (s *ProductSaleInfoListQuery) SetProductIds(v []*string) *ProductSaleInfoListQuery {
	s.ProductIds = v
	return s
}

func (s *ProductSaleInfoListQuery) SetPurchaserId(v string) *ProductSaleInfoListQuery {
	s.PurchaserId = &v
	return s
}

func (s *ProductSaleInfoListQuery) Validate() error {
	return dara.Validate(s)
}

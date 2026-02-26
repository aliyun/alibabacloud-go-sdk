// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRenderProductDTO interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *OrderRenderProductDTO
	GetProductId() *string
	SetPurchaserId(v string) *OrderRenderProductDTO
	GetPurchaserId() *string
	SetQuantity(v int32) *OrderRenderProductDTO
	GetQuantity() *int32
	SetSkuId(v string) *OrderRenderProductDTO
	GetSkuId() *string
}

type OrderRenderProductDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// skuID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s OrderRenderProductDTO) String() string {
	return dara.Prettify(s)
}

func (s OrderRenderProductDTO) GoString() string {
	return s.String()
}

func (s *OrderRenderProductDTO) GetProductId() *string {
	return s.ProductId
}

func (s *OrderRenderProductDTO) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *OrderRenderProductDTO) GetQuantity() *int32 {
	return s.Quantity
}

func (s *OrderRenderProductDTO) GetSkuId() *string {
	return s.SkuId
}

func (s *OrderRenderProductDTO) SetProductId(v string) *OrderRenderProductDTO {
	s.ProductId = &v
	return s
}

func (s *OrderRenderProductDTO) SetPurchaserId(v string) *OrderRenderProductDTO {
	s.PurchaserId = &v
	return s
}

func (s *OrderRenderProductDTO) SetQuantity(v int32) *OrderRenderProductDTO {
	s.Quantity = &v
	return s
}

func (s *OrderRenderProductDTO) SetSkuId(v string) *OrderRenderProductDTO {
	s.SkuId = &v
	return s
}

func (s *OrderRenderProductDTO) Validate() error {
	return dara.Validate(s)
}

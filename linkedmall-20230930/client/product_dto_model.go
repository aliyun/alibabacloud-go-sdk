// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductDTO interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v int64) *ProductDTO
	GetPrice() *int64
	SetProductId(v string) *ProductDTO
	GetProductId() *string
	SetPurchaserId(v string) *ProductDTO
	GetPurchaserId() *string
	SetQuantity(v int32) *ProductDTO
	GetQuantity() *int32
	SetSkuId(v string) *ProductDTO
	GetSkuId() *string
}

type ProductDTO struct {
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
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
	// SKUID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s ProductDTO) String() string {
	return dara.Prettify(s)
}

func (s ProductDTO) GoString() string {
	return s.String()
}

func (s *ProductDTO) GetPrice() *int64 {
	return s.Price
}

func (s *ProductDTO) GetProductId() *string {
	return s.ProductId
}

func (s *ProductDTO) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *ProductDTO) GetQuantity() *int32 {
	return s.Quantity
}

func (s *ProductDTO) GetSkuId() *string {
	return s.SkuId
}

func (s *ProductDTO) SetPrice(v int64) *ProductDTO {
	s.Price = &v
	return s
}

func (s *ProductDTO) SetProductId(v string) *ProductDTO {
	s.ProductId = &v
	return s
}

func (s *ProductDTO) SetPurchaserId(v string) *ProductDTO {
	s.PurchaserId = &v
	return s
}

func (s *ProductDTO) SetQuantity(v int32) *ProductDTO {
	s.Quantity = &v
	return s
}

func (s *ProductDTO) SetSkuId(v string) *ProductDTO {
	s.SkuId = &v
	return s
}

func (s *ProductDTO) Validate() error {
	return dara.Validate(s)
}

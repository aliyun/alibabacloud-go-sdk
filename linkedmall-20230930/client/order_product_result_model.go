// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderProductResult interface {
	dara.Model
	String() string
	GoString() string
	SetCanSell(v bool) *OrderProductResult
	GetCanSell() *bool
	SetFeatures(v map[string]interface{}) *OrderProductResult
	GetFeatures() map[string]interface{}
	SetMessage(v string) *OrderProductResult
	GetMessage() *string
	SetPrice(v int64) *OrderProductResult
	GetPrice() *int64
	SetProductId(v string) *OrderProductResult
	GetProductId() *string
	SetProductPicUrl(v string) *OrderProductResult
	GetProductPicUrl() *string
	SetProductTitle(v string) *OrderProductResult
	GetProductTitle() *string
	SetPurchaserId(v string) *OrderProductResult
	GetPurchaserId() *string
	SetQuantity(v int32) *OrderProductResult
	GetQuantity() *int32
	SetSkuId(v string) *OrderProductResult
	GetSkuId() *string
	SetSkuTitle(v string) *OrderProductResult
	GetSkuTitle() *string
}

type OrderProductResult struct {
	// example:
	//
	// true
	CanSell  *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	Message  *string                `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i4/2216003305543/O1CN01bip3Un1qokG0
	ProductPicUrl *string `json:"productPicUrl,omitempty" xml:"productPicUrl,omitempty"`
	ProductTitle  *string `json:"productTitle,omitempty" xml:"productTitle,omitempty"`
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// example:
	//
	// 1
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// SKUID
	//
	// example:
	//
	// 6600****6737
	SkuId    *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s OrderProductResult) String() string {
	return dara.Prettify(s)
}

func (s OrderProductResult) GoString() string {
	return s.String()
}

func (s *OrderProductResult) GetCanSell() *bool {
	return s.CanSell
}

func (s *OrderProductResult) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *OrderProductResult) GetMessage() *string {
	return s.Message
}

func (s *OrderProductResult) GetPrice() *int64 {
	return s.Price
}

func (s *OrderProductResult) GetProductId() *string {
	return s.ProductId
}

func (s *OrderProductResult) GetProductPicUrl() *string {
	return s.ProductPicUrl
}

func (s *OrderProductResult) GetProductTitle() *string {
	return s.ProductTitle
}

func (s *OrderProductResult) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *OrderProductResult) GetQuantity() *int32 {
	return s.Quantity
}

func (s *OrderProductResult) GetSkuId() *string {
	return s.SkuId
}

func (s *OrderProductResult) GetSkuTitle() *string {
	return s.SkuTitle
}

func (s *OrderProductResult) SetCanSell(v bool) *OrderProductResult {
	s.CanSell = &v
	return s
}

func (s *OrderProductResult) SetFeatures(v map[string]interface{}) *OrderProductResult {
	s.Features = v
	return s
}

func (s *OrderProductResult) SetMessage(v string) *OrderProductResult {
	s.Message = &v
	return s
}

func (s *OrderProductResult) SetPrice(v int64) *OrderProductResult {
	s.Price = &v
	return s
}

func (s *OrderProductResult) SetProductId(v string) *OrderProductResult {
	s.ProductId = &v
	return s
}

func (s *OrderProductResult) SetProductPicUrl(v string) *OrderProductResult {
	s.ProductPicUrl = &v
	return s
}

func (s *OrderProductResult) SetProductTitle(v string) *OrderProductResult {
	s.ProductTitle = &v
	return s
}

func (s *OrderProductResult) SetPurchaserId(v string) *OrderProductResult {
	s.PurchaserId = &v
	return s
}

func (s *OrderProductResult) SetQuantity(v int32) *OrderProductResult {
	s.Quantity = &v
	return s
}

func (s *OrderProductResult) SetSkuId(v string) *OrderProductResult {
	s.SkuId = &v
	return s
}

func (s *OrderProductResult) SetSkuTitle(v string) *OrderProductResult {
	s.SkuTitle = &v
	return s
}

func (s *OrderProductResult) Validate() error {
	return dara.Validate(s)
}

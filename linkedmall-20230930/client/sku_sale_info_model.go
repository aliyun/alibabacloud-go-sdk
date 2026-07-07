// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkuSaleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCanNotSellReason(v string) *SkuSaleInfo
	GetCanNotSellReason() *string
	SetCanSell(v bool) *SkuSaleInfo
	GetCanSell() *bool
	SetDivisionCode(v string) *SkuSaleInfo
	GetDivisionCode() *string
	SetFuzzyQuantity(v string) *SkuSaleInfo
	GetFuzzyQuantity() *string
	SetMarkPrice(v int64) *SkuSaleInfo
	GetMarkPrice() *int64
	SetPrice(v int64) *SkuSaleInfo
	GetPrice() *int64
	SetProductId(v string) *SkuSaleInfo
	GetProductId() *string
	SetQuantity(v int64) *SkuSaleInfo
	GetQuantity() *int64
	SetShopId(v string) *SkuSaleInfo
	GetShopId() *string
	SetSkuId(v string) *SkuSaleInfo
	GetSkuId() *string
	SetSkuStatus(v string) *SkuSaleInfo
	GetSkuStatus() *string
	SetTitle(v string) *SkuSaleInfo
	GetTitle() *string
}

type SkuSaleInfo struct {
	// Reason for not being sellable
	//
	// example:
	//
	// 不可售
	CanNotSellReason *string `json:"canNotSellReason,omitempty" xml:"canNotSellReason,omitempty"`
	// Indicates whether the SKU is sellable
	//
	// example:
	//
	// true
	CanSell *bool `json:"canSell,omitempty" xml:"canSell,omitempty"`
	// Area code
	//
	// example:
	//
	// 330106109
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// Blur inventory availability
	//
	// example:
	//
	// 有货
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// Strikethrough price, in cents
	//
	// example:
	//
	// 999900
	MarkPrice *int64 `json:"markPrice,omitempty" xml:"markPrice,omitempty"`
	// Distributor purchase price, in cents
	//
	// example:
	//
	// 19800
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// Product ID
	//
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// Available inventory
	//
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Shop ID
	//
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// skuId
	//
	// example:
	//
	// 660460842235822081
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// SKU control status
	//
	// example:
	//
	// Online
	SkuStatus *string `json:"skuStatus,omitempty" xml:"skuStatus,omitempty"`
	// SKU title
	//
	// example:
	//
	// 天蓝色
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SkuSaleInfo) String() string {
	return dara.Prettify(s)
}

func (s SkuSaleInfo) GoString() string {
	return s.String()
}

func (s *SkuSaleInfo) GetCanNotSellReason() *string {
	return s.CanNotSellReason
}

func (s *SkuSaleInfo) GetCanSell() *bool {
	return s.CanSell
}

func (s *SkuSaleInfo) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *SkuSaleInfo) GetFuzzyQuantity() *string {
	return s.FuzzyQuantity
}

func (s *SkuSaleInfo) GetMarkPrice() *int64 {
	return s.MarkPrice
}

func (s *SkuSaleInfo) GetPrice() *int64 {
	return s.Price
}

func (s *SkuSaleInfo) GetProductId() *string {
	return s.ProductId
}

func (s *SkuSaleInfo) GetQuantity() *int64 {
	return s.Quantity
}

func (s *SkuSaleInfo) GetShopId() *string {
	return s.ShopId
}

func (s *SkuSaleInfo) GetSkuId() *string {
	return s.SkuId
}

func (s *SkuSaleInfo) GetSkuStatus() *string {
	return s.SkuStatus
}

func (s *SkuSaleInfo) GetTitle() *string {
	return s.Title
}

func (s *SkuSaleInfo) SetCanNotSellReason(v string) *SkuSaleInfo {
	s.CanNotSellReason = &v
	return s
}

func (s *SkuSaleInfo) SetCanSell(v bool) *SkuSaleInfo {
	s.CanSell = &v
	return s
}

func (s *SkuSaleInfo) SetDivisionCode(v string) *SkuSaleInfo {
	s.DivisionCode = &v
	return s
}

func (s *SkuSaleInfo) SetFuzzyQuantity(v string) *SkuSaleInfo {
	s.FuzzyQuantity = &v
	return s
}

func (s *SkuSaleInfo) SetMarkPrice(v int64) *SkuSaleInfo {
	s.MarkPrice = &v
	return s
}

func (s *SkuSaleInfo) SetPrice(v int64) *SkuSaleInfo {
	s.Price = &v
	return s
}

func (s *SkuSaleInfo) SetProductId(v string) *SkuSaleInfo {
	s.ProductId = &v
	return s
}

func (s *SkuSaleInfo) SetQuantity(v int64) *SkuSaleInfo {
	s.Quantity = &v
	return s
}

func (s *SkuSaleInfo) SetShopId(v string) *SkuSaleInfo {
	s.ShopId = &v
	return s
}

func (s *SkuSaleInfo) SetSkuId(v string) *SkuSaleInfo {
	s.SkuId = &v
	return s
}

func (s *SkuSaleInfo) SetSkuStatus(v string) *SkuSaleInfo {
	s.SkuStatus = &v
	return s
}

func (s *SkuSaleInfo) SetTitle(v string) *SkuSaleInfo {
	s.Title = &v
	return s
}

func (s *SkuSaleInfo) Validate() error {
	return dara.Validate(s)
}

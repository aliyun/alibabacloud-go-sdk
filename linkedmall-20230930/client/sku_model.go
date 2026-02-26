// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSku interface {
	dara.Model
	String() string
	GoString() string
	SetBarcode(v string) *Sku
	GetBarcode() *string
	SetCanSell(v bool) *Sku
	GetCanSell() *bool
	SetDiscountRetailPrice(v int64) *Sku
	GetDiscountRetailPrice() *int64
	SetDivisionCode(v string) *Sku
	GetDivisionCode() *string
	SetFuzzyQuantity(v string) *Sku
	GetFuzzyQuantity() *string
	SetMarkPrice(v int64) *Sku
	GetMarkPrice() *int64
	SetPicUrl(v string) *Sku
	GetPicUrl() *string
	SetPlatformPrice(v int64) *Sku
	GetPlatformPrice() *int64
	SetPrice(v int64) *Sku
	GetPrice() *int64
	SetProductId(v string) *Sku
	GetProductId() *string
	SetQuantity(v int64) *Sku
	GetQuantity() *int64
	SetRankValue(v int64) *Sku
	GetRankValue() *int64
	SetShopId(v string) *Sku
	GetShopId() *string
	SetSkuAlias(v string) *Sku
	GetSkuAlias() *string
	SetSkuId(v string) *Sku
	GetSkuId() *string
	SetSkuSpecs(v []*SkuSpec) *Sku
	GetSkuSpecs() []*SkuSpec
	SetSkuSpecsCode(v string) *Sku
	GetSkuSpecsCode() *string
	SetSkuStatus(v string) *Sku
	GetSkuStatus() *string
	SetSuggestedRetailPrice(v int64) *Sku
	GetSuggestedRetailPrice() *int64
	SetTitle(v string) *Sku
	GetTitle() *string
}

type Sku struct {
	// example:
	//
	// 6922454329176
	Barcode *string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// example:
	//
	// true
	CanSell             *bool  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	DiscountRetailPrice *int64 `json:"discountRetailPrice,omitempty" xml:"discountRetailPrice,omitempty"`
	// example:
	//
	// 110000
	DivisionCode  *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// example:
	//
	// 999900
	MarkPrice *int64 `json:"markPrice,omitempty" xml:"markPrice,omitempty"`
	// example:
	//
	// https:////img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2216003305543/O1CN010DEQCX1qokFYGRfPE_!!2216003305543.png
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// 999900
	PlatformPrice *int64 `json:"platformPrice,omitempty" xml:"platformPrice,omitempty"`
	// example:
	//
	// 19800
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 3
	RankValue *int64 `json:"rankValue,omitempty" xml:"rankValue,omitempty"`
	// example:
	//
	// 21000017
	ShopId   *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	SkuAlias *string `json:"skuAlias,omitempty" xml:"skuAlias,omitempty"`
	// skuId
	//
	// example:
	//
	// 660460842235822081
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// sku
	SkuSpecs     []*SkuSpec `json:"skuSpecs,omitempty" xml:"skuSpecs,omitempty" type:"Repeated"`
	SkuSpecsCode *string    `json:"skuSpecsCode,omitempty" xml:"skuSpecsCode,omitempty"`
	// example:
	//
	// Online
	SkuStatus            *string `json:"skuStatus,omitempty" xml:"skuStatus,omitempty"`
	SuggestedRetailPrice *int64  `json:"suggestedRetailPrice,omitempty" xml:"suggestedRetailPrice,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Sku) String() string {
	return dara.Prettify(s)
}

func (s Sku) GoString() string {
	return s.String()
}

func (s *Sku) GetBarcode() *string {
	return s.Barcode
}

func (s *Sku) GetCanSell() *bool {
	return s.CanSell
}

func (s *Sku) GetDiscountRetailPrice() *int64 {
	return s.DiscountRetailPrice
}

func (s *Sku) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *Sku) GetFuzzyQuantity() *string {
	return s.FuzzyQuantity
}

func (s *Sku) GetMarkPrice() *int64 {
	return s.MarkPrice
}

func (s *Sku) GetPicUrl() *string {
	return s.PicUrl
}

func (s *Sku) GetPlatformPrice() *int64 {
	return s.PlatformPrice
}

func (s *Sku) GetPrice() *int64 {
	return s.Price
}

func (s *Sku) GetProductId() *string {
	return s.ProductId
}

func (s *Sku) GetQuantity() *int64 {
	return s.Quantity
}

func (s *Sku) GetRankValue() *int64 {
	return s.RankValue
}

func (s *Sku) GetShopId() *string {
	return s.ShopId
}

func (s *Sku) GetSkuAlias() *string {
	return s.SkuAlias
}

func (s *Sku) GetSkuId() *string {
	return s.SkuId
}

func (s *Sku) GetSkuSpecs() []*SkuSpec {
	return s.SkuSpecs
}

func (s *Sku) GetSkuSpecsCode() *string {
	return s.SkuSpecsCode
}

func (s *Sku) GetSkuStatus() *string {
	return s.SkuStatus
}

func (s *Sku) GetSuggestedRetailPrice() *int64 {
	return s.SuggestedRetailPrice
}

func (s *Sku) GetTitle() *string {
	return s.Title
}

func (s *Sku) SetBarcode(v string) *Sku {
	s.Barcode = &v
	return s
}

func (s *Sku) SetCanSell(v bool) *Sku {
	s.CanSell = &v
	return s
}

func (s *Sku) SetDiscountRetailPrice(v int64) *Sku {
	s.DiscountRetailPrice = &v
	return s
}

func (s *Sku) SetDivisionCode(v string) *Sku {
	s.DivisionCode = &v
	return s
}

func (s *Sku) SetFuzzyQuantity(v string) *Sku {
	s.FuzzyQuantity = &v
	return s
}

func (s *Sku) SetMarkPrice(v int64) *Sku {
	s.MarkPrice = &v
	return s
}

func (s *Sku) SetPicUrl(v string) *Sku {
	s.PicUrl = &v
	return s
}

func (s *Sku) SetPlatformPrice(v int64) *Sku {
	s.PlatformPrice = &v
	return s
}

func (s *Sku) SetPrice(v int64) *Sku {
	s.Price = &v
	return s
}

func (s *Sku) SetProductId(v string) *Sku {
	s.ProductId = &v
	return s
}

func (s *Sku) SetQuantity(v int64) *Sku {
	s.Quantity = &v
	return s
}

func (s *Sku) SetRankValue(v int64) *Sku {
	s.RankValue = &v
	return s
}

func (s *Sku) SetShopId(v string) *Sku {
	s.ShopId = &v
	return s
}

func (s *Sku) SetSkuAlias(v string) *Sku {
	s.SkuAlias = &v
	return s
}

func (s *Sku) SetSkuId(v string) *Sku {
	s.SkuId = &v
	return s
}

func (s *Sku) SetSkuSpecs(v []*SkuSpec) *Sku {
	s.SkuSpecs = v
	return s
}

func (s *Sku) SetSkuSpecsCode(v string) *Sku {
	s.SkuSpecsCode = &v
	return s
}

func (s *Sku) SetSkuStatus(v string) *Sku {
	s.SkuStatus = &v
	return s
}

func (s *Sku) SetSuggestedRetailPrice(v int64) *Sku {
	s.SuggestedRetailPrice = &v
	return s
}

func (s *Sku) SetTitle(v string) *Sku {
	s.Title = &v
	return s
}

func (s *Sku) Validate() error {
	if s.SkuSpecs != nil {
		for _, item := range s.SkuSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

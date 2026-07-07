// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSaleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCanSell(v bool) *ProductSaleInfo
	GetCanSell() *bool
	SetDivisionCode(v string) *ProductSaleInfo
	GetDivisionCode() *string
	SetFuzzyQuantity(v string) *ProductSaleInfo
	GetFuzzyQuantity() *string
	SetLimitRules(v []*LimitRule) *ProductSaleInfo
	GetLimitRules() []*LimitRule
	SetLmItemId(v string) *ProductSaleInfo
	GetLmItemId() *string
	SetProductId(v string) *ProductSaleInfo
	GetProductId() *string
	SetProductStatus(v string) *ProductSaleInfo
	GetProductStatus() *string
	SetQuantity(v int64) *ProductSaleInfo
	GetQuantity() *int64
	SetRequestId(v string) *ProductSaleInfo
	GetRequestId() *string
	SetShopId(v string) *ProductSaleInfo
	GetShopId() *string
	SetSkus(v []*SkuSaleInfo) *ProductSaleInfo
	GetSkus() []*SkuSaleInfo
	SetTitle(v string) *ProductSaleInfo
	GetTitle() *string
}

type ProductSaleInfo struct {
	// Is sellable, calculated value
	//
	// example:
	//
	// true
	CanSell *bool `json:"canSell,omitempty" xml:"canSell,omitempty"`
	// Region code
	//
	// example:
	//
	// 330106109
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// Fuzzy inventory quantity.
	//
	// example:
	//
	// 有货
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// Purchase limit configuration
	LimitRules []*LimitRule `json:"limitRules,omitempty" xml:"limitRules,omitempty" type:"Repeated"`
	// LM product ID
	//
	// example:
	//
	// 21000017-4580902812
	LmItemId *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	// Product ID
	//
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product status
	//
	// example:
	//
	// Online
	ProductStatus *string `json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// Inventory
	//
	// example:
	//
	// 10
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// API request ID
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Channel shop ID
	//
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// SKU collection
	Skus []*SkuSaleInfo `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	// Title
	//
	// example:
	//
	// 发财树
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ProductSaleInfo) String() string {
	return dara.Prettify(s)
}

func (s ProductSaleInfo) GoString() string {
	return s.String()
}

func (s *ProductSaleInfo) GetCanSell() *bool {
	return s.CanSell
}

func (s *ProductSaleInfo) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *ProductSaleInfo) GetFuzzyQuantity() *string {
	return s.FuzzyQuantity
}

func (s *ProductSaleInfo) GetLimitRules() []*LimitRule {
	return s.LimitRules
}

func (s *ProductSaleInfo) GetLmItemId() *string {
	return s.LmItemId
}

func (s *ProductSaleInfo) GetProductId() *string {
	return s.ProductId
}

func (s *ProductSaleInfo) GetProductStatus() *string {
	return s.ProductStatus
}

func (s *ProductSaleInfo) GetQuantity() *int64 {
	return s.Quantity
}

func (s *ProductSaleInfo) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductSaleInfo) GetShopId() *string {
	return s.ShopId
}

func (s *ProductSaleInfo) GetSkus() []*SkuSaleInfo {
	return s.Skus
}

func (s *ProductSaleInfo) GetTitle() *string {
	return s.Title
}

func (s *ProductSaleInfo) SetCanSell(v bool) *ProductSaleInfo {
	s.CanSell = &v
	return s
}

func (s *ProductSaleInfo) SetDivisionCode(v string) *ProductSaleInfo {
	s.DivisionCode = &v
	return s
}

func (s *ProductSaleInfo) SetFuzzyQuantity(v string) *ProductSaleInfo {
	s.FuzzyQuantity = &v
	return s
}

func (s *ProductSaleInfo) SetLimitRules(v []*LimitRule) *ProductSaleInfo {
	s.LimitRules = v
	return s
}

func (s *ProductSaleInfo) SetLmItemId(v string) *ProductSaleInfo {
	s.LmItemId = &v
	return s
}

func (s *ProductSaleInfo) SetProductId(v string) *ProductSaleInfo {
	s.ProductId = &v
	return s
}

func (s *ProductSaleInfo) SetProductStatus(v string) *ProductSaleInfo {
	s.ProductStatus = &v
	return s
}

func (s *ProductSaleInfo) SetQuantity(v int64) *ProductSaleInfo {
	s.Quantity = &v
	return s
}

func (s *ProductSaleInfo) SetRequestId(v string) *ProductSaleInfo {
	s.RequestId = &v
	return s
}

func (s *ProductSaleInfo) SetShopId(v string) *ProductSaleInfo {
	s.ShopId = &v
	return s
}

func (s *ProductSaleInfo) SetSkus(v []*SkuSaleInfo) *ProductSaleInfo {
	s.Skus = v
	return s
}

func (s *ProductSaleInfo) SetTitle(v string) *ProductSaleInfo {
	s.Title = &v
	return s
}

func (s *ProductSaleInfo) Validate() error {
	if s.LimitRules != nil {
		for _, item := range s.LimitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Skus != nil {
		for _, item := range s.Skus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

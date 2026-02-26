// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProduct interface {
	dara.Model
	String() string
	GoString() string
	SetBrandName(v string) *Product
	GetBrandName() *string
	SetCanSell(v bool) *Product
	GetCanSell() *bool
	SetCategoryChain(v []*Category) *Product
	GetCategoryChain() []*Category
	SetCategoryLeafId(v int64) *Product
	GetCategoryLeafId() *int64
	SetDescPath(v string) *Product
	GetDescPath() *string
	SetDivisionCode(v string) *Product
	GetDivisionCode() *string
	SetExtendProperties(v []*ProductExtendProperty) *Product
	GetExtendProperties() []*ProductExtendProperty
	SetFuzzyQuantity(v string) *Product
	GetFuzzyQuantity() *string
	SetImages(v []*string) *Product
	GetImages() []*string
	SetInGroup(v bool) *Product
	GetInGroup() *bool
	SetLimitRules(v []*LimitRule) *Product
	GetLimitRules() []*LimitRule
	SetLmItemId(v string) *Product
	GetLmItemId() *string
	SetPicUrl(v string) *Product
	GetPicUrl() *string
	SetProductId(v string) *Product
	GetProductId() *string
	SetProductSpecs(v []*ProductSpec) *Product
	GetProductSpecs() []*ProductSpec
	SetProductStatus(v string) *Product
	GetProductStatus() *string
	SetProductType(v string) *Product
	GetProductType() *string
	SetProperties(v []*ProductProperty) *Product
	GetProperties() []*ProductProperty
	SetQuantity(v int64) *Product
	GetQuantity() *int64
	SetRequestId(v string) *Product
	GetRequestId() *string
	SetServicePromises(v []*string) *Product
	GetServicePromises() []*string
	SetShopId(v string) *Product
	GetShopId() *string
	SetSkus(v []*Sku) *Product
	GetSkus() []*Sku
	SetSoldQuantity(v string) *Product
	GetSoldQuantity() *string
	SetTaxCode(v string) *Product
	GetTaxCode() *string
	SetTaxRate(v int32) *Product
	GetTaxRate() *int32
	SetTitle(v string) *Product
	GetTitle() *string
}

type Product struct {
	// example:
	//
	// Apple/苹果
	BrandName *string `json:"brandName,omitempty" xml:"brandName,omitempty"`
	// example:
	//
	// true
	CanSell       *bool       `json:"canSell,omitempty" xml:"canSell,omitempty"`
	CategoryChain []*Category `json:"categoryChain,omitempty" xml:"categoryChain,omitempty" type:"Repeated"`
	// example:
	//
	// 201856501
	CategoryLeafId *int64 `json:"categoryLeafId,omitempty" xml:"categoryLeafId,omitempty"`
	// example:
	//
	// https://img.alicdn.com/descpath/O1CN01wciRDp22AEU14435dsf34
	DescPath *string `json:"descPath,omitempty" xml:"descPath,omitempty"`
	// example:
	//
	// 110000
	DivisionCode     *string                  `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	ExtendProperties []*ProductExtendProperty `json:"extendProperties,omitempty" xml:"extendProperties,omitempty" type:"Repeated"`
	// example:
	//
	// 100+
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// images
	Images     []*string    `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	InGroup    *bool        `json:"inGroup,omitempty" xml:"inGroup,omitempty"`
	LimitRules []*LimitRule `json:"limitRules,omitempty" xml:"limitRules,omitempty" type:"Repeated"`
	// example:
	//
	// 21000017-4580902812
	LmItemId *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/2214281521988/O1CN01w4vomR1QYYEx6nyr5_!!2214281521988.jpg
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// productSpecs
	ProductSpecs []*ProductSpec `json:"productSpecs,omitempty" xml:"productSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// Online
	ProductStatus *string `json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// example:
	//
	// Normal
	ProductType *string            `json:"productType,omitempty" xml:"productType,omitempty"`
	Properties  []*ProductProperty `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId       *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServicePromises []*string `json:"servicePromises,omitempty" xml:"servicePromises,omitempty" type:"Repeated"`
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// skus
	Skus []*Sku `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	// example:
	//
	// 100+
	SoldQuantity *string `json:"soldQuantity,omitempty" xml:"soldQuantity,omitempty"`
	// example:
	//
	// 3040203000000000000
	TaxCode *string `json:"taxCode,omitempty" xml:"taxCode,omitempty"`
	// example:
	//
	// 600
	TaxRate *int32  `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Product) String() string {
	return dara.Prettify(s)
}

func (s Product) GoString() string {
	return s.String()
}

func (s *Product) GetBrandName() *string {
	return s.BrandName
}

func (s *Product) GetCanSell() *bool {
	return s.CanSell
}

func (s *Product) GetCategoryChain() []*Category {
	return s.CategoryChain
}

func (s *Product) GetCategoryLeafId() *int64 {
	return s.CategoryLeafId
}

func (s *Product) GetDescPath() *string {
	return s.DescPath
}

func (s *Product) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *Product) GetExtendProperties() []*ProductExtendProperty {
	return s.ExtendProperties
}

func (s *Product) GetFuzzyQuantity() *string {
	return s.FuzzyQuantity
}

func (s *Product) GetImages() []*string {
	return s.Images
}

func (s *Product) GetInGroup() *bool {
	return s.InGroup
}

func (s *Product) GetLimitRules() []*LimitRule {
	return s.LimitRules
}

func (s *Product) GetLmItemId() *string {
	return s.LmItemId
}

func (s *Product) GetPicUrl() *string {
	return s.PicUrl
}

func (s *Product) GetProductId() *string {
	return s.ProductId
}

func (s *Product) GetProductSpecs() []*ProductSpec {
	return s.ProductSpecs
}

func (s *Product) GetProductStatus() *string {
	return s.ProductStatus
}

func (s *Product) GetProductType() *string {
	return s.ProductType
}

func (s *Product) GetProperties() []*ProductProperty {
	return s.Properties
}

func (s *Product) GetQuantity() *int64 {
	return s.Quantity
}

func (s *Product) GetRequestId() *string {
	return s.RequestId
}

func (s *Product) GetServicePromises() []*string {
	return s.ServicePromises
}

func (s *Product) GetShopId() *string {
	return s.ShopId
}

func (s *Product) GetSkus() []*Sku {
	return s.Skus
}

func (s *Product) GetSoldQuantity() *string {
	return s.SoldQuantity
}

func (s *Product) GetTaxCode() *string {
	return s.TaxCode
}

func (s *Product) GetTaxRate() *int32 {
	return s.TaxRate
}

func (s *Product) GetTitle() *string {
	return s.Title
}

func (s *Product) SetBrandName(v string) *Product {
	s.BrandName = &v
	return s
}

func (s *Product) SetCanSell(v bool) *Product {
	s.CanSell = &v
	return s
}

func (s *Product) SetCategoryChain(v []*Category) *Product {
	s.CategoryChain = v
	return s
}

func (s *Product) SetCategoryLeafId(v int64) *Product {
	s.CategoryLeafId = &v
	return s
}

func (s *Product) SetDescPath(v string) *Product {
	s.DescPath = &v
	return s
}

func (s *Product) SetDivisionCode(v string) *Product {
	s.DivisionCode = &v
	return s
}

func (s *Product) SetExtendProperties(v []*ProductExtendProperty) *Product {
	s.ExtendProperties = v
	return s
}

func (s *Product) SetFuzzyQuantity(v string) *Product {
	s.FuzzyQuantity = &v
	return s
}

func (s *Product) SetImages(v []*string) *Product {
	s.Images = v
	return s
}

func (s *Product) SetInGroup(v bool) *Product {
	s.InGroup = &v
	return s
}

func (s *Product) SetLimitRules(v []*LimitRule) *Product {
	s.LimitRules = v
	return s
}

func (s *Product) SetLmItemId(v string) *Product {
	s.LmItemId = &v
	return s
}

func (s *Product) SetPicUrl(v string) *Product {
	s.PicUrl = &v
	return s
}

func (s *Product) SetProductId(v string) *Product {
	s.ProductId = &v
	return s
}

func (s *Product) SetProductSpecs(v []*ProductSpec) *Product {
	s.ProductSpecs = v
	return s
}

func (s *Product) SetProductStatus(v string) *Product {
	s.ProductStatus = &v
	return s
}

func (s *Product) SetProductType(v string) *Product {
	s.ProductType = &v
	return s
}

func (s *Product) SetProperties(v []*ProductProperty) *Product {
	s.Properties = v
	return s
}

func (s *Product) SetQuantity(v int64) *Product {
	s.Quantity = &v
	return s
}

func (s *Product) SetRequestId(v string) *Product {
	s.RequestId = &v
	return s
}

func (s *Product) SetServicePromises(v []*string) *Product {
	s.ServicePromises = v
	return s
}

func (s *Product) SetShopId(v string) *Product {
	s.ShopId = &v
	return s
}

func (s *Product) SetSkus(v []*Sku) *Product {
	s.Skus = v
	return s
}

func (s *Product) SetSoldQuantity(v string) *Product {
	s.SoldQuantity = &v
	return s
}

func (s *Product) SetTaxCode(v string) *Product {
	s.TaxCode = &v
	return s
}

func (s *Product) SetTaxRate(v int32) *Product {
	s.TaxRate = &v
	return s
}

func (s *Product) SetTitle(v string) *Product {
	s.Title = &v
	return s
}

func (s *Product) Validate() error {
	if s.CategoryChain != nil {
		for _, item := range s.CategoryChain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtendProperties != nil {
		for _, item := range s.ExtendProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LimitRules != nil {
		for _, item := range s.LimitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProductSpecs != nil {
		for _, item := range s.ProductSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Properties != nil {
		for _, item := range s.Properties {
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

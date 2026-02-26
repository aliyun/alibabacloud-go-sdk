// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*SearchProductsResponseBodyProducts) *SearchProductsResponseBody
	GetProducts() []*SearchProductsResponseBodyProducts
	SetTotal(v int32) *SearchProductsResponseBody
	GetTotal() *int32
}

type SearchProductsResponseBody struct {
	Products []*SearchProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchProductsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProductsResponseBody) GetProducts() []*SearchProductsResponseBodyProducts {
	return s.Products
}

func (s *SearchProductsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *SearchProductsResponseBody) SetProducts(v []*SearchProductsResponseBodyProducts) *SearchProductsResponseBody {
	s.Products = v
	return s
}

func (s *SearchProductsResponseBody) SetTotal(v int32) *SearchProductsResponseBody {
	s.Total = &v
	return s
}

func (s *SearchProductsResponseBody) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchProductsResponseBodyProducts struct {
	BandName         *string `json:"bandName,omitempty" xml:"bandName,omitempty"`
	CanNotSellReason *string `json:"canNotSellReason,omitempty" xml:"canNotSellReason,omitempty"`
	// example:
	//
	// true
	CanSell       *bool                                              `json:"canSell,omitempty" xml:"canSell,omitempty"`
	CategoryChain []*SearchProductsResponseBodyProductsCategoryChain `json:"categoryChain,omitempty" xml:"categoryChain,omitempty" type:"Repeated"`
	Credit        []*string                                          `json:"credit,omitempty" xml:"credit,omitempty" type:"Repeated"`
	// example:
	//
	// ￥-9998.95 ~ ￥-9977.90
	DiffPrice *string `json:"diffPrice,omitempty" xml:"diffPrice,omitempty"`
	// example:
	//
	// ￥0.05 ~ ￥21.10
	DistributionPrice *string `json:"distributionPrice,omitempty" xml:"distributionPrice,omitempty"`
	// example:
	//
	// -100.00% ~ -99.79%
	DistributionPriceRatio *string `json:"distributionPriceRatio,omitempty" xml:"distributionPriceRatio,omitempty"`
	// example:
	//
	// Taobao
	ExternalPlatformType *string `json:"externalPlatformType,omitempty" xml:"externalPlatformType,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// true
	InGroup *bool `json:"inGroup,omitempty" xml:"inGroup,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	InGroupTime *string `json:"inGroupTime,omitempty" xml:"inGroupTime,omitempty"`
	// example:
	//
	// Low
	InventoryRiskLevel *string `json:"inventoryRiskLevel,omitempty" xml:"inventoryRiskLevel,omitempty"`
	// example:
	//
	// HasInvoice
	InvoiceType *string `json:"invoiceType,omitempty" xml:"invoiceType,omitempty"`
	// example:
	//
	// xxx-xxxxx
	LmItemId *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	// example:
	//
	// https://img.alicdn.com/xxx.jpg
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// ￥9999.00 ~ ￥9999.00
	PlatformPrice *string `json:"platformPrice,omitempty" xml:"platformPrice,omitempty"`
	// example:
	//
	// ￥9999.00 ~ ￥9999.00
	PlatformReservePrice *string `json:"platformReservePrice,omitempty" xml:"platformReservePrice,omitempty"`
	// example:
	//
	// xxxxx
	ProductId   *string `json:"productId,omitempty" xml:"productId,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// xxx
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	// example:
	//
	// 100
	SoldQuantity *string `json:"soldQuantity,omitempty" xml:"soldQuantity,omitempty"`
	// example:
	//
	// 3040203000000000000
	TaxCode *string `json:"taxCode,omitempty" xml:"taxCode,omitempty"`
	TaxRate *int64  `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// example:
	//
	// JingXiao
	TradeMode *string `json:"tradeMode,omitempty" xml:"tradeMode,omitempty"`
}

func (s SearchProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s SearchProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *SearchProductsResponseBodyProducts) GetBandName() *string {
	return s.BandName
}

func (s *SearchProductsResponseBodyProducts) GetCanNotSellReason() *string {
	return s.CanNotSellReason
}

func (s *SearchProductsResponseBodyProducts) GetCanSell() *bool {
	return s.CanSell
}

func (s *SearchProductsResponseBodyProducts) GetCategoryChain() []*SearchProductsResponseBodyProductsCategoryChain {
	return s.CategoryChain
}

func (s *SearchProductsResponseBodyProducts) GetCredit() []*string {
	return s.Credit
}

func (s *SearchProductsResponseBodyProducts) GetDiffPrice() *string {
	return s.DiffPrice
}

func (s *SearchProductsResponseBodyProducts) GetDistributionPrice() *string {
	return s.DistributionPrice
}

func (s *SearchProductsResponseBodyProducts) GetDistributionPriceRatio() *string {
	return s.DistributionPriceRatio
}

func (s *SearchProductsResponseBodyProducts) GetExternalPlatformType() *string {
	return s.ExternalPlatformType
}

func (s *SearchProductsResponseBodyProducts) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SearchProductsResponseBodyProducts) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SearchProductsResponseBodyProducts) GetInGroup() *bool {
	return s.InGroup
}

func (s *SearchProductsResponseBodyProducts) GetInGroupTime() *string {
	return s.InGroupTime
}

func (s *SearchProductsResponseBodyProducts) GetInventoryRiskLevel() *string {
	return s.InventoryRiskLevel
}

func (s *SearchProductsResponseBodyProducts) GetInvoiceType() *string {
	return s.InvoiceType
}

func (s *SearchProductsResponseBodyProducts) GetLmItemId() *string {
	return s.LmItemId
}

func (s *SearchProductsResponseBodyProducts) GetPicUrl() *string {
	return s.PicUrl
}

func (s *SearchProductsResponseBodyProducts) GetPlatformPrice() *string {
	return s.PlatformPrice
}

func (s *SearchProductsResponseBodyProducts) GetPlatformReservePrice() *string {
	return s.PlatformReservePrice
}

func (s *SearchProductsResponseBodyProducts) GetProductId() *string {
	return s.ProductId
}

func (s *SearchProductsResponseBodyProducts) GetProductName() *string {
	return s.ProductName
}

func (s *SearchProductsResponseBodyProducts) GetShopName() *string {
	return s.ShopName
}

func (s *SearchProductsResponseBodyProducts) GetSoldQuantity() *string {
	return s.SoldQuantity
}

func (s *SearchProductsResponseBodyProducts) GetTaxCode() *string {
	return s.TaxCode
}

func (s *SearchProductsResponseBodyProducts) GetTaxRate() *int64 {
	return s.TaxRate
}

func (s *SearchProductsResponseBodyProducts) GetTradeMode() *string {
	return s.TradeMode
}

func (s *SearchProductsResponseBodyProducts) SetBandName(v string) *SearchProductsResponseBodyProducts {
	s.BandName = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetCanNotSellReason(v string) *SearchProductsResponseBodyProducts {
	s.CanNotSellReason = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetCanSell(v bool) *SearchProductsResponseBodyProducts {
	s.CanSell = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetCategoryChain(v []*SearchProductsResponseBodyProductsCategoryChain) *SearchProductsResponseBodyProducts {
	s.CategoryChain = v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetCredit(v []*string) *SearchProductsResponseBodyProducts {
	s.Credit = v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetDiffPrice(v string) *SearchProductsResponseBodyProducts {
	s.DiffPrice = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetDistributionPrice(v string) *SearchProductsResponseBodyProducts {
	s.DistributionPrice = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetDistributionPriceRatio(v string) *SearchProductsResponseBodyProducts {
	s.DistributionPriceRatio = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetExternalPlatformType(v string) *SearchProductsResponseBodyProducts {
	s.ExternalPlatformType = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetGmtCreate(v string) *SearchProductsResponseBodyProducts {
	s.GmtCreate = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetGmtModified(v string) *SearchProductsResponseBodyProducts {
	s.GmtModified = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetInGroup(v bool) *SearchProductsResponseBodyProducts {
	s.InGroup = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetInGroupTime(v string) *SearchProductsResponseBodyProducts {
	s.InGroupTime = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetInventoryRiskLevel(v string) *SearchProductsResponseBodyProducts {
	s.InventoryRiskLevel = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetInvoiceType(v string) *SearchProductsResponseBodyProducts {
	s.InvoiceType = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetLmItemId(v string) *SearchProductsResponseBodyProducts {
	s.LmItemId = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetPicUrl(v string) *SearchProductsResponseBodyProducts {
	s.PicUrl = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetPlatformPrice(v string) *SearchProductsResponseBodyProducts {
	s.PlatformPrice = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetPlatformReservePrice(v string) *SearchProductsResponseBodyProducts {
	s.PlatformReservePrice = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetProductId(v string) *SearchProductsResponseBodyProducts {
	s.ProductId = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetProductName(v string) *SearchProductsResponseBodyProducts {
	s.ProductName = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetShopName(v string) *SearchProductsResponseBodyProducts {
	s.ShopName = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetSoldQuantity(v string) *SearchProductsResponseBodyProducts {
	s.SoldQuantity = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetTaxCode(v string) *SearchProductsResponseBodyProducts {
	s.TaxCode = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetTaxRate(v int64) *SearchProductsResponseBodyProducts {
	s.TaxRate = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) SetTradeMode(v string) *SearchProductsResponseBodyProducts {
	s.TradeMode = &v
	return s
}

func (s *SearchProductsResponseBodyProducts) Validate() error {
	if s.CategoryChain != nil {
		for _, item := range s.CategoryChain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchProductsResponseBodyProductsCategoryChain struct {
	// example:
	//
	// 201792301
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	IsLeaf     *bool  `json:"isLeaf,omitempty" xml:"isLeaf,omitempty"`
	// example:
	//
	// 1
	Level *int32  `json:"level,omitempty" xml:"level,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s SearchProductsResponseBodyProductsCategoryChain) String() string {
	return dara.Prettify(s)
}

func (s SearchProductsResponseBodyProductsCategoryChain) GoString() string {
	return s.String()
}

func (s *SearchProductsResponseBodyProductsCategoryChain) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *SearchProductsResponseBodyProductsCategoryChain) GetIsLeaf() *bool {
	return s.IsLeaf
}

func (s *SearchProductsResponseBodyProductsCategoryChain) GetLevel() *int32 {
	return s.Level
}

func (s *SearchProductsResponseBodyProductsCategoryChain) GetName() *string {
	return s.Name
}

func (s *SearchProductsResponseBodyProductsCategoryChain) GetParentId() *int64 {
	return s.ParentId
}

func (s *SearchProductsResponseBodyProductsCategoryChain) SetCategoryId(v int64) *SearchProductsResponseBodyProductsCategoryChain {
	s.CategoryId = &v
	return s
}

func (s *SearchProductsResponseBodyProductsCategoryChain) SetIsLeaf(v bool) *SearchProductsResponseBodyProductsCategoryChain {
	s.IsLeaf = &v
	return s
}

func (s *SearchProductsResponseBodyProductsCategoryChain) SetLevel(v int32) *SearchProductsResponseBodyProductsCategoryChain {
	s.Level = &v
	return s
}

func (s *SearchProductsResponseBodyProductsCategoryChain) SetName(v string) *SearchProductsResponseBodyProductsCategoryChain {
	s.Name = &v
	return s
}

func (s *SearchProductsResponseBodyProductsCategoryChain) SetParentId(v int64) *SearchProductsResponseBodyProductsCategoryChain {
	s.ParentId = &v
	return s
}

func (s *SearchProductsResponseBodyProductsCategoryChain) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandName(v string) *SearchProductsRequest
	GetBrandName() *string
	SetCategoryIds(v []*string) *SearchProductsRequest
	GetCategoryIds() []*string
	SetCreateEndTime(v string) *SearchProductsRequest
	GetCreateEndTime() *string
	SetCreateStartTime(v string) *SearchProductsRequest
	GetCreateStartTime() *string
	SetDistributionHighPrice(v int64) *SearchProductsRequest
	GetDistributionHighPrice() *int64
	SetDistributionHighPriceRatio(v int64) *SearchProductsRequest
	GetDistributionHighPriceRatio() *int64
	SetDistributionLowPrice(v int64) *SearchProductsRequest
	GetDistributionLowPrice() *int64
	SetDistributionLowPriceRatio(v int64) *SearchProductsRequest
	GetDistributionLowPriceRatio() *int64
	SetHighMarkPrice(v int64) *SearchProductsRequest
	GetHighMarkPrice() *int64
	SetHighPrice(v int64) *SearchProductsRequest
	GetHighPrice() *int64
	SetInGroup(v bool) *SearchProductsRequest
	GetInGroup() *bool
	SetInGroupEndTime(v string) *SearchProductsRequest
	GetInGroupEndTime() *string
	SetInGroupStartTime(v string) *SearchProductsRequest
	GetInGroupStartTime() *string
	SetInventoryRiskLevel(v string) *SearchProductsRequest
	GetInventoryRiskLevel() *string
	SetLmItemId(v string) *SearchProductsRequest
	GetLmItemId() *string
	SetLowMarkPrice(v int64) *SearchProductsRequest
	GetLowMarkPrice() *int64
	SetLowPrice(v int64) *SearchProductsRequest
	GetLowPrice() *int64
	SetModifyEndTime(v string) *SearchProductsRequest
	GetModifyEndTime() *string
	SetModifyStartTime(v string) *SearchProductsRequest
	GetModifyStartTime() *string
	SetOrderBy(v string) *SearchProductsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchProductsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *SearchProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchProductsRequest
	GetPageSize() *int32
	SetPlatform(v string) *SearchProductsRequest
	GetPlatform() *string
	SetProductId(v string) *SearchProductsRequest
	GetProductId() *string
	SetProductName(v string) *SearchProductsRequest
	GetProductName() *string
	SetProductStatus(v string) *SearchProductsRequest
	GetProductStatus() *string
	SetPurchaserId(v string) *SearchProductsRequest
	GetPurchaserId() *string
	SetTaxRate(v string) *SearchProductsRequest
	GetTaxRate() *string
	SetTradeModeAndCredit(v string) *SearchProductsRequest
	GetTradeModeAndCredit() *string
}

type SearchProductsRequest struct {
	BrandName   *string   `json:"brandName,omitempty" xml:"brandName,omitempty"`
	CategoryIds []*string `json:"categoryIds,omitempty" xml:"categoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	CreateEndTime *string `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	CreateStartTime       *string `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	DistributionHighPrice *int64  `json:"distributionHighPrice,omitempty" xml:"distributionHighPrice,omitempty"`
	// example:
	//
	// 244（2.44%）
	DistributionHighPriceRatio *int64 `json:"distributionHighPriceRatio,omitempty" xml:"distributionHighPriceRatio,omitempty"`
	DistributionLowPrice       *int64 `json:"distributionLowPrice,omitempty" xml:"distributionLowPrice,omitempty"`
	// example:
	//
	// 133（1.33%）
	DistributionLowPriceRatio *int64 `json:"distributionLowPriceRatio,omitempty" xml:"distributionLowPriceRatio,omitempty"`
	HighMarkPrice             *int64 `json:"highMarkPrice,omitempty" xml:"highMarkPrice,omitempty"`
	HighPrice                 *int64 `json:"highPrice,omitempty" xml:"highPrice,omitempty"`
	// example:
	//
	// true
	InGroup *bool `json:"inGroup,omitempty" xml:"inGroup,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	InGroupEndTime *string `json:"inGroupEndTime,omitempty" xml:"inGroupEndTime,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	InGroupStartTime *string `json:"inGroupStartTime,omitempty" xml:"inGroupStartTime,omitempty"`
	// example:
	//
	// Low
	InventoryRiskLevel *string `json:"inventoryRiskLevel,omitempty" xml:"inventoryRiskLevel,omitempty"`
	// example:
	//
	// xxx-xxxxx
	LmItemId     *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	LowMarkPrice *int64  `json:"lowMarkPrice,omitempty" xml:"lowMarkPrice,omitempty"`
	LowPrice     *int64  `json:"lowPrice,omitempty" xml:"lowPrice,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	ModifyEndTime *string `json:"modifyEndTime,omitempty" xml:"modifyEndTime,omitempty"`
	// example:
	//
	// 2025-01-02 12:23:34
	//
	// (yyyy-MM-dd HH:mm:ss)
	ModifyStartTime *string `json:"modifyStartTime,omitempty" xml:"modifyStartTime,omitempty"`
	OrderBy         *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// ASC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
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
	// Taobao
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// xxxxxxx
	ProductId   *string `json:"productId,omitempty" xml:"productId,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// Sellable
	ProductStatus *string `json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PIDxxxx
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// example:
	//
	// Rate0
	TaxRate *string `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// example:
	//
	// JingXiao
	TradeModeAndCredit *string `json:"tradeModeAndCredit,omitempty" xml:"tradeModeAndCredit,omitempty"`
}

func (s SearchProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchProductsRequest) GoString() string {
	return s.String()
}

func (s *SearchProductsRequest) GetBrandName() *string {
	return s.BrandName
}

func (s *SearchProductsRequest) GetCategoryIds() []*string {
	return s.CategoryIds
}

func (s *SearchProductsRequest) GetCreateEndTime() *string {
	return s.CreateEndTime
}

func (s *SearchProductsRequest) GetCreateStartTime() *string {
	return s.CreateStartTime
}

func (s *SearchProductsRequest) GetDistributionHighPrice() *int64 {
	return s.DistributionHighPrice
}

func (s *SearchProductsRequest) GetDistributionHighPriceRatio() *int64 {
	return s.DistributionHighPriceRatio
}

func (s *SearchProductsRequest) GetDistributionLowPrice() *int64 {
	return s.DistributionLowPrice
}

func (s *SearchProductsRequest) GetDistributionLowPriceRatio() *int64 {
	return s.DistributionLowPriceRatio
}

func (s *SearchProductsRequest) GetHighMarkPrice() *int64 {
	return s.HighMarkPrice
}

func (s *SearchProductsRequest) GetHighPrice() *int64 {
	return s.HighPrice
}

func (s *SearchProductsRequest) GetInGroup() *bool {
	return s.InGroup
}

func (s *SearchProductsRequest) GetInGroupEndTime() *string {
	return s.InGroupEndTime
}

func (s *SearchProductsRequest) GetInGroupStartTime() *string {
	return s.InGroupStartTime
}

func (s *SearchProductsRequest) GetInventoryRiskLevel() *string {
	return s.InventoryRiskLevel
}

func (s *SearchProductsRequest) GetLmItemId() *string {
	return s.LmItemId
}

func (s *SearchProductsRequest) GetLowMarkPrice() *int64 {
	return s.LowMarkPrice
}

func (s *SearchProductsRequest) GetLowPrice() *int64 {
	return s.LowPrice
}

func (s *SearchProductsRequest) GetModifyEndTime() *string {
	return s.ModifyEndTime
}

func (s *SearchProductsRequest) GetModifyStartTime() *string {
	return s.ModifyStartTime
}

func (s *SearchProductsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchProductsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchProductsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *SearchProductsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *SearchProductsRequest) GetProductName() *string {
	return s.ProductName
}

func (s *SearchProductsRequest) GetProductStatus() *string {
	return s.ProductStatus
}

func (s *SearchProductsRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *SearchProductsRequest) GetTaxRate() *string {
	return s.TaxRate
}

func (s *SearchProductsRequest) GetTradeModeAndCredit() *string {
	return s.TradeModeAndCredit
}

func (s *SearchProductsRequest) SetBrandName(v string) *SearchProductsRequest {
	s.BrandName = &v
	return s
}

func (s *SearchProductsRequest) SetCategoryIds(v []*string) *SearchProductsRequest {
	s.CategoryIds = v
	return s
}

func (s *SearchProductsRequest) SetCreateEndTime(v string) *SearchProductsRequest {
	s.CreateEndTime = &v
	return s
}

func (s *SearchProductsRequest) SetCreateStartTime(v string) *SearchProductsRequest {
	s.CreateStartTime = &v
	return s
}

func (s *SearchProductsRequest) SetDistributionHighPrice(v int64) *SearchProductsRequest {
	s.DistributionHighPrice = &v
	return s
}

func (s *SearchProductsRequest) SetDistributionHighPriceRatio(v int64) *SearchProductsRequest {
	s.DistributionHighPriceRatio = &v
	return s
}

func (s *SearchProductsRequest) SetDistributionLowPrice(v int64) *SearchProductsRequest {
	s.DistributionLowPrice = &v
	return s
}

func (s *SearchProductsRequest) SetDistributionLowPriceRatio(v int64) *SearchProductsRequest {
	s.DistributionLowPriceRatio = &v
	return s
}

func (s *SearchProductsRequest) SetHighMarkPrice(v int64) *SearchProductsRequest {
	s.HighMarkPrice = &v
	return s
}

func (s *SearchProductsRequest) SetHighPrice(v int64) *SearchProductsRequest {
	s.HighPrice = &v
	return s
}

func (s *SearchProductsRequest) SetInGroup(v bool) *SearchProductsRequest {
	s.InGroup = &v
	return s
}

func (s *SearchProductsRequest) SetInGroupEndTime(v string) *SearchProductsRequest {
	s.InGroupEndTime = &v
	return s
}

func (s *SearchProductsRequest) SetInGroupStartTime(v string) *SearchProductsRequest {
	s.InGroupStartTime = &v
	return s
}

func (s *SearchProductsRequest) SetInventoryRiskLevel(v string) *SearchProductsRequest {
	s.InventoryRiskLevel = &v
	return s
}

func (s *SearchProductsRequest) SetLmItemId(v string) *SearchProductsRequest {
	s.LmItemId = &v
	return s
}

func (s *SearchProductsRequest) SetLowMarkPrice(v int64) *SearchProductsRequest {
	s.LowMarkPrice = &v
	return s
}

func (s *SearchProductsRequest) SetLowPrice(v int64) *SearchProductsRequest {
	s.LowPrice = &v
	return s
}

func (s *SearchProductsRequest) SetModifyEndTime(v string) *SearchProductsRequest {
	s.ModifyEndTime = &v
	return s
}

func (s *SearchProductsRequest) SetModifyStartTime(v string) *SearchProductsRequest {
	s.ModifyStartTime = &v
	return s
}

func (s *SearchProductsRequest) SetOrderBy(v string) *SearchProductsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchProductsRequest) SetOrderDirection(v string) *SearchProductsRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchProductsRequest) SetPageNumber(v int32) *SearchProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchProductsRequest) SetPageSize(v int32) *SearchProductsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchProductsRequest) SetPlatform(v string) *SearchProductsRequest {
	s.Platform = &v
	return s
}

func (s *SearchProductsRequest) SetProductId(v string) *SearchProductsRequest {
	s.ProductId = &v
	return s
}

func (s *SearchProductsRequest) SetProductName(v string) *SearchProductsRequest {
	s.ProductName = &v
	return s
}

func (s *SearchProductsRequest) SetProductStatus(v string) *SearchProductsRequest {
	s.ProductStatus = &v
	return s
}

func (s *SearchProductsRequest) SetPurchaserId(v string) *SearchProductsRequest {
	s.PurchaserId = &v
	return s
}

func (s *SearchProductsRequest) SetTaxRate(v string) *SearchProductsRequest {
	s.TaxRate = &v
	return s
}

func (s *SearchProductsRequest) SetTradeModeAndCredit(v string) *SearchProductsRequest {
	s.TradeModeAndCredit = &v
	return s
}

func (s *SearchProductsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeProductsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProductsResponseBody
	GetPageSize() *int32
	SetProductItems(v *DescribeProductsResponseBodyProductItems) *DescribeProductsResponseBody
	GetProductItems() *DescribeProductsResponseBodyProductItems
	SetRequestId(v string) *DescribeProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeProductsResponseBody
	GetTotalCount() *int32
}

type DescribeProductsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductItems *DescribeProductsResponseBodyProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Struct"`
	// example:
	//
	// A077D99E-0C4D-421E-A5D4-F533F6657817
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 75
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProductsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProductsResponseBody) GetProductItems() *DescribeProductsResponseBodyProductItems {
	return s.ProductItems
}

func (s *DescribeProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProductsResponseBody) SetPageNumber(v int32) *DescribeProductsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductsResponseBody) SetPageSize(v int32) *DescribeProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProductsResponseBody) SetProductItems(v *DescribeProductsResponseBodyProductItems) *DescribeProductsResponseBody {
	s.ProductItems = v
	return s
}

func (s *DescribeProductsResponseBody) SetRequestId(v string) *DescribeProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductsResponseBody) SetTotalCount(v int32) *DescribeProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsResponseBodyProductItems struct {
	ProductItem []*DescribeProductsResponseBodyProductItemsProductItem `json:"ProductItem,omitempty" xml:"ProductItem,omitempty" type:"Repeated"`
}

func (s DescribeProductsResponseBodyProductItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyProductItems) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyProductItems) GetProductItem() []*DescribeProductsResponseBodyProductItemsProductItem {
	return s.ProductItem
}

func (s *DescribeProductsResponseBodyProductItems) SetProductItem(v []*DescribeProductsResponseBodyProductItemsProductItem) *DescribeProductsResponseBodyProductItems {
	s.ProductItem = v
	return s
}

func (s *DescribeProductsResponseBodyProductItems) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsResponseBodyProductItemsProductItem struct {
	// example:
	//
	// 53398003
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// cmjj02****
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DeliveryDate *string `json:"DeliveryDate,omitempty" xml:"DeliveryDate,omitempty"`
	DeliveryWay  *string `json:"DeliveryWay,omitempty" xml:"DeliveryWay,omitempty"`
	// example:
	//
	// https://oss.aliyuncs.com/photogallery/photo/1904996544835414/7549/767d6d07-8366-4822-b84e-61f6ea10d146.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// example:
	//
	// {\\"DiscountPrice\\": 0.0, \\"TradePrice\\": 15750.0, \\"Currency\\": \\"CNY\\", \\"OriginalPrice\\": 15750.0, \\"RuleIds\\": {\\"RuleId\\": []}, \\"Coupons\\": {\\"Coupon\\": []}}
	PriceInfo *string `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty"`
	// example:
	//
	// 5.0
	Score            *string `json:"Score,omitempty" xml:"Score,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	SuggestedPrice   *string `json:"SuggestedPrice,omitempty" xml:"SuggestedPrice,omitempty"`
	// example:
	//
	// 228399
	SupplierId   *int64  `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// /products/53616009/cmjj02****.html
	TargetUrl    *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	WarrantyDate *string `json:"WarrantyDate,omitempty" xml:"WarrantyDate,omitempty"`
}

func (s DescribeProductsResponseBodyProductItemsProductItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyProductItemsProductItem) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetCode() *string {
	return s.Code
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetDeliveryDate() *string {
	return s.DeliveryDate
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetDeliveryWay() *string {
	return s.DeliveryWay
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetName() *string {
	return s.Name
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetPriceInfo() *string {
	return s.PriceInfo
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetScore() *string {
	return s.Score
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetSuggestedPrice() *string {
	return s.SuggestedPrice
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetTags() *string {
	return s.Tags
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) GetWarrantyDate() *string {
	return s.WarrantyDate
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetCategoryId(v int64) *DescribeProductsResponseBodyProductItemsProductItem {
	s.CategoryId = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetCode(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Code = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetDeliveryDate(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.DeliveryDate = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetDeliveryWay(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.DeliveryWay = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetImageUrl(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.ImageUrl = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetName(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Name = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetOperationSystem(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.OperationSystem = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetPriceInfo(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.PriceInfo = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetScore(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Score = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetShortDescription(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.ShortDescription = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSuggestedPrice(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SuggestedPrice = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSupplierId(v int64) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SupplierId = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSupplierName(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SupplierName = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetTags(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Tags = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetTargetUrl(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.TargetUrl = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetWarrantyDate(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.WarrantyDate = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) Validate() error {
	return dara.Validate(s)
}

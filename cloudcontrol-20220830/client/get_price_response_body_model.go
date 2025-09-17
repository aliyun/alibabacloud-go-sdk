// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *GetPriceResponseBodyPrice) *GetPriceResponseBody
	GetPrice() *GetPriceResponseBodyPrice
	SetRequestId(v string) *GetPriceResponseBody
	GetRequestId() *string
}

type GetPriceResponseBody struct {
	// The price.
	Price *GetPriceResponseBodyPrice `json:"price,omitempty" xml:"price,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBody) GetPrice() *GetPriceResponseBodyPrice {
	return s.Price
}

func (s *GetPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPriceResponseBody) SetPrice(v *GetPriceResponseBodyPrice) *GetPriceResponseBody {
	s.Price = v
	return s
}

func (s *GetPriceResponseBody) SetRequestId(v string) *GetPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPriceResponseBodyPrice struct {
	// The currency type. Valid values: CNY: Chinese Yuan. USD: US dollar. JPY: Japanese Yen.
	//
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The discount.
	//
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"discountPrice,omitempty" xml:"discountPrice,omitempty"`
	// The order details of the pricing module.
	ModuleDetails []*GetPriceResponseBodyPriceModuleDetails `json:"moduleDetails,omitempty" xml:"moduleDetails,omitempty" type:"Repeated"`
	// The original price.
	//
	// example:
	//
	// 760.0
	OriginalPrice *float32 `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	// The details of the promotion.
	PromotionDetails []*GetPriceResponseBodyPricePromotionDetails `json:"promotionDetails,omitempty" xml:"promotionDetails,omitempty" type:"Repeated"`
	// The discount price.
	//
	// example:
	//
	// 0.0
	TradePrice *float32 `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
}

func (s GetPriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s GetPriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetPriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *GetPriceResponseBodyPrice) GetModuleDetails() []*GetPriceResponseBodyPriceModuleDetails {
	return s.ModuleDetails
}

func (s *GetPriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *GetPriceResponseBodyPrice) GetPromotionDetails() []*GetPriceResponseBodyPricePromotionDetails {
	return s.PromotionDetails
}

func (s *GetPriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *GetPriceResponseBodyPrice) SetCurrency(v string) *GetPriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetDiscountPrice(v float32) *GetPriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetModuleDetails(v []*GetPriceResponseBodyPriceModuleDetails) *GetPriceResponseBodyPrice {
	s.ModuleDetails = v
	return s
}

func (s *GetPriceResponseBodyPrice) SetOriginalPrice(v float32) *GetPriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetPromotionDetails(v []*GetPriceResponseBodyPricePromotionDetails) *GetPriceResponseBodyPrice {
	s.PromotionDetails = v
	return s
}

func (s *GetPriceResponseBodyPrice) SetTradePrice(v float32) *GetPriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *GetPriceResponseBodyPrice) Validate() error {
	return dara.Validate(s)
}

type GetPriceResponseBodyPriceModuleDetails struct {
	// The discount price.
	//
	// example:
	//
	// 0.02
	CostAfterDiscount *float32 `json:"costAfterDiscount,omitempty" xml:"costAfterDiscount,omitempty"`
	// The discount.
	//
	// example:
	//
	// 0.0
	InvoiceDiscount *float32 `json:"invoiceDiscount,omitempty" xml:"invoiceDiscount,omitempty"`
	// The code of the pricing module.
	//
	// example:
	//
	// InstanceRent
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// The name of the pricing module.
	//
	// example:
	//
	// InstanceRent
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1000.0
	OriginalCost *float32 `json:"originalCost,omitempty" xml:"originalCost,omitempty"`
	// The price type.
	//
	// example:
	//
	// 1.0
	PriceType *string `json:"priceType,omitempty" xml:"priceType,omitempty"`
}

func (s GetPriceResponseBodyPriceModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPriceResponseBodyPriceModuleDetails) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *GetPriceResponseBodyPriceModuleDetails) GetPriceType() *string {
	return s.PriceType
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetCostAfterDiscount(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetInvoiceDiscount(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetModuleCode(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetModuleName(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.ModuleName = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetOriginalCost(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.OriginalCost = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetPriceType(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.PriceType = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) Validate() error {
	return dara.Validate(s)
}

type GetPriceResponseBodyPricePromotionDetails struct {
	// The description of the promotion.
	//
	// example:
	//
	// 37284
	PromotionDesc *string `json:"promotionDesc,omitempty" xml:"promotionDesc,omitempty"`
	// The ID of the promotion.
	PromotionId *int64 `json:"promotionId,omitempty" xml:"promotionId,omitempty"`
	// The name of the promotion.
	PromotionName *string `json:"promotionName,omitempty" xml:"promotionName,omitempty"`
}

func (s GetPriceResponseBodyPricePromotionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPriceResponseBodyPricePromotionDetails) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPricePromotionDetails) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetPriceResponseBodyPricePromotionDetails) GetPromotionId() *int64 {
	return s.PromotionId
}

func (s *GetPriceResponseBodyPricePromotionDetails) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionDesc(v string) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionDesc = &v
	return s
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionId(v int64) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionId = &v
	return s
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionName(v string) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionName = &v
	return s
}

func (s *GetPriceResponseBodyPricePromotionDetails) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeMultiPriceResponseBodyPriceInfo) *DescribeMultiPriceResponseBody
	GetPriceInfo() *DescribeMultiPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeMultiPriceResponseBody
	GetRequestId() *string
}

type DescribeMultiPriceResponseBody struct {
	PriceInfo *DescribeMultiPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBody) GetPriceInfo() *DescribeMultiPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeMultiPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiPriceResponseBody) SetPriceInfo(v *DescribeMultiPriceResponseBodyPriceInfo) *DescribeMultiPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeMultiPriceResponseBody) SetRequestId(v string) *DescribeMultiPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiPriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfo struct {
	Price *DescribeMultiPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*DescribeMultiPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMultiPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) GetPrice() *DescribeMultiPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) GetRules() []*DescribeMultiPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) SetPrice(v *DescribeMultiPriceResponseBodyPriceInfoPrice) *DescribeMultiPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) SetRules(v []*DescribeMultiPriceResponseBodyPriceInfoRules) *DescribeMultiPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 534.6
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 6800
	OriginalPrice            *float32                                                    `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetails             []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails `json:"PriceDetails,omitempty" xml:"PriceDetails,omitempty" type:"Repeated"`
	Promotions               []*DescribeMultiPriceResponseBodyPriceInfoPricePromotions   `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	RefundInstanceIdPriceMap map[string]*float32                                         `json:"RefundInstanceIdPriceMap,omitempty" xml:"RefundInstanceIdPriceMap,omitempty"`
	// example:
	//
	// 60.00
	RefundPrice *float32 `json:"RefundPrice,omitempty" xml:"RefundPrice,omitempty"`
	// example:
	//
	// 82.6
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetPriceDetails() []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	return s.PriceDetails
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetPromotions() []*DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	return s.Promotions
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetRefundInstanceIdPriceMap() map[string]*float32 {
	return s.RefundInstanceIdPriceMap
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetRefundPrice() *float32 {
	return s.RefundPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetPriceDetails(v []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.PriceDetails = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeMultiPriceResponseBodyPriceInfoPricePromotions) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetRefundInstanceIdPriceMap(v map[string]*float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.RefundInstanceIdPriceMap = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetRefundPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.RefundPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPrice) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails struct {
	ModuleDetails []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OrderItem   *int32                                                               `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	PriceDetail *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Struct"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) GetModuleDetails() []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	return s.ModuleDetails
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) GetOrderItem() *int32 {
	return s.OrderItem
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) GetPriceDetail() *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	return s.PriceDetail
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetModuleDetails(v []*DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.ModuleDetails = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetOrderItem(v int32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.OrderItem = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) SetPriceDetail(v *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails {
	s.PriceDetail = v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails struct {
	// example:
	//
	// 734.65
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// DesktopType
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// eds.enterprise_office.8c32g
	ModuleValue *string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty"`
	// example:
	//
	// 10900
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 292.2
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetModuleValue() *string {
	return s.ModuleValue
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleCode(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleName(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleName = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetModuleValue(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.ModuleValue = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails {
	s.TradePrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsModuleDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail struct {
	// example:
	//
	// 20.00
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 100.00
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 80.00
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetDiscountPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetOriginalPrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetResourceType(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) SetTradePrice(v float32) *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePriceDetailsPriceDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoPricePromotions struct {
	// example:
	//
	// new
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeMultiPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiPriceResponseBodyPriceInfoRules struct {
	// example:
	//
	// accounts_suspect_users
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hrzdvc
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMultiPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeMultiPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeMultiPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribeMultiPriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

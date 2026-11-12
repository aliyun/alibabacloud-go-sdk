// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody
	GetPriceInfo() *DescribePriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
}

type DescribePriceResponseBody struct {
	// The detailed pricing information.
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetPriceInfo() *DescribePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfo struct {
	// The pricing information.
	Price *DescribePriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The activity rule information.
	Rules []*DescribePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) GetPrice() *DescribePriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribePriceResponseBodyPriceInfo) GetRules() []*DescribePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRules(v []*DescribePriceResponseBodyPriceInfoRules) *DescribePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	// Specifies whether the coupon is selected.
	Coupons []*DescribePriceResponseBodyPriceInfoPriceCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Repeated"`
	// The currency. Valid values:
	//
	// - CNY: Chinese yuan.
	//
	// - USD: US dollar.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The detailed pricing information of the configuration.
	DetailInfos []*DescribePriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Repeated"`
	// The discounted price.
	//
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// Specifies whether a promotion is matched.
	IsContractPromotion *bool `json:"IsContractPromotion,omitempty" xml:"IsContractPromotion,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1580
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotion information.
	Promotions []*DescribePriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The discounted price based on the official website discount.
	//
	// example:
	//
	// 85
	StandardDiscountPrice *float32 `json:"StandardDiscountPrice,omitempty" xml:"StandardDiscountPrice,omitempty"`
	// The discounted price on the official website or the discounted price in the direct sales contract.
	//
	// example:
	//
	// 899
	StandardPrice *float32 `json:"StandardPrice,omitempty" xml:"StandardPrice,omitempty"`
	// The list of price details.
	SubPrices []*DescribePriceResponseBodyPriceInfoPriceSubPrices `json:"SubPrices,omitempty" xml:"SubPrices,omitempty" type:"Repeated"`
	// The transaction price.
	//
	// example:
	//
	// 267.4
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetCoupons() []*DescribePriceResponseBodyPriceInfoPriceCoupons {
	return s.Coupons
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDetailInfos() []*DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	return s.DetailInfos
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetIsContractPromotion() *bool {
	return s.IsContractPromotion
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetPromotions() []*DescribePriceResponseBodyPriceInfoPricePromotions {
	return s.Promotions
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetStandardDiscountPrice() *float32 {
	return s.StandardDiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetStandardPrice() *float32 {
	return s.StandardPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetSubPrices() []*DescribePriceResponseBodyPriceInfoPriceSubPrices {
	return s.SubPrices
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCoupons(v []*DescribePriceResponseBodyPriceInfoPriceCoupons) *DescribePriceResponseBodyPriceInfoPrice {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDetailInfos(v []*DescribePriceResponseBodyPriceInfoPriceDetailInfos) *DescribePriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetIsContractPromotion(v bool) *DescribePriceResponseBodyPriceInfoPrice {
	s.IsContractPromotion = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPricePromotions) *DescribePriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetStandardDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.StandardDiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetStandardPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.StandardPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetSubPrices(v []*DescribePriceResponseBodyPriceInfoPriceSubPrices) *DescribePriceResponseBodyPriceInfoPrice {
	s.SubPrices = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) Validate() error {
	if s.Coupons != nil {
		for _, item := range s.Coupons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DetailInfos != nil {
		for _, item := range s.DetailInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Promotions != nil {
		for _, item := range s.Promotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubPrices != nil {
		for _, item := range s.SubPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceCoupons struct {
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The custom parameter description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The discount.
	//
	// example:
	//
	// 20.0
	DiscountOff *float32 `json:"DiscountOff,omitempty" xml:"DiscountOff,omitempty"`
	// Specifies whether the coupon is selected.
	IsSelected *bool `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	// The command name.
	//
	// example:
	//
	// ACS-SWAS-GenerateRsaKeypair-windows.ps1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// *****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetDiscountOff() *float32 {
	return s.DiscountOff
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetIsSelected() *bool {
	return s.IsSelected
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetCouponNo(v string) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetDescription(v string) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetDiscountOff(v float32) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.DiscountOff = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetIsSelected(v bool) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetName(v string) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) SetOptionCode(v string) *DescribePriceResponseBodyPriceInfoPriceCoupons {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceCoupons) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfos struct {
	// The resource name.
	//
	// example:
	//
	// Server
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1850
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotion information matched by the configuration details.
	Promotions []*DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The transaction price.
	//
	// example:
	//
	// 267.4
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetCommodityType() *string {
	return s.CommodityType
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetPromotions() []*DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions {
	return s.Promotions
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetCommodityType(v string) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.CommodityType = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.Promotions = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) Validate() error {
	if s.Promotions != nil {
		for _, item := range s.Promotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions struct {
	// The discount amount.
	//
	// example:
	//
	// 0
	DiscountOff *float32 `json:"DiscountOff,omitempty" xml:"DiscountOff,omitempty"`
	// The promotion name.
	//
	// example:
	//
	// Discounted price on the official website for compute-optimized instances
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion rule ID.
	//
	// example:
	//
	// 0
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) GetDiscountOff() *float32 {
	return s.DiscountOff
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) SetDiscountOff(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions {
	s.DiscountOff = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) SetName(v string) *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosPromotions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPricePromotions struct {
	// The coupon discount amount.
	//
	// example:
	//
	// 20.0
	DiscountOff *float32 `json:"DiscountOff,omitempty" xml:"DiscountOff,omitempty"`
	// The coupon name.
	//
	// example:
	//
	// ACS-SWAS-ResetPassword-for-linux.sh
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion ID.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetDiscountOff() *float32 {
	return s.DiscountOff
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetDiscountOff(v float32) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.DiscountOff = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetName(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetRuleIds(v []*int64) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPriceSubPrices struct {
	// The order discount amount.
	//
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The instance ID of the Simple Application Server.
	//
	// example:
	//
	// d63a1c00ce14419194ae5bef5e*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether a promotion is matched.
	//
	// example:
	//
	// false
	IsContractPromotion *bool `json:"IsContractPromotion,omitempty" xml:"IsContractPromotion,omitempty"`
	// The original price of the order (CNY).
	//
	// example:
	//
	// 165.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotion name.
	Promotions []*DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The collection of activity rule IDs.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The discounted price based on the official website discount.
	//
	// example:
	//
	// 85.0
	StandardDiscountPrice *float32 `json:"StandardDiscountPrice,omitempty" xml:"StandardDiscountPrice,omitempty"`
	// The discounted price on the official website or the discounted price in the direct sales contract.
	//
	// example:
	//
	// 899.0
	StandardPrice *float32 `json:"StandardPrice,omitempty" xml:"StandardPrice,omitempty"`
	// The transaction price.
	//
	// example:
	//
	// 899.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceSubPrices) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceSubPrices) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetIsContractPromotion() *bool {
	return s.IsContractPromotion
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetPromotions() []*DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions {
	return s.Promotions
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetStandardDiscountPrice() *float32 {
	return s.StandardDiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetStandardPrice() *float32 {
	return s.StandardPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetInstanceId(v string) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.InstanceId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetIsContractPromotion(v bool) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.IsContractPromotion = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.Promotions = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetRuleIds(v []*int64) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetStandardDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.StandardDiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetStandardPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.StandardPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPrices {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPrices) Validate() error {
	if s.Promotions != nil {
		for _, item := range s.Promotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions struct {
	// The coupon discount amount.
	//
	// example:
	//
	// 20.0
	DiscountOff *float32 `json:"DiscountOff,omitempty" xml:"DiscountOff,omitempty"`
	// The name.
	//
	// example:
	//
	// Discount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion rule ID.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) GetDiscountOff() *float32 {
	return s.DiscountOff
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) SetDiscountOff(v float32) *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions {
	s.DiscountOff = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) SetName(v string) *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) SetRuleIds(v []*int64) *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceSubPricesPromotions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoRules struct {
	// The custom parameter description.
	//
	// example:
	//
	// Purchase a one-year subscription to enjoy an xxx discount on the official website price
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 142
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

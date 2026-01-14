// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *DescribeCommodityPriceResponseBody
	GetCurrency() *string
	SetDiscountPrice(v float64) *DescribeCommodityPriceResponseBody
	GetDiscountPrice() *float64
	SetOrderDetails(v []*DescribeCommodityPriceResponseBodyOrderDetails) *DescribeCommodityPriceResponseBody
	GetOrderDetails() []*DescribeCommodityPriceResponseBodyOrderDetails
	SetOriginalPrice(v float64) *DescribeCommodityPriceResponseBody
	GetOriginalPrice() *float64
	SetPromotions(v []*DescribeCommodityPriceResponseBodyPromotions) *DescribeCommodityPriceResponseBody
	GetPromotions() []*DescribeCommodityPriceResponseBodyPromotions
	SetRequestId(v string) *DescribeCommodityPriceResponseBody
	GetRequestId() *string
	SetRuleDetails(v []*DescribeCommodityPriceResponseBodyRuleDetails) *DescribeCommodityPriceResponseBody
	GetRuleDetails() []*DescribeCommodityPriceResponseBodyRuleDetails
	SetTradePrice(v float64) *DescribeCommodityPriceResponseBody
	GetTradePrice() *float64
}

type DescribeCommodityPriceResponseBody struct {
	// The currency unit.
	//
	// 	- China site: **CNY**.
	//
	// 	- International site: **USD**.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount.
	//
	// example:
	//
	// 419.8
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The details of the commodity module.
	OrderDetails []*DescribeCommodityPriceResponseBodyOrderDetails `json:"OrderDetails,omitempty" xml:"OrderDetails,omitempty" type:"Repeated"`
	// The original price.
	//
	// example:
	//
	// 2099
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The details of the coupon.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	Promotions []*DescribeCommodityPriceResponseBodyPromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the discount rules.
	RuleDetails []*DescribeCommodityPriceResponseBodyRuleDetails `json:"RuleDetails,omitempty" xml:"RuleDetails,omitempty" type:"Repeated"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 1679.2
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCommodityPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCommodityPriceResponseBody) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *DescribeCommodityPriceResponseBody) GetOrderDetails() []*DescribeCommodityPriceResponseBodyOrderDetails {
	return s.OrderDetails
}

func (s *DescribeCommodityPriceResponseBody) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *DescribeCommodityPriceResponseBody) GetPromotions() []*DescribeCommodityPriceResponseBodyPromotions {
	return s.Promotions
}

func (s *DescribeCommodityPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommodityPriceResponseBody) GetRuleDetails() []*DescribeCommodityPriceResponseBodyRuleDetails {
	return s.RuleDetails
}

func (s *DescribeCommodityPriceResponseBody) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *DescribeCommodityPriceResponseBody) SetCurrency(v string) *DescribeCommodityPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetDiscountPrice(v float64) *DescribeCommodityPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetOrderDetails(v []*DescribeCommodityPriceResponseBodyOrderDetails) *DescribeCommodityPriceResponseBody {
	s.OrderDetails = v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetOriginalPrice(v float64) *DescribeCommodityPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetPromotions(v []*DescribeCommodityPriceResponseBodyPromotions) *DescribeCommodityPriceResponseBody {
	s.Promotions = v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetRequestId(v string) *DescribeCommodityPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetRuleDetails(v []*DescribeCommodityPriceResponseBodyRuleDetails) *DescribeCommodityPriceResponseBody {
	s.RuleDetails = v
	return s
}

func (s *DescribeCommodityPriceResponseBody) SetTradePrice(v float64) *DescribeCommodityPriceResponseBody {
	s.TradePrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBody) Validate() error {
	if s.OrderDetails != nil {
		for _, item := range s.OrderDetails {
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
	if s.RuleDetails != nil {
		for _, item := range s.RuleDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityPriceResponseBodyOrderDetails struct {
	// The code of the commodity.
	//
	// example:
	//
	// ga_gapluspre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The name of the commodity.
	//
	// example:
	//
	// Global Accelerator_Instance Type (Subscription)
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The discount.
	//
	// example:
	//
	// 419.8
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The information about the commodity module.
	ModuleDetails []*DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Repeated"`
	// The original price.
	//
	// example:
	//
	// 2099
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The details of the discount.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	PromDetails []*DescribeCommodityPriceResponseBodyOrderDetailsPromDetails `json:"PromDetails,omitempty" xml:"PromDetails,omitempty" type:"Repeated"`
	// The number of instances that are purchased.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The IDs of discount rules.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 1679.2
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCommodityPriceResponseBodyOrderDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBodyOrderDetails) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetModuleDetails() []*DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	return s.ModuleDetails
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetPromDetails() []*DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	return s.PromDetails
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetQuantity() *int64 {
	return s.Quantity
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetCommodityCode(v string) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetCommodityName(v string) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.CommodityName = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetDiscountPrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetModuleDetails(v []*DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.ModuleDetails = v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetOriginalPrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetPromDetails(v []*DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.PromDetails = v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetQuantity(v int64) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.Quantity = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetRuleIds(v []*int64) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.RuleIds = v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) SetTradePrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetails {
	s.TradePrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetails) Validate() error {
	if s.ModuleDetails != nil {
		for _, item := range s.ModuleDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PromDetails != nil {
		for _, item := range s.PromDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails struct {
	// The discount.
	//
	// example:
	//
	// 40000.0
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The code of the commodity module.
	//
	// example:
	//
	// spec
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The name of the commodity module.
	//
	// example:
	//
	// Specification
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The original price.
	//
	// example:
	//
	// 200000.0
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 160000.0
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) SetDiscountPrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) SetModuleCode(v string) *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) SetModuleName(v string) *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	s.ModuleName = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) SetOriginalPrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) SetTradePrice(v float64) *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails {
	s.TradePrice = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsModuleDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeCommodityPriceResponseBodyOrderDetailsPromDetails struct {
	// The discounted price.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// 0.01
	FinalPromFee *float64 `json:"FinalPromFee,omitempty" xml:"FinalPromFee,omitempty"`
	// The code of the discount option.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// youhui_quan
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The sub-type of the discount.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// deduct
	PromType *string `json:"PromType,omitempty" xml:"PromType,omitempty"`
	// The ID of the discount item.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// 50003298014****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The name of the discount item.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// coupon
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GetFinalPromFee() *float64 {
	return s.FinalPromFee
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GetPromType() *string {
	return s.PromType
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) SetFinalPromFee(v float64) *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	s.FinalPromFee = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) SetOptionCode(v string) *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	s.OptionCode = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) SetPromType(v string) *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	s.PromType = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) SetPromotionId(v string) *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	s.PromotionId = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) SetPromotionName(v string) *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails {
	s.PromotionName = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyOrderDetailsPromDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeCommodityPriceResponseBodyPromotions struct {
	// The discounted amount.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// 0
	CanPromFee *float64 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// The code of the commodity to which the coupon can be applied.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// youhui_quan
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The name of the coupon.
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The code of the coupon.
	//
	// > 	- `youhuiquan_promotion_option_id_for_blank` indicates coupons that cannot be applied to the commodity.
	//
	// > 	- This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// 50003298014****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Indicates whether the coupon was selected.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter does not take effect for accounts registered on the China site (aliyun.com).
	//
	// example:
	//
	// false
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeCommodityPriceResponseBodyPromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBodyPromotions) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBodyPromotions) GetCanPromFee() *float64 {
	return s.CanPromFee
}

func (s *DescribeCommodityPriceResponseBodyPromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeCommodityPriceResponseBodyPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeCommodityPriceResponseBodyPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *DescribeCommodityPriceResponseBodyPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeCommodityPriceResponseBodyPromotions) SetCanPromFee(v float64) *DescribeCommodityPriceResponseBodyPromotions {
	s.CanPromFee = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyPromotions) SetOptionCode(v string) *DescribeCommodityPriceResponseBodyPromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyPromotions) SetPromotionName(v string) *DescribeCommodityPriceResponseBodyPromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyPromotions) SetPromotionOptionNo(v string) *DescribeCommodityPriceResponseBodyPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyPromotions) SetSelected(v bool) *DescribeCommodityPriceResponseBodyPromotions {
	s.Selected = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyPromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeCommodityPriceResponseBodyRuleDetails struct {
	// The ID of the discount rule.
	//
	// example:
	//
	// 102104100786****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the discount rule.
	//
	// example:
	//
	// GA New Customers Small II Specification Monthly Subscription - 20% Discount
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeCommodityPriceResponseBodyRuleDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponseBodyRuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponseBodyRuleDetails) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCommodityPriceResponseBodyRuleDetails) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeCommodityPriceResponseBodyRuleDetails) SetRuleId(v string) *DescribeCommodityPriceResponseBodyRuleDetails {
	s.RuleId = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyRuleDetails) SetRuleName(v string) *DescribeCommodityPriceResponseBodyRuleDetails {
	s.RuleName = &v
	return s
}

func (s *DescribeCommodityPriceResponseBodyRuleDetails) Validate() error {
	return dara.Validate(s)
}

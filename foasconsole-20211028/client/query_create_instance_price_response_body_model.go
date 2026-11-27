// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCreateInstancePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *QueryCreateInstancePriceResponseBodyPriceInfo) *QueryCreateInstancePriceResponseBody
	GetPriceInfo() *QueryCreateInstancePriceResponseBodyPriceInfo
	SetRequestId(v string) *QueryCreateInstancePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCreateInstancePriceResponseBody
	GetSuccess() *bool
}

type QueryCreateInstancePriceResponseBody struct {
	// The price information, including the price and discount rules.
	PriceInfo *QueryCreateInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCreateInstancePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBody) GetPriceInfo() *QueryCreateInstancePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *QueryCreateInstancePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCreateInstancePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCreateInstancePriceResponseBody) SetPriceInfo(v *QueryCreateInstancePriceResponseBodyPriceInfo) *QueryCreateInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryCreateInstancePriceResponseBody) SetRequestId(v string) *QueryCreateInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBody) SetSuccess(v bool) *QueryCreateInstancePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCreateInstancePriceResponseBodyPriceInfo struct {
	// The error code.
	//
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The price reduction information.
	DepreciateInfo *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount.
	//
	// example:
	//
	// 655.2
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// Indicates whether the key account promotion is applied.
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// The Lingxiao request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	LxRequestId *string `json:"LxRequestId,omitempty" xml:"LxRequestId,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unpaid orders exist. Pay for or cancel the existing orders first
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The module prices.
	ModulePrices []*QueryCreateInstancePriceResponseBodyPriceInfoModulePrices `json:"ModulePrices,omitempty" xml:"ModulePrices,omitempty" type:"Repeated"`
	// The coupon group information.
	OptionalPromotions []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// The original amount.
	//
	// example:
	//
	// 4368
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The promotion rules.
	Rules []*QueryCreateInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The discounted price based on the official website discount.
	//
	// example:
	//
	// 21321
	StandDiscountPrice *string `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The discounted price on the official website or the discounted price in the direct sales contract.
	//
	// example:
	//
	// 32432
	StandPrice *string `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The final amount, which is the original amount minus the discount.
	//
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetCode() *string {
	return s.Code
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetDepreciateInfo() *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetLxRequestId() *string {
	return s.LxRequestId
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetModulePrices() []*QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	return s.ModulePrices
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetOptionalPromotions() []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetRules() []*QueryCreateInstancePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetStandDiscountPrice() *string {
	return s.StandDiscountPrice
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetLxRequestId(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.LxRequestId = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetModulePrices(v []*QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.ModulePrices = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryCreateInstancePriceResponseBodyPriceInfoRules) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) Validate() error {
	if s.DepreciateInfo != nil {
		if err := s.DepreciateInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ModulePrices != nil {
		for _, item := range s.ModulePrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OptionalPromotions != nil {
		for _, item := range s.OptionalPromotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	// The price reduction rate.
	//
	// example:
	//
	// 34%
	CheapRate *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The total price on the official website after the price reduction.
	//
	// example:
	//
	// 8000
	CheapStandAmount *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// Indicates whether to display the price reduction rate.
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The equivalent monthly price.
	//
	// example:
	//
	// 1000
	MonthPrice *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	// The original total price on the official website.
	//
	// example:
	//
	// 12000
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	// The start time of the price reduction, in ISO 8601 format.
	//
	// example:
	//
	// 2024-08-21T02:06:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapRate() *string {
	return s.CheapRate
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapStandAmount() *string {
	return s.CheapStandAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetMonthPrice() *string {
	return s.MonthPrice
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetOriginalStandAmount() *string {
	return s.OriginalStandAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceResponseBodyPriceInfoModulePrices struct {
	// The discount amount.
	//
	// example:
	//
	// 648.0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The unit price of the module.
	//
	// example:
	//
	// 183600.0000
	FinalUnitFee *string `json:"FinalUnitFee,omitempty" xml:"FinalUnitFee,omitempty"`
	// The module code.
	//
	// example:
	//
	// overhead_cu
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The module name.
	//
	// example:
	//
	// Control and Management Services
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The original amount.
	//
	// example:
	//
	// 4320.0
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The payment amount.
	//
	// example:
	//
	// 3672.0
	PayAmount *float32 `json:"PayAmount,omitempty" xml:"PayAmount,omitempty"`
	// The billing method.
	//
	// example:
	//
	// yearPrice
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// The pricing unit.
	//
	// example:
	//
	// Yuan/CU
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The order quantity.
	//
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetFinalUnitFee() *string {
	return s.FinalUnitFee
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetModuleName() *string {
	return s.ModuleName
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetPayAmount() *float32 {
	return s.PayAmount
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetPriceType() *string {
	return s.PriceType
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) GetQuantity() *string {
	return s.Quantity
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetDiscountAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetFinalUnitFee(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.FinalUnitFee = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetModuleCode(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.ModuleCode = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetModuleName(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.ModuleName = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetOriginalAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetPayAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.PayAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetPriceType(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.PriceType = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetPriceUnit(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.PriceUnit = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) SetQuantity(v string) *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices {
	s.Quantity = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoModulePrices) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	// The description of the coupon.
	//
	// example:
	//
	// ¥1,391.5 coupon (valid until 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The name of the coupon.
	//
	// example:
	//
	// ¥1,391.5 coupon
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The ID of the coupon.
	//
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Specifies whether the coupon is selected. Valid values:
	//
	// - true: Selected.
	//
	// - false: Not selected.
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceResponseBodyPriceInfoRules struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// Enjoy a 15% discount on the official website price for a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryCreateInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryCreateInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

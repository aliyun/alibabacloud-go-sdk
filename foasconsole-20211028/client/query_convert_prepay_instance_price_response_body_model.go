// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertPrepayInstancePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) *QueryConvertPrepayInstancePriceResponseBody
	GetPriceInfo() *QueryConvertPrepayInstancePriceResponseBodyPriceInfo
	SetRequestId(v string) *QueryConvertPrepayInstancePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryConvertPrepayInstancePriceResponseBody
	GetSuccess() *bool
}

type QueryConvertPrepayInstancePriceResponseBody struct {
	// The price information, including the price and discount rules.
	PriceInfo *QueryConvertPrepayInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBody) GetPriceInfo() *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *QueryConvertPrepayInstancePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConvertPrepayInstancePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetPriceInfo(v *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) *QueryConvertPrepayInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetRequestId(v string) *QueryConvertPrepayInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetSuccess(v bool) *QueryConvertPrepayInstancePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfo struct {
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
	DepreciateInfo *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount.
	//
	// example:
	//
	// 655.2
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// Indicates whether an enterprise discount is applied.
	//
	// example:
	//
	// true
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
	// An unpaid order exists. Pay for or cancel the existing order first
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The coupon group information.
	OptionalPromotions []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// The original price.
	//
	// example:
	//
	// 4368
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The promotion rules.
	Rules []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The discounted price based on the official website discount.
	//
	// example:
	//
	// 21321
	StandDiscountPrice *string `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The official website discount price or direct sales contract discount price.
	//
	// example:
	//
	// 32432
	StandPrice *string `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The final price, which is the original price minus the discount.
	//
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetCode() *string {
	return s.Code
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetDepreciateInfo() *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetLxRequestId() *string {
	return s.LxRequestId
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetOptionalPromotions() []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetRules() []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetStandDiscountPrice() *string {
	return s.StandDiscountPrice
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetLxRequestId(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.LxRequestId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) Validate() error {
	if s.DepreciateInfo != nil {
		if err := s.DepreciateInfo.Validate(); err != nil {
			return err
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

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	// The price reduction ratio.
	//
	// example:
	//
	// 20%
	CheapRate *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The reduced list price total.
	//
	// example:
	//
	// 8000
	CheapStandAmount *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// Indicates whether to display the price reduction percentage.
	//
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The equivalent monthly price.
	//
	// example:
	//
	// 4000
	MonthPrice *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	// The original list price total.
	//
	// example:
	//
	// 10000
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	// The price reduction start time.
	//
	// example:
	//
	// 2023-03-31T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapRate() *string {
	return s.CheapRate
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapStandAmount() *string {
	return s.CheapStandAmount
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetMonthPrice() *string {
	return s.MonthPrice
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetOriginalStandAmount() *string {
	return s.OriginalStandAmount
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	// The coupon description.
	//
	// example:
	//
	// ¥1,391.5 coupon (valid until 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon name.
	//
	// example:
	//
	// ¥1,391.5 coupon
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The coupon number.
	//
	// example:
	//
	// 500011220***
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Indicates whether the coupon is selected. Valid values:
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

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules struct {
	// The promotion rule description.
	//
	// example:
	//
	// Purchase for 1 year or more and enjoy a 15% discount off the list price
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

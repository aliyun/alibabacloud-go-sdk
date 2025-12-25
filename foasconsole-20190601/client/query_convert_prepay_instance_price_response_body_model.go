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
	PriceInfo *QueryConvertPrepayInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                             `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 655.2
	DiscountAmount     *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool    `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	LxRequestId *string `json:"LxRequestId,omitempty" xml:"LxRequestId,omitempty"`
	// example:
	//
	// 存在未支付订单，请先支付或取消原有订单
	Message            *string                                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                                     `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                      `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                      `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
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
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// example:
	//
	// ￥1,391.5 优惠券 (有效期至 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// ￥1,391.5 优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

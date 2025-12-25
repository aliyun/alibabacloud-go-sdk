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
	PriceInfo *QueryCreateInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                      `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
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
	Message            *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                              `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryCreateInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                               `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                               `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
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
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

type QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
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
	// 50001122001****
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 102007100514****
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

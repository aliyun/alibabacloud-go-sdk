// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertInstancePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *QueryConvertInstancePriceResponseBodyPriceInfo) *QueryConvertInstancePriceResponseBody
	GetPriceInfo() *QueryConvertInstancePriceResponseBodyPriceInfo
	SetRequestId(v string) *QueryConvertInstancePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryConvertInstancePriceResponseBody
	GetSuccess() *bool
}

type QueryConvertInstancePriceResponseBody struct {
	PriceInfo *QueryConvertInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConvertInstancePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBody) GetPriceInfo() *QueryConvertInstancePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *QueryConvertInstancePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConvertInstancePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryConvertInstancePriceResponseBody) SetPriceInfo(v *QueryConvertInstancePriceResponseBodyPriceInfo) *QueryConvertInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryConvertInstancePriceResponseBody) SetRequestId(v string) *QueryConvertInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBody) SetSuccess(v bool) *QueryConvertInstancePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertInstancePriceResponseBodyPriceInfo struct {
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                       `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
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
	Message            *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                               `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryConvertInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetCode() *string {
	return s.Code
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetDepreciateInfo() *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetLxRequestId() *string {
	return s.LxRequestId
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetOptionalPromotions() []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetRules() []*QueryConvertInstancePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetStandDiscountPrice() *string {
	return s.StandDiscountPrice
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetLxRequestId(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.LxRequestId = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryConvertInstancePriceResponseBodyPriceInfoRules) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) Validate() error {
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

type QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapRate() *string {
	return s.CheapRate
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapStandAmount() *string {
	return s.CheapStandAmount
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetMonthPrice() *string {
	return s.MonthPrice
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetOriginalStandAmount() *string {
	return s.OriginalStandAmount
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
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
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryConvertInstancePriceResponseBodyPriceInfoRules struct {
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1020071005141834
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryConvertInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryConvertInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

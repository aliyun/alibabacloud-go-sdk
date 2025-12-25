// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewInstancePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *QueryRenewInstancePriceResponseBodyPriceInfo) *QueryRenewInstancePriceResponseBody
	GetPriceInfo() *QueryRenewInstancePriceResponseBodyPriceInfo
	SetRequestId(v string) *QueryRenewInstancePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRenewInstancePriceResponseBody
	GetSuccess() *bool
}

type QueryRenewInstancePriceResponseBody struct {
	PriceInfo *QueryRenewInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRenewInstancePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBody) GetPriceInfo() *QueryRenewInstancePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *QueryRenewInstancePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRenewInstancePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRenewInstancePriceResponseBody) SetPriceInfo(v *QueryRenewInstancePriceResponseBodyPriceInfo) *QueryRenewInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryRenewInstancePriceResponseBody) SetRequestId(v string) *QueryRenewInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBody) SetSuccess(v bool) *QueryRenewInstancePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRenewInstancePriceResponseBodyPriceInfo struct {
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                     `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
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
	Message            *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                             `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryRenewInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                              `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                              `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetCode() *string {
	return s.Code
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetDepreciateInfo() *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetLxRequestId() *string {
	return s.LxRequestId
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetOptionalPromotions() []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetRules() []*QueryRenewInstancePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetStandDiscountPrice() *string {
	return s.StandDiscountPrice
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetLxRequestId(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.LxRequestId = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryRenewInstancePriceResponseBodyPriceInfoRules) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) Validate() error {
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

type QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapRate() *string {
	return s.CheapRate
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapStandAmount() *string {
	return s.CheapStandAmount
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetMonthPrice() *string {
	return s.MonthPrice
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetOriginalStandAmount() *string {
	return s.OriginalStandAmount
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
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

func (s QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryRenewInstancePriceResponseBodyPriceInfoRules struct {
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1020071005141834
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryRenewInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryRenewInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

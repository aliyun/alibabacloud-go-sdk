// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyInstancePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *QueryModifyInstancePriceResponseBodyPriceInfo) *QueryModifyInstancePriceResponseBody
	GetPriceInfo() *QueryModifyInstancePriceResponseBodyPriceInfo
	SetRequestId(v string) *QueryModifyInstancePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyInstancePriceResponseBody
	GetSuccess() *bool
}

type QueryModifyInstancePriceResponseBody struct {
	PriceInfo *QueryModifyInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyInstancePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBody) GetPriceInfo() *QueryModifyInstancePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *QueryModifyInstancePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyInstancePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyInstancePriceResponseBody) SetPriceInfo(v *QueryModifyInstancePriceResponseBodyPriceInfo) *QueryModifyInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryModifyInstancePriceResponseBody) SetRequestId(v string) *QueryModifyInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBody) SetSuccess(v bool) *QueryModifyInstancePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyInstancePriceResponseBodyPriceInfo struct {
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                      `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
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
	OptionalPromotions []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                              `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryModifyInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                               `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                               `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetCode() *string {
	return s.Code
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetDepreciateInfo() *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetLxRequestId() *string {
	return s.LxRequestId
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetOptionalPromotions() []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetRules() []*QueryModifyInstancePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetStandDiscountPrice() *string {
	return s.StandDiscountPrice
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetLxRequestId(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.LxRequestId = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryModifyInstancePriceResponseBodyPriceInfoRules) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) Validate() error {
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

type QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapRate() *string {
	return s.CheapRate
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetCheapStandAmount() *string {
	return s.CheapStandAmount
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetMonthPrice() *string {
	return s.MonthPrice
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetOriginalStandAmount() *string {
	return s.OriginalStandAmount
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
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

func (s QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GetSelected() *bool {
	return s.Selected
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyInstancePriceResponseBodyPriceInfoRules struct {
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1020071005141834
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryModifyInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryModifyInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}

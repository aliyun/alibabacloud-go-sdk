// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryRenewPriceResponseBodyData) *QueryRenewPriceResponseBody
	GetData() []*QueryRenewPriceResponseBodyData
	SetErrCode(v string) *QueryRenewPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryRenewPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryRenewPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryRenewPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRenewPriceResponseBody
	GetSuccess() *bool
}

type QueryRenewPriceResponseBody struct {
	Data []*QueryRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRenewPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBody) GetData() []*QueryRenewPriceResponseBodyData {
	return s.Data
}

func (s *QueryRenewPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryRenewPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryRenewPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryRenewPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRenewPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRenewPriceResponseBody) SetData(v []*QueryRenewPriceResponseBodyData) *QueryRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryRenewPriceResponseBody) SetErrCode(v string) *QueryRenewPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetErrMessage(v string) *QueryRenewPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetHttpStatusCode(v int32) *QueryRenewPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetRequestId(v string) *QueryRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetSuccess(v bool) *QueryRenewPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRenewPriceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRenewPriceResponseBodyData struct {
	// example:
	//
	// c-96f3bc7f04b2****
	BillingInstanceId *string `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 9812
	DepreciateInfo *QueryRenewPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                             `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryRenewPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 9812
	OriginalAmount *float32                                `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryRenewPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 9812
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 9812
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 9812
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryRenewPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyData) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *QueryRenewPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryRenewPriceResponseBodyData) GetDepreciateInfo() *QueryRenewPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryRenewPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryRenewPriceResponseBodyData) GetOptionalPromotions() []*QueryRenewPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryRenewPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryRenewPriceResponseBodyData) GetRules() []*QueryRenewPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryRenewPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryRenewPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryRenewPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryRenewPriceResponseBodyData) SetBillingInstanceId(v string) *QueryRenewPriceResponseBodyData {
	s.BillingInstanceId = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetCurrency(v string) *QueryRenewPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetDepreciateInfo(v *QueryRenewPriceResponseBodyDataDepreciateInfo) *QueryRenewPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetDiscountAmount(v float32) *QueryRenewPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetOptionalPromotions(v []*QueryRenewPriceResponseBodyDataOptionalPromotions) *QueryRenewPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetOriginalAmount(v float32) *QueryRenewPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetRules(v []*QueryRenewPriceResponseBodyDataRules) *QueryRenewPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryRenewPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetStandPrice(v float32) *QueryRenewPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetTradeAmount(v float32) *QueryRenewPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) Validate() error {
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

type QueryRenewPriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 9812
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 9812
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryRenewPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryRenewPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryRenewPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryRenewPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryRenewPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryRenewPriceResponseBodyDataOptionalPromotions struct {
	// example:
	//
	// youhuiquan_desc
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s QueryRenewPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryRenewPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryRenewPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryRenewPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryRenewPriceResponseBodyDataRules struct {
	// example:
	//
	// 1
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_8syh2j121ns
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8su2i3hsdf128
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryRenewPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryRenewPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryRenewPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryRenewPriceResponseBodyDataRules) SetAmount(v float32) *QueryRenewPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataRules) SetName(v string) *QueryRenewPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryRenewPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

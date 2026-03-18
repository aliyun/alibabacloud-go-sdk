// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyNodeNumberPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyNodeNumberPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyNodeNumberPriceResponseBodyData) *QueryModifyNodeNumberPriceResponseBody
	GetData() *QueryModifyNodeNumberPriceResponseBodyData
	SetErrCode(v string) *QueryModifyNodeNumberPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyNodeNumberPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyNodeNumberPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyNodeNumberPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyNodeNumberPriceResponseBody
	GetSuccess() *bool
}

type QueryModifyNodeNumberPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyNodeNumberPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
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

func (s QueryModifyNodeNumberPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetData() *QueryModifyNodeNumberPriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyNodeNumberPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyNodeNumberPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetData(v *QueryModifyNodeNumberPriceResponseBodyData) *QueryModifyNodeNumberPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetErrCode(v string) *QueryModifyNodeNumberPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetErrMessage(v string) *QueryModifyNodeNumberPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyNodeNumberPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetRequestId(v string) *QueryModifyNodeNumberPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) SetSuccess(v bool) *QueryModifyNodeNumberPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyNodeNumberPriceResponseBodyData struct {
	ComponentPrices []*QueryModifyNodeNumberPriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 0
	DepreciateInfo *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                        `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 26710
	OriginalAmount *float32                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyNodeNumberPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 26710
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 26710
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 26710
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyNodeNumberPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetComponentPrices() []*QueryModifyNodeNumberPriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetDepreciateInfo() *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetOptionalPromotions() []*QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetRules() []*QueryModifyNodeNumberPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetComponentPrices(v []*QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) *QueryModifyNodeNumberPriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetCurrency(v string) *QueryModifyNodeNumberPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetDepreciateInfo(v *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) *QueryModifyNodeNumberPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) *QueryModifyNodeNumberPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetRules(v []*QueryModifyNodeNumberPriceResponseBodyDataRules) *QueryModifyNodeNumberPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyNodeNumberPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetStandPrice(v float32) *QueryModifyNodeNumberPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyData) Validate() error {
	if s.ComponentPrices != nil {
		for _, item := range s.ComponentPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type QueryModifyNodeNumberPriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// RangerUserSync
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 26710
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 26710
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 26710
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 26710
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyNodeNumberPriceResponseBodyDataRules struct {
	// example:
	//
	// 10
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_923su2sf
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8uy3sh12sa
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyNodeNumberPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyNodeNumberPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) SetName(v string) *QueryModifyNodeNumberPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyNodeNumberPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

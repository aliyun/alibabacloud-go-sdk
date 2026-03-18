// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifySpecTypePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifySpecTypePriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifySpecTypePriceResponseBodyData) *QueryModifySpecTypePriceResponseBody
	GetData() *QueryModifySpecTypePriceResponseBodyData
	SetErrCode(v string) *QueryModifySpecTypePriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifySpecTypePriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifySpecTypePriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifySpecTypePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifySpecTypePriceResponseBody
	GetSuccess() *bool
}

type QueryModifySpecTypePriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifySpecTypePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifySpecTypePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifySpecTypePriceResponseBody) GetData() *QueryModifySpecTypePriceResponseBodyData {
	return s.Data
}

func (s *QueryModifySpecTypePriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifySpecTypePriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifySpecTypePriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifySpecTypePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifySpecTypePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifySpecTypePriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifySpecTypePriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetData(v *QueryModifySpecTypePriceResponseBodyData) *QueryModifySpecTypePriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetErrCode(v string) *QueryModifySpecTypePriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetErrMessage(v string) *QueryModifySpecTypePriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetHttpStatusCode(v int32) *QueryModifySpecTypePriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetRequestId(v string) *QueryModifySpecTypePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) SetSuccess(v bool) *QueryModifySpecTypePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifySpecTypePriceResponseBodyData struct {
	ComponentPrices []*QueryModifySpecTypePriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency       *string                                                 `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                      `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifySpecTypePriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 5612
	OriginalAmount *float32                                         `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifySpecTypePriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 5612
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 5612
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 5612
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifySpecTypePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetComponentPrices() []*QueryModifySpecTypePriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetDepreciateInfo() *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetOptionalPromotions() []*QueryModifySpecTypePriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetRules() []*QueryModifySpecTypePriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifySpecTypePriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetComponentPrices(v []*QueryModifySpecTypePriceResponseBodyDataComponentPrices) *QueryModifySpecTypePriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetCurrency(v string) *QueryModifySpecTypePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetDepreciateInfo(v *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) *QueryModifySpecTypePriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifySpecTypePriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetOptionalPromotions(v []*QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) *QueryModifySpecTypePriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifySpecTypePriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetRules(v []*QueryModifySpecTypePriceResponseBodyDataRules) *QueryModifySpecTypePriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifySpecTypePriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetStandPrice(v float32) *QueryModifySpecTypePriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) SetTradeAmount(v float32) *QueryModifySpecTypePriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyData) Validate() error {
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

type QueryModifySpecTypePriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// node_type
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 5612
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 5612
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifySpecTypePriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifySpecTypePriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifySpecTypePriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 5612
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 5612
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifySpecTypePriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifySpecTypePriceResponseBodyDataRules struct {
	// example:
	//
	// 10
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_827231sg1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7usy32gs01
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifySpecTypePriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) SetAmount(v float32) *QueryModifySpecTypePriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) SetName(v string) *QueryModifySpecTypePriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifySpecTypePriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifySpecTypePriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskNumberPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyDiskNumberPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyDiskNumberPriceResponseBodyData) *QueryModifyDiskNumberPriceResponseBody
	GetData() *QueryModifyDiskNumberPriceResponseBodyData
	SetErrCode(v string) *QueryModifyDiskNumberPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyDiskNumberPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyDiskNumberPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyDiskNumberPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyDiskNumberPriceResponseBody
	GetSuccess() *bool
}

type QueryModifyDiskNumberPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the order.
	Data *QueryModifyDiskNumberPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetData() *QueryModifyDiskNumberPriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyDiskNumberPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyDiskNumberPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetData(v *QueryModifyDiskNumberPriceResponseBodyData) *QueryModifyDiskNumberPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetErrCode(v string) *QueryModifyDiskNumberPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetErrMessage(v string) *QueryModifyDiskNumberPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyDiskNumberPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetRequestId(v string) *QueryModifyDiskNumberPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) SetSuccess(v bool) *QueryModifyDiskNumberPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyDiskNumberPriceResponseBodyData struct {
	// The prices of components.
	ComponentPrices []*QueryModifyDiskNumberPriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The total official price after discount.
	//
	// example:
	//
	// 17629
	DepreciateInfo *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount = original amount − payable amount on the bill (including coupon deductions).
	//
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The coupon information.
	OptionalPromotions []*QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// The original amount = catalog price × usage.
	//
	// example:
	//
	// 17629
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The list of rules that match your request. Each item is a matching rule. Only matching rule information and the location of each matching rule are returned.
	Rules []*QueryModifyDiskNumberPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The discounted price.
	//
	// example:
	//
	// 17629
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The official discounted price.
	//
	// example:
	//
	// 17629
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The final amount.
	//
	// example:
	//
	// 17629
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetComponentPrices() []*QueryModifyDiskNumberPriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetDepreciateInfo() *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetOptionalPromotions() []*QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetRules() []*QueryModifyDiskNumberPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetComponentPrices(v []*QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) *QueryModifyDiskNumberPriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetCurrency(v string) *QueryModifyDiskNumberPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetDepreciateInfo(v *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) *QueryModifyDiskNumberPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) *QueryModifyDiskNumberPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetRules(v []*QueryModifyDiskNumberPriceResponseBodyDataRules) *QueryModifyDiskNumberPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyDiskNumberPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetStandPrice(v float32) *QueryModifyDiskNumberPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyData) Validate() error {
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

type QueryModifyDiskNumberPriceResponseBodyDataComponentPrices struct {
	// The name of the component.
	//
	// example:
	//
	// disk
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The discount amount for the order.
	//
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original amount.
	//
	// example:
	//
	// 17629
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The final amount, which equals the original amount minus the discount.
	//
	// example:
	//
	// 17629
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo struct {
	// The discount rate.
	//
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The total official price after discount.
	//
	// example:
	//
	// 17629
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// Indicates whether to show the discount amount.
	//
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The original total official price.
	//
	// example:
	//
	// 17629
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions struct {
	// The description of the coupon.
	//
	// example:
	//
	// youhuiquan_desc
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The name of the coupon.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The ID of the coupon.
	//
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskNumberPriceResponseBodyDataRules struct {
	// The number of Elastic Compute Service (ECS) instances for which you want to query the price. Valid values: 1 to 1000. Default value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 3
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// rule_12hus92
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 6yhsi10223
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyDiskNumberPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) SetName(v string) *QueryModifyDiskNumberPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyDiskNumberPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

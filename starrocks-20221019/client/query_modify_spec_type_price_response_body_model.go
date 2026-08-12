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
	// Details about access denial.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data *QueryModifySpecTypePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
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
	// The component prices.
	ComponentPrices []*QueryModifySpecTypePriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The name of the campaign.
	DepreciateInfo *QueryModifySpecTypePriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount. Formula: Original amount - Payable amount on the bill. The payable amount includes the amount deducted by coupons.
	//
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The coupon information.
	OptionalPromotions []*QueryModifySpecTypePriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// The original price. Formula: List price × Billed usage.
	//
	// example:
	//
	// 5612
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The price calculation rules.
	Rules []*QueryModifySpecTypePriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The discounted price that is based on the official website discount.
	//
	// example:
	//
	// 5612
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The official website discount price.
	//
	// example:
	//
	// 5612
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The amount.
	//
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
	// The component name.
	//
	// example:
	//
	// node_type
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price.
	//
	// example:
	//
	// 5612
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The final price. The final price is the original price minus the discount.
	//
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
	// The price reduction ratio.
	//
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The total official website price after the price reduction.
	//
	// example:
	//
	// 5612
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// Indicates whether to display the price reduction.
	//
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The original total official website price.
	//
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
	// The coupon description.
	//
	// example:
	//
	// youhuiquan_desc
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon name.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The coupon ID.
	//
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
	// The resource count.
	//
	// example:
	//
	// 10
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_827231sg1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The rule ID.
	//
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

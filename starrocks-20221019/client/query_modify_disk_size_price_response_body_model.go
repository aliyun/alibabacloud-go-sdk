// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskSizePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyDiskSizePriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyDiskSizePriceResponseBodyData) *QueryModifyDiskSizePriceResponseBody
	GetData() *QueryModifyDiskSizePriceResponseBodyData
	SetErrCode(v string) *QueryModifyDiskSizePriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyDiskSizePriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyDiskSizePriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyDiskSizePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyDiskSizePriceResponseBody
	GetSuccess() *bool
}

type QueryModifyDiskSizePriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyDiskSizePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyDiskSizePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyDiskSizePriceResponseBody) GetData() *QueryModifyDiskSizePriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyDiskSizePriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyDiskSizePriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyDiskSizePriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyDiskSizePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyDiskSizePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyDiskSizePriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyDiskSizePriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetData(v *QueryModifyDiskSizePriceResponseBodyData) *QueryModifyDiskSizePriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetErrCode(v string) *QueryModifyDiskSizePriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetErrMessage(v string) *QueryModifyDiskSizePriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyDiskSizePriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetRequestId(v string) *QueryModifyDiskSizePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) SetSuccess(v bool) *QueryModifyDiskSizePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyDiskSizePriceResponseBodyData struct {
	ComponentPrices []*QueryModifyDiskSizePriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// true
	DepreciateInfo *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                      `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 9872
	OriginalAmount *float32                                         `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyDiskSizePriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 9872
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 9872
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 9872
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskSizePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetComponentPrices() []*QueryModifyDiskSizePriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetDepreciateInfo() *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetOptionalPromotions() []*QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetRules() []*QueryModifyDiskSizePriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyDiskSizePriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetComponentPrices(v []*QueryModifyDiskSizePriceResponseBodyDataComponentPrices) *QueryModifyDiskSizePriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetCurrency(v string) *QueryModifyDiskSizePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetDepreciateInfo(v *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) *QueryModifyDiskSizePriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyDiskSizePriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) *QueryModifyDiskSizePriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyDiskSizePriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetRules(v []*QueryModifyDiskSizePriceResponseBodyDataRules) *QueryModifyDiskSizePriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyDiskSizePriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetStandPrice(v float32) *QueryModifyDiskSizePriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyDiskSizePriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyData) Validate() error {
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

type QueryModifyDiskSizePriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// disk_size
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 9872
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 9872
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskSizePriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyDiskSizePriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 9872
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// 0
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 9872
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskSizePriceResponseBodyDataRules struct {
	// example:
	//
	// 10
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule-083ja12
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7ysj12ksaa
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyDiskSizePriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyDiskSizePriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) SetName(v string) *QueryModifyDiskSizePriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyDiskSizePriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskPerformanceLevelPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyDiskPerformanceLevelPriceResponseBodyData) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetData() *QueryModifyDiskPerformanceLevelPriceResponseBodyData
	SetErrCode(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyDiskPerformanceLevelPriceResponseBody
	GetSuccess() *bool
}

type QueryModifyDiskPerformanceLevelPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyDiskPerformanceLevelPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 86865DD4-7733-5446-B52B-C9DA71961B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetData() *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetData(v *QueryModifyDiskPerformanceLevelPriceResponseBodyData) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetErrCode(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetErrMessage(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetRequestId(v string) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) SetSuccess(v bool) *QueryModifyDiskPerformanceLevelPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyDiskPerformanceLevelPriceResponseBodyData struct {
	ComponentPrices []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 17629
	DepreciateInfo *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                                  `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 17629
	OriginalAmount *float32                                                     `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 17629
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 17629
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 17629
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetComponentPrices() []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetDepreciateInfo() *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetOptionalPromotions() []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetRules() []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetComponentPrices(v []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetCurrency(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetDepreciateInfo(v *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetRules(v []*QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetStandPrice(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyData) Validate() error {
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

type QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// disk_type
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 17629
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 17629
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 17629
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 17629
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules struct {
	// example:
	//
	// 1
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_128nsg182
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 78hujis12jd
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) SetName(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

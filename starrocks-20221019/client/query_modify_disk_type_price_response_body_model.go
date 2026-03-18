// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskTypePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyDiskTypePriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyDiskTypePriceResponseBodyData) *QueryModifyDiskTypePriceResponseBody
	GetData() *QueryModifyDiskTypePriceResponseBodyData
	SetErrCode(v string) *QueryModifyDiskTypePriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyDiskTypePriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyDiskTypePriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyDiskTypePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyDiskTypePriceResponseBody
	GetSuccess() *bool
}

type QueryModifyDiskTypePriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyDiskTypePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [Region id should be select from set [cn-beijing, cn-hangzhou]]
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyDiskTypePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyDiskTypePriceResponseBody) GetData() *QueryModifyDiskTypePriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyDiskTypePriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyDiskTypePriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyDiskTypePriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyDiskTypePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyDiskTypePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyDiskTypePriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyDiskTypePriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetData(v *QueryModifyDiskTypePriceResponseBodyData) *QueryModifyDiskTypePriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetErrCode(v string) *QueryModifyDiskTypePriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetErrMessage(v string) *QueryModifyDiskTypePriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyDiskTypePriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetRequestId(v string) *QueryModifyDiskTypePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) SetSuccess(v bool) *QueryModifyDiskTypePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyDiskTypePriceResponseBodyData struct {
	ComponentPrices []*QueryModifyDiskTypePriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency       *string                                                 `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                      `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 10923
	OriginalAmount *float32                                         `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyDiskTypePriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 10923
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 10923
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 10923
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskTypePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetComponentPrices() []*QueryModifyDiskTypePriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetDepreciateInfo() *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetOptionalPromotions() []*QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetRules() []*QueryModifyDiskTypePriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyDiskTypePriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetComponentPrices(v []*QueryModifyDiskTypePriceResponseBodyDataComponentPrices) *QueryModifyDiskTypePriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetCurrency(v string) *QueryModifyDiskTypePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetDepreciateInfo(v *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) *QueryModifyDiskTypePriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyDiskTypePriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) *QueryModifyDiskTypePriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyDiskTypePriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetRules(v []*QueryModifyDiskTypePriceResponseBodyDataRules) *QueryModifyDiskTypePriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyDiskTypePriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetStandPrice(v float32) *QueryModifyDiskTypePriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyDiskTypePriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyData) Validate() error {
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

type QueryModifyDiskTypePriceResponseBodyDataComponentPrices struct {
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
	// 10923
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 10923
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyDiskTypePriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyDiskTypePriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 10923
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 10923
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyDiskTypePriceResponseBodyDataRules struct {
	// example:
	//
	// 10
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule-928shy23sa
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7dsu230ks23
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyDiskTypePriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyDiskTypePriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) SetName(v string) *QueryModifyDiskTypePriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyDiskTypePriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

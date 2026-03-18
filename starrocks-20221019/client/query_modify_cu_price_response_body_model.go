// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyCuPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyCuPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyCuPriceResponseBodyData) *QueryModifyCuPriceResponseBody
	GetData() *QueryModifyCuPriceResponseBodyData
	SetErrCode(v string) *QueryModifyCuPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyCuPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyCuPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyCuPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyCuPriceResponseBody
	GetSuccess() *bool
}

type QueryModifyCuPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                             `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyCuPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// B67D142D-D54E-184F-A306-22BDC01B2XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyCuPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyCuPriceResponseBody) GetData() *QueryModifyCuPriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyCuPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyCuPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyCuPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyCuPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyCuPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyCuPriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyCuPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetData(v *QueryModifyCuPriceResponseBodyData) *QueryModifyCuPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetErrCode(v string) *QueryModifyCuPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetErrMessage(v string) *QueryModifyCuPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyCuPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetRequestId(v string) *QueryModifyCuPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) SetSuccess(v bool) *QueryModifyCuPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyCuPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyCuPriceResponseBodyData struct {
	ComponentPrices []*QueryModifyCuPriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 0
	DepreciateInfo *QueryModifyCuPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OptionalPromotions []*QueryModifyCuPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 7986
	OriginalAmount *float32                                   `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyCuPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 7986
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 7986
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 7986
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyCuPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBodyData) GetComponentPrices() []*QueryModifyCuPriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyCuPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyCuPriceResponseBodyData) GetDepreciateInfo() *QueryModifyCuPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyCuPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyCuPriceResponseBodyData) GetOptionalPromotions() []*QueryModifyCuPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyCuPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyCuPriceResponseBodyData) GetRules() []*QueryModifyCuPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyCuPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyCuPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyCuPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyCuPriceResponseBodyData) SetComponentPrices(v []*QueryModifyCuPriceResponseBodyDataComponentPrices) *QueryModifyCuPriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetCurrency(v string) *QueryModifyCuPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetDepreciateInfo(v *QueryModifyCuPriceResponseBodyDataDepreciateInfo) *QueryModifyCuPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyCuPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyCuPriceResponseBodyDataOptionalPromotions) *QueryModifyCuPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyCuPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetRules(v []*QueryModifyCuPriceResponseBodyDataRules) *QueryModifyCuPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyCuPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetStandPrice(v float32) *QueryModifyCuPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyCuPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyData) Validate() error {
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

type QueryModifyCuPriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// cu_num
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 7986
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 7986
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyCuPriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyCuPriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyCuPriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyCuPriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyCuPriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyCuPriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 7986
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 7986
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryModifyCuPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyCuPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyCuPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyCuPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyCuPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyCuPriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryModifyCuPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyCuPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyCuPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyCuPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyCuPriceResponseBodyDataRules struct {
	// example:
	//
	// 1
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_123123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ak72hajsd
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyCuPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyCuPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyCuPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyCuPriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyCuPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataRules) SetName(v string) *QueryModifyCuPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyCuPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyCuPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

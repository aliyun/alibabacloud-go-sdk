// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnableMultiAzPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryEnableMultiAzPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryEnableMultiAzPriceResponseBodyData) *QueryEnableMultiAzPriceResponseBody
	GetData() *QueryEnableMultiAzPriceResponseBodyData
	SetErrCode(v string) *QueryEnableMultiAzPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryEnableMultiAzPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryEnableMultiAzPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryEnableMultiAzPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEnableMultiAzPriceResponseBody
	GetSuccess() *bool
}

type QueryEnableMultiAzPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryEnableMultiAzPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryEnableMultiAzPriceResponseBody) GetData() *QueryEnableMultiAzPriceResponseBodyData {
	return s.Data
}

func (s *QueryEnableMultiAzPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryEnableMultiAzPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryEnableMultiAzPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryEnableMultiAzPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEnableMultiAzPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEnableMultiAzPriceResponseBody) SetAccessDeniedDetail(v string) *QueryEnableMultiAzPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetData(v *QueryEnableMultiAzPriceResponseBodyData) *QueryEnableMultiAzPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetErrCode(v string) *QueryEnableMultiAzPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetErrMessage(v string) *QueryEnableMultiAzPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetHttpStatusCode(v int32) *QueryEnableMultiAzPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetRequestId(v string) *QueryEnableMultiAzPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) SetSuccess(v bool) *QueryEnableMultiAzPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEnableMultiAzPriceResponseBodyData struct {
	ComponentPrices []*QueryEnableMultiAzPriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 7641
	DepreciateInfo *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                     `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	ModuleInstance     []*QueryEnableMultiAzPriceResponseBodyDataModuleInstance     `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
	OptionalPromotions []*QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 7641
	OriginalAmount *float32                                        `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryEnableMultiAzPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 7641
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 7641
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 7641
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetComponentPrices() []*QueryEnableMultiAzPriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetDepreciateInfo() *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetModuleInstance() []*QueryEnableMultiAzPriceResponseBodyDataModuleInstance {
	return s.ModuleInstance
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetOptionalPromotions() []*QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetRules() []*QueryEnableMultiAzPriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryEnableMultiAzPriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetComponentPrices(v []*QueryEnableMultiAzPriceResponseBodyDataComponentPrices) *QueryEnableMultiAzPriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetCurrency(v string) *QueryEnableMultiAzPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetDepreciateInfo(v *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) *QueryEnableMultiAzPriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetDiscountAmount(v float32) *QueryEnableMultiAzPriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetModuleInstance(v []*QueryEnableMultiAzPriceResponseBodyDataModuleInstance) *QueryEnableMultiAzPriceResponseBodyData {
	s.ModuleInstance = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetOptionalPromotions(v []*QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) *QueryEnableMultiAzPriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetOriginalAmount(v float32) *QueryEnableMultiAzPriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetRules(v []*QueryEnableMultiAzPriceResponseBodyDataRules) *QueryEnableMultiAzPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryEnableMultiAzPriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetStandPrice(v float32) *QueryEnableMultiAzPriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) SetTradeAmount(v float32) *QueryEnableMultiAzPriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyData) Validate() error {
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
	if s.ModuleInstance != nil {
		for _, item := range s.ModuleInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type QueryEnableMultiAzPriceResponseBodyDataComponentPrices struct {
	// example:
	//
	// multi-zone
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 7641
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 7641
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryEnableMultiAzPriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 7641
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 7641
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryEnableMultiAzPriceResponseBodyDataModuleInstance struct {
	// example:
	//
	// multi-zone
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// example:
	//
	// zone
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 7641
	StandPrice *string `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 7641
	TotalProductFee *string `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBodyDataModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyDataModuleInstance) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) GetModuleName() *string {
	return s.ModuleName
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) GetTotalProductFee() *string {
	return s.TotalProductFee
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) SetModuleCode(v string) *QueryEnableMultiAzPriceResponseBodyDataModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) SetModuleName(v string) *QueryEnableMultiAzPriceResponseBodyDataModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) SetStandPrice(v string) *QueryEnableMultiAzPriceResponseBodyDataModuleInstance {
	s.StandPrice = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) SetTotalProductFee(v string) *QueryEnableMultiAzPriceResponseBodyDataModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataModuleInstance) Validate() error {
	return dara.Validate(s)
}

type QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions struct {
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

func (s QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryEnableMultiAzPriceResponseBodyDataRules struct {
	// example:
	//
	// 2
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_28sdhsu2320t
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2ju2sd9dss
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryEnableMultiAzPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) SetAmount(v float32) *QueryEnableMultiAzPriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) SetName(v string) *QueryEnableMultiAzPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) SetRuleDescId(v string) *QueryEnableMultiAzPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyChargeTypePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryModifyChargeTypePriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryModifyChargeTypePriceResponseBodyData) *QueryModifyChargeTypePriceResponseBody
	GetData() *QueryModifyChargeTypePriceResponseBodyData
	SetErrCode(v string) *QueryModifyChargeTypePriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryModifyChargeTypePriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryModifyChargeTypePriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryModifyChargeTypePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryModifyChargeTypePriceResponseBody
	GetSuccess() *bool
}

type QueryModifyChargeTypePriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryModifyChargeTypePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s QueryModifyChargeTypePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryModifyChargeTypePriceResponseBody) GetData() *QueryModifyChargeTypePriceResponseBodyData {
	return s.Data
}

func (s *QueryModifyChargeTypePriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryModifyChargeTypePriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryModifyChargeTypePriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryModifyChargeTypePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryModifyChargeTypePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryModifyChargeTypePriceResponseBody) SetAccessDeniedDetail(v string) *QueryModifyChargeTypePriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetData(v *QueryModifyChargeTypePriceResponseBodyData) *QueryModifyChargeTypePriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetErrCode(v string) *QueryModifyChargeTypePriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetErrMessage(v string) *QueryModifyChargeTypePriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetHttpStatusCode(v int32) *QueryModifyChargeTypePriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetRequestId(v string) *QueryModifyChargeTypePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) SetSuccess(v bool) *QueryModifyChargeTypePriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyChargeTypePriceResponseBodyData struct {
	ComponentPrices []*QueryModifyChargeTypePriceResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency       *string                                                   `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                                        `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	ModuleInstance     []*QueryModifyChargeTypePriceResponseBodyDataModuleInstance     `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
	OptionalPromotions []*QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 7986
	OriginalAmount *float32                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryModifyChargeTypePriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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

func (s QueryModifyChargeTypePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetComponentPrices() []*QueryModifyChargeTypePriceResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetDepreciateInfo() *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetModuleInstance() []*QueryModifyChargeTypePriceResponseBodyDataModuleInstance {
	return s.ModuleInstance
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetOptionalPromotions() []*QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetRules() []*QueryModifyChargeTypePriceResponseBodyDataRules {
	return s.Rules
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryModifyChargeTypePriceResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetComponentPrices(v []*QueryModifyChargeTypePriceResponseBodyDataComponentPrices) *QueryModifyChargeTypePriceResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetCurrency(v string) *QueryModifyChargeTypePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetDepreciateInfo(v *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) *QueryModifyChargeTypePriceResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetDiscountAmount(v float32) *QueryModifyChargeTypePriceResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetModuleInstance(v []*QueryModifyChargeTypePriceResponseBodyDataModuleInstance) *QueryModifyChargeTypePriceResponseBodyData {
	s.ModuleInstance = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetOptionalPromotions(v []*QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) *QueryModifyChargeTypePriceResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetOriginalAmount(v float32) *QueryModifyChargeTypePriceResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetRules(v []*QueryModifyChargeTypePriceResponseBodyDataRules) *QueryModifyChargeTypePriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetStandDiscountPrice(v float32) *QueryModifyChargeTypePriceResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetStandPrice(v float32) *QueryModifyChargeTypePriceResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) SetTradeAmount(v float32) *QueryModifyChargeTypePriceResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyData) Validate() error {
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

type QueryModifyChargeTypePriceResponseBodyDataComponentPrices struct {
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

func (s QueryModifyChargeTypePriceResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) SetComponentName(v string) *QueryModifyChargeTypePriceResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo struct {
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

func (s QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryModifyChargeTypePriceResponseBodyDataModuleInstance struct {
	// example:
	//
	// instance_type
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// example:
	//
	// cu_num
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 7986
	StandPrice *string `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 7986
	TotalProductFee *string `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
}

func (s QueryModifyChargeTypePriceResponseBodyDataModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyDataModuleInstance) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) GetModuleName() *string {
	return s.ModuleName
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) GetTotalProductFee() *string {
	return s.TotalProductFee
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) SetModuleCode(v string) *QueryModifyChargeTypePriceResponseBodyDataModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) SetModuleName(v string) *QueryModifyChargeTypePriceResponseBodyDataModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) SetStandPrice(v string) *QueryModifyChargeTypePriceResponseBodyDataModuleInstance {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) SetTotalProductFee(v string) *QueryModifyChargeTypePriceResponseBodyDataModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataModuleInstance) Validate() error {
	return dara.Validate(s)
}

type QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions struct {
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
	// youhuiquan_1238293
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryModifyChargeTypePriceResponseBodyDataRules struct {
	// example:
	//
	// 0
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// rule-12iudfj
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryModifyChargeTypePriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) SetAmount(v float32) *QueryModifyChargeTypePriceResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) SetName(v string) *QueryModifyChargeTypePriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) SetRuleDescId(v string) *QueryModifyChargeTypePriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

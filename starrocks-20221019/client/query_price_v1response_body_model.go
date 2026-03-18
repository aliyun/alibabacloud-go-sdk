// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceV1ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryPriceV1ResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *QueryPriceV1ResponseBodyData) *QueryPriceV1ResponseBody
	GetData() *QueryPriceV1ResponseBodyData
	SetErrCode(v string) *QueryPriceV1ResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryPriceV1ResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryPriceV1ResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryPriceV1ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPriceV1ResponseBody
	GetSuccess() *bool
}

type QueryPriceV1ResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *QueryPriceV1ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s QueryPriceV1ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryPriceV1ResponseBody) GetData() *QueryPriceV1ResponseBodyData {
	return s.Data
}

func (s *QueryPriceV1ResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryPriceV1ResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryPriceV1ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryPriceV1ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPriceV1ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPriceV1ResponseBody) SetAccessDeniedDetail(v string) *QueryPriceV1ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPriceV1ResponseBody) SetData(v *QueryPriceV1ResponseBodyData) *QueryPriceV1ResponseBody {
	s.Data = v
	return s
}

func (s *QueryPriceV1ResponseBody) SetErrCode(v string) *QueryPriceV1ResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryPriceV1ResponseBody) SetErrMessage(v string) *QueryPriceV1ResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryPriceV1ResponseBody) SetHttpStatusCode(v int32) *QueryPriceV1ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryPriceV1ResponseBody) SetRequestId(v string) *QueryPriceV1ResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPriceV1ResponseBody) SetSuccess(v bool) *QueryPriceV1ResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPriceV1ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPriceV1ResponseBodyData struct {
	ComponentPrices []*QueryPriceV1ResponseBodyDataComponentPrices `json:"ComponentPrices,omitempty" xml:"ComponentPrices,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency       *string                                     `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryPriceV1ResponseBodyDataDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DiscountAmount     *float32                                          `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	ModuleInstance     []*QueryPriceV1ResponseBodyDataModuleInstance     `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
	OptionalPromotions []*QueryPriceV1ResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 8094
	OriginalAmount *float32                             `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules          []*QueryPriceV1ResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 8094
	StandDiscountPrice *float32 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// example:
	//
	// 8094
	StandPrice *float32 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 8094
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryPriceV1ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyData) GetComponentPrices() []*QueryPriceV1ResponseBodyDataComponentPrices {
	return s.ComponentPrices
}

func (s *QueryPriceV1ResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryPriceV1ResponseBodyData) GetDepreciateInfo() *QueryPriceV1ResponseBodyDataDepreciateInfo {
	return s.DepreciateInfo
}

func (s *QueryPriceV1ResponseBodyData) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryPriceV1ResponseBodyData) GetModuleInstance() []*QueryPriceV1ResponseBodyDataModuleInstance {
	return s.ModuleInstance
}

func (s *QueryPriceV1ResponseBodyData) GetOptionalPromotions() []*QueryPriceV1ResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *QueryPriceV1ResponseBodyData) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryPriceV1ResponseBodyData) GetRules() []*QueryPriceV1ResponseBodyDataRules {
	return s.Rules
}

func (s *QueryPriceV1ResponseBodyData) GetStandDiscountPrice() *float32 {
	return s.StandDiscountPrice
}

func (s *QueryPriceV1ResponseBodyData) GetStandPrice() *float32 {
	return s.StandPrice
}

func (s *QueryPriceV1ResponseBodyData) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryPriceV1ResponseBodyData) SetComponentPrices(v []*QueryPriceV1ResponseBodyDataComponentPrices) *QueryPriceV1ResponseBodyData {
	s.ComponentPrices = v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetCurrency(v string) *QueryPriceV1ResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetDepreciateInfo(v *QueryPriceV1ResponseBodyDataDepreciateInfo) *QueryPriceV1ResponseBodyData {
	s.DepreciateInfo = v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetDiscountAmount(v float32) *QueryPriceV1ResponseBodyData {
	s.DiscountAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetModuleInstance(v []*QueryPriceV1ResponseBodyDataModuleInstance) *QueryPriceV1ResponseBodyData {
	s.ModuleInstance = v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetOptionalPromotions(v []*QueryPriceV1ResponseBodyDataOptionalPromotions) *QueryPriceV1ResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetOriginalAmount(v float32) *QueryPriceV1ResponseBodyData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetRules(v []*QueryPriceV1ResponseBodyDataRules) *QueryPriceV1ResponseBodyData {
	s.Rules = v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetStandDiscountPrice(v float32) *QueryPriceV1ResponseBodyData {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetStandPrice(v float32) *QueryPriceV1ResponseBodyData {
	s.StandPrice = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) SetTradeAmount(v float32) *QueryPriceV1ResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyData) Validate() error {
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

type QueryPriceV1ResponseBodyDataComponentPrices struct {
	// example:
	//
	// FE
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// example:
	//
	// 3192
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// example:
	//
	// 3192
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryPriceV1ResponseBodyDataComponentPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyDataComponentPrices) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) SetComponentName(v string) *QueryPriceV1ResponseBodyDataComponentPrices {
	s.ComponentName = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) SetDiscountAmount(v float32) *QueryPriceV1ResponseBodyDataComponentPrices {
	s.DiscountAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) SetOriginalAmount(v float32) *QueryPriceV1ResponseBodyDataComponentPrices {
	s.OriginalAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) SetTradeAmount(v float32) *QueryPriceV1ResponseBodyDataComponentPrices {
	s.TradeAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataComponentPrices) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1ResponseBodyDataDepreciateInfo struct {
	// example:
	//
	// 0
	CheapRate *float32 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// example:
	//
	// 8094
	CheapStandAmount *float32 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// example:
	//
	// 8094
	OriginalStandAmount *float32 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s QueryPriceV1ResponseBodyDataDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyDataDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) GetCheapRate() *float32 {
	return s.CheapRate
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) GetCheapStandAmount() *float32 {
	return s.CheapStandAmount
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) GetOriginalStandAmount() *float32 {
	return s.OriginalStandAmount
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) SetCheapRate(v float32) *QueryPriceV1ResponseBodyDataDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) SetCheapStandAmount(v float32) *QueryPriceV1ResponseBodyDataDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) SetIsShow(v bool) *QueryPriceV1ResponseBodyDataDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) SetOriginalStandAmount(v float32) *QueryPriceV1ResponseBodyDataDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1ResponseBodyDataModuleInstance struct {
	// example:
	//
	// cu_num
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// example:
	//
	// CU
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 1622
	StandPrice *string `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 1622
	TotalProductFee *string `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
}

func (s QueryPriceV1ResponseBodyDataModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyDataModuleInstance) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) GetModuleName() *string {
	return s.ModuleName
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) GetStandPrice() *string {
	return s.StandPrice
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) GetTotalProductFee() *string {
	return s.TotalProductFee
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) SetModuleCode(v string) *QueryPriceV1ResponseBodyDataModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) SetModuleName(v string) *QueryPriceV1ResponseBodyDataModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) SetStandPrice(v string) *QueryPriceV1ResponseBodyDataModuleInstance {
	s.StandPrice = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) SetTotalProductFee(v string) *QueryPriceV1ResponseBodyDataModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataModuleInstance) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1ResponseBodyDataOptionalPromotions struct {
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

func (s QueryPriceV1ResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *QueryPriceV1ResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) SetPromotionName(v string) *QueryPriceV1ResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *QueryPriceV1ResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type QueryPriceV1ResponseBodyDataRules struct {
	// example:
	//
	// 1
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// rule_8usi12321sa
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7u22yshaasds
	RuleDescId *string `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s QueryPriceV1ResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1ResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *QueryPriceV1ResponseBodyDataRules) GetAmount() *float32 {
	return s.Amount
}

func (s *QueryPriceV1ResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *QueryPriceV1ResponseBodyDataRules) GetRuleDescId() *string {
	return s.RuleDescId
}

func (s *QueryPriceV1ResponseBodyDataRules) SetAmount(v float32) *QueryPriceV1ResponseBodyDataRules {
	s.Amount = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataRules) SetName(v string) *QueryPriceV1ResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataRules) SetRuleDescId(v string) *QueryPriceV1ResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *QueryPriceV1ResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

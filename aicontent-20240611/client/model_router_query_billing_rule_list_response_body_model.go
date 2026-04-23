// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryBillingRuleListResponseBodyData) *ModelRouterQueryBillingRuleListResponseBody
	GetData() *ModelRouterQueryBillingRuleListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryBillingRuleListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryBillingRuleListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryBillingRuleListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryBillingRuleListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryBillingRuleListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryBillingRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryBillingRuleListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryBillingRuleListResponseBody struct {
	Data *ModelRouterQueryBillingRuleListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryBillingRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetData() *ModelRouterQueryBillingRuleListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryBillingRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetData(v *ModelRouterQueryBillingRuleListResponseBodyData) *ModelRouterQueryBillingRuleListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetErrCode(v string) *ModelRouterQueryBillingRuleListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetErrMessage(v string) *ModelRouterQueryBillingRuleListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryBillingRuleListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetMaxResults(v int32) *ModelRouterQueryBillingRuleListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetNextToken(v string) *ModelRouterQueryBillingRuleListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetRequestId(v string) *ModelRouterQueryBillingRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) SetSuccess(v bool) *ModelRouterQueryBillingRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryBillingRuleListResponseBodyData struct {
	List []*ModelRouterQueryBillingRuleListResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryBillingRuleListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) GetList() []*ModelRouterQueryBillingRuleListResponseBodyDataList {
	return s.List
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) SetList(v []*ModelRouterQueryBillingRuleListResponseBodyDataList) *ModelRouterQueryBillingRuleListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) SetPage(v int32) *ModelRouterQueryBillingRuleListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryBillingRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) SetTotal(v int32) *ModelRouterQueryBillingRuleListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelRouterQueryBillingRuleListResponseBodyDataList struct {
	// example:
	//
	// token_tiered
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// example:
	//
	// 2025-01-01T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 通义千问Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// {}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterQueryBillingRuleListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingRuleListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetModelName() *string {
	return s.ModelName
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetBillingType(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.BillingType = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetDeleteTag(v int32) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.DeleteTag = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetEffectiveTime(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetExpireTime(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetGmtCreate(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetGmtModified(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetId(v int64) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetModelCode(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.ModelCode = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetModelId(v int64) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetModelName(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.ModelName = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetModelType(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.ModelType = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetPricingConfig(v interface{}) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetSymbol(v string) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.Symbol = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) SetVersion(v int32) *ModelRouterQueryBillingRuleListResponseBodyDataList {
	s.Version = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

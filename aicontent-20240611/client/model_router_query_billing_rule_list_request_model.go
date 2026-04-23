// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveOnly(v bool) *ModelRouterQueryBillingRuleListRequest
	GetActiveOnly() *bool
	SetMaxResults(v int32) *ModelRouterQueryBillingRuleListRequest
	GetMaxResults() *int32
	SetModelCode(v string) *ModelRouterQueryBillingRuleListRequest
	GetModelCode() *string
	SetModelId(v int32) *ModelRouterQueryBillingRuleListRequest
	GetModelId() *int32
	SetModelType(v string) *ModelRouterQueryBillingRuleListRequest
	GetModelType() *string
	SetNextToken(v string) *ModelRouterQueryBillingRuleListRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterQueryBillingRuleListRequest
	GetPage() *int32
	SetPageIndex(v int32) *ModelRouterQueryBillingRuleListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryBillingRuleListRequest
	GetPageSize() *int32
}

type ModelRouterQueryBillingRuleListRequest struct {
	// example:
	//
	// true
	ActiveOnly *bool `json:"activeOnly,omitempty" xml:"activeOnly,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// 1
	ModelId *int32 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ModelRouterQueryBillingRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingRuleListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingRuleListRequest) GetActiveOnly() *bool {
	return s.ActiveOnly
}

func (s *ModelRouterQueryBillingRuleListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryBillingRuleListRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelRouterQueryBillingRuleListRequest) GetModelId() *int32 {
	return s.ModelId
}

func (s *ModelRouterQueryBillingRuleListRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterQueryBillingRuleListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryBillingRuleListRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryBillingRuleListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryBillingRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryBillingRuleListRequest) SetActiveOnly(v bool) *ModelRouterQueryBillingRuleListRequest {
	s.ActiveOnly = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetMaxResults(v int32) *ModelRouterQueryBillingRuleListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetModelCode(v string) *ModelRouterQueryBillingRuleListRequest {
	s.ModelCode = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetModelId(v int32) *ModelRouterQueryBillingRuleListRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetModelType(v string) *ModelRouterQueryBillingRuleListRequest {
	s.ModelType = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetNextToken(v string) *ModelRouterQueryBillingRuleListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetPage(v int32) *ModelRouterQueryBillingRuleListRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetPageIndex(v int32) *ModelRouterQueryBillingRuleListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) SetPageSize(v int32) *ModelRouterQueryBillingRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *ListLlmTemplatesRequest
	GetBizType() *int32
	SetLanguage(v string) *ListLlmTemplatesRequest
	GetLanguage() *string
	SetLlmCode(v string) *ListLlmTemplatesRequest
	GetLlmCode() *string
	SetLlmTemplateIds(v []*string) *ListLlmTemplatesRequest
	GetLlmTemplateIds() []*string
	SetModelTemplateId(v string) *ListLlmTemplatesRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListLlmTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLlmTemplatesRequest
	GetPageSize() *int32
	SetProviderTemplateId(v string) *ListLlmTemplatesRequest
	GetProviderTemplateId() *string
	SetSmartModel(v bool) *ListLlmTemplatesRequest
	GetSmartModel() *bool
}

type ListLlmTemplatesRequest struct {
	// The business type. This parameter is required when SmartModel is set to true.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The display language of the model. Valid values:
	//
	// - zh-CN: Simplified Chinese.
	//
	// - en-US: English.
	//
	// - ja-JP: Japanese.
	//
	// - zh-TW: Traditional Chinese.
	//
	// Default value: zh-CN. Currently, zh-TW returns Simplified Chinese, and ja-JP returns English. If the model is not included in Alibaba Cloud Model Studio or the English text is missing, the original source text is retained without automatic translation.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The model code for filtering. Fuzzy match is supported.
	//
	// example:
	//
	// qwen3.6-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// The model template IDs for filtering.
	LlmTemplateIds []*string `json:"LlmTemplateIds,omitempty" xml:"LlmTemplateIds,omitempty" type:"Repeated"`
	// The associated model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The page number. Pages start from 1. A value of 0 returns the same result as 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The model provider template ID.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
	// Specifies whether to query smart models. Valid values:
	//
	// - true: Only LLMs under system-preset smart models are returned. In this case, BizType is required.
	//
	// - false: Smart models are not specifically filtered.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SmartModel *bool `json:"SmartModel,omitempty" xml:"SmartModel,omitempty"`
}

func (s ListLlmTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListLlmTemplatesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListLlmTemplatesRequest) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListLlmTemplatesRequest) GetLlmTemplateIds() []*string {
	return s.LlmTemplateIds
}

func (s *ListLlmTemplatesRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListLlmTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmTemplatesRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListLlmTemplatesRequest) GetSmartModel() *bool {
	return s.SmartModel
}

func (s *ListLlmTemplatesRequest) SetBizType(v int32) *ListLlmTemplatesRequest {
	s.BizType = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetLanguage(v string) *ListLlmTemplatesRequest {
	s.Language = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetLlmCode(v string) *ListLlmTemplatesRequest {
	s.LlmCode = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetLlmTemplateIds(v []*string) *ListLlmTemplatesRequest {
	s.LlmTemplateIds = v
	return s
}

func (s *ListLlmTemplatesRequest) SetModelTemplateId(v string) *ListLlmTemplatesRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetPageNumber(v int32) *ListLlmTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetPageSize(v int32) *ListLlmTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetProviderTemplateId(v string) *ListLlmTemplatesRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListLlmTemplatesRequest) SetSmartModel(v bool) *ListLlmTemplatesRequest {
	s.SmartModel = &v
	return s
}

func (s *ListLlmTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

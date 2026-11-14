// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *ListLlmTemplatesShrinkRequest
	GetBizType() *int32
	SetLanguage(v string) *ListLlmTemplatesShrinkRequest
	GetLanguage() *string
	SetLlmCode(v string) *ListLlmTemplatesShrinkRequest
	GetLlmCode() *string
	SetLlmTemplateIdsShrink(v string) *ListLlmTemplatesShrinkRequest
	GetLlmTemplateIdsShrink() *string
	SetModelTemplateId(v string) *ListLlmTemplatesShrinkRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListLlmTemplatesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLlmTemplatesShrinkRequest
	GetPageSize() *int32
	SetProviderTemplateId(v string) *ListLlmTemplatesShrinkRequest
	GetProviderTemplateId() *string
	SetSmartModel(v bool) *ListLlmTemplatesShrinkRequest
	GetSmartModel() *bool
}

type ListLlmTemplatesShrinkRequest struct {
	// The business type. This parameter is required when SmartModel is set to true.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The display language of the model. Valid values: zh-CN (Simplified Chinese), en-US (English), ja-JP (Japanese), and zh-TW (Traditional Chinese). Default value: zh-CN. Currently, zh-TW returns Simplified Chinese, and ja-JP returns English. If the model is not recorded in Model Studio or the English copy is missing, the original source copy is retained without automatic translation.
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
	LlmTemplateIdsShrink *string `json:"LlmTemplateIds,omitempty" xml:"LlmTemplateIds,omitempty"`
	// The associated model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The page number. Pages start from page 1. If you set this parameter to 0, the system returns the same result as when you set it to 1.
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
	// Specifies whether to query smart models. If you set this parameter to true, only LLMs under system-preset smart models are returned, and BizType is required. Default value: false.
	//
	// example:
	//
	// false
	SmartModel *bool `json:"SmartModel,omitempty" xml:"SmartModel,omitempty"`
}

func (s ListLlmTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesShrinkRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListLlmTemplatesShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListLlmTemplatesShrinkRequest) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListLlmTemplatesShrinkRequest) GetLlmTemplateIdsShrink() *string {
	return s.LlmTemplateIdsShrink
}

func (s *ListLlmTemplatesShrinkRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListLlmTemplatesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmTemplatesShrinkRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListLlmTemplatesShrinkRequest) GetSmartModel() *bool {
	return s.SmartModel
}

func (s *ListLlmTemplatesShrinkRequest) SetBizType(v int32) *ListLlmTemplatesShrinkRequest {
	s.BizType = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetLanguage(v string) *ListLlmTemplatesShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetLlmCode(v string) *ListLlmTemplatesShrinkRequest {
	s.LlmCode = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetLlmTemplateIdsShrink(v string) *ListLlmTemplatesShrinkRequest {
	s.LlmTemplateIdsShrink = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetModelTemplateId(v string) *ListLlmTemplatesShrinkRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetPageNumber(v int32) *ListLlmTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetPageSize(v int32) *ListLlmTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetProviderTemplateId(v string) *ListLlmTemplatesShrinkRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) SetSmartModel(v bool) *ListLlmTemplatesShrinkRequest {
	s.SmartModel = &v
	return s
}

func (s *ListLlmTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

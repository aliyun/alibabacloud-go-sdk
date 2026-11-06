// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderLlmCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListModelProviderLlmCodesRequest
	GetLanguage() *string
	SetLlmCode(v string) *ListModelProviderLlmCodesRequest
	GetLlmCode() *string
	SetModelBrands(v []*string) *ListModelProviderLlmCodesRequest
	GetModelBrands() []*string
	SetPageNumber(v int32) *ListModelProviderLlmCodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProviderLlmCodesRequest
	GetPageSize() *int32
	SetProviderName(v string) *ListModelProviderLlmCodesRequest
	GetProviderName() *string
	SetProviderTemplateId(v string) *ListModelProviderLlmCodesRequest
	GetProviderTemplateId() *string
	SetRecommended(v bool) *ListModelProviderLlmCodesRequest
	GetRecommended() *bool
	SetRecommendedLevels(v []*string) *ListModelProviderLlmCodesRequest
	GetRecommendedLevels() []*string
}

type ListModelProviderLlmCodesRequest struct {
	// The display language for models. Valid values: zh-CN (Simplified Chinese), en-US (English), ja-JP (Japanese), and zh-TW (Traditional Chinese). Default value: zh-CN. Currently, zh-TW returns Simplified Chinese, and ja-JP returns English. If the text is not included in Model Studio or the English text is missing, the original source text is retained without automatic translation.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The LLM code used for fuzzy match filtering.
	//
	// example:
	//
	// qwen3-max
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// The list of model brands as a JSON array string. Supports multi-select filtering and returns results that match any of the specified brands.
	ModelBrands []*string `json:"ModelBrands,omitempty" xml:"ModelBrands,omitempty" type:"Repeated"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 1000. Maximum value: 1000.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the model provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// moonshot
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The ID of the provider template.
	//
	// This parameter is required.
	//
	// example:
	//
	// pt-xxx-001
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
	// Specifies whether to filter by recommended models. If this parameter is not specified, no filtering is applied.
	Recommended *bool `json:"Recommended,omitempty" xml:"Recommended,omitempty"`
	// The list of recommended levels. This parameter takes effect only for wuying-plan. Results must match any of the selected levels and intersect with other filtering conditions such as region and brand. If this parameter is not specified or an empty array is passed, filtering by level is not applied.
	//
	// example:
	//
	// ["Standard","Advanced"]
	RecommendedLevels []*string `json:"RecommendedLevels,omitempty" xml:"RecommendedLevels,omitempty" type:"Repeated"`
}

func (s ListModelProviderLlmCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesRequest) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListModelProviderLlmCodesRequest) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListModelProviderLlmCodesRequest) GetModelBrands() []*string {
	return s.ModelBrands
}

func (s *ListModelProviderLlmCodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProviderLlmCodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProviderLlmCodesRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderLlmCodesRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListModelProviderLlmCodesRequest) GetRecommended() *bool {
	return s.Recommended
}

func (s *ListModelProviderLlmCodesRequest) GetRecommendedLevels() []*string {
	return s.RecommendedLevels
}

func (s *ListModelProviderLlmCodesRequest) SetLanguage(v string) *ListModelProviderLlmCodesRequest {
	s.Language = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetLlmCode(v string) *ListModelProviderLlmCodesRequest {
	s.LlmCode = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetModelBrands(v []*string) *ListModelProviderLlmCodesRequest {
	s.ModelBrands = v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetPageNumber(v int32) *ListModelProviderLlmCodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetPageSize(v int32) *ListModelProviderLlmCodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetProviderName(v string) *ListModelProviderLlmCodesRequest {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetProviderTemplateId(v string) *ListModelProviderLlmCodesRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetRecommended(v bool) *ListModelProviderLlmCodesRequest {
	s.Recommended = &v
	return s
}

func (s *ListModelProviderLlmCodesRequest) SetRecommendedLevels(v []*string) *ListModelProviderLlmCodesRequest {
	s.RecommendedLevels = v
	return s
}

func (s *ListModelProviderLlmCodesRequest) Validate() error {
	return dara.Validate(s)
}

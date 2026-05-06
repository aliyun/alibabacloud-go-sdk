// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type ListLlmTemplatesShrinkRequest struct {
	// example:
	//
	// qwen3.6-plus
	LlmCode              *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	LlmTemplateIdsShrink *string `json:"LlmTemplateIds,omitempty" xml:"LlmTemplateIds,omitempty"`
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s ListLlmTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesShrinkRequest) GoString() string {
	return s.String()
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

func (s *ListLlmTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

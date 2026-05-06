// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type ListLlmTemplatesRequest struct {
	// example:
	//
	// qwen3.6-plus
	LlmCode        *string   `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	LlmTemplateIds []*string `json:"LlmTemplateIds,omitempty" xml:"LlmTemplateIds,omitempty" type:"Repeated"`
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

func (s ListLlmTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesRequest) GoString() string {
	return s.String()
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

func (s *ListLlmTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

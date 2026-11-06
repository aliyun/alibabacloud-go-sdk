// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDistillationTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListDistillationTemplatesRequest
	GetCategory() *string
	SetKeyword(v string) *ListDistillationTemplatesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListDistillationTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDistillationTemplatesRequest
	GetPageSize() *int32
	SetTemplateId(v string) *ListDistillationTemplatesRequest
	GetTemplateId() *string
	SetTemplateKind(v string) *ListDistillationTemplatesRequest
	GetTemplateKind() *string
}

type ListDistillationTemplatesRequest struct {
	// example:
	//
	// reasoning
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 推理
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// scenario
	TemplateKind *string `json:"TemplateKind,omitempty" xml:"TemplateKind,omitempty"`
}

func (s ListDistillationTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDistillationTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDistillationTemplatesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListDistillationTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDistillationTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDistillationTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDistillationTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListDistillationTemplatesRequest) GetTemplateKind() *string {
	return s.TemplateKind
}

func (s *ListDistillationTemplatesRequest) SetCategory(v string) *ListDistillationTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListDistillationTemplatesRequest) SetKeyword(v string) *ListDistillationTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListDistillationTemplatesRequest) SetPageNumber(v int32) *ListDistillationTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDistillationTemplatesRequest) SetPageSize(v int32) *ListDistillationTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDistillationTemplatesRequest) SetTemplateId(v string) *ListDistillationTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListDistillationTemplatesRequest) SetTemplateKind(v string) *ListDistillationTemplatesRequest {
	s.TemplateKind = &v
	return s
}

func (s *ListDistillationTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

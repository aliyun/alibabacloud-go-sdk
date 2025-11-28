// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateMaterialRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateMaterialRequest
	GetNextToken() *string
	SetPage(v int32) *ListTemplateMaterialRequest
	GetPage() *int32
	SetSize(v int32) *ListTemplateMaterialRequest
	GetSize() *int32
	SetTemplateIds(v string) *ListTemplateMaterialRequest
	GetTemplateIds() *string
}

type ListTemplateMaterialRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// M1Ti7gfj7VGDQgD588hxReiw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// M1Ti7gfj7VGDQgD588hxReiw,M1j9tcbkh9YtBw7BdOYlDusQ
	TemplateIds *string `json:"templateIds,omitempty" xml:"templateIds,omitempty"`
}

func (s ListTemplateMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateMaterialRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateMaterialRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateMaterialRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateMaterialRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListTemplateMaterialRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListTemplateMaterialRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *ListTemplateMaterialRequest) SetMaxResults(v int32) *ListTemplateMaterialRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateMaterialRequest) SetNextToken(v string) *ListTemplateMaterialRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateMaterialRequest) SetPage(v int32) *ListTemplateMaterialRequest {
	s.Page = &v
	return s
}

func (s *ListTemplateMaterialRequest) SetSize(v int32) *ListTemplateMaterialRequest {
	s.Size = &v
	return s
}

func (s *ListTemplateMaterialRequest) SetTemplateIds(v string) *ListTemplateMaterialRequest {
	s.TemplateIds = &v
	return s
}

func (s *ListTemplateMaterialRequest) Validate() error {
	return dara.Validate(s)
}

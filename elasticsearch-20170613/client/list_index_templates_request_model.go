// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexTemplate(v string) *ListIndexTemplatesRequest
	GetIndexTemplate() *string
	SetPage(v int32) *ListIndexTemplatesRequest
	GetPage() *int32
	SetSize(v int32) *ListIndexTemplatesRequest
	GetSize() *int32
}

type ListIndexTemplatesRequest struct {
	// example:
	//
	// my-template
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	// example:
	//
	// 5
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 50
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListIndexTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndexTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesRequest) GetIndexTemplate() *string {
	return s.IndexTemplate
}

func (s *ListIndexTemplatesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListIndexTemplatesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListIndexTemplatesRequest) SetIndexTemplate(v string) *ListIndexTemplatesRequest {
	s.IndexTemplate = &v
	return s
}

func (s *ListIndexTemplatesRequest) SetPage(v int32) *ListIndexTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListIndexTemplatesRequest) SetSize(v int32) *ListIndexTemplatesRequest {
	s.Size = &v
	return s
}

func (s *ListIndexTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

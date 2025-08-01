// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNamespacesShrinkRequest
	GetAcceptLanguage() *string
	SetName(v string) *ListNamespacesShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListNamespacesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNamespacesShrinkRequest
	GetPageSize() *int32
	SetRegion(v string) *ListNamespacesShrinkRequest
	GetRegion() *string
	SetTagShrink(v string) *ListNamespacesShrinkRequest
	GetTagShrink() *string
}

type ListNamespacesShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// myNamespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListNamespacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNamespacesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListNamespacesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNamespacesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListNamespacesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListNamespacesShrinkRequest) SetAcceptLanguage(v string) *ListNamespacesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNamespacesShrinkRequest) SetName(v string) *ListNamespacesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListNamespacesShrinkRequest) SetPageNumber(v int32) *ListNamespacesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNamespacesShrinkRequest) SetPageSize(v int32) *ListNamespacesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesShrinkRequest) SetRegion(v string) *ListNamespacesShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListNamespacesShrinkRequest) SetTagShrink(v string) *ListNamespacesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListNamespacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

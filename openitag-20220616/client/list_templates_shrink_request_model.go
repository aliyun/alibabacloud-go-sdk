// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplatesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesShrinkRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListTemplatesShrinkRequest
	GetSearchKey() *string
	SetTypesShrink(v string) *ListTemplatesShrinkRequest
	GetTypesShrink() *string
}

type ListTemplatesShrinkRequest struct {
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search content
	//
	// example:
	//
	// demo
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// List of application types.
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesShrinkRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListTemplatesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListTemplatesShrinkRequest) SetPageNumber(v int32) *ListTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetPageSize(v int32) *ListTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSearchKey(v string) *ListTemplatesShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTypesShrink(v string) *ListTemplatesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

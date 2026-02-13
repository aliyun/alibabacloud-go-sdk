// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceGetListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *AISearchResourceGetListShrinkRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *AISearchResourceGetListShrinkRequest
	GetPageSize() *int32
	SetResourceIdsShrink(v string) *AISearchResourceGetListShrinkRequest
	GetResourceIdsShrink() *string
	SetType(v string) *AISearchResourceGetListShrinkRequest
	GetType() *string
}

type AISearchResourceGetListShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 20
	PageSize          *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceIdsShrink *string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceGetListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceGetListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AISearchResourceGetListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *AISearchResourceGetListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AISearchResourceGetListShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *AISearchResourceGetListShrinkRequest) GetType() *string {
	return s.Type
}

func (s *AISearchResourceGetListShrinkRequest) SetCurrentPage(v int32) *AISearchResourceGetListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *AISearchResourceGetListShrinkRequest) SetPageSize(v int32) *AISearchResourceGetListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *AISearchResourceGetListShrinkRequest) SetResourceIdsShrink(v string) *AISearchResourceGetListShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *AISearchResourceGetListShrinkRequest) SetType(v string) *AISearchResourceGetListShrinkRequest {
	s.Type = &v
	return s
}

func (s *AISearchResourceGetListShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *InsertContentWithOptionsShrinkRequest
	GetContentShrink() *string
	SetDocumentId(v string) *InsertContentWithOptionsShrinkRequest
	GetDocumentId() *string
	SetIndex(v int32) *InsertContentWithOptionsShrinkRequest
	GetIndex() *int32
	SetPathShrink(v string) *InsertContentWithOptionsShrinkRequest
	GetPathShrink() *string
	SetTenantContextShrink(v string) *InsertContentWithOptionsShrinkRequest
	GetTenantContextShrink() *string
}

type InsertContentWithOptionsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// content
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// documentId
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// [0,0]
	PathShrink          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s InsertContentWithOptionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *InsertContentWithOptionsShrinkRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *InsertContentWithOptionsShrinkRequest) GetIndex() *int32 {
	return s.Index
}

func (s *InsertContentWithOptionsShrinkRequest) GetPathShrink() *string {
	return s.PathShrink
}

func (s *InsertContentWithOptionsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InsertContentWithOptionsShrinkRequest) SetContentShrink(v string) *InsertContentWithOptionsShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *InsertContentWithOptionsShrinkRequest) SetDocumentId(v string) *InsertContentWithOptionsShrinkRequest {
	s.DocumentId = &v
	return s
}

func (s *InsertContentWithOptionsShrinkRequest) SetIndex(v int32) *InsertContentWithOptionsShrinkRequest {
	s.Index = &v
	return s
}

func (s *InsertContentWithOptionsShrinkRequest) SetPathShrink(v string) *InsertContentWithOptionsShrinkRequest {
	s.PathShrink = &v
	return s
}

func (s *InsertContentWithOptionsShrinkRequest) SetTenantContextShrink(v string) *InsertContentWithOptionsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertContentWithOptionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

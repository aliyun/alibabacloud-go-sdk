// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *UploadBookShrinkRequest
	GetCategoryId() *string
	SetDocsShrink(v string) *UploadBookShrinkRequest
	GetDocsShrink() *string
	SetWorkspaceId(v string) *UploadBookShrinkRequest
	GetWorkspaceId() *string
}

type UploadBookShrinkRequest struct {
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	DocsShrink *string `json:"Docs,omitempty" xml:"Docs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ipe7d81yq4sl5jmk
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadBookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadBookShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadBookShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *UploadBookShrinkRequest) GetDocsShrink() *string {
	return s.DocsShrink
}

func (s *UploadBookShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadBookShrinkRequest) SetCategoryId(v string) *UploadBookShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UploadBookShrinkRequest) SetDocsShrink(v string) *UploadBookShrinkRequest {
	s.DocsShrink = &v
	return s
}

func (s *UploadBookShrinkRequest) SetWorkspaceId(v string) *UploadBookShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadBookShrinkRequest) Validate() error {
	return dara.Validate(s)
}

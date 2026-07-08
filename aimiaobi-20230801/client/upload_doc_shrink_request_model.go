// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *UploadDocShrinkRequest
	GetCategoryId() *string
	SetDocsShrink(v string) *UploadDocShrinkRequest
	GetDocsShrink() *string
	SetWorkspaceId(v string) *UploadDocShrinkRequest
	GetWorkspaceId() *string
}

type UploadDocShrinkRequest struct {
	// Folder where the document resides. If no value is provided, it defaults to "default".
	//
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Document
	//
	// This parameter is required.
	DocsShrink *string `json:"Docs,omitempty" xml:"Docs,omitempty"`
	// Unique identifier (UUID) of the Alibaba Cloud Model Studio workspace: obtain the [Workspace ID](https://help.aliyun.com/document_detail/2587495.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-yigtrrjl377rcbab
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadDocShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *UploadDocShrinkRequest) GetDocsShrink() *string {
	return s.DocsShrink
}

func (s *UploadDocShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadDocShrinkRequest) SetCategoryId(v string) *UploadDocShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UploadDocShrinkRequest) SetDocsShrink(v string) *UploadDocShrinkRequest {
	s.DocsShrink = &v
	return s
}

func (s *UploadDocShrinkRequest) SetWorkspaceId(v string) *UploadDocShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}

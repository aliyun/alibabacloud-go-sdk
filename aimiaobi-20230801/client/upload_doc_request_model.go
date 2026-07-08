// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *UploadDocRequest
	GetCategoryId() *string
	SetDocs(v []*UploadDocRequestDocs) *UploadDocRequest
	GetDocs() []*UploadDocRequestDocs
	SetWorkspaceId(v string) *UploadDocRequest
	GetWorkspaceId() *string
}

type UploadDocRequest struct {
	// Folder where the document resides. If no value is provided, it defaults to "default".
	//
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Document
	//
	// This parameter is required.
	Docs []*UploadDocRequestDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
	// Unique identifier (UUID) of the Alibaba Cloud Model Studio workspace: obtain the [Workspace ID](https://help.aliyun.com/document_detail/2587495.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-yigtrrjl377rcbab
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadDocRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocRequest) GoString() string {
	return s.String()
}

func (s *UploadDocRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *UploadDocRequest) GetDocs() []*UploadDocRequestDocs {
	return s.Docs
}

func (s *UploadDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadDocRequest) SetCategoryId(v string) *UploadDocRequest {
	s.CategoryId = &v
	return s
}

func (s *UploadDocRequest) SetDocs(v []*UploadDocRequestDocs) *UploadDocRequest {
	s.Docs = v
	return s
}

func (s *UploadDocRequest) SetWorkspaceId(v string) *UploadDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadDocRequest) Validate() error {
	if s.Docs != nil {
		for _, item := range s.Docs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadDocRequestDocs struct {
	// document Name
	//
	// This parameter is required.
	//
	// example:
	//
	// 文档1.pdf
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// file URL
	//
	// This parameter is required.
	//
	// example:
	//
	// http://xxx/ccc.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s UploadDocRequestDocs) String() string {
	return dara.Prettify(s)
}

func (s UploadDocRequestDocs) GoString() string {
	return s.String()
}

func (s *UploadDocRequestDocs) GetDocName() *string {
	return s.DocName
}

func (s *UploadDocRequestDocs) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadDocRequestDocs) SetDocName(v string) *UploadDocRequestDocs {
	s.DocName = &v
	return s
}

func (s *UploadDocRequestDocs) SetFileUrl(v string) *UploadDocRequestDocs {
	s.FileUrl = &v
	return s
}

func (s *UploadDocRequestDocs) Validate() error {
	return dara.Validate(s)
}

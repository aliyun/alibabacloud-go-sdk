// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *UploadBookRequest
	GetCategoryId() *string
	SetDocs(v []*UploadBookRequestDocs) *UploadBookRequest
	GetDocs() []*UploadBookRequestDocs
	SetWorkspaceId(v string) *UploadBookRequest
	GetWorkspaceId() *string
}

type UploadBookRequest struct {
	// Folder ID
	//
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Documents
	//
	// This parameter is required.
	Docs []*UploadBookRequestDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
	// Unique identifier of your Alibaba Cloud Model Studio workspace. [Get your workspace ID](https://help.aliyun.com/document_detail/2782167.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-ipe7d81yq4sl5jmk
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadBookRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadBookRequest) GoString() string {
	return s.String()
}

func (s *UploadBookRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *UploadBookRequest) GetDocs() []*UploadBookRequestDocs {
	return s.Docs
}

func (s *UploadBookRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadBookRequest) SetCategoryId(v string) *UploadBookRequest {
	s.CategoryId = &v
	return s
}

func (s *UploadBookRequest) SetDocs(v []*UploadBookRequestDocs) *UploadBookRequest {
	s.Docs = v
	return s
}

func (s *UploadBookRequest) SetWorkspaceId(v string) *UploadBookRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadBookRequest) Validate() error {
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

type UploadBookRequestDocs struct {
	// Document name
	//
	// example:
	//
	// 文档1.pdf
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// File URL
	//
	// example:
	//
	// http://xxx/ccc.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s UploadBookRequestDocs) String() string {
	return dara.Prettify(s)
}

func (s UploadBookRequestDocs) GoString() string {
	return s.String()
}

func (s *UploadBookRequestDocs) GetDocName() *string {
	return s.DocName
}

func (s *UploadBookRequestDocs) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadBookRequestDocs) SetDocName(v string) *UploadBookRequestDocs {
	s.DocName = &v
	return s
}

func (s *UploadBookRequestDocs) SetFileUrl(v string) *UploadBookRequestDocs {
	s.FileUrl = &v
	return s
}

func (s *UploadBookRequestDocs) Validate() error {
	return dara.Validate(s)
}

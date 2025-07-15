// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentLengthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocName(v string) *GetFileContentLengthRequest
	GetDocName() *string
	SetFileUrl(v string) *GetFileContentLengthRequest
	GetFileUrl() *string
	SetWorkspaceId(v string) *GetFileContentLengthRequest
	GetWorkspaceId() *string
}

type GetFileContentLengthRequest struct {
	// example:
	//
	// test.pdf
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// https://xxx/test.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetFileContentLengthRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentLengthRequest) GoString() string {
	return s.String()
}

func (s *GetFileContentLengthRequest) GetDocName() *string {
	return s.DocName
}

func (s *GetFileContentLengthRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetFileContentLengthRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetFileContentLengthRequest) SetDocName(v string) *GetFileContentLengthRequest {
	s.DocName = &v
	return s
}

func (s *GetFileContentLengthRequest) SetFileUrl(v string) *GetFileContentLengthRequest {
	s.FileUrl = &v
	return s
}

func (s *GetFileContentLengthRequest) SetWorkspaceId(v string) *GetFileContentLengthRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetFileContentLengthRequest) Validate() error {
	return dara.Validate(s)
}

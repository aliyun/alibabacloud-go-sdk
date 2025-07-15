// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportTermsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *SubmitImportTermsTaskRequest
	GetFileKey() *string
	SetWorkspaceId(v string) *SubmitImportTermsTaskRequest
	GetWorkspaceId() *string
}

type SubmitImportTermsTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitImportTermsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportTermsTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitImportTermsTaskRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitImportTermsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitImportTermsTaskRequest) SetFileKey(v string) *SubmitImportTermsTaskRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitImportTermsTaskRequest) SetWorkspaceId(v string) *SubmitImportTermsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitImportTermsTaskRequest) Validate() error {
	return dara.Validate(s)
}

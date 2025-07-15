// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImportTermsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *FetchImportTermsTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *FetchImportTermsTaskRequest
	GetWorkspaceId() *string
}

type FetchImportTermsTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FetchImportTermsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchImportTermsTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchImportTermsTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchImportTermsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *FetchImportTermsTaskRequest) SetTaskId(v string) *FetchImportTermsTaskRequest {
	s.TaskId = &v
	return s
}

func (s *FetchImportTermsTaskRequest) SetWorkspaceId(v string) *FetchImportTermsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *FetchImportTermsTaskRequest) Validate() error {
	return dara.Validate(s)
}

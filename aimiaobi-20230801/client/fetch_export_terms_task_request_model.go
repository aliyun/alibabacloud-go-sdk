// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportTermsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *FetchExportTermsTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *FetchExportTermsTaskRequest
	GetWorkspaceId() *string
}

type FetchExportTermsTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FetchExportTermsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchExportTermsTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchExportTermsTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchExportTermsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *FetchExportTermsTaskRequest) SetTaskId(v string) *FetchExportTermsTaskRequest {
	s.TaskId = &v
	return s
}

func (s *FetchExportTermsTaskRequest) SetWorkspaceId(v string) *FetchExportTermsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *FetchExportTermsTaskRequest) Validate() error {
	return dara.Validate(s)
}

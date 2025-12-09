// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetNotebookTaskStatusRequest
	GetSessionId() *string
	SetTaskId(v string) *GetNotebookTaskStatusRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetNotebookTaskStatusRequest
	GetWorkspaceId() *string
}

type GetNotebookTaskStatusRequest struct {
	// example:
	//
	// 8141456676986429894916354
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// c2b4cb5a-7420-49a8-aa7c-528becd6e1bf
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetNotebookTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetNotebookTaskStatusRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetNotebookTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetNotebookTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetNotebookTaskStatusRequest) SetSessionId(v string) *GetNotebookTaskStatusRequest {
	s.SessionId = &v
	return s
}

func (s *GetNotebookTaskStatusRequest) SetTaskId(v string) *GetNotebookTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetNotebookTaskStatusRequest) SetWorkspaceId(v string) *GetNotebookTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetNotebookTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}

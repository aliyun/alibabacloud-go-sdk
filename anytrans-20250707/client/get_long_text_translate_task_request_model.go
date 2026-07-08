// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLongTextTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetLongTextTranslateTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetLongTextTranslateTaskRequest
	GetWorkspaceId() *string
}

type GetLongTextTranslateTaskRequest struct {
	// The ID of the long-text translation task.
	//
	// example:
	//
	// a8f25f25-0b36-4349-857f-e19a43f69e51
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The ID of the Model Studio workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddri
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetLongTextTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLongTextTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetLongTextTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetLongTextTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetLongTextTranslateTaskRequest) SetTaskId(v string) *GetLongTextTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetLongTextTranslateTaskRequest) SetWorkspaceId(v string) *GetLongTextTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetLongTextTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}

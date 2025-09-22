// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDocTranslateTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetDocTranslateTaskRequest
	GetWorkspaceId() *string
}

type GetDocTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d3a2397bc2c14ab4a2e40a4f5b46241b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetDocTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDocTranslateTaskRequest) SetTaskId(v string) *GetDocTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskRequest) SetWorkspaceId(v string) *GetDocTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDocTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}

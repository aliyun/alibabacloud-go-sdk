// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetImageTranslateTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetImageTranslateTaskRequest
	GetWorkspaceId() *string
}

type GetImageTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2746f4be-cff2-465e-a2c6-12bff30ce0f9
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetImageTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetImageTranslateTaskRequest) SetTaskId(v string) *GetImageTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageTranslateTaskRequest) SetWorkspaceId(v string) *GetImageTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetImageTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHtmlTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetHtmlTranslateTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetHtmlTranslateTaskRequest
	GetWorkspaceId() *string
}

type GetHtmlTranslateTaskRequest struct {
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

func (s GetHtmlTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHtmlTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetHtmlTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetHtmlTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetHtmlTranslateTaskRequest) SetTaskId(v string) *GetHtmlTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetHtmlTranslateTaskRequest) SetWorkspaceId(v string) *GetHtmlTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetHtmlTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}

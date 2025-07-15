// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartClipTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSmartClipTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetSmartClipTaskRequest
	GetWorkspaceId() *string
}

type GetSmartClipTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// task-03d46184ee7d8749
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetSmartClipTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSmartClipTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSmartClipTaskRequest) SetTaskId(v string) *GetSmartClipTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetSmartClipTaskRequest) SetWorkspaceId(v string) *GetSmartClipTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSmartClipTaskRequest) Validate() error {
	return dara.Validate(s)
}

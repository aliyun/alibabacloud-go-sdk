// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDeepWriteTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetDeepWriteTaskRequest
	GetWorkspaceId() *string
}

type GetDeepWriteTaskRequest struct {
	// example:
	//
	// 6d3c0bc9-7561-41a4-be4c-d906abdb40a9
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-odl2p61i4vfbph4g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDeepWriteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeepWriteTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDeepWriteTaskRequest) SetTaskId(v string) *GetDeepWriteTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetDeepWriteTaskRequest) SetWorkspaceId(v string) *GetDeepWriteTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDeepWriteTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeepWriteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CancelDeepWriteTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *CancelDeepWriteTaskRequest
	GetWorkspaceId() *string
}

type CancelDeepWriteTaskRequest struct {
	// example:
	//
	// xbabac91-fdad-44d6-95ce-******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelDeepWriteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDeepWriteTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelDeepWriteTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelDeepWriteTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelDeepWriteTaskRequest) SetTaskId(v string) *CancelDeepWriteTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelDeepWriteTaskRequest) SetWorkspaceId(v string) *CancelDeepWriteTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelDeepWriteTaskRequest) Validate() error {
	return dara.Validate(s)
}

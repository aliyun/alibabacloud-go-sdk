// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDeepWriteTaskResultRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetDeepWriteTaskResultRequest
	GetWorkspaceId() *string
}

type GetDeepWriteTaskResultRequest struct {
	// example:
	//
	// xbabac91-fdad-44d6-95ce-******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-1setzb9xb8m11vrc
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDeepWriteTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeepWriteTaskResultRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDeepWriteTaskResultRequest) SetTaskId(v string) *GetDeepWriteTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDeepWriteTaskResultRequest) SetWorkspaceId(v string) *GetDeepWriteTaskResultRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDeepWriteTaskResultRequest) Validate() error {
	return dara.Validate(s)
}

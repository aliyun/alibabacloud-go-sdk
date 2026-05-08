// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdempotentId(v string) *GetProjectTaskRequest
	GetIdempotentId() *string
	SetTaskId(v string) *GetProjectTaskRequest
	GetTaskId() *string
}

type GetProjectTaskRequest struct {
	// example:
	//
	// 20230823218109326025-1200
	IdempotentId *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetProjectTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *GetProjectTaskRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *GetProjectTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetProjectTaskRequest) SetIdempotentId(v string) *GetProjectTaskRequest {
	s.IdempotentId = &v
	return s
}

func (s *GetProjectTaskRequest) SetTaskId(v string) *GetProjectTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetProjectTaskRequest) Validate() error {
	return dara.Validate(s)
}

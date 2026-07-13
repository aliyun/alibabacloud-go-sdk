// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateInstanceAsyncTaskRequest
	GetInstanceId() *string
	SetIsResume(v bool) *UpdateInstanceAsyncTaskRequest
	GetIsResume() *bool
	SetTaskCode(v string) *UpdateInstanceAsyncTaskRequest
	GetTaskCode() *string
	SetTaskId(v string) *UpdateInstanceAsyncTaskRequest
	GetTaskId() *string
}

type UpdateInstanceAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// at-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsResume   *bool   `json:"IsResume,omitempty" xml:"IsResume,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentteams:pay-order:create
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// task-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateInstanceAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAsyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceAsyncTaskRequest) GetIsResume() *bool {
	return s.IsResume
}

func (s *UpdateInstanceAsyncTaskRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *UpdateInstanceAsyncTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceAsyncTaskRequest) SetInstanceId(v string) *UpdateInstanceAsyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceAsyncTaskRequest) SetIsResume(v bool) *UpdateInstanceAsyncTaskRequest {
	s.IsResume = &v
	return s
}

func (s *UpdateInstanceAsyncTaskRequest) SetTaskCode(v string) *UpdateInstanceAsyncTaskRequest {
	s.TaskCode = &v
	return s
}

func (s *UpdateInstanceAsyncTaskRequest) SetTaskId(v string) *UpdateInstanceAsyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}

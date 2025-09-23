// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *CreateAsyncTaskRequest
	GetResourceGroupId() *string
	SetTaskParams(v string) *CreateAsyncTaskRequest
	GetTaskParams() *string
	SetTaskType(v int32) *CreateAsyncTaskRequest
	GetTaskType() *int32
}

type CreateAsyncTaskRequest struct {
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"timestamp": 1530276554, "instanceId": "ddoscoo-woieuroi234"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAsyncTaskRequest) GetTaskParams() *string {
	return s.TaskParams
}

func (s *CreateAsyncTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int32) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}

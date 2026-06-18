// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteAiOutboundTaskRequest
	GetInstanceId() *string
	SetTaskId(v int64) *DeleteAiOutboundTaskRequest
	GetTaskId() *int64
}

type DeleteAiOutboundTaskRequest struct {
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// You can call the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and check the **Data*	- field in the response, or call the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and check the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAiOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeleteAiOutboundTaskRequest) SetInstanceId(v string) *DeleteAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAiOutboundTaskRequest) SetTaskId(v int64) *DeleteAiOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteAiOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAiOutboundTaskDetailRequest
	GetInstanceId() *string
	SetTaskId(v int64) *GetAiOutboundTaskDetailRequest
	GetTaskId() *int64
}

type GetAiOutboundTaskDetailRequest struct {
	// The ID of the Artificial Intelligence Cloud Call Service (AICCS) instance.
	//
	// You can obtain the instance ID from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task ID.
	//
	// You can call the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) operation and check the **Data*	- parameter in the response, or call the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) operation and check the **TaskId*	- parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAiOutboundTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAiOutboundTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskDetailRequest) SetInstanceId(v string) *GetAiOutboundTaskDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAiOutboundTaskDetailRequest) SetTaskId(v int64) *GetAiOutboundTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}

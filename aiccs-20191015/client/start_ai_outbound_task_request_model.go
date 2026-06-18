// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartAiOutboundTaskRequest
	GetInstanceId() *string
	SetTaskId(v int64) *StartAiOutboundTaskRequest
	GetTaskId() *int64
}

type StartAiOutboundTaskRequest struct {
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// You can invoke the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and view the **Data*	- field in the response, or invoke the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and view the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1763****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *StartAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartAiOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartAiOutboundTaskRequest) SetInstanceId(v string) *StartAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StartAiOutboundTaskRequest) SetTaskId(v int64) *StartAiOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartAiOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

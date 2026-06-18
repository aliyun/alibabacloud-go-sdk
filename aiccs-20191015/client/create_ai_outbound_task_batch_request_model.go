// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateAiOutboundTaskBatchRequest
	GetInstanceId() *string
	SetTaskId(v int64) *CreateAiOutboundTaskBatchRequest
	GetTaskId() *int64
}

type CreateAiOutboundTaskBatchRequest struct {
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
	// You can invoke the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and view the **Data*	- field in the response, or invoke the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and view the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAiOutboundTaskBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskBatchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAiOutboundTaskBatchRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateAiOutboundTaskBatchRequest) SetInstanceId(v string) *CreateAiOutboundTaskBatchRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAiOutboundTaskBatchRequest) SetTaskId(v int64) *CreateAiOutboundTaskBatchRequest {
	s.TaskId = &v
	return s
}

func (s *CreateAiOutboundTaskBatchRequest) Validate() error {
	return dara.Validate(s)
}

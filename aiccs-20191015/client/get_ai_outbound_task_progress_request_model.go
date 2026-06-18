// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchVersion(v int32) *GetAiOutboundTaskProgressRequest
	GetBatchVersion() *int32
	SetInstanceId(v string) *GetAiOutboundTaskProgressRequest
	GetInstanceId() *string
	SetTaskId(v int64) *GetAiOutboundTaskProgressRequest
	GetTaskId() *int64
}

type GetAiOutboundTaskProgressRequest struct {
	// Job batch.
	//
	// > If empty, queries all data under the job.
	//
	// example:
	//
	// 1
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it in the <b>Instance Management</b> section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID.
	//
	// You can invoke the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and check the **Data*	- field in the response, or invoke the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and check the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAiOutboundTaskProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressRequest) GetBatchVersion() *int32 {
	return s.BatchVersion
}

func (s *GetAiOutboundTaskProgressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAiOutboundTaskProgressRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskProgressRequest) SetBatchVersion(v int32) *GetAiOutboundTaskProgressRequest {
	s.BatchVersion = &v
	return s
}

func (s *GetAiOutboundTaskProgressRequest) SetInstanceId(v string) *GetAiOutboundTaskProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAiOutboundTaskProgressRequest) SetTaskId(v int64) *GetAiOutboundTaskProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskProgressRequest) Validate() error {
	return dara.Validate(s)
}

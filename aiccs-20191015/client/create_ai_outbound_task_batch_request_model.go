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
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopAiOutboundTaskRequest
	GetInstanceId() *string
	SetTaskId(v int64) *StopAiOutboundTaskRequest
	GetTaskId() *int64
}

type StopAiOutboundTaskRequest struct {
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

func (s StopAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopAiOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StopAiOutboundTaskRequest) SetInstanceId(v string) *StopAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StopAiOutboundTaskRequest) SetTaskId(v int64) *StopAiOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopAiOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

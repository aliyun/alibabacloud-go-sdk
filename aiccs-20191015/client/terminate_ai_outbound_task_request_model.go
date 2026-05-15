// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *TerminateAiOutboundTaskRequest
	GetInstanceId() *string
	SetTaskId(v int64) *TerminateAiOutboundTaskRequest
	GetTaskId() *int64
}

type TerminateAiOutboundTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// agent_****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1763****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TerminateAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *TerminateAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateAiOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TerminateAiOutboundTaskRequest) SetInstanceId(v string) *TerminateAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateAiOutboundTaskRequest) SetTaskId(v int64) *TerminateAiOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *TerminateAiOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}

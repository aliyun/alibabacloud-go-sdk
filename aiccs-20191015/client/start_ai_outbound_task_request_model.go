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

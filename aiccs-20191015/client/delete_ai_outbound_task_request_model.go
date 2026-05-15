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

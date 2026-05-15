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
	// example:
	//
	// 1
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
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

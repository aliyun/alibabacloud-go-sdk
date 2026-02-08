// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInflightTaskTimeoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *InflightTaskTimeoutRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *InflightTaskTimeoutRequest
	GetInstanceOwnerId() *int64
	SetTaskId(v string) *InflightTaskTimeoutRequest
	GetTaskId() *string
}

type InflightTaskTimeoutRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2d350e38-f561-49b0-85d3-b90d9fc7e052
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Alibaba Cloud account ID to which the instance belongs
	//
	// example:
	//
	// 1864632921948620
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// Job ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bc0585c-0a8a-46d8-b042-23570bbb4855
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InflightTaskTimeoutRequest) String() string {
	return dara.Prettify(s)
}

func (s InflightTaskTimeoutRequest) GoString() string {
	return s.String()
}

func (s *InflightTaskTimeoutRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InflightTaskTimeoutRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *InflightTaskTimeoutRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *InflightTaskTimeoutRequest) SetInstanceId(v string) *InflightTaskTimeoutRequest {
	s.InstanceId = &v
	return s
}

func (s *InflightTaskTimeoutRequest) SetInstanceOwnerId(v int64) *InflightTaskTimeoutRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *InflightTaskTimeoutRequest) SetTaskId(v string) *InflightTaskTimeoutRequest {
	s.TaskId = &v
	return s
}

func (s *InflightTaskTimeoutRequest) Validate() error {
	return dara.Validate(s)
}

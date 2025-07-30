// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceParameterResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ModifyInstanceParameterResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *ModifyInstanceParameterResponseBody
	GetTaskId() *int64
}

type ModifyInstanceParameterResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 578678678
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyInstanceParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceParameterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceParameterResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyInstanceParameterResponseBody) SetInstanceId(v string) *ModifyInstanceParameterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceParameterResponseBody) SetRequestId(v string) *ModifyInstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceParameterResponseBody) SetTaskId(v int64) *ModifyInstanceParameterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyInstanceParameterResponseBody) Validate() error {
	return dara.Validate(s)
}

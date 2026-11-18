// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTairAIGatewayParameterResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ModifyTairAIGatewayParameterResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *ModifyTairAIGatewayParameterResponseBody
	GetTaskId() *int64
}

type ModifyTairAIGatewayParameterResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// tg-2zeb1ce76fee****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created task flow.
	//
	// example:
	//
	// 578678678
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyTairAIGatewayParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayParameterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTairAIGatewayParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTairAIGatewayParameterResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyTairAIGatewayParameterResponseBody) SetInstanceId(v string) *ModifyTairAIGatewayParameterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyTairAIGatewayParameterResponseBody) SetRequestId(v string) *ModifyTairAIGatewayParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTairAIGatewayParameterResponseBody) SetTaskId(v int64) *ModifyTairAIGatewayParameterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyTairAIGatewayParameterResponseBody) Validate() error {
	return dara.Validate(s)
}

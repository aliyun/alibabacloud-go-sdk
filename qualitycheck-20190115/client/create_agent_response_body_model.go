// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentResponseBody
	GetCode() *string
	SetData(v int64) *CreateAgentResponseBody
	GetData() *int64
	SetMessage(v string) *CreateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentResponseBody
	GetSuccess() *bool
}

type CreateAgentResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The AgentId of the task.
	//
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message, if any.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentResponseBody) SetCode(v string) *CreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentResponseBody) SetData(v int64) *CreateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentResponseBody) SetMessage(v string) *CreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentResponseBody
	GetCode() *string
	SetData(v *UpdateAgentResponseBodyData) *UpdateAgentResponseBody
	GetData() *UpdateAgentResponseBodyData
	SetMessage(v string) *UpdateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentResponseBody
	GetSuccess() *bool
}

type UpdateAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidArgument
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 382E6272-XXXXX-A8AF0BFAC1A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentResponseBody) GetData() *UpdateAgentResponseBodyData {
	return s.Data
}

func (s *UpdateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetData(v *UpdateAgentResponseBodyData) *UpdateAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentResponseBodyData struct {
	// Agent ARN
	//
	// example:
	//
	// acs:eventbridge:{region}:{accountId}:agent/{agentName}
	AgentArn *string `json:"AgentArn,omitempty" xml:"AgentArn,omitempty"`
}

func (s UpdateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBodyData) GetAgentArn() *string {
	return s.AgentArn
}

func (s *UpdateAgentResponseBodyData) SetAgentArn(v string) *UpdateAgentResponseBodyData {
	s.AgentArn = &v
	return s
}

func (s *UpdateAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

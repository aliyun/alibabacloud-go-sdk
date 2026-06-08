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
	SetData(v *CreateAgentResponseBodyData) *CreateAgentResponseBody
	GetData() *CreateAgentResponseBodyData
	SetMessage(v string) *CreateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentResponseBody
	GetSuccess() *bool
}

type CreateAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Agent with name \\"XXX\\" already exists for account 12345
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B896B484-XXXXXX-DD0E5C361108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateAgentResponseBody) GetData() *CreateAgentResponseBodyData {
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

func (s *CreateAgentResponseBody) SetData(v *CreateAgentResponseBodyData) *CreateAgentResponseBody {
	s.Data = v
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResponseBodyData struct {
	// Agent ARN
	//
	// example:
	//
	// acs:eventbridge:{region}:{accountId}:agent/{agentName}
	AgentArn *string `json:"AgentArn,omitempty" xml:"AgentArn,omitempty"`
}

func (s CreateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyData) GetAgentArn() *string {
	return s.AgentArn
}

func (s *CreateAgentResponseBodyData) SetAgentArn(v string) *CreateAgentResponseBodyData {
	s.AgentArn = &v
	return s
}

func (s *CreateAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

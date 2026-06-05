// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *CreateAgentResponseBodyAgent) *CreateAgentResponseBody
	GetAgent() *CreateAgentResponseBodyAgent
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
}

type CreateAgentResponseBody struct {
	Agent *CreateAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetAgent() *CreateAgentResponseBodyAgent {
	return s.Agent
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) SetAgent(v *CreateAgentResponseBodyAgent) *CreateAgentResponseBody {
	s.Agent = v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResponseBodyAgent struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyAgent) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateAgentResponseBodyAgent) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *CreateAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *CreateAgentResponseBodyAgent) SetGmtCreateTime(v string) *CreateAgentResponseBodyAgent {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetGmtModifiedTime(v string) *CreateAgentResponseBodyAgent {
	s.GmtModifiedTime = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) SetName(v string) *CreateAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *CreateAgentResponseBodyAgent) Validate() error {
	return dara.Validate(s)
}

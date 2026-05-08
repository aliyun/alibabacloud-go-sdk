// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateAgentResponseBody
	GetAgentId() *string
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
}

type CreateAgentResponseBody struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) SetAgentId(v string) *CreateAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentTeamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ListAgentTeamsRequestBody) *ListAgentTeamsRequest
	GetBody() *ListAgentTeamsRequestBody
}

type ListAgentTeamsRequest struct {
	// The request parameters for querying the agent team list.
	Body *ListAgentTeamsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ListAgentTeamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsRequest) GetBody() *ListAgentTeamsRequestBody {
	return s.Body
}

func (s *ListAgentTeamsRequest) SetBody(v *ListAgentTeamsRequestBody) *ListAgentTeamsRequest {
	s.Body = v
	return s
}

func (s *ListAgentTeamsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentTeamsRequestBody struct {
	// The list of agent IDs for which to query team information.
	//
	// This parameter is required.
	AgentIds []*string `json:"agentIds,omitempty" xml:"agentIds,omitempty" type:"Repeated"`
}

func (s ListAgentTeamsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsRequestBody) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsRequestBody) GetAgentIds() []*string {
	return s.AgentIds
}

func (s *ListAgentTeamsRequestBody) SetAgentIds(v []*string) *ListAgentTeamsRequestBody {
	s.AgentIds = v
	return s
}

func (s *ListAgentTeamsRequestBody) Validate() error {
	return dara.Validate(s)
}

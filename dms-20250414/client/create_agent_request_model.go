// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CreateAgentRequest
	GetAgentName() *string
	SetAgentType(v string) *CreateAgentRequest
	GetAgentType() *string
	SetDescription(v string) *CreateAgentRequest
	GetDescription() *string
	SetExpireAfterSeconds(v int32) *CreateAgentRequest
	GetExpireAfterSeconds() *int32
}

type CreateAgentRequest struct {
	// The agent name. The name must be unique within the same tenant. Maximum length: 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// order-analysis-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The permission inheritance type of the agent, which specifies the permission source. Default value: HUMAN_BOUND.
	//
	// example:
	//
	// HUMAN_BOUND
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The description of the agent. Maximum length: 512 characters.
	//
	// example:
	//
	// An agent for querying and analyzing order data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The validity period of the automatically issued API key, in seconds. Valid values: 1 to 31536000 (up to 1 year).
	//
	// example:
	//
	// 2592000
	ExpireAfterSeconds *int32 `json:"ExpireAfterSeconds,omitempty" xml:"ExpireAfterSeconds,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAgentRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRequest) GetExpireAfterSeconds() *int32 {
	return s.ExpireAfterSeconds
}

func (s *CreateAgentRequest) SetAgentName(v string) *CreateAgentRequest {
	s.AgentName = &v
	return s
}

func (s *CreateAgentRequest) SetAgentType(v string) *CreateAgentRequest {
	s.AgentType = &v
	return s
}

func (s *CreateAgentRequest) SetDescription(v string) *CreateAgentRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentRequest) SetExpireAfterSeconds(v int32) *CreateAgentRequest {
	s.ExpireAfterSeconds = &v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	return dara.Validate(s)
}

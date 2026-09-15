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
	// This parameter is required.
	AgentName          *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentType          *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpireAfterSeconds *int32  `json:"ExpireAfterSeconds,omitempty" xml:"ExpireAfterSeconds,omitempty"`
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

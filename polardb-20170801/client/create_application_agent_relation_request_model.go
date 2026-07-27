// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAgentRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateApplicationAgentRelationRequest
	GetAgentId() *string
	SetApplicationId(v string) *CreateApplicationAgentRelationRequest
	GetApplicationId() *string
	SetToken(v string) *CreateApplicationAgentRelationRequest
	GetToken() *string
}

type CreateApplicationAgentRelationRequest struct {
	// The instance ID of the Agent to attach.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The ID of the Squad application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The authentication token.
	//
	// This parameter is required.
	//
	// example:
	//
	// pas_xxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateApplicationAgentRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAgentRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationAgentRelationRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateApplicationAgentRelationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationAgentRelationRequest) GetToken() *string {
	return s.Token
}

func (s *CreateApplicationAgentRelationRequest) SetAgentId(v string) *CreateApplicationAgentRelationRequest {
	s.AgentId = &v
	return s
}

func (s *CreateApplicationAgentRelationRequest) SetApplicationId(v string) *CreateApplicationAgentRelationRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationAgentRelationRequest) SetToken(v string) *CreateApplicationAgentRelationRequest {
	s.Token = &v
	return s
}

func (s *CreateApplicationAgentRelationRequest) Validate() error {
	return dara.Validate(s)
}

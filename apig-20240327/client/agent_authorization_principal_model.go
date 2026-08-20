// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentAuthorizationPrincipal interface {
	dara.Model
	String() string
	GoString() string
	SetPrincipalId(v string) *AgentAuthorizationPrincipal
	GetPrincipalId() *string
	SetPrincipalType(v string) *AgentAuthorizationPrincipal
	GetPrincipalType() *string
}

type AgentAuthorizationPrincipal struct {
	// The ID of the authorization principal. Specify a consumer ID or consumer group ID based on the value of principalType.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumer-1
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// The type of the authorization principal. Valid values:
	//
	// - Consumer: consumer.
	//
	// - ConsumerGroup: consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Consumer
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
}

func (s AgentAuthorizationPrincipal) String() string {
	return dara.Prettify(s)
}

func (s AgentAuthorizationPrincipal) GoString() string {
	return s.String()
}

func (s *AgentAuthorizationPrincipal) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *AgentAuthorizationPrincipal) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AgentAuthorizationPrincipal) SetPrincipalId(v string) *AgentAuthorizationPrincipal {
	s.PrincipalId = &v
	return s
}

func (s *AgentAuthorizationPrincipal) SetPrincipalType(v string) *AgentAuthorizationPrincipal {
	s.PrincipalType = &v
	return s
}

func (s *AgentAuthorizationPrincipal) Validate() error {
	return dara.Validate(s)
}

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
	// This parameter is required.
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// This parameter is required.
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

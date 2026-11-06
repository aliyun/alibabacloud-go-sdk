// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreatePermissionRulesRequestBody) *CreatePermissionRulesRequest
	GetBody() *CreatePermissionRulesRequestBody
	SetClientToken(v string) *CreatePermissionRulesRequest
	GetClientToken() *string
}

type CreatePermissionRulesRequest struct {
	// The request body for creating permission rules.
	Body *CreatePermissionRulesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The idempotency token.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreatePermissionRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesRequest) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesRequest) GetBody() *CreatePermissionRulesRequestBody {
	return s.Body
}

func (s *CreatePermissionRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePermissionRulesRequest) SetBody(v *CreatePermissionRulesRequestBody) *CreatePermissionRulesRequest {
	s.Body = v
	return s
}

func (s *CreatePermissionRulesRequest) SetClientToken(v string) *CreatePermissionRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePermissionRulesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePermissionRulesRequestBody struct {
	// The list of permission rules to be created.
	//
	// This parameter is required.
	Rules []*CreatePermissionRulesRequestBodyRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s CreatePermissionRulesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesRequestBody) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesRequestBody) GetRules() []*CreatePermissionRulesRequestBodyRules {
	return s.Rules
}

func (s *CreatePermissionRulesRequestBody) SetRules(v []*CreatePermissionRulesRequestBodyRules) *CreatePermissionRulesRequestBody {
	s.Rules = v
	return s
}

func (s *CreatePermissionRulesRequestBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePermissionRulesRequestBodyRules struct {
	// The authorization effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// PERMIT
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The rule enforcement mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"enforcementMode,omitempty" xml:"enforcementMode,omitempty"`
	// The principal ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123456
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// The principal type.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-1234567890
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AGENT
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The rule source.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgentCore
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPECIFIED
	TargetScope *string `json:"targetScope,omitempty" xml:"targetScope,omitempty"`
	// The list of targets for the rule.
	Targets []*string `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s CreatePermissionRulesRequestBodyRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesRequestBodyRules) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesRequestBodyRules) GetEffect() *string {
	return s.Effect
}

func (s *CreatePermissionRulesRequestBodyRules) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *CreatePermissionRulesRequestBodyRules) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreatePermissionRulesRequestBodyRules) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreatePermissionRulesRequestBodyRules) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreatePermissionRulesRequestBodyRules) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreatePermissionRulesRequestBodyRules) GetSource() *string {
	return s.Source
}

func (s *CreatePermissionRulesRequestBodyRules) GetTargetScope() *string {
	return s.TargetScope
}

func (s *CreatePermissionRulesRequestBodyRules) GetTargets() []*string {
	return s.Targets
}

func (s *CreatePermissionRulesRequestBodyRules) SetEffect(v string) *CreatePermissionRulesRequestBodyRules {
	s.Effect = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetEnforcementMode(v string) *CreatePermissionRulesRequestBodyRules {
	s.EnforcementMode = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetPrincipalId(v string) *CreatePermissionRulesRequestBodyRules {
	s.PrincipalId = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetPrincipalType(v string) *CreatePermissionRulesRequestBodyRules {
	s.PrincipalType = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetResourceId(v string) *CreatePermissionRulesRequestBodyRules {
	s.ResourceId = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetResourceType(v string) *CreatePermissionRulesRequestBodyRules {
	s.ResourceType = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetSource(v string) *CreatePermissionRulesRequestBodyRules {
	s.Source = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetTargetScope(v string) *CreatePermissionRulesRequestBodyRules {
	s.TargetScope = &v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) SetTargets(v []*string) *CreatePermissionRulesRequestBodyRules {
	s.Targets = v
	return s
}

func (s *CreatePermissionRulesRequestBodyRules) Validate() error {
	return dara.Validate(s)
}

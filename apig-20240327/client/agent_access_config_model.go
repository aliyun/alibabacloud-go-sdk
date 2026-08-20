// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentAccessConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAiRequestLogEnabled(v bool) *AgentAccessConfig
	GetAiRequestLogEnabled() *bool
	SetAuthorization(v *AgentAccessConfigAuthorization) *AgentAccessConfig
	GetAuthorization() *AgentAccessConfigAuthorization
	SetBasePath(v string) *AgentAccessConfig
	GetBasePath() *string
	SetDomainIds(v []*string) *AgentAccessConfig
	GetDomainIds() []*string
	SetRemoveBasePathOnForward(v bool) *AgentAccessConfig
	GetRemoveBasePathOnForward() *bool
}

type AgentAccessConfig struct {
	// Specifies whether to enable AI request logging. Default value if omitted: true.
	//
	// example:
	//
	// false
	AiRequestLogEnabled *bool `json:"aiRequestLogEnabled,omitempty" xml:"aiRequestLogEnabled,omitempty"`
	// The consumer authorization configuration for Agent access. If omitted, consumer authorization is not enabled.
	Authorization *AgentAccessConfigAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty" type:"Struct"`
	// The base path of the Agent access entry. The path must start with a forward slash (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// /agent
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The list of domain name IDs bound to the Agent access entry. At least one domain name must be specified.
	//
	// This parameter is required.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Specifies whether to remove the base path when forwarding requests to the backend. Default value if omitted: false.
	//
	// example:
	//
	// true
	RemoveBasePathOnForward *bool `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
}

func (s AgentAccessConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentAccessConfig) GoString() string {
	return s.String()
}

func (s *AgentAccessConfig) GetAiRequestLogEnabled() *bool {
	return s.AiRequestLogEnabled
}

func (s *AgentAccessConfig) GetAuthorization() *AgentAccessConfigAuthorization {
	return s.Authorization
}

func (s *AgentAccessConfig) GetBasePath() *string {
	return s.BasePath
}

func (s *AgentAccessConfig) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *AgentAccessConfig) GetRemoveBasePathOnForward() *bool {
	return s.RemoveBasePathOnForward
}

func (s *AgentAccessConfig) SetAiRequestLogEnabled(v bool) *AgentAccessConfig {
	s.AiRequestLogEnabled = &v
	return s
}

func (s *AgentAccessConfig) SetAuthorization(v *AgentAccessConfigAuthorization) *AgentAccessConfig {
	s.Authorization = v
	return s
}

func (s *AgentAccessConfig) SetBasePath(v string) *AgentAccessConfig {
	s.BasePath = &v
	return s
}

func (s *AgentAccessConfig) SetDomainIds(v []*string) *AgentAccessConfig {
	s.DomainIds = v
	return s
}

func (s *AgentAccessConfig) SetRemoveBasePathOnForward(v bool) *AgentAccessConfig {
	s.RemoveBasePathOnForward = &v
	return s
}

func (s *AgentAccessConfig) Validate() error {
	if s.Authorization != nil {
		if err := s.Authorization.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentAccessConfigAuthorization struct {
	// The authentication type of the Agent access entry. Specify this parameter only when enabled is set to true.
	//
	// example:
	//
	// Apikey
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Specifies whether to enable consumer authorization. If set to true, authType must be specified and at least one principal must be provided. If set to false, no principals can be specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of consumers or consumer groups that are granted Agent access permissions. At least one principal must be specified when enabled is set to true.
	Principals []*AgentAuthorizationPrincipal `json:"principals,omitempty" xml:"principals,omitempty" type:"Repeated"`
}

func (s AgentAccessConfigAuthorization) String() string {
	return dara.Prettify(s)
}

func (s AgentAccessConfigAuthorization) GoString() string {
	return s.String()
}

func (s *AgentAccessConfigAuthorization) GetAuthType() *string {
	return s.AuthType
}

func (s *AgentAccessConfigAuthorization) GetEnabled() *bool {
	return s.Enabled
}

func (s *AgentAccessConfigAuthorization) GetPrincipals() []*AgentAuthorizationPrincipal {
	return s.Principals
}

func (s *AgentAccessConfigAuthorization) SetAuthType(v string) *AgentAccessConfigAuthorization {
	s.AuthType = &v
	return s
}

func (s *AgentAccessConfigAuthorization) SetEnabled(v bool) *AgentAccessConfigAuthorization {
	s.Enabled = &v
	return s
}

func (s *AgentAccessConfigAuthorization) SetPrincipals(v []*AgentAuthorizationPrincipal) *AgentAccessConfigAuthorization {
	s.Principals = v
	return s
}

func (s *AgentAccessConfigAuthorization) Validate() error {
	if s.Principals != nil {
		for _, item := range s.Principals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

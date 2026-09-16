// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorUpdateInput interface {
	dara.Model
	String() string
	GoString() string
	SetAuthentication(v *ConnectorAuthenticationUpdateInput) *ConnectorUpdateInput
	GetAuthentication() *ConnectorAuthenticationUpdateInput
	SetCapabilityGrants(v []map[string]interface{}) *ConnectorUpdateInput
	GetCapabilityGrants() []map[string]interface{}
	SetConfiguration(v map[string]interface{}) *ConnectorUpdateInput
	GetConfiguration() map[string]interface{}
	SetDescription(v string) *ConnectorUpdateInput
	GetDescription() *string
	SetDisplayName(v string) *ConnectorUpdateInput
	GetDisplayName() *string
	SetEnabled(v bool) *ConnectorUpdateInput
	GetEnabled() *bool
	SetPolicy(v map[string]interface{}) *ConnectorUpdateInput
	GetPolicy() map[string]interface{}
	SetRuntime(v *ConnectorRuntime) *ConnectorUpdateInput
	GetRuntime() *ConnectorRuntime
	SetTarget(v map[string]interface{}) *ConnectorUpdateInput
	GetTarget() map[string]interface{}
}

type ConnectorUpdateInput struct {
	// The authentication configuration used to replace the existing credentials.
	Authentication *ConnectorAuthenticationUpdateInput `json:"authentication,omitempty" xml:"authentication,omitempty"`
	// The list of capabilities used to replace the existing grants.
	CapabilityGrants []map[string]interface{} `json:"capabilityGrants,omitempty" xml:"capabilityGrants,omitempty" type:"Repeated"`
	// The provider configuration used to update the Connector. Only AlibabaCloudResources allows null. Other providers must provide an object.
	//
	// if can be null:
	// true
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// The description of the Connector.
	//
	// example:
	//
	// Production observability data
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the Connector.
	//
	// example:
	//
	// CMS 2.0 production workspace
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether to enable the Connector.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The execution policy used to replace the existing policy.
	Policy map[string]interface{} `json:"policy,omitempty" xml:"policy,omitempty"`
	// The runtime configuration used to update the Connector.
	Runtime *ConnectorRuntime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The provider target used to update the Connector.
	Target map[string]interface{} `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ConnectorUpdateInput) String() string {
	return dara.Prettify(s)
}

func (s ConnectorUpdateInput) GoString() string {
	return s.String()
}

func (s *ConnectorUpdateInput) GetAuthentication() *ConnectorAuthenticationUpdateInput {
	return s.Authentication
}

func (s *ConnectorUpdateInput) GetCapabilityGrants() []map[string]interface{} {
	return s.CapabilityGrants
}

func (s *ConnectorUpdateInput) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *ConnectorUpdateInput) GetDescription() *string {
	return s.Description
}

func (s *ConnectorUpdateInput) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ConnectorUpdateInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConnectorUpdateInput) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *ConnectorUpdateInput) GetRuntime() *ConnectorRuntime {
	return s.Runtime
}

func (s *ConnectorUpdateInput) GetTarget() map[string]interface{} {
	return s.Target
}

func (s *ConnectorUpdateInput) SetAuthentication(v *ConnectorAuthenticationUpdateInput) *ConnectorUpdateInput {
	s.Authentication = v
	return s
}

func (s *ConnectorUpdateInput) SetCapabilityGrants(v []map[string]interface{}) *ConnectorUpdateInput {
	s.CapabilityGrants = v
	return s
}

func (s *ConnectorUpdateInput) SetConfiguration(v map[string]interface{}) *ConnectorUpdateInput {
	s.Configuration = v
	return s
}

func (s *ConnectorUpdateInput) SetDescription(v string) *ConnectorUpdateInput {
	s.Description = &v
	return s
}

func (s *ConnectorUpdateInput) SetDisplayName(v string) *ConnectorUpdateInput {
	s.DisplayName = &v
	return s
}

func (s *ConnectorUpdateInput) SetEnabled(v bool) *ConnectorUpdateInput {
	s.Enabled = &v
	return s
}

func (s *ConnectorUpdateInput) SetPolicy(v map[string]interface{}) *ConnectorUpdateInput {
	s.Policy = v
	return s
}

func (s *ConnectorUpdateInput) SetRuntime(v *ConnectorRuntime) *ConnectorUpdateInput {
	s.Runtime = v
	return s
}

func (s *ConnectorUpdateInput) SetTarget(v map[string]interface{}) *ConnectorUpdateInput {
	s.Target = v
	return s
}

func (s *ConnectorUpdateInput) Validate() error {
	if s.Authentication != nil {
		if err := s.Authentication.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorCreateInput interface {
	dara.Model
	String() string
	GoString() string
	SetAuthentication(v *ConnectorAuthenticationInput) *ConnectorCreateInput
	GetAuthentication() *ConnectorAuthenticationInput
	SetCapabilityGrants(v []map[string]interface{}) *ConnectorCreateInput
	GetCapabilityGrants() []map[string]interface{}
	SetClientToken(v string) *ConnectorCreateInput
	GetClientToken() *string
	SetConfiguration(v map[string]interface{}) *ConnectorCreateInput
	GetConfiguration() map[string]interface{}
	SetConnectorName(v string) *ConnectorCreateInput
	GetConnectorName() *string
	SetDescription(v string) *ConnectorCreateInput
	GetDescription() *string
	SetDisplayName(v string) *ConnectorCreateInput
	GetDisplayName() *string
	SetEnabled(v bool) *ConnectorCreateInput
	GetEnabled() *bool
	SetPolicy(v map[string]interface{}) *ConnectorCreateInput
	GetPolicy() map[string]interface{}
	SetProvider(v string) *ConnectorCreateInput
	GetProvider() *string
	SetRuntime(v *ConnectorRuntime) *ConnectorCreateInput
	GetRuntime() *ConnectorRuntime
	SetTarget(v map[string]interface{}) *ConnectorCreateInput
	GetTarget() map[string]interface{}
}

type ConnectorCreateInput struct {
	// The authentication configuration used to access the target service.
	//
	// This parameter is required.
	Authentication *ConnectorAuthenticationInput `json:"authentication,omitempty" xml:"authentication,omitempty"`
	// The list of capabilities granted to the connector.
	//
	// This parameter is required.
	CapabilityGrants []map[string]interface{} `json:"capabilityGrants,omitempty" xml:"capabilityGrants,omitempty" type:"Repeated"`
	// Idempotency token
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f73d0f4-3c8a-4eed-91e6-cf2f7ebcb3d7
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Provider configuration
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// Connector name
	//
	// This parameter is required.
	//
	// example:
	//
	// cms2-prod
	ConnectorName *string `json:"connectorName,omitempty" xml:"connectorName,omitempty"`
	// Description
	//
	// example:
	//
	// Production observability data
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display name
	//
	// This parameter is required.
	//
	// example:
	//
	// CMS 2.0 production workspace
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the connector is enabled after creation.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The execution policy of the connector.
	//
	// This parameter is required.
	Policy map[string]interface{} `json:"policy,omitempty" xml:"policy,omitempty"`
	// Provider
	//
	// This parameter is required.
	//
	// example:
	//
	// AlibabaCloudCms
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The runtime configuration of the connector.
	//
	// This parameter is required.
	Runtime *ConnectorRuntime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Provider target
	//
	// This parameter is required.
	Target map[string]interface{} `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ConnectorCreateInput) String() string {
	return dara.Prettify(s)
}

func (s ConnectorCreateInput) GoString() string {
	return s.String()
}

func (s *ConnectorCreateInput) GetAuthentication() *ConnectorAuthenticationInput {
	return s.Authentication
}

func (s *ConnectorCreateInput) GetCapabilityGrants() []map[string]interface{} {
	return s.CapabilityGrants
}

func (s *ConnectorCreateInput) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConnectorCreateInput) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *ConnectorCreateInput) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *ConnectorCreateInput) GetDescription() *string {
	return s.Description
}

func (s *ConnectorCreateInput) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ConnectorCreateInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConnectorCreateInput) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *ConnectorCreateInput) GetProvider() *string {
	return s.Provider
}

func (s *ConnectorCreateInput) GetRuntime() *ConnectorRuntime {
	return s.Runtime
}

func (s *ConnectorCreateInput) GetTarget() map[string]interface{} {
	return s.Target
}

func (s *ConnectorCreateInput) SetAuthentication(v *ConnectorAuthenticationInput) *ConnectorCreateInput {
	s.Authentication = v
	return s
}

func (s *ConnectorCreateInput) SetCapabilityGrants(v []map[string]interface{}) *ConnectorCreateInput {
	s.CapabilityGrants = v
	return s
}

func (s *ConnectorCreateInput) SetClientToken(v string) *ConnectorCreateInput {
	s.ClientToken = &v
	return s
}

func (s *ConnectorCreateInput) SetConfiguration(v map[string]interface{}) *ConnectorCreateInput {
	s.Configuration = v
	return s
}

func (s *ConnectorCreateInput) SetConnectorName(v string) *ConnectorCreateInput {
	s.ConnectorName = &v
	return s
}

func (s *ConnectorCreateInput) SetDescription(v string) *ConnectorCreateInput {
	s.Description = &v
	return s
}

func (s *ConnectorCreateInput) SetDisplayName(v string) *ConnectorCreateInput {
	s.DisplayName = &v
	return s
}

func (s *ConnectorCreateInput) SetEnabled(v bool) *ConnectorCreateInput {
	s.Enabled = &v
	return s
}

func (s *ConnectorCreateInput) SetPolicy(v map[string]interface{}) *ConnectorCreateInput {
	s.Policy = v
	return s
}

func (s *ConnectorCreateInput) SetProvider(v string) *ConnectorCreateInput {
	s.Provider = &v
	return s
}

func (s *ConnectorCreateInput) SetRuntime(v *ConnectorRuntime) *ConnectorCreateInput {
	s.Runtime = v
	return s
}

func (s *ConnectorCreateInput) SetTarget(v map[string]interface{}) *ConnectorCreateInput {
	s.Target = v
	return s
}

func (s *ConnectorCreateInput) Validate() error {
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

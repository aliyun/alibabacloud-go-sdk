// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpTemplateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMcpTemplateConfigRequestBody) *UpdateMcpTemplateConfigRequest
	GetBody() *UpdateMcpTemplateConfigRequestBody
	SetClientToken(v string) *UpdateMcpTemplateConfigRequest
	GetClientToken() *string
	SetTemplateVersion(v string) *UpdateMcpTemplateConfigRequest
	GetTemplateVersion() *string
}

type UpdateMcpTemplateConfigRequest struct {
	// The MCP configuration to update by the specified template version. The configuration must conform to the input schema of the template.
	Body *UpdateMcpTemplateConfigRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template version used for this update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s UpdateMcpTemplateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequest) GetBody() *UpdateMcpTemplateConfigRequestBody {
	return s.Body
}

func (s *UpdateMcpTemplateConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpTemplateConfigRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateMcpTemplateConfigRequest) SetBody(v *UpdateMcpTemplateConfigRequestBody) *UpdateMcpTemplateConfigRequest {
	s.Body = v
	return s
}

func (s *UpdateMcpTemplateConfigRequest) SetClientToken(v string) *UpdateMcpTemplateConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequest) SetTemplateVersion(v string) *UpdateMcpTemplateConfigRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBody struct {
	// The list of remote MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The MCP authentication configuration.
	Auth *UpdateMcpTemplateConfigRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *UpdateMcpTemplateConfigRequestBodyDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description of the MCP service.
	//
	// example:
	//
	// An MCP service for querying the knowledge base
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The OpenAPI configuration used for HTTP-to-MCP conversion, represented as a JSON string.
	//
	// example:
	//
	// {"openapi":"3.0.3","info":{"title":"Knowledge API","version":"1.0.0"},"paths":{}}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateMcpTemplateConfigRequestBody) GetAuth() *UpdateMcpTemplateConfigRequestBodyAuth {
	return s.Auth
}

func (s *UpdateMcpTemplateConfigRequestBody) GetDeploymentConfig() *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	return s.DeploymentConfig
}

func (s *UpdateMcpTemplateConfigRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpTemplateConfigRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *UpdateMcpTemplateConfigRequestBody) SetAddresses(v []*string) *UpdateMcpTemplateConfigRequestBody {
	s.Addresses = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBody) SetAuth(v *UpdateMcpTemplateConfigRequestBodyAuth) *UpdateMcpTemplateConfigRequestBody {
	s.Auth = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBody) SetDeploymentConfig(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) *UpdateMcpTemplateConfigRequestBody {
	s.DeploymentConfig = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBody) SetDescription(v string) *UpdateMcpTemplateConfigRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBody) SetSwaggerConfig(v string) *UpdateMcpTemplateConfigRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBody) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	if s.DeploymentConfig != nil {
		if err := s.DeploymentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyAuth struct {
	// The backend authentication configuration for direct proxy.
	DirectProxy *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether this configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of backend authentication configurations for HTTP-to-MCP conversion.
	HttpToMcp []*UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s UpdateMcpTemplateConfigRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) GetDirectProxy() *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) GetHttpToMcp() []*UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) SetDirectProxy(v *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) *UpdateMcpTemplateConfigRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) SetHttpToMcp(v []*UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) *UpdateMcpTemplateConfigRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuth) Validate() error {
	if s.DirectProxy != nil {
		if err := s.DirectProxy.Validate(); err != nil {
			return err
		}
	}
	if s.HttpToMcp != nil {
		for _, item := range s.HttpToMcp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyAuthDirectProxy struct {
	// The name of the backend authentication request header.
	//
	// example:
	//
	// Authorization
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the backend authentication request header.
	//
	// example:
	//
	// example-api-key
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) SetName(v string) *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) SetValue(v string) *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp struct {
	// The backend authentication credential.
	//
	// example:
	//
	// example-api-key
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The ID of the backend authentication configuration.
	//
	// example:
	//
	// api-key-auth
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the API key parameter.
	//
	// example:
	//
	// X-API-Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The position where the API key is delivered.
	//
	// example:
	//
	// header
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// The backend authentication type.
	//
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) SetCredential(v string) *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) SetId(v string) *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) SetName(v string) *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) SetPosition(v string) *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) SetType(v string) *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfig struct {
	// The MCP ingress access control configuration.
	AccessControl *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// Valid values: Code (ZIP code package) and Container (custom container).
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transform and result enhancement configuration.
	ParameterTransformConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetAccessControl() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetAgentIdentityConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetCodeConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetContainerConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetHookConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetLogConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetMcpConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetNasConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetNetworkConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetOssMountConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetParameterTransformConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetProxyConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) GetRuntimeConfiguration() *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetAccessControl(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetAgentIdentityConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetArtifactType(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetCodeConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetContainerConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetHookConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetLogConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetMcpConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetNasConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetNetworkConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetOssMountConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetParameterTransformConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetProxyConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) SetRuntimeConfiguration(v *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) *UpdateMcpTemplateConfigRequestBodyDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfig) Validate() error {
	if s.AccessControl != nil {
		if err := s.AccessControl.Validate(); err != nil {
			return err
		}
	}
	if s.AgentIdentityConfiguration != nil {
		if err := s.AgentIdentityConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.CodeConfiguration != nil {
		if err := s.CodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.HookConfiguration != nil {
		if err := s.HookConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.McpConfiguration != nil {
		if err := s.McpConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NasConfiguration != nil {
		if err := s.NasConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.OssMountConfiguration != nil {
		if err := s.OssMountConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ParameterTransformConfiguration != nil {
		if err := s.ParameterTransformConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProxyConfiguration != nil {
		if err := s.ProxyConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeConfiguration != nil {
		if err := s.RuntimeConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl struct {
	// The AgentCore Credential referenced when mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Specifies whether to enable ingress access control.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access control mode. Valid values:
	//
	// - ANONYMOUS: anonymous access.
	//
	// - CREDENTIAL: uses AgentCore access credentials.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) SetCredentialId(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) SetMode(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether to enable authorization.
	AuthorizationEnabled *bool `json:"authorizationEnabled,omitempty" xml:"authorizationEnabled,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the credential provider.
	//
	// example:
	//
	// acs:agentidentity:cn-hangzhou:1234567890123456:provider/example
	CredentialProviderArn *string `json:"credentialProviderArn,omitempty" xml:"credentialProviderArn,omitempty"`
	// The credential provider type.
	//
	// example:
	//
	// oauth2
	CredentialProviderType *string `json:"credentialProviderType,omitempty" xml:"credentialProviderType,omitempty"`
	// Specifies whether to enable Agent Identity.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. Use this token to create or update a code deployment after completing the pre-signed upload.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command, with each argument passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the full subcommand, and remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) SetCommand(v []*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) SetLanguage(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration struct {
	// The ACR instance ID.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The startup command.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The container entrypoint arguments.
	Entrypoint []*string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty" type:"Repeated"`
	// The container image address.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/mcp:1.0.0
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The image registry type.
	//
	// example:
	//
	// ACR
	ImageRegistryType *string `json:"imageRegistryType,omitempty" xml:"imageRegistryType,omitempty"`
	// Custom containers must expose a standard MCP endpoint on their own. Set this parameter to SELF_HOSTED.
	//
	// example:
	//
	// SELF_HOSTED
	McpRuntimeMode *string `json:"mcpRuntimeMode,omitempty" xml:"mcpRuntimeMode,omitempty"`
	// Currently fixed to CONTAINER_IMAGE.
	//
	// example:
	//
	// CONTAINER_IMAGE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetCommand(v []*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetImage(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) SetSourceType(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration struct {
	// The PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL hooks are executed in array order.
	Hooks []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) GetHooks() []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) SetHooks(v []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfiguration) Validate() error {
	if s.Hooks != nil {
		for _, item := range s.Hooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks struct {
	// The hook API version.
	//
	// example:
	//
	// 1.0
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The hook description.
	//
	// example:
	//
	// Log MCP tool invocations
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable the hook.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The hook event.
	//
	// example:
	//
	// PRE_CALL_TOOL
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// The hook request headers.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The timeout period, in milliseconds.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The hook callback URL.
	//
	// example:
	//
	// https://example.com/mcp-hook
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetDescription(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetEvent(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) SetUrl(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration struct {
	// Specifies whether to collect instance metrics.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to collect request metrics.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log splitting begin rule for Function Compute (FC).
	//
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// The Logstore name.
	//
	// example:
	//
	// mcp-logs
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The Simple Log Service project name.
	//
	// example:
	//
	// agentcore-mcp-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) SetLogstore(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) SetProject(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration struct {
	// The MCP endpoint path, such as /mcp or /sse.
	//
	// example:
	//
	// /mcp
	EndpointPath *string `json:"endpointPath,omitempty" xml:"endpointPath,omitempty"`
	// The number of concurrent sessions per instance. Currently fixed to 1.
	//
	// example:
	//
	// 1
	SessionConcurrencyPerInstance *int32 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	// The session idle timeout period. Unit: seconds. Default value: 1800.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The maximum session lifetime. Unit: seconds. Default value: 21600.
	//
	// example:
	//
	// 21600
	SessionMaxLifetimeSeconds *int32 `json:"sessionMaxLifetimeSeconds,omitempty" xml:"sessionMaxLifetimeSeconds,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) GetMountPoints() []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) SetGroupId(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) SetMountPoints(v []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) SetUserId(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfiguration) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints struct {
	// Specifies whether to enable TLS.
	EnableTls *bool `json:"enableTls,omitempty" xml:"enableTls,omitempty"`
	// The local mount directory.
	//
	// example:
	//
	// /mnt/data
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// The NAS server address.
	//
	// example:
	//
	// example.cn-hangzhou.nas.aliyuncs.com
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration struct {
	// The network mode.
	//
	// example:
	//
	// PUBLIC
	NetworkMode *string `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-example
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) SetVpcId(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) GetMountPoints() []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) SetMountPoints(v []*UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfiguration) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints struct {
	// The OSS bucket name.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The OSS bucket path.
	//
	// example:
	//
	// /data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The OSS service endpoint.
	//
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The local mount directory.
	//
	// example:
	//
	// /mnt/data
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// Specifies whether the mount point is read-only.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration struct {
	// Specifies whether to enable parameter transform and result enhancement.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The reserved reference to the parameter transform and result enhancement rule set.
	//
	// example:
	//
	// rules-1
	RuleSetId *string `json:"ruleSetId,omitempty" xml:"ruleSetId,omitempty"`
	// The transform rule version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration struct {
	// The number of vCPUs. Default value: 0.25.
	//
	// example:
	//
	// 0.25
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The ephemeral disk size. Unit: MB. Valid values: 512 and 10240.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// The environment variables.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// The ARN of the RAM role used by user code to access downstream Alibaba Cloud resources.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-mcp-execution
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The maximum number of concurrent requests per instance. Default value: 200.
	//
	// example:
	//
	// 200
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// The memory size. Unit: MB. Default value: 512.
	//
	// example:
	//
	// 512
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// The service port. Default value: 9000.
	//
	// example:
	//
	// 9000
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The function timeout period. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetPort(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpTemplateConfigRequestBodyDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMcpMarketItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *InstallMcpMarketItemRequestBody) *InstallMcpMarketItemRequest
	GetBody() *InstallMcpMarketItemRequestBody
	SetClientToken(v string) *InstallMcpMarketItemRequest
	GetClientToken() *string
	SetTemplateVersion(v string) *InstallMcpMarketItemRequest
	GetTemplateVersion() *string
}

type InstallMcpMarketItemRequest struct {
	// The MCP configuration submitted during template installation. The configuration must conform to the input schema of the template.
	Body *InstallMcpMarketItemRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template version to install. You can call GetMcpMarketItem to query available versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s InstallMcpMarketItemRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequest) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequest) GetBody() *InstallMcpMarketItemRequestBody {
	return s.Body
}

func (s *InstallMcpMarketItemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallMcpMarketItemRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *InstallMcpMarketItemRequest) SetBody(v *InstallMcpMarketItemRequestBody) *InstallMcpMarketItemRequest {
	s.Body = v
	return s
}

func (s *InstallMcpMarketItemRequest) SetClientToken(v string) *InstallMcpMarketItemRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallMcpMarketItemRequest) SetTemplateVersion(v string) *InstallMcpMarketItemRequest {
	s.TemplateVersion = &v
	return s
}

func (s *InstallMcpMarketItemRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallMcpMarketItemRequestBody struct {
	// The list of remote MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The MCP authentication configuration.
	Auth *InstallMcpMarketItemRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *InstallMcpMarketItemRequestBodyDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The MCP service description.
	//
	// example:
	//
	// An MCP service for querying the knowledge base
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP service name.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The OpenAPI configuration for HTTP-to-MCP conversion, represented as a JSON string.
	//
	// example:
	//
	// {"openapi":"3.0.3","info":{"title":"Knowledge API","version":"1.0.0"},"paths":{}}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// The MCP type. Valid values:
	//
	// - DIRECT_PROXY: direct proxy.
	//
	// - HTTP_TO_MCP: HTTP-to-MCP conversion.
	//
	// - CODE_PACKAGE: code deployment.
	//
	// example:
	//
	// CODE_PACKAGE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InstallMcpMarketItemRequestBody) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBody) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *InstallMcpMarketItemRequestBody) GetAuth() *InstallMcpMarketItemRequestBodyAuth {
	return s.Auth
}

func (s *InstallMcpMarketItemRequestBody) GetDeploymentConfig() *InstallMcpMarketItemRequestBodyDeploymentConfig {
	return s.DeploymentConfig
}

func (s *InstallMcpMarketItemRequestBody) GetDescription() *string {
	return s.Description
}

func (s *InstallMcpMarketItemRequestBody) GetName() *string {
	return s.Name
}

func (s *InstallMcpMarketItemRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *InstallMcpMarketItemRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *InstallMcpMarketItemRequestBody) GetType() *string {
	return s.Type
}

func (s *InstallMcpMarketItemRequestBody) SetAddresses(v []*string) *InstallMcpMarketItemRequestBody {
	s.Addresses = v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetAuth(v *InstallMcpMarketItemRequestBodyAuth) *InstallMcpMarketItemRequestBody {
	s.Auth = v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetDeploymentConfig(v *InstallMcpMarketItemRequestBodyDeploymentConfig) *InstallMcpMarketItemRequestBody {
	s.DeploymentConfig = v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetDescription(v string) *InstallMcpMarketItemRequestBody {
	s.Description = &v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetName(v string) *InstallMcpMarketItemRequestBody {
	s.Name = &v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetProtocol(v string) *InstallMcpMarketItemRequestBody {
	s.Protocol = &v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetSwaggerConfig(v string) *InstallMcpMarketItemRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *InstallMcpMarketItemRequestBody) SetType(v string) *InstallMcpMarketItemRequestBody {
	s.Type = &v
	return s
}

func (s *InstallMcpMarketItemRequestBody) Validate() error {
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

type InstallMcpMarketItemRequestBodyAuth struct {
	// The backend authentication configuration for direct proxy.
	DirectProxy *InstallMcpMarketItemRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether the configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of backend authentication configurations for HTTP-to-MCP conversion.
	HttpToMcp []*InstallMcpMarketItemRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s InstallMcpMarketItemRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyAuth) GetDirectProxy() *InstallMcpMarketItemRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *InstallMcpMarketItemRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyAuth) GetHttpToMcp() []*InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *InstallMcpMarketItemRequestBodyAuth) SetDirectProxy(v *InstallMcpMarketItemRequestBodyAuthDirectProxy) *InstallMcpMarketItemRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuth) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuth) SetHttpToMcp(v []*InstallMcpMarketItemRequestBodyAuthHttpToMcp) *InstallMcpMarketItemRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuth) Validate() error {
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

type InstallMcpMarketItemRequestBodyAuthDirectProxy struct {
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

func (s InstallMcpMarketItemRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *InstallMcpMarketItemRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *InstallMcpMarketItemRequestBodyAuthDirectProxy) SetName(v string) *InstallMcpMarketItemRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthDirectProxy) SetValue(v string) *InstallMcpMarketItemRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyAuthHttpToMcp struct {
	// The backend authentication credential.
	//
	// example:
	//
	// example-api-key
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The backend authentication configuration ID.
	//
	// example:
	//
	// api-key-auth
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The API key parameter name.
	//
	// example:
	//
	// X-API-Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The position where the API key is passed.
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

func (s InstallMcpMarketItemRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) SetCredential(v string) *InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) SetId(v string) *InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) SetName(v string) *InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) SetPosition(v string) *InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) SetType(v string) *InstallMcpMarketItemRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfig struct {
	// The MCP ingress access control configuration.
	AccessControl *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// Code indicates a ZIP code package. Container indicates a custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfig) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetAccessControl() *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetAgentIdentityConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetCodeConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetContainerConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetHookConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetLogConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetMcpConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetNasConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetNetworkConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetOssMountConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetParameterTransformConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetProxyConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) GetRuntimeConfiguration() *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetAccessControl(v *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetAgentIdentityConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetArtifactType(v string) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetCodeConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetContainerConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetHookConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetLogConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetMcpConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetNasConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetNetworkConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetOssMountConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetParameterTransformConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetProxyConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) SetRuntimeConfiguration(v *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) *InstallMcpMarketItemRequestBodyDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfig) Validate() error {
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

type InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl struct {
	// The AgentCore Credential referenced when mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Specifies whether to enable ingress access control.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// ANONYMOUS indicates anonymous access. CREDENTIAL indicates access using an AgentCore credential.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) SetCredentialId(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) SetMode(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether authorization is enabled.
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
	// Specifies whether Agent Identity is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. Used to create a code deployment after the pre-signed upload is complete. Specify either this parameter or CodePackageUrl.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The public Alibaba Cloud OSS HTTP(S) URL that can be directly passed in when creating a code deployment. Specify either this parameter or CodePackageToken. Only supported by CreateMcp. Not supported for update or query operations.
	//
	// example:
	//
	// https://example-bucket.oss-cn-hangzhou.aliyuncs.com/server.zip
	CodePackageUrl *string `json:"codePackageUrl,omitempty" xml:"codePackageUrl,omitempty"`
	// The full startup command, with arguments passed in sequence by parameter boundary. For example, when using supergateway to start a stdio MCP, pass in supergateway, --stdio, the full subcommand, and remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime: python3.13, nodejs22, or java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageUrl() *string {
	return s.CodePackageUrl
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageUrl(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageUrl = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) SetCommand(v []*string) *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) SetLanguage(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration struct {
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
	// The container image URL.
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
	// Custom containers must expose a standard MCP endpoint. Set this parameter to SELF_HOSTED.
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetCommand(v []*string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetImage(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) SetSourceType(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration struct {
	// Executes PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL hooks in array order.
	Hooks []*InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) GetHooks() []*InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) SetHooks(v []*InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfiguration) Validate() error {
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

type InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks struct {
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetDescription(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetEvent(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) SetUrl(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration struct {
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
	// The Log Service project name.
	//
	// example:
	//
	// agentcore-mcp-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) SetLogstore(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) SetProject(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration struct {
	// The MCP endpoint path. For example, /mcp or /sse.
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) GetMountPoints() []*InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) SetGroupId(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) SetMountPoints(v []*InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) SetUserId(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfiguration) Validate() error {
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

type InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints struct {
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration struct {
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) SetVpcId(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) GetMountPoints() []*InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) SetMountPoints(v []*InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfiguration) Validate() error {
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

type InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints struct {
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration struct {
	// Specifies whether to enable parameter transformation and result enhancement.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The reserved reference to a parameter transformation and result enhancement rule set.
	//
	// example:
	//
	// rules-1
	RuleSetId *string `json:"ruleSetId,omitempty" xml:"ruleSetId,omitempty"`
	// The transformation rule version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) SetEnabled(v bool) *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration struct {
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

func (s InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetPort(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *InstallMcpMarketItemRequestBodyDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

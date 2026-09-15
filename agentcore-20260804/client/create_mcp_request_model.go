// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateMcpRequestBody) *CreateMcpRequest
	GetBody() *CreateMcpRequestBody
	SetClientToken(v string) *CreateMcpRequest
	GetClientToken() *string
}

type CreateMcpRequest struct {
	// The request body.
	Body *CreateMcpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token that ensures idempotency of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpRequest) GetBody() *CreateMcpRequestBody {
	return s.Body
}

func (s *CreateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMcpRequest) SetBody(v *CreateMcpRequestBody) *CreateMcpRequest {
	s.Body = v
	return s
}

func (s *CreateMcpRequest) SetClientToken(v string) *CreateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMcpRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpRequestBody struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. When enabled is set to true: for DIRECT_PROXY, specify directProxy (name/value). For HTTP_TO_MCP, specify the httpToMcp array (each item contains id/type/credential. For apiKey, position/name are also required). Multiple authentication objects are supported, and the first one is used as the default upstream credential. HTTP_TO_MCP credentials are merged into the securitySchemes of the Swagger specification.
	Auth *CreateMcpRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// Custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The code deployment configuration. Required when Type is set to CODE_PACKAGE. When creating a Code artifact, you must specify either CodeConfiguration.CodePackageToken or CodePackageUrl, but not both. CodePackageUrl supports only public Alibaba Cloud OSS HTTP(S) addresses.
	DeploymentConfig *CreateMcpRequestBodyDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// SSE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Required if Type is set to HTTP_TO_MCP.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// The type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECT_PROXY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBody) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpRequestBody) GetAuth() *CreateMcpRequestBodyAuth {
	return s.Auth
}

func (s *CreateMcpRequestBody) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *CreateMcpRequestBody) GetDeploymentConfig() *CreateMcpRequestBodyDeploymentConfig {
	return s.DeploymentConfig
}

func (s *CreateMcpRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpRequestBody) GetType() *string {
	return s.Type
}

func (s *CreateMcpRequestBody) SetAddresses(v []*string) *CreateMcpRequestBody {
	s.Addresses = v
	return s
}

func (s *CreateMcpRequestBody) SetAuth(v *CreateMcpRequestBodyAuth) *CreateMcpRequestBody {
	s.Auth = v
	return s
}

func (s *CreateMcpRequestBody) SetCustomTags(v []*string) *CreateMcpRequestBody {
	s.CustomTags = v
	return s
}

func (s *CreateMcpRequestBody) SetDeploymentConfig(v *CreateMcpRequestBodyDeploymentConfig) *CreateMcpRequestBody {
	s.DeploymentConfig = v
	return s
}

func (s *CreateMcpRequestBody) SetDescription(v string) *CreateMcpRequestBody {
	s.Description = &v
	return s
}

func (s *CreateMcpRequestBody) SetName(v string) *CreateMcpRequestBody {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBody) SetProtocol(v string) *CreateMcpRequestBody {
	s.Protocol = &v
	return s
}

func (s *CreateMcpRequestBody) SetSwaggerConfig(v string) *CreateMcpRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpRequestBody) SetType(v string) *CreateMcpRequestBody {
	s.Type = &v
	return s
}

func (s *CreateMcpRequestBody) Validate() error {
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

type CreateMcpRequestBodyAuth struct {
	// The API key authentication configuration for callers of code-deployed MCP.
	CodePackage *CreateMcpRequestBodyAuthCodePackage `json:"codePackage,omitempty" xml:"codePackage,omitempty" type:"Struct"`
	// The authentication configuration for direct proxy.
	DirectProxy *CreateMcpRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable this configuration.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*CreateMcpRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s CreateMcpRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuth) GetCodePackage() *CreateMcpRequestBodyAuthCodePackage {
	return s.CodePackage
}

func (s *CreateMcpRequestBodyAuth) GetDirectProxy() *CreateMcpRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *CreateMcpRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyAuth) GetHttpToMcp() []*CreateMcpRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *CreateMcpRequestBodyAuth) SetCodePackage(v *CreateMcpRequestBodyAuthCodePackage) *CreateMcpRequestBodyAuth {
	s.CodePackage = v
	return s
}

func (s *CreateMcpRequestBodyAuth) SetDirectProxy(v *CreateMcpRequestBodyAuthDirectProxy) *CreateMcpRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *CreateMcpRequestBodyAuth) SetEnabled(v bool) *CreateMcpRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyAuth) SetHttpToMcp(v []*CreateMcpRequestBodyAuthHttpToMcp) *CreateMcpRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *CreateMcpRequestBodyAuth) Validate() error {
	if s.CodePackage != nil {
		if err := s.CodePackage.Validate(); err != nil {
			return err
		}
	}
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

type CreateMcpRequestBodyAuthCodePackage struct {
	// The API key for authenticating MCP callers.
	//
	// example:
	//
	// example-api-key
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The name of the request header that carries the API key.
	//
	// example:
	//
	// X-API-Key
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
}

func (s CreateMcpRequestBodyAuthCodePackage) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuthCodePackage) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuthCodePackage) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateMcpRequestBodyAuthCodePackage) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateMcpRequestBodyAuthCodePackage) SetApiKey(v string) *CreateMcpRequestBodyAuthCodePackage {
	s.ApiKey = &v
	return s
}

func (s *CreateMcpRequestBodyAuthCodePackage) SetHeaderName(v string) *CreateMcpRequestBodyAuthCodePackage {
	s.HeaderName = &v
	return s
}

func (s *CreateMcpRequestBodyAuthCodePackage) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyAuthDirectProxy struct {
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The authentication parameter value.
	//
	// example:
	//
	// example-credential
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMcpRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *CreateMcpRequestBodyAuthDirectProxy) SetName(v string) *CreateMcpRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBodyAuthDirectProxy) SetValue(v string) *CreateMcpRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *CreateMcpRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyAuthHttpToMcp struct {
	// The authentication credential.
	//
	// example:
	//
	// example-credential
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The authentication scheme ID.
	//
	// example:
	//
	// mcp-1234567890abcdef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The position of the credential.
	//
	// example:
	//
	// header
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// The type.
	//
	// example:
	//
	// basic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetCredential(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetId(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetName(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetPosition(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetType(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfig struct {
	// The MCP ingress access control configuration.
	AccessControl *CreateMcpRequestBodyDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// The artifact type. Valid values:
	//
	// - Code: a ZIP code package.
	//
	// - Container: a custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *CreateMcpRequestBodyDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *CreateMcpRequestBodyDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *CreateMcpRequestBodyDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *CreateMcpRequestBodyDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *CreateMcpRequestBodyDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *CreateMcpRequestBodyDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *CreateMcpRequestBodyDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s CreateMcpRequestBodyDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetAccessControl() *CreateMcpRequestBodyDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetAgentIdentityConfiguration() *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetCodeConfiguration() *CreateMcpRequestBodyDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetContainerConfiguration() *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetHookConfiguration() *CreateMcpRequestBodyDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetLogConfiguration() *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetMcpConfiguration() *CreateMcpRequestBodyDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetNasConfiguration() *CreateMcpRequestBodyDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetNetworkConfiguration() *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetOssMountConfiguration() *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetParameterTransformConfiguration() *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetProxyConfiguration() *CreateMcpRequestBodyDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) GetRuntimeConfiguration() *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetAccessControl(v *CreateMcpRequestBodyDeploymentConfigAccessControl) *CreateMcpRequestBodyDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetAgentIdentityConfiguration(v *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetArtifactType(v string) *CreateMcpRequestBodyDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetCodeConfiguration(v *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetContainerConfiguration(v *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetHookConfiguration(v *CreateMcpRequestBodyDeploymentConfigHookConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetLogConfiguration(v *CreateMcpRequestBodyDeploymentConfigLogConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetMcpConfiguration(v *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetNasConfiguration(v *CreateMcpRequestBodyDeploymentConfigNasConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetNetworkConfiguration(v *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetOssMountConfiguration(v *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetParameterTransformConfiguration(v *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetProxyConfiguration(v *CreateMcpRequestBodyDeploymentConfigProxyConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) SetRuntimeConfiguration(v *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) *CreateMcpRequestBodyDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfig) Validate() error {
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

type CreateMcpRequestBodyDeploymentConfigAccessControl struct {
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
	// - CREDENTIAL: access using an AgentCore credential.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) SetCredentialId(v string) *CreateMcpRequestBodyDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) SetEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) SetMode(v string) *CreateMcpRequestBodyDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether to enable authorization.
	AuthorizationEnabled *bool `json:"authorizationEnabled,omitempty" xml:"authorizationEnabled,omitempty"`
	// The ARN of the credential provider.
	//
	// example:
	//
	// acs:agentidentity:cn-hangzhou:1234567890123456:provider/example
	CredentialProviderArn *string `json:"credentialProviderArn,omitempty" xml:"credentialProviderArn,omitempty"`
	// The type of the credential provider.
	//
	// example:
	//
	// oauth2
	CredentialProviderType *string `json:"credentialProviderType,omitempty" xml:"credentialProviderType,omitempty"`
	// Specifies whether to enable Agent Identity.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. Use this token to create a code deployment after completing the pre-signed upload. Specify either this parameter or CodePackageUrl.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The public Alibaba Cloud OSS HTTP(S) address that you can directly pass in when creating a code deployment. Specify either this parameter or CodePackageToken. Only CreateMcp supports this parameter. Update and query operations do not support this parameter.
	//
	// example:
	//
	// https://example-bucket.oss-cn-hangzhou.aliyuncs.com/server.zip
	CodePackageUrl *string `json:"codePackageUrl,omitempty" xml:"codePackageUrl,omitempty"`
	// The full startup command, with arguments passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass in supergateway, --stdio, the full subcommand, and remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageUrl() *string {
	return s.CodePackageUrl
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *CreateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageUrl(v string) *CreateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageUrl = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) SetCommand(v []*string) *CreateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) SetLanguage(v string) *CreateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigContainerConfiguration struct {
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
	// The MCP Runtime mode. Custom containers must expose a standard MCP on their own. Set this parameter to SELF_HOSTED.
	//
	// example:
	//
	// SELF_HOSTED
	McpRuntimeMode *string `json:"mcpRuntimeMode,omitempty" xml:"mcpRuntimeMode,omitempty"`
	// The container source type. Currently fixed to CONTAINER_IMAGE.
	//
	// example:
	//
	// CONTAINER_IMAGE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetCommand(v []*string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetImage(v string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) SetSourceType(v string) *CreateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigHookConfiguration struct {
	// The hooks executed in array order: PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL.
	Hooks []*CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s CreateMcpRequestBodyDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfiguration) GetHooks() []*CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfiguration) SetHooks(v []*CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) *CreateMcpRequestBodyDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfiguration) Validate() error {
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

type CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks struct {
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
	// The hook timeout period. Unit: milliseconds.
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

func (s CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetDescription(v string) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetEvent(v string) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetUrl(v string) *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigLogConfiguration struct {
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

func (s CreateMcpRequestBodyDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) SetLogstore(v string) *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) SetProject(v string) *CreateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigMcpConfiguration struct {
	// For example, /mcp or /sse.
	//
	// example:
	//
	// /mcp
	EndpointPath *string `json:"endpointPath,omitempty" xml:"endpointPath,omitempty"`
	// Currently fixed to 1.
	//
	// example:
	//
	// 1
	SessionConcurrencyPerInstance *int32 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	// Unit: seconds. Default value: 1800.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// Unit: seconds. Default value: 21600.
	//
	// example:
	//
	// 21600
	SessionMaxLifetimeSeconds *int32 `json:"sessionMaxLifetimeSeconds,omitempty" xml:"sessionMaxLifetimeSeconds,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *CreateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *CreateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *CreateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *CreateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) GetMountPoints() []*CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) SetGroupId(v int32) *CreateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) SetMountPoints(v []*CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) *CreateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) SetUserId(v int32) *CreateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfiguration) Validate() error {
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

type CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints struct {
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

func (s CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigNetworkConfiguration struct {
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
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetVpcId(v string) *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) GetMountPoints() []*CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) SetMountPoints(v []*CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfiguration) Validate() error {
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

type CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints struct {
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

func (s CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration struct {
	// Specifies whether to enable parameter transformation and result enhancement.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The reserved reference to the parameter transformation and result enhancement rule set.
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

func (s CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyDeploymentConfigProxyConfiguration) SetEnabled(v bool) *CreateMcpRequestBodyDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration struct {
	// Unit: cores. Default value: 0.25.
	//
	// example:
	//
	// 0.25
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// Unit: MB. Valid values: 512 and 10240.
	//
	// example:
	//
	// 512
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// The environment variables.
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// The ARN of the RAM role used when user code accesses downstream Alibaba Cloud resources.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/agentcore-mcp-execution
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// Default value: 200.
	//
	// example:
	//
	// 200
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// Unit: MB. Default value: 512.
	//
	// example:
	//
	// 512
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// Default value: 9000.
	//
	// example:
	//
	// 9000
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetPort(v int32) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *CreateMcpRequestBodyDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

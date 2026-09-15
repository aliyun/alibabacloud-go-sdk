// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMcpRequestBody) *UpdateMcpRequest
	GetBody() *UpdateMcpRequestBody
	SetClientToken(v string) *UpdateMcpRequest
	GetClientToken() *string
}

type UpdateMcpRequest struct {
	// The request body.
	Body *UpdateMcpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotency of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequest) GetBody() *UpdateMcpRequestBody {
	return s.Body
}

func (s *UpdateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpRequest) SetBody(v *UpdateMcpRequestBody) *UpdateMcpRequest {
	s.Body = v
	return s
}

func (s *UpdateMcpRequest) SetClientToken(v string) *UpdateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpRequestBody struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. When enabled is set to true: for DIRECT_PROXY, specify directProxy (name/value). For HTTP_TO_MCP, specify the httpToMcp array (each item contains id/type/credential, and apiKey also requires position/name). Multiple authentication objects are supported, and the first one is used as the default upstream credential. HTTP_TO_MCP credentials are merged into the securitySchemes of the Swagger specification.
	Auth *UpdateMcpRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// Custom tags. Multiple tags are supported. Pass an empty list to clear all tags.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// Deployment configuration patch for a CODE_PACKAGE MCP. Object fields are merged hierarchically. To keep the current deployed code package, omit CodeConfiguration.CodePackageToken so the server reuses the existing package. To replace the code package, pass a new non-empty Token. CodePackageUrl is supported only for CreateMcp and is not supported during updates.
	DeploymentConfig *UpdateMcpRequestBodyDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Required if Type is set to HTTP_TO_MCP.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
}

func (s UpdateMcpRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateMcpRequestBody) GetAuth() *UpdateMcpRequestBodyAuth {
	return s.Auth
}

func (s *UpdateMcpRequestBody) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *UpdateMcpRequestBody) GetDeploymentConfig() *UpdateMcpRequestBodyDeploymentConfig {
	return s.DeploymentConfig
}

func (s *UpdateMcpRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *UpdateMcpRequestBody) SetAddresses(v []*string) *UpdateMcpRequestBody {
	s.Addresses = v
	return s
}

func (s *UpdateMcpRequestBody) SetAuth(v *UpdateMcpRequestBodyAuth) *UpdateMcpRequestBody {
	s.Auth = v
	return s
}

func (s *UpdateMcpRequestBody) SetCustomTags(v []*string) *UpdateMcpRequestBody {
	s.CustomTags = v
	return s
}

func (s *UpdateMcpRequestBody) SetDeploymentConfig(v *UpdateMcpRequestBodyDeploymentConfig) *UpdateMcpRequestBody {
	s.DeploymentConfig = v
	return s
}

func (s *UpdateMcpRequestBody) SetDescription(v string) *UpdateMcpRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateMcpRequestBody) SetSwaggerConfig(v string) *UpdateMcpRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *UpdateMcpRequestBody) Validate() error {
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

type UpdateMcpRequestBodyAuth struct {
	// The API key authentication configuration for code-deployed MCP callers.
	CodePackage *UpdateMcpRequestBodyAuthCodePackage `json:"codePackage,omitempty" xml:"codePackage,omitempty" type:"Struct"`
	// The direct proxy authentication configuration.
	DirectProxy *UpdateMcpRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable the configuration.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*UpdateMcpRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s UpdateMcpRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuth) GetCodePackage() *UpdateMcpRequestBodyAuthCodePackage {
	return s.CodePackage
}

func (s *UpdateMcpRequestBodyAuth) GetDirectProxy() *UpdateMcpRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *UpdateMcpRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyAuth) GetHttpToMcp() []*UpdateMcpRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *UpdateMcpRequestBodyAuth) SetCodePackage(v *UpdateMcpRequestBodyAuthCodePackage) *UpdateMcpRequestBodyAuth {
	s.CodePackage = v
	return s
}

func (s *UpdateMcpRequestBodyAuth) SetDirectProxy(v *UpdateMcpRequestBodyAuthDirectProxy) *UpdateMcpRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *UpdateMcpRequestBodyAuth) SetEnabled(v bool) *UpdateMcpRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyAuth) SetHttpToMcp(v []*UpdateMcpRequestBodyAuthHttpToMcp) *UpdateMcpRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *UpdateMcpRequestBodyAuth) Validate() error {
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

type UpdateMcpRequestBodyAuthCodePackage struct {
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

func (s UpdateMcpRequestBodyAuthCodePackage) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuthCodePackage) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuthCodePackage) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateMcpRequestBodyAuthCodePackage) GetHeaderName() *string {
	return s.HeaderName
}

func (s *UpdateMcpRequestBodyAuthCodePackage) SetApiKey(v string) *UpdateMcpRequestBodyAuthCodePackage {
	s.ApiKey = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthCodePackage) SetHeaderName(v string) *UpdateMcpRequestBodyAuthCodePackage {
	s.HeaderName = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthCodePackage) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyAuthDirectProxy struct {
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

func (s UpdateMcpRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) SetName(v string) *UpdateMcpRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) SetValue(v string) *UpdateMcpRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyAuthHttpToMcp struct {
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

func (s UpdateMcpRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetCredential(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetId(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetName(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetPosition(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetType(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfig struct {
	// The MCP ingress access control configuration.
	AccessControl *UpdateMcpRequestBodyDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
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
	CodeConfiguration *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *UpdateMcpRequestBodyDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *UpdateMcpRequestBodyDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *UpdateMcpRequestBodyDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s UpdateMcpRequestBodyDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfig) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetAccessControl() *UpdateMcpRequestBodyDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetAgentIdentityConfiguration() *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetCodeConfiguration() *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetContainerConfiguration() *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetHookConfiguration() *UpdateMcpRequestBodyDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetLogConfiguration() *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetMcpConfiguration() *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetNasConfiguration() *UpdateMcpRequestBodyDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetNetworkConfiguration() *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetOssMountConfiguration() *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetParameterTransformConfiguration() *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetProxyConfiguration() *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) GetRuntimeConfiguration() *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetAccessControl(v *UpdateMcpRequestBodyDeploymentConfigAccessControl) *UpdateMcpRequestBodyDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetAgentIdentityConfiguration(v *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetArtifactType(v string) *UpdateMcpRequestBodyDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetCodeConfiguration(v *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetContainerConfiguration(v *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetHookConfiguration(v *UpdateMcpRequestBodyDeploymentConfigHookConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetLogConfiguration(v *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetMcpConfiguration(v *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetNasConfiguration(v *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetNetworkConfiguration(v *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetOssMountConfiguration(v *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetParameterTransformConfiguration(v *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetProxyConfiguration(v *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) SetRuntimeConfiguration(v *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) *UpdateMcpRequestBodyDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfig) Validate() error {
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

type UpdateMcpRequestBodyDeploymentConfigAccessControl struct {
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
	// - CREDENTIAL: access with an AgentCore credential.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) SetCredentialId(v string) *UpdateMcpRequestBodyDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) SetEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) SetMode(v string) *UpdateMcpRequestBodyDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether to enable authorization.
	AuthorizationEnabled *bool `json:"authorizationEnabled,omitempty" xml:"authorizationEnabled,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the credential provider.
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

func (s UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. After the pre-signed upload is complete, use this token to create or update a code deployment.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command. Pass each argument as a separate element in order. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the full subcommand, and the remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) SetCommand(v []*string) *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) SetLanguage(v string) *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigContainerConfiguration struct {
	// The ID of the Alibaba Cloud Container Registry (ACR) instance.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The startup command.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The container entrypoint arguments.
	Entrypoint []*string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty" type:"Repeated"`
	// The URL of the container image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/mcp:1.0.0
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The type of the image registry.
	//
	// example:
	//
	// ACR
	ImageRegistryType *string `json:"imageRegistryType,omitempty" xml:"imageRegistryType,omitempty"`
	// The MCP runtime mode for the custom container. The custom container must expose a standard MCP endpoint. Set this parameter to SELF_HOSTED.
	//
	// example:
	//
	// SELF_HOSTED
	McpRuntimeMode *string `json:"mcpRuntimeMode,omitempty" xml:"mcpRuntimeMode,omitempty"`
	// The source type of the container. Currently fixed to CONTAINER_IMAGE.
	//
	// example:
	//
	// CONTAINER_IMAGE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetCommand(v []*string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetImage(v string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) SetSourceType(v string) *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigHookConfiguration struct {
	// The list of hooks executed in array order. Supported hook events: PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL.
	Hooks []*UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s UpdateMcpRequestBodyDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfiguration) GetHooks() []*UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfiguration) SetHooks(v []*UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) *UpdateMcpRequestBodyDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfiguration) Validate() error {
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

type UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks struct {
	// The API version of the hook.
	//
	// example:
	//
	// 1.0
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The description of the hook.
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
	// The HTTP request headers for the hook.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The timeout period, in milliseconds.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The callback URL of the hook.
	//
	// example:
	//
	// https://example.com/mcp-hook
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetDescription(v string) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetEvent(v string) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) SetUrl(v string) *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigLogConfiguration struct {
	// Specifies whether to collect instance metrics.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to collect request metrics.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log segmentation start rule for Function Compute.
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

func (s UpdateMcpRequestBodyDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) SetLogstore(v string) *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) SetProject(v string) *UpdateMcpRequestBodyDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigMcpConfiguration struct {
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
	// The session idle timeout. Unit: seconds. Default value: 1800.
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

func (s UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) GetMountPoints() []*UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) SetGroupId(v int32) *UpdateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) SetMountPoints(v []*UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) *UpdateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) SetUserId(v int32) *UpdateMcpRequestBodyDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfiguration) Validate() error {
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

type UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints struct {
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

func (s UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration struct {
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

func (s UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) SetVpcId(v string) *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) GetMountPoints() []*UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) SetMountPoints(v []*UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfiguration) Validate() error {
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

type UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints struct {
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

func (s UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration struct {
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

func (s UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) SetEnabled(v bool) *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration struct {
	// The CPU specification. Unit: cores. Default value: 0.25.
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

func (s UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetPort(v int32) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpRequestBodyDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMcpMarketItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallMcpMarketItemResponseBody
	GetCode() *string
	SetData(v *InstallMcpMarketItemResponseBodyData) *InstallMcpMarketItemResponseBody
	GetData() *InstallMcpMarketItemResponseBodyData
	SetHttpStatusCode(v int32) *InstallMcpMarketItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InstallMcpMarketItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallMcpMarketItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallMcpMarketItemResponseBody
	GetSuccess() *bool
}

type InstallMcpMarketItemResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *InstallMcpMarketItemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, used for locating and troubleshooting issues.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InstallMcpMarketItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBody) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallMcpMarketItemResponseBody) GetData() *InstallMcpMarketItemResponseBodyData {
	return s.Data
}

func (s *InstallMcpMarketItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InstallMcpMarketItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallMcpMarketItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallMcpMarketItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallMcpMarketItemResponseBody) SetCode(v string) *InstallMcpMarketItemResponseBody {
	s.Code = &v
	return s
}

func (s *InstallMcpMarketItemResponseBody) SetData(v *InstallMcpMarketItemResponseBodyData) *InstallMcpMarketItemResponseBody {
	s.Data = v
	return s
}

func (s *InstallMcpMarketItemResponseBody) SetHttpStatusCode(v int32) *InstallMcpMarketItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InstallMcpMarketItemResponseBody) SetMessage(v string) *InstallMcpMarketItemResponseBody {
	s.Message = &v
	return s
}

func (s *InstallMcpMarketItemResponseBody) SetRequestId(v string) *InstallMcpMarketItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBody) SetSuccess(v bool) *InstallMcpMarketItemResponseBody {
	s.Success = &v
	return s
}

func (s *InstallMcpMarketItemResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallMcpMarketItemResponseBodyData struct {
	// The list of remote MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *InstallMcpMarketItemResponseBodyDataDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The MCP service description.
	//
	// example:
	//
	// MCP service for querying knowledge bases
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP service access endpoint.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Function Compute function name corresponding to the code-deployed MCP.
	//
	// example:
	//
	// agentcore-mcp-example
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The marketplace template from which the MCP originates.
	MarketSource *InstallMcpMarketItemResponseBodyDataMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
	// The MCP service ID.
	//
	// example:
	//
	// mcp-server-id
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The MCP service name.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The official purpose tag, managed by the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The MCP service status.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason why the MCP service is in the current status.
	//
	// example:
	//
	// Code package deployment failed
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The template version and input schema bound to the MCP.
	Template *InstallMcpMarketItemResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The MCP type. Valid values:
	//
	// - DIRECT_PROXY: Direct proxy.
	//
	// - HTTP_TO_MCP: HTTP-to-MCP conversion.
	//
	// - CODE_PACKAGE: Code deployment.
	//
	// example:
	//
	// CODE_PACKAGE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Indicates whether the MCP is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *InstallMcpMarketItemResponseBodyData) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *InstallMcpMarketItemResponseBodyData) GetDeploymentConfig() *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	return s.DeploymentConfig
}

func (s *InstallMcpMarketItemResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *InstallMcpMarketItemResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InstallMcpMarketItemResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *InstallMcpMarketItemResponseBodyData) GetMarketSource() *InstallMcpMarketItemResponseBodyDataMarketSource {
	return s.MarketSource
}

func (s *InstallMcpMarketItemResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *InstallMcpMarketItemResponseBodyData) GetName() *string {
	return s.Name
}

func (s *InstallMcpMarketItemResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *InstallMcpMarketItemResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *InstallMcpMarketItemResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *InstallMcpMarketItemResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *InstallMcpMarketItemResponseBodyData) GetTemplate() *InstallMcpMarketItemResponseBodyDataTemplate {
	return s.Template
}

func (s *InstallMcpMarketItemResponseBodyData) GetType() *string {
	return s.Type
}

func (s *InstallMcpMarketItemResponseBodyData) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *InstallMcpMarketItemResponseBodyData) SetAddresses(v []*string) *InstallMcpMarketItemResponseBodyData {
	s.Addresses = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetCustomTags(v []*string) *InstallMcpMarketItemResponseBodyData {
	s.CustomTags = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetDeploymentConfig(v *InstallMcpMarketItemResponseBodyDataDeploymentConfig) *InstallMcpMarketItemResponseBodyData {
	s.DeploymentConfig = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetDescription(v string) *InstallMcpMarketItemResponseBodyData {
	s.Description = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetEndpoint(v string) *InstallMcpMarketItemResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetFunctionName(v string) *InstallMcpMarketItemResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetMarketSource(v *InstallMcpMarketItemResponseBodyDataMarketSource) *InstallMcpMarketItemResponseBodyData {
	s.MarketSource = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetMcpServerId(v string) *InstallMcpMarketItemResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetName(v string) *InstallMcpMarketItemResponseBodyData {
	s.Name = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetOfficialTag(v string) *InstallMcpMarketItemResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetProtocol(v string) *InstallMcpMarketItemResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetStatus(v string) *InstallMcpMarketItemResponseBodyData {
	s.Status = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetStatusReason(v string) *InstallMcpMarketItemResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetTemplate(v *InstallMcpMarketItemResponseBodyDataTemplate) *InstallMcpMarketItemResponseBodyData {
	s.Template = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetType(v string) *InstallMcpMarketItemResponseBodyData {
	s.Type = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) SetUsageActive(v bool) *InstallMcpMarketItemResponseBodyData {
	s.UsageActive = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyData) Validate() error {
	if s.DeploymentConfig != nil {
		if err := s.DeploymentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MarketSource != nil {
		if err := s.MarketSource.Validate(); err != nil {
			return err
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfig struct {
	// The MCP ingress access control settings.
	AccessControl *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// The artifact type. Valid values:
	//
	// - Code: ZIP code package.
	//
	// - Container: custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfig) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetAccessControl() *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetAgentIdentityConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetCodeConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetContainerConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetHookConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetLogConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetMcpConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetNasConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetNetworkConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetOssMountConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetParameterTransformConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetProxyConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) GetRuntimeConfiguration() *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetAccessControl(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetAgentIdentityConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetArtifactType(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetCodeConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetContainerConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetHookConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetLogConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetMcpConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetNasConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetNetworkConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetOssMountConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetParameterTransformConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetProxyConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) SetRuntimeConfiguration(v *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) *InstallMcpMarketItemResponseBodyDataDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfig) Validate() error {
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

type InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl struct {
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
	// - CREDENTIAL: AgentCore credential-based access.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) SetCredentialId(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) SetEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) SetMode(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. This token is used to create or update a code deployment after the presigned upload is complete.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command, with arguments passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the full subcommand, and remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) SetCommand(v []*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) SetLanguage(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration struct {
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
	// The MCP Runtime mode. Custom containers must expose a standard MCP endpoint. Set this parameter to SELF_HOSTED.
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetCommand(v []*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetImage(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) SetSourceType(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration struct {
	// The hooks executed in array order: PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL.
	Hooks []*InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) GetHooks() []*InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) SetHooks(v []*InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfiguration) Validate() error {
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

type InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks struct {
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
	// The timeout period. Unit: milliseconds.
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetDescription(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEvent(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) SetUrl(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) SetLogstore(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) SetProject(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) GetMountPoints() []*InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) SetGroupId(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) SetMountPoints(v []*InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) SetUserId(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfiguration) Validate() error {
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

type InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) SetVpcId(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) GetMountPoints() []*InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) SetMountPoints(v []*InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfiguration) Validate() error {
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

type InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) SetEnabled(v bool) *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration struct {
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

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetPort(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataMarketSource struct {
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataMarketSource) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataMarketSource) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *InstallMcpMarketItemResponseBodyDataMarketSource) SetMarketItemId(v string) *InstallMcpMarketItemResponseBodyDataMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataMarketSource) Validate() error {
	return dara.Validate(s)
}

type InstallMcpMarketItemResponseBodyDataTemplate struct {
	// The template version currently applied to the MCP.
	//
	// example:
	//
	// 1.0.0
	AppliedTemplateVersion *string `json:"appliedTemplateVersion,omitempty" xml:"appliedTemplateVersion,omitempty"`
	// The latest template version.
	//
	// example:
	//
	// 1.1.0
	LatestTemplateVersion *string `json:"latestTemplateVersion,omitempty" xml:"latestTemplateVersion,omitempty"`
	// The template schema version.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The template input schema, represented as a JSON Schema string.
	//
	// example:
	//
	// {"type":"object","properties":{"addresses":{"type":"array","items":{"type":"string"}}}}
	TemplateInputSchema *string `json:"templateInputSchema,omitempty" xml:"templateInputSchema,omitempty"`
	// Indicates whether a newer template version is available for update.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" xml:"updateAvailable,omitempty"`
}

func (s InstallMcpMarketItemResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) GetAppliedTemplateVersion() *string {
	return s.AppliedTemplateVersion
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) GetLatestTemplateVersion() *string {
	return s.LatestTemplateVersion
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) SetAppliedTemplateVersion(v string) *InstallMcpMarketItemResponseBodyDataTemplate {
	s.AppliedTemplateVersion = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) SetLatestTemplateVersion(v string) *InstallMcpMarketItemResponseBodyDataTemplate {
	s.LatestTemplateVersion = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) SetSchemaVersion(v string) *InstallMcpMarketItemResponseBodyDataTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) SetTemplateInputSchema(v string) *InstallMcpMarketItemResponseBodyDataTemplate {
	s.TemplateInputSchema = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) SetUpdateAvailable(v bool) *InstallMcpMarketItemResponseBodyDataTemplate {
	s.UpdateAvailable = &v
	return s
}

func (s *InstallMcpMarketItemResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}

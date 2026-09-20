// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpTemplateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpTemplateConfigResponseBody
	GetCode() *string
	SetData(v *UpdateMcpTemplateConfigResponseBodyData) *UpdateMcpTemplateConfigResponseBody
	GetData() *UpdateMcpTemplateConfigResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMcpTemplateConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMcpTemplateConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpTemplateConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMcpTemplateConfigResponseBody
	GetSuccess() *bool
}

type UpdateMcpTemplateConfigResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *UpdateMcpTemplateConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpTemplateConfigResponseBody) GetData() *UpdateMcpTemplateConfigResponseBodyData {
	return s.Data
}

func (s *UpdateMcpTemplateConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMcpTemplateConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpTemplateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpTemplateConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcpTemplateConfigResponseBody) SetCode(v string) *UpdateMcpTemplateConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) SetData(v *UpdateMcpTemplateConfigResponseBodyData) *UpdateMcpTemplateConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) SetHttpStatusCode(v int32) *UpdateMcpTemplateConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) SetMessage(v string) *UpdateMcpTemplateConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) SetRequestId(v string) *UpdateMcpTemplateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) SetSuccess(v bool) *UpdateMcpTemplateConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpTemplateConfigResponseBodyData struct {
	// The list of remote MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The MCP service description.
	//
	// example:
	//
	// An MCP service for querying knowledge bases
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
	MarketSource *UpdateMcpTemplateConfigResponseBodyDataMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
	// The MCP server ID.
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
	// The official usage tag, managed by the server.
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
	Template *UpdateMcpTemplateConfigResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
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
	// Indicates whether the MCP is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetDeploymentConfig() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	return s.DeploymentConfig
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetMarketSource() *UpdateMcpTemplateConfigResponseBodyDataMarketSource {
	return s.MarketSource
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetTemplate() *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	return s.Template
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetType() *string {
	return s.Type
}

func (s *UpdateMcpTemplateConfigResponseBodyData) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetAddresses(v []*string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Addresses = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetCustomTags(v []*string) *UpdateMcpTemplateConfigResponseBodyData {
	s.CustomTags = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetDeploymentConfig(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) *UpdateMcpTemplateConfigResponseBodyData {
	s.DeploymentConfig = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetDescription(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetEndpoint(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetFunctionName(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetMarketSource(v *UpdateMcpTemplateConfigResponseBodyDataMarketSource) *UpdateMcpTemplateConfigResponseBodyData {
	s.MarketSource = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetMcpServerId(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetName(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetOfficialTag(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetProtocol(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetStatus(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetStatusReason(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetTemplate(v *UpdateMcpTemplateConfigResponseBodyDataTemplate) *UpdateMcpTemplateConfigResponseBodyData {
	s.Template = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetType(v string) *UpdateMcpTemplateConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) SetUsageActive(v bool) *UpdateMcpTemplateConfigResponseBodyData {
	s.UsageActive = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyData) Validate() error {
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

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig struct {
	// The MCP ingress access control settings.
	AccessControl *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
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
	CodeConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transform and result enhancement configuration.
	ParameterTransformConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetAccessControl() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetAgentIdentityConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetCodeConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetContainerConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetHookConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetLogConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetMcpConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetNasConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetNetworkConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetOssMountConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetParameterTransformConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetProxyConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) GetRuntimeConfiguration() *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetAccessControl(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetAgentIdentityConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetArtifactType(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetCodeConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetContainerConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetHookConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetLogConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetMcpConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetNasConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetNetworkConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetOssMountConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetParameterTransformConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetProxyConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) SetRuntimeConfiguration(v *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfig) Validate() error {
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

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) SetCredentialId(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) SetEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) SetMode(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. Use this token to create or update a code deployment after completing the pre-signed upload.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command, with each argument passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the full subcommand, and the remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) SetCommand(v []*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) SetLanguage(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration struct {
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
	// The MCP runtime mode for the custom container. The custom container must expose a standard MCP endpoint on its own. Set this parameter to SELF_HOSTED.
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetCommand(v []*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetImage(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) SetSourceType(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration struct {
	// The hooks executed in array order: PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL.
	Hooks []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) GetHooks() []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) SetHooks(v []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfiguration) Validate() error {
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

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks struct {
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
	// Log MCP tool calls
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetDescription(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEvent(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) SetUrl(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) SetLogstore(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) SetProject(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) GetMountPoints() []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) SetGroupId(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) SetMountPoints(v []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) SetUserId(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfiguration) Validate() error {
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

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) SetVpcId(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) GetMountPoints() []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) SetMountPoints(v []*UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfiguration) Validate() error {
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

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) SetEnabled(v bool) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration struct {
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
	// The ARN of the RAM role used by user code to access downstream Alibaba Cloud resources.
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

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetPort(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataMarketSource struct {
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s UpdateMcpTemplateConfigResponseBodyDataMarketSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataMarketSource) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *UpdateMcpTemplateConfigResponseBodyDataMarketSource) SetMarketItemId(v string) *UpdateMcpTemplateConfigResponseBodyDataMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataMarketSource) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpTemplateConfigResponseBodyDataTemplate struct {
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

func (s UpdateMcpTemplateConfigResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) GetAppliedTemplateVersion() *string {
	return s.AppliedTemplateVersion
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) GetLatestTemplateVersion() *string {
	return s.LatestTemplateVersion
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) SetAppliedTemplateVersion(v string) *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	s.AppliedTemplateVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) SetLatestTemplateVersion(v string) *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	s.LatestTemplateVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) SetSchemaVersion(v string) *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) SetTemplateInputSchema(v string) *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	s.TemplateInputSchema = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) SetUpdateAvailable(v bool) *UpdateMcpTemplateConfigResponseBodyDataTemplate {
	s.UpdateAvailable = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}

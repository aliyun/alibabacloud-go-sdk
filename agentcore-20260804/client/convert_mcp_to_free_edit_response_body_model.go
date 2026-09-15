// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertMcpToFreeEditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertMcpToFreeEditResponseBody
	GetCode() *string
	SetData(v *ConvertMcpToFreeEditResponseBodyData) *ConvertMcpToFreeEditResponseBody
	GetData() *ConvertMcpToFreeEditResponseBodyData
	SetHttpStatusCode(v int32) *ConvertMcpToFreeEditResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ConvertMcpToFreeEditResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertMcpToFreeEditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertMcpToFreeEditResponseBody
	GetSuccess() *bool
}

type ConvertMcpToFreeEditResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ConvertMcpToFreeEditResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The request ID, which is used to locate and troubleshoot requests.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConvertMcpToFreeEditResponseBody) GetData() *ConvertMcpToFreeEditResponseBodyData {
	return s.Data
}

func (s *ConvertMcpToFreeEditResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConvertMcpToFreeEditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertMcpToFreeEditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertMcpToFreeEditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertMcpToFreeEditResponseBody) SetCode(v string) *ConvertMcpToFreeEditResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) SetData(v *ConvertMcpToFreeEditResponseBodyData) *ConvertMcpToFreeEditResponseBody {
	s.Data = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) SetHttpStatusCode(v int32) *ConvertMcpToFreeEditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) SetMessage(v string) *ConvertMcpToFreeEditResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) SetRequestId(v string) *ConvertMcpToFreeEditResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) SetSuccess(v bool) *ConvertMcpToFreeEditResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertMcpToFreeEditResponseBodyData struct {
	// The list of remote MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for the code deployment MCP.
	DeploymentConfig *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
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
	// The Function Compute function name that corresponds to the code deployment MCP.
	//
	// example:
	//
	// agentcore-mcp-example
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The marketplace template from which the MCP service originates.
	MarketSource *ConvertMcpToFreeEditResponseBodyDataMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
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
	// The status of the MCP service.
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
	// The template version and input schema bound to the MCP service.
	Template *ConvertMcpToFreeEditResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The MCP type. Valid values: DIRECT_PROXY (direct proxy), HTTP_TO_MCP (HTTP to MCP), and CODE_PACKAGE (code deployment).
	//
	// example:
	//
	// CODE_PACKAGE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Indicates whether the MCP service is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetDeploymentConfig() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	return s.DeploymentConfig
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetMarketSource() *ConvertMcpToFreeEditResponseBodyDataMarketSource {
	return s.MarketSource
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetTemplate() *ConvertMcpToFreeEditResponseBodyDataTemplate {
	return s.Template
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ConvertMcpToFreeEditResponseBodyData) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetAddresses(v []*string) *ConvertMcpToFreeEditResponseBodyData {
	s.Addresses = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetCustomTags(v []*string) *ConvertMcpToFreeEditResponseBodyData {
	s.CustomTags = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetDeploymentConfig(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) *ConvertMcpToFreeEditResponseBodyData {
	s.DeploymentConfig = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetDescription(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Description = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetEndpoint(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetFunctionName(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetMarketSource(v *ConvertMcpToFreeEditResponseBodyDataMarketSource) *ConvertMcpToFreeEditResponseBodyData {
	s.MarketSource = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetMcpServerId(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetName(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Name = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetOfficialTag(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetProtocol(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetStatus(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Status = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetStatusReason(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetTemplate(v *ConvertMcpToFreeEditResponseBodyDataTemplate) *ConvertMcpToFreeEditResponseBodyData {
	s.Template = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetType(v string) *ConvertMcpToFreeEditResponseBodyData {
	s.Type = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) SetUsageActive(v bool) *ConvertMcpToFreeEditResponseBodyData {
	s.UsageActive = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyData) Validate() error {
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

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfig struct {
	// The MCP ingress access control configuration.
	AccessControl *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The agent identity configuration.
	AgentIdentityConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// The artifact type. Code indicates a ZIP code package. Container indicates a custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transform and result enhancement configuration.
	ParameterTransformConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetAccessControl() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetAgentIdentityConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetCodeConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetContainerConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetHookConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetLogConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetMcpConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetNasConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetNetworkConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetOssMountConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetParameterTransformConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetProxyConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) GetRuntimeConfiguration() *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetAccessControl(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetAgentIdentityConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetArtifactType(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetCodeConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetContainerConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetHookConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetLogConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetMcpConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetNasConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetNetworkConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetOssMountConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetParameterTransformConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetProxyConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) SetRuntimeConfiguration(v *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfig) Validate() error {
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

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl struct {
	// The AgentCore credential referenced when mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Specifies whether to enable ingress access control.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access control mode. ANONYMOUS indicates anonymous access. CREDENTIAL indicates access by using an AgentCore credential.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) SetCredentialId(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) SetEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) SetMode(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration struct {
	// Indicates whether authorization is enabled.
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
	// Indicates whether agent identity is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. This token is used to create or update a code deployment after the pre-signed upload is complete.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command, with arguments passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the full subcommand, and the remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) SetCommand(v []*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) SetLanguage(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration struct {
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
	// The MCP runtime mode. Custom containers must expose a standard MCP on their own. The value is SELF_HOSTED.
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

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetCommand(v []*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetImage(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) SetSourceType(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration struct {
	// The hooks are executed in array order for PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL events.
	Hooks []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) GetHooks() []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) SetHooks(v []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfiguration) Validate() error {
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

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks struct {
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
	// Indicates whether the hook is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The event that triggers the hook.
	//
	// example:
	//
	// PRE_CALL_TOOL
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// The request headers of the hook.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The timeout period of the hook. Unit: milliseconds.
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

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetDescription(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEvent(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) SetUrl(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration struct {
	// Specifies whether to collect instance metrics.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to collect request metrics.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log segmentation begin rule for Function Compute (FC).
	//
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// mcp-logs
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// agentcore-mcp-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) SetLogstore(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) SetProject(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration struct {
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
	// The session idle timeout period, in seconds. Default value: 1800.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The maximum session lifetime, in seconds. Default value: 21600.
	//
	// example:
	//
	// 21600
	SessionMaxLifetimeSeconds *int32 `json:"sessionMaxLifetimeSeconds,omitempty" xml:"sessionMaxLifetimeSeconds,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration struct {
	// The ID of the runtime user group.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The ID of the runtime user.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) GetMountPoints() []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) SetGroupId(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) SetMountPoints(v []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) SetUserId(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfiguration) Validate() error {
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

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints struct {
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

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration struct {
	// The network mode.
	//
	// example:
	//
	// PUBLIC
	NetworkMode *string `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-example
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) SetVpcId(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) GetMountPoints() []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) SetMountPoints(v []*ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfiguration) Validate() error {
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

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The path within the OSS bucket.
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

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration struct {
	// Indicates whether parameter transform and result enhancement is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The reserved reference to the parameter transform and result enhancement rule set.
	//
	// example:
	//
	// rules-1
	RuleSetId *string `json:"ruleSetId,omitempty" xml:"ruleSetId,omitempty"`
	// The version of the transform rule.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) SetEnabled(v bool) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration struct {
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
	// The number of concurrent requests per instance. Default value: 200.
	//
	// example:
	//
	// 200
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// The memory specification. Unit: MB. Default value: 512.
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
	// The function timeout. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetPort(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataMarketSource struct {
	// The ID of the MCP marketplace template.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataMarketSource) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataMarketSource) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *ConvertMcpToFreeEditResponseBodyDataMarketSource) SetMarketItemId(v string) *ConvertMcpToFreeEditResponseBodyDataMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataMarketSource) Validate() error {
	return dara.Validate(s)
}

type ConvertMcpToFreeEditResponseBodyDataTemplate struct {
	// The template version currently applied to the MCP service.
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
	// The schema version of the template.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The input schema of the template, represented as a JSON Schema string.
	//
	// example:
	//
	// {"type":"object","properties":{"addresses":{"type":"array","items":{"type":"string"}}}}
	TemplateInputSchema *string `json:"templateInputSchema,omitempty" xml:"templateInputSchema,omitempty"`
	// Indicates whether a newer template version is available for update.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" xml:"updateAvailable,omitempty"`
}

func (s ConvertMcpToFreeEditResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) GetAppliedTemplateVersion() *string {
	return s.AppliedTemplateVersion
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) GetLatestTemplateVersion() *string {
	return s.LatestTemplateVersion
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) SetAppliedTemplateVersion(v string) *ConvertMcpToFreeEditResponseBodyDataTemplate {
	s.AppliedTemplateVersion = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) SetLatestTemplateVersion(v string) *ConvertMcpToFreeEditResponseBodyDataTemplate {
	s.LatestTemplateVersion = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) SetSchemaVersion(v string) *ConvertMcpToFreeEditResponseBodyDataTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) SetTemplateInputSchema(v string) *ConvertMcpToFreeEditResponseBodyDataTemplate {
	s.TemplateInputSchema = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) SetUpdateAvailable(v bool) *ConvertMcpToFreeEditResponseBodyDataTemplate {
	s.UpdateAvailable = &v
	return s
}

func (s *ConvertMcpToFreeEditResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}

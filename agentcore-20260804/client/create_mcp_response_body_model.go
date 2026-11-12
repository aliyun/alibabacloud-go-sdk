// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpResponseBody
	GetCode() *string
	SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody
	GetData() *CreateMcpResponseBodyData
	SetHttpStatusCode(v int32) *CreateMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpResponseBody
	GetSuccess() *bool
}

type CreateMcpResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateMcpResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Request processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpResponseBody) GetData() *CreateMcpResponseBodyData {
	return s.Data
}

func (s *CreateMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpResponseBody) SetCode(v string) *CreateMcpResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpResponseBody) SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpResponseBody) SetHttpStatusCode(v int32) *CreateMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpResponseBody) SetMessage(v string) *CreateMcpResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpResponseBody) SetRequestId(v string) *CreateMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpResponseBody) SetSuccess(v bool) *CreateMcpResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpResponseBodyData struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. The enabled field indicates whether authentication is enabled. The directProxy field specifies the custom authentication header for the direct proxy. The httpToMcp field specifies the list of OpenAPI credentials for HTTP_TO_MCP.
	Auth *CreateMcpResponseBodyDataAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for the code deployment MCP.
	DeploymentConfig *CreateMcpResponseBodyDataDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description to illustrate the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP endpoint available for users or agents to invoke. This field is empty before deployment is complete.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Function Compute function name corresponding to the CODE_PACKAGE MCP. This field is empty before deployment is complete and for other types.
	//
	// example:
	//
	// agentcore-mcp-example
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The MCP source marketplace template.
	MarketSource *CreateMcpResponseBodyDataMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
	// The MCP server configuration.
	//
	// example:
	//
	// {"type":"object"}
	McpServerConfig *string `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
	// The MCP service ID.
	//
	// example:
	//
	// mcp-1234567890abcdef
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The official usage tag managed by the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// PUBLIC indicates public access. CONTROLLED indicates that access is controlled based on permission rules.
	//
	// example:
	//
	// CONTROLLED
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// SSE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The status.
	//
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status reason.
	//
	// example:
	//
	// Resource processing completed
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The Swagger configuration.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// The bound template version and input schema of the MCP.
	Template *CreateMcpResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// The type.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Indicates whether the MCP is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s CreateMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpResponseBodyData) GetAuth() *CreateMcpResponseBodyDataAuth {
	return s.Auth
}

func (s *CreateMcpResponseBodyData) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *CreateMcpResponseBodyData) GetDeploymentConfig() *CreateMcpResponseBodyDataDeploymentConfig {
	return s.DeploymentConfig
}

func (s *CreateMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateMcpResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *CreateMcpResponseBodyData) GetMarketSource() *CreateMcpResponseBodyDataMarketSource {
	return s.MarketSource
}

func (s *CreateMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *CreateMcpResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *CreateMcpResponseBodyData) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *CreateMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateMcpResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateMcpResponseBodyData) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpResponseBodyData) GetTemplate() *CreateMcpResponseBodyDataTemplate {
	return s.Template
}

func (s *CreateMcpResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateMcpResponseBodyData) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *CreateMcpResponseBodyData) SetAddresses(v []*string) *CreateMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *CreateMcpResponseBodyData) SetAuth(v *CreateMcpResponseBodyDataAuth) *CreateMcpResponseBodyData {
	s.Auth = v
	return s
}

func (s *CreateMcpResponseBodyData) SetCustomTags(v []*string) *CreateMcpResponseBodyData {
	s.CustomTags = v
	return s
}

func (s *CreateMcpResponseBodyData) SetDeploymentConfig(v *CreateMcpResponseBodyDataDeploymentConfig) *CreateMcpResponseBodyData {
	s.DeploymentConfig = v
	return s
}

func (s *CreateMcpResponseBodyData) SetDescription(v string) *CreateMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetEndpoint(v string) *CreateMcpResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetFunctionName(v string) *CreateMcpResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetMarketSource(v *CreateMcpResponseBodyDataMarketSource) *CreateMcpResponseBodyData {
	s.MarketSource = v
	return s
}

func (s *CreateMcpResponseBodyData) SetMcpServerConfig(v string) *CreateMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetMcpServerId(v string) *CreateMcpResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetName(v string) *CreateMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetOfficialTag(v string) *CreateMcpResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetPermissionMode(v string) *CreateMcpResponseBodyData {
	s.PermissionMode = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetProtocol(v string) *CreateMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetStatus(v string) *CreateMcpResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetStatusReason(v string) *CreateMcpResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetSwaggerConfig(v string) *CreateMcpResponseBodyData {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetTemplate(v *CreateMcpResponseBodyDataTemplate) *CreateMcpResponseBodyData {
	s.Template = v
	return s
}

func (s *CreateMcpResponseBodyData) SetType(v string) *CreateMcpResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetUsageActive(v bool) *CreateMcpResponseBodyData {
	s.UsageActive = &v
	return s
}

func (s *CreateMcpResponseBodyData) Validate() error {
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

type CreateMcpResponseBodyDataAuth struct {
	// The caller API Key authentication configuration for code-deployed MCP.
	CodePackage *CreateMcpResponseBodyDataAuthCodePackage `json:"codePackage,omitempty" xml:"codePackage,omitempty" type:"Struct"`
	// The direct proxy authentication configuration.
	DirectProxy *CreateMcpResponseBodyDataAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable the configuration.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*CreateMcpResponseBodyDataAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s CreateMcpResponseBodyDataAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuth) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuth) GetCodePackage() *CreateMcpResponseBodyDataAuthCodePackage {
	return s.CodePackage
}

func (s *CreateMcpResponseBodyDataAuth) GetDirectProxy() *CreateMcpResponseBodyDataAuthDirectProxy {
	return s.DirectProxy
}

func (s *CreateMcpResponseBodyDataAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataAuth) GetHttpToMcp() []*CreateMcpResponseBodyDataAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *CreateMcpResponseBodyDataAuth) SetCodePackage(v *CreateMcpResponseBodyDataAuthCodePackage) *CreateMcpResponseBodyDataAuth {
	s.CodePackage = v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) SetDirectProxy(v *CreateMcpResponseBodyDataAuthDirectProxy) *CreateMcpResponseBodyDataAuth {
	s.DirectProxy = v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) SetEnabled(v bool) *CreateMcpResponseBodyDataAuth {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) SetHttpToMcp(v []*CreateMcpResponseBodyDataAuthHttpToMcp) *CreateMcpResponseBodyDataAuth {
	s.HttpToMcp = v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) Validate() error {
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

type CreateMcpResponseBodyDataAuthCodePackage struct {
	// The API Key used to verify MCP callers.
	//
	// example:
	//
	// example-api-key
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The name of the request header that carries the API Key.
	//
	// example:
	//
	// X-API-Key
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
}

func (s CreateMcpResponseBodyDataAuthCodePackage) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuthCodePackage) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuthCodePackage) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateMcpResponseBodyDataAuthCodePackage) GetHeaderName() *string {
	return s.HeaderName
}

func (s *CreateMcpResponseBodyDataAuthCodePackage) SetApiKey(v string) *CreateMcpResponseBodyDataAuthCodePackage {
	s.ApiKey = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthCodePackage) SetHeaderName(v string) *CreateMcpResponseBodyDataAuthCodePackage {
	s.HeaderName = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthCodePackage) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataAuthDirectProxy struct {
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the authentication parameter.
	//
	// example:
	//
	// example-credential
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMcpResponseBodyDataAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) SetName(v string) *CreateMcpResponseBodyDataAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) SetValue(v string) *CreateMcpResponseBodyDataAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataAuthHttpToMcp struct {
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
	// The location of the credential.
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

func (s CreateMcpResponseBodyDataAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetCredential(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetId(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetName(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetPosition(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetType(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfig struct {
	// The MCP entry access control.
	AccessControl *CreateMcpResponseBodyDataDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The agent identity configuration.
	AgentIdentityConfiguration *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// The artifact type. Code indicates a ZIP code package, and Container indicates a custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s CreateMcpResponseBodyDataDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetAccessControl() *CreateMcpResponseBodyDataDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetAgentIdentityConfiguration() *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetCodeConfiguration() *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetContainerConfiguration() *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetHookConfiguration() *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetLogConfiguration() *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetMcpConfiguration() *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetNasConfiguration() *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetNetworkConfiguration() *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetOssMountConfiguration() *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetParameterTransformConfiguration() *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetProxyConfiguration() *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) GetRuntimeConfiguration() *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetAccessControl(v *CreateMcpResponseBodyDataDeploymentConfigAccessControl) *CreateMcpResponseBodyDataDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetAgentIdentityConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetArtifactType(v string) *CreateMcpResponseBodyDataDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetCodeConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetContainerConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetHookConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetLogConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetMcpConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetNasConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetNetworkConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetOssMountConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetParameterTransformConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetProxyConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) SetRuntimeConfiguration(v *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) *CreateMcpResponseBodyDataDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfig) Validate() error {
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

type CreateMcpResponseBodyDataDeploymentConfigAccessControl struct {
	// The ID of the AgentCore credential referenced when the mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Specifies whether to enable entry access control.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access control mode. ANONYMOUS indicates anonymous access. CREDENTIAL indicates authentication by using AgentCore access credentials.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) SetCredentialId(v string) *CreateMcpResponseBodyDataDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) SetEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) SetMode(v string) *CreateMcpResponseBodyDataDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether to enable authorization.
	AuthorizationEnabled *bool `json:"authorizationEnabled,omitempty" xml:"authorizationEnabled,omitempty"`
	// The credential provider ARN.
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
	// Specifies whether to enable agent identity.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl, which is used to create or update a code deployment after a presigned upload is complete.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The complete startup command, passed sequentially by parameter boundaries. For example, when using supergateway to start a stdio MCP, pass supergateway, --stdio, the complete subcommand, and the remaining parameters.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The runtime of the code package. Valid values: python3.13, nodejs22, and java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetCommand(v []*string) *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetLanguage(v string) *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration struct {
	// The ACR instance ID.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The startup command.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The container entrypoint parameters.
	Entrypoint []*string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty" type:"Repeated"`
	// The container image URL.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/mcp:1.0.0
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The image repository type.
	//
	// example:
	//
	// ACR
	ImageRegistryType *string `json:"imageRegistryType,omitempty" xml:"imageRegistryType,omitempty"`
	// Custom containers currently must expose the standard MCP themselves. Set the value to SELF_HOSTED.
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

func (s CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetCommand(v []*string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetImage(v string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetSourceType(v string) *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigHookConfiguration struct {
	// Executes the PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL hooks in array order.
	Hooks []*CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) GetHooks() []*CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) SetHooks(v []*CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfiguration) Validate() error {
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

type CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks struct {
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
	// Record MCP tool invocation
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
	// The timeout period of the hook, in milliseconds.
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

func (s CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetDescription(v string) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEvent(v string) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetUrl(v string) *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigLogConfiguration struct {
	// Specifies whether to collect instance metrics.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to collect request metrics.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The start rule for FC log splitting.
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

func (s CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) SetLogstore(v string) *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) SetProject(v string) *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration struct {
	// The MCP endpoint path, such as /mcp or /sse.
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
	// The session idle timeout period in seconds. Default value: 1800.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The maximum session lifetime in seconds. Default value: 21600.
	//
	// example:
	//
	// 21600
	SessionMaxLifetimeSeconds *int32 `json:"sessionMaxLifetimeSeconds,omitempty" xml:"sessionMaxLifetimeSeconds,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) GetMountPoints() []*CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) SetGroupId(v int32) *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) SetMountPoints(v []*CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) SetUserId(v int32) *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfiguration) Validate() error {
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

type CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints struct {
	// Specifies whether to enable TLS.
	EnableTls *bool `json:"enableTls,omitempty" xml:"enableTls,omitempty"`
	// The local mount directory.
	//
	// example:
	//
	// /mnt/data
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// The NAS service address.
	//
	// example:
	//
	// example.cn-hangzhou.nas.aliyuncs.com
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration struct {
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

func (s CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetVpcId(v string) *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) GetMountPoints() []*CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) SetMountPoints(v []*CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfiguration) Validate() error {
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

type CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The path in the OSS bucket.
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

func (s CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration struct {
	// Specifies whether to enable parameter transformation and result enhancement.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The reserved parameter transformation and result enhancement rule set reference.
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

func (s CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) SetEnabled(v bool) *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration struct {
	// The CPU specification in cores. Default value: 0.25.
	//
	// example:
	//
	// 0.25
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The ephemeral disk size in MB. Valid values: 512 and 10240.
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
	// The memory specification in MB. Default value: 512.
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
	// The function timeout period in seconds. Default value: 300.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetPort(v int32) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *CreateMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataMarketSource struct {
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s CreateMcpResponseBodyDataMarketSource) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataMarketSource) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *CreateMcpResponseBodyDataMarketSource) SetMarketItemId(v string) *CreateMcpResponseBodyDataMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *CreateMcpResponseBodyDataMarketSource) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataTemplate struct {
	// The applied template version of the current MCP.
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
	// Indicates whether an updated template version is available.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" xml:"updateAvailable,omitempty"`
}

func (s CreateMcpResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataTemplate) GetAppliedTemplateVersion() *string {
	return s.AppliedTemplateVersion
}

func (s *CreateMcpResponseBodyDataTemplate) GetLatestTemplateVersion() *string {
	return s.LatestTemplateVersion
}

func (s *CreateMcpResponseBodyDataTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *CreateMcpResponseBodyDataTemplate) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *CreateMcpResponseBodyDataTemplate) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *CreateMcpResponseBodyDataTemplate) SetAppliedTemplateVersion(v string) *CreateMcpResponseBodyDataTemplate {
	s.AppliedTemplateVersion = &v
	return s
}

func (s *CreateMcpResponseBodyDataTemplate) SetLatestTemplateVersion(v string) *CreateMcpResponseBodyDataTemplate {
	s.LatestTemplateVersion = &v
	return s
}

func (s *CreateMcpResponseBodyDataTemplate) SetSchemaVersion(v string) *CreateMcpResponseBodyDataTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *CreateMcpResponseBodyDataTemplate) SetTemplateInputSchema(v string) *CreateMcpResponseBodyDataTemplate {
	s.TemplateInputSchema = &v
	return s
}

func (s *CreateMcpResponseBodyDataTemplate) SetUpdateAvailable(v bool) *CreateMcpResponseBodyDataTemplate {
	s.UpdateAvailable = &v
	return s
}

func (s *CreateMcpResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}

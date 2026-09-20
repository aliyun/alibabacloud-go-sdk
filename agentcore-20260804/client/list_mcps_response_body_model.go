// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMcpsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody
	GetItems() []*ListMcpsResponseBodyItems
	SetMaxResults(v int32) *ListMcpsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMcpsResponseBody
	GetTotalCount() *int64
}

type ListMcpsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list data.
	Items []*ListMcpsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// Request processing succeeded
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMcpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpsResponseBody) GetItems() []*ListMcpsResponseBodyItems {
	return s.Items
}

func (s *ListMcpsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcpsResponseBody) SetCode(v string) *ListMcpsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpsResponseBody) SetHttpStatusCode(v int32) *ListMcpsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpsResponseBody) SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpsResponseBody) SetMaxResults(v int32) *ListMcpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpsResponseBody) SetMessage(v string) *ListMcpsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpsResponseBody) SetNextToken(v string) *ListMcpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpsResponseBody) SetRequestId(v string) *ListMcpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpsResponseBody) SetSuccess(v bool) *ListMcpsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpsResponseBody) SetTotalCount(v int64) *ListMcpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpsResponseBodyItems struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. enabled indicates whether it is enabled. directProxy specifies custom authentication headers for direct proxy. httpToMcp specifies the OpenAPI credential list for HTTP_TO_MCP.
	Auth *ListMcpsResponseBodyItemsAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The MCP creation time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-23T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *ListMcpsResponseBodyItemsDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP endpoint available for users or agents to call. This value is empty before deployment is complete.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Function Compute function name corresponding to the CODE_PACKAGE MCP. This value is empty before deployment is complete or for other types.
	//
	// example:
	//
	// agentcore-mcp-example
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The MCP marketplace source template.
	MarketSource *ListMcpsResponseBodyItemsMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
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
	// The official tag managed by the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
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
	// The type.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the MCP was last updated, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-23T01:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Indicates whether the MCP is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s ListMcpsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItems) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListMcpsResponseBodyItems) GetAuth() *ListMcpsResponseBodyItemsAuth {
	return s.Auth
}

func (s *ListMcpsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMcpsResponseBodyItems) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *ListMcpsResponseBodyItems) GetDeploymentConfig() *ListMcpsResponseBodyItemsDeploymentConfig {
	return s.DeploymentConfig
}

func (s *ListMcpsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpsResponseBodyItems) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListMcpsResponseBodyItems) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListMcpsResponseBodyItems) GetMarketSource() *ListMcpsResponseBodyItemsMarketSource {
	return s.MarketSource
}

func (s *ListMcpsResponseBodyItems) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItems) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *ListMcpsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMcpsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListMcpsResponseBodyItems) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListMcpsResponseBodyItems) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *ListMcpsResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *ListMcpsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListMcpsResponseBodyItems) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *ListMcpsResponseBodyItems) SetAddresses(v []*string) *ListMcpsResponseBodyItems {
	s.Addresses = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetAuth(v *ListMcpsResponseBodyItemsAuth) *ListMcpsResponseBodyItems {
	s.Auth = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetCreatedAt(v string) *ListMcpsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetCustomTags(v []*string) *ListMcpsResponseBodyItems {
	s.CustomTags = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetDeploymentConfig(v *ListMcpsResponseBodyItemsDeploymentConfig) *ListMcpsResponseBodyItems {
	s.DeploymentConfig = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetDescription(v string) *ListMcpsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetEndpoint(v string) *ListMcpsResponseBodyItems {
	s.Endpoint = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetFunctionName(v string) *ListMcpsResponseBodyItems {
	s.FunctionName = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetMarketSource(v *ListMcpsResponseBodyItemsMarketSource) *ListMcpsResponseBodyItems {
	s.MarketSource = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetMcpServerId(v string) *ListMcpsResponseBodyItems {
	s.McpServerId = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetName(v string) *ListMcpsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetOfficialTag(v string) *ListMcpsResponseBodyItems {
	s.OfficialTag = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetProtocol(v string) *ListMcpsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetStatus(v string) *ListMcpsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetStatusReason(v string) *ListMcpsResponseBodyItems {
	s.StatusReason = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetSwaggerConfig(v string) *ListMcpsResponseBodyItems {
	s.SwaggerConfig = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetType(v string) *ListMcpsResponseBodyItems {
	s.Type = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetUpdatedAt(v string) *ListMcpsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetUsageActive(v bool) *ListMcpsResponseBodyItems {
	s.UsageActive = &v
	return s
}

func (s *ListMcpsResponseBodyItems) Validate() error {
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
	return nil
}

type ListMcpsResponseBodyItemsAuth struct {
	// The caller API Key authentication configuration for code-deployed MCP.
	CodePackage *ListMcpsResponseBodyItemsAuthCodePackage `json:"codePackage,omitempty" xml:"codePackage,omitempty" type:"Struct"`
	// The authentication configuration for direct proxy.
	DirectProxy *ListMcpsResponseBodyItemsAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Indicates whether the configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*ListMcpsResponseBodyItemsAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s ListMcpsResponseBodyItemsAuth) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuth) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuth) GetCodePackage() *ListMcpsResponseBodyItemsAuthCodePackage {
	return s.CodePackage
}

func (s *ListMcpsResponseBodyItemsAuth) GetDirectProxy() *ListMcpsResponseBodyItemsAuthDirectProxy {
	return s.DirectProxy
}

func (s *ListMcpsResponseBodyItemsAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsAuth) GetHttpToMcp() []*ListMcpsResponseBodyItemsAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *ListMcpsResponseBodyItemsAuth) SetCodePackage(v *ListMcpsResponseBodyItemsAuthCodePackage) *ListMcpsResponseBodyItemsAuth {
	s.CodePackage = v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) SetDirectProxy(v *ListMcpsResponseBodyItemsAuthDirectProxy) *ListMcpsResponseBodyItemsAuth {
	s.DirectProxy = v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) SetEnabled(v bool) *ListMcpsResponseBodyItemsAuth {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) SetHttpToMcp(v []*ListMcpsResponseBodyItemsAuthHttpToMcp) *ListMcpsResponseBodyItemsAuth {
	s.HttpToMcp = v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) Validate() error {
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

type ListMcpsResponseBodyItemsAuthCodePackage struct {
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

func (s ListMcpsResponseBodyItemsAuthCodePackage) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuthCodePackage) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuthCodePackage) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListMcpsResponseBodyItemsAuthCodePackage) GetHeaderName() *string {
	return s.HeaderName
}

func (s *ListMcpsResponseBodyItemsAuthCodePackage) SetApiKey(v string) *ListMcpsResponseBodyItemsAuthCodePackage {
	s.ApiKey = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthCodePackage) SetHeaderName(v string) *ListMcpsResponseBodyItemsAuthCodePackage {
	s.HeaderName = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthCodePackage) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsAuthDirectProxy struct {
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

func (s ListMcpsResponseBodyItemsAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) SetName(v string) *ListMcpsResponseBodyItemsAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) SetValue(v string) *ListMcpsResponseBodyItemsAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsAuthHttpToMcp struct {
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

func (s ListMcpsResponseBodyItemsAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetCredential(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetId(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetName(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetPosition(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetType(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfig struct {
	// The MCP ingress access control.
	AccessControl *ListMcpsResponseBodyItemsDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The Agent Identity configuration.
	AgentIdentityConfiguration *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// Code indicates a ZIP code package. Container indicates a custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfig) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetAccessControl() *ListMcpsResponseBodyItemsDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetAgentIdentityConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetCodeConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetContainerConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetHookConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetLogConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetMcpConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetNasConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetNetworkConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetOssMountConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetParameterTransformConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetProxyConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) GetRuntimeConfiguration() *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetAccessControl(v *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetAgentIdentityConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetArtifactType(v string) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetCodeConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetContainerConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetHookConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetLogConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetMcpConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetNasConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetNetworkConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetOssMountConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetParameterTransformConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetProxyConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) SetRuntimeConfiguration(v *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) *ListMcpsResponseBodyItemsDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfig) Validate() error {
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

type ListMcpsResponseBodyItemsDeploymentConfigAccessControl struct {
	// The AgentCore Credential referenced when mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Indicates whether ingress access control is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// ANONYMOUS indicates anonymous access. CREDENTIAL indicates using AgentCore access credentials.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) SetCredentialId(v string) *ListMcpsResponseBodyItemsDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) SetEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) SetMode(v string) *ListMcpsResponseBodyItemsDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration struct {
	// Indicates whether authorization is enabled.
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
	// Indicates whether Agent Identity is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. Use this token to create or update a code deployment after completing the pre-signed upload.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The full startup command, with arguments passed in order by parameter boundary. For example, when using supergateway to start a stdio MCP, pass in supergateway, --stdio, the full subcommand, and the remaining arguments.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// The code package runtime: python3.13, nodejs22, or java17.
	//
	// example:
	//
	// python3.13
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) SetCommand(v []*string) *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) SetLanguage(v string) *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration struct {
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
	// The custom container must expose a standard MCP endpoint on its own. The value is SELF_HOSTED.
	//
	// example:
	//
	// SELF_HOSTED
	McpRuntimeMode *string `json:"mcpRuntimeMode,omitempty" xml:"mcpRuntimeMode,omitempty"`
	// The value is fixed to CONTAINER_IMAGE.
	//
	// example:
	//
	// CONTAINER_IMAGE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetCommand(v []*string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetImage(v string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) SetSourceType(v string) *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration struct {
	// The PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL hooks are executed in array order.
	Hooks []*ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) GetHooks() []*ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) SetHooks(v []*ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfiguration) Validate() error {
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

type ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks struct {
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
	// Record MCP tool calling invokes
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the hook is enabled.
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

func (s ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetDescription(v string) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetEvent(v string) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) SetUrl(v string) *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration struct {
	// Indicates whether instance metrics collection is enabled.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Indicates whether request metrics collection is enabled.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log segmentation begin rule for Function Compute (FC).
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

func (s ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) SetLogstore(v string) *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) SetProject(v string) *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration struct {
	// The MCP endpoint path, such as /mcp or /sse.
	//
	// example:
	//
	// /mcp
	EndpointPath *string `json:"endpointPath,omitempty" xml:"endpointPath,omitempty"`
	// The value is fixed to 1.
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

func (s ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration struct {
	// The runtime user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The runtime user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) GetMountPoints() []*ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) SetGroupId(v int32) *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) SetMountPoints(v []*ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) SetUserId(v int32) *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfiguration) Validate() error {
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

type ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints struct {
	// Indicates whether TLS is enabled.
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

func (s ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration struct {
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

func (s ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) SetVpcId(v string) *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) GetMountPoints() []*ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) SetMountPoints(v []*ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfiguration) Validate() error {
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

type ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints struct {
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
	// Indicates whether the mount point is read-only.
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration struct {
	// Indicates whether parameter transformation and result enhancement is enabled.
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

func (s ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration struct {
	// Indicates whether the MCP proxy is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) SetEnabled(v bool) *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration struct {
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
	// The Alibaba Cloud Resource Name (ARN) of the RAM role used by user code to access downstream Alibaba Cloud resources.
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

func (s ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetPort(v int32) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *ListMcpsResponseBodyItemsDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsMarketSource struct {
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s ListMcpsResponseBodyItemsMarketSource) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsMarketSource) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *ListMcpsResponseBodyItemsMarketSource) SetMarketItemId(v string) *ListMcpsResponseBodyItemsMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *ListMcpsResponseBodyItemsMarketSource) Validate() error {
	return dara.Validate(s)
}

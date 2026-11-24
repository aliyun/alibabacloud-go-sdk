// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpResponseBody
	GetCode() *string
	SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody
	GetData() *GetMcpResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpResponseBody
	GetSuccess() *bool
}

type GetMcpResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetMcpResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpResponseBody) GetData() *GetMcpResponseBodyData {
	return s.Data
}

func (s *GetMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpResponseBody) SetCode(v string) *GetMcpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpResponseBody) SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpResponseBody) SetHttpStatusCode(v int32) *GetMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpResponseBody) SetMessage(v string) *GetMcpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpResponseBody) SetRequestId(v string) *GetMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpResponseBody) SetSuccess(v bool) *GetMcpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyData struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. The enabled field indicates whether it is enabled. The directProxy field specifies the custom authentication header for the direct proxy. The httpToMcp field specifies the list of OpenAPI credentials for HTTP_TO_MCP.
	Auth *GetMcpResponseBodyDataAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The saved legacy or modern backend policy, which is separate from the transport indicated by the protocol and does not indicate verification. This field is not saved for new CODE_PACKAGE deployments. For historical CODE_PACKAGE deployments, the saved value does not control Function Compute session affinity or the gateway protocol. For the actual protocol, refer to OriginProbe.LastResult.ActualConfiguration.BackendMcpProtocol. Returns null for historical deployments with unspecified values and HTTP_TO_MCP deployments.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// The time when the MCP was created, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-23T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The custom tags. Multiple tags are supported.
	CustomTags []*string `json:"customTags,omitempty" xml:"customTags,omitempty" type:"Repeated"`
	// The deployment configuration for code-deployed MCP.
	DeploymentConfig *GetMcpResponseBodyDataDeploymentConfig `json:"deploymentConfig,omitempty" xml:"deploymentConfig,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// Sample description to illustrate the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP endpoint that can be called by users or agents. This field is empty before the deployment is completed.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Function Compute function name for the CODE_PACKAGE MCP. This field is empty before deployment is complete and for other types.
	//
	// example:
	//
	// agentcore-mcp-example
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The MCP source marketplace template.
	MarketSource *GetMcpResponseBodyDataMarketSource `json:"marketSource,omitempty" xml:"marketSource,omitempty" type:"Struct"`
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
	// The official usage tag, which is managed by the server.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// The read-only progress that reuses existing release tasks. It does not change the primary status of the resource. This field is empty if no historical operation information is available.
	Operation *GetMcpResponseBodyDataOperation `json:"operation,omitempty" xml:"operation,omitempty" type:"Struct"`
	// Returns the minimal origin verification summary of the most recent task only for code deployments. This field is empty for direct connection and HTTP-to-MCP deployments. Querying this field does not trigger probing.
	OriginProbe *GetMcpResponseBodyDataOriginProbe `json:"originProbe,omitempty" xml:"originProbe,omitempty" type:"Struct"`
	// The optional saved metadata for the origin MCP protocol date display, in YYYY-MM-DD format. You can specify this field during creation and modify it during updates. If omitted during an update, the current value is retained. This field is used only for saving and displaying information. It is not a software version, does not restrict the external gateway protocol, does not control backend requests, and does not indicate that real-time verification has passed. Returns null if omitted during creation and not updated, missing in historical data, or used in HTTP_TO_MCP deployments. No default value.
	//
	// example:
	//
	// 2025-03-26
	OriginProtocolVersion *string `json:"originProtocolVersion,omitempty" xml:"originProtocolVersion,omitempty"`
	// The access permission mode. Valid values:
	//
	// - PUBLIC: Public access.
	//
	// - CONTROLLED: Access controlled based on permission rules.
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
	// The reason for the status.
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
	// The template schema for the usage-bound MCP version. This field is not returned for regular MCPs.
	Template *GetMcpResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
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
	// The access address of the MCP service.
	//
	// example:
	//
	// https://example.com/artifacts/example.zip
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// Specifies whether the MCP is still subject to the usage constraints of the official template.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s GetMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *GetMcpResponseBodyData) GetAuth() *GetMcpResponseBodyDataAuth {
	return s.Auth
}

func (s *GetMcpResponseBodyData) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *GetMcpResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMcpResponseBodyData) GetCustomTags() []*string {
	return s.CustomTags
}

func (s *GetMcpResponseBodyData) GetDeploymentConfig() *GetMcpResponseBodyDataDeploymentConfig {
	return s.DeploymentConfig
}

func (s *GetMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMcpResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetMcpResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetMcpResponseBodyData) GetMarketSource() *GetMcpResponseBodyDataMarketSource {
	return s.MarketSource
}

func (s *GetMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *GetMcpResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyData) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *GetMcpResponseBodyData) GetOperation() *GetMcpResponseBodyDataOperation {
	return s.Operation
}

func (s *GetMcpResponseBodyData) GetOriginProbe() *GetMcpResponseBodyDataOriginProbe {
	return s.OriginProbe
}

func (s *GetMcpResponseBodyData) GetOriginProtocolVersion() *string {
	return s.OriginProtocolVersion
}

func (s *GetMcpResponseBodyData) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *GetMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMcpResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetMcpResponseBodyData) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *GetMcpResponseBodyData) GetTemplate() *GetMcpResponseBodyDataTemplate {
	return s.Template
}

func (s *GetMcpResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMcpResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMcpResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetMcpResponseBodyData) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *GetMcpResponseBodyData) SetAddresses(v []*string) *GetMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *GetMcpResponseBodyData) SetAuth(v *GetMcpResponseBodyDataAuth) *GetMcpResponseBodyData {
	s.Auth = v
	return s
}

func (s *GetMcpResponseBodyData) SetBackendMcpProtocol(v string) *GetMcpResponseBodyData {
	s.BackendMcpProtocol = &v
	return s
}

func (s *GetMcpResponseBodyData) SetCreatedAt(v string) *GetMcpResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetMcpResponseBodyData) SetCustomTags(v []*string) *GetMcpResponseBodyData {
	s.CustomTags = v
	return s
}

func (s *GetMcpResponseBodyData) SetDeploymentConfig(v *GetMcpResponseBodyDataDeploymentConfig) *GetMcpResponseBodyData {
	s.DeploymentConfig = v
	return s
}

func (s *GetMcpResponseBodyData) SetDescription(v string) *GetMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMcpResponseBodyData) SetEndpoint(v string) *GetMcpResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetMcpResponseBodyData) SetFunctionName(v string) *GetMcpResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *GetMcpResponseBodyData) SetMarketSource(v *GetMcpResponseBodyDataMarketSource) *GetMcpResponseBodyData {
	s.MarketSource = v
	return s
}

func (s *GetMcpResponseBodyData) SetMcpServerConfig(v string) *GetMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetMcpServerId(v string) *GetMcpResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *GetMcpResponseBodyData) SetName(v string) *GetMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyData) SetOfficialTag(v string) *GetMcpResponseBodyData {
	s.OfficialTag = &v
	return s
}

func (s *GetMcpResponseBodyData) SetOperation(v *GetMcpResponseBodyDataOperation) *GetMcpResponseBodyData {
	s.Operation = v
	return s
}

func (s *GetMcpResponseBodyData) SetOriginProbe(v *GetMcpResponseBodyDataOriginProbe) *GetMcpResponseBodyData {
	s.OriginProbe = v
	return s
}

func (s *GetMcpResponseBodyData) SetOriginProtocolVersion(v string) *GetMcpResponseBodyData {
	s.OriginProtocolVersion = &v
	return s
}

func (s *GetMcpResponseBodyData) SetPermissionMode(v string) *GetMcpResponseBodyData {
	s.PermissionMode = &v
	return s
}

func (s *GetMcpResponseBodyData) SetProtocol(v string) *GetMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetMcpResponseBodyData) SetStatus(v string) *GetMcpResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMcpResponseBodyData) SetStatusReason(v string) *GetMcpResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetMcpResponseBodyData) SetSwaggerConfig(v string) *GetMcpResponseBodyData {
	s.SwaggerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetTemplate(v *GetMcpResponseBodyDataTemplate) *GetMcpResponseBodyData {
	s.Template = v
	return s
}

func (s *GetMcpResponseBodyData) SetType(v string) *GetMcpResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMcpResponseBodyData) SetUpdatedAt(v string) *GetMcpResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetMcpResponseBodyData) SetUrl(v string) *GetMcpResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetMcpResponseBodyData) SetUsageActive(v bool) *GetMcpResponseBodyData {
	s.UsageActive = &v
	return s
}

func (s *GetMcpResponseBodyData) Validate() error {
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
	if s.Operation != nil {
		if err := s.Operation.Validate(); err != nil {
			return err
		}
	}
	if s.OriginProbe != nil {
		if err := s.OriginProbe.Validate(); err != nil {
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

type GetMcpResponseBodyDataAuth struct {
	// The caller API key authentication configuration for the code deployment MCP.
	CodePackage *GetMcpResponseBodyDataAuthCodePackage `json:"codePackage,omitempty" xml:"codePackage,omitempty" type:"Struct"`
	// The authentication configuration for the direct proxy.
	DirectProxy *GetMcpResponseBodyDataAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Indicates whether the configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of authentication configurations for HTTP_TO_MCP.
	HttpToMcp []*GetMcpResponseBodyDataAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s GetMcpResponseBodyDataAuth) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuth) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuth) GetCodePackage() *GetMcpResponseBodyDataAuthCodePackage {
	return s.CodePackage
}

func (s *GetMcpResponseBodyDataAuth) GetDirectProxy() *GetMcpResponseBodyDataAuthDirectProxy {
	return s.DirectProxy
}

func (s *GetMcpResponseBodyDataAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataAuth) GetHttpToMcp() []*GetMcpResponseBodyDataAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *GetMcpResponseBodyDataAuth) SetCodePackage(v *GetMcpResponseBodyDataAuthCodePackage) *GetMcpResponseBodyDataAuth {
	s.CodePackage = v
	return s
}

func (s *GetMcpResponseBodyDataAuth) SetDirectProxy(v *GetMcpResponseBodyDataAuthDirectProxy) *GetMcpResponseBodyDataAuth {
	s.DirectProxy = v
	return s
}

func (s *GetMcpResponseBodyDataAuth) SetEnabled(v bool) *GetMcpResponseBodyDataAuth {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataAuth) SetHttpToMcp(v []*GetMcpResponseBodyDataAuthHttpToMcp) *GetMcpResponseBodyDataAuth {
	s.HttpToMcp = v
	return s
}

func (s *GetMcpResponseBodyDataAuth) Validate() error {
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

type GetMcpResponseBodyDataAuthCodePackage struct {
	// The API key used to verify the MCP caller.
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

func (s GetMcpResponseBodyDataAuthCodePackage) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuthCodePackage) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuthCodePackage) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetMcpResponseBodyDataAuthCodePackage) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetMcpResponseBodyDataAuthCodePackage) SetApiKey(v string) *GetMcpResponseBodyDataAuthCodePackage {
	s.ApiKey = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthCodePackage) SetHeaderName(v string) *GetMcpResponseBodyDataAuthCodePackage {
	s.HeaderName = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthCodePackage) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataAuthDirectProxy struct {
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

func (s GetMcpResponseBodyDataAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) SetName(v string) *GetMcpResponseBodyDataAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) SetValue(v string) *GetMcpResponseBodyDataAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataAuthHttpToMcp struct {
	// The authentication credential.
	//
	// example:
	//
	// example-credential
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The ID of the authentication scheme.
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

func (s GetMcpResponseBodyDataAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetCredential(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetId(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetName(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetPosition(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetType(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfig struct {
	// The MCP entry access control configuration.
	AccessControl *GetMcpResponseBodyDataDeploymentConfigAccessControl `json:"accessControl,omitempty" xml:"accessControl,omitempty" type:"Struct"`
	// The agent identity configuration.
	AgentIdentityConfiguration *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration `json:"agentIdentityConfiguration,omitempty" xml:"agentIdentityConfiguration,omitempty" type:"Struct"`
	// The artifact type. Valid values:
	//
	// - Code: ZIP code package.
	//
	// - Container: Custom container.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code package configuration.
	CodeConfiguration *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty" type:"Struct"`
	// The custom container configuration.
	ContainerConfiguration *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty" type:"Struct"`
	// The hook configuration.
	HookConfiguration *GetMcpResponseBodyDataDeploymentConfigHookConfiguration `json:"hookConfiguration,omitempty" xml:"hookConfiguration,omitempty" type:"Struct"`
	// The log configuration.
	LogConfiguration *GetMcpResponseBodyDataDeploymentConfigLogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty" type:"Struct"`
	// The MCP session configuration.
	McpConfiguration *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration `json:"mcpConfiguration,omitempty" xml:"mcpConfiguration,omitempty" type:"Struct"`
	// The NAS storage configuration.
	NasConfiguration *GetMcpResponseBodyDataDeploymentConfigNasConfiguration `json:"nasConfiguration,omitempty" xml:"nasConfiguration,omitempty" type:"Struct"`
	// The network configuration.
	NetworkConfiguration *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The OSS mount configuration.
	OssMountConfiguration *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration `json:"ossMountConfiguration,omitempty" xml:"ossMountConfiguration,omitempty" type:"Struct"`
	// The parameter transformation and result enhancement configuration.
	ParameterTransformConfiguration *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration `json:"parameterTransformConfiguration,omitempty" xml:"parameterTransformConfiguration,omitempty" type:"Struct"`
	// The MCP proxy configuration.
	ProxyConfiguration *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration `json:"proxyConfiguration,omitempty" xml:"proxyConfiguration,omitempty" type:"Struct"`
	// The runtime and resource configuration.
	RuntimeConfiguration *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration `json:"runtimeConfiguration,omitempty" xml:"runtimeConfiguration,omitempty" type:"Struct"`
}

func (s GetMcpResponseBodyDataDeploymentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfig) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetAccessControl() *GetMcpResponseBodyDataDeploymentConfigAccessControl {
	return s.AccessControl
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetAgentIdentityConfiguration() *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	return s.AgentIdentityConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetCodeConfiguration() *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	return s.CodeConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetContainerConfiguration() *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetHookConfiguration() *GetMcpResponseBodyDataDeploymentConfigHookConfiguration {
	return s.HookConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetLogConfiguration() *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	return s.LogConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetMcpConfiguration() *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	return s.McpConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetNasConfiguration() *GetMcpResponseBodyDataDeploymentConfigNasConfiguration {
	return s.NasConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetNetworkConfiguration() *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetOssMountConfiguration() *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration {
	return s.OssMountConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetParameterTransformConfiguration() *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	return s.ParameterTransformConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetProxyConfiguration() *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration {
	return s.ProxyConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) GetRuntimeConfiguration() *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	return s.RuntimeConfiguration
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetAccessControl(v *GetMcpResponseBodyDataDeploymentConfigAccessControl) *GetMcpResponseBodyDataDeploymentConfig {
	s.AccessControl = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetAgentIdentityConfiguration(v *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.AgentIdentityConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetArtifactType(v string) *GetMcpResponseBodyDataDeploymentConfig {
	s.ArtifactType = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetCodeConfiguration(v *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.CodeConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetContainerConfiguration(v *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.ContainerConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetHookConfiguration(v *GetMcpResponseBodyDataDeploymentConfigHookConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.HookConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetLogConfiguration(v *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.LogConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetMcpConfiguration(v *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.McpConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetNasConfiguration(v *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.NasConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetNetworkConfiguration(v *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.NetworkConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetOssMountConfiguration(v *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.OssMountConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetParameterTransformConfiguration(v *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.ParameterTransformConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetProxyConfiguration(v *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.ProxyConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) SetRuntimeConfiguration(v *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) *GetMcpResponseBodyDataDeploymentConfig {
	s.RuntimeConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfig) Validate() error {
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

type GetMcpResponseBodyDataDeploymentConfigAccessControl struct {
	// The ID of the AgentCore credential referenced when the mode is set to CREDENTIAL.
	//
	// example:
	//
	// credential-id
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Specifies whether to enable entry access control.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access control mode. Valid values: ANONYMOUS indicates anonymous access. CREDENTIAL indicates that AgentCore access credentials are used.
	//
	// example:
	//
	// CREDENTIAL
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigAccessControl) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigAccessControl) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) GetMode() *string {
	return s.Mode
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) SetCredentialId(v string) *GetMcpResponseBodyDataDeploymentConfigAccessControl {
	s.CredentialId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) SetEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigAccessControl {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) SetMode(v string) *GetMcpResponseBodyDataDeploymentConfigAccessControl {
	s.Mode = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAccessControl) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration struct {
	// Specifies whether authorization is enabled.
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
	// Specifies whether agent identity is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetAuthorizationEnabled() *bool {
	return s.AuthorizationEnabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetCredentialProviderType() *string {
	return s.CredentialProviderType
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetAuthorizationEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.AuthorizationEnabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderArn(v string) *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderArn = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetCredentialProviderType(v string) *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.CredentialProviderType = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) SetEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigAgentIdentityConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigCodeConfiguration struct {
	// The temporary code package token returned by GetMcpCodePackageUploadUrl. It is used to create or update a code deployment after a presigned upload is completed.
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

func (s GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetCodePackageToken(v string) *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.CodePackageToken = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetCommand(v []*string) *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Command = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) SetLanguage(v string) *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration {
	s.Language = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigCodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigContainerConfiguration struct {
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
	// The MCP runtime mode. Custom containers must expose the standard MCP themselves. The value is SELF_HOSTED.
	//
	// example:
	//
	// SELF_HOSTED
	McpRuntimeMode *string `json:"mcpRuntimeMode,omitempty" xml:"mcpRuntimeMode,omitempty"`
	// The container source type. This value is currently fixed to CONTAINER_IMAGE.
	//
	// example:
	//
	// CONTAINER_IMAGE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetMcpRuntimeMode() *string {
	return s.McpRuntimeMode
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) GetSourceType() *string {
	return s.SourceType
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetAcrInstanceId(v string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetCommand(v []*string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Command = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetEntrypoint(v []*string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Entrypoint = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetImage(v string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.Image = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetImageRegistryType(v string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetMcpRuntimeMode(v string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.McpRuntimeMode = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) SetSourceType(v string) *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration {
	s.SourceType = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigContainerConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigHookConfiguration struct {
	// The ordered list of hooks. The PRE_LIST_TOOLS, PRE_CALL_TOOL, POST_LIST_TOOLS, and POST_CALL_TOOL hooks are executed in the order of the array.
	Hooks []*GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks `json:"hooks,omitempty" xml:"hooks,omitempty" type:"Repeated"`
}

func (s GetMcpResponseBodyDataDeploymentConfigHookConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigHookConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfiguration) GetHooks() []*GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	return s.Hooks
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfiguration) SetHooks(v []*GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) *GetMcpResponseBodyDataDeploymentConfigHookConfiguration {
	s.Hooks = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfiguration) Validate() error {
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

type GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks struct {
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
	// The request headers of the hook.
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

func (s GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetDescription() *string {
	return s.Description
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetEvent() *string {
	return s.Event
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) GetUrl() *string {
	return s.Url
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetApiVersion(v string) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.ApiVersion = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetDescription(v string) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Description = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetEvent(v string) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Event = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetHeaders(v map[string]*string) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Headers = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetTimeout(v int32) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Timeout = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) SetUrl(v string) *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks {
	s.Url = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigHookConfigurationHooks) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigLogConfiguration struct {
	// Specifies whether to collect instance metrics.
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to collect request metrics.
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The start rule for splitting Function Compute logs.
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

func (s GetMcpResponseBodyDataDeploymentConfigLogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) GetProject() *string {
	return s.Project
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) SetEnableInstanceMetrics(v bool) *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) SetEnableRequestMetrics(v bool) *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.EnableRequestMetrics = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) SetLogBeginRule(v string) *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.LogBeginRule = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) SetLogstore(v string) *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.Logstore = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) SetProject(v string) *GetMcpResponseBodyDataDeploymentConfigLogConfiguration {
	s.Project = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigLogConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigMcpConfiguration struct {
	// The MCP endpoint path, such as /mcp or /sse.
	//
	// example:
	//
	// /mcp
	EndpointPath *string `json:"endpointPath,omitempty" xml:"endpointPath,omitempty"`
	// The session concurrency per instance. This value is currently fixed to 1.
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
	// The session mode. This parameter applies only to StreamableHTTP. Default value for creation: STATEFUL. If omitted during an update, the original value is retained. STATELESS disables Function Compute session affinity, and the stdio gateway does not enable stateful mode. Native HTTP services must support stateless MCP themselves.
	//
	// example:
	//
	// STATEFUL
	SessionMode *string `json:"sessionMode,omitempty" xml:"sessionMode,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetEndpointPath() *string {
	return s.EndpointPath
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionConcurrencyPerInstance() *int32 {
	return s.SessionConcurrencyPerInstance
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMaxLifetimeSeconds() *int32 {
	return s.SessionMaxLifetimeSeconds
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) GetSessionMode() *string {
	return s.SessionMode
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetEndpointPath(v string) *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.EndpointPath = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionConcurrencyPerInstance(v int32) *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionIdleTimeoutSeconds(v int32) *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMaxLifetimeSeconds(v int32) *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMaxLifetimeSeconds = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) SetSessionMode(v string) *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration {
	s.SessionMode = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigMcpConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigNasConfiguration struct {
	// The user group ID.
	//
	// example:
	//
	// 1000
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The list of NAS mount points.
	MountPoints []*GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 1000
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigNasConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigNasConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) GetGroupId() *int32 {
	return s.GroupId
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) GetMountPoints() []*GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	return s.MountPoints
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) GetUserId() *int32 {
	return s.UserId
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) SetGroupId(v int32) *GetMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.GroupId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) SetMountPoints(v []*GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) *GetMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.MountPoints = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) SetUserId(v int32) *GetMcpResponseBodyDataDeploymentConfigNasConfiguration {
	s.UserId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfiguration) Validate() error {
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

type GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints struct {
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

func (s GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetEnableTls(v bool) *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.EnableTls = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetMountDir(v string) *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) SetServerAddr(v string) *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNasConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration struct {
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

func (s GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetNetworkMode(v string) *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetSecurityGroupId(v string) *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetVSwitchIds(v []*string) *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VSwitchIds = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) SetVpcId(v string) *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration struct {
	// The list of OSS mount points.
	MountPoints []*GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) GetMountPoints() []*GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	return s.MountPoints
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) SetMountPoints(v []*GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration {
	s.MountPoints = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfiguration) Validate() error {
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

type GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The path of the OSS bucket.
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

func (s GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketName(v string) *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketName = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetBucketPath(v string) *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.BucketPath = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetEndpoint(v string) *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.Endpoint = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetMountDir(v string) *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.MountDir = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) SetReadOnly(v bool) *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigOssMountConfigurationMountPoints) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration struct {
	// Specifies whether parameter transformation and result enhancement are enabled.
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

func (s GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetRuleSetId() *string {
	return s.RuleSetId
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) GetVersion() *string {
	return s.Version
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetRuleSetId(v string) *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.RuleSetId = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) SetVersion(v string) *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration {
	s.Version = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigParameterTransformConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigProxyConfiguration struct {
	// Specifies whether to enable the MCP proxy.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) SetEnabled(v bool) *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigProxyConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration struct {
	// The CPU specifications. Unit: cores. Default value: 0.25.
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
	// The request concurrency per instance. Default value: 200.
	//
	// example:
	//
	// 200
	InstanceConcurrency *int32 `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	// The memory specifications. Unit: MB. Default value: 512.
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

func (s GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetCpu() *float64 {
	return s.Cpu
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetMemory() *int32 {
	return s.Memory
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetCpu(v float64) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Cpu = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetDiskSize(v int32) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.DiskSize = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetEnvironmentVariables(v map[string]*string) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.EnvironmentVariables = v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetExecutionRoleArn(v string) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.ExecutionRoleArn = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetInstanceConcurrency(v int32) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.InstanceConcurrency = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetMemory(v int32) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Memory = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetPort(v int32) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Port = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) SetTimeout(v int32) *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration {
	s.Timeout = &v
	return s
}

func (s *GetMcpResponseBodyDataDeploymentConfigRuntimeConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataMarketSource struct {
	// The MCP marketplace template ID.
	//
	// example:
	//
	// market-1
	MarketItemId *string `json:"marketItemId,omitempty" xml:"marketItemId,omitempty"`
}

func (s GetMcpResponseBodyDataMarketSource) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataMarketSource) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataMarketSource) GetMarketItemId() *string {
	return s.MarketItemId
}

func (s *GetMcpResponseBodyDataMarketSource) SetMarketItemId(v string) *GetMcpResponseBodyDataMarketSource {
	s.MarketItemId = &v
	return s
}

func (s *GetMcpResponseBodyDataMarketSource) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataOperation struct {
	// The desensitized error of the current release operation. Gateway or formal link failures after the origin server passes verification are represented by this field.
	Error *GetMcpResponseBodyDataOperationError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The identifier of the current or most recent release operation.
	//
	// example:
	//
	// 3541f1af-87cc-45d1-a28e-94ec5633b931
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The release stage, such as compute deployment, origin verification, summary saving, gateway release, entry credential synchronization, formal link verification, or deployment submission. Old tasks retain their original stage semantics.
	//
	// example:
	//
	// VALIDATE_ORIGIN
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// The operation status. Valid values: PENDING, RUNNING, SUCCEEDED, and FAILED.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The operation type. Valid value: PUBLISH.
	//
	// example:
	//
	// PUBLISH
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the operation was updated, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-11-19T08:30:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetMcpResponseBodyDataOperation) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOperation) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOperation) GetError() *GetMcpResponseBodyDataOperationError {
	return s.Error
}

func (s *GetMcpResponseBodyDataOperation) GetOperationId() *string {
	return s.OperationId
}

func (s *GetMcpResponseBodyDataOperation) GetStage() *string {
	return s.Stage
}

func (s *GetMcpResponseBodyDataOperation) GetStatus() *string {
	return s.Status
}

func (s *GetMcpResponseBodyDataOperation) GetType() *string {
	return s.Type
}

func (s *GetMcpResponseBodyDataOperation) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMcpResponseBodyDataOperation) SetError(v *GetMcpResponseBodyDataOperationError) *GetMcpResponseBodyDataOperation {
	s.Error = v
	return s
}

func (s *GetMcpResponseBodyDataOperation) SetOperationId(v string) *GetMcpResponseBodyDataOperation {
	s.OperationId = &v
	return s
}

func (s *GetMcpResponseBodyDataOperation) SetStage(v string) *GetMcpResponseBodyDataOperation {
	s.Stage = &v
	return s
}

func (s *GetMcpResponseBodyDataOperation) SetStatus(v string) *GetMcpResponseBodyDataOperation {
	s.Status = &v
	return s
}

func (s *GetMcpResponseBodyDataOperation) SetType(v string) *GetMcpResponseBodyDataOperation {
	s.Type = &v
	return s
}

func (s *GetMcpResponseBodyDataOperation) SetUpdatedAt(v string) *GetMcpResponseBodyDataOperation {
	s.UpdatedAt = &v
	return s
}

func (s *GetMcpResponseBodyDataOperation) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyDataOperationError struct {
	// Returns only bounded diagnostic codes and does not return the original error body from the origin server.
	//
	// example:
	//
	// PROBE_TIMEOUT
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The necessary diagnostics that are trimmed and desensitized. It does not contain credentials or original responses.
	//
	// example:
	//
	// The probe deadline was exceeded.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the operation is suitable for retry based on a restricted policy. This field does not trigger automatic retries.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// The source of the error. Valid values: MAGIC, ROUTER, NETWORK, ORIGIN, and UNKNOWN. The value is determined only based on actual evidence.
	//
	// example:
	//
	// UNKNOWN
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The detection or release stage where the error occurred.
	//
	// example:
	//
	// INITIALIZE
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s GetMcpResponseBodyDataOperationError) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOperationError) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOperationError) GetCode() *string {
	return s.Code
}

func (s *GetMcpResponseBodyDataOperationError) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResponseBodyDataOperationError) GetRetryable() *bool {
	return s.Retryable
}

func (s *GetMcpResponseBodyDataOperationError) GetSource() *string {
	return s.Source
}

func (s *GetMcpResponseBodyDataOperationError) GetStage() *string {
	return s.Stage
}

func (s *GetMcpResponseBodyDataOperationError) SetCode(v string) *GetMcpResponseBodyDataOperationError {
	s.Code = &v
	return s
}

func (s *GetMcpResponseBodyDataOperationError) SetMessage(v string) *GetMcpResponseBodyDataOperationError {
	s.Message = &v
	return s
}

func (s *GetMcpResponseBodyDataOperationError) SetRetryable(v bool) *GetMcpResponseBodyDataOperationError {
	s.Retryable = &v
	return s
}

func (s *GetMcpResponseBodyDataOperationError) SetSource(v string) *GetMcpResponseBodyDataOperationError {
	s.Source = &v
	return s
}

func (s *GetMcpResponseBodyDataOperationError) SetStage(v string) *GetMcpResponseBodyDataOperationError {
	s.Stage = &v
	return s
}

func (s *GetMcpResponseBodyDataOperationError) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataOriginProbe struct {
	// Returned only for CODE_PACKAGE deployments. This field is cleared when a new task starts and is empty if no verification has been performed or if an old task did not perform verification. If a subsequent release fails, the current summary is retained, which does not indicate real-time health or the active version status.
	LastResult *GetMcpResponseBodyDataOriginProbeLastResult `json:"lastResult,omitempty" xml:"lastResult,omitempty" type:"Struct"`
}

func (s GetMcpResponseBodyDataOriginProbe) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOriginProbe) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOriginProbe) GetLastResult() *GetMcpResponseBodyDataOriginProbeLastResult {
	return s.LastResult
}

func (s *GetMcpResponseBodyDataOriginProbe) SetLastResult(v *GetMcpResponseBodyDataOriginProbeLastResult) *GetMcpResponseBodyDataOriginProbe {
	s.LastResult = v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbe) Validate() error {
	if s.LastResult != nil {
		if err := s.LastResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyDataOriginProbeLastResult struct {
	// The actual configuration confirmed by evidence in the current verification. This field is empty if no confirmed information is available and is not inferred from the expected configuration.
	ActualConfiguration *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration `json:"actualConfiguration,omitempty" xml:"actualConfiguration,omitempty" type:"Struct"`
	// The verification error. This field is empty if the verification is successful. It saves only the desensitized diagnostics required for the current verification.
	Error *GetMcpResponseBodyDataOriginProbeLastResultError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The time of the current origin verification, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-11-19T08:30:00Z
	ProbedAt *string `json:"probedAt,omitempty" xml:"probedAt,omitempty"`
	// The verification stage. Valid values: CONNECT, DISCOVER, INITIALIZE, SESSION_CHECK, CLEANUP, and COMPLETE.
	//
	// example:
	//
	// COMPLETE
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// The verification result. Valid values: SUCCESS and FAILED. The result is SUCCESS if the protocol and applicable session checks pass. Not declaring any capabilities does not affect the passing status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMcpResponseBodyDataOriginProbeLastResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOriginProbeLastResult) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) GetActualConfiguration() *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration {
	return s.ActualConfiguration
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) GetError() *GetMcpResponseBodyDataOriginProbeLastResultError {
	return s.Error
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) GetProbedAt() *string {
	return s.ProbedAt
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) GetStage() *string {
	return s.Stage
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) GetStatus() *string {
	return s.Status
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) SetActualConfiguration(v *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) *GetMcpResponseBodyDataOriginProbeLastResult {
	s.ActualConfiguration = v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) SetError(v *GetMcpResponseBodyDataOriginProbeLastResultError) *GetMcpResponseBodyDataOriginProbeLastResult {
	s.Error = v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) SetProbedAt(v string) *GetMcpResponseBodyDataOriginProbeLastResult {
	s.ProbedAt = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) SetStage(v string) *GetMcpResponseBodyDataOriginProbeLastResult {
	s.Stage = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) SetStatus(v string) *GetMcpResponseBodyDataOriginProbeLastResult {
	s.Status = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResult) Validate() error {
	if s.ActualConfiguration != nil {
		if err := s.ActualConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration struct {
	// The verified legacy or modern policy. This field is empty if the policy is not confirmed.
	//
	// example:
	//
	// modern
	BackendMcpProtocol *string `json:"backendMcpProtocol,omitempty" xml:"backendMcpProtocol,omitempty"`
	// The verified SSE or StreamableHTTP transport. This field is empty if the transport is not confirmed.
	//
	// example:
	//
	// StreamableHTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The actual negotiated or selected and verified date version. This field is empty if the version is not confirmed.
	//
	// example:
	//
	// 2025-03-26
	ProtocolVersion *string `json:"protocolVersion,omitempty" xml:"protocolVersion,omitempty"`
	// The actual session mode. Valid values: STATEFUL, STATELESS, and NOT_APPLICABLE. This field is empty if the mode is not confirmed.
	//
	// example:
	//
	// STATELESS
	SessionMode *string `json:"sessionMode,omitempty" xml:"sessionMode,omitempty"`
}

func (s GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) GetBackendMcpProtocol() *string {
	return s.BackendMcpProtocol
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) GetSessionMode() *string {
	return s.SessionMode
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) SetBackendMcpProtocol(v string) *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration {
	s.BackendMcpProtocol = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) SetProtocol(v string) *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration {
	s.Protocol = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) SetProtocolVersion(v string) *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration {
	s.ProtocolVersion = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) SetSessionMode(v string) *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration {
	s.SessionMode = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultActualConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataOriginProbeLastResultError struct {
	// Returns only bounded diagnostic codes and does not return the original error body from the origin server.
	//
	// example:
	//
	// PROBE_TIMEOUT
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The necessary diagnostics that are trimmed and desensitized. It does not contain credentials or original responses.
	//
	// example:
	//
	// The probe deadline was exceeded.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the operation is suitable for retry based on a restricted policy. This field does not trigger automatic retries.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// The source of the error. Valid values: MAGIC, ROUTER, NETWORK, ORIGIN, and UNKNOWN. The value is determined only based on actual evidence.
	//
	// example:
	//
	// UNKNOWN
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The detection or release stage where the error occurred.
	//
	// example:
	//
	// INITIALIZE
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s GetMcpResponseBodyDataOriginProbeLastResultError) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataOriginProbeLastResultError) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) GetCode() *string {
	return s.Code
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) GetRetryable() *bool {
	return s.Retryable
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) GetSource() *string {
	return s.Source
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) GetStage() *string {
	return s.Stage
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) SetCode(v string) *GetMcpResponseBodyDataOriginProbeLastResultError {
	s.Code = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) SetMessage(v string) *GetMcpResponseBodyDataOriginProbeLastResultError {
	s.Message = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) SetRetryable(v bool) *GetMcpResponseBodyDataOriginProbeLastResultError {
	s.Retryable = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) SetSource(v string) *GetMcpResponseBodyDataOriginProbeLastResultError {
	s.Source = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) SetStage(v string) *GetMcpResponseBodyDataOriginProbeLastResultError {
	s.Stage = &v
	return s
}

func (s *GetMcpResponseBodyDataOriginProbeLastResultError) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataTemplate struct {
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
	// Specifies whether an updated template version is available.
	UpdateAvailable *bool `json:"updateAvailable,omitempty" xml:"updateAvailable,omitempty"`
}

func (s GetMcpResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataTemplate) GetAppliedTemplateVersion() *string {
	return s.AppliedTemplateVersion
}

func (s *GetMcpResponseBodyDataTemplate) GetLatestTemplateVersion() *string {
	return s.LatestTemplateVersion
}

func (s *GetMcpResponseBodyDataTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetMcpResponseBodyDataTemplate) GetTemplateInputSchema() *string {
	return s.TemplateInputSchema
}

func (s *GetMcpResponseBodyDataTemplate) GetUpdateAvailable() *bool {
	return s.UpdateAvailable
}

func (s *GetMcpResponseBodyDataTemplate) SetAppliedTemplateVersion(v string) *GetMcpResponseBodyDataTemplate {
	s.AppliedTemplateVersion = &v
	return s
}

func (s *GetMcpResponseBodyDataTemplate) SetLatestTemplateVersion(v string) *GetMcpResponseBodyDataTemplate {
	s.LatestTemplateVersion = &v
	return s
}

func (s *GetMcpResponseBodyDataTemplate) SetSchemaVersion(v string) *GetMcpResponseBodyDataTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *GetMcpResponseBodyDataTemplate) SetTemplateInputSchema(v string) *GetMcpResponseBodyDataTemplate {
	s.TemplateInputSchema = &v
	return s
}

func (s *GetMcpResponseBodyDataTemplate) SetUpdateAvailable(v bool) *GetMcpResponseBodyDataTemplate {
	s.UpdateAvailable = &v
	return s
}

func (s *GetMcpResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}

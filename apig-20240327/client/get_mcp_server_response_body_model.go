// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpServerResponseBody
	GetCode() *string
	SetData(v *GetMcpServerResponseBodyData) *GetMcpServerResponseBody
	GetData() *GetMcpServerResponseBodyData
	SetMessage(v string) *GetMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpServerResponseBody
	GetRequestId() *string
}

type GetMcpServerResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetMcpServerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2A6E90D5-A711-54F4-A489-E33C2021EDDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpServerResponseBody) GetData() *GetMcpServerResponseBodyData {
	return s.Data
}

func (s *GetMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpServerResponseBody) SetCode(v string) *GetMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpServerResponseBody) SetData(v *GetMcpServerResponseBodyData) *GetMcpServerResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpServerResponseBody) SetMessage(v string) *GetMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpServerResponseBody) SetRequestId(v string) *GetMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpServerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServerResponseBodyData struct {
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// A list of assembly sources. This parameter is required when `type` is `AssemblyMCP`.
	AssembledSources []*GetMcpServerResponseBodyDataAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	// The backend service of the route.
	Backend *Backend `json:"backend,omitempty" xml:"backend,omitempty"`
	// Specifies how the MCP Server was created. Valid values:
	//
	// `ApiGatewayHttpToMCP`: The API gateway converts HTTP requests to the MCP protocol. `ApiGatewayMcpHosting`: The API gateway acts as a direct proxy for MCP services. `ApiGatewayAssembly`: The API gateway assembles responses from multiple MCP services. `NacosHttpToMCP`: The API gateway converts HTTP requests to the MCP protocol, with service discovery provided by Nacos. `NacosMcpHosting`: The API gateway acts as a direct proxy for MCP services, with service discovery provided by Nacos.
	//
	// example:
	//
	// ApiGatewayHttpToMCP
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// The API\\"s deployment status in the current environment.
	//
	// example:
	//
	// Deployed
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// The description of the MCP Server.
	//
	// example:
	//
	// 测试专用MCP Server
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A list of domain name IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// A list of domain information.
	DomainInfos []*GetMcpServerResponseBodyDataDomainInfos `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	// The ID of the environment.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The exposed URI path. This parameter is required when `protocol` is `SSE` or `StreamableHTTP` and `type` is `RealMCP`.
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// A list of canary route configurations.
	GrayMcpServerConfigs []*GetMcpServerResponseBodyDataGrayMcpServerConfigs `json:"grayMcpServerConfigs,omitempty" xml:"grayMcpServerConfigs,omitempty" type:"Repeated"`
	// The match rule for the route.
	//
	// example:
	//
	// {\\"product_code\\":\\"apigw\\"}
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The configuration for converting HTTP requests to the MCP protocol.
	//
	// example:
	//
	// MCP Server配置 base64
	McpServerConfig *string `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
	// The attachment ID of the MCP Server plugin configuration.
	//
	// example:
	//
	// pa-adfaefwaef
	McpServerConfigPluginAttachmentId *string `json:"mcpServerConfigPluginAttachmentId,omitempty" xml:"mcpServerConfigPluginAttachmentId,omitempty"`
	// The ID of the MCP Server.
	//
	// example:
	//
	// mcp-adfefz24afg
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The access path to the MCP Server provided by the gateway.
	//
	// example:
	//
	// /mcp-servers/test-mcp
	McpServerPath *string `json:"mcpServerPath,omitempty" xml:"mcpServerPath,omitempty"`
	// Specifies whether to enable MCP observability. The default value is `false`.
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// Information about the MCP Server synchronized from Nacos.
	NacosMcpSyncInfo *GetMcpServerResponseBodyDataNacosMcpSyncInfo `json:"nacosMcpSyncInfo,omitempty" xml:"nacosMcpSyncInfo,omitempty" type:"Struct"`
	// The name of the MCP Server.
	//
	// example:
	//
	// test-mcp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// hr-cr82undlhtgrlej***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
	// The type of the MCP Server.
	//
	// example:
	//
	// RealMCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMcpServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *GetMcpServerResponseBodyData) GetAssembledSources() []*GetMcpServerResponseBodyDataAssembledSources {
	return s.AssembledSources
}

func (s *GetMcpServerResponseBodyData) GetBackend() *Backend {
	return s.Backend
}

func (s *GetMcpServerResponseBodyData) GetCreateFromType() *string {
	return s.CreateFromType
}

func (s *GetMcpServerResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetMcpServerResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMcpServerResponseBodyData) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *GetMcpServerResponseBodyData) GetDomainInfos() []*GetMcpServerResponseBodyDataDomainInfos {
	return s.DomainInfos
}

func (s *GetMcpServerResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetMcpServerResponseBodyData) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *GetMcpServerResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetMcpServerResponseBodyData) GetGrayMcpServerConfigs() []*GetMcpServerResponseBodyDataGrayMcpServerConfigs {
	return s.GrayMcpServerConfigs
}

func (s *GetMcpServerResponseBodyData) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *GetMcpServerResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *GetMcpServerResponseBodyData) GetMcpServerConfigPluginAttachmentId() *string {
	return s.McpServerConfigPluginAttachmentId
}

func (s *GetMcpServerResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpServerResponseBodyData) GetMcpServerPath() *string {
	return s.McpServerPath
}

func (s *GetMcpServerResponseBodyData) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *GetMcpServerResponseBodyData) GetNacosMcpSyncInfo() *GetMcpServerResponseBodyDataNacosMcpSyncInfo {
	return s.NacosMcpSyncInfo
}

func (s *GetMcpServerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMcpServerResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpServerResponseBodyData) GetRouteId() *string {
	return s.RouteId
}

func (s *GetMcpServerResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMcpServerResponseBodyData) SetApiId(v string) *GetMcpServerResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetAssembledSources(v []*GetMcpServerResponseBodyDataAssembledSources) *GetMcpServerResponseBodyData {
	s.AssembledSources = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetBackend(v *Backend) *GetMcpServerResponseBodyData {
	s.Backend = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetCreateFromType(v string) *GetMcpServerResponseBodyData {
	s.CreateFromType = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetDeployStatus(v string) *GetMcpServerResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetDescription(v string) *GetMcpServerResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetDomainIds(v []*string) *GetMcpServerResponseBodyData {
	s.DomainIds = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetDomainInfos(v []*GetMcpServerResponseBodyDataDomainInfos) *GetMcpServerResponseBodyData {
	s.DomainInfos = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetEnvironmentId(v string) *GetMcpServerResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetExposedUriPath(v string) *GetMcpServerResponseBodyData {
	s.ExposedUriPath = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetGatewayId(v string) *GetMcpServerResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetGrayMcpServerConfigs(v []*GetMcpServerResponseBodyDataGrayMcpServerConfigs) *GetMcpServerResponseBodyData {
	s.GrayMcpServerConfigs = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMatch(v *HttpRouteMatch) *GetMcpServerResponseBodyData {
	s.Match = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMcpServerConfig(v string) *GetMcpServerResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMcpServerConfigPluginAttachmentId(v string) *GetMcpServerResponseBodyData {
	s.McpServerConfigPluginAttachmentId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMcpServerId(v string) *GetMcpServerResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMcpServerPath(v string) *GetMcpServerResponseBodyData {
	s.McpServerPath = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetMcpStatisticsEnable(v bool) *GetMcpServerResponseBodyData {
	s.McpStatisticsEnable = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetNacosMcpSyncInfo(v *GetMcpServerResponseBodyDataNacosMcpSyncInfo) *GetMcpServerResponseBodyData {
	s.NacosMcpSyncInfo = v
	return s
}

func (s *GetMcpServerResponseBodyData) SetName(v string) *GetMcpServerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetProtocol(v string) *GetMcpServerResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetRouteId(v string) *GetMcpServerResponseBodyData {
	s.RouteId = &v
	return s
}

func (s *GetMcpServerResponseBodyData) SetType(v string) *GetMcpServerResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMcpServerResponseBodyData) Validate() error {
	if s.AssembledSources != nil {
		for _, item := range s.AssembledSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Backend != nil {
		if err := s.Backend.Validate(); err != nil {
			return err
		}
	}
	if s.DomainInfos != nil {
		for _, item := range s.DomainInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GrayMcpServerConfigs != nil {
		for _, item := range s.GrayMcpServerConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.NacosMcpSyncInfo != nil {
		if err := s.NacosMcpSyncInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServerResponseBodyDataAssembledSources struct {
	// The ID of the MCP Server.
	//
	// example:
	//
	// mcp-xdafeafzz
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name of the MCP Server.
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// A list of MCP tools.
	Tools []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s GetMcpServerResponseBodyDataAssembledSources) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataAssembledSources) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataAssembledSources) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpServerResponseBodyDataAssembledSources) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *GetMcpServerResponseBodyDataAssembledSources) GetTools() []*string {
	return s.Tools
}

func (s *GetMcpServerResponseBodyDataAssembledSources) SetMcpServerId(v string) *GetMcpServerResponseBodyDataAssembledSources {
	s.McpServerId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataAssembledSources) SetMcpServerName(v string) *GetMcpServerResponseBodyDataAssembledSources {
	s.McpServerName = &v
	return s
}

func (s *GetMcpServerResponseBodyDataAssembledSources) SetTools(v []*string) *GetMcpServerResponseBodyDataAssembledSources {
	s.Tools = v
	return s
}

func (s *GetMcpServerResponseBodyDataAssembledSources) Validate() error {
	return dara.Validate(s)
}

type GetMcpServerResponseBodyDataDomainInfos struct {
	// The ID of the domain name.
	//
	// example:
	//
	// www.abc.com
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// verifyicket
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP,HTTPS
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetMcpServerResponseBodyDataDomainInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataDomainInfos) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataDomainInfos) GetDomainId() *string {
	return s.DomainId
}

func (s *GetMcpServerResponseBodyDataDomainInfos) GetName() *string {
	return s.Name
}

func (s *GetMcpServerResponseBodyDataDomainInfos) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpServerResponseBodyDataDomainInfos) SetDomainId(v string) *GetMcpServerResponseBodyDataDomainInfos {
	s.DomainId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataDomainInfos) SetName(v string) *GetMcpServerResponseBodyDataDomainInfos {
	s.Name = &v
	return s
}

func (s *GetMcpServerResponseBodyDataDomainInfos) SetProtocol(v string) *GetMcpServerResponseBodyDataDomainInfos {
	s.Protocol = &v
	return s
}

func (s *GetMcpServerResponseBodyDataDomainInfos) Validate() error {
	return dara.Validate(s)
}

type GetMcpServerResponseBodyDataGrayMcpServerConfigs struct {
	// The backend configuration for the canary route.
	BackendConfig *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// The match rule for the canary route.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The ID of the canary route.
	//
	// example:
	//
	// gray-route-xxx
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigs) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) GetBackendConfig() *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig {
	return s.BackendConfig
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) GetRouteId() *string {
	return s.RouteId
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) SetBackendConfig(v *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) *GetMcpServerResponseBodyDataGrayMcpServerConfigs {
	s.BackendConfig = v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) SetMatch(v *HttpRouteMatch) *GetMcpServerResponseBodyDataGrayMcpServerConfigs {
	s.Match = v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) SetRouteId(v string) *GetMcpServerResponseBodyDataGrayMcpServerConfigs {
	s.RouteId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigs) Validate() error {
	if s.BackendConfig != nil {
		if err := s.BackendConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig struct {
	// The backend scene.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// A list of service configurations.
	Services []*GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) GetServices() []*GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	return s.Services
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) SetScene(v string) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig {
	s.Scene = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) SetServices(v []*GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig {
	s.Services = v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfig) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices struct {
	// The service port.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v2.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) SetPort(v int32) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	s.Port = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) SetProtocol(v string) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) SetServiceId(v string) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) SetVersion(v string) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	s.Version = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) SetWeight(v int32) *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *GetMcpServerResponseBodyDataGrayMcpServerConfigsBackendConfigServices) Validate() error {
	return dara.Validate(s)
}

type GetMcpServerResponseBodyDataNacosMcpSyncInfo struct {
	// The ID of the Nacos instance.
	//
	// example:
	//
	// mse-xxxxx
	ImportInstanceId *string `json:"importInstanceId,omitempty" xml:"importInstanceId,omitempty"`
	// The ID of the synchronized MCP Server.
	//
	// example:
	//
	// test-mcp
	ImportMcpServerId *string `json:"importMcpServerId,omitempty" xml:"importMcpServerId,omitempty"`
	// The Nacos namespace.
	//
	// example:
	//
	// test-ns
	ImportNamespace *string `json:"importNamespace,omitempty" xml:"importNamespace,omitempty"`
}

func (s GetMcpServerResponseBodyDataNacosMcpSyncInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyDataNacosMcpSyncInfo) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) GetImportInstanceId() *string {
	return s.ImportInstanceId
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) GetImportMcpServerId() *string {
	return s.ImportMcpServerId
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) GetImportNamespace() *string {
	return s.ImportNamespace
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) SetImportInstanceId(v string) *GetMcpServerResponseBodyDataNacosMcpSyncInfo {
	s.ImportInstanceId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) SetImportMcpServerId(v string) *GetMcpServerResponseBodyDataNacosMcpSyncInfo {
	s.ImportMcpServerId = &v
	return s
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) SetImportNamespace(v string) *GetMcpServerResponseBodyDataNacosMcpSyncInfo {
	s.ImportNamespace = &v
	return s
}

func (s *GetMcpServerResponseBodyDataNacosMcpSyncInfo) Validate() error {
	return dara.Validate(s)
}

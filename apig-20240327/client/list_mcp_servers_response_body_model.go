// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpServersResponseBody
	GetCode() *string
	SetData(v *ListMcpServersResponseBodyData) *ListMcpServersResponseBody
	GetData() *ListMcpServersResponseBodyData
	SetMessage(v string) *ListMcpServersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMcpServersResponseBody
	GetRequestId() *string
}

type ListMcpServersResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListMcpServersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMcpServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpServersResponseBody) GetData() *ListMcpServersResponseBodyData {
	return s.Data
}

func (s *ListMcpServersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpServersResponseBody) SetCode(v string) *ListMcpServersResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpServersResponseBody) SetData(v *ListMcpServersResponseBodyData) *ListMcpServersResponseBody {
	s.Data = v
	return s
}

func (s *ListMcpServersResponseBody) SetMessage(v string) *ListMcpServersResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpServersResponseBody) SetRequestId(v string) *ListMcpServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpServersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpServersResponseBodyData struct {
	// The list of MCP servers.
	Items []*ListMcpServersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListMcpServersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyData) GetItems() []*ListMcpServersResponseBodyDataItems {
	return s.Items
}

func (s *ListMcpServersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMcpServersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcpServersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListMcpServersResponseBodyData) SetItems(v []*ListMcpServersResponseBodyDataItems) *ListMcpServersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMcpServersResponseBodyData) SetPageNumber(v int32) *ListMcpServersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMcpServersResponseBodyData) SetPageSize(v int32) *ListMcpServersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMcpServersResponseBodyData) SetTotalSize(v int32) *ListMcpServersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListMcpServersResponseBodyData) Validate() error {
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

type ListMcpServersResponseBodyDataItems struct {
	// API ID。
	//
	// example:
	//
	// api-xxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The list of assembly sources. This parameter is required when type is AssemblyMCP.
	AssembledSources []*ListMcpServersResponseBodyDataItemsAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	// The backend service of the route.
	Backend *Backend `json:"backend,omitempty" xml:"backend,omitempty"`
	// The creation source type. Valid values:
	//
	// - ApiGatewayHttpToMCP: gateway-managed HTTP-to-MCP conversion.
	//
	// - ApiGatewayMcpHosting: gateway-managed MCP direct proxy.
	//
	// - ApiGatewayAssembly: gateway MCP assembly.
	//
	// - NacosHttpToMCP: gateway-managed Nacos-synced HTTP-to-MCP conversion.
	//
	// - NacosMcpHosting: gateway-managed Nacos-synced MCP direct proxy.
	//
	// example:
	//
	// ApiGatewayHttpToMCP
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// The deployment status of the API in the current environment.
	//
	// example:
	//
	// Deployed
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// The description.
	//
	// example:
	//
	// 这是xxx的xx项目测试环境
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of domain name IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The list of domain name information.
	DomainInfos []*HttpApiDomainInfo `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The exposed URI path. This parameter is required when protocol is SSE or StreamableHTTP and type is RealMCP.
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// gw-cpv54p5***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The HTTP-to-MCP configuration.
	//
	// example:
	//
	// HTTP转MCP Config base64值
	McpServerConfig *string `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
	// MCP Server ID
	//
	// example:
	//
	// mcp-feaff34va
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The MCP server access path provided by the gateway.
	//
	// example:
	//
	// /mcp-servers/test-mcp
	McpServerPath *string `json:"mcpServerPath,omitempty" xml:"mcpServerPath,omitempty"`
	// Indicates whether MCP observability is enabled. Default value: false.
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// The Nacos-synced managed MCP information.
	NacosMcpSyncInfo *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo `json:"nacosMcpSyncInfo,omitempty" xml:"nacosMcpSyncInfo,omitempty" type:"Struct"`
	// The MCP server name.
	//
	// example:
	//
	// itemcenter-dev-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The route ID associated with the MCP server.
	//
	// example:
	//
	// hr-d11cj86m1hkvop6mp42g
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
	// The MCP server type.
	//
	// example:
	//
	// 可选值：RealMCP、AssemblyMCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMcpServersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyDataItems) GetApiId() *string {
	return s.ApiId
}

func (s *ListMcpServersResponseBodyDataItems) GetAssembledSources() []*ListMcpServersResponseBodyDataItemsAssembledSources {
	return s.AssembledSources
}

func (s *ListMcpServersResponseBodyDataItems) GetBackend() *Backend {
	return s.Backend
}

func (s *ListMcpServersResponseBodyDataItems) GetCreateFromType() *string {
	return s.CreateFromType
}

func (s *ListMcpServersResponseBodyDataItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListMcpServersResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpServersResponseBodyDataItems) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *ListMcpServersResponseBodyDataItems) GetDomainInfos() []*HttpApiDomainInfo {
	return s.DomainInfos
}

func (s *ListMcpServersResponseBodyDataItems) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListMcpServersResponseBodyDataItems) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *ListMcpServersResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListMcpServersResponseBodyDataItems) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *ListMcpServersResponseBodyDataItems) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *ListMcpServersResponseBodyDataItems) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpServersResponseBodyDataItems) GetMcpServerPath() *string {
	return s.McpServerPath
}

func (s *ListMcpServersResponseBodyDataItems) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *ListMcpServersResponseBodyDataItems) GetNacosMcpSyncInfo() *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo {
	return s.NacosMcpSyncInfo
}

func (s *ListMcpServersResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListMcpServersResponseBodyDataItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMcpServersResponseBodyDataItems) GetRouteId() *string {
	return s.RouteId
}

func (s *ListMcpServersResponseBodyDataItems) GetType() *string {
	return s.Type
}

func (s *ListMcpServersResponseBodyDataItems) SetApiId(v string) *ListMcpServersResponseBodyDataItems {
	s.ApiId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetAssembledSources(v []*ListMcpServersResponseBodyDataItemsAssembledSources) *ListMcpServersResponseBodyDataItems {
	s.AssembledSources = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetBackend(v *Backend) *ListMcpServersResponseBodyDataItems {
	s.Backend = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetCreateFromType(v string) *ListMcpServersResponseBodyDataItems {
	s.CreateFromType = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetDeployStatus(v string) *ListMcpServersResponseBodyDataItems {
	s.DeployStatus = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetDescription(v string) *ListMcpServersResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetDomainIds(v []*string) *ListMcpServersResponseBodyDataItems {
	s.DomainIds = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetDomainInfos(v []*HttpApiDomainInfo) *ListMcpServersResponseBodyDataItems {
	s.DomainInfos = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetEnvironmentId(v string) *ListMcpServersResponseBodyDataItems {
	s.EnvironmentId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetExposedUriPath(v string) *ListMcpServersResponseBodyDataItems {
	s.ExposedUriPath = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetGatewayId(v string) *ListMcpServersResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetMatch(v *HttpRouteMatch) *ListMcpServersResponseBodyDataItems {
	s.Match = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetMcpServerConfig(v string) *ListMcpServersResponseBodyDataItems {
	s.McpServerConfig = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetMcpServerId(v string) *ListMcpServersResponseBodyDataItems {
	s.McpServerId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetMcpServerPath(v string) *ListMcpServersResponseBodyDataItems {
	s.McpServerPath = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetMcpStatisticsEnable(v bool) *ListMcpServersResponseBodyDataItems {
	s.McpStatisticsEnable = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetNacosMcpSyncInfo(v *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) *ListMcpServersResponseBodyDataItems {
	s.NacosMcpSyncInfo = v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetName(v string) *ListMcpServersResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetProtocol(v string) *ListMcpServersResponseBodyDataItems {
	s.Protocol = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetRouteId(v string) *ListMcpServersResponseBodyDataItems {
	s.RouteId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) SetType(v string) *ListMcpServersResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItems) Validate() error {
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

type ListMcpServersResponseBodyDataItemsAssembledSources struct {
	// MCP Server ID
	//
	// example:
	//
	// mcp-adfef2334fa
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The MCP server name.
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// The list of MCP tools.
	Tools []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s ListMcpServersResponseBodyDataItemsAssembledSources) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyDataItemsAssembledSources) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) GetTools() []*string {
	return s.Tools
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) SetMcpServerId(v string) *ListMcpServersResponseBodyDataItemsAssembledSources {
	s.McpServerId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) SetMcpServerName(v string) *ListMcpServersResponseBodyDataItemsAssembledSources {
	s.McpServerName = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) SetTools(v []*string) *ListMcpServersResponseBodyDataItemsAssembledSources {
	s.Tools = v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsAssembledSources) Validate() error {
	return dara.Validate(s)
}

type ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo struct {
	// The Nacos instance.
	//
	// example:
	//
	// mse-faefrefxz
	ImportInstanceId *string `json:"importInstanceId,omitempty" xml:"importInstanceId,omitempty"`
	// The synced MCP server ID.
	//
	// example:
	//
	// 同步的MCP Server ID
	ImportMcpServerId *string `json:"importMcpServerId,omitempty" xml:"importMcpServerId,omitempty"`
	// The Nacos namespace.
	//
	// example:
	//
	// test-ns
	ImportNamespace *string `json:"importNamespace,omitempty" xml:"importNamespace,omitempty"`
}

func (s ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) GetImportInstanceId() *string {
	return s.ImportInstanceId
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) GetImportMcpServerId() *string {
	return s.ImportMcpServerId
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) GetImportNamespace() *string {
	return s.ImportNamespace
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) SetImportInstanceId(v string) *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo {
	s.ImportInstanceId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) SetImportMcpServerId(v string) *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo {
	s.ImportMcpServerId = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) SetImportNamespace(v string) *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo {
	s.ImportNamespace = &v
	return s
}

func (s *ListMcpServersResponseBodyDataItemsNacosMcpSyncInfo) Validate() error {
	return dara.Validate(s)
}

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
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetMcpServerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
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
	return dara.Validate(s)
}

type GetMcpServerResponseBodyData struct {
	AssembledSources []*GetMcpServerResponseBodyDataAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	Backend          *Backend                                        `json:"backend,omitempty" xml:"backend,omitempty"`
	// example:
	//
	// ApiGatewayHttpToMCP
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// example:
	//
	// Deployed
	DeployStatus *string                                    `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	Description  *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	DomainIds    []*string                                  `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	DomainInfos  []*GetMcpServerResponseBodyDataDomainInfos `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// {\\"product_code\\":\\"apigw\\"}
	Match           *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	McpServerConfig *string         `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
	// example:
	//
	// pa-adfaefwaef
	McpServerConfigPluginAttachmentId *string `json:"mcpServerConfigPluginAttachmentId,omitempty" xml:"mcpServerConfigPluginAttachmentId,omitempty"`
	// MCP Server ID
	//
	// example:
	//
	// mcp-adfefz24afg
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// example:
	//
	// /mcp-servers/test-mcp
	McpServerPath *string `json:"mcpServerPath,omitempty" xml:"mcpServerPath,omitempty"`
	// example:
	//
	// false
	McpStatisticsEnable *bool                                         `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	NacosMcpSyncInfo    *GetMcpServerResponseBodyDataNacosMcpSyncInfo `json:"nacosMcpSyncInfo,omitempty" xml:"nacosMcpSyncInfo,omitempty" type:"Struct"`
	// example:
	//
	// test-mcp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// hr-cr82undlhtgrlej***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
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
	return dara.Validate(s)
}

type GetMcpServerResponseBodyDataAssembledSources struct {
	// MCP Server ID
	//
	// example:
	//
	// mcp-xdafeafzz
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// example:
	//
	// test-mcp
	McpServerName *string   `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	Tools         []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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
	// example:
	//
	// www.abc.com
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// example:
	//
	// verifyicket
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

type GetMcpServerResponseBodyDataNacosMcpSyncInfo struct {
	// example:
	//
	// mse-xxxxx
	ImportInstanceId *string `json:"importInstanceId,omitempty" xml:"importInstanceId,omitempty"`
	// example:
	//
	// test-mcp
	ImportMcpServerId *string `json:"importMcpServerId,omitempty" xml:"importMcpServerId,omitempty"`
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

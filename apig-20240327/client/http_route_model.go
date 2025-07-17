// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRoute interface {
	dara.Model
	String() string
	GoString() string
	SetBackend(v *Backend) *HttpRoute
	GetBackend() *Backend
	SetCreateTimestamp(v int64) *HttpRoute
	GetCreateTimestamp() *int64
	SetDeployStatus(v string) *HttpRoute
	GetDeployStatus() *string
	SetDescription(v string) *HttpRoute
	GetDescription() *string
	SetDomainInfos(v []*HttpRouteDomainInfos) *HttpRoute
	GetDomainInfos() []*HttpRouteDomainInfos
	SetEnvironmentInfo(v *HttpRouteEnvironmentInfo) *HttpRoute
	GetEnvironmentInfo() *HttpRouteEnvironmentInfo
	SetGatewayStatus(v map[string]*string) *HttpRoute
	GetGatewayStatus() map[string]*string
	SetMatch(v *HttpRouteMatch) *HttpRoute
	GetMatch() *HttpRouteMatch
	SetMcpServerInfo(v *HttpRouteMcpServerInfo) *HttpRoute
	GetMcpServerInfo() *HttpRouteMcpServerInfo
	SetName(v string) *HttpRoute
	GetName() *string
	SetRouteId(v string) *HttpRoute
	GetRouteId() *string
	SetUpdateTimestamp(v int64) *HttpRoute
	GetUpdateTimestamp() *int64
}

type HttpRoute struct {
	Backend         *Backend                  `json:"backend,omitempty" xml:"backend,omitempty"`
	CreateTimestamp *int64                    `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	DeployStatus    *string                   `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	Description     *string                   `json:"description,omitempty" xml:"description,omitempty"`
	DomainInfos     []*HttpRouteDomainInfos   `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	EnvironmentInfo *HttpRouteEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	GatewayStatus   map[string]*string        `json:"gatewayStatus,omitempty" xml:"gatewayStatus,omitempty"`
	Match           *HttpRouteMatch           `json:"match,omitempty" xml:"match,omitempty"`
	McpServerInfo   *HttpRouteMcpServerInfo   `json:"mcpServerInfo,omitempty" xml:"mcpServerInfo,omitempty" type:"Struct"`
	Name            *string                   `json:"name,omitempty" xml:"name,omitempty"`
	RouteId         *string                   `json:"routeId,omitempty" xml:"routeId,omitempty"`
	UpdateTimestamp *int64                    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s HttpRoute) String() string {
	return dara.Prettify(s)
}

func (s HttpRoute) GoString() string {
	return s.String()
}

func (s *HttpRoute) GetBackend() *Backend {
	return s.Backend
}

func (s *HttpRoute) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *HttpRoute) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *HttpRoute) GetDescription() *string {
	return s.Description
}

func (s *HttpRoute) GetDomainInfos() []*HttpRouteDomainInfos {
	return s.DomainInfos
}

func (s *HttpRoute) GetEnvironmentInfo() *HttpRouteEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *HttpRoute) GetGatewayStatus() map[string]*string {
	return s.GatewayStatus
}

func (s *HttpRoute) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *HttpRoute) GetMcpServerInfo() *HttpRouteMcpServerInfo {
	return s.McpServerInfo
}

func (s *HttpRoute) GetName() *string {
	return s.Name
}

func (s *HttpRoute) GetRouteId() *string {
	return s.RouteId
}

func (s *HttpRoute) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *HttpRoute) SetBackend(v *Backend) *HttpRoute {
	s.Backend = v
	return s
}

func (s *HttpRoute) SetCreateTimestamp(v int64) *HttpRoute {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpRoute) SetDeployStatus(v string) *HttpRoute {
	s.DeployStatus = &v
	return s
}

func (s *HttpRoute) SetDescription(v string) *HttpRoute {
	s.Description = &v
	return s
}

func (s *HttpRoute) SetDomainInfos(v []*HttpRouteDomainInfos) *HttpRoute {
	s.DomainInfos = v
	return s
}

func (s *HttpRoute) SetEnvironmentInfo(v *HttpRouteEnvironmentInfo) *HttpRoute {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpRoute) SetGatewayStatus(v map[string]*string) *HttpRoute {
	s.GatewayStatus = v
	return s
}

func (s *HttpRoute) SetMatch(v *HttpRouteMatch) *HttpRoute {
	s.Match = v
	return s
}

func (s *HttpRoute) SetMcpServerInfo(v *HttpRouteMcpServerInfo) *HttpRoute {
	s.McpServerInfo = v
	return s
}

func (s *HttpRoute) SetName(v string) *HttpRoute {
	s.Name = &v
	return s
}

func (s *HttpRoute) SetRouteId(v string) *HttpRoute {
	s.RouteId = &v
	return s
}

func (s *HttpRoute) SetUpdateTimestamp(v int64) *HttpRoute {
	s.UpdateTimestamp = &v
	return s
}

func (s *HttpRoute) Validate() error {
	return dara.Validate(s)
}

type HttpRouteDomainInfos struct {
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpRouteDomainInfos) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteDomainInfos) GoString() string {
	return s.String()
}

func (s *HttpRouteDomainInfos) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpRouteDomainInfos) GetName() *string {
	return s.Name
}

func (s *HttpRouteDomainInfos) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpRouteDomainInfos) SetDomainId(v string) *HttpRouteDomainInfos {
	s.DomainId = &v
	return s
}

func (s *HttpRouteDomainInfos) SetName(v string) *HttpRouteDomainInfos {
	s.Name = &v
	return s
}

func (s *HttpRouteDomainInfos) SetProtocol(v string) *HttpRouteDomainInfos {
	s.Protocol = &v
	return s
}

func (s *HttpRouteDomainInfos) Validate() error {
	return dara.Validate(s)
}

type HttpRouteEnvironmentInfo struct {
	Alias         *string                               `json:"alias,omitempty" xml:"alias,omitempty"`
	EnvironmentId *string                               `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpRouteEnvironmentInfoGatewayInfo  `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	Name          *string                               `json:"name,omitempty" xml:"name,omitempty"`
	SubDomains    []*HttpRouteEnvironmentInfoSubDomains `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
}

func (s HttpRouteEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpRouteEnvironmentInfo) GetAlias() *string {
	return s.Alias
}

func (s *HttpRouteEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpRouteEnvironmentInfo) GetGatewayInfo() *HttpRouteEnvironmentInfoGatewayInfo {
	return s.GatewayInfo
}

func (s *HttpRouteEnvironmentInfo) GetName() *string {
	return s.Name
}

func (s *HttpRouteEnvironmentInfo) GetSubDomains() []*HttpRouteEnvironmentInfoSubDomains {
	return s.SubDomains
}

func (s *HttpRouteEnvironmentInfo) SetAlias(v string) *HttpRouteEnvironmentInfo {
	s.Alias = &v
	return s
}

func (s *HttpRouteEnvironmentInfo) SetEnvironmentId(v string) *HttpRouteEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *HttpRouteEnvironmentInfo) SetGatewayInfo(v *HttpRouteEnvironmentInfoGatewayInfo) *HttpRouteEnvironmentInfo {
	s.GatewayInfo = v
	return s
}

func (s *HttpRouteEnvironmentInfo) SetName(v string) *HttpRouteEnvironmentInfo {
	s.Name = &v
	return s
}

func (s *HttpRouteEnvironmentInfo) SetSubDomains(v []*HttpRouteEnvironmentInfoSubDomains) *HttpRouteEnvironmentInfo {
	s.SubDomains = v
	return s
}

func (s *HttpRouteEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type HttpRouteEnvironmentInfoGatewayInfo struct {
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpRouteEnvironmentInfoGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteEnvironmentInfoGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) GetName() *string {
	return s.Name
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) SetGatewayId(v string) *HttpRouteEnvironmentInfoGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) SetName(v string) *HttpRouteEnvironmentInfoGatewayInfo {
	s.Name = &v
	return s
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type HttpRouteEnvironmentInfoSubDomains struct {
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Internet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Protocol    *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpRouteEnvironmentInfoSubDomains) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteEnvironmentInfoSubDomains) GoString() string {
	return s.String()
}

func (s *HttpRouteEnvironmentInfoSubDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpRouteEnvironmentInfoSubDomains) GetName() *string {
	return s.Name
}

func (s *HttpRouteEnvironmentInfoSubDomains) GetNetworkType() *string {
	return s.NetworkType
}

func (s *HttpRouteEnvironmentInfoSubDomains) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpRouteEnvironmentInfoSubDomains) SetDomainId(v string) *HttpRouteEnvironmentInfoSubDomains {
	s.DomainId = &v
	return s
}

func (s *HttpRouteEnvironmentInfoSubDomains) SetName(v string) *HttpRouteEnvironmentInfoSubDomains {
	s.Name = &v
	return s
}

func (s *HttpRouteEnvironmentInfoSubDomains) SetNetworkType(v string) *HttpRouteEnvironmentInfoSubDomains {
	s.NetworkType = &v
	return s
}

func (s *HttpRouteEnvironmentInfoSubDomains) SetProtocol(v string) *HttpRouteEnvironmentInfoSubDomains {
	s.Protocol = &v
	return s
}

func (s *HttpRouteEnvironmentInfoSubDomains) Validate() error {
	return dara.Validate(s)
}

type HttpRouteMcpServerInfo struct {
	CreateFromType    *string                               `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	ImportInstanceId  *string                               `json:"importInstanceId,omitempty" xml:"importInstanceId,omitempty"`
	ImportMcpServerId *string                               `json:"importMcpServerId,omitempty" xml:"importMcpServerId,omitempty"`
	ImportNamespace   *string                               `json:"importNamespace,omitempty" xml:"importNamespace,omitempty"`
	McpRouteConfig    *HttpRouteMcpServerInfoMcpRouteConfig `json:"mcpRouteConfig,omitempty" xml:"mcpRouteConfig,omitempty" type:"Struct"`
	McpServerConfig   *string                               `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
}

func (s HttpRouteMcpServerInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMcpServerInfo) GoString() string {
	return s.String()
}

func (s *HttpRouteMcpServerInfo) GetCreateFromType() *string {
	return s.CreateFromType
}

func (s *HttpRouteMcpServerInfo) GetImportInstanceId() *string {
	return s.ImportInstanceId
}

func (s *HttpRouteMcpServerInfo) GetImportMcpServerId() *string {
	return s.ImportMcpServerId
}

func (s *HttpRouteMcpServerInfo) GetImportNamespace() *string {
	return s.ImportNamespace
}

func (s *HttpRouteMcpServerInfo) GetMcpRouteConfig() *HttpRouteMcpServerInfoMcpRouteConfig {
	return s.McpRouteConfig
}

func (s *HttpRouteMcpServerInfo) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *HttpRouteMcpServerInfo) SetCreateFromType(v string) *HttpRouteMcpServerInfo {
	s.CreateFromType = &v
	return s
}

func (s *HttpRouteMcpServerInfo) SetImportInstanceId(v string) *HttpRouteMcpServerInfo {
	s.ImportInstanceId = &v
	return s
}

func (s *HttpRouteMcpServerInfo) SetImportMcpServerId(v string) *HttpRouteMcpServerInfo {
	s.ImportMcpServerId = &v
	return s
}

func (s *HttpRouteMcpServerInfo) SetImportNamespace(v string) *HttpRouteMcpServerInfo {
	s.ImportNamespace = &v
	return s
}

func (s *HttpRouteMcpServerInfo) SetMcpRouteConfig(v *HttpRouteMcpServerInfoMcpRouteConfig) *HttpRouteMcpServerInfo {
	s.McpRouteConfig = v
	return s
}

func (s *HttpRouteMcpServerInfo) SetMcpServerConfig(v string) *HttpRouteMcpServerInfo {
	s.McpServerConfig = &v
	return s
}

func (s *HttpRouteMcpServerInfo) Validate() error {
	return dara.Validate(s)
}

type HttpRouteMcpServerInfoMcpRouteConfig struct {
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	Protocol       *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpRouteMcpServerInfoMcpRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMcpServerInfoMcpRouteConfig) GoString() string {
	return s.String()
}

func (s *HttpRouteMcpServerInfoMcpRouteConfig) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *HttpRouteMcpServerInfoMcpRouteConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpRouteMcpServerInfoMcpRouteConfig) SetExposedUriPath(v string) *HttpRouteMcpServerInfoMcpRouteConfig {
	s.ExposedUriPath = &v
	return s
}

func (s *HttpRouteMcpServerInfoMcpRouteConfig) SetProtocol(v string) *HttpRouteMcpServerInfoMcpRouteConfig {
	s.Protocol = &v
	return s
}

func (s *HttpRouteMcpServerInfoMcpRouteConfig) Validate() error {
	return dara.Validate(s)
}

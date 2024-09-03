// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Attachment struct {
	AttachResourceIds  []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	AttachResourceType *string   `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	EnvironmentId      *string   `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayId          *string   `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	PolicyAttachmentId *string   `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
}

func (s Attachment) String() string {
	return tea.Prettify(s)
}

func (s Attachment) GoString() string {
	return s.String()
}

func (s *Attachment) SetAttachResourceIds(v []*string) *Attachment {
	s.AttachResourceIds = v
	return s
}

func (s *Attachment) SetAttachResourceType(v string) *Attachment {
	s.AttachResourceType = &v
	return s
}

func (s *Attachment) SetEnvironmentId(v string) *Attachment {
	s.EnvironmentId = &v
	return s
}

func (s *Attachment) SetGatewayId(v string) *Attachment {
	s.GatewayId = &v
	return s
}

func (s *Attachment) SetPolicyAttachmentId(v string) *Attachment {
	s.PolicyAttachmentId = &v
	return s
}

type CheckServiceLinkedRoleResult struct {
	Existed *bool `json:"existed,omitempty" xml:"existed,omitempty"`
}

func (s CheckServiceLinkedRoleResult) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResult) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResult) SetExisted(v bool) *CheckServiceLinkedRoleResult {
	s.Existed = &v
	return s
}

type DashboardFilter struct {
	// example:
	//
	// test
	RouteName *string `json:"routeName,omitempty" xml:"routeName,omitempty"`
}

func (s DashboardFilter) String() string {
	return tea.Prettify(s)
}

func (s DashboardFilter) GoString() string {
	return s.String()
}

func (s *DashboardFilter) SetRouteName(v string) *DashboardFilter {
	s.RouteName = &v
	return s
}

type DomainInfo struct {
	CertIdentifier  *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	CreateFrom      *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	CreateTimestamp *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	DomainId        *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	ForceHttps      *bool   `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol        *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTimestamp *int64  `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s DomainInfo) String() string {
	return tea.Prettify(s)
}

func (s DomainInfo) GoString() string {
	return s.String()
}

func (s *DomainInfo) SetCertIdentifier(v string) *DomainInfo {
	s.CertIdentifier = &v
	return s
}

func (s *DomainInfo) SetCreateFrom(v string) *DomainInfo {
	s.CreateFrom = &v
	return s
}

func (s *DomainInfo) SetCreateTimestamp(v int64) *DomainInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *DomainInfo) SetDomainId(v string) *DomainInfo {
	s.DomainId = &v
	return s
}

func (s *DomainInfo) SetForceHttps(v bool) *DomainInfo {
	s.ForceHttps = &v
	return s
}

func (s *DomainInfo) SetName(v string) *DomainInfo {
	s.Name = &v
	return s
}

func (s *DomainInfo) SetProtocol(v string) *DomainInfo {
	s.Protocol = &v
	return s
}

func (s *DomainInfo) SetStatus(v string) *DomainInfo {
	s.Status = &v
	return s
}

func (s *DomainInfo) SetUpdateTimestamp(v int64) *DomainInfo {
	s.UpdateTimestamp = &v
	return s
}

type EnvironmentInfo struct {
	Alias           *string          `json:"alias,omitempty" xml:"alias,omitempty"`
	CreateTimestamp *int64           `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Default         *bool            `json:"default,omitempty" xml:"default,omitempty"`
	Description     *string          `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentId   *string          `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo     *GatewayInfo     `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	Name            *string          `json:"name,omitempty" xml:"name,omitempty"`
	SubDomainInfos  []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	UpdateTimestamp *int64           `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s EnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s EnvironmentInfo) GoString() string {
	return s.String()
}

func (s *EnvironmentInfo) SetAlias(v string) *EnvironmentInfo {
	s.Alias = &v
	return s
}

func (s *EnvironmentInfo) SetCreateTimestamp(v int64) *EnvironmentInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *EnvironmentInfo) SetDefault(v bool) *EnvironmentInfo {
	s.Default = &v
	return s
}

func (s *EnvironmentInfo) SetDescription(v string) *EnvironmentInfo {
	s.Description = &v
	return s
}

func (s *EnvironmentInfo) SetEnvironmentId(v string) *EnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *EnvironmentInfo) SetGatewayInfo(v *GatewayInfo) *EnvironmentInfo {
	s.GatewayInfo = v
	return s
}

func (s *EnvironmentInfo) SetName(v string) *EnvironmentInfo {
	s.Name = &v
	return s
}

func (s *EnvironmentInfo) SetSubDomainInfos(v []*SubDomainInfo) *EnvironmentInfo {
	s.SubDomainInfos = v
	return s
}

func (s *EnvironmentInfo) SetUpdateTimestamp(v int64) *EnvironmentInfo {
	s.UpdateTimestamp = &v
	return s
}

type GatewayInfo struct {
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayInfo) GoString() string {
	return s.String()
}

func (s *GatewayInfo) SetGatewayId(v string) *GatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *GatewayInfo) SetName(v string) *GatewayInfo {
	s.Name = &v
	return s
}

type GatewayLogConfig struct {
	SlsConfig *GatewayLogConfigSlsConfig `json:"slsConfig,omitempty" xml:"slsConfig,omitempty" type:"Struct"`
}

func (s GatewayLogConfig) String() string {
	return tea.Prettify(s)
}

func (s GatewayLogConfig) GoString() string {
	return s.String()
}

func (s *GatewayLogConfig) SetSlsConfig(v *GatewayLogConfigSlsConfig) *GatewayLogConfig {
	s.SlsConfig = v
	return s
}

type GatewayLogConfigSlsConfig struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GatewayLogConfigSlsConfig) String() string {
	return tea.Prettify(s)
}

func (s GatewayLogConfigSlsConfig) GoString() string {
	return s.String()
}

func (s *GatewayLogConfigSlsConfig) SetEnable(v bool) *GatewayLogConfigSlsConfig {
	s.Enable = &v
	return s
}

type GatewayRouteBackend struct {
	Services []*GatewayRouteBackendServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// example:
	//
	// Single
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GatewayRouteBackend) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteBackend) GoString() string {
	return s.String()
}

func (s *GatewayRouteBackend) SetServices(v []*GatewayRouteBackendServices) *GatewayRouteBackend {
	s.Services = v
	return s
}

func (s *GatewayRouteBackend) SetType(v string) *GatewayRouteBackend {
	s.Type = &v
	return s
}

type GatewayRouteBackendServices struct {
	// example:
	//
	// gs-cq2bmmdlhtgj***
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// example:
	//
	// item-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// port
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 49
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GatewayRouteBackendServices) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteBackendServices) GoString() string {
	return s.String()
}

func (s *GatewayRouteBackendServices) SetGatewayServiceId(v string) *GatewayRouteBackendServices {
	s.GatewayServiceId = &v
	return s
}

func (s *GatewayRouteBackendServices) SetName(v string) *GatewayRouteBackendServices {
	s.Name = &v
	return s
}

func (s *GatewayRouteBackendServices) SetPort(v int32) *GatewayRouteBackendServices {
	s.Port = &v
	return s
}

func (s *GatewayRouteBackendServices) SetProtocol(v string) *GatewayRouteBackendServices {
	s.Protocol = &v
	return s
}

func (s *GatewayRouteBackendServices) SetVersion(v string) *GatewayRouteBackendServices {
	s.Version = &v
	return s
}

func (s *GatewayRouteBackendServices) SetWeight(v float32) *GatewayRouteBackendServices {
	s.Weight = &v
	return s
}

type GatewayRouteBackendConfig struct {
	Services []*GatewayRouteBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// example:
	//
	// Single
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GatewayRouteBackendConfig) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteBackendConfig) GoString() string {
	return s.String()
}

func (s *GatewayRouteBackendConfig) SetServices(v []*GatewayRouteBackendConfigServices) *GatewayRouteBackendConfig {
	s.Services = v
	return s
}

func (s *GatewayRouteBackendConfig) SetType(v string) *GatewayRouteBackendConfig {
	s.Type = &v
	return s
}

type GatewayRouteBackendConfigServices struct {
	// example:
	//
	// gs-cq2bmmdlhtgj***
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// example:
	//
	// port
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// K8S
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// 49
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GatewayRouteBackendConfigServices) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteBackendConfigServices) GoString() string {
	return s.String()
}

func (s *GatewayRouteBackendConfigServices) SetGatewayServiceId(v string) *GatewayRouteBackendConfigServices {
	s.GatewayServiceId = &v
	return s
}

func (s *GatewayRouteBackendConfigServices) SetPort(v int32) *GatewayRouteBackendConfigServices {
	s.Port = &v
	return s
}

func (s *GatewayRouteBackendConfigServices) SetProtocol(v string) *GatewayRouteBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *GatewayRouteBackendConfigServices) SetSourceType(v string) *GatewayRouteBackendConfigServices {
	s.SourceType = &v
	return s
}

func (s *GatewayRouteBackendConfigServices) SetWeight(v float32) *GatewayRouteBackendConfigServices {
	s.Weight = &v
	return s
}

type GatewayRouteDomainConfig struct {
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
}

func (s GatewayRouteDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteDomainConfig) GoString() string {
	return s.String()
}

func (s *GatewayRouteDomainConfig) SetDomainIds(v []*string) *GatewayRouteDomainConfig {
	s.DomainIds = v
	return s
}

type GatewayRouteDomainInfo struct {
	Domains []*GatewayRouteDomainInfoDomains `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GatewayRouteDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteDomainInfo) GoString() string {
	return s.String()
}

func (s *GatewayRouteDomainInfo) SetDomains(v []*GatewayRouteDomainInfoDomains) *GatewayRouteDomainInfo {
	s.Domains = v
	return s
}

type GatewayRouteDomainInfoDomains struct {
	// example:
	//
	// d-cp82or5l***
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// example:
	//
	// item.dev
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GatewayRouteDomainInfoDomains) String() string {
	return tea.Prettify(s)
}

func (s GatewayRouteDomainInfoDomains) GoString() string {
	return s.String()
}

func (s *GatewayRouteDomainInfoDomains) SetDomainId(v string) *GatewayRouteDomainInfoDomains {
	s.DomainId = &v
	return s
}

func (s *GatewayRouteDomainInfoDomains) SetName(v string) *GatewayRouteDomainInfoDomains {
	s.Name = &v
	return s
}

func (s *GatewayRouteDomainInfoDomains) SetProtocol(v string) *GatewayRouteDomainInfoDomains {
	s.Protocol = &v
	return s
}

type GatewayService struct {
	Addresses        []*string              `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	CreateTimestamp  *int64                 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	GatewayServiceId *string                `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	HealthCheck      *ServiceHealthCheck    `json:"healthCheck,omitempty" xml:"healthCheck,omitempty"`
	HealthStatus     *string                `json:"healthStatus,omitempty" xml:"healthStatus,omitempty"`
	Name             *string                `json:"name,omitempty" xml:"name,omitempty"`
	Namespace        *string                `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Ports            []*GatewayServicePorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// LATEST
	Qualifier          *string   `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	SourceType         *string   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UnhealthyEndpoints []*string `json:"unhealthyEndpoints,omitempty" xml:"unhealthyEndpoints,omitempty" type:"Repeated"`
	UpdateTimestamp    *int64    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GatewayService) String() string {
	return tea.Prettify(s)
}

func (s GatewayService) GoString() string {
	return s.String()
}

func (s *GatewayService) SetAddresses(v []*string) *GatewayService {
	s.Addresses = v
	return s
}

func (s *GatewayService) SetCreateTimestamp(v int64) *GatewayService {
	s.CreateTimestamp = &v
	return s
}

func (s *GatewayService) SetGatewayServiceId(v string) *GatewayService {
	s.GatewayServiceId = &v
	return s
}

func (s *GatewayService) SetHealthCheck(v *ServiceHealthCheck) *GatewayService {
	s.HealthCheck = v
	return s
}

func (s *GatewayService) SetHealthStatus(v string) *GatewayService {
	s.HealthStatus = &v
	return s
}

func (s *GatewayService) SetName(v string) *GatewayService {
	s.Name = &v
	return s
}

func (s *GatewayService) SetNamespace(v string) *GatewayService {
	s.Namespace = &v
	return s
}

func (s *GatewayService) SetPorts(v []*GatewayServicePorts) *GatewayService {
	s.Ports = v
	return s
}

func (s *GatewayService) SetQualifier(v string) *GatewayService {
	s.Qualifier = &v
	return s
}

func (s *GatewayService) SetSourceType(v string) *GatewayService {
	s.SourceType = &v
	return s
}

func (s *GatewayService) SetUnhealthyEndpoints(v []*string) *GatewayService {
	s.UnhealthyEndpoints = v
	return s
}

func (s *GatewayService) SetUpdateTimestamp(v int64) *GatewayService {
	s.UpdateTimestamp = &v
	return s
}

type GatewayServicePorts struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Port     *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GatewayServicePorts) String() string {
	return tea.Prettify(s)
}

func (s GatewayServicePorts) GoString() string {
	return s.String()
}

func (s *GatewayServicePorts) SetName(v string) *GatewayServicePorts {
	s.Name = &v
	return s
}

func (s *GatewayServicePorts) SetPort(v int32) *GatewayServicePorts {
	s.Port = &v
	return s
}

func (s *GatewayServicePorts) SetProtocol(v string) *GatewayServicePorts {
	s.Protocol = &v
	return s
}

type GatewayServiceSource struct {
	Bound                  *bool                                       `json:"bound,omitempty" xml:"bound,omitempty"`
	CreateTimestamp        *int64                                      `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	K8sServiceSourceInfo   *GatewayServiceSourceK8sServiceSourceInfo   `json:"k8sServiceSourceInfo,omitempty" xml:"k8sServiceSourceInfo,omitempty" type:"Struct"`
	NacosServiceSourceInfo *GatewayServiceSourceNacosServiceSourceInfo `json:"nacosServiceSourceInfo,omitempty" xml:"nacosServiceSourceInfo,omitempty" type:"Struct"`
	Name                   *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	ServiceSourceId        *string                                     `json:"serviceSourceId,omitempty" xml:"serviceSourceId,omitempty"`
	Type                   *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTimestamp        *int64                                      `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GatewayServiceSource) String() string {
	return tea.Prettify(s)
}

func (s GatewayServiceSource) GoString() string {
	return s.String()
}

func (s *GatewayServiceSource) SetBound(v bool) *GatewayServiceSource {
	s.Bound = &v
	return s
}

func (s *GatewayServiceSource) SetCreateTimestamp(v int64) *GatewayServiceSource {
	s.CreateTimestamp = &v
	return s
}

func (s *GatewayServiceSource) SetK8sServiceSourceInfo(v *GatewayServiceSourceK8sServiceSourceInfo) *GatewayServiceSource {
	s.K8sServiceSourceInfo = v
	return s
}

func (s *GatewayServiceSource) SetNacosServiceSourceInfo(v *GatewayServiceSourceNacosServiceSourceInfo) *GatewayServiceSource {
	s.NacosServiceSourceInfo = v
	return s
}

func (s *GatewayServiceSource) SetName(v string) *GatewayServiceSource {
	s.Name = &v
	return s
}

func (s *GatewayServiceSource) SetServiceSourceId(v string) *GatewayServiceSource {
	s.ServiceSourceId = &v
	return s
}

func (s *GatewayServiceSource) SetType(v string) *GatewayServiceSource {
	s.Type = &v
	return s
}

func (s *GatewayServiceSource) SetUpdateTimestamp(v int64) *GatewayServiceSource {
	s.UpdateTimestamp = &v
	return s
}

type GatewayServiceSourceK8sServiceSourceInfo struct {
	ClusterId     *string                                                `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	IngressConfig *GatewayServiceSourceK8sServiceSourceInfoIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
}

func (s GatewayServiceSourceK8sServiceSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayServiceSourceK8sServiceSourceInfo) GoString() string {
	return s.String()
}

func (s *GatewayServiceSourceK8sServiceSourceInfo) SetClusterId(v string) *GatewayServiceSourceK8sServiceSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *GatewayServiceSourceK8sServiceSourceInfo) SetIngressConfig(v *GatewayServiceSourceK8sServiceSourceInfoIngressConfig) *GatewayServiceSourceK8sServiceSourceInfo {
	s.IngressConfig = v
	return s
}

type GatewayServiceSourceK8sServiceSourceInfoIngressConfig struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s GatewayServiceSourceK8sServiceSourceInfoIngressConfig) String() string {
	return tea.Prettify(s)
}

func (s GatewayServiceSourceK8sServiceSourceInfoIngressConfig) GoString() string {
	return s.String()
}

func (s *GatewayServiceSourceK8sServiceSourceInfoIngressConfig) SetEnable(v bool) *GatewayServiceSourceK8sServiceSourceInfoIngressConfig {
	s.Enable = &v
	return s
}

func (s *GatewayServiceSourceK8sServiceSourceInfoIngressConfig) SetIngressClass(v string) *GatewayServiceSourceK8sServiceSourceInfoIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *GatewayServiceSourceK8sServiceSourceInfoIngressConfig) SetOverrideIngressIp(v bool) *GatewayServiceSourceK8sServiceSourceInfoIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *GatewayServiceSourceK8sServiceSourceInfoIngressConfig) SetWatchNamespace(v string) *GatewayServiceSourceK8sServiceSourceInfoIngressConfig {
	s.WatchNamespace = &v
	return s
}

type GatewayServiceSourceNacosServiceSourceInfo struct {
	Address    *string `json:"address,omitempty" xml:"address,omitempty"`
	ClusterId  *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GatewayServiceSourceNacosServiceSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayServiceSourceNacosServiceSourceInfo) GoString() string {
	return s.String()
}

func (s *GatewayServiceSourceNacosServiceSourceInfo) SetAddress(v string) *GatewayServiceSourceNacosServiceSourceInfo {
	s.Address = &v
	return s
}

func (s *GatewayServiceSourceNacosServiceSourceInfo) SetClusterId(v string) *GatewayServiceSourceNacosServiceSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *GatewayServiceSourceNacosServiceSourceInfo) SetInstanceId(v string) *GatewayServiceSourceNacosServiceSourceInfo {
	s.InstanceId = &v
	return s
}

type HttpApiApiInfo struct {
	// example:
	//
	// /v1
	BasePath     *string                       `json:"basePath,omitempty" xml:"basePath,omitempty"`
	Description  *string                       `json:"description,omitempty" xml:"description,omitempty"`
	Environments []*HttpApiApiInfoEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// example:
	//
	// test
	Name        *string             `json:"name,omitempty" xml:"name,omitempty"`
	Protocols   []*string           `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	VersionInfo *HttpApiVersionInfo `json:"versionInfo,omitempty" xml:"versionInfo,omitempty"`
}

func (s HttpApiApiInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfo) SetBasePath(v string) *HttpApiApiInfo {
	s.BasePath = &v
	return s
}

func (s *HttpApiApiInfo) SetDescription(v string) *HttpApiApiInfo {
	s.Description = &v
	return s
}

func (s *HttpApiApiInfo) SetEnvironments(v []*HttpApiApiInfoEnvironments) *HttpApiApiInfo {
	s.Environments = v
	return s
}

func (s *HttpApiApiInfo) SetHttpApiId(v string) *HttpApiApiInfo {
	s.HttpApiId = &v
	return s
}

func (s *HttpApiApiInfo) SetName(v string) *HttpApiApiInfo {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfo) SetProtocols(v []*string) *HttpApiApiInfo {
	s.Protocols = v
	return s
}

func (s *HttpApiApiInfo) SetVersionInfo(v *HttpApiVersionInfo) *HttpApiApiInfo {
	s.VersionInfo = v
	return s
}

type HttpApiApiInfoEnvironments struct {
	// example:
	//
	// test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// example:
	//
	// Service
	BackendType        *string                                       `json:"backendType,omitempty" xml:"backendType,omitempty"`
	CloudProductConfig *HttpApiApiInfoEnvironmentsCloudProductConfig `json:"cloudProductConfig,omitempty" xml:"cloudProductConfig,omitempty" type:"Struct"`
	CustomDomains      []*HttpApiDomainInfo                          `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	DnsConfigs         []*HttpApiApiInfoEnvironmentsDnsConfigs       `json:"dnsConfigs,omitempty" xml:"dnsConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentId *string                                `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpApiApiInfoEnvironmentsGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Published
	PublishStatus  *string                                     `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	ServiceConfigs []*HttpApiApiInfoEnvironmentsServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	VipConfigs     []*HttpApiApiInfoEnvironmentsVipConfigs     `json:"vipConfigs,omitempty" xml:"vipConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiApiInfoEnvironments) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironments) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironments) SetAlias(v string) *HttpApiApiInfoEnvironments {
	s.Alias = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetBackendScene(v string) *HttpApiApiInfoEnvironments {
	s.BackendScene = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetBackendType(v string) *HttpApiApiInfoEnvironments {
	s.BackendType = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetCloudProductConfig(v *HttpApiApiInfoEnvironmentsCloudProductConfig) *HttpApiApiInfoEnvironments {
	s.CloudProductConfig = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiApiInfoEnvironments {
	s.CustomDomains = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetDnsConfigs(v []*HttpApiApiInfoEnvironmentsDnsConfigs) *HttpApiApiInfoEnvironments {
	s.DnsConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetEnvironmentId(v string) *HttpApiApiInfoEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetGatewayInfo(v *HttpApiApiInfoEnvironmentsGatewayInfo) *HttpApiApiInfoEnvironments {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetName(v string) *HttpApiApiInfoEnvironments {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetPublishStatus(v string) *HttpApiApiInfoEnvironments {
	s.PublishStatus = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetServiceConfigs(v []*HttpApiApiInfoEnvironmentsServiceConfigs) *HttpApiApiInfoEnvironments {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetVipConfigs(v []*HttpApiApiInfoEnvironmentsVipConfigs) *HttpApiApiInfoEnvironments {
	s.VipConfigs = v
	return s
}

type HttpApiApiInfoEnvironmentsCloudProductConfig struct {
	// example:
	//
	// CS
	CloudProductType        *string                                                                `json:"cloudProductType,omitempty" xml:"cloudProductType,omitempty"`
	ContainerServiceConfigs []*HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs `json:"containerServiceConfigs,omitempty" xml:"containerServiceConfigs,omitempty" type:"Repeated"`
	FunctionConfigs         []*HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs         `json:"functionConfigs,omitempty" xml:"functionConfigs,omitempty" type:"Repeated"`
	MseNacosConfigs         []*HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs         `json:"mseNacosConfigs,omitempty" xml:"mseNacosConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfig) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfig) SetCloudProductType(v string) *HttpApiApiInfoEnvironmentsCloudProductConfig {
	s.CloudProductType = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfig) SetContainerServiceConfigs(v []*HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) *HttpApiApiInfoEnvironmentsCloudProductConfig {
	s.ContainerServiceConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfig) SetFunctionConfigs(v []*HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) *HttpApiApiInfoEnvironmentsCloudProductConfig {
	s.FunctionConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfig) SetMseNacosConfigs(v []*HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) *HttpApiApiInfoEnvironmentsCloudProductConfig {
	s.MseNacosConfigs = v
	return s
}

type HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetGatewayServiceId(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetName(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetNamespace(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetPort(v int32) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetProtocol(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsCloudProductConfigContainerServiceConfigs {
	s.Weight = &v
	return s
}

type HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// test-function
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) SetGatewayServiceId(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) SetName(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) SetQualifier(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs {
	s.Qualifier = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsCloudProductConfigFunctionConfigs {
	s.Weight = &v
	return s
}

type HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string                        `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// springboot-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetGatewayServiceId(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetGroupName(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.GroupName = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetName(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetNamespace(v string) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsCloudProductConfigMseNacosConfigs {
	s.Weight = &v
	return s
}

type HttpApiApiInfoEnvironmentsDnsConfigs struct {
	DnsList []*string                      `json:"dnsList,omitempty" xml:"dnsList,omitempty" type:"Repeated"`
	Match   *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsDnsConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsDnsConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsDnsConfigs) SetDnsList(v []*string) *HttpApiApiInfoEnvironmentsDnsConfigs {
	s.DnsList = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsDnsConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsDnsConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsDnsConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsDnsConfigs {
	s.Weight = &v
	return s
}

type HttpApiApiInfoEnvironmentsGatewayInfo struct {
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsGatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) SetGatewayId(v string) *HttpApiApiInfoEnvironmentsGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) SetName(v string) *HttpApiApiInfoEnvironmentsGatewayInfo {
	s.Name = &v
	return s
}

type HttpApiApiInfoEnvironmentsServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 8080
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetGatewayServiceId(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetName(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetPort(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetProtocol(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetVersion(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Weight = &v
	return s
}

type HttpApiApiInfoEnvironmentsVipConfigs struct {
	Endpoints []*string                      `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsVipConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsVipConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsVipConfigs) SetEndpoints(v []*string) *HttpApiApiInfoEnvironmentsVipConfigs {
	s.Endpoints = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsVipConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsVipConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsVipConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsVipConfigs {
	s.Weight = &v
	return s
}

type HttpApiBackendMatchCondition struct {
	// example:
	//
	// color
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// Query
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// gray
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpApiBackendMatchCondition) String() string {
	return tea.Prettify(s)
}

func (s HttpApiBackendMatchCondition) GoString() string {
	return s.String()
}

func (s *HttpApiBackendMatchCondition) SetKey(v string) *HttpApiBackendMatchCondition {
	s.Key = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetOperator(v string) *HttpApiBackendMatchCondition {
	s.Operator = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetType(v string) *HttpApiBackendMatchCondition {
	s.Type = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetValue(v string) *HttpApiBackendMatchCondition {
	s.Value = &v
	return s
}

type HttpApiBackendMatchConditions struct {
	Conditions []*HttpApiBackendMatchCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
}

func (s HttpApiBackendMatchConditions) String() string {
	return tea.Prettify(s)
}

func (s HttpApiBackendMatchConditions) GoString() string {
	return s.String()
}

func (s *HttpApiBackendMatchConditions) SetConditions(v []*HttpApiBackendMatchCondition) *HttpApiBackendMatchConditions {
	s.Conditions = v
	return s
}

func (s *HttpApiBackendMatchConditions) SetDefault(v bool) *HttpApiBackendMatchConditions {
	s.Default = &v
	return s
}

type HttpApiDomainInfo struct {
	// example:
	//
	// d-xxx
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// example:
	//
	// www.example.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiDomainInfo) GoString() string {
	return s.String()
}

func (s *HttpApiDomainInfo) SetDomainId(v string) *HttpApiDomainInfo {
	s.DomainId = &v
	return s
}

func (s *HttpApiDomainInfo) SetName(v string) *HttpApiDomainInfo {
	s.Name = &v
	return s
}

func (s *HttpApiDomainInfo) SetProtocol(v string) *HttpApiDomainInfo {
	s.Protocol = &v
	return s
}

type HttpApiInfoByName struct {
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	VersionEnabled    *bool             `json:"versionEnabled,omitempty" xml:"versionEnabled,omitempty"`
	VersionedHttpApis []*HttpApiApiInfo `json:"versionedHttpApis,omitempty" xml:"versionedHttpApis,omitempty" type:"Repeated"`
}

func (s HttpApiInfoByName) String() string {
	return tea.Prettify(s)
}

func (s HttpApiInfoByName) GoString() string {
	return s.String()
}

func (s *HttpApiInfoByName) SetName(v string) *HttpApiInfoByName {
	s.Name = &v
	return s
}

func (s *HttpApiInfoByName) SetVersionEnabled(v bool) *HttpApiInfoByName {
	s.VersionEnabled = &v
	return s
}

func (s *HttpApiInfoByName) SetVersionedHttpApis(v []*HttpApiApiInfo) *HttpApiInfoByName {
	s.VersionedHttpApis = v
	return s
}

type HttpApiMockContract struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 200
	ResponseCode    *int32  `json:"responseCode,omitempty" xml:"responseCode,omitempty"`
	ResponseContent *string `json:"responseContent,omitempty" xml:"responseContent,omitempty"`
}

func (s HttpApiMockContract) String() string {
	return tea.Prettify(s)
}

func (s HttpApiMockContract) GoString() string {
	return s.String()
}

func (s *HttpApiMockContract) SetEnable(v bool) *HttpApiMockContract {
	s.Enable = &v
	return s
}

func (s *HttpApiMockContract) SetResponseCode(v int32) *HttpApiMockContract {
	s.ResponseCode = &v
	return s
}

func (s *HttpApiMockContract) SetResponseContent(v string) *HttpApiMockContract {
	s.ResponseContent = &v
	return s
}

type HttpApiOperation struct {
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GET
	Method *string              `json:"method,omitempty" xml:"method,omitempty"`
	Mock   *HttpApiMockContract `json:"mock,omitempty" xml:"mock,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GetUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /user
	Path     *string                  `json:"path,omitempty" xml:"path,omitempty"`
	Request  *HttpApiRequestContract  `json:"request,omitempty" xml:"request,omitempty"`
	Response *HttpApiResponseContract `json:"response,omitempty" xml:"response,omitempty"`
}

func (s HttpApiOperation) String() string {
	return tea.Prettify(s)
}

func (s HttpApiOperation) GoString() string {
	return s.String()
}

func (s *HttpApiOperation) SetDescription(v string) *HttpApiOperation {
	s.Description = &v
	return s
}

func (s *HttpApiOperation) SetMethod(v string) *HttpApiOperation {
	s.Method = &v
	return s
}

func (s *HttpApiOperation) SetMock(v *HttpApiMockContract) *HttpApiOperation {
	s.Mock = v
	return s
}

func (s *HttpApiOperation) SetName(v string) *HttpApiOperation {
	s.Name = &v
	return s
}

func (s *HttpApiOperation) SetPath(v string) *HttpApiOperation {
	s.Path = &v
	return s
}

func (s *HttpApiOperation) SetRequest(v *HttpApiRequestContract) *HttpApiOperation {
	s.Request = v
	return s
}

func (s *HttpApiOperation) SetResponse(v *HttpApiResponseContract) *HttpApiOperation {
	s.Response = v
	return s
}

type HttpApiOperationInfo struct {
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// GET
	Method *string              `json:"method,omitempty" xml:"method,omitempty"`
	Mock   *HttpApiMockContract `json:"mock,omitempty" xml:"mock,omitempty"`
	// example:
	//
	// GetUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// op-xxx
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// example:
	//
	// /user/123
	Path     *string                  `json:"path,omitempty" xml:"path,omitempty"`
	Request  *HttpApiRequestContract  `json:"request,omitempty" xml:"request,omitempty"`
	Response *HttpApiResponseContract `json:"response,omitempty" xml:"response,omitempty"`
}

func (s HttpApiOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiOperationInfo) GoString() string {
	return s.String()
}

func (s *HttpApiOperationInfo) SetCreateTimestamp(v int64) *HttpApiOperationInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiOperationInfo) SetDescription(v string) *HttpApiOperationInfo {
	s.Description = &v
	return s
}

func (s *HttpApiOperationInfo) SetMethod(v string) *HttpApiOperationInfo {
	s.Method = &v
	return s
}

func (s *HttpApiOperationInfo) SetMock(v *HttpApiMockContract) *HttpApiOperationInfo {
	s.Mock = v
	return s
}

func (s *HttpApiOperationInfo) SetName(v string) *HttpApiOperationInfo {
	s.Name = &v
	return s
}

func (s *HttpApiOperationInfo) SetOperationId(v string) *HttpApiOperationInfo {
	s.OperationId = &v
	return s
}

func (s *HttpApiOperationInfo) SetPath(v string) *HttpApiOperationInfo {
	s.Path = &v
	return s
}

func (s *HttpApiOperationInfo) SetRequest(v *HttpApiRequestContract) *HttpApiOperationInfo {
	s.Request = v
	return s
}

func (s *HttpApiOperationInfo) SetResponse(v *HttpApiResponseContract) *HttpApiOperationInfo {
	s.Response = v
	return s
}

type HttpApiParameter struct {
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	ExampleValue *string `json:"exampleValue,omitempty" xml:"exampleValue,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiParameter) String() string {
	return tea.Prettify(s)
}

func (s HttpApiParameter) GoString() string {
	return s.String()
}

func (s *HttpApiParameter) SetDefaultValue(v string) *HttpApiParameter {
	s.DefaultValue = &v
	return s
}

func (s *HttpApiParameter) SetDescription(v string) *HttpApiParameter {
	s.Description = &v
	return s
}

func (s *HttpApiParameter) SetExampleValue(v string) *HttpApiParameter {
	s.ExampleValue = &v
	return s
}

func (s *HttpApiParameter) SetName(v string) *HttpApiParameter {
	s.Name = &v
	return s
}

func (s *HttpApiParameter) SetRequired(v bool) *HttpApiParameter {
	s.Required = &v
	return s
}

func (s *HttpApiParameter) SetType(v string) *HttpApiParameter {
	s.Type = &v
	return s
}

type HttpApiPublishRevisionInfo struct {
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// example:
	//
	// Service
	BackendType        *string                                       `json:"backendType,omitempty" xml:"backendType,omitempty"`
	CloudProductConfig *HttpApiPublishRevisionInfoCloudProductConfig `json:"cloudProductConfig,omitempty" xml:"cloudProductConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1718807057927
	CreateTimestamp *int64                                  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	CustomDomains   []*HttpApiDomainInfo                    `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	DnsConfigs      []*HttpApiPublishRevisionInfoDnsConfigs `json:"dnsConfigs,omitempty" xml:"dnsConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentInfo *HttpApiPublishRevisionInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsCurrentVersion *bool                   `json:"isCurrentVersion,omitempty" xml:"isCurrentVersion,omitempty"`
	Operations       []*HttpApiOperationInfo `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// example:
	//
	// apr-xxx
	RevisionId     *string                                     `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
	ServiceConfigs []*HttpApiPublishRevisionInfoServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx.com
	SubDomains []*HttpApiDomainInfo                    `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
	VipConfigs []*HttpApiPublishRevisionInfoVipConfigs `json:"vipConfigs,omitempty" xml:"vipConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfo) SetBackendScene(v string) *HttpApiPublishRevisionInfo {
	s.BackendScene = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetBackendType(v string) *HttpApiPublishRevisionInfo {
	s.BackendType = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCloudProductConfig(v *HttpApiPublishRevisionInfoCloudProductConfig) *HttpApiPublishRevisionInfo {
	s.CloudProductConfig = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCreateTimestamp(v int64) *HttpApiPublishRevisionInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.CustomDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetDnsConfigs(v []*HttpApiPublishRevisionInfoDnsConfigs) *HttpApiPublishRevisionInfo {
	s.DnsConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetEnvironmentInfo(v *HttpApiPublishRevisionInfoEnvironmentInfo) *HttpApiPublishRevisionInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetIsCurrentVersion(v bool) *HttpApiPublishRevisionInfo {
	s.IsCurrentVersion = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetOperations(v []*HttpApiOperationInfo) *HttpApiPublishRevisionInfo {
	s.Operations = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetRevisionId(v string) *HttpApiPublishRevisionInfo {
	s.RevisionId = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetServiceConfigs(v []*HttpApiPublishRevisionInfoServiceConfigs) *HttpApiPublishRevisionInfo {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetSubDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.SubDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetVipConfigs(v []*HttpApiPublishRevisionInfoVipConfigs) *HttpApiPublishRevisionInfo {
	s.VipConfigs = v
	return s
}

type HttpApiPublishRevisionInfoCloudProductConfig struct {
	// example:
	//
	// FC
	CloudProductType        *string                                                                `json:"cloudProductType,omitempty" xml:"cloudProductType,omitempty"`
	ContainerServiceConfigs []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs `json:"containerServiceConfigs,omitempty" xml:"containerServiceConfigs,omitempty" type:"Repeated"`
	FunctionConfigs         []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs         `json:"functionConfigs,omitempty" xml:"functionConfigs,omitempty" type:"Repeated"`
	MseNacosConfigs         []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs         `json:"mseNacosConfigs,omitempty" xml:"mseNacosConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetCloudProductType(v string) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.CloudProductType = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetContainerServiceConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.ContainerServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetFunctionConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.FunctionConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetMseNacosConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.MseNacosConfigs = v
	return s
}

type HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 100
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetWeight(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Weight = &v
	return s
}

type HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// demo-function
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetQualifier(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Qualifier = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Weight = &v
	return s
}

type HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string                        `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// spring-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGroupName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GroupName = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Weight = &v
	return s
}

type HttpApiPublishRevisionInfoDnsConfigs struct {
	DnsList []*string                      `json:"dnsList,omitempty" xml:"dnsList,omitempty" type:"Repeated"`
	Match   *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoDnsConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoDnsConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetDnsList(v []*string) *HttpApiPublishRevisionInfoDnsConfigs {
	s.DnsList = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Weight = &v
	return s
}

type HttpApiPublishRevisionInfoEnvironmentInfo struct {
	// example:
	//
	// 测试
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// env-xxx
	EnvironmentId *string                                               `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetAlias(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Alias = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetEnvironmentId(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetGatewayInfo(v *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Name = &v
	return s
}

type HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo struct {
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// 实例1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetGatewayId(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.Name = &v
	return s
}

type HttpApiPublishRevisionInfoServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetVersion(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Weight = &v
	return s
}

type HttpApiPublishRevisionInfoVipConfigs struct {
	Endpoints []*string                      `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoVipConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiPublishRevisionInfoVipConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetEndpoints(v []*string) *HttpApiPublishRevisionInfoVipConfigs {
	s.Endpoints = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoVipConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoVipConfigs {
	s.Weight = &v
	return s
}

type HttpApiRequestContract struct {
	Body             *HttpApiRequestContractBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	HeaderParameters []*HttpApiParameter         `json:"headerParameters,omitempty" xml:"headerParameters,omitempty" type:"Repeated"`
	PathParameters   []*HttpApiParameter         `json:"pathParameters,omitempty" xml:"pathParameters,omitempty" type:"Repeated"`
	QueryParameters  []*HttpApiParameter         `json:"queryParameters,omitempty" xml:"queryParameters,omitempty" type:"Repeated"`
}

func (s HttpApiRequestContract) String() string {
	return tea.Prettify(s)
}

func (s HttpApiRequestContract) GoString() string {
	return s.String()
}

func (s *HttpApiRequestContract) SetBody(v *HttpApiRequestContractBody) *HttpApiRequestContract {
	s.Body = v
	return s
}

func (s *HttpApiRequestContract) SetHeaderParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.HeaderParameters = v
	return s
}

func (s *HttpApiRequestContract) SetPathParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.PathParameters = v
	return s
}

func (s *HttpApiRequestContract) SetQueryParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.QueryParameters = v
	return s
}

type HttpApiRequestContractBody struct {
	// example:
	//
	// application/json
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"key":"value"}
	Example    *string `json:"example,omitempty" xml:"example,omitempty"`
	JsonSchema *string `json:"jsonSchema,omitempty" xml:"jsonSchema,omitempty"`
}

func (s HttpApiRequestContractBody) String() string {
	return tea.Prettify(s)
}

func (s HttpApiRequestContractBody) GoString() string {
	return s.String()
}

func (s *HttpApiRequestContractBody) SetContentType(v string) *HttpApiRequestContractBody {
	s.ContentType = &v
	return s
}

func (s *HttpApiRequestContractBody) SetDescription(v string) *HttpApiRequestContractBody {
	s.Description = &v
	return s
}

func (s *HttpApiRequestContractBody) SetExample(v string) *HttpApiRequestContractBody {
	s.Example = &v
	return s
}

func (s *HttpApiRequestContractBody) SetJsonSchema(v string) *HttpApiRequestContractBody {
	s.JsonSchema = &v
	return s
}

type HttpApiResponseContract struct {
	// This parameter is required.
	//
	// example:
	//
	// application/json
	ContentType *string                         `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Items       []*HttpApiResponseContractItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s HttpApiResponseContract) String() string {
	return tea.Prettify(s)
}

func (s HttpApiResponseContract) GoString() string {
	return s.String()
}

func (s *HttpApiResponseContract) SetContentType(v string) *HttpApiResponseContract {
	s.ContentType = &v
	return s
}

func (s *HttpApiResponseContract) SetItems(v []*HttpApiResponseContractItems) *HttpApiResponseContract {
	s.Items = v
	return s
}

type HttpApiResponseContractItems struct {
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 正常接口响应
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"result": "ok"}
	Example    *string `json:"example,omitempty" xml:"example,omitempty"`
	JsonSchema *string `json:"jsonSchema,omitempty" xml:"jsonSchema,omitempty"`
}

func (s HttpApiResponseContractItems) String() string {
	return tea.Prettify(s)
}

func (s HttpApiResponseContractItems) GoString() string {
	return s.String()
}

func (s *HttpApiResponseContractItems) SetCode(v int32) *HttpApiResponseContractItems {
	s.Code = &v
	return s
}

func (s *HttpApiResponseContractItems) SetDescription(v string) *HttpApiResponseContractItems {
	s.Description = &v
	return s
}

func (s *HttpApiResponseContractItems) SetExample(v string) *HttpApiResponseContractItems {
	s.Example = &v
	return s
}

func (s *HttpApiResponseContractItems) SetJsonSchema(v string) *HttpApiResponseContractItems {
	s.JsonSchema = &v
	return s
}

type HttpApiVersionConfig struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// my-version
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// example:
	//
	// myVersion
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// example:
	//
	// Query
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HttpApiVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpApiVersionConfig) GoString() string {
	return s.String()
}

func (s *HttpApiVersionConfig) SetEnable(v bool) *HttpApiVersionConfig {
	s.Enable = &v
	return s
}

func (s *HttpApiVersionConfig) SetHeaderName(v string) *HttpApiVersionConfig {
	s.HeaderName = &v
	return s
}

func (s *HttpApiVersionConfig) SetQueryName(v string) *HttpApiVersionConfig {
	s.QueryName = &v
	return s
}

func (s *HttpApiVersionConfig) SetScheme(v string) *HttpApiVersionConfig {
	s.Scheme = &v
	return s
}

func (s *HttpApiVersionConfig) SetVersion(v string) *HttpApiVersionConfig {
	s.Version = &v
	return s
}

type HttpApiVersionInfo struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// my-version
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// example:
	//
	// myVersion
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// example:
	//
	// Query
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HttpApiVersionInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiVersionInfo) GoString() string {
	return s.String()
}

func (s *HttpApiVersionInfo) SetEnable(v bool) *HttpApiVersionInfo {
	s.Enable = &v
	return s
}

func (s *HttpApiVersionInfo) SetHeaderName(v string) *HttpApiVersionInfo {
	s.HeaderName = &v
	return s
}

func (s *HttpApiVersionInfo) SetQueryName(v string) *HttpApiVersionInfo {
	s.QueryName = &v
	return s
}

func (s *HttpApiVersionInfo) SetScheme(v string) *HttpApiVersionInfo {
	s.Scheme = &v
	return s
}

func (s *HttpApiVersionInfo) SetVersion(v string) *HttpApiVersionInfo {
	s.Version = &v
	return s
}

type HttpDubboTranscoder struct {
	DubboServiceGroup   *string                             `json:"dubboServiceGroup,omitempty" xml:"dubboServiceGroup,omitempty"`
	DubboServiceName    *string                             `json:"dubboServiceName,omitempty" xml:"dubboServiceName,omitempty"`
	DubboServiceVersion *string                             `json:"dubboServiceVersion,omitempty" xml:"dubboServiceVersion,omitempty"`
	MothedMapList       []*HttpDubboTranscoderMothedMapList `json:"mothedMapList,omitempty" xml:"mothedMapList,omitempty" type:"Repeated"`
}

func (s HttpDubboTranscoder) String() string {
	return tea.Prettify(s)
}

func (s HttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoder) SetDubboServiceGroup(v string) *HttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *HttpDubboTranscoder) SetDubboServiceName(v string) *HttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *HttpDubboTranscoder) SetDubboServiceVersion(v string) *HttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *HttpDubboTranscoder) SetMothedMapList(v []*HttpDubboTranscoderMothedMapList) *HttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

type HttpDubboTranscoderMothedMapList struct {
	DubboMothedName *string `json:"dubboMothedName,omitempty" xml:"dubboMothedName,omitempty"`
	// example:
	//
	// ALL_GET
	HttpMothed *string `json:"httpMothed,omitempty" xml:"httpMothed,omitempty"`
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath    *string                                          `json:"mothedpath,omitempty" xml:"mothedpath,omitempty"`
	ParamMapsList []*HttpDubboTranscoderMothedMapListParamMapsList `json:"paramMapsList,omitempty" xml:"paramMapsList,omitempty" type:"Repeated"`
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string   `json:"passThroughAllHeaders,omitempty" xml:"passThroughAllHeaders,omitempty"`
	PassThroughList       []*string `json:"passThroughList,omitempty" xml:"passThroughList,omitempty" type:"Repeated"`
}

func (s HttpDubboTranscoderMothedMapList) String() string {
	return tea.Prettify(s)
}

func (s HttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *HttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *HttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetMothedpath(v string) *HttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetParamMapsList(v []*HttpDubboTranscoderMothedMapListParamMapsList) *HttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *HttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *HttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

type HttpDubboTranscoderMothedMapListParamMapsList struct {
	// example:
	//
	// name
	ExtractKey *string `json:"extractKey,omitempty" xml:"extractKey,omitempty"`
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"extractKeySpec,omitempty" xml:"extractKeySpec,omitempty"`
	// example:
	//
	// java.lang.String
	MappingType *string `json:"mappingType,omitempty" xml:"mappingType,omitempty"`
}

func (s HttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return tea.Prettify(s)
}

func (s HttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

type HttpRouteMatch struct {
	Headers []*HttpRouteMatchHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IgnoreUriCase *bool                        `json:"ignoreUriCase,omitempty" xml:"ignoreUriCase,omitempty"`
	Methods       []*string                    `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	Path          *HttpRouteMatchPath          `json:"path,omitempty" xml:"path,omitempty" type:"Struct"`
	QueryParams   []*HttpRouteMatchQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
}

func (s HttpRouteMatch) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteMatch) GoString() string {
	return s.String()
}

func (s *HttpRouteMatch) SetHeaders(v []*HttpRouteMatchHeaders) *HttpRouteMatch {
	s.Headers = v
	return s
}

func (s *HttpRouteMatch) SetIgnoreUriCase(v bool) *HttpRouteMatch {
	s.IgnoreUriCase = &v
	return s
}

func (s *HttpRouteMatch) SetMethods(v []*string) *HttpRouteMatch {
	s.Methods = v
	return s
}

func (s *HttpRouteMatch) SetPath(v *HttpRouteMatchPath) *HttpRouteMatch {
	s.Path = v
	return s
}

func (s *HttpRouteMatch) SetQueryParams(v []*HttpRouteMatchQueryParams) *HttpRouteMatch {
	s.QueryParams = v
	return s
}

type HttpRouteMatchHeaders struct {
	// example:
	//
	// dev
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Exact
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteMatchHeaders) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchHeaders) SetName(v string) *HttpRouteMatchHeaders {
	s.Name = &v
	return s
}

func (s *HttpRouteMatchHeaders) SetType(v string) *HttpRouteMatchHeaders {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchHeaders) SetValue(v string) *HttpRouteMatchHeaders {
	s.Value = &v
	return s
}

type HttpRouteMatchPath struct {
	// example:
	//
	// Prefix
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// /user
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchPath) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteMatchPath) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchPath) SetType(v string) *HttpRouteMatchPath {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchPath) SetValue(v string) *HttpRouteMatchPath {
	s.Value = &v
	return s
}

type HttpRouteMatchQueryParams struct {
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Exact
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 17
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchQueryParams) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteMatchQueryParams) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchQueryParams) SetName(v string) *HttpRouteMatchQueryParams {
	s.Name = &v
	return s
}

func (s *HttpRouteMatchQueryParams) SetType(v string) *HttpRouteMatchQueryParams {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchQueryParams) SetValue(v string) *HttpRouteMatchQueryParams {
	s.Value = &v
	return s
}

type PolicyClassInfo struct {
	Alias                   *string   `json:"alias,omitempty" xml:"alias,omitempty"`
	AttachableResourceTypes []*string `json:"attachableResourceTypes,omitempty" xml:"attachableResourceTypes,omitempty" type:"Repeated"`
	ClassId                 *string   `json:"classId,omitempty" xml:"classId,omitempty"`
	ConfigExample           *string   `json:"configExample,omitempty" xml:"configExample,omitempty"`
	Deprecated              *bool     `json:"deprecated,omitempty" xml:"deprecated,omitempty"`
	Description             *string   `json:"description,omitempty" xml:"description,omitempty"`
	Direction               *string   `json:"direction,omitempty" xml:"direction,omitempty"`
	EnableLog               *bool     `json:"enableLog,omitempty" xml:"enableLog,omitempty"`
	ExecutePriority         *string   `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	ExecuteStage            *string   `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	Name                    *string   `json:"name,omitempty" xml:"name,omitempty"`
	Type                    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Version                 *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PolicyClassInfo) String() string {
	return tea.Prettify(s)
}

func (s PolicyClassInfo) GoString() string {
	return s.String()
}

func (s *PolicyClassInfo) SetAlias(v string) *PolicyClassInfo {
	s.Alias = &v
	return s
}

func (s *PolicyClassInfo) SetAttachableResourceTypes(v []*string) *PolicyClassInfo {
	s.AttachableResourceTypes = v
	return s
}

func (s *PolicyClassInfo) SetClassId(v string) *PolicyClassInfo {
	s.ClassId = &v
	return s
}

func (s *PolicyClassInfo) SetConfigExample(v string) *PolicyClassInfo {
	s.ConfigExample = &v
	return s
}

func (s *PolicyClassInfo) SetDeprecated(v bool) *PolicyClassInfo {
	s.Deprecated = &v
	return s
}

func (s *PolicyClassInfo) SetDescription(v string) *PolicyClassInfo {
	s.Description = &v
	return s
}

func (s *PolicyClassInfo) SetDirection(v string) *PolicyClassInfo {
	s.Direction = &v
	return s
}

func (s *PolicyClassInfo) SetEnableLog(v bool) *PolicyClassInfo {
	s.EnableLog = &v
	return s
}

func (s *PolicyClassInfo) SetExecutePriority(v string) *PolicyClassInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PolicyClassInfo) SetExecuteStage(v string) *PolicyClassInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PolicyClassInfo) SetName(v string) *PolicyClassInfo {
	s.Name = &v
	return s
}

func (s *PolicyClassInfo) SetType(v string) *PolicyClassInfo {
	s.Type = &v
	return s
}

func (s *PolicyClassInfo) SetVersion(v string) *PolicyClassInfo {
	s.Version = &v
	return s
}

type PolicyDetailInfo struct {
	ClassId     *string `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName   *string `json:"className,omitempty" xml:"className,omitempty"`
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	PolicyId    *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s PolicyDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s PolicyDetailInfo) GoString() string {
	return s.String()
}

func (s *PolicyDetailInfo) SetClassId(v string) *PolicyDetailInfo {
	s.ClassId = &v
	return s
}

func (s *PolicyDetailInfo) SetClassName(v string) *PolicyDetailInfo {
	s.ClassName = &v
	return s
}

func (s *PolicyDetailInfo) SetConfig(v string) *PolicyDetailInfo {
	s.Config = &v
	return s
}

func (s *PolicyDetailInfo) SetDescription(v string) *PolicyDetailInfo {
	s.Description = &v
	return s
}

func (s *PolicyDetailInfo) SetName(v string) *PolicyDetailInfo {
	s.Name = &v
	return s
}

func (s *PolicyDetailInfo) SetPolicyId(v string) *PolicyDetailInfo {
	s.PolicyId = &v
	return s
}

type PolicyInfo struct {
	Attachments     []*Attachment `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	ClassAlias      *string       `json:"classAlias,omitempty" xml:"classAlias,omitempty"`
	ClassName       *string       `json:"className,omitempty" xml:"className,omitempty"`
	Config          *string       `json:"config,omitempty" xml:"config,omitempty"`
	Direction       *string       `json:"direction,omitempty" xml:"direction,omitempty"`
	ExecutePriority *string       `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	ExecuteStage    *string       `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	Name            *string       `json:"name,omitempty" xml:"name,omitempty"`
	PolicyId        *string       `json:"policyId,omitempty" xml:"policyId,omitempty"`
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PolicyInfo) String() string {
	return tea.Prettify(s)
}

func (s PolicyInfo) GoString() string {
	return s.String()
}

func (s *PolicyInfo) SetAttachments(v []*Attachment) *PolicyInfo {
	s.Attachments = v
	return s
}

func (s *PolicyInfo) SetClassAlias(v string) *PolicyInfo {
	s.ClassAlias = &v
	return s
}

func (s *PolicyInfo) SetClassName(v string) *PolicyInfo {
	s.ClassName = &v
	return s
}

func (s *PolicyInfo) SetConfig(v string) *PolicyInfo {
	s.Config = &v
	return s
}

func (s *PolicyInfo) SetDirection(v string) *PolicyInfo {
	s.Direction = &v
	return s
}

func (s *PolicyInfo) SetExecutePriority(v string) *PolicyInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PolicyInfo) SetExecuteStage(v string) *PolicyInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PolicyInfo) SetName(v string) *PolicyInfo {
	s.Name = &v
	return s
}

func (s *PolicyInfo) SetPolicyId(v string) *PolicyInfo {
	s.PolicyId = &v
	return s
}

func (s *PolicyInfo) SetType(v string) *PolicyInfo {
	s.Type = &v
	return s
}

type ResourceStatistic struct {
	ResourceCount *int32  `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
	ResourceType  *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ResourceStatistic) String() string {
	return tea.Prettify(s)
}

func (s ResourceStatistic) GoString() string {
	return s.String()
}

func (s *ResourceStatistic) SetResourceCount(v int32) *ResourceStatistic {
	s.ResourceCount = &v
	return s
}

func (s *ResourceStatistic) SetResourceType(v string) *ResourceStatistic {
	s.ResourceType = &v
	return s
}

type RouteRulesConflictInfo struct {
	Conflicts  []*RouteRulesConflictInfoConflicts `json:"conflicts,omitempty" xml:"conflicts,omitempty" type:"Repeated"`
	DomainInfo *RouteRulesConflictInfoDomainInfo  `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Struct"`
}

func (s RouteRulesConflictInfo) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfo) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfo) SetConflicts(v []*RouteRulesConflictInfoConflicts) *RouteRulesConflictInfo {
	s.Conflicts = v
	return s
}

func (s *RouteRulesConflictInfo) SetDomainInfo(v *RouteRulesConflictInfoDomainInfo) *RouteRulesConflictInfo {
	s.DomainInfo = v
	return s
}

type RouteRulesConflictInfoConflicts struct {
	Details         []*RouteRulesConflictInfoConflictsDetails       `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	EnvironmentInfo *RouteRulesConflictInfoConflictsEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	ResourceId      *string                                         `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName    *string                                         `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType    *string                                         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RouteRulesConflictInfoConflicts) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflicts) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflicts) SetDetails(v []*RouteRulesConflictInfoConflictsDetails) *RouteRulesConflictInfoConflicts {
	s.Details = v
	return s
}

func (s *RouteRulesConflictInfoConflicts) SetEnvironmentInfo(v *RouteRulesConflictInfoConflictsEnvironmentInfo) *RouteRulesConflictInfoConflicts {
	s.EnvironmentInfo = v
	return s
}

func (s *RouteRulesConflictInfoConflicts) SetResourceId(v string) *RouteRulesConflictInfoConflicts {
	s.ResourceId = &v
	return s
}

func (s *RouteRulesConflictInfoConflicts) SetResourceName(v string) *RouteRulesConflictInfoConflicts {
	s.ResourceName = &v
	return s
}

func (s *RouteRulesConflictInfoConflicts) SetResourceType(v string) *RouteRulesConflictInfoConflicts {
	s.ResourceType = &v
	return s
}

type RouteRulesConflictInfoConflictsDetails struct {
	ConflictingMatch *RouteRulesConflictInfoConflictsDetailsConflictingMatch `json:"conflictingMatch,omitempty" xml:"conflictingMatch,omitempty" type:"Struct"`
	DetectedMatch    *RouteRulesConflictInfoConflictsDetailsDetectedMatch    `json:"detectedMatch,omitempty" xml:"detectedMatch,omitempty" type:"Struct"`
	Level            *string                                                 `json:"level,omitempty" xml:"level,omitempty"`
}

func (s RouteRulesConflictInfoConflictsDetails) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsDetails) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsDetails) SetConflictingMatch(v *RouteRulesConflictInfoConflictsDetailsConflictingMatch) *RouteRulesConflictInfoConflictsDetails {
	s.ConflictingMatch = v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetails) SetDetectedMatch(v *RouteRulesConflictInfoConflictsDetailsDetectedMatch) *RouteRulesConflictInfoConflictsDetails {
	s.DetectedMatch = v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetails) SetLevel(v string) *RouteRulesConflictInfoConflictsDetails {
	s.Level = &v
	return s
}

type RouteRulesConflictInfoConflictsDetailsConflictingMatch struct {
	Match         *HttpRouteMatch                                                      `json:"match,omitempty" xml:"match,omitempty"`
	OperationInfo *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s RouteRulesConflictInfoConflictsDetailsConflictingMatch) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsDetailsConflictingMatch) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsDetailsConflictingMatch) SetMatch(v *HttpRouteMatch) *RouteRulesConflictInfoConflictsDetailsConflictingMatch {
	s.Match = v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetailsConflictingMatch) SetOperationInfo(v *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo) *RouteRulesConflictInfoConflictsDetailsConflictingMatch {
	s.OperationInfo = v
	return s
}

type RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetName(v string) *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetOperationId(v string) *RouteRulesConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.OperationId = &v
	return s
}

type RouteRulesConflictInfoConflictsDetailsDetectedMatch struct {
	Match         *HttpRouteMatch                                                   `json:"match,omitempty" xml:"match,omitempty"`
	OperationInfo *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s RouteRulesConflictInfoConflictsDetailsDetectedMatch) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsDetailsDetectedMatch) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsDetailsDetectedMatch) SetMatch(v *HttpRouteMatch) *RouteRulesConflictInfoConflictsDetailsDetectedMatch {
	s.Match = v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetailsDetectedMatch) SetOperationInfo(v *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo) *RouteRulesConflictInfoConflictsDetailsDetectedMatch {
	s.OperationInfo = v
	return s
}

type RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetName(v string) *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetOperationId(v string) *RouteRulesConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.OperationId = &v
	return s
}

type RouteRulesConflictInfoConflictsEnvironmentInfo struct {
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RouteRulesConflictInfoConflictsEnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoConflictsEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoConflictsEnvironmentInfo) SetEnvironmentId(v string) *RouteRulesConflictInfoConflictsEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *RouteRulesConflictInfoConflictsEnvironmentInfo) SetName(v string) *RouteRulesConflictInfoConflictsEnvironmentInfo {
	s.Name = &v
	return s
}

type RouteRulesConflictInfoDomainInfo struct {
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RouteRulesConflictInfoDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s RouteRulesConflictInfoDomainInfo) GoString() string {
	return s.String()
}

func (s *RouteRulesConflictInfoDomainInfo) SetDomainId(v string) *RouteRulesConflictInfoDomainInfo {
	s.DomainId = &v
	return s
}

func (s *RouteRulesConflictInfoDomainInfo) SetName(v string) *RouteRulesConflictInfoDomainInfo {
	s.Name = &v
	return s
}

type ServiceHealthCheck struct {
	// example:
	//
	// true
	Enable           *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	HealthyThreshold *int32  `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	HttpHost         *string `json:"httpHost,omitempty" xml:"httpHost,omitempty"`
	HttpPath         *string `json:"httpPath,omitempty" xml:"httpPath,omitempty"`
	Interval         *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// TCP
	Protocol           *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Timeout            *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	UnhealthyThreshold *int32  `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s ServiceHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s ServiceHealthCheck) GoString() string {
	return s.String()
}

func (s *ServiceHealthCheck) SetEnable(v bool) *ServiceHealthCheck {
	s.Enable = &v
	return s
}

func (s *ServiceHealthCheck) SetHealthyThreshold(v int32) *ServiceHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *ServiceHealthCheck) SetHttpHost(v string) *ServiceHealthCheck {
	s.HttpHost = &v
	return s
}

func (s *ServiceHealthCheck) SetHttpPath(v string) *ServiceHealthCheck {
	s.HttpPath = &v
	return s
}

func (s *ServiceHealthCheck) SetInterval(v int32) *ServiceHealthCheck {
	s.Interval = &v
	return s
}

func (s *ServiceHealthCheck) SetProtocol(v string) *ServiceHealthCheck {
	s.Protocol = &v
	return s
}

func (s *ServiceHealthCheck) SetTimeout(v int32) *ServiceHealthCheck {
	s.Timeout = &v
	return s
}

func (s *ServiceHealthCheck) SetUnhealthyThreshold(v int32) *ServiceHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

type ServiceLinkedRole struct {
	Arn                      *string `json:"arn,omitempty" xml:"arn,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty" xml:"assumeRolePolicyDocument,omitempty"`
	CreateDate               *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	Description              *string `json:"description,omitempty" xml:"description,omitempty"`
	IsServiceLinkedRole      *bool   `json:"isServiceLinkedRole,omitempty" xml:"isServiceLinkedRole,omitempty"`
	RoleId                   *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName                 *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	RolePrincipalName        *string `json:"rolePrincipalName,omitempty" xml:"rolePrincipalName,omitempty"`
}

func (s ServiceLinkedRole) String() string {
	return tea.Prettify(s)
}

func (s ServiceLinkedRole) GoString() string {
	return s.String()
}

func (s *ServiceLinkedRole) SetArn(v string) *ServiceLinkedRole {
	s.Arn = &v
	return s
}

func (s *ServiceLinkedRole) SetAssumeRolePolicyDocument(v string) *ServiceLinkedRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *ServiceLinkedRole) SetCreateDate(v string) *ServiceLinkedRole {
	s.CreateDate = &v
	return s
}

func (s *ServiceLinkedRole) SetDescription(v string) *ServiceLinkedRole {
	s.Description = &v
	return s
}

func (s *ServiceLinkedRole) SetIsServiceLinkedRole(v bool) *ServiceLinkedRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ServiceLinkedRole) SetRoleId(v string) *ServiceLinkedRole {
	s.RoleId = &v
	return s
}

func (s *ServiceLinkedRole) SetRoleName(v string) *ServiceLinkedRole {
	s.RoleName = &v
	return s
}

func (s *ServiceLinkedRole) SetRolePrincipalName(v string) *ServiceLinkedRole {
	s.RolePrincipalName = &v
	return s
}

type SslCertMetaInfo struct {
	Algorithm          *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	CertId             *int64  `json:"certId,omitempty" xml:"certId,omitempty"`
	CertIdentifier     *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	CertName           *string `json:"certName,omitempty" xml:"certName,omitempty"`
	CommonName         *string `json:"commonName,omitempty" xml:"commonName,omitempty"`
	Domain             *string `json:"domain,omitempty" xml:"domain,omitempty"`
	DomainMatchCert    *bool   `json:"domainMatchCert,omitempty" xml:"domainMatchCert,omitempty"`
	Fingerprint        *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	InstanceId         *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsChainCompleted   *bool   `json:"isChainCompleted,omitempty" xml:"isChainCompleted,omitempty"`
	Issuer             *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	KeySize            *string `json:"keySize,omitempty" xml:"keySize,omitempty"`
	Md5                *string `json:"md5,omitempty" xml:"md5,omitempty"`
	NotAfterTimestamp  *int64  `json:"notAfterTimestamp,omitempty" xml:"notAfterTimestamp,omitempty"`
	NotBeforeTimestamp *int64  `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	Sans               *string `json:"sans,omitempty" xml:"sans,omitempty"`
	SerialNo           *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	Sha2               *string `json:"sha2,omitempty" xml:"sha2,omitempty"`
	SignAlgorithm      *string `json:"signAlgorithm,omitempty" xml:"signAlgorithm,omitempty"`
}

func (s SslCertMetaInfo) String() string {
	return tea.Prettify(s)
}

func (s SslCertMetaInfo) GoString() string {
	return s.String()
}

func (s *SslCertMetaInfo) SetAlgorithm(v string) *SslCertMetaInfo {
	s.Algorithm = &v
	return s
}

func (s *SslCertMetaInfo) SetCertId(v int64) *SslCertMetaInfo {
	s.CertId = &v
	return s
}

func (s *SslCertMetaInfo) SetCertIdentifier(v string) *SslCertMetaInfo {
	s.CertIdentifier = &v
	return s
}

func (s *SslCertMetaInfo) SetCertName(v string) *SslCertMetaInfo {
	s.CertName = &v
	return s
}

func (s *SslCertMetaInfo) SetCommonName(v string) *SslCertMetaInfo {
	s.CommonName = &v
	return s
}

func (s *SslCertMetaInfo) SetDomain(v string) *SslCertMetaInfo {
	s.Domain = &v
	return s
}

func (s *SslCertMetaInfo) SetDomainMatchCert(v bool) *SslCertMetaInfo {
	s.DomainMatchCert = &v
	return s
}

func (s *SslCertMetaInfo) SetFingerprint(v string) *SslCertMetaInfo {
	s.Fingerprint = &v
	return s
}

func (s *SslCertMetaInfo) SetInstanceId(v string) *SslCertMetaInfo {
	s.InstanceId = &v
	return s
}

func (s *SslCertMetaInfo) SetIsChainCompleted(v bool) *SslCertMetaInfo {
	s.IsChainCompleted = &v
	return s
}

func (s *SslCertMetaInfo) SetIssuer(v string) *SslCertMetaInfo {
	s.Issuer = &v
	return s
}

func (s *SslCertMetaInfo) SetKeySize(v string) *SslCertMetaInfo {
	s.KeySize = &v
	return s
}

func (s *SslCertMetaInfo) SetMd5(v string) *SslCertMetaInfo {
	s.Md5 = &v
	return s
}

func (s *SslCertMetaInfo) SetNotAfterTimestamp(v int64) *SslCertMetaInfo {
	s.NotAfterTimestamp = &v
	return s
}

func (s *SslCertMetaInfo) SetNotBeforeTimestamp(v int64) *SslCertMetaInfo {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *SslCertMetaInfo) SetSans(v string) *SslCertMetaInfo {
	s.Sans = &v
	return s
}

func (s *SslCertMetaInfo) SetSerialNo(v string) *SslCertMetaInfo {
	s.SerialNo = &v
	return s
}

func (s *SslCertMetaInfo) SetSha2(v string) *SslCertMetaInfo {
	s.Sha2 = &v
	return s
}

func (s *SslCertMetaInfo) SetSignAlgorithm(v string) *SslCertMetaInfo {
	s.SignAlgorithm = &v
	return s
}

type SubDomainInfo struct {
	DomainId    *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Protocol    *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s SubDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s SubDomainInfo) GoString() string {
	return s.String()
}

func (s *SubDomainInfo) SetDomainId(v string) *SubDomainInfo {
	s.DomainId = &v
	return s
}

func (s *SubDomainInfo) SetName(v string) *SubDomainInfo {
	s.Name = &v
	return s
}

func (s *SubDomainInfo) SetNetworkType(v string) *SubDomainInfo {
	s.NetworkType = &v
	return s
}

func (s *SubDomainInfo) SetProtocol(v string) *SubDomainInfo {
	s.Protocol = &v
	return s
}

type AddGatewaySecurityGroupRuleRequest struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	PortRanges  []*string `json:"portRanges,omitempty" xml:"portRanges,omitempty" type:"Repeated"`
	// example:
	//
	// sg-wz929kxhcdp****
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s AddGatewaySecurityGroupRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleRequest) SetDescription(v string) *AddGatewaySecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleRequest) SetPortRanges(v []*string) *AddGatewaySecurityGroupRuleRequest {
	s.PortRanges = v
	return s
}

func (s *AddGatewaySecurityGroupRuleRequest) SetSecurityGroupId(v string) *AddGatewaySecurityGroupRuleRequest {
	s.SecurityGroupId = &v
	return s
}

type AddGatewaySecurityGroupRuleResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2A6E90D5-A711-54F4-A489-E33C2021EDDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddGatewaySecurityGroupRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetCode(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetMessage(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponseBody) SetRequestId(v string) *AddGatewaySecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

type AddGatewaySecurityGroupRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewaySecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewaySecurityGroupRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleResponse) SetHeaders(v map[string]*string) *AddGatewaySecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponse) SetStatusCode(v int32) *AddGatewaySecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleResponse) SetBody(v *AddGatewaySecurityGroupRuleResponseBody) *AddGatewaySecurityGroupRuleResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// example:
	//
	// 194445-cn-hangzhou
	CaCertIndentifier *string `json:"caCertIndentifier,omitempty" xml:"caCertIndentifier,omitempty"`
	// example:
	//
	// 194445-cn-hangzhou
	CertIndentifier *string `json:"certIndentifier,omitempty" xml:"certIndentifier,omitempty"`
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// TLS1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// example:
	//
	// TLS1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetCaCertIndentifier(v string) *CreateDomainRequest {
	s.CaCertIndentifier = &v
	return s
}

func (s *CreateDomainRequest) SetCertIndentifier(v string) *CreateDomainRequest {
	s.CertIndentifier = &v
	return s
}

func (s *CreateDomainRequest) SetForceHttps(v bool) *CreateDomainRequest {
	s.ForceHttps = &v
	return s
}

func (s *CreateDomainRequest) SetHttp2Option(v string) *CreateDomainRequest {
	s.Http2Option = &v
	return s
}

func (s *CreateDomainRequest) SetName(v string) *CreateDomainRequest {
	s.Name = &v
	return s
}

func (s *CreateDomainRequest) SetProtocol(v string) *CreateDomainRequest {
	s.Protocol = &v
	return s
}

func (s *CreateDomainRequest) SetTlsMax(v string) *CreateDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *CreateDomainRequest) SetTlsMin(v string) *CreateDomainRequest {
	s.TlsMin = &v
	return s
}

type CreateDomainResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0C2D1C68-0D93-5561-8EE6-FDB7BF067A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetCode(v string) *CreateDomainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDomainResponseBody) SetData(v *CreateDomainResponseBodyData) *CreateDomainResponseBody {
	s.Data = v
	return s
}

func (s *CreateDomainResponseBody) SetMessage(v string) *CreateDomainResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponseBodyData struct {
	// example:
	//
	// d-cpu1aullhtgkidg7sa4g
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
}

func (s CreateDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyData) SetDomainId(v string) *CreateDomainResponseBodyData {
	s.DomainId = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateEnvironmentRequest struct {
	// This parameter is required.
	Alias       *string `json:"alias,omitempty" xml:"alias,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) SetAlias(v string) *CreateEnvironmentRequest {
	s.Alias = &v
	return s
}

func (s *CreateEnvironmentRequest) SetDescription(v string) *CreateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *CreateEnvironmentRequest) SetGatewayId(v string) *CreateEnvironmentRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetName(v string) *CreateEnvironmentRequest {
	s.Name = &v
	return s
}

type CreateEnvironmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3C3B9A12-3868-5EB9-8BEA-F99E03DD125C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) SetCode(v string) *CreateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetData(v *CreateEnvironmentResponseBodyData) *CreateEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentResponseBody) SetMessage(v string) *CreateEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetRequestId(v string) *CreateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

type CreateEnvironmentResponseBodyData struct {
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s CreateEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBodyData) SetEnvironmentId(v string) *CreateEnvironmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

type CreateEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetStatusCode(v int32) *CreateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *CreateEnvironmentResponseBody) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

type CreateGatewayRouteRequest struct {
	BackendConfig *GatewayRouteBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	DomainConfig  *GatewayRouteDomainConfig  `json:"domainConfig,omitempty" xml:"domainConfig,omitempty"`
	Match         *HttpRouteMatch            `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// itemcenter-route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateGatewayRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRouteRequest) SetBackendConfig(v *GatewayRouteBackendConfig) *CreateGatewayRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *CreateGatewayRouteRequest) SetDescription(v string) *CreateGatewayRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayRouteRequest) SetDomainConfig(v *GatewayRouteDomainConfig) *CreateGatewayRouteRequest {
	s.DomainConfig = v
	return s
}

func (s *CreateGatewayRouteRequest) SetMatch(v *HttpRouteMatch) *CreateGatewayRouteRequest {
	s.Match = v
	return s
}

func (s *CreateGatewayRouteRequest) SetName(v string) *CreateGatewayRouteRequest {
	s.Name = &v
	return s
}

type CreateGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateGatewayRouteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayRouteResponseBody) SetCode(v string) *CreateGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayRouteResponseBody) SetData(v *CreateGatewayRouteResponseBodyData) *CreateGatewayRouteResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayRouteResponseBody) SetMessage(v string) *CreateGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayRouteResponseBody) SetRequestId(v string) *CreateGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type CreateGatewayRouteResponseBodyData struct {
	// example:
	//
	// gr-cpumc37d*****
	GatewayRouteId *string `json:"gatewayRouteId,omitempty" xml:"gatewayRouteId,omitempty"`
}

func (s CreateGatewayRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGatewayRouteResponseBodyData) SetGatewayRouteId(v string) *CreateGatewayRouteResponseBodyData {
	s.GatewayRouteId = &v
	return s
}

type CreateGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayRouteResponse) SetHeaders(v map[string]*string) *CreateGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayRouteResponse) SetStatusCode(v int32) *CreateGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayRouteResponse) SetBody(v *CreateGatewayRouteResponseBody) *CreateGatewayRouteResponse {
	s.Body = v
	return s
}

type CreateGatewayServiceRequest struct {
	GatewayServiceConfigs []*CreateGatewayServiceRequestGatewayServiceConfigs `json:"gatewayServiceConfigs,omitempty" xml:"gatewayServiceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CreateGatewayServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceRequest) SetGatewayServiceConfigs(v []*CreateGatewayServiceRequestGatewayServiceConfigs) *CreateGatewayServiceRequest {
	s.GatewayServiceConfigs = v
	return s
}

func (s *CreateGatewayServiceRequest) SetSourceType(v string) *CreateGatewayServiceRequest {
	s.SourceType = &v
	return s
}

type CreateGatewayServiceRequestGatewayServiceConfigs struct {
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// example:
	//
	// group-1
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// itemcenter-provider
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// MSE_NACOS
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s CreateGatewayServiceRequestGatewayServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceRequestGatewayServiceConfigs) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetAddresses(v []*string) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.Addresses = v
	return s
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetGroupName(v string) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.GroupName = &v
	return s
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetName(v string) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.Name = &v
	return s
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetNamespace(v string) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetPort(v int32) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.Port = &v
	return s
}

func (s *CreateGatewayServiceRequestGatewayServiceConfigs) SetQualifier(v string) *CreateGatewayServiceRequestGatewayServiceConfigs {
	s.Qualifier = &v
	return s
}

type CreateGatewayServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateGatewayServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9BA5586D-0EAE-5F78-B704-1A8DBADE46DA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceResponseBody) SetCode(v string) *CreateGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayServiceResponseBody) SetData(v *CreateGatewayServiceResponseBodyData) *CreateGatewayServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayServiceResponseBody) SetMessage(v string) *CreateGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayServiceResponseBody) SetRequestId(v string) *CreateGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateGatewayServiceResponseBodyData struct {
	GatewayServiceIds []*string `json:"gatewayServiceIds,omitempty" xml:"gatewayServiceIds,omitempty" type:"Repeated"`
}

func (s CreateGatewayServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceResponseBodyData) SetGatewayServiceIds(v []*string) *CreateGatewayServiceResponseBodyData {
	s.GatewayServiceIds = v
	return s
}

type CreateGatewayServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceResponse) SetHeaders(v map[string]*string) *CreateGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayServiceResponse) SetStatusCode(v int32) *CreateGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayServiceResponse) SetBody(v *CreateGatewayServiceResponseBody) *CreateGatewayServiceResponse {
	s.Body = v
	return s
}

type CreateGatewayServiceVersionRequest struct {
	Labels []*CreateGatewayServiceVersionRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// example:
	//
	// v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateGatewayServiceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceVersionRequest) SetLabels(v []*CreateGatewayServiceVersionRequestLabels) *CreateGatewayServiceVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateGatewayServiceVersionRequest) SetName(v string) *CreateGatewayServiceVersionRequest {
	s.Name = &v
	return s
}

type CreateGatewayServiceVersionRequestLabels struct {
	// example:
	//
	// app
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// itemcenter-blue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateGatewayServiceVersionRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceVersionRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceVersionRequestLabels) SetKey(v string) *CreateGatewayServiceVersionRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateGatewayServiceVersionRequestLabels) SetValue(v string) *CreateGatewayServiceVersionRequestLabels {
	s.Value = &v
	return s
}

type CreateGatewayServiceVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceVersionResponseBody) SetCode(v string) *CreateGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayServiceVersionResponseBody) SetMessage(v string) *CreateGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayServiceVersionResponseBody) SetRequestId(v string) *CreateGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateGatewayServiceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *CreateGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayServiceVersionResponse) SetStatusCode(v int32) *CreateGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayServiceVersionResponse) SetBody(v *CreateGatewayServiceVersionResponseBody) *CreateGatewayServiceVersionResponse {
	s.Body = v
	return s
}

type CreateHttpApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath    *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name          *string               `json:"name,omitempty" xml:"name,omitempty"`
	Protocols     []*string             `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s CreateHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRequest) SetBasePath(v string) *CreateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *CreateHttpApiRequest) SetDescription(v string) *CreateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *CreateHttpApiRequest) SetName(v string) *CreateHttpApiRequest {
	s.Name = &v
	return s
}

func (s *CreateHttpApiRequest) SetProtocols(v []*string) *CreateHttpApiRequest {
	s.Protocols = v
	return s
}

func (s *CreateHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiRequest {
	s.VersionConfig = v
	return s
}

type CreateHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponseBody) SetCode(v string) *CreateHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiResponseBody) SetData(v *CreateHttpApiResponseBodyData) *CreateHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiResponseBody) SetMessage(v string) *CreateHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiResponseBody) SetRequestId(v string) *CreateHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type CreateHttpApiResponseBodyData struct {
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateHttpApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponseBodyData) SetHttpApiId(v string) *CreateHttpApiResponseBodyData {
	s.HttpApiId = &v
	return s
}

func (s *CreateHttpApiResponseBodyData) SetName(v string) *CreateHttpApiResponseBodyData {
	s.Name = &v
	return s
}

type CreateHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponse) SetHeaders(v map[string]*string) *CreateHttpApiResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiResponse) SetStatusCode(v int32) *CreateHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiResponse) SetBody(v *CreateHttpApiResponseBody) *CreateHttpApiResponse {
	s.Body = v
	return s
}

type CreateHttpApiOperationRequest struct {
	Operations []*HttpApiOperation `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
}

func (s CreateHttpApiOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiOperationRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationRequest) SetOperations(v []*HttpApiOperation) *CreateHttpApiOperationRequest {
	s.Operations = v
	return s
}

type CreateHttpApiOperationResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateHttpApiOperationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBody) SetCode(v string) *CreateHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetData(v *CreateHttpApiOperationResponseBodyData) *CreateHttpApiOperationResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetMessage(v string) *CreateHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetRequestId(v string) *CreateHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

type CreateHttpApiOperationResponseBodyData struct {
	Operations []*CreateHttpApiOperationResponseBodyDataOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
}

func (s CreateHttpApiOperationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBodyData) SetOperations(v []*CreateHttpApiOperationResponseBodyDataOperations) *CreateHttpApiOperationResponseBodyData {
	s.Operations = v
	return s
}

type CreateHttpApiOperationResponseBodyDataOperations struct {
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s CreateHttpApiOperationResponseBodyDataOperations) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiOperationResponseBodyDataOperations) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBodyDataOperations) SetOperationId(v string) *CreateHttpApiOperationResponseBodyDataOperations {
	s.OperationId = &v
	return s
}

type CreateHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponse) SetHeaders(v map[string]*string) *CreateHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiOperationResponse) SetStatusCode(v int32) *CreateHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiOperationResponse) SetBody(v *CreateHttpApiOperationResponseBody) *CreateHttpApiOperationResponse {
	s.Body = v
	return s
}

type CreateServiceSourceRequest struct {
	K8sServiceSourceConfig   *CreateServiceSourceRequestK8sServiceSourceConfig   `json:"k8sServiceSourceConfig,omitempty" xml:"k8sServiceSourceConfig,omitempty" type:"Struct"`
	NacosServiceSourceConfig *CreateServiceSourceRequestNacosServiceSourceConfig `json:"nacosServiceSourceConfig,omitempty" xml:"nacosServiceSourceConfig,omitempty" type:"Struct"`
	Type                     *string                                             `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateServiceSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceRequest) SetK8sServiceSourceConfig(v *CreateServiceSourceRequestK8sServiceSourceConfig) *CreateServiceSourceRequest {
	s.K8sServiceSourceConfig = v
	return s
}

func (s *CreateServiceSourceRequest) SetNacosServiceSourceConfig(v *CreateServiceSourceRequestNacosServiceSourceConfig) *CreateServiceSourceRequest {
	s.NacosServiceSourceConfig = v
	return s
}

func (s *CreateServiceSourceRequest) SetType(v string) *CreateServiceSourceRequest {
	s.Type = &v
	return s
}

type CreateServiceSourceRequestK8sServiceSourceConfig struct {
	AuthorizeSecurityGroupRules []*CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules `json:"authorizeSecurityGroupRules,omitempty" xml:"authorizeSecurityGroupRules,omitempty" type:"Repeated"`
	// example:
	//
	// c4a21b3560fad4ec299f3e******
	ClusterId     *string                                                        `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	IngressConfig *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
}

func (s CreateServiceSourceRequestK8sServiceSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceRequestK8sServiceSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfig) SetAuthorizeSecurityGroupRules(v []*CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) *CreateServiceSourceRequestK8sServiceSourceConfig {
	s.AuthorizeSecurityGroupRules = v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfig) SetClusterId(v string) *CreateServiceSourceRequestK8sServiceSourceConfig {
	s.ClusterId = &v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfig) SetIngressConfig(v *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) *CreateServiceSourceRequestK8sServiceSourceConfig {
	s.IngressConfig = v
	return s
}

type CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules struct {
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	PortRanges      []*string `json:"portRanges,omitempty" xml:"portRanges,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) SetDescription(v string) *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules {
	s.Description = &v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) SetPortRanges(v []*string) *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules {
	s.PortRanges = v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules) SetSecurityGroupId(v string) *CreateServiceSourceRequestK8sServiceSourceConfigAuthorizeSecurityGroupRules {
	s.SecurityGroupId = &v
	return s
}

type CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetEnable(v bool) *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.Enable = &v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetIngressClass(v string) *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetOverrideIngressIp(v bool) *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetWatchNamespace(v string) *CreateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.WatchNamespace = &v
	return s
}

type CreateServiceSourceRequestNacosServiceSourceConfig struct {
	// example:
	//
	// mse-cn-328fc8***
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateServiceSourceRequestNacosServiceSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceRequestNacosServiceSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceRequestNacosServiceSourceConfig) SetInstanceId(v string) *CreateServiceSourceRequestNacosServiceSourceConfig {
	s.InstanceId = &v
	return s
}

type CreateServiceSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateServiceSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854CD0D4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceResponseBody) SetCode(v string) *CreateServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceSourceResponseBody) SetData(v *CreateServiceSourceResponseBodyData) *CreateServiceSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceSourceResponseBody) SetMessage(v string) *CreateServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceSourceResponseBody) SetRequestId(v string) *CreateServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceSourceResponseBodyData struct {
	// example:
	//
	// gss-cpqots5lht****
	ServiceSourceId *string `json:"serviceSourceId,omitempty" xml:"serviceSourceId,omitempty"`
}

func (s CreateServiceSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceResponseBodyData) SetServiceSourceId(v string) *CreateServiceSourceResponseBodyData {
	s.ServiceSourceId = &v
	return s
}

type CreateServiceSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceSourceResponse) SetHeaders(v map[string]*string) *CreateServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceSourceResponse) SetStatusCode(v int32) *CreateServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceSourceResponse) SetBody(v *CreateServiceSourceResponseBody) *CreateServiceSourceResponse {
	s.Body = v
	return s
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A60EE5CA-1294-532A-9775-8D2FD1C6EFBF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetCode(v string) *DeleteDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDomainResponseBody) SetMessage(v string) *DeleteDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61E30D3-579A-5B43-994E-31E02EDC9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) SetCode(v string) *DeleteEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetMessage(v string) *DeleteEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetRequestId(v string) *DeleteEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentResponse) SetStatusCode(v int32) *DeleteEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvironmentResponse) SetBody(v *DeleteEnvironmentResponseBody) *DeleteEnvironmentResponse {
	s.Body = v
	return s
}

type DeleteGatewayResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DE97DFDB-7DF0-5AB9-941C-10D27D769E4B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetCode(v string) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponseBody) SetCode(v string) *DeleteGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetMessage(v string) *DeleteGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetRequestId(v string) *DeleteGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponse) SetHeaders(v map[string]*string) *DeleteGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayRouteResponse) SetStatusCode(v int32) *DeleteGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayRouteResponse) SetBody(v *DeleteGatewayRouteResponseBody) *DeleteGatewayRouteResponse {
	s.Body = v
	return s
}

type DeleteGatewayServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceResponseBody) SetCode(v string) *DeleteGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetMessage(v string) *DeleteGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayServiceResponseBody) SetRequestId(v string) *DeleteGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceResponse) SetHeaders(v map[string]*string) *DeleteGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayServiceResponse) SetStatusCode(v int32) *DeleteGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayServiceResponse) SetBody(v *DeleteGatewayServiceResponseBody) *DeleteGatewayServiceResponse {
	s.Body = v
	return s
}

type DeleteGatewayServiceVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceVersionResponseBody) SetCode(v string) *DeleteGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetMessage(v string) *DeleteGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetRequestId(v string) *DeleteGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayServiceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *DeleteGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayServiceVersionResponse) SetStatusCode(v int32) *DeleteGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponse) SetBody(v *DeleteGatewayServiceVersionResponseBody) *DeleteGatewayServiceVersionResponse {
	s.Body = v
	return s
}

type DeleteHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiResponseBody) SetCode(v string) *DeleteHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiResponseBody) SetMessage(v string) *DeleteHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiResponseBody) SetRequestId(v string) *DeleteHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiResponse) SetHeaders(v map[string]*string) *DeleteHttpApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiResponse) SetStatusCode(v int32) *DeleteHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiResponse) SetBody(v *DeleteHttpApiResponseBody) *DeleteHttpApiResponse {
	s.Body = v
	return s
}

type DeleteHttpApiOperationResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiOperationResponseBody) SetCode(v string) *DeleteHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiOperationResponseBody) SetMessage(v string) *DeleteHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiOperationResponseBody) SetRequestId(v string) *DeleteHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiOperationResponse) SetHeaders(v map[string]*string) *DeleteHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiOperationResponse) SetStatusCode(v int32) *DeleteHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiOperationResponse) SetBody(v *DeleteHttpApiOperationResponseBody) *DeleteHttpApiOperationResponse {
	s.Body = v
	return s
}

type DeleteServiceSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceSourceResponseBody) SetCode(v string) *DeleteServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetMessage(v string) *DeleteServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceSourceResponseBody) SetRequestId(v string) *DeleteServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceSourceResponse) SetHeaders(v map[string]*string) *DeleteServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceSourceResponse) SetStatusCode(v int32) *DeleteServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceSourceResponse) SetBody(v *DeleteServiceSourceResponseBody) *DeleteServiceSourceResponse {
	s.Body = v
	return s
}

type GetDomainResponseBody struct {
	// example:
	//
	// Ok
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) SetCode(v string) *GetDomainResponseBody {
	s.Code = &v
	return s
}

func (s *GetDomainResponseBody) SetData(v *GetDomainResponseBodyData) *GetDomainResponseBody {
	s.Data = v
	return s
}

func (s *GetDomainResponseBody) SetMessage(v string) *GetDomainResponseBody {
	s.Message = &v
	return s
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainResponseBodyData struct {
	// example:
	//
	// RSA
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// example:
	//
	// 223576-cn-hangzhou
	CaCertIndentifier *string `json:"caCertIndentifier,omitempty" xml:"caCertIndentifier,omitempty"`
	// example:
	//
	// 123576-cn-hangzhou
	CertIndentifier *string `json:"certIndentifier,omitempty" xml:"certIndentifier,omitempty"`
	// example:
	//
	// test-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// false
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// example:
	//
	// d-cq1m3utlhtgvgkv7sitg
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// example:
	//
	// Alibaba
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1719386834548
	NotAfterTimstamp *int64 `json:"notAfterTimstamp,omitempty" xml:"notAfterTimstamp,omitempty"`
	// example:
	//
	// 1719386834548
	NotBeforeTimestamp *int64 `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// aliyun.com
	Sans *string `json:"sans,omitempty" xml:"sans,omitempty"`
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
	// example:
	//
	// 1719386834548
	Updatetimestamp *int64 `json:"updatetimestamp,omitempty" xml:"updatetimestamp,omitempty"`
}

func (s GetDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyData) SetAlgorithm(v string) *GetDomainResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCaCertIndentifier(v string) *GetDomainResponseBodyData {
	s.CaCertIndentifier = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCertIndentifier(v string) *GetDomainResponseBodyData {
	s.CertIndentifier = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCertName(v string) *GetDomainResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCreateFrom(v string) *GetDomainResponseBodyData {
	s.CreateFrom = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCreateTimestamp(v int64) *GetDomainResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetDefault(v bool) *GetDomainResponseBodyData {
	s.Default = &v
	return s
}

func (s *GetDomainResponseBodyData) SetDomainId(v string) *GetDomainResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *GetDomainResponseBodyData) SetForceHttps(v bool) *GetDomainResponseBodyData {
	s.ForceHttps = &v
	return s
}

func (s *GetDomainResponseBodyData) SetHttp2Option(v string) *GetDomainResponseBodyData {
	s.Http2Option = &v
	return s
}

func (s *GetDomainResponseBodyData) SetIssuer(v string) *GetDomainResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *GetDomainResponseBodyData) SetName(v string) *GetDomainResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDomainResponseBodyData) SetNotAfterTimstamp(v int64) *GetDomainResponseBodyData {
	s.NotAfterTimstamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetNotBeforeTimestamp(v int64) *GetDomainResponseBodyData {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetProtocol(v string) *GetDomainResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetDomainResponseBodyData) SetSans(v string) *GetDomainResponseBodyData {
	s.Sans = &v
	return s
}

func (s *GetDomainResponseBodyData) SetTlsMax(v string) *GetDomainResponseBodyData {
	s.TlsMax = &v
	return s
}

func (s *GetDomainResponseBodyData) SetTlsMin(v string) *GetDomainResponseBodyData {
	s.TlsMin = &v
	return s
}

func (s *GetDomainResponseBodyData) SetUpdatetimestamp(v int64) *GetDomainResponseBodyData {
	s.Updatetimestamp = &v
	return s
}

type GetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

type GetEnvironmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3F8EE674-BB08-5E92-BE6F-E4756A748B0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBody) SetCode(v string) *GetEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentResponseBody) SetMessage(v string) *GetEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetRequestId(v string) *GetEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

type GetEnvironmentResponseBodyData struct {
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// true
	Default     *bool   `json:"default,omitempty" xml:"default,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string      `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// example:
	//
	// test
	Name           *string          `json:"name,omitempty" xml:"name,omitempty"`
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyData) SetAlias(v string) *GetEnvironmentResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCreateTimestamp(v int64) *GetEnvironmentResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDefault(v bool) *GetEnvironmentResponseBodyData {
	s.Default = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDescription(v string) *GetEnvironmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetEnvironmentId(v string) *GetEnvironmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetGatewayInfo(v *GatewayInfo) *GetEnvironmentResponseBodyData {
	s.GatewayInfo = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetName(v string) *GetEnvironmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetSubDomainInfos(v []*SubDomainInfo) *GetEnvironmentResponseBodyData {
	s.SubDomainInfos = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetUpdateTimestamp(v int64) *GetEnvironmentResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

type GetEnvironmentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetStatusCode(v int32) *GetEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetGatewayResponseBody struct {
	// example:
	//
	// Ok
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetGatewayResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) SetCode(v string) *GetGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetMessage(v string) *GetGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResponseBody) SetRequestId(v string) *GetGatewayResponseBody {
	s.RequestId = &v
	return s
}

type GetGatewayResponseBodyData struct {
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64                                    `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Environments    []*GetGatewayResponseBodyDataEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// 1719386834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId     *string                                    `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	LoadBalancers []*GetGatewayResponseBodyDataLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Replicas      *string                                  `json:"replicas,omitempty" xml:"replicas,omitempty"`
	SecurityGroup *GetGatewayResponseBodyDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64                             `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	VSwitch         *GetGatewayResponseBodyDataVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// example:
	//
	// 2.0.2
	Version *string                            `json:"version,omitempty" xml:"version,omitempty"`
	Vpc     *GetGatewayResponseBodyDataVpc     `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	Zones   []*GetGatewayResponseBodyDataZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetGatewayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyData) SetChargeType(v string) *GetGatewayResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateFrom(v string) *GetGatewayResponseBodyData {
	s.CreateFrom = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEnvironments(v []*GetGatewayResponseBodyDataEnvironments) *GetGatewayResponseBodyData {
	s.Environments = v
	return s
}

func (s *GetGatewayResponseBodyData) SetExpireTimestamp(v int64) *GetGatewayResponseBodyData {
	s.ExpireTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayId(v string) *GetGatewayResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLoadBalancers(v []*GetGatewayResponseBodyDataLoadBalancers) *GetGatewayResponseBodyData {
	s.LoadBalancers = v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetReplicas(v string) *GetGatewayResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSecurityGroup(v *GetGatewayResponseBodyDataSecurityGroup) *GetGatewayResponseBodyData {
	s.SecurityGroup = v
	return s
}

func (s *GetGatewayResponseBodyData) SetSpec(v string) *GetGatewayResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatus(v string) *GetGatewayResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetTargetVersion(v string) *GetGatewayResponseBodyData {
	s.TargetVersion = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetUpdateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVSwitch(v *GetGatewayResponseBodyDataVSwitch) *GetGatewayResponseBodyData {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyData) SetVersion(v string) *GetGatewayResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVpc(v *GetGatewayResponseBodyDataVpc) *GetGatewayResponseBodyData {
	s.Vpc = v
	return s
}

func (s *GetGatewayResponseBodyData) SetZones(v []*GetGatewayResponseBodyDataZones) *GetGatewayResponseBodyData {
	s.Zones = v
	return s
}

type GetGatewayResponseBodyDataEnvironments struct {
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// env-cp9uhudlht***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// default-gw-cp9ugg5***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetGatewayResponseBodyDataEnvironments) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataEnvironments) SetAlias(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Alias = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetEnvironmentId(v string) *GetGatewayResponseBodyDataEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetName(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Name = &v
	return s
}

type GetGatewayResponseBodyDataLoadBalancers struct {
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// example:
	//
	// nlb-xoh3pghru7c***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// example:
	//
	// Managed
	Mode  *string                                         `json:"mode,omitempty" xml:"mode,omitempty"`
	Ports []*GetGatewayResponseBodyDataLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancers) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddress(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Address = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressIpVersion(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetGatewayDefault(v bool) *GetGatewayResponseBodyDataLoadBalancers {
	s.GatewayDefault = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetLoadBalancerId(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetMode(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Mode = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetPorts(v []*GetGatewayResponseBodyDataLoadBalancersPorts) *GetGatewayResponseBodyDataLoadBalancers {
	s.Ports = v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetStatus(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Type = &v
	return s
}

type GetGatewayResponseBodyDataLoadBalancersPorts struct {
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetPort(v int32) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Port = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetProtocol(v string) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Protocol = &v
	return s
}

type GetGatewayResponseBodyDataSecurityGroup struct {
	// example:
	//
	// APIG-sg-gw-cq7ke5ll***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// sg-bp16tafq9***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s GetGatewayResponseBodyDataSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataSecurityGroup) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetName(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetSecurityGroupId(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

type GetGatewayResponseBodyDataVSwitch struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataVSwitch) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVSwitch) SetName(v string) *GetGatewayResponseBodyDataVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataVSwitch {
	s.VSwitchId = &v
	return s
}

type GetGatewayResponseBodyDataVpc struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-bp1llj52lvj6xc***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetGatewayResponseBodyDataVpc) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataVpc) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVpc) SetName(v string) *GetGatewayResponseBodyDataVpc {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVpc) SetVpcId(v string) *GetGatewayResponseBodyDataVpc {
	s.VpcId = &v
	return s
}

type GetGatewayResponseBodyDataZones struct {
	Name    *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	VSwitch *GetGatewayResponseBodyDataZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetGatewayResponseBodyDataZones) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataZones) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZones) SetName(v string) *GetGatewayResponseBodyDataZones {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetVSwitch(v *GetGatewayResponseBodyDataZonesVSwitch) *GetGatewayResponseBodyDataZones {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetZoneId(v string) *GetGatewayResponseBodyDataZones {
	s.ZoneId = &v
	return s
}

type GetGatewayResponseBodyDataZonesVSwitch struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataZonesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataZonesVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetName(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.VSwitchId = &v
	return s
}

type GetGatewayResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayResponse) SetHeaders(v map[string]*string) *GetGatewayResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayResponse) SetStatusCode(v int32) *GetGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayResponse) SetBody(v *GetGatewayResponseBody) *GetGatewayResponse {
	s.Body = v
	return s
}

type GetGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetGatewayRouteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteResponseBody) SetCode(v string) *GetGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayRouteResponseBody) SetData(v *GetGatewayRouteResponseBodyData) *GetGatewayRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayRouteResponseBody) SetMessage(v string) *GetGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayRouteResponseBody) SetRequestId(v string) *GetGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type GetGatewayRouteResponseBodyData struct {
	Backend *GatewayRouteBackend `json:"backend,omitempty" xml:"backend,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64                  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Description     *string                 `json:"description,omitempty" xml:"description,omitempty"`
	DomainInfo      *GatewayRouteDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty"`
	// example:
	//
	// gr-cptf6e7d5l***
	GatewayRouteId *string         `json:"gatewayRouteId,omitempty" xml:"gatewayRouteId,omitempty"`
	Match          *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// itemcenter-pre-route
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// NotPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetGatewayRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteResponseBodyData) SetBackend(v *GatewayRouteBackend) *GetGatewayRouteResponseBodyData {
	s.Backend = v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetCreateTimestamp(v int64) *GetGatewayRouteResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetDescription(v string) *GetGatewayRouteResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetDomainInfo(v *GatewayRouteDomainInfo) *GetGatewayRouteResponseBodyData {
	s.DomainInfo = v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetGatewayRouteId(v string) *GetGatewayRouteResponseBodyData {
	s.GatewayRouteId = &v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetMatch(v *HttpRouteMatch) *GetGatewayRouteResponseBodyData {
	s.Match = v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetName(v string) *GetGatewayRouteResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetStatus(v string) *GetGatewayRouteResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteResponseBodyData) SetUpdateTimestamp(v int64) *GetGatewayRouteResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

type GetGatewayRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteResponse) SetHeaders(v map[string]*string) *GetGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayRouteResponse) SetStatusCode(v int32) *GetGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayRouteResponse) SetBody(v *GetGatewayRouteResponseBody) *GetGatewayRouteResponse {
	s.Body = v
	return s
}

type GetGatewayServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GatewayService `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceResponseBody) SetCode(v string) *GetGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayServiceResponseBody) SetData(v *GatewayService) *GetGatewayServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayServiceResponseBody) SetMessage(v string) *GetGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayServiceResponseBody) SetRequestId(v string) *GetGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetGatewayServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceResponse) SetHeaders(v map[string]*string) *GetGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayServiceResponse) SetStatusCode(v int32) *GetGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayServiceResponse) SetBody(v *GetGatewayServiceResponseBody) *GetGatewayServiceResponse {
	s.Body = v
	return s
}

type GetHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string         `json:"code,omitempty" xml:"code,omitempty"`
	Data *HttpApiApiInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiResponseBody) SetCode(v string) *GetHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiResponseBody) SetData(v *HttpApiApiInfo) *GetHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiResponseBody) SetMessage(v string) *GetHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiResponseBody) SetRequestId(v string) *GetHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type GetHttpApiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiResponse) SetHeaders(v map[string]*string) *GetHttpApiResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiResponse) SetStatusCode(v int32) *GetHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiResponse) SetBody(v *GetHttpApiResponseBody) *GetHttpApiResponse {
	s.Body = v
	return s
}

type GetHttpApiOperationResponseBody struct {
	// example:
	//
	// Ok
	Code *string               `json:"code,omitempty" xml:"code,omitempty"`
	Data *HttpApiOperationInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B725275B-50C6-5A49-A9FD-F0332FCB3351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHttpApiOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiOperationResponseBody) SetCode(v string) *GetHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetData(v *HttpApiOperationInfo) *GetHttpApiOperationResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetMessage(v string) *GetHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiOperationResponseBody) SetRequestId(v string) *GetHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

type GetHttpApiOperationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiOperationResponse) SetHeaders(v map[string]*string) *GetHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiOperationResponse) SetStatusCode(v int32) *GetHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiOperationResponse) SetBody(v *GetHttpApiOperationResponseBody) *GetHttpApiOperationResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetGatewayId(v string) *ListDomainsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListDomainsRequest) SetNameLike(v string) *ListDomainsRequest {
	s.NameLike = &v
	return s
}

func (s *ListDomainsRequest) SetPageNumber(v int32) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int32) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

type ListDomainsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListDomainsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61E30D3-579A-5B43-994E-31E02EDC9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetCode(v string) *ListDomainsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDomainsResponseBody) SetData(v *ListDomainsResponseBodyData) *ListDomainsResponseBody {
	s.Data = v
	return s
}

func (s *ListDomainsResponseBody) SetMessage(v string) *ListDomainsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListDomainsResponseBodyData struct {
	Items []*DomainInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 9
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDomainsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyData) SetItems(v []*DomainInfo) *ListDomainsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDomainsResponseBodyData) SetPageNumber(v int32) *ListDomainsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsResponseBodyData) SetPageSize(v int32) *ListDomainsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDomainsResponseBodyData) SetTotalSize(v int32) *ListDomainsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListEnvironmentsRequest struct {
	AliasLike *string `json:"aliasLike,omitempty" xml:"aliasLike,omitempty"`
	// example:
	//
	// gw-cptv6ktlhtgnqr73h8d1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test-gw
	GatewayNameLike *string `json:"gatewayNameLike,omitempty" xml:"gatewayNameLike,omitempty"`
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) SetAliasLike(v string) *ListEnvironmentsRequest {
	s.AliasLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetGatewayId(v string) *ListEnvironmentsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetGatewayNameLike(v string) *ListEnvironmentsRequest {
	s.GatewayNameLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetNameLike(v string) *ListEnvironmentsRequest {
	s.NameLike = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageNumber(v int32) *ListEnvironmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int32) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71E20
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) SetCode(v string) *ListEnvironmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetMessage(v string) *ListEnvironmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetRequestId(v string) *ListEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

type ListEnvironmentsResponseBodyData struct {
	Items []*EnvironmentInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) SetItems(v []*EnvironmentInfo) *ListEnvironmentsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageNumber(v int32) *ListEnvironmentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageSize(v int32) *ListEnvironmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetTotalSize(v int32) *ListEnvironmentsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListEnvironmentsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetStatusCode(v int32) *ListEnvironmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type ListGatewayRoutesRequest struct {
	// example:
	//
	// itemcenter
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// pre-itemcenter-router
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// /user
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// Published
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListGatewayRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRoutesRequest) SetKeyword(v string) *ListGatewayRoutesRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewayRoutesRequest) SetName(v string) *ListGatewayRoutesRequest {
	s.Name = &v
	return s
}

func (s *ListGatewayRoutesRequest) SetPageNumber(v int32) *ListGatewayRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRoutesRequest) SetPageSize(v int32) *ListGatewayRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRoutesRequest) SetPath(v string) *ListGatewayRoutesRequest {
	s.Path = &v
	return s
}

func (s *ListGatewayRoutesRequest) SetStatus(v string) *ListGatewayRoutesRequest {
	s.Status = &v
	return s
}

type ListGatewayRoutesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListGatewayRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayRoutesResponseBody) SetCode(v string) *ListGatewayRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayRoutesResponseBody) SetData(v *ListGatewayRoutesResponseBodyData) *ListGatewayRoutesResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayRoutesResponseBody) SetMessage(v string) *ListGatewayRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayRoutesResponseBody) SetRequestId(v string) *ListGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

type ListGatewayRoutesResponseBodyData struct {
	Items []*ListGatewayRoutesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGatewayRoutesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayRoutesResponseBodyData) SetItems(v []*ListGatewayRoutesResponseBodyDataItems) *ListGatewayRoutesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayRoutesResponseBodyData) SetPageNumber(v int32) *ListGatewayRoutesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyData) SetPageSize(v int32) *ListGatewayRoutesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyData) SetTotalSize(v int32) *ListGatewayRoutesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListGatewayRoutesResponseBodyDataItems struct {
	Backend *GatewayRouteBackend `json:"backend,omitempty" xml:"backend,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64                  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Description     *string                 `json:"description,omitempty" xml:"description,omitempty"`
	DomainInfo      *GatewayRouteDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty"`
	// example:
	//
	// gr-cqa8oddlhtg***
	GatewayRouteId *string         `json:"gatewayRouteId,omitempty" xml:"gatewayRouteId,omitempty"`
	Match          *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// pre-itemcenter-router
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// NotPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s ListGatewayRoutesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRoutesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetBackend(v *GatewayRouteBackend) *ListGatewayRoutesResponseBodyDataItems {
	s.Backend = v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetCreateTimestamp(v int64) *ListGatewayRoutesResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetDescription(v string) *ListGatewayRoutesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetDomainInfo(v *GatewayRouteDomainInfo) *ListGatewayRoutesResponseBodyDataItems {
	s.DomainInfo = v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetGatewayRouteId(v string) *ListGatewayRoutesResponseBodyDataItems {
	s.GatewayRouteId = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetMatch(v *HttpRouteMatch) *ListGatewayRoutesResponseBodyDataItems {
	s.Match = v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetName(v string) *ListGatewayRoutesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetStatus(v string) *ListGatewayRoutesResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListGatewayRoutesResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListGatewayRoutesResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

type ListGatewayRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayRoutesResponse) SetHeaders(v map[string]*string) *ListGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayRoutesResponse) SetStatusCode(v int32) *ListGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayRoutesResponse) SetBody(v *ListGatewayRoutesResponseBody) *ListGatewayRoutesResponse {
	s.Body = v
	return s
}

type ListGatewayServicesRequest struct {
	// example:
	//
	// itemcenter-provider
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListGatewayServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayServicesRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayServicesRequest) SetName(v string) *ListGatewayServicesRequest {
	s.Name = &v
	return s
}

func (s *ListGatewayServicesRequest) SetPageNumber(v int32) *ListGatewayServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayServicesRequest) SetPageSize(v int32) *ListGatewayServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayServicesRequest) SetSourceType(v string) *ListGatewayServicesRequest {
	s.SourceType = &v
	return s
}

type ListGatewayServicesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListGatewayServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayServicesResponseBody) SetCode(v string) *ListGatewayServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayServicesResponseBody) SetData(v *ListGatewayServicesResponseBodyData) *ListGatewayServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayServicesResponseBody) SetMessage(v string) *ListGatewayServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayServicesResponseBody) SetRequestId(v string) *ListGatewayServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListGatewayServicesResponseBodyData struct {
	Items []*GatewayService `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 6
	TotalSize *int64 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGatewayServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayServicesResponseBodyData) SetItems(v []*GatewayService) *ListGatewayServicesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayServicesResponseBodyData) SetPageNumber(v int32) *ListGatewayServicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayServicesResponseBodyData) SetPageSize(v int32) *ListGatewayServicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayServicesResponseBodyData) SetTotalSize(v int64) *ListGatewayServicesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListGatewayServicesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayServicesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayServicesResponse) SetHeaders(v map[string]*string) *ListGatewayServicesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayServicesResponse) SetStatusCode(v int32) *ListGatewayServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayServicesResponse) SetBody(v *ListGatewayServicesResponseBody) *ListGatewayServicesResponse {
	s.Body = v
	return s
}

type ListGatewaysRequest struct {
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// dev
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequest) SetGatewayId(v string) *ListGatewaysRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysRequest) SetKeyword(v string) *ListGatewaysRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewaysRequest) SetName(v string) *ListGatewaysRequest {
	s.Name = &v
	return s
}

func (s *ListGatewaysRequest) SetPageNumber(v int32) *ListGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysRequest) SetPageSize(v int32) *ListGatewaysRequest {
	s.PageSize = &v
	return s
}

type ListGatewaysResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListGatewaysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBody) SetCode(v string) *ListGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewaysResponseBody) SetData(v *ListGatewaysResponseBodyData) *ListGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewaysResponseBody) SetMessage(v string) *ListGatewaysResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewaysResponseBody) SetRequestId(v string) *ListGatewaysResponseBody {
	s.RequestId = &v
	return s
}

type ListGatewaysResponseBodyData struct {
	Items []*ListGatewaysResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 6
	TotalSize *int64 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGatewaysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyData) SetItems(v []*ListGatewaysResponseBodyDataItems) *ListGatewaysResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewaysResponseBodyData) SetPageNumber(v int32) *ListGatewaysResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysResponseBodyData) SetPageSize(v int32) *ListGatewaysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysResponseBodyData) SetTotalSize(v int64) *ListGatewaysResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListGatewaysResponseBodyDataItems struct {
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 172086834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// example:
	//
	// gw-cpv54p5***
	GatewayId     *string                                           `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	LoadBalancers []*ListGatewaysResponseBodyDataItemsLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Replicas      *string                                         `json:"replicas,omitempty" xml:"replicas,omitempty"`
	SecurityGroup *ListGatewaysResponseBodyDataItemsSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64                                    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	VSwitch         *ListGatewaysResponseBodyDataItemsVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// example:
	//
	// 2.0.2
	Version *string                                   `json:"version,omitempty" xml:"version,omitempty"`
	Vpc     *ListGatewaysResponseBodyDataItemsVpc     `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	Zones   []*ListGatewaysResponseBodyDataItemsZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s ListGatewaysResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItems) SetChargeType(v string) *ListGatewaysResponseBodyDataItems {
	s.ChargeType = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetCreateFrom(v string) *ListGatewaysResponseBodyDataItems {
	s.CreateFrom = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetCreateTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetExpireTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.ExpireTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetGatewayId(v string) *ListGatewaysResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetLoadBalancers(v []*ListGatewaysResponseBodyDataItemsLoadBalancers) *ListGatewaysResponseBodyDataItems {
	s.LoadBalancers = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetName(v string) *ListGatewaysResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetReplicas(v string) *ListGatewaysResponseBodyDataItems {
	s.Replicas = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetSecurityGroup(v *ListGatewaysResponseBodyDataItemsSecurityGroup) *ListGatewaysResponseBodyDataItems {
	s.SecurityGroup = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetSpec(v string) *ListGatewaysResponseBodyDataItems {
	s.Spec = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetStatus(v string) *ListGatewaysResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetTargetVersion(v string) *ListGatewaysResponseBodyDataItems {
	s.TargetVersion = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListGatewaysResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVSwitch(v *ListGatewaysResponseBodyDataItemsVSwitch) *ListGatewaysResponseBodyDataItems {
	s.VSwitch = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVersion(v string) *ListGatewaysResponseBodyDataItems {
	s.Version = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetVpc(v *ListGatewaysResponseBodyDataItemsVpc) *ListGatewaysResponseBodyDataItems {
	s.Vpc = v
	return s
}

func (s *ListGatewaysResponseBodyDataItems) SetZones(v []*ListGatewaysResponseBodyDataItemsZones) *ListGatewaysResponseBodyDataItems {
	s.Zones = v
	return s
}

type ListGatewaysResponseBodyDataItemsLoadBalancers struct {
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// example:
	//
	// nlb-xqwioje1c91r***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// example:
	//
	// Managed
	Mode  *string                                                `json:"mode,omitempty" xml:"mode,omitempty"`
	Ports []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddress(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Address = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddressIpVersion(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetAddressType(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetGatewayDefault(v bool) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.GatewayDefault = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetLoadBalancerId(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetMode(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Mode = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetPorts(v []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Ports = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetStatus(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Status = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancers) SetType(v string) *ListGatewaysResponseBodyDataItemsLoadBalancers {
	s.Type = &v
	return s
}

type ListGatewaysResponseBodyDataItemsLoadBalancersPorts struct {
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancersPorts) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsLoadBalancersPorts) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) SetPort(v int32) *ListGatewaysResponseBodyDataItemsLoadBalancersPorts {
	s.Port = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsLoadBalancersPorts) SetProtocol(v string) *ListGatewaysResponseBodyDataItemsLoadBalancersPorts {
	s.Protocol = &v
	return s
}

type ListGatewaysResponseBodyDataItemsSecurityGroup struct {
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsSecurityGroup) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsSecurityGroup) SetSecurityGroupId(v string) *ListGatewaysResponseBodyDataItemsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

type ListGatewaysResponseBodyDataItemsVSwitch struct {
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsVSwitch) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsVSwitch) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsVSwitch) SetVSwitchId(v string) *ListGatewaysResponseBodyDataItemsVSwitch {
	s.VSwitchId = &v
	return s
}

type ListGatewaysResponseBodyDataItemsVpc struct {
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsVpc) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsVpc) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsVpc) SetVpcId(v string) *ListGatewaysResponseBodyDataItemsVpc {
	s.VpcId = &v
	return s
}

type ListGatewaysResponseBodyDataItemsZones struct {
	VSwitch *ListGatewaysResponseBodyDataItemsZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	ZoneId  *string                                        `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsZones) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsZones) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsZones) SetVSwitch(v *ListGatewaysResponseBodyDataItemsZonesVSwitch) *ListGatewaysResponseBodyDataItemsZones {
	s.VSwitch = v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsZones) SetZoneId(v string) *ListGatewaysResponseBodyDataItemsZones {
	s.ZoneId = &v
	return s
}

type ListGatewaysResponseBodyDataItemsZonesVSwitch struct {
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsZonesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsZonesVSwitch) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsZonesVSwitch) SetVSwitchId(v string) *ListGatewaysResponseBodyDataItemsZonesVSwitch {
	s.VSwitchId = &v
	return s
}

type ListGatewaysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponse) SetHeaders(v map[string]*string) *ListGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListGatewaysResponse) SetStatusCode(v int32) *ListGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewaysResponse) SetBody(v *ListGatewaysResponseBody) *ListGatewaysResponse {
	s.Body = v
	return s
}

type ListHttpApiOperationsRequest struct {
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// GetUser
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// /v1
	PathLike *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
}

func (s ListHttpApiOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsRequest) SetMethod(v string) *ListHttpApiOperationsRequest {
	s.Method = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetName(v string) *ListHttpApiOperationsRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetNameLike(v string) *ListHttpApiOperationsRequest {
	s.NameLike = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPageNumber(v int32) *ListHttpApiOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPageSize(v int32) *ListHttpApiOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetPathLike(v string) *ListHttpApiOperationsRequest {
	s.PathLike = &v
	return s
}

type ListHttpApiOperationsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListHttpApiOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHttpApiOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponseBody) SetCode(v string) *ListHttpApiOperationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetData(v *ListHttpApiOperationsResponseBodyData) *ListHttpApiOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetMessage(v string) *ListHttpApiOperationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetRequestId(v string) *ListHttpApiOperationsResponseBody {
	s.RequestId = &v
	return s
}

type ListHttpApiOperationsResponseBodyData struct {
	Items []*HttpApiOperationInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApiOperationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponseBodyData) SetItems(v []*HttpApiOperationInfo) *ListHttpApiOperationsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetPageNumber(v int32) *ListHttpApiOperationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetPageSize(v int32) *ListHttpApiOperationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetTotalSize(v int32) *ListHttpApiOperationsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListHttpApiOperationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApiOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApiOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponse) SetHeaders(v map[string]*string) *ListHttpApiOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApiOperationsResponse) SetStatusCode(v int32) *ListHttpApiOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApiOperationsResponse) SetBody(v *ListHttpApiOperationsResponseBody) *ListHttpApiOperationsResponse {
	s.Body = v
	return s
}

type ListHttpApisRequest struct {
	// example:
	//
	// test-
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PublishedOnly *bool  `json:"publishedOnly,omitempty" xml:"publishedOnly,omitempty"`
}

func (s ListHttpApisRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApisRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApisRequest) SetKeyword(v string) *ListHttpApisRequest {
	s.Keyword = &v
	return s
}

func (s *ListHttpApisRequest) SetName(v string) *ListHttpApisRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApisRequest) SetPageNumber(v int32) *ListHttpApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApisRequest) SetPageSize(v int32) *ListHttpApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApisRequest) SetPublishedOnly(v bool) *ListHttpApisRequest {
	s.PublishedOnly = &v
	return s
}

type ListHttpApisResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListHttpApisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHttpApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponseBody) SetCode(v string) *ListHttpApisResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApisResponseBody) SetData(v *ListHttpApisResponseBodyData) *ListHttpApisResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApisResponseBody) SetMessage(v string) *ListHttpApisResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApisResponseBody) SetRequestId(v string) *ListHttpApisResponseBody {
	s.RequestId = &v
	return s
}

type ListHttpApisResponseBodyData struct {
	Items []*HttpApiInfoByName `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponseBodyData) SetItems(v []*HttpApiInfoByName) *ListHttpApisResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApisResponseBodyData) SetPageNumber(v int32) *ListHttpApisResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApisResponseBodyData) SetPageSize(v int32) *ListHttpApisResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApisResponseBodyData) SetTotalSize(v int32) *ListHttpApisResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListHttpApisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApisResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApisResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponse) SetHeaders(v map[string]*string) *ListHttpApisResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApisResponse) SetStatusCode(v int32) *ListHttpApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApisResponse) SetBody(v *ListHttpApisResponseBody) *ListHttpApisResponse {
	s.Body = v
	return s
}

type OfflineGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OfflineGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineGatewayRouteResponseBody) SetCode(v string) *OfflineGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetMessage(v string) *OfflineGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetRequestId(v string) *OfflineGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type OfflineGatewayRouteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *OfflineGatewayRouteResponse) SetHeaders(v map[string]*string) *OfflineGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *OfflineGatewayRouteResponse) SetStatusCode(v int32) *OfflineGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineGatewayRouteResponse) SetBody(v *OfflineGatewayRouteResponseBody) *OfflineGatewayRouteResponse {
	s.Body = v
	return s
}

type OfflineHttpApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s OfflineHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineHttpApiRequest) GoString() string {
	return s.String()
}

func (s *OfflineHttpApiRequest) SetEnvironmentId(v string) *OfflineHttpApiRequest {
	s.EnvironmentId = &v
	return s
}

type OfflineHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 53DE779E-422D-56EB-B84C-62D75CA5E8DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OfflineHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineHttpApiResponseBody) SetCode(v string) *OfflineHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineHttpApiResponseBody) SetMessage(v string) *OfflineHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineHttpApiResponseBody) SetRequestId(v string) *OfflineHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type OfflineHttpApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineHttpApiResponse) GoString() string {
	return s.String()
}

func (s *OfflineHttpApiResponse) SetHeaders(v map[string]*string) *OfflineHttpApiResponse {
	s.Headers = v
	return s
}

func (s *OfflineHttpApiResponse) SetStatusCode(v int32) *OfflineHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineHttpApiResponse) SetBody(v *OfflineHttpApiResponseBody) *OfflineHttpApiResponse {
	s.Body = v
	return s
}

type PublishGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PublishGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *PublishGatewayRouteResponseBody) SetCode(v string) *PublishGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *PublishGatewayRouteResponseBody) SetMessage(v string) *PublishGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *PublishGatewayRouteResponseBody) SetRequestId(v string) *PublishGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type PublishGatewayRouteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *PublishGatewayRouteResponse) SetHeaders(v map[string]*string) *PublishGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *PublishGatewayRouteResponse) SetStatusCode(v int32) *PublishGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishGatewayRouteResponse) SetBody(v *PublishGatewayRouteResponseBody) *PublishGatewayRouteResponse {
	s.Body = v
	return s
}

type PublishHttpApiRequest struct {
	// example:
	//
	// true
	AllowOverwrite *bool                             `json:"allowOverwrite,omitempty" xml:"allowOverwrite,omitempty"`
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Environment    *PublishHttpApiRequestEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// example:
	//
	// apr-xxx
	RevisionId *string `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
}

func (s PublishHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequest) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequest) SetAllowOverwrite(v bool) *PublishHttpApiRequest {
	s.AllowOverwrite = &v
	return s
}

func (s *PublishHttpApiRequest) SetDescription(v string) *PublishHttpApiRequest {
	s.Description = &v
	return s
}

func (s *PublishHttpApiRequest) SetEnvironment(v *PublishHttpApiRequestEnvironment) *PublishHttpApiRequest {
	s.Environment = v
	return s
}

func (s *PublishHttpApiRequest) SetRevisionId(v string) *PublishHttpApiRequest {
	s.RevisionId = &v
	return s
}

type PublishHttpApiRequestEnvironment struct {
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// example:
	//
	// Service
	BackendType        *string                                             `json:"backendType,omitempty" xml:"backendType,omitempty"`
	CloudProductConfig *PublishHttpApiRequestEnvironmentCloudProductConfig `json:"cloudProductConfig,omitempty" xml:"cloudProductConfig,omitempty" type:"Struct"`
	CustomDomainIds    []*string                                           `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	DnsConfigs         []*PublishHttpApiRequestEnvironmentDnsConfigs       `json:"dnsConfigs,omitempty" xml:"dnsConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentId  *string                                           `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	ServiceConfigs []*PublishHttpApiRequestEnvironmentServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	VipConfigs     []*PublishHttpApiRequestEnvironmentVipConfigs     `json:"vipConfigs,omitempty" xml:"vipConfigs,omitempty" type:"Repeated"`
}

func (s PublishHttpApiRequestEnvironment) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironment) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironment) SetBackendScene(v string) *PublishHttpApiRequestEnvironment {
	s.BackendScene = &v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetBackendType(v string) *PublishHttpApiRequestEnvironment {
	s.BackendType = &v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetCloudProductConfig(v *PublishHttpApiRequestEnvironmentCloudProductConfig) *PublishHttpApiRequestEnvironment {
	s.CloudProductConfig = v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetCustomDomainIds(v []*string) *PublishHttpApiRequestEnvironment {
	s.CustomDomainIds = v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetDnsConfigs(v []*PublishHttpApiRequestEnvironmentDnsConfigs) *PublishHttpApiRequestEnvironment {
	s.DnsConfigs = v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetEnvironmentId(v string) *PublishHttpApiRequestEnvironment {
	s.EnvironmentId = &v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetServiceConfigs(v []*PublishHttpApiRequestEnvironmentServiceConfigs) *PublishHttpApiRequestEnvironment {
	s.ServiceConfigs = v
	return s
}

func (s *PublishHttpApiRequestEnvironment) SetVipConfigs(v []*PublishHttpApiRequestEnvironmentVipConfigs) *PublishHttpApiRequestEnvironment {
	s.VipConfigs = v
	return s
}

type PublishHttpApiRequestEnvironmentCloudProductConfig struct {
	// example:
	//
	// FC
	CloudProductType        *string                                                                      `json:"cloudProductType,omitempty" xml:"cloudProductType,omitempty"`
	ContainerServiceConfigs []*PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs `json:"containerServiceConfigs,omitempty" xml:"containerServiceConfigs,omitempty" type:"Repeated"`
	FunctionConfigs         []*PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs         `json:"functionConfigs,omitempty" xml:"functionConfigs,omitempty" type:"Repeated"`
	MseNacosConfigs         []*PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs         `json:"mseNacosConfigs,omitempty" xml:"mseNacosConfigs,omitempty" type:"Repeated"`
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfig) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfig) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfig) SetCloudProductType(v string) *PublishHttpApiRequestEnvironmentCloudProductConfig {
	s.CloudProductType = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfig) SetContainerServiceConfigs(v []*PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) *PublishHttpApiRequestEnvironmentCloudProductConfig {
	s.ContainerServiceConfigs = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfig) SetFunctionConfigs(v []*PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) *PublishHttpApiRequestEnvironmentCloudProductConfig {
	s.FunctionConfigs = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfig) SetMseNacosConfigs(v []*PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) *PublishHttpApiRequestEnvironmentCloudProductConfig {
	s.MseNacosConfigs = v
	return s
}

type PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs struct {
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// test-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetName(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Name = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetNamespace(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetPort(v int32) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Port = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetProtocol(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs) SetWeight(v int32) *PublishHttpApiRequestEnvironmentCloudProductConfigContainerServiceConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs struct {
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// fc-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// LATEST
	Quanlifer *string `json:"quanlifer,omitempty" xml:"quanlifer,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) SetName(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs {
	s.Name = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) SetQuanlifer(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs {
	s.Quanlifer = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs) SetWeight(v int32) *PublishHttpApiRequestEnvironmentCloudProductConfigFunctionConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs struct {
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string                        `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// provider
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 100
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) SetGroupName(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs {
	s.GroupName = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) SetName(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs {
	s.Name = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) SetNamespace(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs {
	s.Namespace = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs) SetWeight(v string) *PublishHttpApiRequestEnvironmentCloudProductConfigMseNacosConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiRequestEnvironmentDnsConfigs struct {
	DnsList []*string                      `json:"dnsList,omitempty" xml:"dnsList,omitempty" type:"Repeated"`
	Match   *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentDnsConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentDnsConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentDnsConfigs) SetDnsList(v []*string) *PublishHttpApiRequestEnvironmentDnsConfigs {
	s.DnsList = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentDnsConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentDnsConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentDnsConfigs) SetWeight(v int32) *PublishHttpApiRequestEnvironmentDnsConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiRequestEnvironmentServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentServiceConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetGatewayServiceId(v string) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetPort(v int32) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.Port = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetProtocol(v string) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetVersion(v string) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.Version = &v
	return s
}

func (s *PublishHttpApiRequestEnvironmentServiceConfigs) SetWeight(v int32) *PublishHttpApiRequestEnvironmentServiceConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiRequestEnvironmentVipConfigs struct {
	Endpoints []*string                      `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s PublishHttpApiRequestEnvironmentVipConfigs) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiRequestEnvironmentVipConfigs) GoString() string {
	return s.String()
}

func (s *PublishHttpApiRequestEnvironmentVipConfigs) SetEndpoints(v []*string) *PublishHttpApiRequestEnvironmentVipConfigs {
	s.Endpoints = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentVipConfigs) SetMatch(v *HttpApiBackendMatchConditions) *PublishHttpApiRequestEnvironmentVipConfigs {
	s.Match = v
	return s
}

func (s *PublishHttpApiRequestEnvironmentVipConfigs) SetWeight(v int32) *PublishHttpApiRequestEnvironmentVipConfigs {
	s.Weight = &v
	return s
}

type PublishHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 4BACB05C-3FE2-588F-9148-700C5C026B74
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PublishHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *PublishHttpApiResponseBody) SetCode(v string) *PublishHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *PublishHttpApiResponseBody) SetMessage(v string) *PublishHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *PublishHttpApiResponseBody) SetRequestId(v string) *PublishHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type PublishHttpApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishHttpApiResponse) GoString() string {
	return s.String()
}

func (s *PublishHttpApiResponse) SetHeaders(v map[string]*string) *PublishHttpApiResponse {
	s.Headers = v
	return s
}

func (s *PublishHttpApiResponse) SetStatusCode(v int32) *PublishHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishHttpApiResponse) SetBody(v *PublishHttpApiResponseBody) *PublishHttpApiResponse {
	s.Body = v
	return s
}

type UpdateDomainRequest struct {
	// example:
	//
	// 123455-cn-hangzhou
	CaCertIndentifier *string `json:"caCertIndentifier,omitempty" xml:"caCertIndentifier,omitempty"`
	// example:
	//
	// 123458-cn-hangzhou
	CertIndentifier *string `json:"certIndentifier,omitempty" xml:"certIndentifier,omitempty"`
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) SetCaCertIndentifier(v string) *UpdateDomainRequest {
	s.CaCertIndentifier = &v
	return s
}

func (s *UpdateDomainRequest) SetCertIndentifier(v string) *UpdateDomainRequest {
	s.CertIndentifier = &v
	return s
}

func (s *UpdateDomainRequest) SetForceHttps(v bool) *UpdateDomainRequest {
	s.ForceHttps = &v
	return s
}

func (s *UpdateDomainRequest) SetHttp2Option(v string) *UpdateDomainRequest {
	s.Http2Option = &v
	return s
}

func (s *UpdateDomainRequest) SetProtocol(v string) *UpdateDomainRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateDomainRequest) SetTlsMax(v string) *UpdateDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *UpdateDomainRequest) SetTlsMin(v string) *UpdateDomainRequest {
	s.TlsMin = &v
	return s
}

type UpdateDomainResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 4BACB05C-3FE2-588F-9148-700C5C026B74
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBody) SetCode(v string) *UpdateDomainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDomainResponseBody) SetData(v *UpdateDomainResponseBodyData) *UpdateDomainResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDomainResponseBody) SetMessage(v string) *UpdateDomainResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDomainResponseBody) SetRequestId(v string) *UpdateDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainResponseBodyData struct {
	DeployRevisionId *string `json:"deployRevisionId,omitempty" xml:"deployRevisionId,omitempty"`
}

func (s UpdateDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBodyData) SetDeployRevisionId(v string) *UpdateDomainResponseBodyData {
	s.DeployRevisionId = &v
	return s
}

type UpdateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponse) SetHeaders(v map[string]*string) *UpdateDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainResponse) SetStatusCode(v int32) *UpdateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainResponse) SetBody(v *UpdateDomainResponseBody) *UpdateDomainResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentRequest struct {
	// This parameter is required.
	Alias       *string `json:"alias,omitempty" xml:"alias,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) SetAlias(v string) *UpdateEnvironmentRequest {
	s.Alias = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetDescription(v string) *UpdateEnvironmentRequest {
	s.Description = &v
	return s
}

type UpdateEnvironmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 52FB803B-3CD8-5FF8-AAE9-C2B841F6A483
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) SetCode(v string) *UpdateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetMessage(v string) *UpdateEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetRequestId(v string) *UpdateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetStatusCode(v int32) *UpdateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

type UpdateGatewayRouteRequest struct {
	BackendConfig *GatewayRouteBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty"`
	Description   *string                    `json:"description,omitempty" xml:"description,omitempty"`
	DomainConfig  *GatewayRouteDomainConfig  `json:"domainConfig,omitempty" xml:"domainConfig,omitempty"`
	Match         *HttpRouteMatch            `json:"match,omitempty" xml:"match,omitempty"`
}

func (s UpdateGatewayRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequest) SetBackendConfig(v *GatewayRouteBackendConfig) *UpdateGatewayRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDescription(v string) *UpdateGatewayRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDomainConfig(v *GatewayRouteDomainConfig) *UpdateGatewayRouteRequest {
	s.DomainConfig = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetMatch(v *HttpRouteMatch) *UpdateGatewayRouteRequest {
	s.Match = v
	return s
}

type UpdateGatewayRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B725275B-50C6-5A49-A9FD-F0332FCB3351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteResponseBody) SetCode(v string) *UpdateGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetMessage(v string) *UpdateGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetRequestId(v string) *UpdateGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteResponse) SetStatusCode(v int32) *UpdateGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteResponse) SetBody(v *UpdateGatewayRouteResponseBody) *UpdateGatewayRouteResponse {
	s.Body = v
	return s
}

type UpdateGatewayServiceRequest struct {
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s UpdateGatewayServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceRequest) SetAddresses(v []*string) *UpdateGatewayServiceRequest {
	s.Addresses = v
	return s
}

func (s *UpdateGatewayServiceRequest) SetPort(v int32) *UpdateGatewayServiceRequest {
	s.Port = &v
	return s
}

type UpdateGatewayServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceResponseBody) SetCode(v string) *UpdateGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetMessage(v string) *UpdateGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceResponseBody) SetRequestId(v string) *UpdateGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceResponse) SetStatusCode(v int32) *UpdateGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceResponse) SetBody(v *UpdateGatewayServiceResponseBody) *UpdateGatewayServiceResponse {
	s.Body = v
	return s
}

type UpdateGatewayServiceVersionRequest struct {
	Labels []*UpdateGatewayServiceVersionRequestLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
}

func (s UpdateGatewayServiceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionRequest) SetLabels(v []*UpdateGatewayServiceVersionRequestLabels) *UpdateGatewayServiceVersionRequest {
	s.Labels = v
	return s
}

type UpdateGatewayServiceVersionRequestLabels struct {
	// example:
	//
	// topology.kubernetes.io/zone
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateGatewayServiceVersionRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceVersionRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionRequestLabels) SetKey(v string) *UpdateGatewayServiceVersionRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateGatewayServiceVersionRequestLabels) SetValue(v string) *UpdateGatewayServiceVersionRequestLabels {
	s.Value = &v
	return s
}

type UpdateGatewayServiceVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionResponseBody) SetCode(v string) *UpdateGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetMessage(v string) *UpdateGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponseBody) SetRequestId(v string) *UpdateGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayServiceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceVersionResponse) SetStatusCode(v int32) *UpdateGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponse) SetBody(v *UpdateGatewayServiceVersionResponseBody) *UpdateGatewayServiceVersionResponse {
	s.Body = v
	return s
}

type UpdateHttpApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath      *string               `json:"basePath,omitempty" xml:"basePath,omitempty"`
	Description   *string               `json:"description,omitempty" xml:"description,omitempty"`
	Protocols     []*string             `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s UpdateHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRequest) SetBasePath(v string) *UpdateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *UpdateHttpApiRequest) SetDescription(v string) *UpdateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRequest) SetProtocols(v []*string) *UpdateHttpApiRequest {
	s.Protocols = v
	return s
}

func (s *UpdateHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *UpdateHttpApiRequest {
	s.VersionConfig = v
	return s
}

type UpdateHttpApiResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiResponseBody) SetCode(v string) *UpdateHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiResponseBody) SetMessage(v string) *UpdateHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiResponseBody) SetRequestId(v string) *UpdateHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiResponse) SetHeaders(v map[string]*string) *UpdateHttpApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiResponse) SetStatusCode(v int32) *UpdateHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiResponse) SetBody(v *UpdateHttpApiResponseBody) *UpdateHttpApiResponse {
	s.Body = v
	return s
}

type UpdateHttpApiOperationRequest struct {
	Operation *HttpApiOperation `json:"operation,omitempty" xml:"operation,omitempty"`
}

func (s UpdateHttpApiOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiOperationRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationRequest) SetOperation(v *HttpApiOperation) *UpdateHttpApiOperationRequest {
	s.Operation = v
	return s
}

type UpdateHttpApiOperationResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationResponseBody) SetCode(v string) *UpdateHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiOperationResponseBody) SetMessage(v string) *UpdateHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiOperationResponseBody) SetRequestId(v string) *UpdateHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationResponse) SetHeaders(v map[string]*string) *UpdateHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiOperationResponse) SetStatusCode(v int32) *UpdateHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiOperationResponse) SetBody(v *UpdateHttpApiOperationResponseBody) *UpdateHttpApiOperationResponse {
	s.Body = v
	return s
}

type UpdateServiceSourceRequest struct {
	K8sServiceSourceConfig *UpdateServiceSourceRequestK8sServiceSourceConfig `json:"k8sServiceSourceConfig,omitempty" xml:"k8sServiceSourceConfig,omitempty" type:"Struct"`
}

func (s UpdateServiceSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceRequest) SetK8sServiceSourceConfig(v *UpdateServiceSourceRequestK8sServiceSourceConfig) *UpdateServiceSourceRequest {
	s.K8sServiceSourceConfig = v
	return s
}

type UpdateServiceSourceRequestK8sServiceSourceConfig struct {
	IngressConfig *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
}

func (s UpdateServiceSourceRequestK8sServiceSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSourceRequestK8sServiceSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceRequestK8sServiceSourceConfig) SetIngressConfig(v *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) *UpdateServiceSourceRequestK8sServiceSourceConfig {
	s.IngressConfig = v
	return s
}

type UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetEnable(v bool) *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.Enable = &v
	return s
}

func (s *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetIngressClass(v string) *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetOverrideIngressIp(v bool) *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig) SetWatchNamespace(v string) *UpdateServiceSourceRequestK8sServiceSourceConfigIngressConfig {
	s.WatchNamespace = &v
	return s
}

type UpdateServiceSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceResponseBody) SetCode(v string) *UpdateServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetMessage(v string) *UpdateServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetRequestId(v string) *UpdateServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceResponse) SetHeaders(v map[string]*string) *UpdateServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceSourceResponse) SetStatusCode(v int32) *UpdateServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceSourceResponse) SetBody(v *UpdateServiceSourceResponseBody) *UpdateServiceSourceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("apig"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权网关访问服务的安全组
//
// @param request - AddGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewaySecurityGroupRuleResponse
func (client *Client) AddGatewaySecurityGroupRuleWithOptions(gatewayId *string, request *AddGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddGatewaySecurityGroupRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		body["portRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["securityGroupId"] = request.SecurityGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGatewaySecurityGroupRule"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/security-group-rules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGatewaySecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权网关访问服务的安全组
//
// @param request - AddGatewaySecurityGroupRuleRequest
//
// @return AddGatewaySecurityGroupRuleResponse
func (client *Client) AddGatewaySecurityGroupRule(gatewayId *string, request *AddGatewaySecurityGroupRuleRequest) (_result *AddGatewaySecurityGroupRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGatewaySecurityGroupRuleResponse{}
	_body, _err := client.AddGatewaySecurityGroupRuleWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建域名
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertIndentifier)) {
		body["caCertIndentifier"] = request.CaCertIndentifier
	}

	if !tea.BoolValue(util.IsUnset(request.CertIndentifier)) {
		body["certIndentifier"] = request.CertIndentifier
	}

	if !tea.BoolValue(util.IsUnset(request.ForceHttps)) {
		body["forceHttps"] = request.ForceHttps
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Option)) {
		body["http2Option"] = request.Http2Option
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.TlsMax)) {
		body["tlsMax"] = request.TlsMax
	}

	if !tea.BoolValue(util.IsUnset(request.TlsMin)) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建域名
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateEnvironment
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithOptions(request *CreateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEnvironment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/environments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateEnvironment
//
// @param request - CreateEnvironmentRequest
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironment(request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - CreateGatewayRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayRouteResponse
func (client *Client) CreateGatewayRouteWithOptions(gatewayId *string, request *CreateGatewayRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendConfig)) {
		body["backendConfig"] = request.BackendConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DomainConfig)) {
		body["domainConfig"] = request.DomainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Match)) {
		body["match"] = request.Match
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - CreateGatewayRouteRequest
//
// @return CreateGatewayRouteResponse
func (client *Client) CreateGatewayRoute(gatewayId *string, request *CreateGatewayRouteRequest) (_result *CreateGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayRouteResponse{}
	_body, _err := client.CreateGatewayRouteWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务
//
// @param request - CreateGatewayServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayServiceResponse
func (client *Client) CreateGatewayServiceWithOptions(gatewayId *string, request *CreateGatewayServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayServiceConfigs)) {
		body["gatewayServiceConfigs"] = request.GatewayServiceConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务
//
// @param request - CreateGatewayServiceRequest
//
// @return CreateGatewayServiceResponse
func (client *Client) CreateGatewayService(gatewayId *string, request *CreateGatewayServiceRequest) (_result *CreateGatewayServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayServiceResponse{}
	_body, _err := client.CreateGatewayServiceWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务版本
//
// @param request - CreateGatewayServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayServiceVersionResponse
func (client *Client) CreateGatewayServiceVersionWithOptions(gatewayId *string, gatewayServiceId *string, request *CreateGatewayServiceVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayServiceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayServiceVersion"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId)) + "/service-versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务版本
//
// @param request - CreateGatewayServiceVersionRequest
//
// @return CreateGatewayServiceVersionResponse
func (client *Client) CreateGatewayServiceVersion(gatewayId *string, gatewayServiceId *string, request *CreateGatewayServiceVersionRequest) (_result *CreateGatewayServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayServiceVersionResponse{}
	_body, _err := client.CreateGatewayServiceVersionWithOptions(gatewayId, gatewayServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个HTTP类型的API
//
// @param request - CreateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiResponse
func (client *Client) CreateHttpApiWithOptions(request *CreateHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		body["protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.VersionConfig)) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个HTTP类型的API
//
// @param request - CreateHttpApiRequest
//
// @return CreateHttpApiResponse
func (client *Client) CreateHttpApi(request *CreateHttpApiRequest) (_result *CreateHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiResponse{}
	_body, _err := client.CreateHttpApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为HTTP API创建Operation
//
// @param request - CreateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiOperationResponse
func (client *Client) CreateHttpApiOperationWithOptions(httpApiId *string, request *CreateHttpApiOperationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHttpApiOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operations)) {
		body["operations"] = request.Operations
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHttpApiOperation"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/operations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为HTTP API创建Operation
//
// @param request - CreateHttpApiOperationRequest
//
// @return CreateHttpApiOperationResponse
func (client *Client) CreateHttpApiOperation(httpApiId *string, request *CreateHttpApiOperationRequest) (_result *CreateHttpApiOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiOperationResponse{}
	_body, _err := client.CreateHttpApiOperationWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务来源
//
// @param request - CreateServiceSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceSourceResponse
func (client *Client) CreateServiceSourceWithOptions(gatewayId *string, request *CreateServiceSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sServiceSourceConfig)) {
		body["k8sServiceSourceConfig"] = request.K8sServiceSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.NacosServiceSourceConfig)) {
		body["nacosServiceSourceConfig"] = request.NacosServiceSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceSource"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/service-sources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务来源
//
// @param request - CreateServiceSourceRequest
//
// @return CreateServiceSourceResponse
func (client *Client) CreateServiceSource(gatewayId *string, request *CreateServiceSourceRequest) (_result *CreateServiceSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceSourceResponse{}
	_body, _err := client.CreateServiceSourceWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DeleteDomain
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(domainId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DeleteDomain
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(domainId *string) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(domainId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DeleteEnvironment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithOptions(environmentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnvironment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(environmentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DeleteEnvironment
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironment(environmentId *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(environmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(gatewayId *string) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayRouteResponse
func (client *Client) DeleteGatewayRouteWithOptions(gatewayId *string, gatewayRouteId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayRouteId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @return DeleteGatewayRouteResponse
func (client *Client) DeleteGatewayRoute(gatewayId *string, gatewayRouteId *string) (_result *DeleteGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.DeleteGatewayRouteWithOptions(gatewayId, gatewayRouteId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayServiceResponse
func (client *Client) DeleteGatewayServiceWithOptions(gatewayId *string, gatewayServiceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务
//
// @return DeleteGatewayServiceResponse
func (client *Client) DeleteGatewayService(gatewayId *string, gatewayServiceId *string) (_result *DeleteGatewayServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayServiceResponse{}
	_body, _err := client.DeleteGatewayServiceWithOptions(gatewayId, gatewayServiceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除服务版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayServiceVersionResponse
func (client *Client) DeleteGatewayServiceVersionWithOptions(gatewayId *string, gatewayServiceId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayServiceVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayServiceVersion"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId)) + "/service-versions/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务版本
//
// @return DeleteGatewayServiceVersionResponse
func (client *Client) DeleteGatewayServiceVersion(gatewayId *string, gatewayServiceId *string, name *string) (_result *DeleteGatewayServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayServiceVersionResponse{}
	_body, _err := client.DeleteGatewayServiceVersionWithOptions(gatewayId, gatewayServiceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除HTTP API
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiResponse
func (client *Client) DeleteHttpApiWithOptions(httpApiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHttpApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除HTTP API
//
// @return DeleteHttpApiResponse
func (client *Client) DeleteHttpApi(httpApiId *string) (_result *DeleteHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiResponse{}
	_body, _err := client.DeleteHttpApiWithOptions(httpApiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Operation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiOperationResponse
func (client *Client) DeleteHttpApiOperationWithOptions(httpApiId *string, operationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHttpApiOperationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHttpApiOperation"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/operations/" + tea.StringValue(openapiutil.GetEncodeParam(operationId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Operation
//
// @return DeleteHttpApiOperationResponse
func (client *Client) DeleteHttpApiOperation(httpApiId *string, operationId *string) (_result *DeleteHttpApiOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiOperationResponse{}
	_body, _err := client.DeleteHttpApiOperationWithOptions(httpApiId, operationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除服务来源
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceSourceResponse
func (client *Client) DeleteServiceSourceWithOptions(gatewayId *string, serviceSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceSource"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/service-sources/" + tea.StringValue(openapiutil.GetEncodeParam(serviceSourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务来源
//
// @return DeleteServiceSourceResponse
func (client *Client) DeleteServiceSource(gatewayId *string, serviceSourceId *string) (_result *DeleteServiceSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceSourceResponse{}
	_body, _err := client.DeleteServiceSourceWithOptions(gatewayId, serviceSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(domainId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名详情
//
// @return GetDomainResponse
func (client *Client) GetDomain(domainId *string) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(domainId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetEnvironment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithOptions(environmentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(environmentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetEnvironment
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironment(environmentId *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(environmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网关实例详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayResponse
func (client *Client) GetGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGateway"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关实例详情
//
// @return GetGatewayResponse
func (client *Client) GetGateway(gatewayId *string) (_result *GetGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayResponse{}
	_body, _err := client.GetGatewayWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayRouteResponse
func (client *Client) GetGatewayRouteWithOptions(gatewayId *string, gatewayRouteId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGatewayRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayRouteId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @return GetGatewayRouteResponse
func (client *Client) GetGatewayRoute(gatewayId *string, gatewayRouteId *string) (_result *GetGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayRouteResponse{}
	_body, _err := client.GetGatewayRouteWithOptions(gatewayId, gatewayRouteId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayServiceResponse
func (client *Client) GetGatewayServiceWithOptions(gatewayId *string, gatewayServiceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGatewayServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务
//
// @return GetGatewayServiceResponse
func (client *Client) GetGatewayService(gatewayId *string, gatewayServiceId *string) (_result *GetGatewayServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayServiceResponse{}
	_body, _err := client.GetGatewayServiceWithOptions(gatewayId, gatewayServiceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取HttpApi
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiResponse
func (client *Client) GetHttpApiWithOptions(httpApiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHttpApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取HttpApi
//
// @return GetHttpApiResponse
func (client *Client) GetHttpApi(httpApiId *string) (_result *GetHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiResponse{}
	_body, _err := client.GetHttpApiWithOptions(httpApiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取Operation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiOperationResponse
func (client *Client) GetHttpApiOperationWithOptions(httpApiId *string, operationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHttpApiOperationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHttpApiOperation"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/operations/" + tea.StringValue(openapiutil.GetEncodeParam(operationId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取Operation
//
// @return GetHttpApiOperationResponse
func (client *Client) GetHttpApiOperation(httpApiId *string, operationId *string) (_result *GetHttpApiOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiOperationResponse{}
	_body, _err := client.GetHttpApiOperationWithOptions(httpApiId, operationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListDomains
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NameLike)) {
		query["nameLike"] = request.NameLike
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ListDomains
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListEnvironments
//
// @param request - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithOptions(request *ListEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasLike)) {
		query["aliasLike"] = request.AliasLike
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayNameLike)) {
		query["gatewayNameLike"] = request.GatewayNameLike
	}

	if !tea.BoolValue(util.IsUnset(request.NameLike)) {
		query["nameLike"] = request.NameLike
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/environments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ListEnvironments
//
// @param request - ListEnvironmentsRequest
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironments(request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - ListGatewayRoutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayRoutesResponse
func (client *Client) ListGatewayRoutesWithOptions(gatewayId *string, request *ListGatewayRoutesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewayRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayRoutes"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - ListGatewayRoutesRequest
//
// @return ListGatewayRoutesResponse
func (client *Client) ListGatewayRoutes(gatewayId *string, request *ListGatewayRoutesRequest) (_result *ListGatewayRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayRoutesResponse{}
	_body, _err := client.ListGatewayRoutesWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务列表
//
// @param request - ListGatewayServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayServicesResponse
func (client *Client) ListGatewayServicesWithOptions(gatewayId *string, request *ListGatewayServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewayServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayServices"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务列表
//
// @param request - ListGatewayServicesRequest
//
// @return ListGatewayServicesResponse
func (client *Client) ListGatewayServices(gatewayId *string, request *ListGatewayServicesRequest) (_result *ListGatewayServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayServicesResponse{}
	_body, _err := client.ListGatewayServicesWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取已经创建的云原生网关列表
//
// @param request - ListGatewaysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewaysResponse
func (client *Client) ListGatewaysWithOptions(request *ListGatewaysRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGateways"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取已经创建的云原生网关列表
//
// @param request - ListGatewaysRequest
//
// @return ListGatewaysResponse
func (client *Client) ListGateways(request *ListGatewaysRequest) (_result *ListGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewaysResponse{}
	_body, _err := client.ListGatewaysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举Operation
//
// @param request - ListHttpApiOperationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiOperationsResponse
func (client *Client) ListHttpApiOperationsWithOptions(httpApiId *string, request *ListHttpApiOperationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHttpApiOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NameLike)) {
		query["nameLike"] = request.NameLike
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PathLike)) {
		query["pathLike"] = request.PathLike
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHttpApiOperations"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/operations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHttpApiOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举Operation
//
// @param request - ListHttpApiOperationsRequest
//
// @return ListHttpApiOperationsResponse
func (client *Client) ListHttpApiOperations(httpApiId *string, request *ListHttpApiOperationsRequest) (_result *ListHttpApiOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApiOperationsResponse{}
	_body, _err := client.ListHttpApiOperationsWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举HTTP API
//
// @param request - ListHttpApisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApisResponse
func (client *Client) ListHttpApisWithOptions(request *ListHttpApisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHttpApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublishedOnly)) {
		query["publishedOnly"] = request.PublishedOnly
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHttpApis"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHttpApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举HTTP API
//
// @param request - ListHttpApisRequest
//
// @return ListHttpApisResponse
func (client *Client) ListHttpApis(request *ListHttpApisRequest) (_result *ListHttpApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApisResponse{}
	_body, _err := client.ListHttpApisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布路由
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineGatewayRouteResponse
func (client *Client) OfflineGatewayRouteWithOptions(gatewayId *string, gatewayRouteId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OfflineGatewayRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayRouteId)) + "/offline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OfflineGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布路由
//
// @return OfflineGatewayRouteResponse
func (client *Client) OfflineGatewayRoute(gatewayId *string, gatewayRouteId *string) (_result *OfflineGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineGatewayRouteResponse{}
	_body, _err := client.OfflineGatewayRouteWithOptions(gatewayId, gatewayRouteId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下线已发布的HTTP API
//
// @param request - OfflineHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineHttpApiResponse
func (client *Client) OfflineHttpApiWithOptions(httpApiId *string, request *OfflineHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OfflineHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/offline"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OfflineHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线已发布的HTTP API
//
// @param request - OfflineHttpApiRequest
//
// @return OfflineHttpApiResponse
func (client *Client) OfflineHttpApi(httpApiId *string, request *OfflineHttpApiRequest) (_result *OfflineHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineHttpApiResponse{}
	_body, _err := client.OfflineHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布路由
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishGatewayRouteResponse
func (client *Client) PublishGatewayRouteWithOptions(gatewayId *string, gatewayRouteId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishGatewayRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayRouteId)) + "/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布路由
//
// @return PublishGatewayRouteResponse
func (client *Client) PublishGatewayRoute(gatewayId *string, gatewayRouteId *string) (_result *PublishGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishGatewayRouteResponse{}
	_body, _err := client.PublishGatewayRouteWithOptions(gatewayId, gatewayRouteId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布HTTP API
//
// @param request - PublishHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishHttpApiResponse
func (client *Client) PublishHttpApiWithOptions(httpApiId *string, request *PublishHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowOverwrite)) {
		body["allowOverwrite"] = request.AllowOverwrite
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionId)) {
		body["revisionId"] = request.RevisionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布HTTP API
//
// @param request - PublishHttpApiRequest
//
// @return PublishHttpApiResponse
func (client *Client) PublishHttpApi(httpApiId *string, request *PublishHttpApiRequest) (_result *PublishHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishHttpApiResponse{}
	_body, _err := client.PublishHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// UpdateDomain
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithOptions(domainId *string, request *UpdateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertIndentifier)) {
		body["caCertIndentifier"] = request.CaCertIndentifier
	}

	if !tea.BoolValue(util.IsUnset(request.CertIndentifier)) {
		body["certIndentifier"] = request.CertIndentifier
	}

	if !tea.BoolValue(util.IsUnset(request.ForceHttps)) {
		body["forceHttps"] = request.ForceHttps
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Option)) {
		body["http2Option"] = request.Http2Option
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.TlsMax)) {
		body["tlsMax"] = request.TlsMax
	}

	if !tea.BoolValue(util.IsUnset(request.TlsMin)) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDomain"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// UpdateDomain
//
// @param request - UpdateDomainRequest
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomain(domainId *string, request *UpdateDomainRequest) (_result *UpdateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDomainResponse{}
	_body, _err := client.UpdateDomainWithOptions(domainId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// UpdateEnvironment
//
// @param request - UpdateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithOptions(environmentId *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnvironment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/environments/" + tea.StringValue(openapiutil.GetEncodeParam(environmentId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// UpdateEnvironment
//
// @param request - UpdateEnvironmentRequest
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironment(environmentId *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(environmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - UpdateGatewayRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteResponse
func (client *Client) UpdateGatewayRouteWithOptions(gatewayId *string, gatewayRouteId *string, request *UpdateGatewayRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendConfig)) {
		body["backendConfig"] = request.BackendConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DomainConfig)) {
		body["domainConfig"] = request.DomainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Match)) {
		body["match"] = request.Match
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/http-routes/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayRouteId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由
//
// @param request - UpdateGatewayRouteRequest
//
// @return UpdateGatewayRouteResponse
func (client *Client) UpdateGatewayRoute(gatewayId *string, gatewayRouteId *string, request *UpdateGatewayRouteRequest) (_result *UpdateGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayRouteResponse{}
	_body, _err := client.UpdateGatewayRouteWithOptions(gatewayId, gatewayRouteId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateGatewayServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceResponse
func (client *Client) UpdateGatewayServiceWithOptions(gatewayId *string, gatewayServiceId *string, request *UpdateGatewayServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Addresses)) {
		body["addresses"] = request.Addresses
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateGatewayServiceRequest
//
// @return UpdateGatewayServiceResponse
func (client *Client) UpdateGatewayService(gatewayId *string, gatewayServiceId *string, request *UpdateGatewayServiceRequest) (_result *UpdateGatewayServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayServiceResponse{}
	_body, _err := client.UpdateGatewayServiceWithOptions(gatewayId, gatewayServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务版本
//
// @param request - UpdateGatewayServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceVersionResponse
func (client *Client) UpdateGatewayServiceVersionWithOptions(gatewayId *string, gatewayServiceId *string, name *string, request *UpdateGatewayServiceVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayServiceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayServiceVersion"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-services/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayServiceId)) + "/service-versions/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务版本
//
// @param request - UpdateGatewayServiceVersionRequest
//
// @return UpdateGatewayServiceVersionResponse
func (client *Client) UpdateGatewayServiceVersion(gatewayId *string, gatewayServiceId *string, name *string, request *UpdateGatewayServiceVersionRequest) (_result *UpdateGatewayServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayServiceVersionResponse{}
	_body, _err := client.UpdateGatewayServiceVersionWithOptions(gatewayId, gatewayServiceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新HTTP API
//
// @param request - UpdateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiResponse
func (client *Client) UpdateHttpApiWithOptions(httpApiId *string, request *UpdateHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		body["protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.VersionConfig)) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新HTTP API
//
// @param request - UpdateHttpApiRequest
//
// @return UpdateHttpApiResponse
func (client *Client) UpdateHttpApi(httpApiId *string, request *UpdateHttpApiRequest) (_result *UpdateHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiResponse{}
	_body, _err := client.UpdateHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Operation
//
// @param request - UpdateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiOperationResponse
func (client *Client) UpdateHttpApiOperationWithOptions(httpApiId *string, operationId *string, request *UpdateHttpApiOperationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHttpApiOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["operation"] = request.Operation
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHttpApiOperation"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/operations/" + tea.StringValue(openapiutil.GetEncodeParam(operationId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Operation
//
// @param request - UpdateHttpApiOperationRequest
//
// @return UpdateHttpApiOperationResponse
func (client *Client) UpdateHttpApiOperation(httpApiId *string, operationId *string, request *UpdateHttpApiOperationRequest) (_result *UpdateHttpApiOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiOperationResponse{}
	_body, _err := client.UpdateHttpApiOperationWithOptions(httpApiId, operationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务来源
//
// @param request - UpdateServiceSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceSourceResponse
func (client *Client) UpdateServiceSourceWithOptions(gatewayId *string, serviceSourceId *string, request *UpdateServiceSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sServiceSourceConfig)) {
		body["k8sServiceSourceConfig"] = request.K8sServiceSourceConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceSource"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/service-sources/" + tea.StringValue(openapiutil.GetEncodeParam(serviceSourceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务来源
//
// @param request - UpdateServiceSourceRequest
//
// @return UpdateServiceSourceResponse
func (client *Client) UpdateServiceSource(gatewayId *string, serviceSourceId *string, request *UpdateServiceSourceRequest) (_result *UpdateServiceSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceSourceResponse{}
	_body, _err := client.UpdateServiceSourceWithOptions(gatewayId, serviceSourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

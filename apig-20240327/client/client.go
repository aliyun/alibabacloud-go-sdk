// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AiServiceConfig struct {
	// example:
	//
	// https://dashscope.aliyun.com
	Address           *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKeys           []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	EnableHealthCheck *bool     `json:"enableHealthCheck,omitempty" xml:"enableHealthCheck,omitempty"`
	Protocols         []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s AiServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s AiServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfig) SetAddress(v string) *AiServiceConfig {
	s.Address = &v
	return s
}

func (s *AiServiceConfig) SetApiKeys(v []*string) *AiServiceConfig {
	s.ApiKeys = v
	return s
}

func (s *AiServiceConfig) SetEnableHealthCheck(v bool) *AiServiceConfig {
	s.EnableHealthCheck = &v
	return s
}

func (s *AiServiceConfig) SetProtocols(v []*string) *AiServiceConfig {
	s.Protocols = v
	return s
}

func (s *AiServiceConfig) SetProvider(v string) *AiServiceConfig {
	s.Provider = &v
	return s
}

type AkSkIdentityConfig struct {
	Ak           *string `json:"ak,omitempty" xml:"ak,omitempty"`
	GenerateMode *string `json:"generateMode,omitempty" xml:"generateMode,omitempty"`
	Sk           *string `json:"sk,omitempty" xml:"sk,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AkSkIdentityConfig) String() string {
	return tea.Prettify(s)
}

func (s AkSkIdentityConfig) GoString() string {
	return s.String()
}

func (s *AkSkIdentityConfig) SetAk(v string) *AkSkIdentityConfig {
	s.Ak = &v
	return s
}

func (s *AkSkIdentityConfig) SetGenerateMode(v string) *AkSkIdentityConfig {
	s.GenerateMode = &v
	return s
}

func (s *AkSkIdentityConfig) SetSk(v string) *AkSkIdentityConfig {
	s.Sk = &v
	return s
}

func (s *AkSkIdentityConfig) SetType(v string) *AkSkIdentityConfig {
	s.Type = &v
	return s
}

type ApiKeyIdentityConfig struct {
	Apikey       *string                           `json:"apikey,omitempty" xml:"apikey,omitempty"`
	ApikeySource *ApiKeyIdentityConfigApikeySource `json:"apikeySource,omitempty" xml:"apikeySource,omitempty" type:"Struct"`
	GenerateMode *string                           `json:"generateMode,omitempty" xml:"generateMode,omitempty"`
	Type         *string                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApiKeyIdentityConfig) String() string {
	return tea.Prettify(s)
}

func (s ApiKeyIdentityConfig) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfig) SetApikey(v string) *ApiKeyIdentityConfig {
	s.Apikey = &v
	return s
}

func (s *ApiKeyIdentityConfig) SetApikeySource(v *ApiKeyIdentityConfigApikeySource) *ApiKeyIdentityConfig {
	s.ApikeySource = v
	return s
}

func (s *ApiKeyIdentityConfig) SetGenerateMode(v string) *ApiKeyIdentityConfig {
	s.GenerateMode = &v
	return s
}

func (s *ApiKeyIdentityConfig) SetType(v string) *ApiKeyIdentityConfig {
	s.Type = &v
	return s
}

type ApiKeyIdentityConfigApikeySource struct {
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ApiKeyIdentityConfigApikeySource) String() string {
	return tea.Prettify(s)
}

func (s ApiKeyIdentityConfigApikeySource) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfigApikeySource) SetSource(v string) *ApiKeyIdentityConfigApikeySource {
	s.Source = &v
	return s
}

func (s *ApiKeyIdentityConfigApikeySource) SetValue(v string) *ApiKeyIdentityConfigApikeySource {
	s.Value = &v
	return s
}

type ApiRouteConflictInfo struct {
	Conflicts  []*ApiRouteConflictInfoConflicts `json:"conflicts,omitempty" xml:"conflicts,omitempty" type:"Repeated"`
	DomainInfo *ApiRouteConflictInfoDomainInfo  `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfo) SetConflicts(v []*ApiRouteConflictInfoConflicts) *ApiRouteConflictInfo {
	s.Conflicts = v
	return s
}

func (s *ApiRouteConflictInfo) SetDomainInfo(v *ApiRouteConflictInfoDomainInfo) *ApiRouteConflictInfo {
	s.DomainInfo = v
	return s
}

type ApiRouteConflictInfoConflicts struct {
	Details         []*ApiRouteConflictInfoConflictsDetails       `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	EnvironmentInfo *ApiRouteConflictInfoConflictsEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	ResourceId      *string                                       `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName    *string                                       `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType    *string                                       `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	RouteInfo       *ApiRouteConflictInfoConflictsRouteInfo       `json:"routeInfo,omitempty" xml:"routeInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflicts) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflicts) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflicts) SetDetails(v []*ApiRouteConflictInfoConflictsDetails) *ApiRouteConflictInfoConflicts {
	s.Details = v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetEnvironmentInfo(v *ApiRouteConflictInfoConflictsEnvironmentInfo) *ApiRouteConflictInfoConflicts {
	s.EnvironmentInfo = v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceId(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceId = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceName(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceName = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetResourceType(v string) *ApiRouteConflictInfoConflicts {
	s.ResourceType = &v
	return s
}

func (s *ApiRouteConflictInfoConflicts) SetRouteInfo(v *ApiRouteConflictInfoConflictsRouteInfo) *ApiRouteConflictInfoConflicts {
	s.RouteInfo = v
	return s
}

type ApiRouteConflictInfoConflictsDetails struct {
	ConflictingMatch *ApiRouteConflictInfoConflictsDetailsConflictingMatch `json:"conflictingMatch,omitempty" xml:"conflictingMatch,omitempty" type:"Struct"`
	DetectedMatch    *ApiRouteConflictInfoConflictsDetailsDetectedMatch    `json:"detectedMatch,omitempty" xml:"detectedMatch,omitempty" type:"Struct"`
	Level            *string                                               `json:"level,omitempty" xml:"level,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetails) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetails) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetails) SetConflictingMatch(v *ApiRouteConflictInfoConflictsDetailsConflictingMatch) *ApiRouteConflictInfoConflictsDetails {
	s.ConflictingMatch = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetails) SetDetectedMatch(v *ApiRouteConflictInfoConflictsDetailsDetectedMatch) *ApiRouteConflictInfoConflictsDetails {
	s.DetectedMatch = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetails) SetLevel(v string) *ApiRouteConflictInfoConflictsDetails {
	s.Level = &v
	return s
}

type ApiRouteConflictInfoConflictsDetailsConflictingMatch struct {
	Match         *HttpRouteMatch                                                    `json:"match,omitempty" xml:"match,omitempty"`
	OperationInfo *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatch) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatch) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) SetMatch(v *HttpRouteMatch) *ApiRouteConflictInfoConflictsDetailsConflictingMatch {
	s.Match = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatch) SetOperationInfo(v *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) *ApiRouteConflictInfoConflictsDetailsConflictingMatch {
	s.OperationInfo = v
	return s
}

type ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetName(v string) *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo) SetOperationId(v string) *ApiRouteConflictInfoConflictsDetailsConflictingMatchOperationInfo {
	s.OperationId = &v
	return s
}

type ApiRouteConflictInfoConflictsDetailsDetectedMatch struct {
	Match         *HttpRouteMatch                                                 `json:"match,omitempty" xml:"match,omitempty"`
	OperationInfo *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo `json:"operationInfo,omitempty" xml:"operationInfo,omitempty" type:"Struct"`
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatch) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatch) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) SetMatch(v *HttpRouteMatch) *ApiRouteConflictInfoConflictsDetailsDetectedMatch {
	s.Match = v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatch) SetOperationInfo(v *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) *ApiRouteConflictInfoConflictsDetailsDetectedMatch {
	s.OperationInfo = v
	return s
}

type ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetName(v string) *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo) SetOperationId(v string) *ApiRouteConflictInfoConflictsDetailsDetectedMatchOperationInfo {
	s.OperationId = &v
	return s
}

type ApiRouteConflictInfoConflictsEnvironmentInfo struct {
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiRouteConflictInfoConflictsEnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) SetEnvironmentId(v string) *ApiRouteConflictInfoConflictsEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsEnvironmentInfo) SetName(v string) *ApiRouteConflictInfoConflictsEnvironmentInfo {
	s.Name = &v
	return s
}

type ApiRouteConflictInfoConflictsRouteInfo struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s ApiRouteConflictInfoConflictsRouteInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoConflictsRouteInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) SetName(v string) *ApiRouteConflictInfoConflictsRouteInfo {
	s.Name = &v
	return s
}

func (s *ApiRouteConflictInfoConflictsRouteInfo) SetRouteId(v string) *ApiRouteConflictInfoConflictsRouteInfo {
	s.RouteId = &v
	return s
}

type ApiRouteConflictInfoDomainInfo struct {
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ApiRouteConflictInfoDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiRouteConflictInfoDomainInfo) GoString() string {
	return s.String()
}

func (s *ApiRouteConflictInfoDomainInfo) SetDomainId(v string) *ApiRouteConflictInfoDomainInfo {
	s.DomainId = &v
	return s
}

func (s *ApiRouteConflictInfoDomainInfo) SetName(v string) *ApiRouteConflictInfoDomainInfo {
	s.Name = &v
	return s
}

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

type AuthorizationResourceInfo struct {
	EnvironmentId    *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	ParentResourceId *string `json:"parentResourceId,omitempty" xml:"parentResourceId,omitempty"`
	ResourceId       *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s AuthorizationResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s AuthorizationResourceInfo) GoString() string {
	return s.String()
}

func (s *AuthorizationResourceInfo) SetEnvironmentId(v string) *AuthorizationResourceInfo {
	s.EnvironmentId = &v
	return s
}

func (s *AuthorizationResourceInfo) SetParentResourceId(v string) *AuthorizationResourceInfo {
	s.ParentResourceId = &v
	return s
}

func (s *AuthorizationResourceInfo) SetResourceId(v string) *AuthorizationResourceInfo {
	s.ResourceId = &v
	return s
}

type Backend struct {
	// example:
	//
	// Single
	Scene    *string            `json:"scene,omitempty" xml:"scene,omitempty"`
	Services []*BackendServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s Backend) String() string {
	return tea.Prettify(s)
}

func (s Backend) GoString() string {
	return s.String()
}

func (s *Backend) SetScene(v string) *Backend {
	s.Scene = &v
	return s
}

func (s *Backend) SetServices(v []*BackendServices) *Backend {
	s.Services = v
	return s
}

type BackendServices struct {
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
	// service-cq2bmmdlhtgj***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BackendServices) String() string {
	return tea.Prettify(s)
}

func (s BackendServices) GoString() string {
	return s.String()
}

func (s *BackendServices) SetName(v string) *BackendServices {
	s.Name = &v
	return s
}

func (s *BackendServices) SetPort(v int32) *BackendServices {
	s.Port = &v
	return s
}

func (s *BackendServices) SetProtocol(v string) *BackendServices {
	s.Protocol = &v
	return s
}

func (s *BackendServices) SetServiceId(v string) *BackendServices {
	s.ServiceId = &v
	return s
}

func (s *BackendServices) SetVersion(v string) *BackendServices {
	s.Version = &v
	return s
}

func (s *BackendServices) SetWeight(v int32) *BackendServices {
	s.Weight = &v
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
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *DomainInfo) SetResourceGroupId(v string) *DomainInfo {
	s.ResourceGroupId = &v
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
	Alias           *string      `json:"alias,omitempty" xml:"alias,omitempty"`
	CreateTimestamp *int64       `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Default         *bool        `json:"default,omitempty" xml:"default,omitempty"`
	Description     *string      `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentId   *string      `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo     *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	Name            *string      `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string          `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *EnvironmentInfo) SetResourceGroupId(v string) *EnvironmentInfo {
	s.ResourceGroupId = &v
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
	EngineVersion *string             `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	GatewayId     *string             `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Name          *string             `json:"name,omitempty" xml:"name,omitempty"`
	VpcInfo       *GatewayInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s GatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayInfo) GoString() string {
	return s.String()
}

func (s *GatewayInfo) SetEngineVersion(v string) *GatewayInfo {
	s.EngineVersion = &v
	return s
}

func (s *GatewayInfo) SetGatewayId(v string) *GatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *GatewayInfo) SetName(v string) *GatewayInfo {
	s.Name = &v
	return s
}

func (s *GatewayInfo) SetVpcInfo(v *GatewayInfoVpcInfo) *GatewayInfo {
	s.VpcInfo = v
	return s
}

type GatewayInfoVpcInfo struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GatewayInfoVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s GatewayInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *GatewayInfoVpcInfo) SetName(v string) *GatewayInfoVpcInfo {
	s.Name = &v
	return s
}

func (s *GatewayInfoVpcInfo) SetVpcId(v string) *GatewayInfoVpcInfo {
	s.VpcId = &v
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

type HttpApiApiInfo struct {
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// example:
	//
	// /v1
	BasePath      *string                       `json:"basePath,omitempty" xml:"basePath,omitempty"`
	DeployConfigs []*HttpApiDeployConfig        `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	Description   *string                       `json:"description,omitempty" xml:"description,omitempty"`
	Environments  []*HttpApiApiInfoEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// example:
	//
	// api-xxx
	HttpApiId   *string                    `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	IngressInfo *HttpApiApiInfoIngressInfo `json:"ingressInfo,omitempty" xml:"ingressInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name      *string   `json:"name,omitempty" xml:"name,omitempty"`
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// Rest
	Type        *string             `json:"type,omitempty" xml:"type,omitempty"`
	VersionInfo *HttpApiVersionInfo `json:"versionInfo,omitempty" xml:"versionInfo,omitempty"`
}

func (s HttpApiApiInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfo) SetAiProtocols(v []*string) *HttpApiApiInfo {
	s.AiProtocols = v
	return s
}

func (s *HttpApiApiInfo) SetBasePath(v string) *HttpApiApiInfo {
	s.BasePath = &v
	return s
}

func (s *HttpApiApiInfo) SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiApiInfo {
	s.DeployConfigs = v
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

func (s *HttpApiApiInfo) SetIngressInfo(v *HttpApiApiInfoIngressInfo) *HttpApiApiInfo {
	s.IngressInfo = v
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

func (s *HttpApiApiInfo) SetResourceGroupId(v string) *HttpApiApiInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *HttpApiApiInfo) SetType(v string) *HttpApiApiInfo {
	s.Type = &v
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
	BackendType   *string              `json:"backendType,omitempty" xml:"backendType,omitempty"`
	CustomDomains []*HttpApiDomainInfo `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	// example:
	//
	// Deployed
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// example:
	//
	// env-xxx
	EnvironmentId *string                                `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpApiApiInfoEnvironmentsGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name           *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	ServiceConfigs []*HttpApiApiInfoEnvironmentsServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	SubDomains     []*HttpApiApiInfoEnvironmentsSubDomains     `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
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

func (s *HttpApiApiInfoEnvironments) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiApiInfoEnvironments {
	s.CustomDomains = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetDeployStatus(v string) *HttpApiApiInfoEnvironments {
	s.DeployStatus = &v
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

func (s *HttpApiApiInfoEnvironments) SetServiceConfigs(v []*HttpApiApiInfoEnvironmentsServiceConfigs) *HttpApiApiInfoEnvironments {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetSubDomains(v []*HttpApiApiInfoEnvironmentsSubDomains) *HttpApiApiInfoEnvironments {
	s.SubDomains = v
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
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
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

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetServiceId(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.ServiceId = &v
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

type HttpApiApiInfoEnvironmentsSubDomains struct {
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
	// Internet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsSubDomains) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsSubDomains) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetDomainId(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.DomainId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetName(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetNetworkType(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.NetworkType = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetProtocol(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.Protocol = &v
	return s
}

type HttpApiApiInfoIngressInfo struct {
	EnvironmentInfo   *HttpApiApiInfoIngressInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	IngressClass      *string                                   `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	K8sClusterInfo    *HttpApiApiInfoIngressInfoK8sClusterInfo  `json:"k8sClusterInfo,omitempty" xml:"k8sClusterInfo,omitempty" type:"Struct"`
	OverrideIngressIp *bool                                     `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	SourceId          *string                                   `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	WatchNamespace    *string                                   `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s HttpApiApiInfoIngressInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoIngressInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfo) SetEnvironmentInfo(v *HttpApiApiInfoIngressInfoEnvironmentInfo) *HttpApiApiInfoIngressInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetIngressClass(v string) *HttpApiApiInfoIngressInfo {
	s.IngressClass = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetK8sClusterInfo(v *HttpApiApiInfoIngressInfoK8sClusterInfo) *HttpApiApiInfoIngressInfo {
	s.K8sClusterInfo = v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetOverrideIngressIp(v bool) *HttpApiApiInfoIngressInfo {
	s.OverrideIngressIp = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetSourceId(v string) *HttpApiApiInfoIngressInfo {
	s.SourceId = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetWatchNamespace(v string) *HttpApiApiInfoIngressInfo {
	s.WatchNamespace = &v
	return s
}

type HttpApiApiInfoIngressInfoEnvironmentInfo struct {
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s HttpApiApiInfoIngressInfoEnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoIngressInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfoEnvironmentInfo) SetEnvironmentId(v string) *HttpApiApiInfoIngressInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

type HttpApiApiInfoIngressInfoK8sClusterInfo struct {
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s HttpApiApiInfoIngressInfoK8sClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpApiApiInfoIngressInfoK8sClusterInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfoK8sClusterInfo) SetClusterId(v string) *HttpApiApiInfoIngressInfoK8sClusterInfo {
	s.ClusterId = &v
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

type HttpApiDeployConfig struct {
	// example:
	//
	// true
	AutoDeploy *bool `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	// example:
	//
	// SingleService
	BackendScene    *string   `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	CustomDomainIds []*string `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentId  *string                              `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	PolicyConfigs  []*HttpApiDeployConfigPolicyConfigs  `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	ServiceConfigs []*HttpApiDeployConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiDeployConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpApiDeployConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfig) SetAutoDeploy(v bool) *HttpApiDeployConfig {
	s.AutoDeploy = &v
	return s
}

func (s *HttpApiDeployConfig) SetBackendScene(v string) *HttpApiDeployConfig {
	s.BackendScene = &v
	return s
}

func (s *HttpApiDeployConfig) SetCustomDomainIds(v []*string) *HttpApiDeployConfig {
	s.CustomDomainIds = v
	return s
}

func (s *HttpApiDeployConfig) SetEnvironmentId(v string) *HttpApiDeployConfig {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiDeployConfig) SetPolicyConfigs(v []*HttpApiDeployConfigPolicyConfigs) *HttpApiDeployConfig {
	s.PolicyConfigs = v
	return s
}

func (s *HttpApiDeployConfig) SetServiceConfigs(v []*HttpApiDeployConfigServiceConfigs) *HttpApiDeployConfig {
	s.ServiceConfigs = v
	return s
}

type HttpApiDeployConfigPolicyConfigs struct {
	AiFallbackConfig *HttpApiDeployConfigPolicyConfigsAiFallbackConfig `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// AiFallback
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigs) SetAiFallbackConfig(v *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) *HttpApiDeployConfigPolicyConfigs {
	s.AiFallbackConfig = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) SetEnable(v bool) *HttpApiDeployConfigPolicyConfigs {
	s.Enable = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) SetType(v string) *HttpApiDeployConfigPolicyConfigs {
	s.Type = &v
	return s
}

type HttpApiDeployConfigPolicyConfigsAiFallbackConfig struct {
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfig) String() string {
	return tea.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) SetServiceIds(v []*string) *HttpApiDeployConfigPolicyConfigsAiFallbackConfig {
	s.ServiceIds = v
	return s
}

type HttpApiDeployConfigServiceConfigs struct {
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// qwen-*
	ModelNamePattern *string `json:"modelNamePattern,omitempty" xml:"modelNamePattern,omitempty"`
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// 100
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiDeployConfigServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s HttpApiDeployConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigServiceConfigs) SetModelName(v string) *HttpApiDeployConfigServiceConfigs {
	s.ModelName = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetModelNamePattern(v string) *HttpApiDeployConfigServiceConfigs {
	s.ModelNamePattern = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetServiceId(v string) *HttpApiDeployConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetWeight(v int64) *HttpApiDeployConfigServiceConfigs {
	s.Weight = &v
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
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *HttpApiInfoByName) SetType(v string) *HttpApiInfoByName {
	s.Type = &v
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

type HttpRoute struct {
	Backend         *Backend                  `json:"backend,omitempty" xml:"backend,omitempty"`
	CreateTimestamp *int64                    `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	DeployStatus    *string                   `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	Description     *string                   `json:"description,omitempty" xml:"description,omitempty"`
	DomainInfos     []*HttpRouteDomainInfos   `json:"domainInfos,omitempty" xml:"domainInfos,omitempty" type:"Repeated"`
	EnvironmentInfo *HttpRouteEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	Match           *HttpRouteMatch           `json:"match,omitempty" xml:"match,omitempty"`
	Name            *string                   `json:"name,omitempty" xml:"name,omitempty"`
	RouteId         *string                   `json:"routeId,omitempty" xml:"routeId,omitempty"`
	UpdateTimestamp *int64                    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s HttpRoute) String() string {
	return tea.Prettify(s)
}

func (s HttpRoute) GoString() string {
	return s.String()
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

func (s *HttpRoute) SetMatch(v *HttpRouteMatch) *HttpRoute {
	s.Match = v
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

type HttpRouteDomainInfos struct {
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpRouteDomainInfos) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteDomainInfos) GoString() string {
	return s.String()
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

type HttpRouteEnvironmentInfo struct {
	Alias         *string                               `json:"alias,omitempty" xml:"alias,omitempty"`
	EnvironmentId *string                               `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpRouteEnvironmentInfoGatewayInfo  `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	Name          *string                               `json:"name,omitempty" xml:"name,omitempty"`
	SubDomains    []*HttpRouteEnvironmentInfoSubDomains `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
}

func (s HttpRouteEnvironmentInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteEnvironmentInfo) GoString() string {
	return s.String()
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

type HttpRouteEnvironmentInfoGatewayInfo struct {
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpRouteEnvironmentInfoGatewayInfo) String() string {
	return tea.Prettify(s)
}

func (s HttpRouteEnvironmentInfoGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) SetGatewayId(v string) *HttpRouteEnvironmentInfoGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpRouteEnvironmentInfoGatewayInfo) SetName(v string) *HttpRouteEnvironmentInfoGatewayInfo {
	s.Name = &v
	return s
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
	return tea.Prettify(s)
}

func (s HttpRouteEnvironmentInfoSubDomains) GoString() string {
	return s.String()
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

type JwtIdentityConfig struct {
	Jwks             *string                            `json:"jwks,omitempty" xml:"jwks,omitempty"`
	JwtPayloadConfig *JwtIdentityConfigJwtPayloadConfig `json:"jwtPayloadConfig,omitempty" xml:"jwtPayloadConfig,omitempty" type:"Struct"`
	JwtTokenConfig   *JwtIdentityConfigJwtTokenConfig   `json:"jwtTokenConfig,omitempty" xml:"jwtTokenConfig,omitempty" type:"Struct"`
	SecretType       *string                            `json:"secretType,omitempty" xml:"secretType,omitempty"`
	Type             *string                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s JwtIdentityConfig) String() string {
	return tea.Prettify(s)
}

func (s JwtIdentityConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfig) SetJwks(v string) *JwtIdentityConfig {
	s.Jwks = &v
	return s
}

func (s *JwtIdentityConfig) SetJwtPayloadConfig(v *JwtIdentityConfigJwtPayloadConfig) *JwtIdentityConfig {
	s.JwtPayloadConfig = v
	return s
}

func (s *JwtIdentityConfig) SetJwtTokenConfig(v *JwtIdentityConfigJwtTokenConfig) *JwtIdentityConfig {
	s.JwtTokenConfig = v
	return s
}

func (s *JwtIdentityConfig) SetSecretType(v string) *JwtIdentityConfig {
	s.SecretType = &v
	return s
}

func (s *JwtIdentityConfig) SetType(v string) *JwtIdentityConfig {
	s.Type = &v
	return s
}

type JwtIdentityConfigJwtPayloadConfig struct {
	PayloadKeyName  *string `json:"payloadKeyName,omitempty" xml:"payloadKeyName,omitempty"`
	PayloadKeyValue *string `json:"payloadKeyValue,omitempty" xml:"payloadKeyValue,omitempty"`
}

func (s JwtIdentityConfigJwtPayloadConfig) String() string {
	return tea.Prettify(s)
}

func (s JwtIdentityConfigJwtPayloadConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfigJwtPayloadConfig) SetPayloadKeyName(v string) *JwtIdentityConfigJwtPayloadConfig {
	s.PayloadKeyName = &v
	return s
}

func (s *JwtIdentityConfigJwtPayloadConfig) SetPayloadKeyValue(v string) *JwtIdentityConfigJwtPayloadConfig {
	s.PayloadKeyValue = &v
	return s
}

type JwtIdentityConfigJwtTokenConfig struct {
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
	Pass     *bool   `json:"pass,omitempty" xml:"pass,omitempty"`
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	Prefix   *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s JwtIdentityConfigJwtTokenConfig) String() string {
	return tea.Prettify(s)
}

func (s JwtIdentityConfigJwtTokenConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfigJwtTokenConfig) SetKey(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Key = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPass(v bool) *JwtIdentityConfigJwtTokenConfig {
	s.Pass = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPosition(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Position = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPrefix(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Prefix = &v
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

type Service struct {
	Addresses       []*string        `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	AiServiceConfig *AiServiceConfig `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	CreateTimestamp *int64           `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// publich
	GroupName    *string             `json:"groupName,omitempty" xml:"groupName,omitempty"`
	HealthCheck  *ServiceHealthCheck `json:"healthCheck,omitempty" xml:"healthCheck,omitempty"`
	HealthStatus *string             `json:"healthStatus,omitempty" xml:"healthStatus,omitempty"`
	Name         *string             `json:"name,omitempty" xml:"name,omitempty"`
	Namespace    *string             `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Ports        []*ServicePorts     `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// rg-xxx
	ResourceGroupId    *string   `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ServiceId          *string   `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	SourceType         *string   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UnhealthyEndpoints []*string `json:"unhealthyEndpoints,omitempty" xml:"unhealthyEndpoints,omitempty" type:"Repeated"`
	UpdateTimestamp    *int64    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s Service) String() string {
	return tea.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) SetAddresses(v []*string) *Service {
	s.Addresses = v
	return s
}

func (s *Service) SetAiServiceConfig(v *AiServiceConfig) *Service {
	s.AiServiceConfig = v
	return s
}

func (s *Service) SetCreateTimestamp(v int64) *Service {
	s.CreateTimestamp = &v
	return s
}

func (s *Service) SetGatewayId(v string) *Service {
	s.GatewayId = &v
	return s
}

func (s *Service) SetGroupName(v string) *Service {
	s.GroupName = &v
	return s
}

func (s *Service) SetHealthCheck(v *ServiceHealthCheck) *Service {
	s.HealthCheck = v
	return s
}

func (s *Service) SetHealthStatus(v string) *Service {
	s.HealthStatus = &v
	return s
}

func (s *Service) SetName(v string) *Service {
	s.Name = &v
	return s
}

func (s *Service) SetNamespace(v string) *Service {
	s.Namespace = &v
	return s
}

func (s *Service) SetPorts(v []*ServicePorts) *Service {
	s.Ports = v
	return s
}

func (s *Service) SetProtocol(v string) *Service {
	s.Protocol = &v
	return s
}

func (s *Service) SetQualifier(v string) *Service {
	s.Qualifier = &v
	return s
}

func (s *Service) SetResourceGroupId(v string) *Service {
	s.ResourceGroupId = &v
	return s
}

func (s *Service) SetServiceId(v string) *Service {
	s.ServiceId = &v
	return s
}

func (s *Service) SetSourceType(v string) *Service {
	s.SourceType = &v
	return s
}

func (s *Service) SetUnhealthyEndpoints(v []*string) *Service {
	s.UnhealthyEndpoints = v
	return s
}

func (s *Service) SetUpdateTimestamp(v int64) *Service {
	s.UpdateTimestamp = &v
	return s
}

type ServicePorts struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Port     *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ServicePorts) String() string {
	return tea.Prettify(s)
}

func (s ServicePorts) GoString() string {
	return s.String()
}

func (s *ServicePorts) SetName(v string) *ServicePorts {
	s.Name = &v
	return s
}

func (s *ServicePorts) SetPort(v int32) *ServicePorts {
	s.Port = &v
	return s
}

func (s *ServicePorts) SetProtocol(v string) *ServicePorts {
	s.Protocol = &v
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
	// Description of the security group rule.
	//
	// example:
	//
	// 商品中心访问安全组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Port ranges.
	PortRanges []*string `json:"portRanges,omitempty" xml:"portRanges,omitempty" type:"Repeated"`
	// Security group ID.
	//
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// CA Certificate Identifier.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	// Certificate Unique Identifier.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	// Set the HTTPS protocol type, whether to enable forced HTTPS redirection.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// HTTP/2 settings.
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// Domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol type supported by the domain.
	//
	// - HTTP: Supports only HTTP protocol.
	//
	// - HTTPS: Supports only HTTPS protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol        *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Maximum TLS protocol version, supports up to TLS 1.3.
	//
	// example:
	//
	// TLS1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// Minimum TLS protocol version, supports down to TLS 1.0.
	//
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

func (s *CreateDomainRequest) SetCaCertIdentifier(v string) *CreateDomainRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *CreateDomainRequest) SetCertIdentifier(v string) *CreateDomainRequest {
	s.CertIdentifier = &v
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

func (s *CreateDomainRequest) SetResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceGroupId = &v
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
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Domain ID.
	//
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
	// Environment alias.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试环境
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// Description of the environment, which can include information such as the purpose of the environment and its owner.
	//
	// example:
	//
	// 这是xxx的xx项目测试环境
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Gateway ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Environment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *CreateEnvironmentRequest) SetResourceGroupId(v string) *CreateEnvironmentRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateEnvironmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreateEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Environment ID.
	//
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

type CreateHttpApiRequest struct {
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// Base path of the API, which must start with a \\"/\\".
	//
	// example:
	//
	// /v1
	BasePath      *string                `json:"basePath,omitempty" xml:"basePath,omitempty"`
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// Description of the API.
	//
	// example:
	//
	// 测试专用API。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Configuration information for the HTTP Ingress API.
	IngressConfig *CreateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// Name of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// List of API access protocols.
	Protocols       []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	ResourceGroupId *string   `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Type of the HTTP API.
	//
	// - Http
	//
	// - Rest
	//
	// - WebSocket
	//
	// - HttpIngress
	//
	// example:
	//
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Versioning configuration for the API.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s CreateHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRequest) SetAiProtocols(v []*string) *CreateHttpApiRequest {
	s.AiProtocols = v
	return s
}

func (s *CreateHttpApiRequest) SetBasePath(v string) *CreateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *CreateHttpApiRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRequest {
	s.DeployConfigs = v
	return s
}

func (s *CreateHttpApiRequest) SetDescription(v string) *CreateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *CreateHttpApiRequest) SetIngressConfig(v *CreateHttpApiRequestIngressConfig) *CreateHttpApiRequest {
	s.IngressConfig = v
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

func (s *CreateHttpApiRequest) SetResourceGroupId(v string) *CreateHttpApiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHttpApiRequest) SetType(v string) *CreateHttpApiRequest {
	s.Type = &v
	return s
}

func (s *CreateHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiRequest {
	s.VersionConfig = v
	return s
}

type CreateHttpApiRequestIngressConfig struct {
	// Environment ID.
	//
	// example:
	//
	// env-cq146allhtgk***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Ingress Class being listened to.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// Whether to update the address in the Ingress Status.
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// Source ID.
	//
	// example:
	//
	// src-crdddallhtgtr***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// Namespace being watched.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s CreateHttpApiRequestIngressConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRequestIngressConfig) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRequestIngressConfig) SetEnvironmentId(v string) *CreateHttpApiRequestIngressConfig {
	s.EnvironmentId = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetIngressClass(v string) *CreateHttpApiRequestIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetOverrideIngressIp(v bool) *CreateHttpApiRequestIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetSourceId(v string) *CreateHttpApiRequestIngressConfig {
	s.SourceId = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetWatchNamespace(v string) *CreateHttpApiRequestIngressConfig {
	s.WatchNamespace = &v
	return s
}

type CreateHttpApiResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API information.
	Data *CreateHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// HTTP API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// Name of the API.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// List of operation definitions.
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Operation information.
	Data *CreateHttpApiOperationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// Operation information.
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
	// Operation ID.
	//
	// example:
	//
	// op-xxx
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

type DeleteDomainResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the request chain.
	//
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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

type DeleteHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message,
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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

type GetDomainRequest struct {
	WithStatistics *bool `json:"withStatistics,omitempty" xml:"withStatistics,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetWithStatistics(v bool) *GetDomainRequest {
	s.WithStatistics = &v
	return s
}

type GetDomainResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *GetDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Encryption algorithm name
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Cloud Shield CA certificate identity.
	//
	// example:
	//
	// 223576-cn-hangzhou
	CaCertIndentifier *string `json:"caCertIndentifier,omitempty" xml:"caCertIndentifier,omitempty"`
	// Cloud Shield certificate identity.
	//
	// example:
	//
	// 123576-cn-hangzhou
	CertIndentifier *string `json:"certIndentifier,omitempty" xml:"certIndentifier,omitempty"`
	// Certificate name
	//
	// example:
	//
	// test-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// Where it was created from.
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// Creation timestamp.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Whether it is the default domain.
	//
	// example:
	//
	// false
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// Domain ID.
	//
	// example:
	//
	// d-cq1m3utlhtgvgkv7sitg
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// Setting for HTTPS protocol type, whether to enable forced HTTPS redirection.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// HTTP/2 setting.
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// Certificate issuer.
	//
	// example:
	//
	// Alibaba
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// Domain name.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Certificate expiration time.
	//
	// example:
	//
	// 1719386834548
	NotAfterTimstamp *int64 `json:"notAfterTimstamp,omitempty" xml:"notAfterTimstamp,omitempty"`
	// Certificate effective time.
	//
	// example:
	//
	// 1719386834548
	NotBeforeTimestamp *int64 `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	// The protocol types supported by the domain.
	//
	// - HTTP: Supports only HTTP protocol.
	//
	// - HTTPS: Supports only HTTPS protocol.
	//
	// example:
	//
	// HTTP
	Protocol        *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// All domain names bound to the certificate.
	//
	// example:
	//
	// aliyun.com
	Sans           *string                                  `json:"sans,omitempty" xml:"sans,omitempty"`
	StatisticsInfo *GetDomainResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
	// Maximum TLS protocol version, supports up to TLS 1.3.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// Minimum TLS protocol version, supports down to TLS 1.0.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
	// Update timestamp.
	//
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

func (s *GetDomainResponseBodyData) SetResourceGroupId(v string) *GetDomainResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDomainResponseBodyData) SetSans(v string) *GetDomainResponseBodyData {
	s.Sans = &v
	return s
}

func (s *GetDomainResponseBodyData) SetStatisticsInfo(v *GetDomainResponseBodyDataStatisticsInfo) *GetDomainResponseBodyData {
	s.StatisticsInfo = v
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

type GetDomainResponseBodyDataStatisticsInfo struct {
	ResourceStatistics []*ResourceStatistic `json:"resourceStatistics,omitempty" xml:"resourceStatistics,omitempty" type:"Repeated"`
	TotalCount         *string              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDomainResponseBodyDataStatisticsInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBodyDataStatisticsInfo) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDataStatisticsInfo) SetResourceStatistics(v []*ResourceStatistic) *GetDomainResponseBodyDataStatisticsInfo {
	s.ResourceStatistics = v
	return s
}

func (s *GetDomainResponseBodyDataStatisticsInfo) SetTotalCount(v string) *GetDomainResponseBodyDataStatisticsInfo {
	s.TotalCount = &v
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

type GetEnvironmentRequest struct {
	WithStatistics *bool `json:"withStatistics,omitempty" xml:"withStatistics,omitempty"`
	// Option for vpc info.
	WithVpcInfo *bool `json:"withVpcInfo,omitempty" xml:"withVpcInfo,omitempty"`
}

func (s GetEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *GetEnvironmentRequest) SetWithStatistics(v bool) *GetEnvironmentRequest {
	s.WithStatistics = &v
	return s
}

func (s *GetEnvironmentRequest) SetWithVpcInfo(v bool) *GetEnvironmentRequest {
	s.WithVpcInfo = &v
	return s
}

type GetEnvironmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *GetEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Environment alias.
	//
	// example:
	//
	// 测试环境
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// Creation timestamp.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Whether it is the default environment.
	//
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// Environment description.
	//
	// example:
	//
	// 这是xxx的xx项目测试环境
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway information
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// Environment name.
	//
	// example:
	//
	// test
	Name            *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string                                       `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	StatisticsInfo  *GetEnvironmentResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
	// List of subdomains.
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// Update timestamp.
	//
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

func (s *GetEnvironmentResponseBodyData) SetResourceGroupId(v string) *GetEnvironmentResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetStatisticsInfo(v *GetEnvironmentResponseBodyDataStatisticsInfo) *GetEnvironmentResponseBodyData {
	s.StatisticsInfo = v
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

type GetEnvironmentResponseBodyDataStatisticsInfo struct {
	ResourceStatistics []*ResourceStatistic `json:"resourceStatistics,omitempty" xml:"resourceStatistics,omitempty" type:"Repeated"`
	TotalCount         *int32               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetEnvironmentResponseBodyDataStatisticsInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyDataStatisticsInfo) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) SetResourceStatistics(v []*ResourceStatistic) *GetEnvironmentResponseBodyDataStatisticsInfo {
	s.ResourceStatistics = v
	return s
}

func (s *GetEnvironmentResponseBodyDataStatisticsInfo) SetTotalCount(v int32) *GetEnvironmentResponseBodyDataStatisticsInfo {
	s.TotalCount = &v
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *GetGatewayResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// Charge type
	//
	// - POSTPAY: Postpaid (pay-as-you-go)
	//
	// - PREPAY: Prepaid (subscription)
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// Source of gateway creation:
	//
	// - Console: Console.
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// Creation timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// List of environments associated with the gateway.
	Environments []*GetGatewayResponseBodyDataEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// Expiration timestamp for subscription. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// Gateway ID.
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// List of entry addresses for the gateway.
	LoadBalancers []*GetGatewayResponseBodyDataLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// Gateway name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Number of gateway instance nodes.
	//
	// example:
	//
	// 2
	Replicas        *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The security group of the gateway.
	SecurityGroup *GetGatewayResponseBodyDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// Gateway specification:
	//
	// - apigw.small.x1: Small specification.
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// Gateway status:
	//
	// - Running: Running.
	//
	// - Creating: Creating.
	//
	// - CreateFailed: Creation failed.
	//
	// - Upgrading: Upgrading.
	//
	// - UpgradeFailed: Upgrade failed.
	//
	// - Restarting: Restarting.
	//
	// - RestartFailed: Restart failed.
	//
	// - Deleting: Deleting.
	//
	// - DeleteFailed: Deletion failed.
	//
	// example:
	//
	// Running
	Status *string                           `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*GetGatewayResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Target version of the gateway. When it is inconsistent with the current version, an upgrade can be performed.
	//
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// Update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	// The virtual switch associated with the gateway.
	VSwitch *GetGatewayResponseBodyDataVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// Gateway version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The VPC (Virtual Private Cloud) associated with the gateway.
	Vpc *GetGatewayResponseBodyDataVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	// List of availability zones associated with the gateway.
	Zones []*GetGatewayResponseBodyDataZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
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

func (s *GetGatewayResponseBodyData) SetResourceGroupId(v string) *GetGatewayResponseBodyData {
	s.ResourceGroupId = &v
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

func (s *GetGatewayResponseBodyData) SetTags(v []*GetGatewayResponseBodyDataTags) *GetGatewayResponseBodyData {
	s.Tags = v
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
	// The environment alias.
	//
	// example:
	//
	// 默认环境
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// Environment ID.
	//
	// example:
	//
	// env-cp9uhudlht***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The environment name。
	//
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
	// The address of the load balancer.
	//
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The IP version of the protocol:
	//
	// - ipv4: IPv4 type.
	//
	// - ipv6: IPv6 type.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// Load balancer address type:
	//
	// - Internet: Public.
	//
	// - Intranet: Private.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// Whether it is the default entry address for the gateway.
	//
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// Load balancer ID.
	//
	// example:
	//
	// nlb-xoh3pghru7c***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The provision mode of the load balancer for the gateway:
	//
	// - Managed: Managed by the Cloud Native API Gateway.
	//
	// example:
	//
	// Managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// List of listening ports.
	Ports []*GetGatewayResponseBodyDataLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// The status of the load balancer:
	//
	// - Ready: Available.
	//
	// - NotCreate: Not associated with an instance.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of load balancer:
	//
	// - NLB: Network Load Balancer.
	//
	// - CLB: Classic Load Balancer.
	//
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
	// Port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Protocol:
	//
	// - TCP
	//
	// - UDP
	//
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
	// Security group name.
	//
	// example:
	//
	// APIG-sg-gw-cq7ke5ll***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Security group ID.
	//
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

type GetGatewayResponseBodyDataTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetGatewayResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataTags) SetKey(v string) *GetGatewayResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetGatewayResponseBodyDataTags) SetValue(v string) *GetGatewayResponseBodyDataTags {
	s.Value = &v
	return s
}

type GetGatewayResponseBodyDataVSwitch struct {
	// Virtual switch name.
	//
	// example:
	//
	// 杭州VPC虚拟交换机
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Virtual switch ID.
	//
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
	// VPC gateway name.
	//
	// example:
	//
	// 杭州VPC
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// VPC network ID.
	//
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
	// Availability zone name.
	//
	// example:
	//
	// 杭州可用区E
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Virtual switch.
	VSwitch *GetGatewayResponseBodyDataZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// Availability zone ID.
	//
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
	// Virtual switch name.
	//
	// example:
	//
	// 杭州VPC虚拟交换机
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Virtual switch ID.
	//
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

type GetHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API information.
	Data *HttpApiApiInfo `json:"data,omitempty" xml:"data,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Operation information.
	Data *HttpApiOperationInfo `json:"data,omitempty" xml:"data,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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

type GetHttpApiRouteResponseBody struct {
	// example:
	//
	// Ok
	Code *string    `json:"code,omitempty" xml:"code,omitempty"`
	Data *HttpRoute `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHttpApiRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiRouteResponseBody) SetCode(v string) *GetHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetData(v *HttpRoute) *GetHttpApiRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetMessage(v string) *GetHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiRouteResponseBody) SetRequestId(v string) *GetHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

type GetHttpApiRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiRouteResponse) SetHeaders(v map[string]*string) *GetHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiRouteResponse) SetStatusCode(v int32) *GetHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiRouteResponse) SetBody(v *GetHttpApiRouteResponseBody) *GetHttpApiRouteResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// Gateway Id.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Domain name, fuzzy search.
	//
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// Page number, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page, default is 10.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *ListDomainsRequest) SetResourceGroupId(v string) *ListDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListDomainsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *ListDomainsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// List of domain information.
	Items []*DomainInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items.
	//
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
	// Environment alias, fuzzy search.
	//
	// example:
	//
	// 测试
	AliasLike *string `json:"aliasLike,omitempty" xml:"aliasLike,omitempty"`
	// Gateway ID, exact search.
	//
	// example:
	//
	// gw-cptv6ktlhtgnqr73h8d1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Gateway name, fuzzy search.
	//
	// example:
	//
	// test-gw
	GatewayNameLike *string `json:"gatewayNameLike,omitempty" xml:"gatewayNameLike,omitempty"`
	// Environment name, fuzzy search.
	//
	// example:
	//
	// test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// Page number, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *ListEnvironmentsRequest) SetResourceGroupId(v string) *ListEnvironmentsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Paged query environment list response.
	Data *ListEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the call chain.
	//
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
	// List of environment information.
	Items []*EnvironmentInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items.
	//
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

type ListGatewaysRequest struct {
	// Query exactly by gateway ID.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Keyword, search with full match, case-insensitive.
	//
	// example:
	//
	// dev
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// Query exactly by gateway name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize        *int32                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string                    `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*ListGatewaysRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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

func (s *ListGatewaysRequest) SetResourceGroupId(v string) *ListGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewaysRequest) SetTags(v []*ListGatewaysRequestTags) *ListGatewaysRequest {
	s.Tags = v
	return s
}

type ListGatewaysRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewaysRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysRequestTags) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequestTags) SetKey(v string) *ListGatewaysRequestTags {
	s.Key = &v
	return s
}

func (s *ListGatewaysRequestTags) SetValue(v string) *ListGatewaysRequestTags {
	s.Value = &v
	return s
}

type ListGatewaysShrinkRequest struct {
	// Query exactly by gateway ID.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Keyword, search with full match, case-insensitive.
	//
	// example:
	//
	// dev
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// Query exactly by gateway name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	TagsShrink      *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListGatewaysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysShrinkRequest) SetGatewayId(v string) *ListGatewaysShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetKeyword(v string) *ListGatewaysShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetName(v string) *ListGatewaysShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetPageNumber(v int32) *ListGatewaysShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetPageSize(v int32) *ListGatewaysShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetResourceGroupId(v string) *ListGatewaysShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewaysShrinkRequest) SetTagsShrink(v string) *ListGatewaysShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListGatewaysResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Result of the gateway list query.
	Data *ListGatewaysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// Gateway list
	Items []*ListGatewaysResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total count.
	//
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
	// Charge type
	//
	// - POSTPAY: Postpaid (pay-as-you-go)
	//
	// - PREPAY: Prepaid (subscription)
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// Source of gateway creation:
	//
	// - Console: Console.
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// Creation timestamp, in milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Expiration timestamp for the prepaid (annual or monthly) plan. Unit: milliseconds.
	//
	// example:
	//
	// 172086834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// Gateway ID.
	//
	// example:
	//
	// gw-cpv54p5***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// List of gateway entry addresses.
	LoadBalancers []*ListGatewaysResponseBodyDataItemsLoadBalancers `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	// Gateway name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Number of gateway instance nodes.
	//
	// example:
	//
	// 2
	Replicas        *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Security Group.
	SecurityGroup *ListGatewaysResponseBodyDataItemsSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// Gateway specification:
	//
	// - apigw.small.x1: Small specification.
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// Gateway status:
	//
	// - Running: Running.
	//
	// - Creating: Creating.
	//
	// - CreateFailed: Creation failed.
	//
	// - Upgrading: Upgrading.
	//
	// - UpgradeFailed: Upgrade failed.
	//
	// - Restarting: Restarting.
	//
	// - RestartFailed: Restart failed.
	//
	// - Deleting: Deleting.
	//
	// - DeleteFailed: Deletion failed.
	//
	// example:
	//
	// Running
	Status *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListGatewaysResponseBodyDataItemsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Target version of the gateway. When it is inconsistent with `version`, a version upgrade can be performed.
	//
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// Update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64                                    `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	VSwitch         *ListGatewaysResponseBodyDataItemsVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// Gateway version.
	//
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

func (s *ListGatewaysResponseBodyDataItems) SetResourceGroupId(v string) *ListGatewaysResponseBodyDataItems {
	s.ResourceGroupId = &v
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

func (s *ListGatewaysResponseBodyDataItems) SetTags(v []*ListGatewaysResponseBodyDataItemsTags) *ListGatewaysResponseBodyDataItems {
	s.Tags = v
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
	// Load balancer address.
	//
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// IP version:
	//
	// - ipv4: IPv4.
	//
	// - ipv6: IPv6.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// Load balancer address type:
	//
	// - Internet: Public network.
	//
	// - Intranet: Private network.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// Indicates whether this is the default entry address for the gateway.
	//
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// Load balancer ID.
	//
	// example:
	//
	// nlb-xqwioje1c91r***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// Mode of load balancer provision for the gateway:
	//
	// - Managed: Managed by the Cloud Native API Gateway.
	//
	// example:
	//
	// Managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// List of listening ports.
	Ports []*ListGatewaysResponseBodyDataItemsLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// Status of the load balancer:
	//
	// - Ready: Available.
	//
	// - NotCreate: No associated instance.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Type of load balancer for the gateway:
	//
	// - NLB: Network Load Balancer.
	//
	// - CLB: Classic Load Balancer.
	//
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
	// Port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Protocol:
	//
	// - TCP
	//
	// - UDP
	//
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
	// The Security Group ID.
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

type ListGatewaysResponseBodyDataItemsTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsTags) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsTags) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetTagKey(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.TagKey = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetTagValue(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.TagValue = &v
	return s
}

type ListGatewaysResponseBodyDataItemsVSwitch struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxx
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
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxxxx
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
	// The vSwitch.
	VSwitch *ListGatewaysResponseBodyDataItemsZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The ID of the current zone.
	//
	// example:
	//
	// cn-hangzhou
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxx
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
	// Filter the operation list based on a specific consumer authorization rule ID, and the interface list in the response only contains authorized operations.
	//
	// example:
	//
	// cas-xxx
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// List interfaces by Method.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// Search operations by exact name.
	//
	// example:
	//
	// getUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Search operations by name prefix.
	//
	// example:
	//
	// GetUser
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// Page number, starting from 1, default is 1 if not specified.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, valid range [1, 100], default is 10 if not specified.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Search operations by path prefix.
	//
	// example:
	//
	// /v1
	PathLike *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
	// Each operation information in the response carries a list of authorization rules for the specified consumer under the specified environment ID. The withConsumerInEnvironmentId field needs to be additionally specified.
	//
	// example:
	//
	// env-xxx
	WithConsumerInEnvironmentId *string `json:"withConsumerInEnvironmentId,omitempty" xml:"withConsumerInEnvironmentId,omitempty"`
	// Each operation information in the response carries a list of authorization rules for the specified consumer under the specified environment ID. The withConsumerInEnvironmentId field needs to be additionally specified.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
}

func (s ListHttpApiOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsRequest) SetConsumerAuthorizationRuleId(v string) *ListHttpApiOperationsRequest {
	s.ConsumerAuthorizationRuleId = &v
	return s
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

func (s *ListHttpApiOperationsRequest) SetWithConsumerInEnvironmentId(v string) *ListHttpApiOperationsRequest {
	s.WithConsumerInEnvironmentId = &v
	return s
}

func (s *ListHttpApiOperationsRequest) SetWithConsumerInfoById(v string) *ListHttpApiOperationsRequest {
	s.WithConsumerInfoById = &v
	return s
}

type ListHttpApiOperationsResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// List of operations.
	Data *ListHttpApiOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// List of operations.
	Items []*HttpApiOperationInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total count.
	//
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
	// Cloud-native API Gateway ID.
	//
	// example:
	//
	// gw-cq2avtllh****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Search keyword, supports fuzzy search by API name or exact search by API ID.
	//
	// example:
	//
	// test-
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// Exact search by name.
	//
	// example:
	//
	// login
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page number, starting from 1, default is 1 if not provided.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, valid range [1, 100], default is 10 if not provided.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Type of HTTP API. Supports multiple types, separated by ",".
	//
	// - Http
	//
	// - Rest
	//
	// - WebSocket
	//
	// - HttpIngress
	//
	// example:
	//
	// Http,Rest
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
	// Each API information in the response carries consumer authentication policy information under the specified environment ID.
	//
	// example:
	//
	// env-xxx
	WithAuthPolicyInEnvironmentId *string `json:"withAuthPolicyInEnvironmentId,omitempty" xml:"withAuthPolicyInEnvironmentId,omitempty"`
	// Each API information in the response carries a list of authorization rules for the specified consumer ID.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	WithEnvironmentInfo  *bool   `json:"withEnvironmentInfo,omitempty" xml:"withEnvironmentInfo,omitempty"`
}

func (s ListHttpApisRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApisRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApisRequest) SetGatewayId(v string) *ListHttpApisRequest {
	s.GatewayId = &v
	return s
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

func (s *ListHttpApisRequest) SetResourceGroupId(v string) *ListHttpApisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHttpApisRequest) SetTypes(v string) *ListHttpApisRequest {
	s.Types = &v
	return s
}

func (s *ListHttpApisRequest) SetWithAuthPolicyInEnvironmentId(v string) *ListHttpApisRequest {
	s.WithAuthPolicyInEnvironmentId = &v
	return s
}

func (s *ListHttpApisRequest) SetWithConsumerInfoById(v string) *ListHttpApisRequest {
	s.WithConsumerInfoById = &v
	return s
}

func (s *ListHttpApisRequest) SetWithEnvironmentInfo(v bool) *ListHttpApisRequest {
	s.WithEnvironmentInfo = &v
	return s
}

type ListHttpApisResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API list.
	Data *ListHttpApisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// API information.
	Items []*HttpApiInfoByName `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total count.
	//
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

type UpdateDomainRequest struct {
	// Cloud Shield CA certificate identifier.
	//
	// example:
	//
	// 123455-cn-hangzhou
	CaCertIndentifier *string `json:"caCertIndentifier,omitempty" xml:"caCertIndentifier,omitempty"`
	// Cloud Shield certificate identifier.
	//
	// example:
	//
	// 123458-cn-hangzhou
	CertIndentifier *string `json:"certIndentifier,omitempty" xml:"certIndentifier,omitempty"`
	// Set the HTTPS protocol type, whether to enable forced HTTPS redirection.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// HTTP/2 settings.
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// The protocol type supported by the domain.
	//
	// - HTTP: Supports only HTTP protocol.
	//
	// - HTTPS: Supports only HTTPS protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Maximum TLS protocol version, supports up to TLS 1.3.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// Minimum TLS protocol version, supports down to TLS 1.0.
	//
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
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *UpdateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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
	// Deploy revision id.
	//
	// example:
	//
	// apr-xxx
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
	// Environment alias.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试环境
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// Description of the environment, which can include information such as the purpose of the environment and its users.
	//
	// example:
	//
	// 这是xx的xx项目测试环境
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
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
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

type UpdateHttpApiRequest struct {
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// Base path of the API, which must start with a \\"/\\".
	//
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath      *string                `json:"basePath,omitempty" xml:"basePath,omitempty"`
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// API description.
	//
	// example:
	//
	// 更新API描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Configuration information for the HTTP Ingress API.
	IngressConfig *UpdateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// List of API access protocols.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// API versioning configuration.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s UpdateHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRequest) SetAiProtocols(v []*string) *UpdateHttpApiRequest {
	s.AiProtocols = v
	return s
}

func (s *UpdateHttpApiRequest) SetBasePath(v string) *UpdateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *UpdateHttpApiRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRequest {
	s.DeployConfigs = v
	return s
}

func (s *UpdateHttpApiRequest) SetDescription(v string) *UpdateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRequest) SetIngressConfig(v *UpdateHttpApiRequestIngressConfig) *UpdateHttpApiRequest {
	s.IngressConfig = v
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

type UpdateHttpApiRequestIngressConfig struct {
	// Environment ID.
	//
	// example:
	//
	// env-cr6ql0tlhtgmc****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Ingress Class being listened to.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// Whether to update the address in the Ingress Status.
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// Source ID.
	//
	// example:
	//
	// src-crdddallhtgtr****
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// Watched namespace.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s UpdateHttpApiRequestIngressConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRequestIngressConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRequestIngressConfig) SetEnvironmentId(v string) *UpdateHttpApiRequestIngressConfig {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetIngressClass(v string) *UpdateHttpApiRequestIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetOverrideIngressIp(v bool) *UpdateHttpApiRequestIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetSourceId(v string) *UpdateHttpApiRequestIngressConfig {
	s.SourceId = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetWatchNamespace(v string) *UpdateHttpApiRequestIngressConfig {
	s.WatchNamespace = &v
	return s
}

type UpdateHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
	// operation definition.
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
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
// Authorize the security group for gateway to access services
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
// Authorize the security group for gateway to access services
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
// Create Domain
//
// Description:
//
// Create Domain.
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
	if !tea.BoolValue(util.IsUnset(request.CaCertIdentifier)) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		body["certIdentifier"] = request.CertIdentifier
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
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
// Create Domain
//
// Description:
//
// Create Domain.
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
// Description:
//
// Create environment.
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
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
// Description:
//
// Create environment.
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
// Create an API of HTTP type
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
	if !tea.BoolValue(util.IsUnset(request.AiProtocols)) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.DeployConfigs)) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IngressConfig)) {
		body["ingressConfig"] = request.IngressConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		body["protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
// Create an API of HTTP type
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
// Create an Operation for HTTP API
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
// Create an Operation for HTTP API
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
// Delete Gateway
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
// Delete Gateway
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
// Delete HTTP API
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
// Delete HTTP API
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
// Delete Operation
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
// Delete Operation
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
// Query domain details
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(domainId *string, request *GetDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WithStatistics)) {
		query["withStatistics"] = request.WithStatistics
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
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
// Query domain details
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(domainId *string, request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(domainId, request, headers, runtime)
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
// @param request - GetEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithOptions(environmentId *string, request *GetEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WithStatistics)) {
		query["withStatistics"] = request.WithStatistics
	}

	if !tea.BoolValue(util.IsUnset(request.WithVpcInfo)) {
		query["withVpcInfo"] = request.WithVpcInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
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
// @param request - GetEnvironmentRequest
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironment(environmentId *string, request *GetEnvironmentRequest) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(environmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get a gateway.
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
// Get a gateway.
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
// Read HttpApi
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
// Read HttpApi
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
// Get Operation
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
// Get Operation
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
// 获取HttpApi的路由详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiRouteResponse
func (client *Client) GetHttpApiRouteWithOptions(httpApiId *string, routeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHttpApiRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHttpApiRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/routes/" + tea.StringValue(openapiutil.GetEncodeParam(routeId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHttpApiRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取HttpApi的路由详情
//
// @return GetHttpApiRouteResponse
func (client *Client) GetHttpApiRoute(httpApiId *string, routeId *string) (_result *GetHttpApiRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiRouteResponse{}
	_body, _err := client.GetHttpApiRouteWithOptions(httpApiId, routeId, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
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
// Retrieve the list of created cloud-native gateways
//
// @param tmpReq - ListGatewaysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewaysResponse
func (client *Client) ListGatewaysWithOptions(tmpReq *ListGatewaysRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewaysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGatewaysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
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
// Retrieve the list of created cloud-native gateways
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
// List Operations
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
	if !tea.BoolValue(util.IsUnset(request.ConsumerAuthorizationRuleId)) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

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

	if !tea.BoolValue(util.IsUnset(request.WithConsumerInEnvironmentId)) {
		query["withConsumerInEnvironmentId"] = request.WithConsumerInEnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.WithConsumerInfoById)) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
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
// List Operations
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
// List HTTP APIs
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	if !tea.BoolValue(util.IsUnset(request.WithAuthPolicyInEnvironmentId)) {
		query["withAuthPolicyInEnvironmentId"] = request.WithAuthPolicyInEnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.WithConsumerInfoById)) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !tea.BoolValue(util.IsUnset(request.WithEnvironmentInfo)) {
		query["withEnvironmentInfo"] = request.WithEnvironmentInfo
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
// List HTTP APIs
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
// Update HTTP API
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
	if !tea.BoolValue(util.IsUnset(request.AiProtocols)) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.DeployConfigs)) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IngressConfig)) {
		body["ingressConfig"] = request.IngressConfig
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
// Update HTTP API
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
// Update Operation
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
// Update Operation
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

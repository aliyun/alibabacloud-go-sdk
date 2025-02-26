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
	ApikeySource *ApiKeyIdentityConfigApikeySource  `json:"apikeySource,omitempty" xml:"apikeySource,omitempty" type:"Struct"`
	Credentials  []*ApiKeyIdentityConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	Type         *string                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApiKeyIdentityConfig) String() string {
	return tea.Prettify(s)
}

func (s ApiKeyIdentityConfig) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfig) SetApikeySource(v *ApiKeyIdentityConfigApikeySource) *ApiKeyIdentityConfig {
	s.ApikeySource = v
	return s
}

func (s *ApiKeyIdentityConfig) SetCredentials(v []*ApiKeyIdentityConfigCredentials) *ApiKeyIdentityConfig {
	s.Credentials = v
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

type ApiKeyIdentityConfigCredentials struct {
	Apikey       *string `json:"apikey,omitempty" xml:"apikey,omitempty"`
	GenerateMode *string `json:"generateMode,omitempty" xml:"generateMode,omitempty"`
}

func (s ApiKeyIdentityConfigCredentials) String() string {
	return tea.Prettify(s)
}

func (s ApiKeyIdentityConfigCredentials) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfigCredentials) SetApikey(v string) *ApiKeyIdentityConfigCredentials {
	s.Apikey = &v
	return s
}

func (s *ApiKeyIdentityConfigCredentials) SetGenerateMode(v string) *ApiKeyIdentityConfigCredentials {
	s.GenerateMode = &v
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

type AuthConfig struct {
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
}

func (s AuthConfig) String() string {
	return tea.Prettify(s)
}

func (s AuthConfig) GoString() string {
	return s.String()
}

func (s *AuthConfig) SetAuthMode(v string) *AuthConfig {
	s.AuthMode = &v
	return s
}

func (s *AuthConfig) SetAuthType(v string) *AuthConfig {
	s.AuthType = &v
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

type ConsumerInfo struct {
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	Enable     *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ConsumerInfo) String() string {
	return tea.Prettify(s)
}

func (s ConsumerInfo) GoString() string {
	return s.String()
}

func (s *ConsumerInfo) SetConsumerId(v string) *ConsumerInfo {
	s.ConsumerId = &v
	return s
}

func (s *ConsumerInfo) SetEnable(v bool) *ConsumerInfo {
	s.Enable = &v
	return s
}

func (s *ConsumerInfo) SetName(v string) *ConsumerInfo {
	s.Name = &v
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
	AiProtocols []*string   `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	AuthConfig  *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// /v1
	BasePath      *string                       `json:"basePath,omitempty" xml:"basePath,omitempty"`
	DeployConfigs []*HttpApiDeployConfig        `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	Description   *string                       `json:"description,omitempty" xml:"description,omitempty"`
	EnabelAuth    *bool                         `json:"enabelAuth,omitempty" xml:"enabelAuth,omitempty"`
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

func (s *HttpApiApiInfo) SetAuthConfig(v *AuthConfig) *HttpApiApiInfo {
	s.AuthConfig = v
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

func (s *HttpApiApiInfo) SetEnabelAuth(v bool) *HttpApiApiInfo {
	s.EnabelAuth = &v
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
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableAuth  *bool   `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
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

func (s *HttpApiOperation) SetAuthConfig(v *AuthConfig) *HttpApiOperation {
	s.AuthConfig = v
	return s
}

func (s *HttpApiOperation) SetDescription(v string) *HttpApiOperation {
	s.Description = &v
	return s
}

func (s *HttpApiOperation) SetEnableAuth(v bool) *HttpApiOperation {
	s.EnableAuth = &v
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
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableAuth  *bool   `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
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

func (s *HttpApiOperationInfo) SetAuthConfig(v *AuthConfig) *HttpApiOperationInfo {
	s.AuthConfig = v
	return s
}

func (s *HttpApiOperationInfo) SetCreateTimestamp(v int64) *HttpApiOperationInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiOperationInfo) SetDescription(v string) *HttpApiOperationInfo {
	s.Description = &v
	return s
}

func (s *HttpApiOperationInfo) SetEnableAuth(v bool) *HttpApiOperationInfo {
	s.EnableAuth = &v
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

type ParentResourceInfo struct {
	ApiInfo      *HttpApiApiInfo `json:"apiInfo,omitempty" xml:"apiInfo,omitempty"`
	ResourceType *string         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ParentResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ParentResourceInfo) GoString() string {
	return s.String()
}

func (s *ParentResourceInfo) SetApiInfo(v *HttpApiApiInfo) *ParentResourceInfo {
	s.ApiInfo = v
	return s
}

func (s *ParentResourceInfo) SetResourceType(v string) *ParentResourceInfo {
	s.ResourceType = &v
	return s
}

type PluginClassInfo struct {
	Alias                      *string `json:"alias,omitempty" xml:"alias,omitempty"`
	ConfigExample              *string `json:"configExample,omitempty" xml:"configExample,omitempty"`
	Description                *string `json:"description,omitempty" xml:"description,omitempty"`
	ExecutePriority            *int32  `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	ExecuteStage               *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	ImageName                  *string `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InnerPlugin                *bool   `json:"innerPlugin,omitempty" xml:"innerPlugin,omitempty"`
	Mode                       *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Name                       *string `json:"name,omitempty" xml:"name,omitempty"`
	PluginClassId              *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	Source                     *string `json:"source,omitempty" xml:"source,omitempty"`
	SupportedMinGatewayVersion *string `json:"supportedMinGatewayVersion,omitempty" xml:"supportedMinGatewayVersion,omitempty"`
	Type                       *string `json:"type,omitempty" xml:"type,omitempty"`
	Version                    *string `json:"version,omitempty" xml:"version,omitempty"`
	VersionDescription         *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	WasmLanguage               *string `json:"wasmLanguage,omitempty" xml:"wasmLanguage,omitempty"`
	WasmUrl                    *string `json:"wasmUrl,omitempty" xml:"wasmUrl,omitempty"`
}

func (s PluginClassInfo) String() string {
	return tea.Prettify(s)
}

func (s PluginClassInfo) GoString() string {
	return s.String()
}

func (s *PluginClassInfo) SetAlias(v string) *PluginClassInfo {
	s.Alias = &v
	return s
}

func (s *PluginClassInfo) SetConfigExample(v string) *PluginClassInfo {
	s.ConfigExample = &v
	return s
}

func (s *PluginClassInfo) SetDescription(v string) *PluginClassInfo {
	s.Description = &v
	return s
}

func (s *PluginClassInfo) SetExecutePriority(v int32) *PluginClassInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PluginClassInfo) SetExecuteStage(v string) *PluginClassInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PluginClassInfo) SetImageName(v string) *PluginClassInfo {
	s.ImageName = &v
	return s
}

func (s *PluginClassInfo) SetInnerPlugin(v bool) *PluginClassInfo {
	s.InnerPlugin = &v
	return s
}

func (s *PluginClassInfo) SetMode(v string) *PluginClassInfo {
	s.Mode = &v
	return s
}

func (s *PluginClassInfo) SetName(v string) *PluginClassInfo {
	s.Name = &v
	return s
}

func (s *PluginClassInfo) SetPluginClassId(v string) *PluginClassInfo {
	s.PluginClassId = &v
	return s
}

func (s *PluginClassInfo) SetSource(v string) *PluginClassInfo {
	s.Source = &v
	return s
}

func (s *PluginClassInfo) SetSupportedMinGatewayVersion(v string) *PluginClassInfo {
	s.SupportedMinGatewayVersion = &v
	return s
}

func (s *PluginClassInfo) SetType(v string) *PluginClassInfo {
	s.Type = &v
	return s
}

func (s *PluginClassInfo) SetVersion(v string) *PluginClassInfo {
	s.Version = &v
	return s
}

func (s *PluginClassInfo) SetVersionDescription(v string) *PluginClassInfo {
	s.VersionDescription = &v
	return s
}

func (s *PluginClassInfo) SetWasmLanguage(v string) *PluginClassInfo {
	s.WasmLanguage = &v
	return s
}

func (s *PluginClassInfo) SetWasmUrl(v string) *PluginClassInfo {
	s.WasmUrl = &v
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

type ResourceInfo struct {
	ResourceId      *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
}

func (s ResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ResourceInfo) GoString() string {
	return s.String()
}

func (s *ResourceInfo) SetResourceId(v string) *ResourceInfo {
	s.ResourceId = &v
	return s
}

func (s *ResourceInfo) SetResourceName(v string) *ResourceInfo {
	s.ResourceName = &v
	return s
}

func (s *ResourceInfo) SetResourceType(v string) *ResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *ResourceInfo) SetResourceVersion(v string) *ResourceInfo {
	s.ResourceVersion = &v
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

type TlsCipherSuitesConfig struct {
	ConfigType     *string                                `json:"configType,omitempty" xml:"configType,omitempty"`
	TlsCipherSuite []*TlsCipherSuitesConfigTlsCipherSuite `json:"tlsCipherSuite,omitempty" xml:"tlsCipherSuite,omitempty" type:"Repeated"`
}

func (s TlsCipherSuitesConfig) String() string {
	return tea.Prettify(s)
}

func (s TlsCipherSuitesConfig) GoString() string {
	return s.String()
}

func (s *TlsCipherSuitesConfig) SetConfigType(v string) *TlsCipherSuitesConfig {
	s.ConfigType = &v
	return s
}

func (s *TlsCipherSuitesConfig) SetTlsCipherSuite(v []*TlsCipherSuitesConfigTlsCipherSuite) *TlsCipherSuitesConfig {
	s.TlsCipherSuite = v
	return s
}

type TlsCipherSuitesConfigTlsCipherSuite struct {
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	SupportVersions []*string `json:"supportVersions,omitempty" xml:"supportVersions,omitempty" type:"Repeated"`
}

func (s TlsCipherSuitesConfigTlsCipherSuite) String() string {
	return tea.Prettify(s)
}

func (s TlsCipherSuitesConfigTlsCipherSuite) GoString() string {
	return s.String()
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) SetName(v string) *TlsCipherSuitesConfigTlsCipherSuite {
	s.Name = &v
	return s
}

func (s *TlsCipherSuitesConfigTlsCipherSuite) SetSupportVersions(v []*string) *TlsCipherSuitesConfigTlsCipherSuite {
	s.SupportVersions = v
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

type ChangeResourceGroupRequest struct {
	// example:
	//
	// rg-aekzdrfx2xdnaja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// gw-ct4i14um1hkn0tpqfae0
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// apig
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetService(v string) *ChangeResourceGroupRequest {
	s.Service = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// example:
	//
	// 59F86F37-787A-52DB-9475-DB5A255517F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
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
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzoiafjtr7zyq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The TLS cipher suites config.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
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

func (s *CreateDomainRequest) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *CreateDomainRequest {
	s.TlsCipherSuitesConfig = v
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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmycs5expl7oq
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
	// The AI protocols list.
	AiProtocols []*string   `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	AuthConfig  *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// Base path of the API, which must start with a \\"/\\".
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The deploy configs.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// Description of the API.
	//
	// example:
	//
	// 测试专用API。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableAuth  *bool   `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
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
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzgvmlotionbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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

func (s *CreateHttpApiRequest) SetAuthConfig(v *AuthConfig) *CreateHttpApiRequest {
	s.AuthConfig = v
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

func (s *CreateHttpApiRequest) SetEnableAuth(v bool) *CreateHttpApiRequest {
	s.EnableAuth = &v
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

type CreateHttpApiRouteRequest struct {
	// Backend service configuration for the route.
	BackendConfig *CreateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Route description.
	//
	// example:
	//
	// 用户登陆路由。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubcv***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Route match rules.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// Route name.
	//
	// example:
	//
	// login
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateHttpApiRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequest) SetBackendConfig(v *CreateHttpApiRouteRequestBackendConfig) *CreateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetDescription(v string) *CreateHttpApiRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateHttpApiRouteRequest) SetDomainIds(v []*string) *CreateHttpApiRouteRequest {
	s.DomainIds = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetEnvironmentId(v string) *CreateHttpApiRouteRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateHttpApiRouteRequest) SetMatch(v *HttpRouteMatch) *CreateHttpApiRouteRequest {
	s.Match = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetName(v string) *CreateHttpApiRouteRequest {
	s.Name = &v
	return s
}

type CreateHttpApiRouteRequestBackendConfig struct {
	// Backend service scenario.
	//
	// - SingleService: Single service.
	//
	// - MultiServiceByRatio: Multiple services with ratio-based canary release.
	//
	// - Mock: Mock service.
	//
	// - Redirect: Redirect service.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// List of backend services.
	Services []*CreateHttpApiRouteRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s CreateHttpApiRouteRequestBackendConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequestBackendConfig) SetScene(v string) *CreateHttpApiRouteRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfig) SetServices(v []*CreateHttpApiRouteRequestBackendConfigServices) *CreateHttpApiRouteRequestBackendConfig {
	s.Services = v
	return s
}

type CreateHttpApiRouteRequestBackendConfigServices struct {
	// Service port, not provided for dynamic ports.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol:
	//
	// - HTTP.
	//
	// - HTTPS.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Service ID.
	//
	// example:
	//
	// svc-crbgq0dlhtgr***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service version, valid only in label-based scenarios.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Percentage value of traffic distribution.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateHttpApiRouteRequestBackendConfigServices) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetPort(v int32) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetProtocol(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetServiceId(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetVersion(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetWeight(v int32) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Weight = &v
	return s
}

type CreateHttpApiRouteResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreateHttpApiRouteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponseBody) SetCode(v string) *CreateHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetData(v *CreateHttpApiRouteResponseBodyData) *CreateHttpApiRouteResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetMessage(v string) *CreateHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiRouteResponseBody) SetRequestId(v string) *CreateHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

type CreateHttpApiRouteResponseBodyData struct {
	// Route ID。
	//
	// example:
	//
	// hr-cr82undlhtgrlej***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s CreateHttpApiRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponseBodyData) SetRouteId(v string) *CreateHttpApiRouteResponseBodyData {
	s.RouteId = &v
	return s
}

type CreateHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponse) SetHeaders(v map[string]*string) *CreateHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiRouteResponse) SetStatusCode(v int32) *CreateHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiRouteResponse) SetBody(v *CreateHttpApiRouteResponseBody) *CreateHttpApiRouteResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Timeout
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetClassName(v string) *CreatePolicyRequest {
	s.ClassName = &v
	return s
}

func (s *CreatePolicyRequest) SetConfig(v string) *CreatePolicyRequest {
	s.Config = &v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

type CreatePolicyResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreatePolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E7406754***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetCode(v string) *CreatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyResponseBody) SetData(v *CreatePolicyResponseBodyData) *CreatePolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyResponseBody) SetMessage(v string) *CreatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyResponseBodyData struct {
	// example:
	//
	// p-cq7l5s5lhtgi6qasr***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s CreatePolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBodyData) SetPolicyId(v string) *CreatePolicyResponseBodyData {
	s.PolicyId = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type CreatePolicyAttachmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// api-cu07jj6m1hkokaus***
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qas***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// p-cq787etlhtghrptjg***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s CreatePolicyAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentRequest) SetAttachResourceId(v string) *CreatePolicyAttachmentRequest {
	s.AttachResourceId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetAttachResourceType(v string) *CreatePolicyAttachmentRequest {
	s.AttachResourceType = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetEnvironmentId(v string) *CreatePolicyAttachmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetGatewayId(v string) *CreatePolicyAttachmentRequest {
	s.GatewayId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetPolicyId(v string) *CreatePolicyAttachmentRequest {
	s.PolicyId = &v
	return s
}

type CreatePolicyAttachmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreatePolicyAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C64***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePolicyAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponseBody) SetCode(v string) *CreatePolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetData(v *CreatePolicyAttachmentResponseBodyData) *CreatePolicyAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetMessage(v string) *CreatePolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyAttachmentResponseBody) SetRequestId(v string) *CreatePolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyAttachmentResponseBodyData struct {
	// example:
	//
	// pr-cqooju5lhtgquuj6***
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
}

func (s CreatePolicyAttachmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponseBodyData) SetPolicyAttachmentId(v string) *CreatePolicyAttachmentResponseBodyData {
	s.PolicyAttachmentId = &v
	return s
}

type CreatePolicyAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponse) SetHeaders(v map[string]*string) *CreatePolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyAttachmentResponse) SetStatusCode(v int32) *CreatePolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyAttachmentResponse) SetBody(v *CreatePolicyAttachmentResponseBody) *CreatePolicyAttachmentResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	// example:
	//
	// gw-cq7l5s5lhtg***
	GatewayId       *string                               `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ResourceGroupId *string                               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ServiceConfigs  []*CreateServiceRequestServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetGatewayId(v string) *CreateServiceRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateServiceRequest) SetResourceGroupId(v string) *CreateServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceConfigs(v []*CreateServiceRequestServiceConfigs) *CreateServiceRequest {
	s.ServiceConfigs = v
	return s
}

func (s *CreateServiceRequest) SetSourceType(v string) *CreateServiceRequest {
	s.SourceType = &v
	return s
}

type CreateServiceRequestServiceConfigs struct {
	Addresses       []*string        `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	AiServiceConfig *AiServiceConfig `json:"aiServiceConfig,omitempty" xml:"aiServiceConfig,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// user-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PUBLIC
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s CreateServiceRequestServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceConfigs) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceConfigs) SetAddresses(v []*string) *CreateServiceRequestServiceConfigs {
	s.Addresses = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetAiServiceConfig(v *AiServiceConfig) *CreateServiceRequestServiceConfigs {
	s.AiServiceConfig = v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetGroupName(v string) *CreateServiceRequestServiceConfigs {
	s.GroupName = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetName(v string) *CreateServiceRequestServiceConfigs {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetNamespace(v string) *CreateServiceRequestServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *CreateServiceRequestServiceConfigs) SetQualifier(v string) *CreateServiceRequestServiceConfigs {
	s.Qualifier = &v
	return s
}

type CreateServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854CD0D4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetCode(v string) *CreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceResponseBody) SetData(v *CreateServiceResponseBodyData) *CreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceResponseBody) SetMessage(v string) *CreateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceResponseBodyData struct {
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s CreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyData) SetServiceIds(v []*string) *CreateServiceResponseBodyData {
	s.ServiceIds = v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
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

type DeleteGatewaySecurityGroupRuleRequest struct {
	// Whether to cascade delete the security group rules.
	//
	// example:
	//
	// true
	CascadingDelete *bool `json:"cascadingDelete,omitempty" xml:"cascadingDelete,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleRequest) SetCascadingDelete(v bool) *DeleteGatewaySecurityGroupRuleRequest {
	s.CascadingDelete = &v
	return s
}

type DeleteGatewaySecurityGroupRuleResponseBody struct {
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
	// 8F94B3CC-F4BA-511E-8367-ECBBE486E595
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetCode(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetMessage(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponseBody) SetRequestId(v string) *DeleteGatewaySecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewaySecurityGroupRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewaySecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewaySecurityGroupRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewaySecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetStatusCode(v int32) *DeleteGatewaySecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySecurityGroupRuleResponse) SetBody(v *DeleteGatewaySecurityGroupRuleResponseBody) *DeleteGatewaySecurityGroupRuleResponse {
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

type DeleteHttpApiRouteResponseBody struct {
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
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiRouteResponseBody) SetCode(v string) *DeleteHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiRouteResponseBody) SetMessage(v string) *DeleteHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiRouteResponseBody) SetRequestId(v string) *DeleteHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiRouteResponse) SetHeaders(v map[string]*string) *DeleteHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiRouteResponse) SetStatusCode(v int32) *DeleteHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiRouteResponse) SetBody(v *DeleteHttpApiRouteResponseBody) *DeleteHttpApiRouteResponse {
	s.Body = v
	return s
}

type DeletePolicyResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E7406754***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) SetCode(v string) *DeletePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyResponseBody) SetMessage(v string) *DeletePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

type DeletePolicyAttachmentResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePolicyAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyAttachmentResponseBody) SetCode(v string) *DeletePolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyAttachmentResponseBody) SetMessage(v string) *DeletePolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyAttachmentResponseBody) SetRequestId(v string) *DeletePolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyAttachmentResponse) SetHeaders(v map[string]*string) *DeletePolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyAttachmentResponse) SetStatusCode(v int32) *DeletePolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyAttachmentResponse) SetBody(v *DeletePolicyAttachmentResponseBody) *DeletePolicyAttachmentResponse {
	s.Body = v
	return s
}

type DeployHttpApiRequest struct {
	RestApiConfig *DeployHttpApiRequestRestApiConfig `json:"restApiConfig,omitempty" xml:"restApiConfig,omitempty" type:"Struct"`
	// example:
	//
	// hr-cr82undlhtgrl***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s DeployHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiRequest) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequest) SetRestApiConfig(v *DeployHttpApiRequestRestApiConfig) *DeployHttpApiRequest {
	s.RestApiConfig = v
	return s
}

func (s *DeployHttpApiRequest) SetRouteId(v string) *DeployHttpApiRequest {
	s.RouteId = &v
	return s
}

type DeployHttpApiRequestRestApiConfig struct {
	Description *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Environment *DeployHttpApiRequestRestApiConfigEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// example:
	//
	// apr-xxx
	RevisionId *string `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
}

func (s DeployHttpApiRequestRestApiConfig) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfig) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfig) SetDescription(v string) *DeployHttpApiRequestRestApiConfig {
	s.Description = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetEnvironment(v *DeployHttpApiRequestRestApiConfigEnvironment) *DeployHttpApiRequestRestApiConfig {
	s.Environment = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetRevisionId(v string) *DeployHttpApiRequestRestApiConfig {
	s.RevisionId = &v
	return s
}

type DeployHttpApiRequestRestApiConfigEnvironment struct {
	// example:
	//
	// SingleService
	BackendScene    *string   `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	CustomDomainIds []*string `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId  *string                                                       `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	ServiceConfigs []*DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s DeployHttpApiRequestRestApiConfigEnvironment) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfigEnvironment) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetBackendScene(v string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.BackendScene = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetCustomDomainIds(v []*string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.CustomDomainIds = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetEnvironmentId(v string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.EnvironmentId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetServiceConfigs(v []*DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.ServiceConfigs = v
	return s
}

type DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs struct {
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
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
	// svc-cr6pk4tlhtgm58e***
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

func (s DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Match = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetPort(v int32) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Port = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetProtocol(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetServiceId(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetVersion(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Version = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetWeight(v int32) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Weight = &v
	return s
}

type DeployHttpApiResponseBody struct {
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
	// 0C2D1C68-0D93-5561-8EE6-FDB7BF067A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeployHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeployHttpApiResponseBody) SetCode(v string) *DeployHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeployHttpApiResponseBody) SetMessage(v string) *DeployHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeployHttpApiResponseBody) SetRequestId(v string) *DeployHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type DeployHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployHttpApiResponse) GoString() string {
	return s.String()
}

func (s *DeployHttpApiResponse) SetHeaders(v map[string]*string) *DeployHttpApiResponse {
	s.Headers = v
	return s
}

func (s *DeployHttpApiResponse) SetStatusCode(v int32) *DeployHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployHttpApiResponse) SetBody(v *DeployHttpApiResponseBody) *DeployHttpApiResponse {
	s.Body = v
	return s
}

type ExportHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API definition information.
	Data *ExportHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 4BACB05C-3FE2-588F-9148-700C5C026B74
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExportHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *ExportHttpApiResponseBody) SetCode(v string) *ExportHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *ExportHttpApiResponseBody) SetData(v *ExportHttpApiResponseBodyData) *ExportHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *ExportHttpApiResponseBody) SetMessage(v string) *ExportHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *ExportHttpApiResponseBody) SetRequestId(v string) *ExportHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type ExportHttpApiResponseBodyData struct {
	// Base64编码的API定义。
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
}

func (s ExportHttpApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExportHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExportHttpApiResponseBodyData) SetSpecContentBase64(v string) *ExportHttpApiResponseBodyData {
	s.SpecContentBase64 = &v
	return s
}

type ExportHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportHttpApiResponse) GoString() string {
	return s.String()
}

func (s *ExportHttpApiResponse) SetHeaders(v map[string]*string) *ExportHttpApiResponse {
	s.Headers = v
	return s
}

func (s *ExportHttpApiResponse) SetStatusCode(v int32) *ExportHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportHttpApiResponse) SetBody(v *ExportHttpApiResponseBody) *ExportHttpApiResponse {
	s.Body = v
	return s
}

type GetDashboardRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string                    `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	ApiId          *string                    `json:"apiId,omitempty" xml:"apiId,omitempty"`
	Filter         *GetDashboardRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// example:
	//
	// PLUGIN
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pls-dn82a9djd8z****
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// example:
	//
	// SLS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardRequest) SetAcceptLanguage(v string) *GetDashboardRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetDashboardRequest) SetApiId(v string) *GetDashboardRequest {
	s.ApiId = &v
	return s
}

func (s *GetDashboardRequest) SetFilter(v *GetDashboardRequestFilter) *GetDashboardRequest {
	s.Filter = v
	return s
}

func (s *GetDashboardRequest) SetName(v string) *GetDashboardRequest {
	s.Name = &v
	return s
}

func (s *GetDashboardRequest) SetPluginClassId(v string) *GetDashboardRequest {
	s.PluginClassId = &v
	return s
}

func (s *GetDashboardRequest) SetSource(v string) *GetDashboardRequest {
	s.Source = &v
	return s
}

type GetDashboardRequestFilter struct {
	// example:
	//
	// test-route
	RouteName *string `json:"routeName,omitempty" xml:"routeName,omitempty"`
}

func (s GetDashboardRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardRequestFilter) GoString() string {
	return s.String()
}

func (s *GetDashboardRequestFilter) SetRouteName(v string) *GetDashboardRequestFilter {
	s.RouteName = &v
	return s
}

type GetDashboardShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	ApiId          *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	FilterShrink   *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// PLUGIN
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pls-dn82a9djd8z****
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// example:
	//
	// SLS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetDashboardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardShrinkRequest) SetAcceptLanguage(v string) *GetDashboardShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetApiId(v string) *GetDashboardShrinkRequest {
	s.ApiId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetFilterShrink(v string) *GetDashboardShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetName(v string) *GetDashboardShrinkRequest {
	s.Name = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetPluginClassId(v string) *GetDashboardShrinkRequest {
	s.PluginClassId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetSource(v string) *GetDashboardShrinkRequest {
	s.Source = &v
	return s
}

type GetDashboardResponseBody struct {
	// example:
	//
	// 200
	Code *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetDashboardResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Ok
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetDashboardResponseBody) SetCode(v int32) *GetDashboardResponseBody {
	s.Code = &v
	return s
}

func (s *GetDashboardResponseBody) SetData(v *GetDashboardResponseBodyData) *GetDashboardResponseBody {
	s.Data = v
	return s
}

func (s *GetDashboardResponseBody) SetErrorCode(v string) *GetDashboardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDashboardResponseBody) SetMessage(v string) *GetDashboardResponseBody {
	s.Message = &v
	return s
}

func (s *GetDashboardResponseBody) SetRequestId(v string) *GetDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDashboardResponseBody) SetSuccess(v bool) *GetDashboardResponseBody {
	s.Success = &v
	return s
}

type GetDashboardResponseBodyData struct {
	// example:
	//
	// gw-co370icmjeu****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// PLUGIN
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// APIG Plugin
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://sls.console.aliyun.com/lognext/project/xxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDashboardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDashboardResponseBodyData) SetGatewayId(v string) *GetDashboardResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetName(v string) *GetDashboardResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetTitle(v string) *GetDashboardResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetDashboardResponseBodyData) SetUrl(v string) *GetDashboardResponseBodyData {
	s.Url = &v
	return s
}

type GetDashboardResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetDashboardResponse) SetHeaders(v map[string]*string) *GetDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetDashboardResponse) SetStatusCode(v int32) *GetDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDashboardResponse) SetBody(v *GetDashboardResponseBody) *GetDashboardResponse {
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
	Algorithm        *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	CertIdentifier   *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
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
	Sans                  *string                                  `json:"sans,omitempty" xml:"sans,omitempty"`
	StatisticsInfo        *GetDomainResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
	TlsCipherSuitesConfig *TlsCipherSuitesConfig                   `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
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

func (s *GetDomainResponseBodyData) SetCaCertIdentifier(v string) *GetDomainResponseBodyData {
	s.CaCertIdentifier = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCertIdentifier(v string) *GetDomainResponseBodyData {
	s.CertIdentifier = &v
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

func (s *GetDomainResponseBodyData) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *GetDomainResponseBodyData {
	s.TlsCipherSuitesConfig = v
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
	// Indicates whether to return online resource info.
	//
	// example:
	//
	// true
	WithStatistics *bool `json:"withStatistics,omitempty" xml:"withStatistics,omitempty"`
	// Option for vpc info.
	//
	// example:
	//
	// true
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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzzzntl5njbpi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Related resource information.
	StatisticsInfo *GetEnvironmentResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
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
	// The array of related resource information.
	ResourceStatistics []*ResourceStatistic `json:"resourceStatistics,omitempty" xml:"resourceStatistics,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2s3cvc4jzfxi
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
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The resource tags.
	Tags []*GetGatewayResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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
	// The tag key of the resource.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// zhangsan
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
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Route detail data.
	Data *HttpRoute `json:"data,omitempty" xml:"data,omitempty"`
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

type GetPolicyResponseBody struct {
	// example:
	//
	// Ok
	Code *string           `json:"code,omitempty" xml:"code,omitempty"`
	Data *PolicyDetailInfo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32A***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) SetCode(v string) *GetPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyResponseBody) SetData(v *PolicyDetailInfo) *GetPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyResponseBody) SetMessage(v string) *GetPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetHeaders(v map[string]*string) *GetPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyResponse) SetStatusCode(v int32) *GetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyResponse) SetBody(v *GetPolicyResponseBody) *GetPolicyResponse {
	s.Body = v
	return s
}

type GetPolicyAttachmentResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetPolicyAttachmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2C3B9A12-3868-5EB9-fBEA-F99E03DD1***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPolicyAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponseBody) SetCode(v string) *GetPolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetData(v *GetPolicyAttachmentResponseBodyData) *GetPolicyAttachmentResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetMessage(v string) *GetPolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyAttachmentResponseBody) SetRequestId(v string) *GetPolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyAttachmentResponseBodyData struct {
	// example:
	//
	// op-csbkd9llhtgqhqua***
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// example:
	//
	// Operation
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// env-cq7l5s5lhtgi6qa***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// pr-cqoojualhtgquuj***
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
	// example:
	//
	// p-cq7l5s5bblhtgi6qas***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s GetPolicyAttachmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyAttachmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponseBodyData) SetAttachResourceId(v string) *GetPolicyAttachmentResponseBodyData {
	s.AttachResourceId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetAttachResourceType(v string) *GetPolicyAttachmentResponseBodyData {
	s.AttachResourceType = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetConfig(v string) *GetPolicyAttachmentResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetEnvironmentId(v string) *GetPolicyAttachmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetGatewayId(v string) *GetPolicyAttachmentResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetPolicyAttachmentId(v string) *GetPolicyAttachmentResponseBodyData {
	s.PolicyAttachmentId = &v
	return s
}

func (s *GetPolicyAttachmentResponseBodyData) SetPolicyId(v string) *GetPolicyAttachmentResponseBodyData {
	s.PolicyId = &v
	return s
}

type GetPolicyAttachmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponse) SetHeaders(v map[string]*string) *GetPolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyAttachmentResponse) SetStatusCode(v int32) *GetPolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyAttachmentResponse) SetBody(v *GetPolicyAttachmentResponseBody) *GetPolicyAttachmentResponse {
	s.Body = v
	return s
}

type GetResourceOverviewResponseBody struct {
	// example:
	//
	// Ok
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetResourceOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD19A442-93C5-5C97-AFA0-B9C57EBD781B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetResourceOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBody) SetCode(v string) *GetResourceOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetData(v *GetResourceOverviewResponseBodyData) *GetResourceOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceOverviewResponseBody) SetMessage(v string) *GetResourceOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetRequestId(v string) *GetResourceOverviewResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceOverviewResponseBodyData struct {
	Api     *GetResourceOverviewResponseBodyDataApi     `json:"api,omitempty" xml:"api,omitempty" type:"Struct"`
	Gateway *GetResourceOverviewResponseBodyDataGateway `json:"gateway,omitempty" xml:"gateway,omitempty" type:"Struct"`
}

func (s GetResourceOverviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyData) SetApi(v *GetResourceOverviewResponseBodyDataApi) *GetResourceOverviewResponseBodyData {
	s.Api = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) SetGateway(v *GetResourceOverviewResponseBodyDataGateway) *GetResourceOverviewResponseBodyData {
	s.Gateway = v
	return s
}

type GetResourceOverviewResponseBodyDataApi struct {
	// example:
	//
	// 1
	PublishedCount *int64 `json:"publishedCount,omitempty" xml:"publishedCount,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataApi) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataApi) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataApi) SetPublishedCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.PublishedCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataApi) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.TotalCount = &v
	return s
}

type GetResourceOverviewResponseBodyDataGateway struct {
	// example:
	//
	// 1
	RunningCount *int64 `json:"runningCount,omitempty" xml:"runningCount,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataGateway) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataGateway) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetRunningCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.RunningCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.TotalCount = &v
	return s
}

type GetResourceOverviewResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponse) SetHeaders(v map[string]*string) *GetResourceOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetResourceOverviewResponse) SetStatusCode(v int32) *GetResourceOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceOverviewResponse) SetBody(v *GetResourceOverviewResponseBody) *GetResourceOverviewResponse {
	s.Body = v
	return s
}

type GetServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string  `json:"code,omitempty" xml:"code,omitempty"`
	Data *Service `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetCode(v string) *GetServiceResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceResponseBody) SetData(v *Service) *GetServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceResponseBody) SetMessage(v string) *GetServiceResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetTraceConfigRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
}

func (s GetTraceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetTraceConfigRequest) SetAcceptLanguage(v string) *GetTraceConfigRequest {
	s.AcceptLanguage = &v
	return s
}

type GetTraceConfigResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTraceConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32AF2C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTraceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponseBody) SetCode(v int32) *GetTraceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetData(v *GetTraceConfigResponseBodyData) *GetTraceConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetTraceConfigResponseBody) SetMessage(v string) *GetTraceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetRequestId(v string) *GetTraceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceConfigResponseBody) SetSuccess(v bool) *GetTraceConfigResponseBody {
	s.Success = &v
	return s
}

type GetTraceConfigResponseBodyData struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 50
	SampleRatio *int32 `json:"sampleRatio,omitempty" xml:"sampleRatio,omitempty"`
	// example:
	//
	// ss-co370icmjeu****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// 8090
	ServicePort *string `json:"servicePort,omitempty" xml:"servicePort,omitempty"`
	// example:
	//
	// SKYWALKING
	TraceType *string `json:"traceType,omitempty" xml:"traceType,omitempty"`
}

func (s GetTraceConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTraceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponseBodyData) SetEnable(v bool) *GetTraceConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetSampleRatio(v int32) *GetTraceConfigResponseBodyData {
	s.SampleRatio = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetServiceId(v string) *GetTraceConfigResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetServicePort(v string) *GetTraceConfigResponseBodyData {
	s.ServicePort = &v
	return s
}

func (s *GetTraceConfigResponseBodyData) SetTraceType(v string) *GetTraceConfigResponseBodyData {
	s.TraceType = &v
	return s
}

type GetTraceConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTraceConfigResponse) SetHeaders(v map[string]*string) *GetTraceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTraceConfigResponse) SetStatusCode(v int32) *GetTraceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceConfigResponse) SetBody(v *GetTraceConfigResponseBody) *GetTraceConfigResponse {
	s.Body = v
	return s
}

type ImportHttpApiRequest struct {
	// The description of the imported API. If not provided, it will be extracted from the API definition. The length is limited to 255 bytes.
	//
	// example:
	//
	// 测试专用API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Pre-import check. If enabled, only the check will be performed without importing the API.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The name of the imported API. If not provided, it will be extracted from the API definition file. If the API name and version configuration already exist, this import will update the existing API definition based on the `strategy` field.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3q4zjh7fkki
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Base64 encoded API definition, supporting OAS2.0 and OAS3.0 specifications, in YAML or JSON format. It has a higher priority than the `specFileUrl` parameter. If the file size exceeds 10MB, please use the `specFileUrl` parameter to transfer.
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	// Download URL for the API definition file, which must be accessible from the public network or via an internal OSS download address within the same region. The file URL must have download permissions. For non-publicly readable OSS file links, refer to the documentation at https://help.aliyun.com/zh/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects, and provide a URL with download permissions. Currently, only API definition files stored on OSS are supported.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/my-api/api.yaml
	SpecFileUrl *string `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	// The oss config info.
	SpecOssConfig *ImportHttpApiRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// When the imported API name and version management match an existing API, you need to specify an update strategy.
	//
	// - SpecOnly: Completely based on the imported file.
	//
	// - SpecFirst: Prioritize the imported file, add new interfaces and update existing ones, while keeping interfaces not mentioned in the file unchanged.
	//
	// - ExistFirst: Prioritize the existing API, only add new interfaces without updating existing ones.
	//
	// If not specified, the default strategy is ExistFirst.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// If this field is specified, this import will update the specified API instead of importing or searching for an existing API based on the API name and version management configuration. The target API must be of REST type.
	//
	// example:
	//
	// api-xxxx
	TargetHttpApiId *string `json:"targetHttpApiId,omitempty" xml:"targetHttpApiId,omitempty"`
	// API versioning configuration. If versioning is enabled, and the version number and API name match an existing API, this import will be treated as an update. If versioning is not enabled, and the API name matches an existing API, this import will also be treated as an update.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s ImportHttpApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiRequest) GoString() string {
	return s.String()
}

func (s *ImportHttpApiRequest) SetDescription(v string) *ImportHttpApiRequest {
	s.Description = &v
	return s
}

func (s *ImportHttpApiRequest) SetDryRun(v bool) *ImportHttpApiRequest {
	s.DryRun = &v
	return s
}

func (s *ImportHttpApiRequest) SetName(v string) *ImportHttpApiRequest {
	s.Name = &v
	return s
}

func (s *ImportHttpApiRequest) SetResourceGroupId(v string) *ImportHttpApiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecContentBase64(v string) *ImportHttpApiRequest {
	s.SpecContentBase64 = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecFileUrl(v string) *ImportHttpApiRequest {
	s.SpecFileUrl = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecOssConfig(v *ImportHttpApiRequestSpecOssConfig) *ImportHttpApiRequest {
	s.SpecOssConfig = v
	return s
}

func (s *ImportHttpApiRequest) SetStrategy(v string) *ImportHttpApiRequest {
	s.Strategy = &v
	return s
}

func (s *ImportHttpApiRequest) SetTargetHttpApiId(v string) *ImportHttpApiRequest {
	s.TargetHttpApiId = &v
	return s
}

func (s *ImportHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *ImportHttpApiRequest {
	s.VersionConfig = v
	return s
}

type ImportHttpApiRequestSpecOssConfig struct {
	// The name of the Object Storage Service (OSS) bucket that stores the dictionary file.
	//
	// example:
	//
	// api-1
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The file path.
	//
	// example:
	//
	// /test/swagger.json
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ImportHttpApiRequestSpecOssConfig) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiRequestSpecOssConfig) GoString() string {
	return s.String()
}

func (s *ImportHttpApiRequestSpecOssConfig) SetBucketName(v string) *ImportHttpApiRequestSpecOssConfig {
	s.BucketName = &v
	return s
}

func (s *ImportHttpApiRequestSpecOssConfig) SetObjectKey(v string) *ImportHttpApiRequestSpecOssConfig {
	s.ObjectKey = &v
	return s
}

func (s *ImportHttpApiRequestSpecOssConfig) SetRegionId(v string) *ImportHttpApiRequestSpecOssConfig {
	s.RegionId = &v
	return s
}

type ImportHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API information.
	Data *ImportHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// CE857A85-251D-5018-8103-A38957D71E20
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ImportHttpApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBody) SetCode(v string) *ImportHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *ImportHttpApiResponseBody) SetData(v *ImportHttpApiResponseBodyData) *ImportHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *ImportHttpApiResponseBody) SetMessage(v string) *ImportHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *ImportHttpApiResponseBody) SetRequestId(v string) *ImportHttpApiResponseBody {
	s.RequestId = &v
	return s
}

type ImportHttpApiResponseBodyData struct {
	// Pre-import check result.
	DryRunInfo *ImportHttpApiResponseBodyDataDryRunInfo `json:"dryRunInfo,omitempty" xml:"dryRunInfo,omitempty" type:"Struct"`
	// HTTP API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// API name.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyData) SetDryRunInfo(v *ImportHttpApiResponseBodyDataDryRunInfo) *ImportHttpApiResponseBodyData {
	s.DryRunInfo = v
	return s
}

func (s *ImportHttpApiResponseBodyData) SetHttpApiId(v string) *ImportHttpApiResponseBodyData {
	s.HttpApiId = &v
	return s
}

func (s *ImportHttpApiResponseBodyData) SetName(v string) *ImportHttpApiResponseBodyData {
	s.Name = &v
	return s
}

type ImportHttpApiResponseBodyDataDryRunInfo struct {
	// Error messages. If there are any error messages, the API cannot be imported successfully.
	ErrorMessages []*string `json:"errorMessages,omitempty" xml:"errorMessages,omitempty" type:"Repeated"`
	// 已存在的API信息。若该字段非空，则导入动作将更新该API。
	ExistHttpApiInfo *HttpApiApiInfo `json:"existHttpApiInfo,omitempty" xml:"existHttpApiInfo,omitempty"`
	// List of data structures that failed the precheck.
	FailureComponents []*ImportHttpApiResponseBodyDataDryRunInfoFailureComponents `json:"failureComponents,omitempty" xml:"failureComponents,omitempty" type:"Repeated"`
	// List of APIs that failed the precheck.
	FailureOperations []*ImportHttpApiResponseBodyDataDryRunInfoFailureOperations `json:"failureOperations,omitempty" xml:"failureOperations,omitempty" type:"Repeated"`
	// List of data structures that passed the precheck.
	SuccessComponents []*ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents `json:"successComponents,omitempty" xml:"successComponents,omitempty" type:"Repeated"`
	// List of successfully pre-checked interfaces.
	SuccessOperations []*ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations `json:"successOperations,omitempty" xml:"successOperations,omitempty" type:"Repeated"`
	// Warning messages. If there are any warning messages, some interfaces or data interfaces may not be imported successfully.
	WarningMessages []*string `json:"warningMessages,omitempty" xml:"warningMessages,omitempty" type:"Repeated"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfo) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfo) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetErrorMessages(v []*string) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.ErrorMessages = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetExistHttpApiInfo(v *HttpApiApiInfo) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.ExistHttpApiInfo = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetFailureComponents(v []*ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.FailureComponents = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetFailureOperations(v []*ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.FailureOperations = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetSuccessComponents(v []*ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.SuccessComponents = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetSuccessOperations(v []*ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.SuccessOperations = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetWarningMessages(v []*string) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.WarningMessages = v
	return s
}

type ImportHttpApiResponseBodyDataDryRunInfoFailureComponents struct {
	// Error message.
	//
	// example:
	//
	// 数据结构定义有误。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 数据结构名称。
	//
	// example:
	//
	// orderDTO
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) SetErrorMessage(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents {
	s.ErrorMessage = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents {
	s.Name = &v
	return s
}

type ImportHttpApiResponseBodyDataDryRunInfoFailureOperations struct {
	// Error message
	//
	// example:
	//
	// 缺少响应定义。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// API method.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// API path.
	//
	// example:
	//
	// /v1/orders
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetErrorMessage(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.ErrorMessage = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetMethod(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.Method = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetPath(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.Path = &v
	return s
}

type ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents struct {
	// Action to be executed after the precheck.
	//
	// - Create: Create
	//
	// - Update: Update
	//
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// Data structure name.
	//
	// example:
	//
	// userDTO
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) SetAction(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents {
	s.Action = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents {
	s.Name = &v
	return s
}

type ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations struct {
	// Action to be executed after the precheck.
	//
	// - Create: Create
	//
	// - Update: Update
	//
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// API method.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// API name.
	//
	// example:
	//
	// CreateUser
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// API path.
	//
	// example:
	//
	// /v1/users
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetAction(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Action = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetMethod(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Method = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Name = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetPath(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Path = &v
	return s
}

type ImportHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportHttpApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportHttpApiResponse) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponse) SetHeaders(v map[string]*string) *ImportHttpApiResponse {
	s.Headers = v
	return s
}

func (s *ImportHttpApiResponse) SetStatusCode(v int32) *ImportHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportHttpApiResponse) SetBody(v *ImportHttpApiResponseBody) *ImportHttpApiResponse {
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-ahr5uil8raz0rq3b
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek2sy66mftleiq
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
	// Query exactly by Gateway ID.
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
	// Query exactly by Gateway name.
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz3wes3hnre5a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Filter list by tags
	Tag []*ListGatewaysRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
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

func (s *ListGatewaysRequest) SetTag(v []*ListGatewaysRequestTag) *ListGatewaysRequest {
	s.Tag = v
	return s
}

type ListGatewaysRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewaysRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequestTag) SetKey(v string) *ListGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *ListGatewaysRequestTag) SetValue(v string) *ListGatewaysRequestTag {
	s.Value = &v
	return s
}

type ListGatewaysShrinkRequest struct {
	// Query exactly by Gateway ID.
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
	// Query exactly by Gateway name.
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz3wes3hnre5a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Filter list by tags
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
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

func (s *ListGatewaysShrinkRequest) SetTagShrink(v string) *ListGatewaysShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListGatewaysResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Result of gateway list query.
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
	// Expiration timestamp for prepaid (annual or monthly) subscriptions. Unit: milliseconds.
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
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-xxx
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
	Status         *string          `json:"status,omitempty" xml:"status,omitempty"`
	SubDomainInfos []*SubDomainInfo `json:"subDomainInfos,omitempty" xml:"subDomainInfos,omitempty" type:"Repeated"`
	// Array of tags.
	Tags []*ListGatewaysResponseBodyDataItemsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Target version of the gateway. When it is inconsistent with the current `version`, an upgrade can be performed.
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
	// The vSwtich.
	VSwitch *ListGatewaysResponseBodyDataItemsVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// Gateway version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The VPC.
	Vpc *ListGatewaysResponseBodyDataItemsVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	// The Zones.
	Zones []*ListGatewaysResponseBodyDataItemsZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
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

func (s *ListGatewaysResponseBodyDataItems) SetSubDomainInfos(v []*SubDomainInfo) *ListGatewaysResponseBodyDataItems {
	s.SubDomainInfos = v
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
	// Load balancer provision mode for the gateway:
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
	//
	// example:
	//
	// sg-xxxx
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
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewaysResponseBodyDataItemsTags) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataItemsTags) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetKey(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.Key = &v
	return s
}

func (s *ListGatewaysResponseBodyDataItemsTags) SetValue(v string) *ListGatewaysResponseBodyDataItemsTags {
	s.Value = &v
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
	// cn-hangzhou-f
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
	// Plugin ID, use this plugin ID to retrieve the plugin release information.
	//
	// example:
	//
	// pl-xxx
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
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

func (s *ListHttpApiOperationsRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApiOperationsRequest {
	s.WithPluginAttachmentByPluginId = &v
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

type ListHttpApiRoutesRequest struct {
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// example:
	//
	// NotDeployed
	DeployStatuses *string `json:"deployStatuses,omitempty" xml:"deployStatuses,omitempty"`
	DomainId       *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// item
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
	PathLike                       *string `json:"pathLike,omitempty" xml:"pathLike,omitempty"`
	WithAuthPolicyInfo             *bool   `json:"withAuthPolicyInfo,omitempty" xml:"withAuthPolicyInfo,omitempty"`
	WithConsumerInfoById           *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
}

func (s ListHttpApiRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesRequest) SetConsumerAuthorizationRuleId(v string) *ListHttpApiRoutesRequest {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetDeployStatuses(v string) *ListHttpApiRoutesRequest {
	s.DeployStatuses = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetDomainId(v string) *ListHttpApiRoutesRequest {
	s.DomainId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetEnvironmentId(v string) *ListHttpApiRoutesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetGatewayId(v string) *ListHttpApiRoutesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetName(v string) *ListHttpApiRoutesRequest {
	s.Name = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetNameLike(v string) *ListHttpApiRoutesRequest {
	s.NameLike = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPageNumber(v int32) *ListHttpApiRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPageSize(v int32) *ListHttpApiRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetPathLike(v string) *ListHttpApiRoutesRequest {
	s.PathLike = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithAuthPolicyInfo(v bool) *ListHttpApiRoutesRequest {
	s.WithAuthPolicyInfo = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithConsumerInfoById(v string) *ListHttpApiRoutesRequest {
	s.WithConsumerInfoById = &v
	return s
}

func (s *ListHttpApiRoutesRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApiRoutesRequest {
	s.WithPluginAttachmentByPluginId = &v
	return s
}

type ListHttpApiRoutesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListHttpApiRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CBEEB8C1-108E-50F0-9BEA-DED79553C309
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHttpApiRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponseBody) SetCode(v string) *ListHttpApiRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetData(v *ListHttpApiRoutesResponseBodyData) *ListHttpApiRoutesResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetMessage(v string) *ListHttpApiRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetRequestId(v string) *ListHttpApiRoutesResponseBody {
	s.RequestId = &v
	return s
}

type ListHttpApiRoutesResponseBodyData struct {
	Items []*HttpRoute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 9
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApiRoutesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponseBodyData) SetItems(v []*HttpRoute) *ListHttpApiRoutesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetPageNumber(v int32) *ListHttpApiRoutesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetPageSize(v int32) *ListHttpApiRoutesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetTotalSize(v int32) *ListHttpApiRoutesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListHttpApiRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApiRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApiRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHttpApiRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponse) SetHeaders(v map[string]*string) *ListHttpApiRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApiRoutesResponse) SetStatusCode(v int32) *ListHttpApiRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApiRoutesResponse) SetBody(v *ListHttpApiRoutesResponseBody) *ListHttpApiRoutesResponse {
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
	WithAuthPolicyList            *bool   `json:"withAuthPolicyList,omitempty" xml:"withAuthPolicyList,omitempty"`
	// Each API information in the response carries a list of authorization rules for the specified consumer ID.
	//
	// example:
	//
	// cs-xxx
	WithConsumerInfoById           *string `json:"withConsumerInfoById,omitempty" xml:"withConsumerInfoById,omitempty"`
	WithEnvironmentInfo            *bool   `json:"withEnvironmentInfo,omitempty" xml:"withEnvironmentInfo,omitempty"`
	WithEnvironmentInfoById        *string `json:"withEnvironmentInfoById,omitempty" xml:"withEnvironmentInfoById,omitempty"`
	WithIngressInfo                *bool   `json:"withIngressInfo,omitempty" xml:"withIngressInfo,omitempty"`
	WithPluginAttachmentByPluginId *string `json:"withPluginAttachmentByPluginId,omitempty" xml:"withPluginAttachmentByPluginId,omitempty"`
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

func (s *ListHttpApisRequest) SetWithAuthPolicyList(v bool) *ListHttpApisRequest {
	s.WithAuthPolicyList = &v
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

func (s *ListHttpApisRequest) SetWithEnvironmentInfoById(v string) *ListHttpApisRequest {
	s.WithEnvironmentInfoById = &v
	return s
}

func (s *ListHttpApisRequest) SetWithIngressInfo(v bool) *ListHttpApisRequest {
	s.WithIngressInfo = &v
	return s
}

func (s *ListHttpApisRequest) SetWithPluginAttachmentByPluginId(v string) *ListHttpApisRequest {
	s.WithPluginAttachmentByPluginId = &v
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

type ListPolicyClassesRequest struct {
	// example:
	//
	// Operation
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
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
	// FlowControl
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPolicyClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyClassesRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesRequest) SetAttachResourceType(v string) *ListPolicyClassesRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPolicyClassesRequest) SetDirection(v string) *ListPolicyClassesRequest {
	s.Direction = &v
	return s
}

func (s *ListPolicyClassesRequest) SetPageNumber(v int32) *ListPolicyClassesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyClassesRequest) SetPageSize(v int32) *ListPolicyClassesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyClassesRequest) SetType(v string) *ListPolicyClassesRequest {
	s.Type = &v
	return s
}

type ListPolicyClassesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListPolicyClassesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 23B45FA9-7208-5E55-B5CE-B6B2567DD822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPolicyClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponseBody) SetCode(v string) *ListPolicyClassesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPolicyClassesResponseBody) SetData(v *ListPolicyClassesResponseBodyData) *ListPolicyClassesResponseBody {
	s.Data = v
	return s
}

func (s *ListPolicyClassesResponseBody) SetMessage(v string) *ListPolicyClassesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPolicyClassesResponseBody) SetRequestId(v string) *ListPolicyClassesResponseBody {
	s.RequestId = &v
	return s
}

type ListPolicyClassesResponseBodyData struct {
	Items []*PolicyClassInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListPolicyClassesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyClassesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponseBodyData) SetItems(v []*PolicyClassInfo) *ListPolicyClassesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetPageNumber(v int32) *ListPolicyClassesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetPageSize(v int32) *ListPolicyClassesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetTotalSize(v int32) *ListPolicyClassesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListPolicyClassesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyClassesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyClassesResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponse) SetHeaders(v map[string]*string) *ListPolicyClassesResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyClassesResponse) SetStatusCode(v int32) *ListPolicyClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyClassesResponse) SetBody(v *ListPolicyClassesResponseBody) *ListPolicyClassesResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// example:
	//
	// gw-cpv4sqdl*****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// user-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetGatewayId(v string) *ListServicesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListServicesRequest) SetName(v string) *ListServicesRequest {
	s.Name = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int32) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetSourceType(v string) *ListServicesRequest {
	s.SourceType = &v
	return s
}

type ListServicesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetCode(v string) *ListServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListServicesResponseBody) SetData(v *ListServicesResponseBodyData) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetMessage(v string) *ListServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListServicesResponseBodyData struct {
	Items []*Service `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 18
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyData) SetItems(v []*Service) *ListServicesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListServicesResponseBodyData) SetPageNumber(v int32) *ListServicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBodyData) SetPageSize(v int32) *ListServicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBodyData) SetTotalSize(v int32) *ListServicesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSslCertsRequest struct {
	// example:
	//
	// ali
	CertNameLike *string `json:"certNameLike,omitempty" xml:"certNameLike,omitempty"`
	// example:
	//
	// fun.iot.evideocloud.com.cn
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSslCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSslCertsRequest) GoString() string {
	return s.String()
}

func (s *ListSslCertsRequest) SetCertNameLike(v string) *ListSslCertsRequest {
	s.CertNameLike = &v
	return s
}

func (s *ListSslCertsRequest) SetDomainName(v string) *ListSslCertsRequest {
	s.DomainName = &v
	return s
}

func (s *ListSslCertsRequest) SetPageNumber(v int32) *ListSslCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSslCertsRequest) SetPageSize(v int32) *ListSslCertsRequest {
	s.PageSize = &v
	return s
}

type ListSslCertsResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListSslCertsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// AADF7197-3384-52AF-A2DE-A66696734129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSslCertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSslCertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBody) SetCode(v string) *ListSslCertsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSslCertsResponseBody) SetData(v *ListSslCertsResponseBodyData) *ListSslCertsResponseBody {
	s.Data = v
	return s
}

func (s *ListSslCertsResponseBody) SetMessage(v string) *ListSslCertsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSslCertsResponseBody) SetRequestId(v string) *ListSslCertsResponseBody {
	s.RequestId = &v
	return s
}

type ListSslCertsResponseBodyData struct {
	Items []*SslCertMetaInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListSslCertsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSslCertsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBodyData) SetItems(v []*SslCertMetaInfo) *ListSslCertsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSslCertsResponseBodyData) SetPageNumber(v int32) *ListSslCertsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSslCertsResponseBodyData) SetPageSize(v int32) *ListSslCertsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSslCertsResponseBodyData) SetTotalSize(v int32) *ListSslCertsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListSslCertsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSslCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSslCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSslCertsResponse) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponse) SetHeaders(v map[string]*string) *ListSslCertsResponse {
	s.Headers = v
	return s
}

func (s *ListSslCertsResponse) SetStatusCode(v int32) *ListSslCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSslCertsResponse) SetBody(v *ListSslCertsResponseBody) *ListSslCertsResponse {
	s.Body = v
	return s
}

type ListZonesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListZonesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// E8079207-B651-592A-A565-23E9EE5673B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) SetCode(v string) *ListZonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListZonesResponseBody) SetData(v *ListZonesResponseBodyData) *ListZonesResponseBody {
	s.Data = v
	return s
}

func (s *ListZonesResponseBody) SetMessage(v string) *ListZonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

type ListZonesResponseBodyData struct {
	Items []*ListZonesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListZonesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyData) SetItems(v []*ListZonesResponseBodyDataItems) *ListZonesResponseBodyData {
	s.Items = v
	return s
}

type ListZonesResponseBodyDataItems struct {
	// example:
	//
	// cn-shenzhen-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListZonesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyDataItems) SetZoneId(v string) *ListZonesResponseBodyDataItems {
	s.ZoneId = &v
	return s
}

type ListZonesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponse) GoString() string {
	return s.String()
}

func (s *ListZonesResponse) SetHeaders(v map[string]*string) *ListZonesResponse {
	s.Headers = v
	return s
}

func (s *ListZonesResponse) SetStatusCode(v int32) *ListZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZonesResponse) SetBody(v *ListZonesResponseBody) *ListZonesResponse {
	s.Body = v
	return s
}

type RestartGatewayResponseBody struct {
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
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RestartGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RestartGatewayResponseBody) SetCode(v string) *RestartGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *RestartGatewayResponseBody) SetMessage(v string) *RestartGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *RestartGatewayResponseBody) SetRequestId(v string) *RestartGatewayResponseBody {
	s.RequestId = &v
	return s
}

type RestartGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartGatewayResponse) GoString() string {
	return s.String()
}

func (s *RestartGatewayResponse) SetHeaders(v map[string]*string) *RestartGatewayResponse {
	s.Headers = v
	return s
}

func (s *RestartGatewayResponse) SetStatusCode(v int32) *RestartGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartGatewayResponse) SetBody(v *RestartGatewayResponseBody) *RestartGatewayResponse {
	s.Body = v
	return s
}

type UpdateDomainRequest struct {
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	CertIdentifier   *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
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
	// TLS Cipher Suite Configuration.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
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

func (s *UpdateDomainRequest) SetCaCertIdentifier(v string) *UpdateDomainRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *UpdateDomainRequest) SetCertIdentifier(v string) *UpdateDomainRequest {
	s.CertIdentifier = &v
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

func (s *UpdateDomainRequest) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *UpdateDomainRequest {
	s.TlsCipherSuitesConfig = v
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

type UpdateGatewayFeatureRequest struct {
	// Parameter value.
	//
	// example:
	//
	// "true"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateGatewayFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureRequest) SetValue(v string) *UpdateGatewayFeatureRequest {
	s.Value = &v
	return s
}

type UpdateGatewayFeatureResponseBody struct {
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

func (s UpdateGatewayFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureResponseBody) SetCode(v string) *UpdateGatewayFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayFeatureResponseBody) SetMessage(v string) *UpdateGatewayFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayFeatureResponseBody) SetRequestId(v string) *UpdateGatewayFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayFeatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureResponse) SetHeaders(v map[string]*string) *UpdateGatewayFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayFeatureResponse) SetStatusCode(v int32) *UpdateGatewayFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayFeatureResponse) SetBody(v *UpdateGatewayFeatureResponseBody) *UpdateGatewayFeatureResponse {
	s.Body = v
	return s
}

type UpdateGatewayNameRequest struct {
	// Gateway name.
	//
	// example:
	//
	// dev-itemcenter-router
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateGatewayNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameRequest) SetName(v string) *UpdateGatewayNameRequest {
	s.Name = &v
	return s
}

type UpdateGatewayNameResponseBody struct {
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
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponseBody) SetCode(v string) *UpdateGatewayNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetMessage(v string) *UpdateGatewayNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetRequestId(v string) *UpdateGatewayNameResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponse) SetHeaders(v map[string]*string) *UpdateGatewayNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayNameResponse) SetStatusCode(v int32) *UpdateGatewayNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayNameResponse) SetBody(v *UpdateGatewayNameResponseBody) *UpdateGatewayNameResponse {
	s.Body = v
	return s
}

type UpdateHttpApiRequest struct {
	// List of AI protocols.
	AiProtocols []*string   `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	AuthConfig  *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// Base path of the API, which must start with a \\"/\\".
	//
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// List of API deployment configurations.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// API description.
	//
	// example:
	//
	// 更新API描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableAuth  *bool   `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
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

func (s *UpdateHttpApiRequest) SetAuthConfig(v *AuthConfig) *UpdateHttpApiRequest {
	s.AuthConfig = v
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

func (s *UpdateHttpApiRequest) SetEnableAuth(v bool) *UpdateHttpApiRequest {
	s.EnableAuth = &v
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

type UpdateHttpApiRouteRequest struct {
	// Backend service configuration for the route.
	BackendConfig *UpdateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Route description.
	//
	// example:
	//
	// test route
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// List of domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Environment ID.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
}

func (s UpdateHttpApiRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequest) SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDescription(v string) *UpdateHttpApiRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDomainIds(v []*string) *UpdateHttpApiRouteRequest {
	s.DomainIds = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetEnvironmentId(v string) *UpdateHttpApiRouteRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest {
	s.Match = v
	return s
}

type UpdateHttpApiRouteRequestBackendConfig struct {
	// Backend service scenario.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// List of backend services.
	Services []*UpdateHttpApiRouteRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequestBackendConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetScene(v string) *UpdateHttpApiRouteRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetServices(v []*UpdateHttpApiRouteRequestBackendConfigServices) *UpdateHttpApiRouteRequestBackendConfig {
	s.Services = v
	return s
}

type UpdateHttpApiRouteRequestBackendConfigServices struct {
	// Service port, do not pass in for dynamic ports.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol:
	//
	// - HTTP.
	//
	// - HTTPS.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Service ID.
	//
	// example:
	//
	// svc-cr6pk4tlhtgm58e***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Percentage value of traffic ratio.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s UpdateHttpApiRouteRequestBackendConfigServices) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetPort(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetProtocol(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetServiceId(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetVersion(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetWeight(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Weight = &v
	return s
}

type UpdateHttpApiRouteResponseBody struct {
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
	// CBEEB8C1-108E-50F0-9BEA-DED79553C309
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateHttpApiRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteResponseBody) SetCode(v string) *UpdateHttpApiRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpApiRouteResponseBody) SetMessage(v string) *UpdateHttpApiRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpApiRouteResponseBody) SetRequestId(v string) *UpdateHttpApiRouteResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteResponse) SetHeaders(v map[string]*string) *UpdateHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiRouteResponse) SetStatusCode(v int32) *UpdateHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiRouteResponse) SetBody(v *UpdateHttpApiRouteResponseBody) *UpdateHttpApiRouteResponse {
	s.Body = v
	return s
}

type UpdatePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// celue-timeout-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) SetConfig(v string) *UpdatePolicyRequest {
	s.Config = &v
	return s
}

func (s *UpdatePolicyRequest) SetDescription(v string) *UpdatePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyRequest) SetName(v string) *UpdatePolicyRequest {
	s.Name = &v
	return s
}

type UpdatePolicyResponseBody struct {
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
	// C67DED2B-F19B-5BEC-88C1-D6EB854C***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) SetCode(v string) *UpdatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetMessage(v string) *UpdatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetRequestId(v string) *UpdatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponse) SetHeaders(v map[string]*string) *UpdatePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyResponse) SetStatusCode(v int32) *UpdatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyResponse) SetBody(v *UpdatePolicyResponseBody) *UpdatePolicyResponse {
	s.Body = v
	return s
}

type UpgradeGatewayRequest struct {
	// Gateway version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayRequest) SetVersion(v string) *UpgradeGatewayRequest {
	s.Version = &v
	return s
}

type UpgradeGatewayResponseBody struct {
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

func (s UpgradeGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponseBody) SetCode(v string) *UpgradeGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetMessage(v string) *UpgradeGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetRequestId(v string) *UpgradeGatewayResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponse) SetHeaders(v map[string]*string) *UpgradeGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeGatewayResponse) SetStatusCode(v int32) *UpgradeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeGatewayResponse) SetBody(v *UpgradeGatewayResponseBody) *UpgradeGatewayResponse {
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddGatewaySecurityGroupRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddGatewaySecurityGroupRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/move-resource-group"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.TlsCipherSuitesConfig)) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDomainResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDomainResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Deprecated: OpenAPI CreateEnvironment is deprecated
//
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
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI CreateEnvironment is deprecated
//
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
// Deprecated
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

	if !tea.BoolValue(util.IsUnset(request.AuthConfig)) {
		body["authConfig"] = request.AuthConfig
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

	if !tea.BoolValue(util.IsUnset(request.EnableAuth)) {
		body["enableAuth"] = request.EnableAuth
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateHttpApiOperationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateHttpApiOperationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Create a route for HttpApi
//
// @param request - CreateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiRouteResponse
func (client *Client) CreateHttpApiRouteWithOptions(httpApiId *string, request *CreateHttpApiRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHttpApiRouteResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DomainIds)) {
		body["domainIds"] = request.DomainIds
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
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
		Action:      tea.String("CreateHttpApiRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/routes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateHttpApiRouteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateHttpApiRouteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Create a route for HttpApi
//
// @param request - CreateHttpApiRouteRequest
//
// @return CreateHttpApiRouteResponse
func (client *Client) CreateHttpApiRoute(httpApiId *string, request *CreateHttpApiRouteRequest) (_result *CreateHttpApiRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiRouteResponse{}
	_body, _err := client.CreateHttpApiRouteWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param request - CreatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassName)) {
		body["className"] = request.ClassName
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicy"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/policies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建策略
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建策略资源挂载
//
// @param request - CreatePolicyAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyAttachmentResponse
func (client *Client) CreatePolicyAttachmentWithOptions(request *CreatePolicyAttachmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePolicyAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachResourceId)) {
		body["attachResourceId"] = request.AttachResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.AttachResourceType)) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		body["policyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicyAttachment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/policy-attachments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePolicyAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePolicyAttachmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建策略资源挂载
//
// @param request - CreatePolicyAttachmentRequest
//
// @return CreatePolicyAttachmentResponse
func (client *Client) CreatePolicyAttachment(request *CreatePolicyAttachmentRequest) (_result *CreatePolicyAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePolicyAttachmentResponse{}
	_body, _err := client.CreatePolicyAttachmentWithOptions(request, headers, runtime)
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
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConfigs)) {
		body["serviceConfigs"] = request.ServiceConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建服务
//
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDomainResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDomainResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Deprecated: OpenAPI DeleteEnvironment is deprecated
//
// Summary:
//
// DeleteEnvironment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI DeleteEnvironment is deprecated
//
// Summary:
//
// DeleteEnvironment
//
// @return DeleteEnvironmentResponse
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteGatewayResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteGatewayResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Delete the security group rule of a gateway
//
// @param request - DeleteGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewaySecurityGroupRuleResponse
func (client *Client) DeleteGatewaySecurityGroupRuleWithOptions(gatewayId *string, securityGroupRuleId *string, request *DeleteGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewaySecurityGroupRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CascadingDelete)) {
		query["cascadingDelete"] = request.CascadingDelete
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewaySecurityGroupRule"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/security-group-rules/" + tea.StringValue(openapiutil.GetEncodeParam(securityGroupRuleId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteGatewaySecurityGroupRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteGatewaySecurityGroupRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Delete the security group rule of a gateway
//
// @param request - DeleteGatewaySecurityGroupRuleRequest
//
// @return DeleteGatewaySecurityGroupRuleResponse
func (client *Client) DeleteGatewaySecurityGroupRule(gatewayId *string, securityGroupRuleId *string, request *DeleteGatewaySecurityGroupRuleRequest) (_result *DeleteGatewaySecurityGroupRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewaySecurityGroupRuleResponse{}
	_body, _err := client.DeleteGatewaySecurityGroupRuleWithOptions(gatewayId, securityGroupRuleId, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteHttpApiOperationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteHttpApiOperationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Delete the route of an HttpApi
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiRouteResponse
func (client *Client) DeleteHttpApiRouteWithOptions(httpApiId *string, routeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHttpApiRouteResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHttpApiRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/routes/" + tea.StringValue(openapiutil.GetEncodeParam(routeId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteHttpApiRouteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteHttpApiRouteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Delete the route of an HttpApi
//
// @return DeleteHttpApiRouteResponse
func (client *Client) DeleteHttpApiRoute(httpApiId *string, routeId *string) (_result *DeleteHttpApiRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiRouteResponse{}
	_body, _err := client.DeleteHttpApiRouteWithOptions(httpApiId, routeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除策略
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(policyId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicy"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/policies/" + tea.StringValue(openapiutil.GetEncodeParam(policyId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeletePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeletePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除策略
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(policyId *string) (_result *DeletePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除策略资源挂载
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyAttachmentResponse
func (client *Client) DeletePolicyAttachmentWithOptions(policyAttachmentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePolicyAttachmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicyAttachment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/policy-attachments/" + tea.StringValue(openapiutil.GetEncodeParam(policyAttachmentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeletePolicyAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeletePolicyAttachmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除策略资源挂载
//
// @return DeletePolicyAttachmentResponse
func (client *Client) DeletePolicyAttachment(policyAttachmentId *string) (_result *DeletePolicyAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyAttachmentResponse{}
	_body, _err := client.DeletePolicyAttachmentWithOptions(policyAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 部署HttpApi
//
// @param request - DeployHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployHttpApiResponse
func (client *Client) DeployHttpApiWithOptions(httpApiId *string, request *DeployHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RestApiConfig)) {
		body["restApiConfig"] = request.RestApiConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RouteId)) {
		body["routeId"] = request.RouteId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/deploy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeployHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeployHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 部署HttpApi
//
// @param request - DeployHttpApiRequest
//
// @return DeployHttpApiResponse
func (client *Client) DeployHttpApi(httpApiId *string, request *DeployHttpApiRequest) (_result *DeployHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployHttpApiResponse{}
	_body, _err := client.DeployHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Export HTTP API
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportHttpApiResponse
func (client *Client) ExportHttpApiWithOptions(httpApiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExportHttpApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ExportHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/export"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ExportHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ExportHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Export HTTP API
//
// @return ExportHttpApiResponse
func (client *Client) ExportHttpApi(httpApiId *string) (_result *ExportHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportHttpApiResponse{}
	_body, _err := client.ExportHttpApiWithOptions(httpApiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取监控/日志大盘接口
//
// @param tmpReq - GetDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDashboardResponse
func (client *Client) GetDashboardWithOptions(gatewayId *string, tmpReq *GetDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDashboardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDashboardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["apiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PluginClassId)) {
		query["pluginClassId"] = request.PluginClassId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDashboard"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/dashboards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDashboardResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDashboardResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取监控/日志大盘接口
//
// @param request - GetDashboardRequest
//
// @return GetDashboardResponse
func (client *Client) GetDashboard(gatewayId *string, request *GetDashboardRequest) (_result *GetDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDashboardResponse{}
	_body, _err := client.GetDashboardWithOptions(gatewayId, request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDomainResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDomainResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Deprecated: OpenAPI GetEnvironment is deprecated
//
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
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI GetEnvironment is deprecated
//
// Summary:
//
// GetEnvironment
//
// @param request - GetEnvironmentRequest
//
// @return GetEnvironmentResponse
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetGatewayResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetGatewayResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHttpApiOperationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHttpApiOperationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Get the details of an HttpApi route
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHttpApiRouteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHttpApiRouteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get the details of an HttpApi route
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
// GetPolicy
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(policyId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicy"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/policies/" + tea.StringValue(openapiutil.GetEncodeParam(policyId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// GetPolicy
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(policyId *string) (_result *GetPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询策略资源挂载
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyAttachmentResponse
func (client *Client) GetPolicyAttachmentWithOptions(policyAttachmentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPolicyAttachmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicyAttachment"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/policy-attachments/" + tea.StringValue(openapiutil.GetEncodeParam(policyAttachmentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPolicyAttachmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPolicyAttachmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询策略资源挂载
//
// @return GetPolicyAttachmentResponse
func (client *Client) GetPolicyAttachment(policyAttachmentId *string) (_result *GetPolicyAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyAttachmentResponse{}
	_body, _err := client.GetPolicyAttachmentWithOptions(policyAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源概览信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceOverviewResponse
func (client *Client) GetResourceOverviewWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceOverviewResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceOverview"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/overview/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceOverviewResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceOverviewResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取资源概览信息
//
// @return GetResourceOverviewResponse
func (client *Client) GetResourceOverview() (_result *GetResourceOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceOverviewResponse{}
	_body, _err := client.GetResourceOverviewWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(serviceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/services/" + tea.StringValue(openapiutil.GetEncodeParam(serviceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取服务详情
//
// @return GetServiceResponse
func (client *Client) GetService(serviceId *string) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取链路追踪配置
//
// @param request - GetTraceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceConfigResponse
func (client *Client) GetTraceConfigWithOptions(gatewayId *string, request *GetTraceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTraceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTraceConfig"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/trace"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTraceConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTraceConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取链路追踪配置
//
// @param request - GetTraceConfigRequest
//
// @return GetTraceConfigResponse
func (client *Client) GetTraceConfig(gatewayId *string, request *GetTraceConfigRequest) (_result *GetTraceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTraceConfigResponse{}
	_body, _err := client.GetTraceConfigWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Import HttpApi
//
// @param request - ImportHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportHttpApiResponse
func (client *Client) ImportHttpApiWithOptions(request *ImportHttpApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImportHttpApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["dryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecContentBase64)) {
		body["specContentBase64"] = request.SpecContentBase64
	}

	if !tea.BoolValue(util.IsUnset(request.SpecFileUrl)) {
		body["specFileUrl"] = request.SpecFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SpecOssConfig)) {
		body["specOssConfig"] = request.SpecOssConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Strategy)) {
		body["strategy"] = request.Strategy
	}

	if !tea.BoolValue(util.IsUnset(request.TargetHttpApiId)) {
		body["targetHttpApiId"] = request.TargetHttpApiId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionConfig)) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportHttpApi"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ImportHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ImportHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Import HttpApi
//
// @param request - ImportHttpApiRequest
//
// @return ImportHttpApiResponse
func (client *Client) ImportHttpApi(request *ImportHttpApiRequest) (_result *ImportHttpApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportHttpApiResponse{}
	_body, _err := client.ImportHttpApiWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDomainsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDomainsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Deprecated: OpenAPI ListEnvironments is deprecated
//
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
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnvironmentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnvironmentsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI ListEnvironments is deprecated
//
// Summary:
//
// ListEnvironments
//
// @param request - ListEnvironmentsRequest
//
// @return ListEnvironmentsResponse
// Deprecated
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
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListGatewaysResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListGatewaysResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

	if !tea.BoolValue(util.IsUnset(request.WithPluginAttachmentByPluginId)) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListHttpApiOperationsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListHttpApiOperationsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 创建HttpApi的路由
//
// @param request - ListHttpApiRoutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiRoutesResponse
func (client *Client) ListHttpApiRoutesWithOptions(httpApiId *string, request *ListHttpApiRoutesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHttpApiRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerAuthorizationRuleId)) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployStatuses)) {
		query["deployStatuses"] = request.DeployStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["domainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		query["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["gatewayId"] = request.GatewayId
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

	if !tea.BoolValue(util.IsUnset(request.WithAuthPolicyInfo)) {
		query["withAuthPolicyInfo"] = request.WithAuthPolicyInfo
	}

	if !tea.BoolValue(util.IsUnset(request.WithConsumerInfoById)) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !tea.BoolValue(util.IsUnset(request.WithPluginAttachmentByPluginId)) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHttpApiRoutes"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/routes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListHttpApiRoutesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListHttpApiRoutesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建HttpApi的路由
//
// @param request - ListHttpApiRoutesRequest
//
// @return ListHttpApiRoutesResponse
func (client *Client) ListHttpApiRoutes(httpApiId *string, request *ListHttpApiRoutesRequest) (_result *ListHttpApiRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApiRoutesResponse{}
	_body, _err := client.ListHttpApiRoutesWithOptions(httpApiId, request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.WithAuthPolicyList)) {
		query["withAuthPolicyList"] = request.WithAuthPolicyList
	}

	if !tea.BoolValue(util.IsUnset(request.WithConsumerInfoById)) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !tea.BoolValue(util.IsUnset(request.WithEnvironmentInfo)) {
		query["withEnvironmentInfo"] = request.WithEnvironmentInfo
	}

	if !tea.BoolValue(util.IsUnset(request.WithEnvironmentInfoById)) {
		query["withEnvironmentInfoById"] = request.WithEnvironmentInfoById
	}

	if !tea.BoolValue(util.IsUnset(request.WithIngressInfo)) {
		query["withIngressInfo"] = request.WithIngressInfo
	}

	if !tea.BoolValue(util.IsUnset(request.WithPluginAttachmentByPluginId)) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListHttpApisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListHttpApisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// ListPolicyClasses。
//
// @param request - ListPolicyClassesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyClassesResponse
func (client *Client) ListPolicyClassesWithOptions(request *ListPolicyClassesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPolicyClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachResourceType)) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicyClasses"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/policy-classes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPolicyClassesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPolicyClassesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// ListPolicyClasses。
//
// @param request - ListPolicyClassesRequest
//
// @return ListPolicyClassesResponse
func (client *Client) ListPolicyClasses(request *ListPolicyClassesRequest) (_result *ListPolicyClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPolicyClassesResponse{}
	_body, _err := client.ListPolicyClassesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务列表
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["gatewayId"] = request.GatewayId
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

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListServicesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListServicesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取服务列表
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListSslCerts。
//
// @param request - ListSslCertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSslCertsResponse
func (client *Client) ListSslCertsWithOptions(request *ListSslCertsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSslCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertNameLike)) {
		query["certNameLike"] = request.CertNameLike
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["domainName"] = request.DomainName
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
		Action:      tea.String("ListSslCerts"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/ssl/certs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSslCertsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSslCertsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// ListSslCerts。
//
// @param request - ListSslCertsRequest
//
// @return ListSslCertsResponse
func (client *Client) ListSslCerts(request *ListSslCertsRequest) (_result *ListSslCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSslCertsResponse{}
	_body, _err := client.ListSslCertsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云原生API网关region下的可用区
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZonesResponse
func (client *Client) ListZonesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListZones"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/zones"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListZonesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListZonesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取云原生API网关region下的可用区
//
// @return ListZonesResponse
func (client *Client) ListZones() (_result *ListZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListZonesResponse{}
	_body, _err := client.ListZonesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gateway Restart
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartGatewayResponse
func (client *Client) RestartGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RestartGateway"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartGatewayResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartGatewayResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Gateway Restart
//
// @return RestartGatewayResponse
func (client *Client) RestartGateway(gatewayId *string) (_result *RestartGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartGatewayResponse{}
	_body, _err := client.RestartGatewayWithOptions(gatewayId, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.TlsCipherSuitesConfig)) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDomainResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDomainResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Deprecated: OpenAPI UpdateEnvironment is deprecated
//
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
// Deprecated
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateEnvironmentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateEnvironmentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI UpdateEnvironment is deprecated
//
// Summary:
//
// UpdateEnvironment
//
// @param request - UpdateEnvironmentRequest
//
// @return UpdateEnvironmentResponse
// Deprecated
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
// Get the feature configuration of the gateway
//
// @param request - UpdateGatewayFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayFeatureResponse
func (client *Client) UpdateGatewayFeatureWithOptions(gatewayId *string, name *string, request *UpdateGatewayFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayFeature"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/gateway-features/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateGatewayFeatureResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateGatewayFeatureResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get the feature configuration of the gateway
//
// @param request - UpdateGatewayFeatureRequest
//
// @return UpdateGatewayFeatureResponse
func (client *Client) UpdateGatewayFeature(gatewayId *string, name *string, request *UpdateGatewayFeatureRequest) (_result *UpdateGatewayFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayFeatureResponse{}
	_body, _err := client.UpdateGatewayFeatureWithOptions(gatewayId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Change the name of a gateway instance
//
// @param request - UpdateGatewayNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayNameWithOptions(gatewayId *string, request *UpdateGatewayNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayName"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/name"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateGatewayNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateGatewayNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Change the name of a gateway instance
//
// @param request - UpdateGatewayNameRequest
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayName(gatewayId *string, request *UpdateGatewayNameRequest) (_result *UpdateGatewayNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.UpdateGatewayNameWithOptions(gatewayId, request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.AuthConfig)) {
		body["authConfig"] = request.AuthConfig
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

	if !tea.BoolValue(util.IsUnset(request.EnableAuth)) {
		body["enableAuth"] = request.EnableAuth
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateHttpApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateHttpApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateHttpApiOperationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateHttpApiOperationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Summary:
//
// Update the route of HttpApi
//
// @param request - UpdateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiRouteResponse
func (client *Client) UpdateHttpApiRouteWithOptions(httpApiId *string, routeId *string, request *UpdateHttpApiRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHttpApiRouteResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DomainIds)) {
		body["domainIds"] = request.DomainIds
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentId)) {
		body["environmentId"] = request.EnvironmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Match)) {
		body["match"] = request.Match
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHttpApiRoute"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/http-apis/" + tea.StringValue(openapiutil.GetEncodeParam(httpApiId)) + "/routes/" + tea.StringValue(openapiutil.GetEncodeParam(routeId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateHttpApiRouteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateHttpApiRouteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Update the route of HttpApi
//
// @param request - UpdateHttpApiRouteRequest
//
// @return UpdateHttpApiRouteResponse
func (client *Client) UpdateHttpApiRoute(httpApiId *string, routeId *string, request *UpdateHttpApiRouteRequest) (_result *UpdateHttpApiRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiRouteResponse{}
	_body, _err := client.UpdateHttpApiRouteWithOptions(httpApiId, routeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新策略
//
// @param request - UpdatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithOptions(policyId *string, request *UpdatePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePolicy"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/policies/" + tea.StringValue(openapiutil.GetEncodeParam(policyId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdatePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdatePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新策略
//
// @param request - UpdatePolicyRequest
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicy(policyId *string, request *UpdatePolicyRequest) (_result *UpdatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrade the gateway version
//
// @param request - UpgradeGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeGatewayResponse
func (client *Client) UpgradeGatewayWithOptions(gatewayId *string, request *UpgradeGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeGateway"),
		Version:     tea.String("2024-03-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(gatewayId)) + "/upgrade"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeGatewayResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeGatewayResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrade the gateway version
//
// @param request - UpgradeGatewayRequest
//
// @return UpgradeGatewayResponse
func (client *Client) UpgradeGateway(gatewayId *string, request *UpgradeGatewayRequest) (_result *UpgradeGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.UpgradeGatewayWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

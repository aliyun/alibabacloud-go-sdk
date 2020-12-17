// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FindAllServiceRequest struct {
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
}

func (s FindAllServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceRequest) GoString() string {
	return s.String()
}

func (s *FindAllServiceRequest) SetPageNumber(v int64) *FindAllServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *FindAllServiceRequest) SetPageSize(v string) *FindAllServiceRequest {
	s.PageSize = &v
	return s
}

func (s *FindAllServiceRequest) SetName(v string) *FindAllServiceRequest {
	s.Name = &v
	return s
}

func (s *FindAllServiceRequest) SetAliasName(v string) *FindAllServiceRequest {
	s.AliasName = &v
	return s
}

func (s *FindAllServiceRequest) SetSourceType(v int64) *FindAllServiceRequest {
	s.SourceType = &v
	return s
}

func (s *FindAllServiceRequest) SetIsHealth(v bool) *FindAllServiceRequest {
	s.IsHealth = &v
	return s
}

type FindAllServiceResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *FindAllServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindAllServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FindAllServiceResponseBody) SetCode(v int64) *FindAllServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FindAllServiceResponseBody) SetData(v *FindAllServiceResponseBodyData) *FindAllServiceResponseBody {
	s.Data = v
	return s
}

func (s *FindAllServiceResponseBody) SetMessage(v string) *FindAllServiceResponseBody {
	s.Message = &v
	return s
}

type FindAllServiceResponseBodyData struct {
	// list
	List []*FindAllServiceResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s FindAllServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindAllServiceResponseBodyData) SetList(v []*FindAllServiceResponseBodyDataList) *FindAllServiceResponseBodyData {
	s.List = v
	return s
}

func (s *FindAllServiceResponseBodyData) SetTotalCount(v int64) *FindAllServiceResponseBodyData {
	s.TotalCount = &v
	return s
}

type FindAllServiceResponseBodyDataList struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*FindAllServiceResponseBodyDataListServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindAllServiceResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *FindAllServiceResponseBodyDataList) SetAliasName(v string) *FindAllServiceResponseBodyDataList {
	s.AliasName = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetCreationDateTime(v string) *FindAllServiceResponseBodyDataList {
	s.CreationDateTime = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetDescription(v string) *FindAllServiceResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetId(v int64) *FindAllServiceResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetIsAutoRefresh(v bool) *FindAllServiceResponseBodyDataList {
	s.IsAutoRefresh = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetIsHealth(v bool) *FindAllServiceResponseBodyDataList {
	s.IsHealth = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetName(v string) *FindAllServiceResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetRegistryId(v string) *FindAllServiceResponseBodyDataList {
	s.RegistryId = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetServiceEnds(v []*FindAllServiceResponseBodyDataListServiceEnds) *FindAllServiceResponseBodyDataList {
	s.ServiceEnds = v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetServiceNameInRegistry(v string) *FindAllServiceResponseBodyDataList {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetSourceType(v int64) *FindAllServiceResponseBodyDataList {
	s.SourceType = &v
	return s
}

func (s *FindAllServiceResponseBodyDataList) SetUpdateDateTime(v string) *FindAllServiceResponseBodyDataList {
	s.UpdateDateTime = &v
	return s
}

type FindAllServiceResponseBodyDataListServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindAllServiceResponseBodyDataListServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceResponseBodyDataListServiceEnds) GoString() string {
	return s.String()
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetCreationDateTime(v string) *FindAllServiceResponseBodyDataListServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetId(v int64) *FindAllServiceResponseBodyDataListServiceEnds {
	s.Id = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetIpAddress(v string) *FindAllServiceResponseBodyDataListServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetPort(v string) *FindAllServiceResponseBodyDataListServiceEnds {
	s.Port = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetServiceId(v int64) *FindAllServiceResponseBodyDataListServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetStatus(v int64) *FindAllServiceResponseBodyDataListServiceEnds {
	s.Status = &v
	return s
}

func (s *FindAllServiceResponseBodyDataListServiceEnds) SetUpdateDateTime(v string) *FindAllServiceResponseBodyDataListServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type FindAllServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindAllServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindAllServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s FindAllServiceResponse) GoString() string {
	return s.String()
}

func (s *FindAllServiceResponse) SetHeaders(v map[string]*string) *FindAllServiceResponse {
	s.Headers = v
	return s
}

func (s *FindAllServiceResponse) SetBody(v *FindAllServiceResponseBody) *FindAllServiceResponse {
	s.Body = v
	return s
}

type CreateApiRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedServices
	AttachedServices []*int64 `json:"attachedServices,omitempty" xml:"attachedServices,omitempty" type:"Repeated"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiRequest) GoString() string {
	return s.String()
}

func (s *CreateApiRequest) SetAliasName(v string) *CreateApiRequest {
	s.AliasName = &v
	return s
}

func (s *CreateApiRequest) SetAttachedServices(v []*int64) *CreateApiRequest {
	s.AttachedServices = v
	return s
}

func (s *CreateApiRequest) SetBasePath(v string) *CreateApiRequest {
	s.BasePath = &v
	return s
}

func (s *CreateApiRequest) SetDescription(v string) *CreateApiRequest {
	s.Description = &v
	return s
}

func (s *CreateApiRequest) SetName(v string) *CreateApiRequest {
	s.Name = &v
	return s
}

func (s *CreateApiRequest) SetStatus(v int64) *CreateApiRequest {
	s.Status = &v
	return s
}

type CreateApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiResponseBody) SetCode(v int64) *CreateApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiResponseBody) SetData(v map[string]interface{}) *CreateApiResponseBody {
	s.Data = v
	return s
}

func (s *CreateApiResponseBody) SetMessage(v string) *CreateApiResponseBody {
	s.Message = &v
	return s
}

type CreateApiResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiResponse) GoString() string {
	return s.String()
}

func (s *CreateApiResponse) SetHeaders(v map[string]*string) *CreateApiResponse {
	s.Headers = v
	return s
}

func (s *CreateApiResponse) SetBody(v *CreateApiResponseBody) *CreateApiResponse {
	s.Body = v
	return s
}

type GetGatewayByIdResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetGatewayByIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetGatewayByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayByIdResponseBody) SetCode(v int64) *GetGatewayByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayByIdResponseBody) SetData(v []*GetGatewayByIdResponseBodyData) *GetGatewayByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayByIdResponseBody) SetMessage(v string) *GetGatewayByIdResponseBody {
	s.Message = &v
	return s
}

type GetGatewayByIdResponseBodyData struct {
	// armsInfo
	ArmsInfo *GetGatewayByIdResponseBodyDataArmsInfo `json:"armsInfo,omitempty" xml:"armsInfo,omitempty" type:"Struct"`
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbAccessAddr
	SlbAccessAddr *string `json:"slbAccessAddr,omitempty" xml:"slbAccessAddr,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
}

func (s GetGatewayByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayByIdResponseBodyData) SetArmsInfo(v *GetGatewayByIdResponseBodyDataArmsInfo) *GetGatewayByIdResponseBodyData {
	s.ArmsInfo = v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetAutoCreateSlb(v bool) *GetGatewayByIdResponseBodyData {
	s.AutoCreateSlb = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetBasePath(v string) *GetGatewayByIdResponseBodyData {
	s.BasePath = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetCreationDateTime(v string) *GetGatewayByIdResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetEdasNamespaceId(v string) *GetGatewayByIdResponseBodyData {
	s.EdasNamespaceId = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetGatewayType(v string) *GetGatewayByIdResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetId(v int64) *GetGatewayByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetName(v string) *GetGatewayByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetPodCidr(v string) *GetGatewayByIdResponseBodyData {
	s.PodCidr = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetRegion(v string) *GetGatewayByIdResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetRegionName(v string) *GetGatewayByIdResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetReplica(v int64) *GetGatewayByIdResponseBodyData {
	s.Replica = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetRuntimeOn(v string) *GetGatewayByIdResponseBodyData {
	s.RuntimeOn = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetSecurityGroup(v string) *GetGatewayByIdResponseBodyData {
	s.SecurityGroup = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetSlb(v string) *GetGatewayByIdResponseBodyData {
	s.Slb = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetSlbAccessAddr(v string) *GetGatewayByIdResponseBodyData {
	s.SlbAccessAddr = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetStatus(v string) *GetGatewayByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetVpc(v string) *GetGatewayByIdResponseBodyData {
	s.Vpc = &v
	return s
}

func (s *GetGatewayByIdResponseBodyData) SetVswitch(v string) *GetGatewayByIdResponseBodyData {
	s.Vswitch = &v
	return s
}

type GetGatewayByIdResponseBodyDataArmsInfo struct {
	// appId
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// licenseKey
	LicenseKey *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
}

func (s GetGatewayByIdResponseBodyDataArmsInfo) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayByIdResponseBodyDataArmsInfo) GoString() string {
	return s.String()
}

func (s *GetGatewayByIdResponseBodyDataArmsInfo) SetAppId(v string) *GetGatewayByIdResponseBodyDataArmsInfo {
	s.AppId = &v
	return s
}

func (s *GetGatewayByIdResponseBodyDataArmsInfo) SetAppName(v string) *GetGatewayByIdResponseBodyDataArmsInfo {
	s.AppName = &v
	return s
}

func (s *GetGatewayByIdResponseBodyDataArmsInfo) SetDescription(v string) *GetGatewayByIdResponseBodyDataArmsInfo {
	s.Description = &v
	return s
}

func (s *GetGatewayByIdResponseBodyDataArmsInfo) SetLicenseKey(v string) *GetGatewayByIdResponseBodyDataArmsInfo {
	s.LicenseKey = &v
	return s
}

type GetGatewayByIdResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGatewayByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayByIdResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayByIdResponse) SetHeaders(v map[string]*string) *GetGatewayByIdResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayByIdResponse) SetBody(v *GetGatewayByIdResponseBody) *GetGatewayByIdResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// gatewayId
	GatewayId *int64 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetAliasName(v string) *CreatePolicyRequest {
	s.AliasName = &v
	return s
}

func (s *CreatePolicyRequest) SetContent(v string) *CreatePolicyRequest {
	s.Content = &v
	return s
}

func (s *CreatePolicyRequest) SetGatewayId(v int64) *CreatePolicyRequest {
	s.GatewayId = &v
	return s
}

func (s *CreatePolicyRequest) SetId(v int64) *CreatePolicyRequest {
	s.Id = &v
	return s
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyGroup(v string) *CreatePolicyRequest {
	s.PolicyGroup = &v
	return s
}

func (s *CreatePolicyRequest) SetType(v int64) *CreatePolicyRequest {
	s.Type = &v
	return s
}

type CreatePolicyResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetCode(v int64) *CreatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyResponseBody) SetData(v map[string]interface{}) *CreatePolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyResponseBody) SetMessage(v string) *CreatePolicyResponseBody {
	s.Message = &v
	return s
}

type CreatePolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type GetServiceInstanceForRegistryByServiceNameRequest struct {
	// serviceName
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetServiceInstanceForRegistryByServiceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceForRegistryByServiceNameRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceForRegistryByServiceNameRequest) SetServiceName(v string) *GetServiceInstanceForRegistryByServiceNameRequest {
	s.ServiceName = &v
	return s
}

type GetServiceInstanceForRegistryByServiceNameResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetServiceInstanceForRegistryByServiceNameResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetServiceInstanceForRegistryByServiceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceForRegistryByServiceNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBody) SetCode(v int64) *GetServiceInstanceForRegistryByServiceNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBody) SetData(v []*GetServiceInstanceForRegistryByServiceNameResponseBodyData) *GetServiceInstanceForRegistryByServiceNameResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBody) SetMessage(v string) *GetServiceInstanceForRegistryByServiceNameResponseBody {
	s.Message = &v
	return s
}

type GetServiceInstanceForRegistryByServiceNameResponseBodyData struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// metaInfo
	MetaInfo    *string   `json:"metaInfo,omitempty" xml:"metaInfo,omitempty"`
	ServiceEnds []*string `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceName
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s GetServiceInstanceForRegistryByServiceNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceForRegistryByServiceNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBodyData) SetId(v int64) *GetServiceInstanceForRegistryByServiceNameResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBodyData) SetMetaInfo(v string) *GetServiceInstanceForRegistryByServiceNameResponseBodyData {
	s.MetaInfo = &v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBodyData) SetServiceEnds(v []*string) *GetServiceInstanceForRegistryByServiceNameResponseBodyData {
	s.ServiceEnds = v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponseBodyData) SetServiceName(v string) *GetServiceInstanceForRegistryByServiceNameResponseBodyData {
	s.ServiceName = &v
	return s
}

type GetServiceInstanceForRegistryByServiceNameResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceInstanceForRegistryByServiceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceInstanceForRegistryByServiceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceForRegistryByServiceNameResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceForRegistryByServiceNameResponse) SetHeaders(v map[string]*string) *GetServiceInstanceForRegistryByServiceNameResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceForRegistryByServiceNameResponse) SetBody(v *GetServiceInstanceForRegistryByServiceNameResponseBody) *GetServiceInstanceForRegistryByServiceNameResponse {
	s.Body = v
	return s
}

type DeleteServiceResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetCode(v int64) *DeleteServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceResponseBody) SetData(v map[string]interface{}) *DeleteServiceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServiceResponseBody) SetMessage(v string) *DeleteServiceResponseBody {
	s.Message = &v
	return s
}

type DeleteServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type UpdateRegistryRequest struct {
	// address
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// gatewayId
	GatewayId *int64 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateRegistryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistryRequest) SetAddress(v string) *UpdateRegistryRequest {
	s.Address = &v
	return s
}

func (s *UpdateRegistryRequest) SetDescription(v string) *UpdateRegistryRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistryRequest) SetGatewayId(v int64) *UpdateRegistryRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateRegistryRequest) SetId(v string) *UpdateRegistryRequest {
	s.Id = &v
	return s
}

func (s *UpdateRegistryRequest) SetName(v string) *UpdateRegistryRequest {
	s.Name = &v
	return s
}

func (s *UpdateRegistryRequest) SetType(v int64) *UpdateRegistryRequest {
	s.Type = &v
	return s
}

type UpdateRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegistryResponseBody) SetCode(v int64) *UpdateRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRegistryResponseBody) SetData(v map[string]interface{}) *UpdateRegistryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRegistryResponseBody) SetMessage(v string) *UpdateRegistryResponseBody {
	s.Message = &v
	return s
}

type UpdateRegistryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistryResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegistryResponse) SetHeaders(v map[string]*string) *UpdateRegistryResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegistryResponse) SetBody(v *UpdateRegistryResponseBody) *UpdateRegistryResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbSpec
	SlbSpec *string `json:"slbSpec,omitempty" xml:"slbSpec,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
	// zone
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetAutoCreateSlb(v bool) *CreateGatewayRequest {
	s.AutoCreateSlb = &v
	return s
}

func (s *CreateGatewayRequest) SetBasePath(v string) *CreateGatewayRequest {
	s.BasePath = &v
	return s
}

func (s *CreateGatewayRequest) SetEdasNamespaceId(v string) *CreateGatewayRequest {
	s.EdasNamespaceId = &v
	return s
}

func (s *CreateGatewayRequest) SetGatewayType(v string) *CreateGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetPodCidr(v string) *CreateGatewayRequest {
	s.PodCidr = &v
	return s
}

func (s *CreateGatewayRequest) SetRegion(v string) *CreateGatewayRequest {
	s.Region = &v
	return s
}

func (s *CreateGatewayRequest) SetRegionName(v string) *CreateGatewayRequest {
	s.RegionName = &v
	return s
}

func (s *CreateGatewayRequest) SetReplica(v int64) *CreateGatewayRequest {
	s.Replica = &v
	return s
}

func (s *CreateGatewayRequest) SetRuntimeOn(v string) *CreateGatewayRequest {
	s.RuntimeOn = &v
	return s
}

func (s *CreateGatewayRequest) SetSecurityGroup(v string) *CreateGatewayRequest {
	s.SecurityGroup = &v
	return s
}

func (s *CreateGatewayRequest) SetSlb(v string) *CreateGatewayRequest {
	s.Slb = &v
	return s
}

func (s *CreateGatewayRequest) SetSlbSpec(v string) *CreateGatewayRequest {
	s.SlbSpec = &v
	return s
}

func (s *CreateGatewayRequest) SetVpc(v string) *CreateGatewayRequest {
	s.Vpc = &v
	return s
}

func (s *CreateGatewayRequest) SetVswitch(v string) *CreateGatewayRequest {
	s.Vswitch = &v
	return s
}

func (s *CreateGatewayRequest) SetZone(v string) *CreateGatewayRequest {
	s.Zone = &v
	return s
}

type CreateGatewayResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetCode(v int64) *CreateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayResponseBody) SetData(v map[string]interface{}) *CreateGatewayResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayResponseBody) SetMessage(v string) *CreateGatewayResponseBody {
	s.Message = &v
	return s
}

type CreateGatewayResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CheckServiceHealthRequest struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// operationIds
	OperationIds []*int64 `json:"operationIds,omitempty" xml:"operationIds,omitempty" type:"Repeated"`
}

func (s CheckServiceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceHealthRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceHealthRequest) SetId(v int64) *CheckServiceHealthRequest {
	s.Id = &v
	return s
}

func (s *CheckServiceHealthRequest) SetOperationIds(v []*int64) *CheckServiceHealthRequest {
	s.OperationIds = v
	return s
}

type CheckServiceHealthResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*CheckServiceHealthResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CheckServiceHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceHealthResponseBody) SetCode(v int64) *CheckServiceHealthResponseBody {
	s.Code = &v
	return s
}

func (s *CheckServiceHealthResponseBody) SetData(v []*CheckServiceHealthResponseBodyData) *CheckServiceHealthResponseBody {
	s.Data = v
	return s
}

func (s *CheckServiceHealthResponseBody) SetMessage(v string) *CheckServiceHealthResponseBody {
	s.Message = &v
	return s
}

type CheckServiceHealthResponseBodyData struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*CheckServiceHealthResponseBodyDataServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s CheckServiceHealthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceHealthResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckServiceHealthResponseBodyData) SetAliasName(v string) *CheckServiceHealthResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetCreationDateTime(v string) *CheckServiceHealthResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetDescription(v string) *CheckServiceHealthResponseBodyData {
	s.Description = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetId(v int64) *CheckServiceHealthResponseBodyData {
	s.Id = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetIsAutoRefresh(v bool) *CheckServiceHealthResponseBodyData {
	s.IsAutoRefresh = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetIsHealth(v bool) *CheckServiceHealthResponseBodyData {
	s.IsHealth = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetName(v string) *CheckServiceHealthResponseBodyData {
	s.Name = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetRegistryId(v string) *CheckServiceHealthResponseBodyData {
	s.RegistryId = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetServiceEnds(v []*CheckServiceHealthResponseBodyDataServiceEnds) *CheckServiceHealthResponseBodyData {
	s.ServiceEnds = v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetServiceNameInRegistry(v string) *CheckServiceHealthResponseBodyData {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetSourceType(v int64) *CheckServiceHealthResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *CheckServiceHealthResponseBodyData) SetUpdateDateTime(v string) *CheckServiceHealthResponseBodyData {
	s.UpdateDateTime = &v
	return s
}

type CheckServiceHealthResponseBodyDataServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s CheckServiceHealthResponseBodyDataServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceHealthResponseBodyDataServiceEnds) GoString() string {
	return s.String()
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetCreationDateTime(v string) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetId(v int64) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.Id = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetIpAddress(v string) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetPort(v string) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.Port = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetServiceId(v int64) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetStatus(v int64) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.Status = &v
	return s
}

func (s *CheckServiceHealthResponseBodyDataServiceEnds) SetUpdateDateTime(v string) *CheckServiceHealthResponseBodyDataServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type CheckServiceHealthResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceHealthResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceHealthResponse) SetHeaders(v map[string]*string) *CheckServiceHealthResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceHealthResponse) SetBody(v *CheckServiceHealthResponseBody) *CheckServiceHealthResponse {
	s.Body = v
	return s
}

type CreatePolicyToApiRequest struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePolicyToApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyToApiRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyToApiRequest) SetCreationDateTime(v string) *CreatePolicyToApiRequest {
	s.CreationDateTime = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetDirection(v string) *CreatePolicyToApiRequest {
	s.Direction = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPolicyAliasName(v string) *CreatePolicyToApiRequest {
	s.PolicyAliasName = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPolicyContent(v string) *CreatePolicyToApiRequest {
	s.PolicyContent = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPolicyGroup(v string) *CreatePolicyToApiRequest {
	s.PolicyGroup = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPolicyId(v int64) *CreatePolicyToApiRequest {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPolicyName(v string) *CreatePolicyToApiRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetPriority(v int64) *CreatePolicyToApiRequest {
	s.Priority = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetScope(v string) *CreatePolicyToApiRequest {
	s.Scope = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetStatus(v bool) *CreatePolicyToApiRequest {
	s.Status = &v
	return s
}

func (s *CreatePolicyToApiRequest) SetType(v int64) *CreatePolicyToApiRequest {
	s.Type = &v
	return s
}

type CreatePolicyToApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreatePolicyToApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyToApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyToApiResponseBody) SetCode(v int64) *CreatePolicyToApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyToApiResponseBody) SetData(v map[string]interface{}) *CreatePolicyToApiResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyToApiResponseBody) SetMessage(v string) *CreatePolicyToApiResponseBody {
	s.Message = &v
	return s
}

type CreatePolicyToApiResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolicyToApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePolicyToApiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyToApiResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyToApiResponse) SetHeaders(v map[string]*string) *CreatePolicyToApiResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyToApiResponse) SetBody(v *CreatePolicyToApiResponseBody) *CreatePolicyToApiResponse {
	s.Body = v
	return s
}

type DetachPolicyResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DetachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponseBody) SetCode(v int64) *DetachPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DetachPolicyResponseBody) SetData(v map[string]interface{}) *DetachPolicyResponseBody {
	s.Data = v
	return s
}

func (s *DetachPolicyResponseBody) SetMessage(v string) *DetachPolicyResponseBody {
	s.Message = &v
	return s
}

type DetachPolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponse) SetHeaders(v map[string]*string) *DetachPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyResponse) SetBody(v *DetachPolicyResponseBody) *DetachPolicyResponse {
	s.Body = v
	return s
}

type FindTemplateResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *FindTemplateResponseBody) SetCode(v int64) *FindTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *FindTemplateResponseBody) SetData(v map[string]interface{}) *FindTemplateResponseBody {
	s.Data = v
	return s
}

func (s *FindTemplateResponseBody) SetMessage(v string) *FindTemplateResponseBody {
	s.Message = &v
	return s
}

type FindTemplateResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s FindTemplateResponse) GoString() string {
	return s.String()
}

func (s *FindTemplateResponse) SetHeaders(v map[string]*string) *FindTemplateResponse {
	s.Headers = v
	return s
}

func (s *FindTemplateResponse) SetBody(v *FindTemplateResponseBody) *FindTemplateResponse {
	s.Body = v
	return s
}

type ValidateRegistryAddressRequest struct {
	// address
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ValidateRegistryAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateRegistryAddressRequest) GoString() string {
	return s.String()
}

func (s *ValidateRegistryAddressRequest) SetAddress(v string) *ValidateRegistryAddressRequest {
	s.Address = &v
	return s
}

func (s *ValidateRegistryAddressRequest) SetType(v int64) *ValidateRegistryAddressRequest {
	s.Type = &v
	return s
}

type ValidateRegistryAddressResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidateRegistryAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateRegistryAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateRegistryAddressResponseBody) SetCode(v int64) *ValidateRegistryAddressResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateRegistryAddressResponseBody) SetData(v map[string]interface{}) *ValidateRegistryAddressResponseBody {
	s.Data = v
	return s
}

func (s *ValidateRegistryAddressResponseBody) SetMessage(v string) *ValidateRegistryAddressResponseBody {
	s.Message = &v
	return s
}

type ValidateRegistryAddressResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateRegistryAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateRegistryAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateRegistryAddressResponse) GoString() string {
	return s.String()
}

func (s *ValidateRegistryAddressResponse) SetHeaders(v map[string]*string) *ValidateRegistryAddressResponse {
	s.Headers = v
	return s
}

func (s *ValidateRegistryAddressResponse) SetBody(v *ValidateRegistryAddressResponseBody) *ValidateRegistryAddressResponse {
	s.Body = v
	return s
}

type GetApiDetailResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetApiDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetApiDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBody) SetCode(v int64) *GetApiDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiDetailResponseBody) SetData(v []*GetApiDetailResponseBodyData) *GetApiDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetApiDetailResponseBody) SetMessage(v string) *GetApiDetailResponseBody {
	s.Message = &v
	return s
}

type GetApiDetailResponseBodyData struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedServices
	AttachedServices []*GetApiDetailResponseBodyDataAttachedServices `json:"attachedServices,omitempty" xml:"attachedServices,omitempty" type:"Repeated"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// owneredPolicies
	OwneredPolicies []*GetApiDetailResponseBodyDataOwneredPolicies `json:"owneredPolicies,omitempty" xml:"owneredPolicies,omitempty" type:"Repeated"`
	// A short description of struct
	PublishedGateway *GetApiDetailResponseBodyDataPublishedGateway `json:"publishedGateway,omitempty" xml:"publishedGateway,omitempty" type:"Struct"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetApiDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBodyData) SetAliasName(v string) *GetApiDetailResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetAttachedServices(v []*GetApiDetailResponseBodyDataAttachedServices) *GetApiDetailResponseBodyData {
	s.AttachedServices = v
	return s
}

func (s *GetApiDetailResponseBodyData) SetBasePath(v string) *GetApiDetailResponseBodyData {
	s.BasePath = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetCreationDateTime(v string) *GetApiDetailResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetDescription(v string) *GetApiDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetId(v int64) *GetApiDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetName(v string) *GetApiDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetOwneredPolicies(v []*GetApiDetailResponseBodyDataOwneredPolicies) *GetApiDetailResponseBodyData {
	s.OwneredPolicies = v
	return s
}

func (s *GetApiDetailResponseBodyData) SetPublishedGateway(v *GetApiDetailResponseBodyDataPublishedGateway) *GetApiDetailResponseBodyData {
	s.PublishedGateway = v
	return s
}

func (s *GetApiDetailResponseBodyData) SetStatus(v string) *GetApiDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApiDetailResponseBodyData) SetUpdateDateTime(v string) *GetApiDetailResponseBodyData {
	s.UpdateDateTime = &v
	return s
}

type GetApiDetailResponseBodyDataAttachedServices struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*GetApiDetailResponseBodyDataAttachedServicesServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetApiDetailResponseBodyDataAttachedServices) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBodyDataAttachedServices) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetAliasName(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.AliasName = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetCreationDateTime(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.CreationDateTime = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetDescription(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.Description = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetId(v int64) *GetApiDetailResponseBodyDataAttachedServices {
	s.Id = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetIsAutoRefresh(v bool) *GetApiDetailResponseBodyDataAttachedServices {
	s.IsAutoRefresh = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetIsHealth(v bool) *GetApiDetailResponseBodyDataAttachedServices {
	s.IsHealth = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetName(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.Name = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetRegistryId(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.RegistryId = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetServiceEnds(v []*GetApiDetailResponseBodyDataAttachedServicesServiceEnds) *GetApiDetailResponseBodyDataAttachedServices {
	s.ServiceEnds = v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetServiceNameInRegistry(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetSourceType(v int64) *GetApiDetailResponseBodyDataAttachedServices {
	s.SourceType = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServices) SetUpdateDateTime(v string) *GetApiDetailResponseBodyDataAttachedServices {
	s.UpdateDateTime = &v
	return s
}

type GetApiDetailResponseBodyDataAttachedServicesServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetApiDetailResponseBodyDataAttachedServicesServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBodyDataAttachedServicesServiceEnds) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetCreationDateTime(v string) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetId(v int64) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.Id = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetIpAddress(v string) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetPort(v string) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.Port = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetServiceId(v int64) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetStatus(v int64) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.Status = &v
	return s
}

func (s *GetApiDetailResponseBodyDataAttachedServicesServiceEnds) SetUpdateDateTime(v string) *GetApiDetailResponseBodyDataAttachedServicesServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type GetApiDetailResponseBodyDataOwneredPolicies struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetApiDetailResponseBodyDataOwneredPolicies) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBodyDataOwneredPolicies) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetApiId(v int64) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.ApiId = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetApiName(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.ApiName = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetCreationDateTime(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.CreationDateTime = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetDirection(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Direction = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetId(v int64) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Id = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPolicyAliasName(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.PolicyAliasName = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPolicyContent(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.PolicyContent = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPolicyGroup(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.PolicyGroup = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPolicyId(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.PolicyId = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPolicyName(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.PolicyName = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetPriority(v int64) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Priority = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetScope(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Scope = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetStatus(v bool) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Status = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetType(v int64) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.Type = &v
	return s
}

func (s *GetApiDetailResponseBodyDataOwneredPolicies) SetUpdateDateTime(v string) *GetApiDetailResponseBodyDataOwneredPolicies {
	s.UpdateDateTime = &v
	return s
}

type GetApiDetailResponseBodyDataPublishedGateway struct {
	// armsInfo
	ArmsInfo *string `json:"armsInfo,omitempty" xml:"armsInfo,omitempty"`
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbAccessAddr
	SlbAccessAddr *string `json:"slbAccessAddr,omitempty" xml:"slbAccessAddr,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
}

func (s GetApiDetailResponseBodyDataPublishedGateway) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponseBodyDataPublishedGateway) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetArmsInfo(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.ArmsInfo = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetAutoCreateSlb(v bool) *GetApiDetailResponseBodyDataPublishedGateway {
	s.AutoCreateSlb = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetBasePath(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.BasePath = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetCreationDateTime(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.CreationDateTime = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetEdasNamespaceId(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.EdasNamespaceId = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetGatewayType(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.GatewayType = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetId(v int64) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Id = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetName(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Name = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetPodCidr(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.PodCidr = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetRegion(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Region = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetRegionName(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.RegionName = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetReplica(v int64) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Replica = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetRuntimeOn(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.RuntimeOn = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetSecurityGroup(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.SecurityGroup = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetSlb(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Slb = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetSlbAccessAddr(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.SlbAccessAddr = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetStatus(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Status = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetVpc(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Vpc = &v
	return s
}

func (s *GetApiDetailResponseBodyDataPublishedGateway) SetVswitch(v string) *GetApiDetailResponseBodyDataPublishedGateway {
	s.Vswitch = &v
	return s
}

type GetApiDetailResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApiDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApiDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApiDetailResponse) GoString() string {
	return s.String()
}

func (s *GetApiDetailResponse) SetHeaders(v map[string]*string) *GetApiDetailResponse {
	s.Headers = v
	return s
}

func (s *GetApiDetailResponse) SetBody(v *GetApiDetailResponseBody) *GetApiDetailResponse {
	s.Body = v
	return s
}

type CreateSpecialRouteForRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateSpecialRouteForRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpecialRouteForRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpecialRouteForRegistryResponseBody) SetCode(v int64) *CreateSpecialRouteForRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSpecialRouteForRegistryResponseBody) SetData(v map[string]interface{}) *CreateSpecialRouteForRegistryResponseBody {
	s.Data = v
	return s
}

func (s *CreateSpecialRouteForRegistryResponseBody) SetMessage(v string) *CreateSpecialRouteForRegistryResponseBody {
	s.Message = &v
	return s
}

type CreateSpecialRouteForRegistryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSpecialRouteForRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpecialRouteForRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpecialRouteForRegistryResponse) GoString() string {
	return s.String()
}

func (s *CreateSpecialRouteForRegistryResponse) SetHeaders(v map[string]*string) *CreateSpecialRouteForRegistryResponse {
	s.Headers = v
	return s
}

func (s *CreateSpecialRouteForRegistryResponse) SetBody(v *CreateSpecialRouteForRegistryResponseBody) *CreateSpecialRouteForRegistryResponse {
	s.Body = v
	return s
}

type PublishApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s PublishApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishApiResponseBody) GoString() string {
	return s.String()
}

func (s *PublishApiResponseBody) SetCode(v int64) *PublishApiResponseBody {
	s.Code = &v
	return s
}

func (s *PublishApiResponseBody) SetData(v map[string]interface{}) *PublishApiResponseBody {
	s.Data = v
	return s
}

func (s *PublishApiResponseBody) SetMessage(v string) *PublishApiResponseBody {
	s.Message = &v
	return s
}

type PublishApiResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishApiResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishApiResponse) GoString() string {
	return s.String()
}

func (s *PublishApiResponse) SetHeaders(v map[string]*string) *PublishApiResponse {
	s.Headers = v
	return s
}

func (s *PublishApiResponse) SetBody(v *PublishApiResponseBody) *PublishApiResponse {
	s.Body = v
	return s
}

type CreateGatewayLogEtlResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateGatewayLogEtlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayLogEtlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayLogEtlResponseBody) SetCode(v int64) *CreateGatewayLogEtlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayLogEtlResponseBody) SetData(v map[string]interface{}) *CreateGatewayLogEtlResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayLogEtlResponseBody) SetMessage(v string) *CreateGatewayLogEtlResponseBody {
	s.Message = &v
	return s
}

type CreateGatewayLogEtlResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGatewayLogEtlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayLogEtlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayLogEtlResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayLogEtlResponse) SetHeaders(v map[string]*string) *CreateGatewayLogEtlResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayLogEtlResponse) SetBody(v *CreateGatewayLogEtlResponseBody) *CreateGatewayLogEtlResponse {
	s.Body = v
	return s
}

type FindPoliciesRequest struct {
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
}

func (s FindPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesRequest) GoString() string {
	return s.String()
}

func (s *FindPoliciesRequest) SetPageNumber(v int64) *FindPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *FindPoliciesRequest) SetPageSize(v int64) *FindPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *FindPoliciesRequest) SetName(v string) *FindPoliciesRequest {
	s.Name = &v
	return s
}

func (s *FindPoliciesRequest) SetAliasName(v string) *FindPoliciesRequest {
	s.AliasName = &v
	return s
}

func (s *FindPoliciesRequest) SetType(v int64) *FindPoliciesRequest {
	s.Type = &v
	return s
}

func (s *FindPoliciesRequest) SetGroup(v string) *FindPoliciesRequest {
	s.Group = &v
	return s
}

type FindPoliciesResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *FindPoliciesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *FindPoliciesResponseBody) SetCode(v int64) *FindPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *FindPoliciesResponseBody) SetData(v *FindPoliciesResponseBodyData) *FindPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *FindPoliciesResponseBody) SetMessage(v string) *FindPoliciesResponseBody {
	s.Message = &v
	return s
}

type FindPoliciesResponseBodyData struct {
	// list
	List []*FindPoliciesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s FindPoliciesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindPoliciesResponseBodyData) SetList(v []*FindPoliciesResponseBodyDataList) *FindPoliciesResponseBodyData {
	s.List = v
	return s
}

func (s *FindPoliciesResponseBodyData) SetTotalCount(v int64) *FindPoliciesResponseBodyData {
	s.TotalCount = &v
	return s
}

type FindPoliciesResponseBodyDataList struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedApi
	AttachedApi []*FindPoliciesResponseBodyDataListAttachedApi `json:"attachedApi,omitempty" xml:"attachedApi,omitempty" type:"Repeated"`
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyTypeName
	PolicyTypeName *string `json:"policyTypeName,omitempty" xml:"policyTypeName,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindPoliciesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *FindPoliciesResponseBodyDataList) SetAliasName(v string) *FindPoliciesResponseBodyDataList {
	s.AliasName = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetAttachedApi(v []*FindPoliciesResponseBodyDataListAttachedApi) *FindPoliciesResponseBodyDataList {
	s.AttachedApi = v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetContent(v string) *FindPoliciesResponseBodyDataList {
	s.Content = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetCreationDateTime(v string) *FindPoliciesResponseBodyDataList {
	s.CreationDateTime = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetId(v int64) *FindPoliciesResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetName(v string) *FindPoliciesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetPolicyGroup(v string) *FindPoliciesResponseBodyDataList {
	s.PolicyGroup = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetPolicyTypeName(v string) *FindPoliciesResponseBodyDataList {
	s.PolicyTypeName = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetType(v int64) *FindPoliciesResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *FindPoliciesResponseBodyDataList) SetUpdateDateTime(v string) *FindPoliciesResponseBodyDataList {
	s.UpdateDateTime = &v
	return s
}

type FindPoliciesResponseBodyDataListAttachedApi struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindPoliciesResponseBodyDataListAttachedApi) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesResponseBodyDataListAttachedApi) GoString() string {
	return s.String()
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetApiId(v int64) *FindPoliciesResponseBodyDataListAttachedApi {
	s.ApiId = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetApiName(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.ApiName = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetCreationDateTime(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.CreationDateTime = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetDirection(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Direction = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetId(v int64) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Id = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPolicyAliasName(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.PolicyAliasName = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPolicyContent(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.PolicyContent = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPolicyGroup(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.PolicyGroup = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPolicyId(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.PolicyId = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPolicyName(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.PolicyName = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetPriority(v int64) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Priority = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetScope(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Scope = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetStatus(v bool) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Status = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetType(v int64) *FindPoliciesResponseBodyDataListAttachedApi {
	s.Type = &v
	return s
}

func (s *FindPoliciesResponseBodyDataListAttachedApi) SetUpdateDateTime(v string) *FindPoliciesResponseBodyDataListAttachedApi {
	s.UpdateDateTime = &v
	return s
}

type FindPoliciesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s FindPoliciesResponse) GoString() string {
	return s.String()
}

func (s *FindPoliciesResponse) SetHeaders(v map[string]*string) *FindPoliciesResponse {
	s.Headers = v
	return s
}

func (s *FindPoliciesResponse) SetBody(v *FindPoliciesResponseBody) *FindPoliciesResponse {
	s.Body = v
	return s
}

type AttachPolicyRequest struct {
	// data
	Data []*AttachPolicyRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s AttachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) SetData(v []*AttachPolicyRequestData) *AttachPolicyRequest {
	s.Data = v
	return s
}

type AttachPolicyRequestData struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AttachPolicyRequestData) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyRequestData) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequestData) SetCreationDateTime(v string) *AttachPolicyRequestData {
	s.CreationDateTime = &v
	return s
}

func (s *AttachPolicyRequestData) SetDirection(v string) *AttachPolicyRequestData {
	s.Direction = &v
	return s
}

func (s *AttachPolicyRequestData) SetPolicyAliasName(v string) *AttachPolicyRequestData {
	s.PolicyAliasName = &v
	return s
}

func (s *AttachPolicyRequestData) SetPolicyContent(v string) *AttachPolicyRequestData {
	s.PolicyContent = &v
	return s
}

func (s *AttachPolicyRequestData) SetPolicyGroup(v string) *AttachPolicyRequestData {
	s.PolicyGroup = &v
	return s
}

func (s *AttachPolicyRequestData) SetPolicyId(v int64) *AttachPolicyRequestData {
	s.PolicyId = &v
	return s
}

func (s *AttachPolicyRequestData) SetPolicyName(v string) *AttachPolicyRequestData {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyRequestData) SetPriority(v int64) *AttachPolicyRequestData {
	s.Priority = &v
	return s
}

func (s *AttachPolicyRequestData) SetScope(v string) *AttachPolicyRequestData {
	s.Scope = &v
	return s
}

func (s *AttachPolicyRequestData) SetStatus(v bool) *AttachPolicyRequestData {
	s.Status = &v
	return s
}

func (s *AttachPolicyRequestData) SetType(v int64) *AttachPolicyRequestData {
	s.Type = &v
	return s
}

type AttachPolicyResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AttachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponseBody) SetCode(v int64) *AttachPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *AttachPolicyResponseBody) SetData(v map[string]interface{}) *AttachPolicyResponseBody {
	s.Data = v
	return s
}

func (s *AttachPolicyResponseBody) SetMessage(v string) *AttachPolicyResponseBody {
	s.Message = &v
	return s
}

type AttachPolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponse) SetHeaders(v map[string]*string) *AttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyResponse) SetBody(v *AttachPolicyResponseBody) *AttachPolicyResponse {
	s.Body = v
	return s
}

type FindRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*FindRegistryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *FindRegistryResponseBody) SetCode(v int64) *FindRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *FindRegistryResponseBody) SetData(v []*FindRegistryResponseBodyData) *FindRegistryResponseBody {
	s.Data = v
	return s
}

func (s *FindRegistryResponseBody) SetMessage(v string) *FindRegistryResponseBody {
	s.Message = &v
	return s
}

type FindRegistryResponseBodyData struct {
	// address
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// gatewayId
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindRegistryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindRegistryResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindRegistryResponseBodyData) SetAddress(v string) *FindRegistryResponseBodyData {
	s.Address = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetCreationDateTime(v string) *FindRegistryResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetDescription(v string) *FindRegistryResponseBodyData {
	s.Description = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetGatewayId(v string) *FindRegistryResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetId(v int64) *FindRegistryResponseBodyData {
	s.Id = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetName(v string) *FindRegistryResponseBodyData {
	s.Name = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetType(v int64) *FindRegistryResponseBodyData {
	s.Type = &v
	return s
}

func (s *FindRegistryResponseBodyData) SetUpdateDateTime(v string) *FindRegistryResponseBodyData {
	s.UpdateDateTime = &v
	return s
}

type FindRegistryResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s FindRegistryResponse) GoString() string {
	return s.String()
}

func (s *FindRegistryResponse) SetHeaders(v map[string]*string) *FindRegistryResponse {
	s.Headers = v
	return s
}

func (s *FindRegistryResponse) SetBody(v *FindRegistryResponseBody) *FindRegistryResponse {
	s.Body = v
	return s
}

type GetAuthTicketByIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// cookie
	Cookie map[string]interface{} `json:"cookie,omitempty" xml:"cookie,omitempty"`
}

func (s GetAuthTicketByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTicketByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAuthTicketByIdHeaders) SetCommonHeaders(v map[string]*string) *GetAuthTicketByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAuthTicketByIdHeaders) SetCookie(v map[string]interface{}) *GetAuthTicketByIdHeaders {
	s.Cookie = v
	return s
}

type GetAuthTicketByIdShrinkHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// cookie
	CookieShrink *string `json:"cookie,omitempty" xml:"cookie,omitempty"`
}

func (s GetAuthTicketByIdShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTicketByIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetAuthTicketByIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetAuthTicketByIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAuthTicketByIdShrinkHeaders) SetCookieShrink(v string) *GetAuthTicketByIdShrinkHeaders {
	s.CookieShrink = &v
	return s
}

type GetAuthTicketByIdResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetAuthTicketByIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAuthTicketByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTicketByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTicketByIdResponseBody) SetCode(v int64) *GetAuthTicketByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthTicketByIdResponseBody) SetData(v []*GetAuthTicketByIdResponseBodyData) *GetAuthTicketByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthTicketByIdResponseBody) SetMessage(v string) *GetAuthTicketByIdResponseBody {
	s.Message = &v
	return s
}

type GetAuthTicketByIdResponseBodyData struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// serverKey
	ServerKey *string `json:"serverKey,omitempty" xml:"serverKey,omitempty"`
	// ticketType
	TicketType *string `json:"ticketType,omitempty" xml:"ticketType,omitempty"`
	// validEndTime
	ValidEndTime *string `json:"validEndTime,omitempty" xml:"validEndTime,omitempty"`
	// validStartTime
	ValidStartTime *string `json:"validStartTime,omitempty" xml:"validStartTime,omitempty"`
}

func (s GetAuthTicketByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTicketByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthTicketByIdResponseBodyData) SetClientToken(v string) *GetAuthTicketByIdResponseBodyData {
	s.ClientToken = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetComment(v string) *GetAuthTicketByIdResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetId(v int64) *GetAuthTicketByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetName(v string) *GetAuthTicketByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetServerKey(v string) *GetAuthTicketByIdResponseBodyData {
	s.ServerKey = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetTicketType(v string) *GetAuthTicketByIdResponseBodyData {
	s.TicketType = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetValidEndTime(v string) *GetAuthTicketByIdResponseBodyData {
	s.ValidEndTime = &v
	return s
}

func (s *GetAuthTicketByIdResponseBodyData) SetValidStartTime(v string) *GetAuthTicketByIdResponseBodyData {
	s.ValidStartTime = &v
	return s
}

type GetAuthTicketByIdResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthTicketByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthTicketByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTicketByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTicketByIdResponse) SetHeaders(v map[string]*string) *GetAuthTicketByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTicketByIdResponse) SetBody(v *GetAuthTicketByIdResponseBody) *GetAuthTicketByIdResponse {
	s.Body = v
	return s
}

type CreateRegistryRequest struct {
	// address
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// gatewayId
	GatewayId *int64 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRegistryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistryRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistryRequest) SetAddress(v string) *CreateRegistryRequest {
	s.Address = &v
	return s
}

func (s *CreateRegistryRequest) SetDescription(v string) *CreateRegistryRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistryRequest) SetGatewayId(v int64) *CreateRegistryRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateRegistryRequest) SetId(v string) *CreateRegistryRequest {
	s.Id = &v
	return s
}

func (s *CreateRegistryRequest) SetName(v string) *CreateRegistryRequest {
	s.Name = &v
	return s
}

func (s *CreateRegistryRequest) SetType(v int64) *CreateRegistryRequest {
	s.Type = &v
	return s
}

type CreateRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRegistryResponseBody) SetCode(v int64) *CreateRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRegistryResponseBody) SetData(v map[string]interface{}) *CreateRegistryResponseBody {
	s.Data = v
	return s
}

func (s *CreateRegistryResponseBody) SetMessage(v string) *CreateRegistryResponseBody {
	s.Message = &v
	return s
}

type CreateRegistryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistryResponse) GoString() string {
	return s.String()
}

func (s *CreateRegistryResponse) SetHeaders(v map[string]*string) *CreateRegistryResponse {
	s.Headers = v
	return s
}

func (s *CreateRegistryResponse) SetBody(v *CreateRegistryResponseBody) *CreateRegistryResponse {
	s.Body = v
	return s
}

type RecycleApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s RecycleApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecycleApiResponseBody) GoString() string {
	return s.String()
}

func (s *RecycleApiResponseBody) SetCode(v int64) *RecycleApiResponseBody {
	s.Code = &v
	return s
}

func (s *RecycleApiResponseBody) SetData(v map[string]interface{}) *RecycleApiResponseBody {
	s.Data = v
	return s
}

func (s *RecycleApiResponseBody) SetMessage(v string) *RecycleApiResponseBody {
	s.Message = &v
	return s
}

type RecycleApiResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecycleApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecycleApiResponse) String() string {
	return tea.Prettify(s)
}

func (s RecycleApiResponse) GoString() string {
	return s.String()
}

func (s *RecycleApiResponse) SetHeaders(v map[string]*string) *RecycleApiResponse {
	s.Headers = v
	return s
}

func (s *RecycleApiResponse) SetBody(v *RecycleApiResponseBody) *RecycleApiResponse {
	s.Body = v
	return s
}

type CreateAuthTicketRequest struct {
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// gatewayId
	GatewayId *int64 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ticketType
	TicketType    *string `json:"ticketType,omitempty" xml:"ticketType,omitempty"`
	ValidDuration *int64  `json:"validDuration,omitempty" xml:"validDuration,omitempty"`
}

func (s CreateAuthTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthTicketRequest) SetComment(v string) *CreateAuthTicketRequest {
	s.Comment = &v
	return s
}

func (s *CreateAuthTicketRequest) SetGatewayId(v int64) *CreateAuthTicketRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateAuthTicketRequest) SetName(v string) *CreateAuthTicketRequest {
	s.Name = &v
	return s
}

func (s *CreateAuthTicketRequest) SetTicketType(v string) *CreateAuthTicketRequest {
	s.TicketType = &v
	return s
}

func (s *CreateAuthTicketRequest) SetValidDuration(v int64) *CreateAuthTicketRequest {
	s.ValidDuration = &v
	return s
}

type CreateAuthTicketResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateAuthTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthTicketResponseBody) SetCode(v int64) *CreateAuthTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAuthTicketResponseBody) SetData(v map[string]interface{}) *CreateAuthTicketResponseBody {
	s.Data = v
	return s
}

func (s *CreateAuthTicketResponseBody) SetMessage(v string) *CreateAuthTicketResponseBody {
	s.Message = &v
	return s
}

type CreateAuthTicketResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAuthTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAuthTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthTicketResponse) SetHeaders(v map[string]*string) *CreateAuthTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthTicketResponse) SetBody(v *CreateAuthTicketResponseBody) *CreateAuthTicketResponse {
	s.Body = v
	return s
}

type DeleteGatewayResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetCode(v int64) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetData(v map[string]interface{}) *DeleteGatewayResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type FindServiceResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*FindServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FindServiceResponseBody) SetCode(v int64) *FindServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FindServiceResponseBody) SetData(v []*FindServiceResponseBodyData) *FindServiceResponseBody {
	s.Data = v
	return s
}

func (s *FindServiceResponseBody) SetMessage(v string) *FindServiceResponseBody {
	s.Message = &v
	return s
}

type FindServiceResponseBodyData struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*FindServiceResponseBodyDataServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindServiceResponseBodyData) SetAliasName(v string) *FindServiceResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *FindServiceResponseBodyData) SetCreationDateTime(v string) *FindServiceResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *FindServiceResponseBodyData) SetDescription(v string) *FindServiceResponseBodyData {
	s.Description = &v
	return s
}

func (s *FindServiceResponseBodyData) SetId(v int64) *FindServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *FindServiceResponseBodyData) SetIsAutoRefresh(v bool) *FindServiceResponseBodyData {
	s.IsAutoRefresh = &v
	return s
}

func (s *FindServiceResponseBodyData) SetIsHealth(v bool) *FindServiceResponseBodyData {
	s.IsHealth = &v
	return s
}

func (s *FindServiceResponseBodyData) SetName(v string) *FindServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *FindServiceResponseBodyData) SetRegistryId(v string) *FindServiceResponseBodyData {
	s.RegistryId = &v
	return s
}

func (s *FindServiceResponseBodyData) SetServiceEnds(v []*FindServiceResponseBodyDataServiceEnds) *FindServiceResponseBodyData {
	s.ServiceEnds = v
	return s
}

func (s *FindServiceResponseBodyData) SetServiceNameInRegistry(v string) *FindServiceResponseBodyData {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *FindServiceResponseBodyData) SetSourceType(v int64) *FindServiceResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *FindServiceResponseBodyData) SetUpdateDateTime(v string) *FindServiceResponseBodyData {
	s.UpdateDateTime = &v
	return s
}

type FindServiceResponseBodyDataServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindServiceResponseBodyDataServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s FindServiceResponseBodyDataServiceEnds) GoString() string {
	return s.String()
}

func (s *FindServiceResponseBodyDataServiceEnds) SetCreationDateTime(v string) *FindServiceResponseBodyDataServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetId(v int64) *FindServiceResponseBodyDataServiceEnds {
	s.Id = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetIpAddress(v string) *FindServiceResponseBodyDataServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetPort(v string) *FindServiceResponseBodyDataServiceEnds {
	s.Port = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetServiceId(v int64) *FindServiceResponseBodyDataServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetStatus(v int64) *FindServiceResponseBodyDataServiceEnds {
	s.Status = &v
	return s
}

func (s *FindServiceResponseBodyDataServiceEnds) SetUpdateDateTime(v string) *FindServiceResponseBodyDataServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type FindServiceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s FindServiceResponse) GoString() string {
	return s.String()
}

func (s *FindServiceResponse) SetHeaders(v map[string]*string) *FindServiceResponse {
	s.Headers = v
	return s
}

func (s *FindServiceResponse) SetBody(v *FindServiceResponseBody) *FindServiceResponse {
	s.Body = v
	return s
}

type DeletePolicyByIdResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeletePolicyByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyByIdResponseBody) SetCode(v int64) *DeletePolicyByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyByIdResponseBody) SetData(v map[string]interface{}) *DeletePolicyByIdResponseBody {
	s.Data = v
	return s
}

func (s *DeletePolicyByIdResponseBody) SetMessage(v string) *DeletePolicyByIdResponseBody {
	s.Message = &v
	return s
}

type DeletePolicyByIdResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePolicyByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePolicyByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyByIdResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyByIdResponse) SetHeaders(v map[string]*string) *DeletePolicyByIdResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyByIdResponse) SetBody(v *DeletePolicyByIdResponseBody) *DeletePolicyByIdResponse {
	s.Body = v
	return s
}

type DeleteApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiResponseBody) SetCode(v int64) *DeleteApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApiResponseBody) SetData(v map[string]interface{}) *DeleteApiResponseBody {
	s.Data = v
	return s
}

func (s *DeleteApiResponseBody) SetMessage(v string) *DeleteApiResponseBody {
	s.Message = &v
	return s
}

type DeleteApiResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiResponse) SetHeaders(v map[string]*string) *DeleteApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiResponse) SetBody(v *DeleteApiResponseBody) *DeleteApiResponse {
	s.Body = v
	return s
}

type FindAuthTicketsRequest struct {
	// gatewayId
	GatewayId *int64 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s FindAuthTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s FindAuthTicketsRequest) GoString() string {
	return s.String()
}

func (s *FindAuthTicketsRequest) SetGatewayId(v int64) *FindAuthTicketsRequest {
	s.GatewayId = &v
	return s
}

func (s *FindAuthTicketsRequest) SetName(v string) *FindAuthTicketsRequest {
	s.Name = &v
	return s
}

func (s *FindAuthTicketsRequest) SetPageNumber(v int64) *FindAuthTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *FindAuthTicketsRequest) SetPageSize(v int64) *FindAuthTicketsRequest {
	s.PageSize = &v
	return s
}

type FindAuthTicketsResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *FindAuthTicketsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindAuthTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindAuthTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *FindAuthTicketsResponseBody) SetCode(v int64) *FindAuthTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *FindAuthTicketsResponseBody) SetData(v *FindAuthTicketsResponseBodyData) *FindAuthTicketsResponseBody {
	s.Data = v
	return s
}

func (s *FindAuthTicketsResponseBody) SetMessage(v string) *FindAuthTicketsResponseBody {
	s.Message = &v
	return s
}

type FindAuthTicketsResponseBodyData struct {
	// list
	List []*FindAuthTicketsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s FindAuthTicketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindAuthTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindAuthTicketsResponseBodyData) SetList(v []*FindAuthTicketsResponseBodyDataList) *FindAuthTicketsResponseBodyData {
	s.List = v
	return s
}

func (s *FindAuthTicketsResponseBodyData) SetTotalCount(v int64) *FindAuthTicketsResponseBodyData {
	s.TotalCount = &v
	return s
}

type FindAuthTicketsResponseBodyDataList struct {
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// serverKey
	ServerKey *string `json:"serverKey,omitempty" xml:"serverKey,omitempty"`
	// ticketType
	TicketType *string `json:"ticketType,omitempty" xml:"ticketType,omitempty"`
	// validEndTime
	ValidEndTime *string `json:"validEndTime,omitempty" xml:"validEndTime,omitempty"`
	// validStartTime
	ValidStartTime *string `json:"validStartTime,omitempty" xml:"validStartTime,omitempty"`
}

func (s FindAuthTicketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s FindAuthTicketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *FindAuthTicketsResponseBodyDataList) SetClientToken(v string) *FindAuthTicketsResponseBodyDataList {
	s.ClientToken = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetComment(v string) *FindAuthTicketsResponseBodyDataList {
	s.Comment = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetId(v int64) *FindAuthTicketsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetName(v string) *FindAuthTicketsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetServerKey(v string) *FindAuthTicketsResponseBodyDataList {
	s.ServerKey = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetTicketType(v string) *FindAuthTicketsResponseBodyDataList {
	s.TicketType = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetValidEndTime(v string) *FindAuthTicketsResponseBodyDataList {
	s.ValidEndTime = &v
	return s
}

func (s *FindAuthTicketsResponseBodyDataList) SetValidStartTime(v string) *FindAuthTicketsResponseBodyDataList {
	s.ValidStartTime = &v
	return s
}

type FindAuthTicketsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindAuthTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindAuthTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s FindAuthTicketsResponse) GoString() string {
	return s.String()
}

func (s *FindAuthTicketsResponse) SetHeaders(v map[string]*string) *FindAuthTicketsResponse {
	s.Headers = v
	return s
}

func (s *FindAuthTicketsResponse) SetBody(v *FindAuthTicketsResponseBody) *FindAuthTicketsResponse {
	s.Body = v
	return s
}

type UpdatePolicyRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) SetAliasName(v string) *UpdatePolicyRequest {
	s.AliasName = &v
	return s
}

func (s *UpdatePolicyRequest) SetContent(v string) *UpdatePolicyRequest {
	s.Content = &v
	return s
}

func (s *UpdatePolicyRequest) SetId(v int64) *UpdatePolicyRequest {
	s.Id = &v
	return s
}

func (s *UpdatePolicyRequest) SetName(v string) *UpdatePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolicyRequest) SetPolicyGroup(v string) *UpdatePolicyRequest {
	s.PolicyGroup = &v
	return s
}

func (s *UpdatePolicyRequest) SetType(v int64) *UpdatePolicyRequest {
	s.Type = &v
	return s
}

type UpdatePolicyResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) SetCode(v int64) *UpdatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetData(v map[string]interface{}) *UpdatePolicyResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePolicyResponseBody) SetMessage(v string) *UpdatePolicyResponseBody {
	s.Message = &v
	return s
}

type UpdatePolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdatePolicyResponse) SetBody(v *UpdatePolicyResponseBody) *UpdatePolicyResponse {
	s.Body = v
	return s
}

type UpdateAuthTicketRequest struct {
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateAuthTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthTicketRequest) SetComment(v string) *UpdateAuthTicketRequest {
	s.Comment = &v
	return s
}

func (s *UpdateAuthTicketRequest) SetId(v int64) *UpdateAuthTicketRequest {
	s.Id = &v
	return s
}

type UpdateAuthTicketResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateAuthTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthTicketResponseBody) SetCode(v int64) *UpdateAuthTicketResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAuthTicketResponseBody) SetData(v map[string]interface{}) *UpdateAuthTicketResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAuthTicketResponseBody) SetMessage(v string) *UpdateAuthTicketResponseBody {
	s.Message = &v
	return s
}

type UpdateAuthTicketResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAuthTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAuthTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthTicketResponse) SetHeaders(v map[string]*string) *UpdateAuthTicketResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthTicketResponse) SetBody(v *UpdateAuthTicketResponseBody) *UpdateAuthTicketResponse {
	s.Body = v
	return s
}

type InstallArmsAgentResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InstallArmsAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallArmsAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallArmsAgentResponseBody) SetCode(v int64) *InstallArmsAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallArmsAgentResponseBody) SetData(v map[string]interface{}) *InstallArmsAgentResponseBody {
	s.Data = v
	return s
}

func (s *InstallArmsAgentResponseBody) SetMessage(v string) *InstallArmsAgentResponseBody {
	s.Message = &v
	return s
}

type InstallArmsAgentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallArmsAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallArmsAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallArmsAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallArmsAgentResponse) SetHeaders(v map[string]*string) *InstallArmsAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallArmsAgentResponse) SetBody(v *InstallArmsAgentResponseBody) *InstallArmsAgentResponse {
	s.Body = v
	return s
}

type DeleteAuthTicketResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteAuthTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthTicketResponseBody) SetCode(v int64) *DeleteAuthTicketResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAuthTicketResponseBody) SetData(v map[string]interface{}) *DeleteAuthTicketResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAuthTicketResponseBody) SetMessage(v string) *DeleteAuthTicketResponseBody {
	s.Message = &v
	return s
}

type DeleteAuthTicketResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAuthTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAuthTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthTicketResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthTicketResponse) SetHeaders(v map[string]*string) *DeleteAuthTicketResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthTicketResponse) SetBody(v *DeleteAuthTicketResponseBody) *DeleteAuthTicketResponse {
	s.Body = v
	return s
}

type GetPolicyByIdResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetPolicyByIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetPolicyByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyByIdResponseBody) SetCode(v int64) *GetPolicyByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyByIdResponseBody) SetData(v []*GetPolicyByIdResponseBodyData) *GetPolicyByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyByIdResponseBody) SetMessage(v string) *GetPolicyByIdResponseBody {
	s.Message = &v
	return s
}

type GetPolicyByIdResponseBodyData struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedApi
	AttachedApi []*GetPolicyByIdResponseBodyDataAttachedApi `json:"attachedApi,omitempty" xml:"attachedApi,omitempty" type:"Repeated"`
	// content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyTypeName
	PolicyTypeName *string `json:"policyTypeName,omitempty" xml:"policyTypeName,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetPolicyByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPolicyByIdResponseBodyData) SetAliasName(v string) *GetPolicyByIdResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetAttachedApi(v []*GetPolicyByIdResponseBodyDataAttachedApi) *GetPolicyByIdResponseBodyData {
	s.AttachedApi = v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetContent(v string) *GetPolicyByIdResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetCreationDateTime(v string) *GetPolicyByIdResponseBodyData {
	s.CreationDateTime = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetId(v int64) *GetPolicyByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetName(v string) *GetPolicyByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetPolicyGroup(v string) *GetPolicyByIdResponseBodyData {
	s.PolicyGroup = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetPolicyTypeName(v string) *GetPolicyByIdResponseBodyData {
	s.PolicyTypeName = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetType(v int64) *GetPolicyByIdResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetPolicyByIdResponseBodyData) SetUpdateDateTime(v string) *GetPolicyByIdResponseBodyData {
	s.UpdateDateTime = &v
	return s
}

type GetPolicyByIdResponseBodyDataAttachedApi struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetPolicyByIdResponseBodyDataAttachedApi) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyByIdResponseBodyDataAttachedApi) GoString() string {
	return s.String()
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetApiId(v int64) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.ApiId = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetApiName(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.ApiName = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetCreationDateTime(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.CreationDateTime = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetDirection(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Direction = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetId(v int64) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Id = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPolicyAliasName(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.PolicyAliasName = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPolicyContent(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.PolicyContent = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPolicyGroup(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.PolicyGroup = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPolicyId(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPolicyName(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetPriority(v int64) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Priority = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetScope(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Scope = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetStatus(v bool) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Status = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetType(v int64) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.Type = &v
	return s
}

func (s *GetPolicyByIdResponseBodyDataAttachedApi) SetUpdateDateTime(v string) *GetPolicyByIdResponseBodyDataAttachedApi {
	s.UpdateDateTime = &v
	return s
}

type GetPolicyByIdResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPolicyByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPolicyByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyByIdResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyByIdResponse) SetHeaders(v map[string]*string) *GetPolicyByIdResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyByIdResponse) SetBody(v *GetPolicyByIdResponseBody) *GetPolicyByIdResponse {
	s.Body = v
	return s
}

type DeleteRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistryResponseBody) SetCode(v int64) *DeleteRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRegistryResponseBody) SetData(v map[string]interface{}) *DeleteRegistryResponseBody {
	s.Data = v
	return s
}

func (s *DeleteRegistryResponseBody) SetMessage(v string) *DeleteRegistryResponseBody {
	s.Message = &v
	return s
}

type DeleteRegistryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistryResponse) SetHeaders(v map[string]*string) *DeleteRegistryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistryResponse) SetBody(v *DeleteRegistryResponseBody) *DeleteRegistryResponse {
	s.Body = v
	return s
}

type GetPolicyOwnedByApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]*DataValue `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetPolicyOwnedByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyOwnedByApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyOwnedByApiResponseBody) SetCode(v int64) *GetPolicyOwnedByApiResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyOwnedByApiResponseBody) SetData(v map[string]*DataValue) *GetPolicyOwnedByApiResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyOwnedByApiResponseBody) SetMessage(v string) *GetPolicyOwnedByApiResponseBody {
	s.Message = &v
	return s
}

type GetPolicyOwnedByApiResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPolicyOwnedByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPolicyOwnedByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyOwnedByApiResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyOwnedByApiResponse) SetHeaders(v map[string]*string) *GetPolicyOwnedByApiResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyOwnedByApiResponse) SetBody(v *GetPolicyOwnedByApiResponseBody) *GetPolicyOwnedByApiResponse {
	s.Body = v
	return s
}

type UpdateApiRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedServices
	AttachedServices []*UpdateApiRequestAttachedServices `json:"attachedServices,omitempty" xml:"attachedServices,omitempty" type:"Repeated"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// owneredPolicies
	OwneredPolicies []*UpdateApiRequestOwneredPolicies `json:"owneredPolicies,omitempty" xml:"owneredPolicies,omitempty" type:"Repeated"`
	// A short description of struct
	PublishedGateway *UpdateApiRequestPublishedGateway `json:"publishedGateway,omitempty" xml:"publishedGateway,omitempty" type:"Struct"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s UpdateApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiRequest) SetAliasName(v string) *UpdateApiRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateApiRequest) SetAttachedServices(v []*UpdateApiRequestAttachedServices) *UpdateApiRequest {
	s.AttachedServices = v
	return s
}

func (s *UpdateApiRequest) SetBasePath(v string) *UpdateApiRequest {
	s.BasePath = &v
	return s
}

func (s *UpdateApiRequest) SetCreationDateTime(v string) *UpdateApiRequest {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateApiRequest) SetDescription(v string) *UpdateApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiRequest) SetId(v int64) *UpdateApiRequest {
	s.Id = &v
	return s
}

func (s *UpdateApiRequest) SetName(v string) *UpdateApiRequest {
	s.Name = &v
	return s
}

func (s *UpdateApiRequest) SetOwneredPolicies(v []*UpdateApiRequestOwneredPolicies) *UpdateApiRequest {
	s.OwneredPolicies = v
	return s
}

func (s *UpdateApiRequest) SetPublishedGateway(v *UpdateApiRequestPublishedGateway) *UpdateApiRequest {
	s.PublishedGateway = v
	return s
}

func (s *UpdateApiRequest) SetStatus(v string) *UpdateApiRequest {
	s.Status = &v
	return s
}

func (s *UpdateApiRequest) SetUpdateDateTime(v string) *UpdateApiRequest {
	s.UpdateDateTime = &v
	return s
}

type UpdateApiRequestAttachedServices struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*UpdateApiRequestAttachedServicesServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s UpdateApiRequestAttachedServices) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiRequestAttachedServices) GoString() string {
	return s.String()
}

func (s *UpdateApiRequestAttachedServices) SetAliasName(v string) *UpdateApiRequestAttachedServices {
	s.AliasName = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetCreationDateTime(v string) *UpdateApiRequestAttachedServices {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetDescription(v string) *UpdateApiRequestAttachedServices {
	s.Description = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetId(v int64) *UpdateApiRequestAttachedServices {
	s.Id = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetIsAutoRefresh(v bool) *UpdateApiRequestAttachedServices {
	s.IsAutoRefresh = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetIsHealth(v bool) *UpdateApiRequestAttachedServices {
	s.IsHealth = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetName(v string) *UpdateApiRequestAttachedServices {
	s.Name = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetRegistryId(v string) *UpdateApiRequestAttachedServices {
	s.RegistryId = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetServiceEnds(v []*UpdateApiRequestAttachedServicesServiceEnds) *UpdateApiRequestAttachedServices {
	s.ServiceEnds = v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetServiceNameInRegistry(v string) *UpdateApiRequestAttachedServices {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetSourceType(v int64) *UpdateApiRequestAttachedServices {
	s.SourceType = &v
	return s
}

func (s *UpdateApiRequestAttachedServices) SetUpdateDateTime(v string) *UpdateApiRequestAttachedServices {
	s.UpdateDateTime = &v
	return s
}

type UpdateApiRequestAttachedServicesServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s UpdateApiRequestAttachedServicesServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiRequestAttachedServicesServiceEnds) GoString() string {
	return s.String()
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetCreationDateTime(v string) *UpdateApiRequestAttachedServicesServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetId(v int64) *UpdateApiRequestAttachedServicesServiceEnds {
	s.Id = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetIpAddress(v string) *UpdateApiRequestAttachedServicesServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetPort(v string) *UpdateApiRequestAttachedServicesServiceEnds {
	s.Port = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetServiceId(v int64) *UpdateApiRequestAttachedServicesServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetStatus(v int64) *UpdateApiRequestAttachedServicesServiceEnds {
	s.Status = &v
	return s
}

func (s *UpdateApiRequestAttachedServicesServiceEnds) SetUpdateDateTime(v string) *UpdateApiRequestAttachedServicesServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type UpdateApiRequestOwneredPolicies struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s UpdateApiRequestOwneredPolicies) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiRequestOwneredPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApiRequestOwneredPolicies) SetApiId(v int64) *UpdateApiRequestOwneredPolicies {
	s.ApiId = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetApiName(v string) *UpdateApiRequestOwneredPolicies {
	s.ApiName = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetCreationDateTime(v string) *UpdateApiRequestOwneredPolicies {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetDirection(v string) *UpdateApiRequestOwneredPolicies {
	s.Direction = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetId(v int64) *UpdateApiRequestOwneredPolicies {
	s.Id = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPolicyAliasName(v string) *UpdateApiRequestOwneredPolicies {
	s.PolicyAliasName = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPolicyContent(v string) *UpdateApiRequestOwneredPolicies {
	s.PolicyContent = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPolicyGroup(v string) *UpdateApiRequestOwneredPolicies {
	s.PolicyGroup = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPolicyId(v string) *UpdateApiRequestOwneredPolicies {
	s.PolicyId = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPolicyName(v string) *UpdateApiRequestOwneredPolicies {
	s.PolicyName = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetPriority(v int64) *UpdateApiRequestOwneredPolicies {
	s.Priority = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetScope(v string) *UpdateApiRequestOwneredPolicies {
	s.Scope = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetStatus(v bool) *UpdateApiRequestOwneredPolicies {
	s.Status = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetType(v int64) *UpdateApiRequestOwneredPolicies {
	s.Type = &v
	return s
}

func (s *UpdateApiRequestOwneredPolicies) SetUpdateDateTime(v string) *UpdateApiRequestOwneredPolicies {
	s.UpdateDateTime = &v
	return s
}

type UpdateApiRequestPublishedGateway struct {
	// armsInfo
	ArmsInfo *string `json:"armsInfo,omitempty" xml:"armsInfo,omitempty"`
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbAccessAddr
	SlbAccessAddr *string `json:"slbAccessAddr,omitempty" xml:"slbAccessAddr,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
}

func (s UpdateApiRequestPublishedGateway) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiRequestPublishedGateway) GoString() string {
	return s.String()
}

func (s *UpdateApiRequestPublishedGateway) SetArmsInfo(v string) *UpdateApiRequestPublishedGateway {
	s.ArmsInfo = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetAutoCreateSlb(v bool) *UpdateApiRequestPublishedGateway {
	s.AutoCreateSlb = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetBasePath(v string) *UpdateApiRequestPublishedGateway {
	s.BasePath = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetCreationDateTime(v string) *UpdateApiRequestPublishedGateway {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetEdasNamespaceId(v string) *UpdateApiRequestPublishedGateway {
	s.EdasNamespaceId = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetGatewayType(v string) *UpdateApiRequestPublishedGateway {
	s.GatewayType = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetId(v int64) *UpdateApiRequestPublishedGateway {
	s.Id = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetName(v string) *UpdateApiRequestPublishedGateway {
	s.Name = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetPodCidr(v string) *UpdateApiRequestPublishedGateway {
	s.PodCidr = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetRegion(v string) *UpdateApiRequestPublishedGateway {
	s.Region = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetRegionName(v string) *UpdateApiRequestPublishedGateway {
	s.RegionName = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetReplica(v int64) *UpdateApiRequestPublishedGateway {
	s.Replica = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetRuntimeOn(v string) *UpdateApiRequestPublishedGateway {
	s.RuntimeOn = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetSecurityGroup(v string) *UpdateApiRequestPublishedGateway {
	s.SecurityGroup = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetSlb(v string) *UpdateApiRequestPublishedGateway {
	s.Slb = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetSlbAccessAddr(v string) *UpdateApiRequestPublishedGateway {
	s.SlbAccessAddr = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetStatus(v string) *UpdateApiRequestPublishedGateway {
	s.Status = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetVpc(v string) *UpdateApiRequestPublishedGateway {
	s.Vpc = &v
	return s
}

func (s *UpdateApiRequestPublishedGateway) SetVswitch(v string) *UpdateApiRequestPublishedGateway {
	s.Vswitch = &v
	return s
}

type UpdateApiResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiResponseBody) SetCode(v int64) *UpdateApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApiResponseBody) SetData(v map[string]interface{}) *UpdateApiResponseBody {
	s.Data = v
	return s
}

func (s *UpdateApiResponseBody) SetMessage(v string) *UpdateApiResponseBody {
	s.Message = &v
	return s
}

type UpdateApiResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiResponse) SetHeaders(v map[string]*string) *UpdateApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiResponse) SetBody(v *UpdateApiResponseBody) *UpdateApiResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// metaInfo
	MetaInfo []*string `json:"metaInfo,omitempty" xml:"metaInfo,omitempty" type:"Repeated"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *int64 `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*string `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetAliasName(v string) *CreateServiceRequest {
	s.AliasName = &v
	return s
}

func (s *CreateServiceRequest) SetDescription(v string) *CreateServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceRequest) SetIsAutoRefresh(v bool) *CreateServiceRequest {
	s.IsAutoRefresh = &v
	return s
}

func (s *CreateServiceRequest) SetMetaInfo(v []*string) *CreateServiceRequest {
	s.MetaInfo = v
	return s
}

func (s *CreateServiceRequest) SetName(v string) *CreateServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceRequest) SetRegistryId(v int64) *CreateServiceRequest {
	s.RegistryId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceEnds(v []*string) *CreateServiceRequest {
	s.ServiceEnds = v
	return s
}

func (s *CreateServiceRequest) SetServiceNameInRegistry(v string) *CreateServiceRequest {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *CreateServiceRequest) SetSourceType(v int64) *CreateServiceRequest {
	s.SourceType = &v
	return s
}

type CreateServiceResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetCode(v int64) *CreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceResponseBody) SetData(v map[string]interface{}) *CreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceResponseBody) SetMessage(v string) *CreateServiceResponseBody {
	s.Message = &v
	return s
}

type CreateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type SaveAllPoliciesRequest struct {
	// data
	Data []*SaveAllPoliciesRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s SaveAllPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAllPoliciesRequest) GoString() string {
	return s.String()
}

func (s *SaveAllPoliciesRequest) SetData(v []*SaveAllPoliciesRequestData) *SaveAllPoliciesRequest {
	s.Data = v
	return s
}

type SaveAllPoliciesRequestData struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s SaveAllPoliciesRequestData) String() string {
	return tea.Prettify(s)
}

func (s SaveAllPoliciesRequestData) GoString() string {
	return s.String()
}

func (s *SaveAllPoliciesRequestData) SetApiId(v int64) *SaveAllPoliciesRequestData {
	s.ApiId = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetApiName(v string) *SaveAllPoliciesRequestData {
	s.ApiName = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetCreationDateTime(v string) *SaveAllPoliciesRequestData {
	s.CreationDateTime = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetDirection(v string) *SaveAllPoliciesRequestData {
	s.Direction = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetId(v int64) *SaveAllPoliciesRequestData {
	s.Id = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPolicyAliasName(v string) *SaveAllPoliciesRequestData {
	s.PolicyAliasName = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPolicyContent(v string) *SaveAllPoliciesRequestData {
	s.PolicyContent = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPolicyGroup(v string) *SaveAllPoliciesRequestData {
	s.PolicyGroup = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPolicyId(v string) *SaveAllPoliciesRequestData {
	s.PolicyId = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPolicyName(v string) *SaveAllPoliciesRequestData {
	s.PolicyName = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetPriority(v int64) *SaveAllPoliciesRequestData {
	s.Priority = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetScope(v string) *SaveAllPoliciesRequestData {
	s.Scope = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetStatus(v bool) *SaveAllPoliciesRequestData {
	s.Status = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetType(v int64) *SaveAllPoliciesRequestData {
	s.Type = &v
	return s
}

func (s *SaveAllPoliciesRequestData) SetUpdateDateTime(v string) *SaveAllPoliciesRequestData {
	s.UpdateDateTime = &v
	return s
}

type SaveAllPoliciesResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s SaveAllPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAllPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAllPoliciesResponseBody) SetCode(v int64) *SaveAllPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *SaveAllPoliciesResponseBody) SetData(v map[string]interface{}) *SaveAllPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *SaveAllPoliciesResponseBody) SetMessage(v string) *SaveAllPoliciesResponseBody {
	s.Message = &v
	return s
}

type SaveAllPoliciesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveAllPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAllPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAllPoliciesResponse) GoString() string {
	return s.String()
}

func (s *SaveAllPoliciesResponse) SetHeaders(v map[string]*string) *SaveAllPoliciesResponse {
	s.Headers = v
	return s
}

func (s *SaveAllPoliciesResponse) SetBody(v *SaveAllPoliciesResponseBody) *SaveAllPoliciesResponse {
	s.Body = v
	return s
}

type UpdateGatewayRequest struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// replica
	Replica *string `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s UpdateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRequest) SetId(v int64) *UpdateGatewayRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRequest) SetReplica(v string) *UpdateGatewayRequest {
	s.Replica = &v
	return s
}

type UpdateGatewayResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponseBody) SetCode(v int64) *UpdateGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetData(v map[string]interface{}) *UpdateGatewayResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayResponseBody) SetMessage(v string) *UpdateGatewayResponseBody {
	s.Message = &v
	return s
}

type UpdateGatewayResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponse) SetHeaders(v map[string]*string) *UpdateGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayResponse) SetBody(v *UpdateGatewayResponseBody) *UpdateGatewayResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *int64 `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*string `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetAliasName(v string) *UpdateServiceRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateServiceRequest) SetCreationDateTime(v string) *UpdateServiceRequest {
	s.CreationDateTime = &v
	return s
}

func (s *UpdateServiceRequest) SetDescription(v string) *UpdateServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateServiceRequest) SetId(v int64) *UpdateServiceRequest {
	s.Id = &v
	return s
}

func (s *UpdateServiceRequest) SetIsAutoRefresh(v bool) *UpdateServiceRequest {
	s.IsAutoRefresh = &v
	return s
}

func (s *UpdateServiceRequest) SetIsHealth(v bool) *UpdateServiceRequest {
	s.IsHealth = &v
	return s
}

func (s *UpdateServiceRequest) SetName(v string) *UpdateServiceRequest {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequest) SetRegistryId(v int64) *UpdateServiceRequest {
	s.RegistryId = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceEnds(v []*string) *UpdateServiceRequest {
	s.ServiceEnds = v
	return s
}

func (s *UpdateServiceRequest) SetServiceNameInRegistry(v string) *UpdateServiceRequest {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *UpdateServiceRequest) SetSourceType(v int64) *UpdateServiceRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateServiceRequest) SetUpdateDateTime(v string) *UpdateServiceRequest {
	s.UpdateDateTime = &v
	return s
}

type UpdateServiceResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetCode(v int64) *UpdateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceResponseBody) SetData(v map[string]interface{}) *UpdateServiceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceResponseBody) SetMessage(v string) *UpdateServiceResponseBody {
	s.Message = &v
	return s
}

type UpdateServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type FindApisByPagingRequest struct {
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
}

func (s FindApisByPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingRequest) GoString() string {
	return s.String()
}

func (s *FindApisByPagingRequest) SetPageNumber(v int64) *FindApisByPagingRequest {
	s.PageNumber = &v
	return s
}

func (s *FindApisByPagingRequest) SetPageSize(v int64) *FindApisByPagingRequest {
	s.PageSize = &v
	return s
}

func (s *FindApisByPagingRequest) SetStatus(v string) *FindApisByPagingRequest {
	s.Status = &v
	return s
}

func (s *FindApisByPagingRequest) SetName(v string) *FindApisByPagingRequest {
	s.Name = &v
	return s
}

func (s *FindApisByPagingRequest) SetAliasName(v string) *FindApisByPagingRequest {
	s.AliasName = &v
	return s
}

type FindApisByPagingResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *FindApisByPagingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindApisByPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBody) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBody) SetCode(v int64) *FindApisByPagingResponseBody {
	s.Code = &v
	return s
}

func (s *FindApisByPagingResponseBody) SetData(v *FindApisByPagingResponseBodyData) *FindApisByPagingResponseBody {
	s.Data = v
	return s
}

func (s *FindApisByPagingResponseBody) SetMessage(v string) *FindApisByPagingResponseBody {
	s.Message = &v
	return s
}

type FindApisByPagingResponseBodyData struct {
	// list
	List []*FindApisByPagingResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s FindApisByPagingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyData) SetList(v []*FindApisByPagingResponseBodyDataList) *FindApisByPagingResponseBodyData {
	s.List = v
	return s
}

func (s *FindApisByPagingResponseBodyData) SetTotalCount(v int64) *FindApisByPagingResponseBodyData {
	s.TotalCount = &v
	return s
}

type FindApisByPagingResponseBodyDataList struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// attachedServices
	AttachedServices []*FindApisByPagingResponseBodyDataListAttachedServices `json:"attachedServices,omitempty" xml:"attachedServices,omitempty" type:"Repeated"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// owneredPolicies
	OwneredPolicies []*FindApisByPagingResponseBodyDataListOwneredPolicies `json:"owneredPolicies,omitempty" xml:"owneredPolicies,omitempty" type:"Repeated"`
	// A short description of struct
	PublishedGateway *FindApisByPagingResponseBodyDataListPublishedGateway `json:"publishedGateway,omitempty" xml:"publishedGateway,omitempty" type:"Struct"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindApisByPagingResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyDataList) SetAliasName(v string) *FindApisByPagingResponseBodyDataList {
	s.AliasName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetAttachedServices(v []*FindApisByPagingResponseBodyDataListAttachedServices) *FindApisByPagingResponseBodyDataList {
	s.AttachedServices = v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetBasePath(v string) *FindApisByPagingResponseBodyDataList {
	s.BasePath = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetCreationDateTime(v string) *FindApisByPagingResponseBodyDataList {
	s.CreationDateTime = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetDescription(v string) *FindApisByPagingResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetId(v int64) *FindApisByPagingResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetName(v string) *FindApisByPagingResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetOwneredPolicies(v []*FindApisByPagingResponseBodyDataListOwneredPolicies) *FindApisByPagingResponseBodyDataList {
	s.OwneredPolicies = v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetPublishedGateway(v *FindApisByPagingResponseBodyDataListPublishedGateway) *FindApisByPagingResponseBodyDataList {
	s.PublishedGateway = v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetStatus(v string) *FindApisByPagingResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataList) SetUpdateDateTime(v string) *FindApisByPagingResponseBodyDataList {
	s.UpdateDateTime = &v
	return s
}

type FindApisByPagingResponseBodyDataListAttachedServices struct {
	// aliasName
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isAutoRefresh
	IsAutoRefresh *bool `json:"isAutoRefresh,omitempty" xml:"isAutoRefresh,omitempty"`
	// isHealth
	IsHealth *bool `json:"isHealth,omitempty" xml:"isHealth,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// registryId
	RegistryId *string `json:"registryId,omitempty" xml:"registryId,omitempty"`
	// serviceEnds
	ServiceEnds []*FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceNameInRegistry
	ServiceNameInRegistry *string `json:"serviceNameInRegistry,omitempty" xml:"serviceNameInRegistry,omitempty"`
	// sourceType
	SourceType *int64 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindApisByPagingResponseBodyDataListAttachedServices) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyDataListAttachedServices) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetAliasName(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.AliasName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetCreationDateTime(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.CreationDateTime = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetDescription(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.Description = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetId(v int64) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.Id = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetIsAutoRefresh(v bool) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.IsAutoRefresh = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetIsHealth(v bool) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.IsHealth = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetName(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.Name = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetRegistryId(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.RegistryId = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetServiceEnds(v []*FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.ServiceEnds = v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetServiceNameInRegistry(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetSourceType(v int64) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.SourceType = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServices) SetUpdateDateTime(v string) *FindApisByPagingResponseBodyDataListAttachedServices {
	s.UpdateDateTime = &v
	return s
}

type FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds struct {
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// serviceId
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetCreationDateTime(v string) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.CreationDateTime = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetId(v int64) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.Id = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetIpAddress(v string) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.IpAddress = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetPort(v string) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.Port = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetServiceId(v int64) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.ServiceId = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetStatus(v int64) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.Status = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds) SetUpdateDateTime(v string) *FindApisByPagingResponseBodyDataListAttachedServicesServiceEnds {
	s.UpdateDateTime = &v
	return s
}

type FindApisByPagingResponseBodyDataListOwneredPolicies struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s FindApisByPagingResponseBodyDataListOwneredPolicies) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyDataListOwneredPolicies) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetApiId(v int64) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.ApiId = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetApiName(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.ApiName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetCreationDateTime(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.CreationDateTime = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetDirection(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Direction = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetId(v int64) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Id = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPolicyAliasName(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.PolicyAliasName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPolicyContent(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.PolicyContent = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPolicyGroup(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.PolicyGroup = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPolicyId(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.PolicyId = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPolicyName(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.PolicyName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetPriority(v int64) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Priority = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetScope(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Scope = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetStatus(v bool) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Status = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetType(v int64) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.Type = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListOwneredPolicies) SetUpdateDateTime(v string) *FindApisByPagingResponseBodyDataListOwneredPolicies {
	s.UpdateDateTime = &v
	return s
}

type FindApisByPagingResponseBodyDataListPublishedGateway struct {
	// armsInfo
	ArmsInfo *string `json:"armsInfo,omitempty" xml:"armsInfo,omitempty"`
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbAccessAddr
	SlbAccessAddr *string `json:"slbAccessAddr,omitempty" xml:"slbAccessAddr,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
}

func (s FindApisByPagingResponseBodyDataListPublishedGateway) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponseBodyDataListPublishedGateway) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetArmsInfo(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.ArmsInfo = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetAutoCreateSlb(v bool) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.AutoCreateSlb = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetBasePath(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.BasePath = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetCreationDateTime(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.CreationDateTime = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetEdasNamespaceId(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.EdasNamespaceId = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetGatewayType(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.GatewayType = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetId(v int64) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Id = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetName(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Name = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetPodCidr(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.PodCidr = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetRegion(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Region = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetRegionName(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.RegionName = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetReplica(v int64) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Replica = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetRuntimeOn(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.RuntimeOn = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetSecurityGroup(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.SecurityGroup = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetSlb(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Slb = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetSlbAccessAddr(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.SlbAccessAddr = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetStatus(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Status = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetVpc(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Vpc = &v
	return s
}

func (s *FindApisByPagingResponseBodyDataListPublishedGateway) SetVswitch(v string) *FindApisByPagingResponseBodyDataListPublishedGateway {
	s.Vswitch = &v
	return s
}

type FindApisByPagingResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindApisByPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindApisByPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s FindApisByPagingResponse) GoString() string {
	return s.String()
}

func (s *FindApisByPagingResponse) SetHeaders(v map[string]*string) *FindApisByPagingResponse {
	s.Headers = v
	return s
}

func (s *FindApisByPagingResponse) SetBody(v *FindApisByPagingResponseBody) *FindApisByPagingResponse {
	s.Body = v
	return s
}

type UpdateServiceEndsRequest struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// serviceNodes
	ServiceNodes []*UpdateServiceEndsRequestServiceNodes `json:"serviceNodes,omitempty" xml:"serviceNodes,omitempty" type:"Repeated"`
}

func (s UpdateServiceEndsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEndsRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndsRequest) SetId(v int64) *UpdateServiceEndsRequest {
	s.Id = &v
	return s
}

func (s *UpdateServiceEndsRequest) SetServiceNodes(v []*UpdateServiceEndsRequestServiceNodes) *UpdateServiceEndsRequest {
	s.ServiceNodes = v
	return s
}

type UpdateServiceEndsRequestServiceNodes struct {
	// port
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// ipAddress
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// status
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateServiceEndsRequestServiceNodes) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEndsRequestServiceNodes) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndsRequestServiceNodes) SetPort(v string) *UpdateServiceEndsRequestServiceNodes {
	s.Port = &v
	return s
}

func (s *UpdateServiceEndsRequestServiceNodes) SetIpAddress(v string) *UpdateServiceEndsRequestServiceNodes {
	s.IpAddress = &v
	return s
}

func (s *UpdateServiceEndsRequestServiceNodes) SetStatus(v int64) *UpdateServiceEndsRequestServiceNodes {
	s.Status = &v
	return s
}

type UpdateServiceEndsResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateServiceEndsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEndsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndsResponseBody) SetCode(v int64) *UpdateServiceEndsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceEndsResponseBody) SetData(v map[string]interface{}) *UpdateServiceEndsResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceEndsResponseBody) SetMessage(v string) *UpdateServiceEndsResponseBody {
	s.Message = &v
	return s
}

type UpdateServiceEndsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceEndsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceEndsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceEndsResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndsResponse) SetHeaders(v map[string]*string) *UpdateServiceEndsResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceEndsResponse) SetBody(v *UpdateServiceEndsResponseBody) *UpdateServiceEndsResponse {
	s.Body = v
	return s
}

type FindGatewaysRequest struct {
	// gatewayUniqueId
	GatewayUniqueId *string `json:"gatewayUniqueId,omitempty" xml:"gatewayUniqueId,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// gatewayTypes
	GatewayTypes *string `json:"gatewayTypes,omitempty" xml:"gatewayTypes,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// pageNumber
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s FindGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysRequest) GoString() string {
	return s.String()
}

func (s *FindGatewaysRequest) SetGatewayUniqueId(v string) *FindGatewaysRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *FindGatewaysRequest) SetName(v string) *FindGatewaysRequest {
	s.Name = &v
	return s
}

func (s *FindGatewaysRequest) SetRegion(v string) *FindGatewaysRequest {
	s.Region = &v
	return s
}

func (s *FindGatewaysRequest) SetGatewayTypes(v string) *FindGatewaysRequest {
	s.GatewayTypes = &v
	return s
}

func (s *FindGatewaysRequest) SetStatus(v string) *FindGatewaysRequest {
	s.Status = &v
	return s
}

func (s *FindGatewaysRequest) SetPageNumber(v string) *FindGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *FindGatewaysRequest) SetPageSize(v string) *FindGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *FindGatewaysRequest) SetNamespace(v string) *FindGatewaysRequest {
	s.Namespace = &v
	return s
}

type FindGatewaysResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *FindGatewaysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *FindGatewaysResponseBody) SetCode(v int64) *FindGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *FindGatewaysResponseBody) SetData(v *FindGatewaysResponseBodyData) *FindGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *FindGatewaysResponseBody) SetMessage(v string) *FindGatewaysResponseBody {
	s.Message = &v
	return s
}

type FindGatewaysResponseBodyData struct {
	// list
	List []*FindGatewaysResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s FindGatewaysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindGatewaysResponseBodyData) SetList(v []*FindGatewaysResponseBodyDataList) *FindGatewaysResponseBodyData {
	s.List = v
	return s
}

func (s *FindGatewaysResponseBodyData) SetTotalCount(v int64) *FindGatewaysResponseBodyData {
	s.TotalCount = &v
	return s
}

type FindGatewaysResponseBodyDataList struct {
	// armsInfo
	ArmsInfo *FindGatewaysResponseBodyDataListArmsInfo `json:"armsInfo,omitempty" xml:"armsInfo,omitempty" type:"Struct"`
	// autoCreateSlb
	AutoCreateSlb *bool `json:"autoCreateSlb,omitempty" xml:"autoCreateSlb,omitempty"`
	// basePath
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// edasNamespaceId
	EdasNamespaceId *string `json:"edasNamespaceId,omitempty" xml:"edasNamespaceId,omitempty"`
	// gatewayType
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// podCidr
	PodCidr *string `json:"podCidr,omitempty" xml:"podCidr,omitempty"`
	// region
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// regionName
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// replica
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// runtimeOn
	RuntimeOn *string `json:"runtimeOn,omitempty" xml:"runtimeOn,omitempty"`
	// securityGroup
	SecurityGroup *string `json:"securityGroup,omitempty" xml:"securityGroup,omitempty"`
	// slb
	Slb *string `json:"slb,omitempty" xml:"slb,omitempty"`
	// slbAccessAddr
	SlbAccessAddr *string `json:"slbAccessAddr,omitempty" xml:"slbAccessAddr,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// vpc
	Vpc *string `json:"vpc,omitempty" xml:"vpc,omitempty"`
	// vswitch
	Vswitch *string `json:"vswitch,omitempty" xml:"vswitch,omitempty"`
}

func (s FindGatewaysResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *FindGatewaysResponseBodyDataList) SetArmsInfo(v *FindGatewaysResponseBodyDataListArmsInfo) *FindGatewaysResponseBodyDataList {
	s.ArmsInfo = v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetAutoCreateSlb(v bool) *FindGatewaysResponseBodyDataList {
	s.AutoCreateSlb = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetBasePath(v string) *FindGatewaysResponseBodyDataList {
	s.BasePath = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetCreationDateTime(v string) *FindGatewaysResponseBodyDataList {
	s.CreationDateTime = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetEdasNamespaceId(v string) *FindGatewaysResponseBodyDataList {
	s.EdasNamespaceId = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetGatewayType(v string) *FindGatewaysResponseBodyDataList {
	s.GatewayType = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetId(v int64) *FindGatewaysResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetName(v string) *FindGatewaysResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetPodCidr(v string) *FindGatewaysResponseBodyDataList {
	s.PodCidr = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetRegion(v string) *FindGatewaysResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetRegionName(v string) *FindGatewaysResponseBodyDataList {
	s.RegionName = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetReplica(v int64) *FindGatewaysResponseBodyDataList {
	s.Replica = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetRuntimeOn(v string) *FindGatewaysResponseBodyDataList {
	s.RuntimeOn = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetSecurityGroup(v string) *FindGatewaysResponseBodyDataList {
	s.SecurityGroup = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetSlb(v string) *FindGatewaysResponseBodyDataList {
	s.Slb = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetSlbAccessAddr(v string) *FindGatewaysResponseBodyDataList {
	s.SlbAccessAddr = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetStatus(v string) *FindGatewaysResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetVpc(v string) *FindGatewaysResponseBodyDataList {
	s.Vpc = &v
	return s
}

func (s *FindGatewaysResponseBodyDataList) SetVswitch(v string) *FindGatewaysResponseBodyDataList {
	s.Vswitch = &v
	return s
}

type FindGatewaysResponseBodyDataListArmsInfo struct {
	// appId
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// licenseKey
	LicenseKey *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
}

func (s FindGatewaysResponseBodyDataListArmsInfo) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysResponseBodyDataListArmsInfo) GoString() string {
	return s.String()
}

func (s *FindGatewaysResponseBodyDataListArmsInfo) SetAppId(v string) *FindGatewaysResponseBodyDataListArmsInfo {
	s.AppId = &v
	return s
}

func (s *FindGatewaysResponseBodyDataListArmsInfo) SetAppName(v string) *FindGatewaysResponseBodyDataListArmsInfo {
	s.AppName = &v
	return s
}

func (s *FindGatewaysResponseBodyDataListArmsInfo) SetDescription(v string) *FindGatewaysResponseBodyDataListArmsInfo {
	s.Description = &v
	return s
}

func (s *FindGatewaysResponseBodyDataListArmsInfo) SetLicenseKey(v string) *FindGatewaysResponseBodyDataListArmsInfo {
	s.LicenseKey = &v
	return s
}

type FindGatewaysResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FindGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FindGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s FindGatewaysResponse) GoString() string {
	return s.String()
}

func (s *FindGatewaysResponse) SetHeaders(v map[string]*string) *FindGatewaysResponse {
	s.Headers = v
	return s
}

func (s *FindGatewaysResponse) SetBody(v *FindGatewaysResponseBody) *FindGatewaysResponse {
	s.Body = v
	return s
}

type GetAllRegistryRequest struct {
	// pageNumber
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// groupBy
	GroupBy *bool `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetAllRegistryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegistryRequest) GoString() string {
	return s.String()
}

func (s *GetAllRegistryRequest) SetPageNumber(v int64) *GetAllRegistryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAllRegistryRequest) SetPageSize(v int64) *GetAllRegistryRequest {
	s.PageSize = &v
	return s
}

func (s *GetAllRegistryRequest) SetName(v string) *GetAllRegistryRequest {
	s.Name = &v
	return s
}

func (s *GetAllRegistryRequest) SetType(v int64) *GetAllRegistryRequest {
	s.Type = &v
	return s
}

func (s *GetAllRegistryRequest) SetGroupBy(v bool) *GetAllRegistryRequest {
	s.GroupBy = &v
	return s
}

type GetAllRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *GetAllRegistryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAllRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllRegistryResponseBody) SetCode(v int64) *GetAllRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllRegistryResponseBody) SetData(v *GetAllRegistryResponseBodyData) *GetAllRegistryResponseBody {
	s.Data = v
	return s
}

func (s *GetAllRegistryResponseBody) SetMessage(v string) *GetAllRegistryResponseBody {
	s.Message = &v
	return s
}

type GetAllRegistryResponseBodyData struct {
	// list
	List []*GetAllRegistryResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// totalCount
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetAllRegistryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegistryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllRegistryResponseBodyData) SetList(v []*GetAllRegistryResponseBodyDataList) *GetAllRegistryResponseBodyData {
	s.List = v
	return s
}

func (s *GetAllRegistryResponseBodyData) SetTotalCount(v int64) *GetAllRegistryResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetAllRegistryResponseBodyDataList struct {
	// address
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// gatewayId
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s GetAllRegistryResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegistryResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAllRegistryResponseBodyDataList) SetAddress(v string) *GetAllRegistryResponseBodyDataList {
	s.Address = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetCreationDateTime(v string) *GetAllRegistryResponseBodyDataList {
	s.CreationDateTime = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetDescription(v string) *GetAllRegistryResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetGatewayId(v string) *GetAllRegistryResponseBodyDataList {
	s.GatewayId = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetId(v int64) *GetAllRegistryResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetName(v string) *GetAllRegistryResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetType(v int64) *GetAllRegistryResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *GetAllRegistryResponseBodyDataList) SetUpdateDateTime(v string) *GetAllRegistryResponseBodyDataList {
	s.UpdateDateTime = &v
	return s
}

type GetAllRegistryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllRegistryResponse) GoString() string {
	return s.String()
}

func (s *GetAllRegistryResponse) SetHeaders(v map[string]*string) *GetAllRegistryResponse {
	s.Headers = v
	return s
}

func (s *GetAllRegistryResponse) SetBody(v *GetAllRegistryResponseBody) *GetAllRegistryResponse {
	s.Body = v
	return s
}

type PullServiceInfoFromRegistryResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*PullServiceInfoFromRegistryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s PullServiceInfoFromRegistryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullServiceInfoFromRegistryResponseBody) GoString() string {
	return s.String()
}

func (s *PullServiceInfoFromRegistryResponseBody) SetCode(v int64) *PullServiceInfoFromRegistryResponseBody {
	s.Code = &v
	return s
}

func (s *PullServiceInfoFromRegistryResponseBody) SetData(v []*PullServiceInfoFromRegistryResponseBodyData) *PullServiceInfoFromRegistryResponseBody {
	s.Data = v
	return s
}

func (s *PullServiceInfoFromRegistryResponseBody) SetMessage(v string) *PullServiceInfoFromRegistryResponseBody {
	s.Message = &v
	return s
}

type PullServiceInfoFromRegistryResponseBodyData struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// metaInfo
	MetaInfo    *string   `json:"metaInfo,omitempty" xml:"metaInfo,omitempty"`
	ServiceEnds []*string `json:"serviceEnds,omitempty" xml:"serviceEnds,omitempty" type:"Repeated"`
	// serviceName
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s PullServiceInfoFromRegistryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PullServiceInfoFromRegistryResponseBodyData) GoString() string {
	return s.String()
}

func (s *PullServiceInfoFromRegistryResponseBodyData) SetId(v int64) *PullServiceInfoFromRegistryResponseBodyData {
	s.Id = &v
	return s
}

func (s *PullServiceInfoFromRegistryResponseBodyData) SetMetaInfo(v string) *PullServiceInfoFromRegistryResponseBodyData {
	s.MetaInfo = &v
	return s
}

func (s *PullServiceInfoFromRegistryResponseBodyData) SetServiceEnds(v []*string) *PullServiceInfoFromRegistryResponseBodyData {
	s.ServiceEnds = v
	return s
}

func (s *PullServiceInfoFromRegistryResponseBodyData) SetServiceName(v string) *PullServiceInfoFromRegistryResponseBodyData {
	s.ServiceName = &v
	return s
}

type PullServiceInfoFromRegistryResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullServiceInfoFromRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullServiceInfoFromRegistryResponse) String() string {
	return tea.Prettify(s)
}

func (s PullServiceInfoFromRegistryResponse) GoString() string {
	return s.String()
}

func (s *PullServiceInfoFromRegistryResponse) SetHeaders(v map[string]*string) *PullServiceInfoFromRegistryResponse {
	s.Headers = v
	return s
}

func (s *PullServiceInfoFromRegistryResponse) SetBody(v *PullServiceInfoFromRegistryResponseBody) *PullServiceInfoFromRegistryResponse {
	s.Body = v
	return s
}

type DataValue struct {
	// apiId
	ApiId *int64 `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// apiName
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// creationDateTime
	CreationDateTime *string `json:"creationDateTime,omitempty" xml:"creationDateTime,omitempty"`
	// direction
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// policyAliasName
	PolicyAliasName *string `json:"policyAliasName,omitempty" xml:"policyAliasName,omitempty"`
	// policyContent
	PolicyContent *string `json:"policyContent,omitempty" xml:"policyContent,omitempty"`
	// policyGroup
	PolicyGroup *string `json:"policyGroup,omitempty" xml:"policyGroup,omitempty"`
	// policyId
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// policyName
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// priority
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// status
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
	// type
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// updateDateTime
	UpdateDateTime *string `json:"updateDateTime,omitempty" xml:"updateDateTime,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetApiId(v int64) *DataValue {
	s.ApiId = &v
	return s
}

func (s *DataValue) SetApiName(v string) *DataValue {
	s.ApiName = &v
	return s
}

func (s *DataValue) SetCreationDateTime(v string) *DataValue {
	s.CreationDateTime = &v
	return s
}

func (s *DataValue) SetDirection(v string) *DataValue {
	s.Direction = &v
	return s
}

func (s *DataValue) SetId(v int64) *DataValue {
	s.Id = &v
	return s
}

func (s *DataValue) SetPolicyAliasName(v string) *DataValue {
	s.PolicyAliasName = &v
	return s
}

func (s *DataValue) SetPolicyContent(v string) *DataValue {
	s.PolicyContent = &v
	return s
}

func (s *DataValue) SetPolicyGroup(v string) *DataValue {
	s.PolicyGroup = &v
	return s
}

func (s *DataValue) SetPolicyId(v string) *DataValue {
	s.PolicyId = &v
	return s
}

func (s *DataValue) SetPolicyName(v string) *DataValue {
	s.PolicyName = &v
	return s
}

func (s *DataValue) SetPriority(v int64) *DataValue {
	s.Priority = &v
	return s
}

func (s *DataValue) SetScope(v string) *DataValue {
	s.Scope = &v
	return s
}

func (s *DataValue) SetStatus(v bool) *DataValue {
	s.Status = &v
	return s
}

func (s *DataValue) SetType(v int64) *DataValue {
	s.Type = &v
	return s
}

func (s *DataValue) SetUpdateDateTime(v string) *DataValue {
	s.UpdateDateTime = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("microgw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * findAllService
 */
func (client *Client) FindAllService(gatewayId *string, request *FindAllServiceRequest) (_result *FindAllServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindAllServiceResponse{}
	_body, _err := client.FindAllServiceWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindAllServiceWithOptions(gatewayId *string, request *FindAllServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindAllServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsHealth)) {
		query["isHealth"] = request.IsHealth
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FindAllServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("FindAllService"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/service"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createApi
 */
func (client *Client) CreateApi(gatewayId *string, request *CreateApiRequest) (_result *CreateApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApiResponse{}
	_body, _err := client.CreateApiWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApiWithOptions(gatewayId *string, request *CreateApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.AttachedServices)) {
		body["attachedServices"] = request.AttachedServices
	}

	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateApiResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/api"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getGatewayById
 */
func (client *Client) GetGatewayById(gatewayId *string) (_result *GetGatewayByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayByIdResponse{}
	_body, _err := client.GetGatewayByIdWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayByIdWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGatewayByIdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetGatewayByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetGatewayById"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createPolicy
 */
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

func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroup)) {
		body["policyGroup"] = request.PolicyGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatePolicy"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/policy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getServiceInstanceForRegistryByServiceName
 */
func (client *Client) GetServiceInstanceForRegistryByServiceName(gatewayId *string, registryId *string, request *GetServiceInstanceForRegistryByServiceNameRequest) (_result *GetServiceInstanceForRegistryByServiceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceInstanceForRegistryByServiceNameResponse{}
	_body, _err := client.GetServiceInstanceForRegistryByServiceNameWithOptions(gatewayId, registryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceInstanceForRegistryByServiceNameWithOptions(gatewayId *string, registryId *string, request *GetServiceInstanceForRegistryByServiceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceInstanceForRegistryByServiceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["serviceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetServiceInstanceForRegistryByServiceNameResponse{}
	_body, _err := client.DoROARequest(tea.String("GetServiceInstanceForRegistryByServiceName"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/registry/{registryId}/service"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * deleteService
 */
func (client *Client) DeleteService(serviceId *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(serviceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteService"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/service/"+tea.StringValue(serviceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * UpdateRegistry
 */
func (client *Client) UpdateRegistry(registryId *string, request *UpdateRegistryRequest) (_result *UpdateRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRegistryResponse{}
	_body, _err := client.UpdateRegistryWithOptions(registryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRegistryWithOptions(registryId *string, request *UpdateRegistryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRegistryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/registry/"+tea.StringValue(registryId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createGateway
 */
func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateSlb)) {
		body["autoCreateSlb"] = request.AutoCreateSlb
	}

	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.EdasNamespaceId)) {
		body["edasNamespaceId"] = request.EdasNamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayType)) {
		body["gatewayType"] = request.GatewayType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PodCidr)) {
		body["podCidr"] = request.PodCidr
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionName)) {
		body["regionName"] = request.RegionName
	}

	if !tea.BoolValue(util.IsUnset(request.Replica)) {
		body["replica"] = request.Replica
	}

	if !tea.BoolValue(util.IsUnset(request.RuntimeOn)) {
		body["runtimeOn"] = request.RuntimeOn
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroup)) {
		body["securityGroup"] = request.SecurityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Slb)) {
		body["slb"] = request.Slb
	}

	if !tea.BoolValue(util.IsUnset(request.SlbSpec)) {
		body["slbSpec"] = request.SlbSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Vpc)) {
		body["vpc"] = request.Vpc
	}

	if !tea.BoolValue(util.IsUnset(request.Vswitch)) {
		body["vswitch"] = request.Vswitch
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["zone"] = request.Zone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGateway"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/gateway"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * checkServiceHealth
 */
func (client *Client) CheckServiceHealth(request *CheckServiceHealthRequest) (_result *CheckServiceHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckServiceHealthResponse{}
	_body, _err := client.CheckServiceHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceHealthWithOptions(request *CheckServiceHealthRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckServiceHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OperationIds)) {
		body["operationIds"] = request.OperationIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CheckServiceHealthResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckServiceHealth"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/service/check"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createPolicyToApi
 */
func (client *Client) CreatePolicyToApi(apiId *string, request *CreatePolicyToApiRequest) (_result *CreatePolicyToApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePolicyToApiResponse{}
	_body, _err := client.CreatePolicyToApiWithOptions(apiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolicyToApiWithOptions(apiId *string, request *CreatePolicyToApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePolicyToApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreationDateTime)) {
		body["creationDateTime"] = request.CreationDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		body["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyAliasName)) {
		body["policyAliasName"] = request.PolicyAliasName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyContent)) {
		body["policyContent"] = request.PolicyContent
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroup)) {
		body["policyGroup"] = request.PolicyGroup
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		body["policyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		body["policyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreatePolicyToApiResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatePolicyToApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/policy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * detachPolicy
 */
func (client *Client) DetachPolicy(apiId *string, policyId *string) (_result *DetachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DetachPolicyResponse{}
	_body, _err := client.DetachPolicyWithOptions(apiId, policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachPolicyWithOptions(apiId *string, policyId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DetachPolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("DetachPolicy"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/detach/{policyId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * findTemplate
 */
func (client *Client) FindTemplate(apiId *string) (_result *FindTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindTemplateResponse{}
	_body, _err := client.FindTemplateWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindTemplateWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &FindTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("FindTemplate"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/policy/template"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * validateRegistryAddress
 */
func (client *Client) ValidateRegistryAddress(gatewayId *string, request *ValidateRegistryAddressRequest) (_result *ValidateRegistryAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateRegistryAddressResponse{}
	_body, _err := client.ValidateRegistryAddressWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateRegistryAddressWithOptions(gatewayId *string, request *ValidateRegistryAddressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateRegistryAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ValidateRegistryAddressResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateRegistryAddress"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/registry/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getApiDetail
 */
func (client *Client) GetApiDetail(apiId *string) (_result *GetApiDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApiDetailResponse{}
	_body, _err := client.GetApiDetailWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApiDetailWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetApiDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetApiDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApiDetail"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createSpecialRouteForRegistry
 */
func (client *Client) CreateSpecialRouteForRegistry(gatewayId *string) (_result *CreateSpecialRouteForRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSpecialRouteForRegistryResponse{}
	_body, _err := client.CreateSpecialRouteForRegistryWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSpecialRouteForRegistryWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSpecialRouteForRegistryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateSpecialRouteForRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSpecialRouteForRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/registry/special/route"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * publishApi
 */
func (client *Client) PublishApi(apiId *string) (_result *PublishApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishApiResponse{}
	_body, _err := client.PublishApiWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishApiWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &PublishApiResponse{}
	_body, _err := client.DoROARequest(tea.String("PublishApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/publish"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * CreateGatewayLogEtl
 */
func (client *Client) CreateGatewayLogEtl(gatewayId *string, regionId *string) (_result *CreateGatewayLogEtlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayLogEtlResponse{}
	_body, _err := client.CreateGatewayLogEtlWithOptions(gatewayId, regionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayLogEtlWithOptions(gatewayId *string, regionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayLogEtlResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateGatewayLogEtlResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateGatewayLogEtl"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/sls/gateway/"+tea.StringValue(gatewayId)+"/region/{regionId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * FindPolicies
 */
func (client *Client) FindPolicies(gatewayId *string, request *FindPoliciesRequest) (_result *FindPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindPoliciesResponse{}
	_body, _err := client.FindPoliciesWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindPoliciesWithOptions(gatewayId *string, request *FindPoliciesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["group"] = request.Group
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FindPoliciesResponse{}
	_body, _err := client.DoROARequest(tea.String("FindPolicies"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/policy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * attachPolicy
 */
func (client *Client) AttachPolicy(apiId *string, request *AttachPolicyRequest) (_result *AttachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachPolicyResponse{}
	_body, _err := client.AttachPolicyWithOptions(apiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachPolicyWithOptions(apiId *string, request *AttachPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Data),
	}
	_result = &AttachPolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("AttachPolicy"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/attach"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * findRegistry
 */
func (client *Client) FindRegistry(registryId *string) (_result *FindRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindRegistryResponse{}
	_body, _err := client.FindRegistryWithOptions(registryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindRegistryWithOptions(registryId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindRegistryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &FindRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("FindRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/registry/"+tea.StringValue(registryId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getAuthTicketById
 */
func (client *Client) GetAuthTicketById(ticketId *string) (_result *GetAuthTicketByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAuthTicketByIdHeaders{}
	_result = &GetAuthTicketByIdResponse{}
	_body, _err := client.GetAuthTicketByIdWithOptions(ticketId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthTicketByIdWithOptions(ticketId *string, tmpHeader *GetAuthTicketByIdHeaders, runtime *util.RuntimeOptions) (_result *GetAuthTicketByIdResponse, _err error) {
	headers := &GetAuthTicketByIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.Cookie)) {
		headers.CookieShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.Cookie, tea.String("cookie"), tea.String("json"))
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CookieShrink)) {
		realHeaders["cookie"] = headers.CookieShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetAuthTicketByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAuthTicketById"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/auth/"+tea.StringValue(ticketId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * CreateRegistry
 */
func (client *Client) CreateRegistry(gatewayId *string, request *CreateRegistryRequest) (_result *CreateRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRegistryResponse{}
	_body, _err := client.CreateRegistryWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRegistryWithOptions(gatewayId *string, request *CreateRegistryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRegistryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/registry"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * recycleApi
 */
func (client *Client) RecycleApi(apiId *string) (_result *RecycleApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecycleApiResponse{}
	_body, _err := client.RecycleApiWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecycleApiWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecycleApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &RecycleApiResponse{}
	_body, _err := client.DoROARequest(tea.String("RecycleApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/recycle"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createAuthTicket
 */
func (client *Client) CreateAuthTicket(request *CreateAuthTicketRequest) (_result *CreateAuthTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAuthTicketResponse{}
	_body, _err := client.CreateAuthTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAuthTicketWithOptions(request *CreateAuthTicketRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAuthTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		body["gatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TicketType)) {
		body["ticketType"] = request.TicketType
	}

	if !tea.BoolValue(util.IsUnset(request.ValidDuration)) {
		body["validDuration"] = request.ValidDuration
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateAuthTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateAuthTicket"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/auth"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * deleteGateway
 */
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

func (client *Client) DeleteGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteGateway"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * findService
 */
func (client *Client) FindService(serviceId *string) (_result *FindServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindServiceResponse{}
	_body, _err := client.FindServiceWithOptions(serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindServiceWithOptions(serviceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &FindServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("FindService"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/service/"+tea.StringValue(serviceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * DeletePolicyById
 */
func (client *Client) DeletePolicyById(policyId *string) (_result *DeletePolicyByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyByIdResponse{}
	_body, _err := client.DeletePolicyByIdWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePolicyByIdWithOptions(policyId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePolicyByIdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeletePolicyByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePolicyById"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/policy/"+tea.StringValue(policyId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * deleteApi
 */
func (client *Client) DeleteApi(apiId *string) (_result *DeleteApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApiResponse{}
	_body, _err := client.DeleteApiWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApiWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteApiResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * findAuthTickets
 */
func (client *Client) FindAuthTickets(request *FindAuthTicketsRequest) (_result *FindAuthTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindAuthTicketsResponse{}
	_body, _err := client.FindAuthTicketsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindAuthTicketsWithOptions(request *FindAuthTicketsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindAuthTicketsResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FindAuthTicketsResponse{}
	_body, _err := client.DoROARequest(tea.String("FindAuthTickets"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/auth"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * updatePolicy
 */
func (client *Client) UpdatePolicy(request *UpdatePolicyRequest) (_result *UpdatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePolicyWithOptions(request *UpdatePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyGroup)) {
		body["policyGroup"] = request.PolicyGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePolicy"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/policy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * updateAuthTicket
 */
func (client *Client) UpdateAuthTicket(request *UpdateAuthTicketRequest) (_result *UpdateAuthTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAuthTicketResponse{}
	_body, _err := client.UpdateAuthTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAuthTicketWithOptions(request *UpdateAuthTicketRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAuthTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateAuthTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateAuthTicket"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/auth"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * installArmsAgent
 */
func (client *Client) InstallArmsAgent(gatewayId *string) (_result *InstallArmsAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallArmsAgentResponse{}
	_body, _err := client.InstallArmsAgentWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallArmsAgentWithOptions(gatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallArmsAgentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &InstallArmsAgentResponse{}
	_body, _err := client.DoROARequest(tea.String("InstallArmsAgent"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/agent/"+tea.StringValue(gatewayId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * deleteAuthTicket
 */
func (client *Client) DeleteAuthTicket(ticketId *string) (_result *DeleteAuthTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAuthTicketResponse{}
	_body, _err := client.DeleteAuthTicketWithOptions(ticketId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAuthTicketWithOptions(ticketId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAuthTicketResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteAuthTicketResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteAuthTicket"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/auth/"+tea.StringValue(ticketId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * GetPolicyById
 */
func (client *Client) GetPolicyById(policyId *string) (_result *GetPolicyByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyByIdResponse{}
	_body, _err := client.GetPolicyByIdWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPolicyByIdWithOptions(policyId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPolicyByIdResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetPolicyByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPolicyById"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/policy/"+tea.StringValue(policyId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * deleteRegistry
 */
func (client *Client) DeleteRegistry(registryId *string) (_result *DeleteRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRegistryResponse{}
	_body, _err := client.DeleteRegistryWithOptions(registryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRegistryWithOptions(registryId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRegistryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/registry/"+tea.StringValue(registryId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getPolicyOwnedByApi
 */
func (client *Client) GetPolicyOwnedByApi(apiId *string) (_result *GetPolicyOwnedByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyOwnedByApiResponse{}
	_body, _err := client.GetPolicyOwnedByApiWithOptions(apiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPolicyOwnedByApiWithOptions(apiId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPolicyOwnedByApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetPolicyOwnedByApiResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPolicyOwnedByApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/policy"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * updateApi
 */
func (client *Client) UpdateApi(apiId *string, request *UpdateApiRequest) (_result *UpdateApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApiResponse{}
	_body, _err := client.UpdateApiWithOptions(apiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApiWithOptions(apiId *string, request *UpdateApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.AttachedServices)) {
		body["attachedServices"] = request.AttachedServices
	}

	if !tea.BoolValue(util.IsUnset(request.BasePath)) {
		body["basePath"] = request.BasePath
	}

	if !tea.BoolValue(util.IsUnset(request.CreationDateTime)) {
		body["creationDateTime"] = request.CreationDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwneredPolicies)) {
		body["owneredPolicies"] = request.OwneredPolicies
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PublishedGateway))) {
		body["publishedGateway"] = request.PublishedGateway
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateDateTime)) {
		body["updateDateTime"] = request.UpdateDateTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateApiResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateApi"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * /gateway/{gatewayId}/service
 */
func (client *Client) CreateService(gatewayId *string, request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(gatewayId *string, request *CreateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRefresh)) {
		body["isAutoRefresh"] = request.IsAutoRefresh
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		body["metaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegistryId)) {
		body["registryId"] = request.RegistryId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEnds)) {
		body["serviceEnds"] = request.ServiceEnds
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNameInRegistry)) {
		body["serviceNameInRegistry"] = request.ServiceNameInRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateService"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/service"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * saveAllPolicies
 */
func (client *Client) SaveAllPolicies(apiId *string, request *SaveAllPoliciesRequest) (_result *SaveAllPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveAllPoliciesResponse{}
	_body, _err := client.SaveAllPoliciesWithOptions(apiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAllPoliciesWithOptions(apiId *string, request *SaveAllPoliciesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveAllPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Data),
	}
	_result = &SaveAllPoliciesResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveAllPolicies"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/api/"+tea.StringValue(apiId)+"/policies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * updateGateway
 */
func (client *Client) UpdateGateway(request *UpdateGatewayRequest) (_result *UpdateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayResponse{}
	_body, _err := client.UpdateGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayWithOptions(request *UpdateGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Replica)) {
		body["replica"] = request.Replica
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGateway"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/gateway"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * updateService
 */
func (client *Client) UpdateService(serviceId *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["aliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.CreationDateTime)) {
		body["creationDateTime"] = request.CreationDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRefresh)) {
		body["isAutoRefresh"] = request.IsAutoRefresh
	}

	if !tea.BoolValue(util.IsUnset(request.IsHealth)) {
		body["isHealth"] = request.IsHealth
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegistryId)) {
		body["registryId"] = request.RegistryId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceEnds)) {
		body["serviceEnds"] = request.ServiceEnds
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNameInRegistry)) {
		body["serviceNameInRegistry"] = request.ServiceNameInRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateDateTime)) {
		body["updateDateTime"] = request.UpdateDateTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateService"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/service/"+tea.StringValue(serviceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * findApisByPaging
 */
func (client *Client) FindApisByPaging(gatewayId *string, request *FindApisByPagingRequest) (_result *FindApisByPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindApisByPagingResponse{}
	_body, _err := client.FindApisByPagingWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindApisByPagingWithOptions(gatewayId *string, request *FindApisByPagingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindApisByPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		query["aliasName"] = request.AliasName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FindApisByPagingResponse{}
	_body, _err := client.DoROARequest(tea.String("FindApisByPaging"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/api"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceEnds(serviceId *string, request *UpdateServiceEndsRequest) (_result *UpdateServiceEndsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceEndsResponse{}
	_body, _err := client.UpdateServiceEndsWithOptions(serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceEndsWithOptions(serviceId *string, request *UpdateServiceEndsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceEndsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNodes)) {
		body["serviceNodes"] = request.ServiceNodes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateServiceEndsResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateServiceEnds"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/service/"+tea.StringValue(serviceId)+"/serviceends"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FindGateways(request *FindGatewaysRequest) (_result *FindGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FindGatewaysResponse{}
	_body, _err := client.FindGatewaysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FindGatewaysWithOptions(request *FindGatewaysRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FindGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayUniqueId)) {
		query["gatewayUniqueId"] = request.GatewayUniqueId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayTypes)) {
		query["gatewayTypes"] = request.GatewayTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FindGatewaysResponse{}
	_body, _err := client.DoROARequest(tea.String("FindGateways"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * getAllRegistry
 */
func (client *Client) GetAllRegistry(gatewayId *string, request *GetAllRegistryRequest) (_result *GetAllRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAllRegistryResponse{}
	_body, _err := client.GetAllRegistryWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllRegistryWithOptions(gatewayId *string, request *GetAllRegistryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAllRegistryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["groupBy"] = request.GroupBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetAllRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAllRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/gateway/"+tea.StringValue(gatewayId)+"/registry"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * pullServiceInfoFromRegistry
 */
func (client *Client) PullServiceInfoFromRegistry(registryId *string) (_result *PullServiceInfoFromRegistryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PullServiceInfoFromRegistryResponse{}
	_body, _err := client.PullServiceInfoFromRegistryWithOptions(registryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullServiceInfoFromRegistryWithOptions(registryId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PullServiceInfoFromRegistryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &PullServiceInfoFromRegistryResponse{}
	_body, _err := client.DoROARequest(tea.String("PullServiceInfoFromRegistry"), tea.String("2020-08-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/registry/"+tea.StringValue(registryId)+"/pull"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

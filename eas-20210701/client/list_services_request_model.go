// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListServicesRequest
	GetAccessibility() *string
	SetAutoscalerEnabled(v bool) *ListServicesRequest
	GetAutoscalerEnabled() *bool
	SetCallerUid(v string) *ListServicesRequest
	GetCallerUid() *string
	SetCronscalerEnabled(v bool) *ListServicesRequest
	GetCronscalerEnabled() *bool
	SetFilter(v string) *ListServicesRequest
	GetFilter() *string
	SetGateway(v string) *ListServicesRequest
	GetGateway() *string
	SetGroupName(v string) *ListServicesRequest
	GetGroupName() *string
	SetIncludeNoWorkspace(v bool) *ListServicesRequest
	GetIncludeNoWorkspace() *bool
	SetLabel(v map[string]*string) *ListServicesRequest
	GetLabel() map[string]*string
	SetOrder(v string) *ListServicesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServicesRequest
	GetPageSize() *int32
	SetParentServiceUid(v string) *ListServicesRequest
	GetParentServiceUid() *string
	SetQuotaId(v string) *ListServicesRequest
	GetQuotaId() *string
	SetResourceAliasName(v string) *ListServicesRequest
	GetResourceAliasName() *string
	SetResourceBurstable(v bool) *ListServicesRequest
	GetResourceBurstable() *bool
	SetResourceId(v string) *ListServicesRequest
	GetResourceId() *string
	SetResourceName(v string) *ListServicesRequest
	GetResourceName() *string
	SetResourceType(v string) *ListServicesRequest
	GetResourceType() *string
	SetRole(v string) *ListServicesRequest
	GetRole() *string
	SetServiceName(v string) *ListServicesRequest
	GetServiceName() *string
	SetServiceStatus(v string) *ListServicesRequest
	GetServiceStatus() *string
	SetServiceType(v string) *ListServicesRequest
	GetServiceType() *string
	SetServiceUid(v string) *ListServicesRequest
	GetServiceUid() *string
	SetSort(v string) *ListServicesRequest
	GetSort() *string
	SetTrafficState(v string) *ListServicesRequest
	GetTrafficState() *string
	SetWorkspaceId(v string) *ListServicesRequest
	GetWorkspaceId() *string
}

type ListServicesRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Specifies whether to enable Auto Scaling for the service.
	//
	// example:
	//
	// true
	AutoscalerEnabled *bool `json:"AutoscalerEnabled,omitempty" xml:"AutoscalerEnabled,omitempty"`
	// The UID of the account that created the service.
	//
	// example:
	//
	// 19989224166xxxxxxx
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// Specifies whether to enable scheduled auto scaling for the service.
	//
	// example:
	//
	// true
	CronscalerEnabled *bool `json:"CronscalerEnabled,omitempty" xml:"CronscalerEnabled,omitempty"`
	// The keyword for a fuzzy search. This parameter supports fuzzy searches by service name only.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The name of the service group. To learn how to obtain this name, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// foo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to include services that do not belong to any workspace. The default value is true.
	//
	// example:
	//
	// true
	IncludeNoWorkspace *bool `json:"IncludeNoWorkspace,omitempty" xml:"IncludeNoWorkspace,omitempty"`
	// Filters services by label.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sort order. Valid values:
	//
	// - `desc` (default): descending.
	//
	// - `asc`: ascending.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the results to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of services to return per page. The default value is 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the primary service. This parameter applies to member services in a service group.
	//
	// example:
	//
	// eas-m-ijafy3c8cxxxx
	ParentServiceUid *string `json:"ParentServiceUid,omitempty" xml:"ParentServiceUid,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// quota1****
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The custom name of the resource group.
	//
	// example:
	//
	// example
	ResourceAliasName *string `json:"ResourceAliasName,omitempty" xml:"ResourceAliasName,omitempty"`
	// Specifies whether to enable a burstable resource pool for the service.
	//
	// example:
	//
	// true
	ResourceBurstable *bool `json:"ResourceBurstable,omitempty" xml:"ResourceBurstable,omitempty"`
	// The ID of the resource group. To learn how to query for this ID, see [ListResources](https://help.aliyun.com/document_detail/412133.html).
	//
	// example:
	//
	// eas-r-asdas****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Deprecated
	//
	// The name or ID of the service\\"s resource group.
	//
	// example:
	//
	// eas-r-hd0qwy8cxxxx
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of resource the service uses. Valid values:
	//
	// - PublicResource
	//
	// - DedicatedResource
	//
	// - Lingjun
	//
	// - SelfManagedLingjun
	//
	// example:
	//
	// PublicResource
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service role.
	//
	// example:
	//
	// LLMGateway
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The service name.
	//
	// example:
	//
	// echo_test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The status of the service.
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The service type. Valid values:
	//
	// - Async
	//
	// - Standard
	//
	// - Queue
	//
	// - LLM
	//
	// - RAG
	//
	// - Serverless
	//
	// - LLMGatewayService
	//
	// - OfflineTask
	//
	// - SDCluster
	//
	// - ScalableJob
	//
	// - ScalableJobService
	//
	// - AssistantJob
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service UID.
	//
	// example:
	//
	// eas-m-c9iw3yitxxxx
	ServiceUid *string `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	// The sort field. By default, results are sorted by timestamp in descending order.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies whether the service accepts group traffic. This parameter applies only to services within a service group.
	//
	// example:
	//
	// grouping
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListServicesRequest) GetAutoscalerEnabled() *bool {
	return s.AutoscalerEnabled
}

func (s *ListServicesRequest) GetCallerUid() *string {
	return s.CallerUid
}

func (s *ListServicesRequest) GetCronscalerEnabled() *bool {
	return s.CronscalerEnabled
}

func (s *ListServicesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListServicesRequest) GetGateway() *string {
	return s.Gateway
}

func (s *ListServicesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListServicesRequest) GetIncludeNoWorkspace() *bool {
	return s.IncludeNoWorkspace
}

func (s *ListServicesRequest) GetLabel() map[string]*string {
	return s.Label
}

func (s *ListServicesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServicesRequest) GetParentServiceUid() *string {
	return s.ParentServiceUid
}

func (s *ListServicesRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListServicesRequest) GetResourceAliasName() *string {
	return s.ResourceAliasName
}

func (s *ListServicesRequest) GetResourceBurstable() *bool {
	return s.ResourceBurstable
}

func (s *ListServicesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListServicesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListServicesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListServicesRequest) GetRole() *string {
	return s.Role
}

func (s *ListServicesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServicesRequest) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListServicesRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesRequest) GetServiceUid() *string {
	return s.ServiceUid
}

func (s *ListServicesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListServicesRequest) GetTrafficState() *string {
	return s.TrafficState
}

func (s *ListServicesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListServicesRequest) SetAccessibility(v string) *ListServicesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListServicesRequest) SetAutoscalerEnabled(v bool) *ListServicesRequest {
	s.AutoscalerEnabled = &v
	return s
}

func (s *ListServicesRequest) SetCallerUid(v string) *ListServicesRequest {
	s.CallerUid = &v
	return s
}

func (s *ListServicesRequest) SetCronscalerEnabled(v bool) *ListServicesRequest {
	s.CronscalerEnabled = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v string) *ListServicesRequest {
	s.Filter = &v
	return s
}

func (s *ListServicesRequest) SetGateway(v string) *ListServicesRequest {
	s.Gateway = &v
	return s
}

func (s *ListServicesRequest) SetGroupName(v string) *ListServicesRequest {
	s.GroupName = &v
	return s
}

func (s *ListServicesRequest) SetIncludeNoWorkspace(v bool) *ListServicesRequest {
	s.IncludeNoWorkspace = &v
	return s
}

func (s *ListServicesRequest) SetLabel(v map[string]*string) *ListServicesRequest {
	s.Label = v
	return s
}

func (s *ListServicesRequest) SetOrder(v string) *ListServicesRequest {
	s.Order = &v
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

func (s *ListServicesRequest) SetParentServiceUid(v string) *ListServicesRequest {
	s.ParentServiceUid = &v
	return s
}

func (s *ListServicesRequest) SetQuotaId(v string) *ListServicesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListServicesRequest) SetResourceAliasName(v string) *ListServicesRequest {
	s.ResourceAliasName = &v
	return s
}

func (s *ListServicesRequest) SetResourceBurstable(v bool) *ListServicesRequest {
	s.ResourceBurstable = &v
	return s
}

func (s *ListServicesRequest) SetResourceId(v string) *ListServicesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListServicesRequest) SetResourceName(v string) *ListServicesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListServicesRequest) SetResourceType(v string) *ListServicesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServicesRequest) SetRole(v string) *ListServicesRequest {
	s.Role = &v
	return s
}

func (s *ListServicesRequest) SetServiceName(v string) *ListServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesRequest) SetServiceStatus(v string) *ListServicesRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesRequest) SetServiceUid(v string) *ListServicesRequest {
	s.ServiceUid = &v
	return s
}

func (s *ListServicesRequest) SetSort(v string) *ListServicesRequest {
	s.Sort = &v
	return s
}

func (s *ListServicesRequest) SetTrafficState(v string) *ListServicesRequest {
	s.TrafficState = &v
	return s
}

func (s *ListServicesRequest) SetWorkspaceId(v string) *ListServicesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}

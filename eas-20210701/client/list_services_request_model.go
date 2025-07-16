// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoscalerEnabled(v bool) *ListServicesRequest
	GetAutoscalerEnabled() *bool
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
	SetWorkspaceId(v string) *ListServicesRequest
	GetWorkspaceId() *string
}

type ListServicesRequest struct {
	AutoscalerEnabled *bool `json:"AutoscalerEnabled,omitempty" xml:"AutoscalerEnabled,omitempty"`
	CronscalerEnabled *bool `json:"CronscalerEnabled,omitempty" xml:"CronscalerEnabled,omitempty"`
	// The field that is used for fuzzy matches. The system performs fuzzy matches only by service name.
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
	// The name of the service group. For more information about how to query the name of a service group, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// foo
	GroupName          *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IncludeNoWorkspace *bool   `json:"IncludeNoWorkspace,omitempty" xml:"IncludeNoWorkspace,omitempty"`
	// The tag that is used to filter services.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default): The query results are sorted in descending order.
	//
	// 	- asc: The query results are sorted in ascending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the primary service that corresponds to the Band member service.
	//
	// example:
	//
	// eas-m-ijafy3c8cxxxx
	ParentServiceUid *string `json:"ParentServiceUid,omitempty" xml:"ParentServiceUid,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// quota12345
	QuotaId           *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	ResourceAliasName *string `json:"ResourceAliasName,omitempty" xml:"ResourceAliasName,omitempty"`
	ResourceBurstable *bool   `json:"ResourceBurstable,omitempty" xml:"ResourceBurstable,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Deprecated
	//
	// The name or ID of the resource group to which the service belongs.
	//
	// example:
	//
	// eas-r-hd0qwy8cxxxx
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The server role.
	//
	// Valid values:
	//
	// 	- DataLoader
	//
	// 	- FrontEnd
	//
	// 	- DataSet
	//
	// 	- SDProxy
	//
	// 	- LLMSscheduler
	//
	// 	- ScalableJob
	//
	// 	- LLMGateway
	//
	// 	- Job
	//
	// 	- Queue
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
	// The service state.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Failed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Complete
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cloning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Waiting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- HotUpdate
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Committing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Developing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Scaling
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The service type. Valid values:
	//
	// 	- Async
	//
	// 	- Standard
	//
	// 	- Offline Task
	//
	// 	- Proxima
	//
	// Valid values:
	//
	// 	- Async
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- OfflineTask
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Proxima
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The user ID (UID) of the service.
	//
	// example:
	//
	// eas-m-c9iw3yitxxxx
	ServiceUid *string `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	// The sort field. By default, the query results are sorted by the timestamp type in descending order.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetAutoscalerEnabled() *bool {
	return s.AutoscalerEnabled
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

func (s *ListServicesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListServicesRequest) SetAutoscalerEnabled(v bool) *ListServicesRequest {
	s.AutoscalerEnabled = &v
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

func (s *ListServicesRequest) SetWorkspaceId(v string) *ListServicesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}

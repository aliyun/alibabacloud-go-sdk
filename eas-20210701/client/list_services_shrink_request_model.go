// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoscalerEnabled(v bool) *ListServicesShrinkRequest
	GetAutoscalerEnabled() *bool
	SetCronscalerEnabled(v bool) *ListServicesShrinkRequest
	GetCronscalerEnabled() *bool
	SetFilter(v string) *ListServicesShrinkRequest
	GetFilter() *string
	SetGateway(v string) *ListServicesShrinkRequest
	GetGateway() *string
	SetGroupName(v string) *ListServicesShrinkRequest
	GetGroupName() *string
	SetIncludeNoWorkspace(v bool) *ListServicesShrinkRequest
	GetIncludeNoWorkspace() *bool
	SetLabelShrink(v string) *ListServicesShrinkRequest
	GetLabelShrink() *string
	SetOrder(v string) *ListServicesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListServicesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServicesShrinkRequest
	GetPageSize() *int32
	SetParentServiceUid(v string) *ListServicesShrinkRequest
	GetParentServiceUid() *string
	SetQuotaId(v string) *ListServicesShrinkRequest
	GetQuotaId() *string
	SetResourceAliasName(v string) *ListServicesShrinkRequest
	GetResourceAliasName() *string
	SetResourceBurstable(v bool) *ListServicesShrinkRequest
	GetResourceBurstable() *bool
	SetResourceId(v string) *ListServicesShrinkRequest
	GetResourceId() *string
	SetResourceName(v string) *ListServicesShrinkRequest
	GetResourceName() *string
	SetResourceType(v string) *ListServicesShrinkRequest
	GetResourceType() *string
	SetRole(v string) *ListServicesShrinkRequest
	GetRole() *string
	SetServiceName(v string) *ListServicesShrinkRequest
	GetServiceName() *string
	SetServiceStatus(v string) *ListServicesShrinkRequest
	GetServiceStatus() *string
	SetServiceType(v string) *ListServicesShrinkRequest
	GetServiceType() *string
	SetServiceUid(v string) *ListServicesShrinkRequest
	GetServiceUid() *string
	SetSort(v string) *ListServicesShrinkRequest
	GetSort() *string
	SetTrafficState(v string) *ListServicesShrinkRequest
	GetTrafficState() *string
	SetWorkspaceId(v string) *ListServicesShrinkRequest
	GetWorkspaceId() *string
}

type ListServicesShrinkRequest struct {
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
	LabelShrink *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
	Sort         *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServicesShrinkRequest) GetAutoscalerEnabled() *bool {
	return s.AutoscalerEnabled
}

func (s *ListServicesShrinkRequest) GetCronscalerEnabled() *bool {
	return s.CronscalerEnabled
}

func (s *ListServicesShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListServicesShrinkRequest) GetGateway() *string {
	return s.Gateway
}

func (s *ListServicesShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListServicesShrinkRequest) GetIncludeNoWorkspace() *bool {
	return s.IncludeNoWorkspace
}

func (s *ListServicesShrinkRequest) GetLabelShrink() *string {
	return s.LabelShrink
}

func (s *ListServicesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListServicesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServicesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServicesShrinkRequest) GetParentServiceUid() *string {
	return s.ParentServiceUid
}

func (s *ListServicesShrinkRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListServicesShrinkRequest) GetResourceAliasName() *string {
	return s.ResourceAliasName
}

func (s *ListServicesShrinkRequest) GetResourceBurstable() *bool {
	return s.ResourceBurstable
}

func (s *ListServicesShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListServicesShrinkRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListServicesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListServicesShrinkRequest) GetRole() *string {
	return s.Role
}

func (s *ListServicesShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServicesShrinkRequest) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListServicesShrinkRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesShrinkRequest) GetServiceUid() *string {
	return s.ServiceUid
}

func (s *ListServicesShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListServicesShrinkRequest) GetTrafficState() *string {
	return s.TrafficState
}

func (s *ListServicesShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListServicesShrinkRequest) SetAutoscalerEnabled(v bool) *ListServicesShrinkRequest {
	s.AutoscalerEnabled = &v
	return s
}

func (s *ListServicesShrinkRequest) SetCronscalerEnabled(v bool) *ListServicesShrinkRequest {
	s.CronscalerEnabled = &v
	return s
}

func (s *ListServicesShrinkRequest) SetFilter(v string) *ListServicesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListServicesShrinkRequest) SetGateway(v string) *ListServicesShrinkRequest {
	s.Gateway = &v
	return s
}

func (s *ListServicesShrinkRequest) SetGroupName(v string) *ListServicesShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetIncludeNoWorkspace(v bool) *ListServicesShrinkRequest {
	s.IncludeNoWorkspace = &v
	return s
}

func (s *ListServicesShrinkRequest) SetLabelShrink(v string) *ListServicesShrinkRequest {
	s.LabelShrink = &v
	return s
}

func (s *ListServicesShrinkRequest) SetOrder(v string) *ListServicesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListServicesShrinkRequest) SetPageNumber(v int32) *ListServicesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesShrinkRequest) SetPageSize(v int32) *ListServicesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesShrinkRequest) SetParentServiceUid(v string) *ListServicesShrinkRequest {
	s.ParentServiceUid = &v
	return s
}

func (s *ListServicesShrinkRequest) SetQuotaId(v string) *ListServicesShrinkRequest {
	s.QuotaId = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceAliasName(v string) *ListServicesShrinkRequest {
	s.ResourceAliasName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceBurstable(v bool) *ListServicesShrinkRequest {
	s.ResourceBurstable = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceId(v string) *ListServicesShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceName(v string) *ListServicesShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceType(v string) *ListServicesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServicesShrinkRequest) SetRole(v string) *ListServicesShrinkRequest {
	s.Role = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceName(v string) *ListServicesShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceStatus(v string) *ListServicesShrinkRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceType(v string) *ListServicesShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceUid(v string) *ListServicesShrinkRequest {
	s.ServiceUid = &v
	return s
}

func (s *ListServicesShrinkRequest) SetSort(v string) *ListServicesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListServicesShrinkRequest) SetTrafficState(v string) *ListServicesShrinkRequest {
	s.TrafficState = &v
	return s
}

func (s *ListServicesShrinkRequest) SetWorkspaceId(v string) *ListServicesShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListServicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

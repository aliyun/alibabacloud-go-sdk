// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListServiceInstancesRequest
	GetFilter() *string
	SetHostIP(v string) *ListServiceInstancesRequest
	GetHostIP() *string
	SetInstanceIP(v string) *ListServiceInstancesRequest
	GetInstanceIP() *string
	SetInstanceName(v string) *ListServiceInstancesRequest
	GetInstanceName() *string
	SetInstanceStatus(v string) *ListServiceInstancesRequest
	GetInstanceStatus() *string
	SetInstanceType(v string) *ListServiceInstancesRequest
	GetInstanceType() *string
	SetIsSpot(v bool) *ListServiceInstancesRequest
	GetIsSpot() *bool
	SetMemberType(v string) *ListServiceInstancesRequest
	GetMemberType() *string
	SetOrder(v string) *ListServiceInstancesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListServiceInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServiceInstancesRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListServiceInstancesRequest
	GetResourceType() *string
	SetRole(v string) *ListServiceInstancesRequest
	GetRole() *string
	SetSort(v string) *ListServiceInstancesRequest
	GetSort() *string
}

type ListServiceInstancesRequest struct {
	// The keyword used to query instances. Instances can be queried based on instance name, instance IP address, IP address of the server where the instance resides, and instance type.
	//
	// example:
	//
	// 10.118.xx.xx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The IP address of the server where the instance resides.
	//
	// example:
	//
	// 10.224.xx.xx
	HostIP *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.224.xx.xx
	InstanceIP *string `json:"InstanceIP,omitempty" xml:"InstanceIP,omitempty"`
	// The instance name.
	//
	// example:
	//
	// foo-bdc5xxxx-8l7rk
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance state.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the instance is a preemptible instance.
	//
	// example:
	//
	// false
	IsSpot     *bool   `json:"IsSpot,omitempty" xml:"IsSpot,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The sorting order.
	//
	// Valid values:
	//
	// 	- asc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The instances are sorted in ascending order.
	//
	// 	- desc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The instances are sorted in descending order.
	//
	// example:
	//
	// desc
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource group to which the instance belongs.
	//
	// Valid values:
	//
	// 	- PublicResource
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DedicatedResource
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PublicResource
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service role.
	//
	// Valid values:
	//
	// 	- DataSet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     dataset service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- SDProxy
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Stable-Diffusion proxy service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     standard service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Queue
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     queue service
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Queue
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The field that you use to sort the query results.
	//
	// 	- Set the value to StartTime.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The value specifies that the query results are sorted based on the time when the instances were created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// StartTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListServiceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListServiceInstancesRequest) GetHostIP() *string {
	return s.HostIP
}

func (s *ListServiceInstancesRequest) GetInstanceIP() *string {
	return s.InstanceIP
}

func (s *ListServiceInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListServiceInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListServiceInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListServiceInstancesRequest) GetIsSpot() *bool {
	return s.IsSpot
}

func (s *ListServiceInstancesRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *ListServiceInstancesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListServiceInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceInstancesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListServiceInstancesRequest) GetRole() *string {
	return s.Role
}

func (s *ListServiceInstancesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListServiceInstancesRequest) SetFilter(v string) *ListServiceInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListServiceInstancesRequest) SetHostIP(v string) *ListServiceInstancesRequest {
	s.HostIP = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceIP(v string) *ListServiceInstancesRequest {
	s.InstanceIP = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceName(v string) *ListServiceInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceStatus(v string) *ListServiceInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceType(v string) *ListServiceInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListServiceInstancesRequest) SetIsSpot(v bool) *ListServiceInstancesRequest {
	s.IsSpot = &v
	return s
}

func (s *ListServiceInstancesRequest) SetMemberType(v string) *ListServiceInstancesRequest {
	s.MemberType = &v
	return s
}

func (s *ListServiceInstancesRequest) SetOrder(v string) *ListServiceInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListServiceInstancesRequest) SetPageNumber(v int32) *ListServiceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesRequest) SetPageSize(v int32) *ListServiceInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceType(v string) *ListServiceInstancesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRole(v string) *ListServiceInstancesRequest {
	s.Role = &v
	return s
}

func (s *ListServiceInstancesRequest) SetSort(v string) *ListServiceInstancesRequest {
	s.Sort = &v
	return s
}

func (s *ListServiceInstancesRequest) Validate() error {
	return dara.Validate(s)
}

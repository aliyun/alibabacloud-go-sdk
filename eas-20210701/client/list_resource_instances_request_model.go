// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ListResourceInstancesRequest
	GetChargeType() *string
	SetFilter(v string) *ListResourceInstancesRequest
	GetFilter() *string
	SetInstanceIP(v string) *ListResourceInstancesRequest
	GetInstanceIP() *string
	SetInstanceId(v string) *ListResourceInstancesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListResourceInstancesRequest
	GetInstanceName() *string
	SetInstanceStatus(v string) *ListResourceInstancesRequest
	GetInstanceStatus() *string
	SetLabel(v map[string]*string) *ListResourceInstancesRequest
	GetLabel() map[string]*string
	SetOrder(v string) *ListResourceInstancesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListResourceInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstancesRequest
	GetPageSize() *int32
	SetSort(v string) *ListResourceInstancesRequest
	GetSort() *string
	SetZone(v string) *ListResourceInstancesRequest
	GetZone() *string
}

type ListResourceInstancesRequest struct {
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The keyword used to query instances. Instances can be queried by instance ID or instance IP address.
	//
	// example:
	//
	// 10.224.xx.xx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.224.xx.xx
	InstanceIP *string `json:"InstanceIP,omitempty" xml:"InstanceIP,omitempty"`
	// The instance ID. For more information about how to query the instance ID, see [ListResourceInstances](https://help.aliyun.com/document_detail/412129.html).
	//
	// example:
	//
	// i-bp1jd6x3uotsv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// e-xxxx***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance state.
	//
	// Valid values:
	//
	// 	- Ready-SchedulingDisabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance is available but unschedulable
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Ready
	//
	//     <!-- -->
	//
	//     : The instance
	//
	//     <!-- -->
	//
	//     is running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- NotReady
	//
	//     <!-- -->
	//
	//     : The instance is unready.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     : The instance has stopped.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NotReady-SchedulingDisabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance is unavailable and unschedulable
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Attaching
	//
	//     <!-- -->
	//
	//     : The instance
	//
	//     <!-- -->
	//
	//     is starting
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     : The instance is being deleted.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed: The instance failed to be created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Ready
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The tag.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sorting order.
	//
	// Valid values:
	//
	// 	- asc: The instances are sorted in ascending order.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- desc
	//
	//     <!-- -->
	//
	//     : The instances are sorted in descending order.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- CreateTime
	//
	//     <!-- -->
	//
	//     : The instances are sorted based on the time when the instances were created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MemoryUsed
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instances are sorted based on the memory usage of the instances
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- GpuUsed
	//
	//     <!-- -->
	//
	//     : The instances are sorted based on the
	//
	//     <!-- -->
	//
	//     GPU usage of the instances.
	//
	//     <!-- -->
	//
	// 	- ExpireTime: The instances are sorted based on the time when the instances expired.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CpuUsed
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instances are sorted based on the CPU utilization of the instances.
	//
	//     <!-- -->
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s ListResourceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListResourceInstancesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListResourceInstancesRequest) GetInstanceIP() *string {
	return s.InstanceIP
}

func (s *ListResourceInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListResourceInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListResourceInstancesRequest) GetLabel() map[string]*string {
	return s.Label
}

func (s *ListResourceInstancesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstancesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourceInstancesRequest) GetZone() *string {
	return s.Zone
}

func (s *ListResourceInstancesRequest) SetChargeType(v string) *ListResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListResourceInstancesRequest) SetFilter(v string) *ListResourceInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceIP(v string) *ListResourceInstancesRequest {
	s.InstanceIP = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceId(v string) *ListResourceInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceName(v string) *ListResourceInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceStatus(v string) *ListResourceInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListResourceInstancesRequest) SetLabel(v map[string]*string) *ListResourceInstancesRequest {
	s.Label = v
	return s
}

func (s *ListResourceInstancesRequest) SetOrder(v string) *ListResourceInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageNumber(v int32) *ListResourceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageSize(v int32) *ListResourceInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesRequest) SetSort(v string) *ListResourceInstancesRequest {
	s.Sort = &v
	return s
}

func (s *ListResourceInstancesRequest) SetZone(v string) *ListResourceInstancesRequest {
	s.Zone = &v
	return s
}

func (s *ListResourceInstancesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ListResourceInstancesShrinkRequest
	GetChargeType() *string
	SetFilter(v string) *ListResourceInstancesShrinkRequest
	GetFilter() *string
	SetInstanceIP(v string) *ListResourceInstancesShrinkRequest
	GetInstanceIP() *string
	SetInstanceId(v string) *ListResourceInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListResourceInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceStatus(v string) *ListResourceInstancesShrinkRequest
	GetInstanceStatus() *string
	SetLabelShrink(v string) *ListResourceInstancesShrinkRequest
	GetLabelShrink() *string
	SetOrder(v string) *ListResourceInstancesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListResourceInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstancesShrinkRequest
	GetPageSize() *int32
	SetSort(v string) *ListResourceInstancesShrinkRequest
	GetSort() *string
	SetZone(v string) *ListResourceInstancesShrinkRequest
	GetZone() *string
}

type ListResourceInstancesShrinkRequest struct {
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
	LabelShrink *string `json:"Label,omitempty" xml:"Label,omitempty"`
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

func (s ListResourceInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListResourceInstancesShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListResourceInstancesShrinkRequest) GetInstanceIP() *string {
	return s.InstanceIP
}

func (s *ListResourceInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListResourceInstancesShrinkRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListResourceInstancesShrinkRequest) GetLabelShrink() *string {
	return s.LabelShrink
}

func (s *ListResourceInstancesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstancesShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourceInstancesShrinkRequest) GetZone() *string {
	return s.Zone
}

func (s *ListResourceInstancesShrinkRequest) SetChargeType(v string) *ListResourceInstancesShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetFilter(v string) *ListResourceInstancesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetInstanceIP(v string) *ListResourceInstancesShrinkRequest {
	s.InstanceIP = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetInstanceId(v string) *ListResourceInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetInstanceName(v string) *ListResourceInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetInstanceStatus(v string) *ListResourceInstancesShrinkRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetLabelShrink(v string) *ListResourceInstancesShrinkRequest {
	s.LabelShrink = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetOrder(v string) *ListResourceInstancesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetPageNumber(v int32) *ListResourceInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetPageSize(v int32) *ListResourceInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetSort(v string) *ListResourceInstancesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) SetZone(v string) *ListResourceInstancesShrinkRequest {
	s.Zone = &v
	return s
}

func (s *ListResourceInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

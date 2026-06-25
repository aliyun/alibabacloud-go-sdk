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
	// Filters instances by billing method. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A keyword for the search. You can search by instance ID or IP address.
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
	// Filter by instance ID. For more information, see [ListResourceInstances](https://help.aliyun.com/document_detail/412129.html).
	//
	// example:
	//
	// i-bp1jd6x3uots****a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Filter by instance name.
	//
	// example:
	//
	// e-xxxx***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Ready
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Filter by label.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances to return on each page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The zone of the instance.
	//
	// example:
	//
	// J
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

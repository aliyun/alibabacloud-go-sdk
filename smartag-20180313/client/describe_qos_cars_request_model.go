// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosCarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeQosCarsRequest
	GetDescription() *string
	SetOrder(v string) *DescribeQosCarsRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeQosCarsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeQosCarsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeQosCarsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosCarsRequest
	GetPageSize() *int32
	SetQosCarId(v string) *DescribeQosCarsRequest
	GetQosCarId() *string
	SetQosId(v string) *DescribeQosCarsRequest
	GetQosId() *string
	SetRegionId(v string) *DescribeQosCarsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeQosCarsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeQosCarsRequest
	GetResourceOwnerId() *int64
}

type DescribeQosCarsRequest struct {
	// The description of the traffic throttling rule.
	//
	// example:
	//
	// testdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The sorting method of the traffic throttling rules. Valid values:
	//
	// 	- **asc**: sorted in ascending order. This is the default value.
	//
	// 	- **desc**: sorted in descending order.
	//
	// By default, traffic throttling rules are sorted in ascending order of priority.
	//
	// example:
	//
	// asc
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the traffic throttling rule.
	//
	// example:
	//
	// qoscar-n5k8g97lihlph****
	QosCarId *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-awfxl1adxeqyk****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the region where the QoS policy is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeQosCarsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosCarsRequest) GoString() string {
	return s.String()
}

func (s *DescribeQosCarsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeQosCarsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeQosCarsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeQosCarsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeQosCarsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosCarsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosCarsRequest) GetQosCarId() *string {
	return s.QosCarId
}

func (s *DescribeQosCarsRequest) GetQosId() *string {
	return s.QosId
}

func (s *DescribeQosCarsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeQosCarsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeQosCarsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeQosCarsRequest) SetDescription(v string) *DescribeQosCarsRequest {
	s.Description = &v
	return s
}

func (s *DescribeQosCarsRequest) SetOrder(v string) *DescribeQosCarsRequest {
	s.Order = &v
	return s
}

func (s *DescribeQosCarsRequest) SetOwnerAccount(v string) *DescribeQosCarsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeQosCarsRequest) SetOwnerId(v int64) *DescribeQosCarsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeQosCarsRequest) SetPageNumber(v int32) *DescribeQosCarsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosCarsRequest) SetPageSize(v int32) *DescribeQosCarsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeQosCarsRequest) SetQosCarId(v string) *DescribeQosCarsRequest {
	s.QosCarId = &v
	return s
}

func (s *DescribeQosCarsRequest) SetQosId(v string) *DescribeQosCarsRequest {
	s.QosId = &v
	return s
}

func (s *DescribeQosCarsRequest) SetRegionId(v string) *DescribeQosCarsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeQosCarsRequest) SetResourceOwnerAccount(v string) *DescribeQosCarsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeQosCarsRequest) SetResourceOwnerId(v int64) *DescribeQosCarsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeQosCarsRequest) Validate() error {
	return dara.Validate(s)
}

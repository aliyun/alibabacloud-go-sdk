// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribe95TrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDay(v string) *Describe95TrafficRequest
	GetDay() *string
	SetInstanceId(v string) *Describe95TrafficRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *Describe95TrafficRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *Describe95TrafficRequest
	GetOwnerId() *int64
	SetRegionId(v string) *Describe95TrafficRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *Describe95TrafficRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *Describe95TrafficRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *Describe95TrafficRequest
	GetResourceType() *string
}

type Describe95TrafficRequest struct {
	// The date in UTC+8. Format: year-month-day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-22
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-wz9j19xrwf78fvz7*****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to cbwp, which specifies Internet Shared Bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s Describe95TrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficRequest) GoString() string {
	return s.String()
}

func (s *Describe95TrafficRequest) GetDay() *string {
	return s.Day
}

func (s *Describe95TrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *Describe95TrafficRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *Describe95TrafficRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *Describe95TrafficRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *Describe95TrafficRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *Describe95TrafficRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *Describe95TrafficRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *Describe95TrafficRequest) SetDay(v string) *Describe95TrafficRequest {
	s.Day = &v
	return s
}

func (s *Describe95TrafficRequest) SetInstanceId(v string) *Describe95TrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *Describe95TrafficRequest) SetOwnerAccount(v string) *Describe95TrafficRequest {
	s.OwnerAccount = &v
	return s
}

func (s *Describe95TrafficRequest) SetOwnerId(v int64) *Describe95TrafficRequest {
	s.OwnerId = &v
	return s
}

func (s *Describe95TrafficRequest) SetRegionId(v string) *Describe95TrafficRequest {
	s.RegionId = &v
	return s
}

func (s *Describe95TrafficRequest) SetResourceOwnerAccount(v string) *Describe95TrafficRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *Describe95TrafficRequest) SetResourceOwnerId(v int64) *Describe95TrafficRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *Describe95TrafficRequest) SetResourceType(v string) *Describe95TrafficRequest {
	s.ResourceType = &v
	return s
}

func (s *Describe95TrafficRequest) Validate() error {
	return dara.Validate(s)
}

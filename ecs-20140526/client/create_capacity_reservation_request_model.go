// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *CreateCapacityReservationRequestPrivatePoolOptions) *CreateCapacityReservationRequest
	GetPrivatePoolOptions() *CreateCapacityReservationRequestPrivatePoolOptions
	SetClientToken(v string) *CreateCapacityReservationRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCapacityReservationRequest
	GetDescription() *string
	SetEndTime(v string) *CreateCapacityReservationRequest
	GetEndTime() *string
	SetEndTimeType(v string) *CreateCapacityReservationRequest
	GetEndTimeType() *string
	SetInstanceAmount(v int32) *CreateCapacityReservationRequest
	GetInstanceAmount() *int32
	SetInstanceChargeType(v string) *CreateCapacityReservationRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *CreateCapacityReservationRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *CreateCapacityReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCapacityReservationRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateCapacityReservationRequest
	GetPlatform() *string
	SetRegionId(v string) *CreateCapacityReservationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCapacityReservationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCapacityReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCapacityReservationRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CreateCapacityReservationRequest
	GetStartTime() *string
	SetTag(v []*CreateCapacityReservationRequestTag) *CreateCapacityReservationRequest
	GetTag() []*CreateCapacityReservationRequestTag
	SetZoneId(v []*string) *CreateCapacityReservationRequest
	GetZoneId() []*string
}

type CreateCapacityReservationRequest struct {
	PrivatePoolOptions *CreateCapacityReservationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// A client-generated token that ensures the request is idempotent. You can use the same token to retry a request. The `ClientToken` value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the capacity reservation. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty string.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the capacity reservation. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2021-10-30T06:32:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The release mode of the capacity reservation. Valid values:
	//
	// - Limited: The capacity reservation is automatically released at a specific time. You must also specify the `EndTime` parameter.
	//
	// - Unlimited: The capacity reservation must be released manually.
	//
	// example:
	//
	// Unlimited
	EndTimeType *string `json:"EndTimeType,omitempty" xml:"EndTimeType,omitempty"`
	// The number of instances of the specified instance type for which to reserve capacity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	InstanceAmount     *int32  `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type for which to reserve capacity. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to view the instance types that ECS provides.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system of the image used by the instance. This parameter corresponds to the `Platform` parameter of a regional reserved instance. If this platform matches the platform of a regional reserved instance, the regional reserved instance can be used to offset the costs of unused capacity in the reservation. Valid values:
	//
	// - Windows: Windows Server operating systems.
	//
	// - Linux: Linux and Unix-like operating systems.
	//
	// Default value: Linux.
	//
	// > This parameter is not yet available for use.
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the region in which to create the capacity reservation. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the capacity reservation belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when the capacity reservation takes effect. The capacity reservation takes effect immediately after it is created.
	//
	// > If you do not specify this parameter, the capacity reservation takes effect immediately.
	//
	// example:
	//
	// 2021-10-30T05:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags to add to the capacity reservation.
	Tag []*CreateCapacityReservationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone in which you want to create the capacity reservation. A capacity reservation can reserve resources within only one zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s CreateCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequest) GetPrivatePoolOptions() *CreateCapacityReservationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateCapacityReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCapacityReservationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCapacityReservationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCapacityReservationRequest) GetEndTimeType() *string {
	return s.EndTimeType
}

func (s *CreateCapacityReservationRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *CreateCapacityReservationRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateCapacityReservationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCapacityReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCapacityReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCapacityReservationRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCapacityReservationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCapacityReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCapacityReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCapacityReservationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCapacityReservationRequest) GetTag() []*CreateCapacityReservationRequestTag {
	return s.Tag
}

func (s *CreateCapacityReservationRequest) GetZoneId() []*string {
	return s.ZoneId
}

func (s *CreateCapacityReservationRequest) SetPrivatePoolOptions(v *CreateCapacityReservationRequestPrivatePoolOptions) *CreateCapacityReservationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateCapacityReservationRequest) SetClientToken(v string) *CreateCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetDescription(v string) *CreateCapacityReservationRequest {
	s.Description = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetEndTime(v string) *CreateCapacityReservationRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetEndTimeType(v string) *CreateCapacityReservationRequest {
	s.EndTimeType = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetInstanceAmount(v int32) *CreateCapacityReservationRequest {
	s.InstanceAmount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetInstanceChargeType(v string) *CreateCapacityReservationRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetInstanceType(v string) *CreateCapacityReservationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetOwnerAccount(v string) *CreateCapacityReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetOwnerId(v int64) *CreateCapacityReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetPlatform(v string) *CreateCapacityReservationRequest {
	s.Platform = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetRegionId(v string) *CreateCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceGroupId(v string) *CreateCapacityReservationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceOwnerAccount(v string) *CreateCapacityReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceOwnerId(v int64) *CreateCapacityReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetStartTime(v string) *CreateCapacityReservationRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetTag(v []*CreateCapacityReservationRequestTag) *CreateCapacityReservationRequest {
	s.Tag = v
	return s
}

func (s *CreateCapacityReservationRequest) SetZoneId(v []*string) *CreateCapacityReservationRequest {
	s.ZoneId = v
	return s
}

func (s *CreateCapacityReservationRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCapacityReservationRequestPrivatePoolOptions struct {
	// The type of the private pool that is generated after the capacity reservation takes effect. Valid values:
	//
	// - Open: open mode. When you launch an instance, it is automatically matched with the capacity of an open private pool. If no suitable private pool capacity is available, the instance is launched by using public pool resources.
	//
	// - Target: targeted mode. The instance is launched by using the capacity of a specified private pool. If the capacity is unavailable, the instance fails to launch.
	//
	// Default value: Open.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// The name of the capacity reservation. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// crpTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCapacityReservationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateCapacityReservationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) SetName(v string) *CreateCapacityReservationRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateCapacityReservationRequestTag struct {
	// The tag key of the capacity reservation. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the capacity reservation. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCapacityReservationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCapacityReservationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCapacityReservationRequestTag) SetKey(v string) *CreateCapacityReservationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCapacityReservationRequestTag) SetValue(v string) *CreateCapacityReservationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCapacityReservationRequestTag) Validate() error {
	return dara.Validate(s)
}

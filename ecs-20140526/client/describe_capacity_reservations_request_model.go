// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeCapacityReservationsRequestPrivatePoolOptions) *DescribeCapacityReservationsRequest
	GetPrivatePoolOptions() *DescribeCapacityReservationsRequestPrivatePoolOptions
	SetInstanceChargeType(v string) *DescribeCapacityReservationsRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeCapacityReservationsRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeCapacityReservationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCapacityReservationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeCapacityReservationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCapacityReservationsRequest
	GetOwnerId() *int64
	SetPlatform(v string) *DescribeCapacityReservationsRequest
	GetPlatform() *string
	SetRegionId(v string) *DescribeCapacityReservationsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCapacityReservationsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCapacityReservationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCapacityReservationsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeCapacityReservationsRequest
	GetStatus() *string
	SetTag(v []*DescribeCapacityReservationsRequestTag) *DescribeCapacityReservationsRequest
	GetTag() []*DescribeCapacityReservationsRequestTag
	SetZoneId(v string) *DescribeCapacityReservationsRequest
	GetZoneId() *string
}

type DescribeCapacityReservationsRequest struct {
	PrivatePoolOptions *DescribeCapacityReservationsRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// - PostPaid: pay-as-you-go.
	//
	// - PrePaid: subscription.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type. You can use the instance type to query only active capacity reservations. Released capacity reservations can be queried only by using PrivatePoolOptions.Ids.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the capacity reservation query. Obtain the value from the result of the previous request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system of the instance. Valid values:
	//
	// - windows: queries only capacity reservations for Windows instances.
	//
	// - linux: queries only capacity reservations for Linux instances.
	//
	// - all: queries all capacity reservations.
	//
	// Default value: all.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the capacity reservation. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// >Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the capacity reservation. Valid values:
	//
	// - All: all states.
	//
	// - Pending: initializing. A capacity reservation that takes effect at a specified time enters the initializing state first.
	//
	// - Preparing: being prepared. A capacity reservation that takes effect at a specified time is in the Preparing state during the resource delivery phase.
	//
	// - Prepared: to take effect. A capacity reservation that takes effect at a specified time is in the Prepared state after resource delivery is complete but before the service takes effect.
	//
	// - Active: active.
	//
	// - Released: released, including manual release and automatic release upon expiration.
	//
	// If you do not specify this parameter, capacity reservations in all states except Pending and Released are queried.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags bound to the capacity reservation.
	Tag []*DescribeCapacityReservationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the capacity reservation.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCapacityReservationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsRequest) GetPrivatePoolOptions() *DescribeCapacityReservationsRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeCapacityReservationsRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeCapacityReservationsRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCapacityReservationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCapacityReservationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCapacityReservationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCapacityReservationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCapacityReservationsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeCapacityReservationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCapacityReservationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCapacityReservationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCapacityReservationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCapacityReservationsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCapacityReservationsRequest) GetTag() []*DescribeCapacityReservationsRequestTag {
	return s.Tag
}

func (s *DescribeCapacityReservationsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeCapacityReservationsRequest) SetPrivatePoolOptions(v *DescribeCapacityReservationsRequestPrivatePoolOptions) *DescribeCapacityReservationsRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetInstanceChargeType(v string) *DescribeCapacityReservationsRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetInstanceType(v string) *DescribeCapacityReservationsRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetMaxResults(v int32) *DescribeCapacityReservationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetNextToken(v string) *DescribeCapacityReservationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetOwnerAccount(v string) *DescribeCapacityReservationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetOwnerId(v int64) *DescribeCapacityReservationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetPlatform(v string) *DescribeCapacityReservationsRequest {
	s.Platform = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetRegionId(v string) *DescribeCapacityReservationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetResourceGroupId(v string) *DescribeCapacityReservationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetResourceOwnerAccount(v string) *DescribeCapacityReservationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetResourceOwnerId(v int64) *DescribeCapacityReservationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetStatus(v string) *DescribeCapacityReservationsRequest {
	s.Status = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetTag(v []*DescribeCapacityReservationsRequestTag) *DescribeCapacityReservationsRequest {
	s.Tag = v
	return s
}

func (s *DescribeCapacityReservationsRequest) SetZoneId(v string) *DescribeCapacityReservationsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeCapacityReservationsRequest) Validate() error {
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

type DescribeCapacityReservationsRequestPrivatePoolOptions struct {
	// The list of capacity reservation IDs. The value can be a JSON array that consists of up to 100 IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["crp-bp1gubrkqutenqdd****", "crp-bp67acfmxazb5****"]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DescribeCapacityReservationsRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsRequestPrivatePoolOptions) GetIds() *string {
	return s.Ids
}

func (s *DescribeCapacityReservationsRequestPrivatePoolOptions) SetIds(v string) *DescribeCapacityReservationsRequestPrivatePoolOptions {
	s.Ids = &v
	return s
}

func (s *DescribeCapacityReservationsRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeCapacityReservationsRequestTag struct {
	// The tag key. N indicates that you can set multiple tag keys for filtering. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1000. If you use multiple tags to filter resources, the resource count with all specified tags attached cannot exceed 1000. If the resource count exceeds 1000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. N indicates that you can set multiple tag values for filtering. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCapacityReservationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCapacityReservationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCapacityReservationsRequestTag) SetKey(v string) *DescribeCapacityReservationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCapacityReservationsRequestTag) SetValue(v string) *DescribeCapacityReservationsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCapacityReservationsRequestTag) Validate() error {
	return dara.Validate(s)
}

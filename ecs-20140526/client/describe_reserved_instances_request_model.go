// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationType(v string) *DescribeReservedInstancesRequest
	GetAllocationType() *string
	SetInstanceType(v string) *DescribeReservedInstancesRequest
	GetInstanceType() *string
	SetInstanceTypeFamily(v string) *DescribeReservedInstancesRequest
	GetInstanceTypeFamily() *string
	SetLockReason(v string) *DescribeReservedInstancesRequest
	GetLockReason() *string
	SetOfferingType(v string) *DescribeReservedInstancesRequest
	GetOfferingType() *string
	SetOwnerAccount(v string) *DescribeReservedInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeReservedInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeReservedInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReservedInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeReservedInstancesRequest
	GetRegionId() *string
	SetReservedInstanceId(v []*string) *DescribeReservedInstancesRequest
	GetReservedInstanceId() []*string
	SetReservedInstanceName(v string) *DescribeReservedInstancesRequest
	GetReservedInstanceName() *string
	SetResourceOwnerAccount(v string) *DescribeReservedInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeReservedInstancesRequest
	GetResourceOwnerId() *int64
	SetScope(v string) *DescribeReservedInstancesRequest
	GetScope() *string
	SetStatus(v []*string) *DescribeReservedInstancesRequest
	GetStatus() []*string
	SetTag(v []*DescribeReservedInstancesRequestTag) *DescribeReservedInstancesRequest
	GetTag() []*DescribeReservedInstancesRequestTag
	SetZoneId(v string) *DescribeReservedInstancesRequest
	GetZoneId() *string
}

type DescribeReservedInstancesRequest struct {
	// The allocation type. Valid values:
	//
	// - Normal: queries reserved instances under the current account.
	//
	// - Shared: queries reserved instances that have been shared between the current account and linked accounts.
	//
	// Default value: Normal.
	//
	// example:
	//
	// Normal
	AllocationType *string `json:"AllocationType,omitempty" xml:"AllocationType,omitempty"`
	// The instance type that the reserved instance can be applied to. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// > This is the instance type selected when you purchased the reserved instance. During actual deduction, region-level reserved instances support size-flexible deduction within the same instance family.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family that the reserved instance can be applied to. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The lock type. Valid values:
	//
	// - financial: The account has an overdue payment or the service has expired.
	//
	// - security: Locked for security reasons.
	//
	// example:
	//
	// security
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The payment type of the reserved instance. Valid values:
	//
	// - No Upfront: no upfront.
	//
	// - Partial Upfront: partial upfront.
	//
	// - All Upfront: all upfront.
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the reserved instance list. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the reserved instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of reserved instances. Array length: 1 to 100.
	//
	// example:
	//
	// ri-bpzhex2ulpzf53****
	ReservedInstanceId []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	// The name of the reserved instance.
	//
	// > Only exact match is supported. Fuzzy match is not supported.
	//
	// example:
	//
	// testReservedInstanceName
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scope of the reserved instance. Valid values:
	//
	//
	//
	// - Region: regional.
	//
	// - Zone: zonal.
	//
	// example:
	//
	// Region
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The statuses of reserved instances.
	//
	// example:
	//
	// Active
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags. Array length: 1 to 20.
	Tag []*DescribeReservedInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the instance. This parameter is required and takes effect only when Scope is set to Zone. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// example:
	//
	// cn-hangzhou-z
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeReservedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesRequest) GetAllocationType() *string {
	return s.AllocationType
}

func (s *DescribeReservedInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeReservedInstancesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeReservedInstancesRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeReservedInstancesRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeReservedInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeReservedInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReservedInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReservedInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReservedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReservedInstancesRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *DescribeReservedInstancesRequest) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *DescribeReservedInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReservedInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeReservedInstancesRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeReservedInstancesRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeReservedInstancesRequest) GetTag() []*DescribeReservedInstancesRequestTag {
	return s.Tag
}

func (s *DescribeReservedInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeReservedInstancesRequest) SetAllocationType(v string) *DescribeReservedInstancesRequest {
	s.AllocationType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetInstanceType(v string) *DescribeReservedInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetInstanceTypeFamily(v string) *DescribeReservedInstancesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetLockReason(v string) *DescribeReservedInstancesRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOfferingType(v string) *DescribeReservedInstancesRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOwnerAccount(v string) *DescribeReservedInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOwnerId(v int64) *DescribeReservedInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetPageNumber(v int32) *DescribeReservedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetPageSize(v int32) *DescribeReservedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetRegionId(v string) *DescribeReservedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetReservedInstanceId(v []*string) *DescribeReservedInstancesRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetReservedInstanceName(v string) *DescribeReservedInstancesRequest {
	s.ReservedInstanceName = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetResourceOwnerAccount(v string) *DescribeReservedInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetResourceOwnerId(v int64) *DescribeReservedInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetScope(v string) *DescribeReservedInstancesRequest {
	s.Scope = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetStatus(v []*string) *DescribeReservedInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetTag(v []*DescribeReservedInstancesRequestTag) *DescribeReservedInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetZoneId(v string) *DescribeReservedInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) Validate() error {
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

type DescribeReservedInstancesRequestTag struct {
	// The tag key of the reserved instance. The tag key cannot be an empty string and can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// > If you use a single tag to filter resources, the resource count with the specified tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count with all specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query resources.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the reserved instance. The tag value cannot be an empty string and can be up to 128 characters in length. It cannot start with acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReservedInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeReservedInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeReservedInstancesRequestTag) SetKey(v string) *DescribeReservedInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeReservedInstancesRequestTag) SetValue(v string) *DescribeReservedInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeReservedInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}

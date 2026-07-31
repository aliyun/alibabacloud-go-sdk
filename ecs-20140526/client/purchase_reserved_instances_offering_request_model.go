// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseReservedInstancesOfferingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *PurchaseReservedInstancesOfferingRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *PurchaseReservedInstancesOfferingRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *PurchaseReservedInstancesOfferingRequest
	GetClientToken() *string
	SetDescription(v string) *PurchaseReservedInstancesOfferingRequest
	GetDescription() *string
	SetInstanceAmount(v int32) *PurchaseReservedInstancesOfferingRequest
	GetInstanceAmount() *int32
	SetInstanceType(v string) *PurchaseReservedInstancesOfferingRequest
	GetInstanceType() *string
	SetOfferingType(v string) *PurchaseReservedInstancesOfferingRequest
	GetOfferingType() *string
	SetOwnerAccount(v string) *PurchaseReservedInstancesOfferingRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PurchaseReservedInstancesOfferingRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *PurchaseReservedInstancesOfferingRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *PurchaseReservedInstancesOfferingRequest
	GetPeriodUnit() *string
	SetPlatform(v string) *PurchaseReservedInstancesOfferingRequest
	GetPlatform() *string
	SetRegionId(v string) *PurchaseReservedInstancesOfferingRequest
	GetRegionId() *string
	SetReservedInstanceName(v string) *PurchaseReservedInstancesOfferingRequest
	GetReservedInstanceName() *string
	SetResourceGroupId(v string) *PurchaseReservedInstancesOfferingRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *PurchaseReservedInstancesOfferingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PurchaseReservedInstancesOfferingRequest
	GetResourceOwnerId() *int64
	SetScope(v string) *PurchaseReservedInstancesOfferingRequest
	GetScope() *string
	SetStartTime(v string) *PurchaseReservedInstancesOfferingRequest
	GetStartTime() *string
	SetTag(v []*PurchaseReservedInstancesOfferingRequestTag) *PurchaseReservedInstancesOfferingRequest
	GetTag() []*PurchaseReservedInstancesOfferingRequestTag
	SetZoneId(v string) *PurchaseReservedInstancesOfferingRequest
	GetZoneId() *string
}

type PurchaseReservedInstancesOfferingRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false (default): disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period, in months. This parameter takes effect only when AutoRenew is set to True.
	//
	// <props="china">Valid values of AutoRenewPeriod: 1, 12, 36, and 60.
	//
	// <props="china">If PeriodUnit is set to Month, the default value is 1.
	//
	// <props="intl">Valid values of AutoRenewPeriod: 12 and 36.
	//
	// If PeriodUnit is set to Year, the default value is 12.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the reserved instance. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// Default value: empty.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of pay-as-you-go instances of the same instance type that the reserved instance can match at the same time. Valid values: 1 to 50.
	//
	// Default value: 1.
	//
	// example:
	//
	// 3
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The instance type that the reserved instance can match.
	//
	// > Applicable instance types are continuously updated. For more information, see [Overview of reserved instances](~~100370#3c1b682051vt4~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The payment type of the reserved instance. Valid values:
	//
	//
	//
	// - No Upfront: no upfront.
	//
	// - Partial Upfront: partial upfront.
	//
	// - All Upfront: all upfront.
	//
	// Default value: All Upfront.
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The term of the reserved instance.
	//
	// <props="intl">Valid values: 1 and 3.
	//
	//
	// <props="china">
	//
	// - If `PeriodUnit` is set to `Year`, valid values are 1, 3, and 5.
	//
	// - If `PeriodUnit` is set to `Month`, the valid value is 1.
	//
	//
	//
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the term of the reserved instance.
	//
	// <props="intl">Valid values: Year.
	//
	// <props="intl">Default value: Year.
	//
	// <props="china">Valid values: Year and Month.
	//
	// <props="china">Default value: Month.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The operating system type of the image used by the instance. Valid values:
	//
	// - Windows: Windows Server operating system.
	//
	// - Linux: Linux and Unix-like operating system.
	//
	// Default value: Linux.
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the reserved instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testReservedInstanceName
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp199lyny9b3****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// Default value: Region.
	//
	// example:
	//
	// Zone
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The effective period of the reserved instance. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in UTC+0. Format: `yyyy-MM-ddTHHZ`.
	//
	// > If you do not specify this parameter, the effective period starts from the nearest time frame by default. For example, if you successfully purchase a reserved instance at 2024-11-01 13:45:35, the effective period starts at 2024-11-01 13:00:00.
	//
	// example:
	//
	// 2024-07-04T15Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags. Array length: 0 to 20.
	Tag []*PurchaseReservedInstancesOfferingRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID. This parameter is required and takes effect only when `Scope` is set to `Zone`. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s PurchaseReservedInstancesOfferingRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseReservedInstancesOfferingRequest) GoString() string {
	return s.String()
}

func (s *PurchaseReservedInstancesOfferingRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *PurchaseReservedInstancesOfferingRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *PurchaseReservedInstancesOfferingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PurchaseReservedInstancesOfferingRequest) GetDescription() *string {
	return s.Description
}

func (s *PurchaseReservedInstancesOfferingRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *PurchaseReservedInstancesOfferingRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *PurchaseReservedInstancesOfferingRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *PurchaseReservedInstancesOfferingRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PurchaseReservedInstancesOfferingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PurchaseReservedInstancesOfferingRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *PurchaseReservedInstancesOfferingRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *PurchaseReservedInstancesOfferingRequest) GetPlatform() *string {
	return s.Platform
}

func (s *PurchaseReservedInstancesOfferingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PurchaseReservedInstancesOfferingRequest) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *PurchaseReservedInstancesOfferingRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *PurchaseReservedInstancesOfferingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PurchaseReservedInstancesOfferingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PurchaseReservedInstancesOfferingRequest) GetScope() *string {
	return s.Scope
}

func (s *PurchaseReservedInstancesOfferingRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PurchaseReservedInstancesOfferingRequest) GetTag() []*PurchaseReservedInstancesOfferingRequestTag {
	return s.Tag
}

func (s *PurchaseReservedInstancesOfferingRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *PurchaseReservedInstancesOfferingRequest) SetAutoRenew(v bool) *PurchaseReservedInstancesOfferingRequest {
	s.AutoRenew = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetAutoRenewPeriod(v int32) *PurchaseReservedInstancesOfferingRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetClientToken(v string) *PurchaseReservedInstancesOfferingRequest {
	s.ClientToken = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetDescription(v string) *PurchaseReservedInstancesOfferingRequest {
	s.Description = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetInstanceAmount(v int32) *PurchaseReservedInstancesOfferingRequest {
	s.InstanceAmount = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetInstanceType(v string) *PurchaseReservedInstancesOfferingRequest {
	s.InstanceType = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetOfferingType(v string) *PurchaseReservedInstancesOfferingRequest {
	s.OfferingType = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetOwnerAccount(v string) *PurchaseReservedInstancesOfferingRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetOwnerId(v int64) *PurchaseReservedInstancesOfferingRequest {
	s.OwnerId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetPeriod(v int32) *PurchaseReservedInstancesOfferingRequest {
	s.Period = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetPeriodUnit(v string) *PurchaseReservedInstancesOfferingRequest {
	s.PeriodUnit = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetPlatform(v string) *PurchaseReservedInstancesOfferingRequest {
	s.Platform = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetRegionId(v string) *PurchaseReservedInstancesOfferingRequest {
	s.RegionId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetReservedInstanceName(v string) *PurchaseReservedInstancesOfferingRequest {
	s.ReservedInstanceName = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetResourceGroupId(v string) *PurchaseReservedInstancesOfferingRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetResourceOwnerAccount(v string) *PurchaseReservedInstancesOfferingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetResourceOwnerId(v int64) *PurchaseReservedInstancesOfferingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetScope(v string) *PurchaseReservedInstancesOfferingRequest {
	s.Scope = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetStartTime(v string) *PurchaseReservedInstancesOfferingRequest {
	s.StartTime = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetTag(v []*PurchaseReservedInstancesOfferingRequestTag) *PurchaseReservedInstancesOfferingRequest {
	s.Tag = v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) SetZoneId(v string) *PurchaseReservedInstancesOfferingRequest {
	s.ZoneId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequest) Validate() error {
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

type PurchaseReservedInstancesOfferingRequestTag struct {
	// The tag key of the reserved instance. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the reserved instance. The tag value cannot be an empty string. The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PurchaseReservedInstancesOfferingRequestTag) String() string {
	return dara.Prettify(s)
}

func (s PurchaseReservedInstancesOfferingRequestTag) GoString() string {
	return s.String()
}

func (s *PurchaseReservedInstancesOfferingRequestTag) GetKey() *string {
	return s.Key
}

func (s *PurchaseReservedInstancesOfferingRequestTag) GetValue() *string {
	return s.Value
}

func (s *PurchaseReservedInstancesOfferingRequestTag) SetKey(v string) *PurchaseReservedInstancesOfferingRequestTag {
	s.Key = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequestTag) SetValue(v string) *PurchaseReservedInstancesOfferingRequestTag {
	s.Value = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingRequestTag) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDedicatedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAttributes(v *AllocateDedicatedHostsRequestNetworkAttributes) *AllocateDedicatedHostsRequest
	GetNetworkAttributes() *AllocateDedicatedHostsRequestNetworkAttributes
	SetActionOnMaintenance(v string) *AllocateDedicatedHostsRequest
	GetActionOnMaintenance() *string
	SetAutoPlacement(v string) *AllocateDedicatedHostsRequest
	GetAutoPlacement() *string
	SetAutoReleaseTime(v string) *AllocateDedicatedHostsRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *AllocateDedicatedHostsRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *AllocateDedicatedHostsRequest
	GetAutoRenewPeriod() *int32
	SetChargeType(v string) *AllocateDedicatedHostsRequest
	GetChargeType() *string
	SetClientToken(v string) *AllocateDedicatedHostsRequest
	GetClientToken() *string
	SetCpuOverCommitRatio(v float32) *AllocateDedicatedHostsRequest
	GetCpuOverCommitRatio() *float32
	SetDedicatedHostClusterId(v string) *AllocateDedicatedHostsRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostName(v string) *AllocateDedicatedHostsRequest
	GetDedicatedHostName() *string
	SetDedicatedHostType(v string) *AllocateDedicatedHostsRequest
	GetDedicatedHostType() *string
	SetDescription(v string) *AllocateDedicatedHostsRequest
	GetDescription() *string
	SetMinQuantity(v int32) *AllocateDedicatedHostsRequest
	GetMinQuantity() *int32
	SetOwnerAccount(v string) *AllocateDedicatedHostsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateDedicatedHostsRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *AllocateDedicatedHostsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *AllocateDedicatedHostsRequest
	GetPeriodUnit() *string
	SetQuantity(v int32) *AllocateDedicatedHostsRequest
	GetQuantity() *int32
	SetRegionId(v string) *AllocateDedicatedHostsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AllocateDedicatedHostsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AllocateDedicatedHostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateDedicatedHostsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*AllocateDedicatedHostsRequestTag) *AllocateDedicatedHostsRequest
	GetTag() []*AllocateDedicatedHostsRequestTag
	SetZoneId(v string) *AllocateDedicatedHostsRequest
	GetZoneId() *string
}

type AllocateDedicatedHostsRequest struct {
	NetworkAttributes *AllocateDedicatedHostsRequestNetworkAttributes `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
	// The policy used to migrate the instances deployed on the dedicated host when the dedicated host fails or needs to be repaired online. Valid values:
	//
	// - Migrate: The instances are migrated to another physical server and restarted.
	//
	//   Default value when cloud disks are attached to the dedicated host: Migrate.
	//
	// - Stop: The instances are stopped on the current dedicated host. After the dedicated host is confirmed to be irreparable, the instances are migrated to another physical server and restarted.
	//
	//   Default value when local disks are attached to the dedicated host: Stop.
	//
	// example:
	//
	// Migrate
	ActionOnMaintenance *string `json:"ActionOnMaintenance,omitempty" xml:"ActionOnMaintenance,omitempty"`
	// Specifies whether to add the dedicated host to the automatic deployment resource pool. If you create an instance on a dedicated host without specifying **DedicatedHostId**, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. For more information, see [Automatic deployment](https://help.aliyun.com/document_detail/118938.html). Valid values:
	//
	// - on: adds the dedicated host to the automatic deployment resource pool.
	//
	// - off: does not add the dedicated host to the automatic deployment resource pool.
	//
	// Default value: on.
	//
	// > If you do not want the dedicated host to be added to the automatic deployment resource pool, set this parameter to off.
	//
	// example:
	//
	// off
	AutoPlacement *string `json:"AutoPlacement,omitempty" xml:"AutoPlacement,omitempty"`
	// The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC+0.
	//
	// > - The earliest release time must be at least half an hour from the current time.
	//
	// > - The latest release time must be at most three years from the current time.
	//
	// > - If the value of seconds (ss) is not 00, it is automatically set to 00.
	//
	// example:
	//
	// 2019-08-21T12:30:24Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal for the subscription dedicated host.
	//
	// > The **AutoRenew*	- parameter takes effect only when **ChargeType*	- is set to PrePaid.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. The **AutoRenewPeriod*	- parameter takes effect and is required only when **AutoRenew*	- is set to true. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The billing method of the dedicated host. Valid values:
	//
	// <props="china">
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your payment method supports balance payment or credit payment. Otherwise, the `InvalidPayMethod` error is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	//
	//
	//
	// <props="intl">
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your payment method supports credit payment. Otherwise, the `InvalidPayMethod` error is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	//
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The CPU overcommit ratio. Only the custom instance types g6s, c6s, and r6s support the CPU overcommit ratio. Valid values: 1 to 5.
	//
	// The CPU overcommit ratio affects the number of available vCPUs on a dedicated host. The number of available vCPUs on a dedicated host = Number of physical CPU cores × 2 × CPU overcommit ratio. For example, the number of physical CPU cores on each g6s host is 52. If you set the CPU overcommit ratio to 4, the total number of vCPUs on the dedicated host is 416. For scenarios that do not require strict CPU stability or have low CPU loads, such as development and testing environments, you can increase the CPU overcommit ratio to increase the number of available vCPUs and deploy more ECS instances of the same specifications, which reduces the unit deployment cost.
	//
	// example:
	//
	// 1
	CpuOverCommitRatio *float32 `json:"CpuOverCommitRatio,omitempty" xml:"CpuOverCommitRatio,omitempty"`
	// The ID of the dedicated host cluster to which the dedicated host belongs.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The name of the dedicated host. The name must be 2 to 128 characters in length and can contain Unicode characters under the Letter category, which includes characters from various scripts such as English, Chinese, and digits. The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// myDDH
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The type of the dedicated host. You can call [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) to query the most recent list of dedicated host types.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddh.c5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This-is-my-DDH
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The minimum number of dedicated hosts to create. Valid values: 1 to 100.
	//
	// > If the active stock of dedicated hosts is less than the minimum number, the dedicated host creation fails.
	//
	// example:
	//
	// 2
	MinQuantity  *int32  `json:"MinQuantity,omitempty" xml:"MinQuantity,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the dedicated host. The `Period` parameter takes effect and is required only when `ChargeType` is set to `PrePaid`. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	//
	//
	// <props="intl">
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// <props="china">
	//
	// - Week
	//
	// - Month
	//
	// - Year
	//
	//
	//
	// <props="intl">
	//
	// - Month
	//
	// - Year
	//
	//
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of dedicated hosts to create. Valid values: 1 to 100.
	//
	// Default value: 1.
	//
	// example:
	//
	// 2
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID of the dedicated host. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph***
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*AllocateDedicatedHostsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the dedicated host.
	//
	// Default value: empty, which indicates that the system selects a zone.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AllocateDedicatedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsRequest) GetNetworkAttributes() *AllocateDedicatedHostsRequestNetworkAttributes {
	return s.NetworkAttributes
}

func (s *AllocateDedicatedHostsRequest) GetActionOnMaintenance() *string {
	return s.ActionOnMaintenance
}

func (s *AllocateDedicatedHostsRequest) GetAutoPlacement() *string {
	return s.AutoPlacement
}

func (s *AllocateDedicatedHostsRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *AllocateDedicatedHostsRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *AllocateDedicatedHostsRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *AllocateDedicatedHostsRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *AllocateDedicatedHostsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateDedicatedHostsRequest) GetCpuOverCommitRatio() *float32 {
	return s.CpuOverCommitRatio
}

func (s *AllocateDedicatedHostsRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *AllocateDedicatedHostsRequest) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *AllocateDedicatedHostsRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *AllocateDedicatedHostsRequest) GetDescription() *string {
	return s.Description
}

func (s *AllocateDedicatedHostsRequest) GetMinQuantity() *int32 {
	return s.MinQuantity
}

func (s *AllocateDedicatedHostsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateDedicatedHostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateDedicatedHostsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *AllocateDedicatedHostsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *AllocateDedicatedHostsRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *AllocateDedicatedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateDedicatedHostsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateDedicatedHostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateDedicatedHostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateDedicatedHostsRequest) GetTag() []*AllocateDedicatedHostsRequestTag {
	return s.Tag
}

func (s *AllocateDedicatedHostsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AllocateDedicatedHostsRequest) SetNetworkAttributes(v *AllocateDedicatedHostsRequestNetworkAttributes) *AllocateDedicatedHostsRequest {
	s.NetworkAttributes = v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetActionOnMaintenance(v string) *AllocateDedicatedHostsRequest {
	s.ActionOnMaintenance = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetAutoPlacement(v string) *AllocateDedicatedHostsRequest {
	s.AutoPlacement = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetAutoReleaseTime(v string) *AllocateDedicatedHostsRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetAutoRenew(v bool) *AllocateDedicatedHostsRequest {
	s.AutoRenew = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetAutoRenewPeriod(v int32) *AllocateDedicatedHostsRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetChargeType(v string) *AllocateDedicatedHostsRequest {
	s.ChargeType = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetClientToken(v string) *AllocateDedicatedHostsRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetCpuOverCommitRatio(v float32) *AllocateDedicatedHostsRequest {
	s.CpuOverCommitRatio = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetDedicatedHostClusterId(v string) *AllocateDedicatedHostsRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetDedicatedHostName(v string) *AllocateDedicatedHostsRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetDedicatedHostType(v string) *AllocateDedicatedHostsRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetDescription(v string) *AllocateDedicatedHostsRequest {
	s.Description = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetMinQuantity(v int32) *AllocateDedicatedHostsRequest {
	s.MinQuantity = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetOwnerAccount(v string) *AllocateDedicatedHostsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetOwnerId(v int64) *AllocateDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetPeriod(v int32) *AllocateDedicatedHostsRequest {
	s.Period = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetPeriodUnit(v string) *AllocateDedicatedHostsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetQuantity(v int32) *AllocateDedicatedHostsRequest {
	s.Quantity = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetRegionId(v string) *AllocateDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetResourceGroupId(v string) *AllocateDedicatedHostsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetResourceOwnerAccount(v string) *AllocateDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetResourceOwnerId(v int64) *AllocateDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetTag(v []*AllocateDedicatedHostsRequestTag) *AllocateDedicatedHostsRequest {
	s.Tag = v
	return s
}

func (s *AllocateDedicatedHostsRequest) SetZoneId(v string) *AllocateDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *AllocateDedicatedHostsRequest) Validate() error {
	if s.NetworkAttributes != nil {
		if err := s.NetworkAttributes.Validate(); err != nil {
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

type AllocateDedicatedHostsRequestNetworkAttributes struct {
	// The timeout period of a UDP session for load balancing connections to the dedicated host. Unit: seconds. Valid values: 15 to 310.
	//
	// example:
	//
	// 60
	SlbUdpTimeout *int32 `json:"SlbUdpTimeout,omitempty" xml:"SlbUdpTimeout,omitempty"`
	// The timeout period of a UDP session between a user and a cloud service running on the dedicated host. Unit: seconds. Valid values: 15 to 310.
	//
	// example:
	//
	// 60
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s AllocateDedicatedHostsRequestNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsRequestNetworkAttributes) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsRequestNetworkAttributes) GetSlbUdpTimeout() *int32 {
	return s.SlbUdpTimeout
}

func (s *AllocateDedicatedHostsRequestNetworkAttributes) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *AllocateDedicatedHostsRequestNetworkAttributes) SetSlbUdpTimeout(v int32) *AllocateDedicatedHostsRequestNetworkAttributes {
	s.SlbUdpTimeout = &v
	return s
}

func (s *AllocateDedicatedHostsRequestNetworkAttributes) SetUdpTimeout(v int32) *AllocateDedicatedHostsRequestNetworkAttributes {
	s.UdpTimeout = &v
	return s
}

func (s *AllocateDedicatedHostsRequestNetworkAttributes) Validate() error {
	return dara.Validate(s)
}

type AllocateDedicatedHostsRequestTag struct {
	// The tag key of the dedicated host. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// Environment
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated host. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// Production
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AllocateDedicatedHostsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsRequestTag) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsRequestTag) GetKey() *string {
	return s.Key
}

func (s *AllocateDedicatedHostsRequestTag) GetValue() *string {
	return s.Value
}

func (s *AllocateDedicatedHostsRequestTag) SetKey(v string) *AllocateDedicatedHostsRequestTag {
	s.Key = &v
	return s
}

func (s *AllocateDedicatedHostsRequestTag) SetValue(v string) *AllocateDedicatedHostsRequestTag {
	s.Value = &v
	return s
}

func (s *AllocateDedicatedHostsRequestTag) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetAutoRenew() *bool
	SetAutoRenewWithEcs(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetAutoRenewWithEcs() *string
	SetDedicatedHostIds(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetDedicatedHostIds() *string
	SetDuration(v int32) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetDuration() *int32
	SetOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriodUnit(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostAutoRenewAttributeRequest struct {
	// Specifies whether to enable auto-renewal for the subscription dedicated host. Valid values:
	//
	// - true: Enables auto-renewal for the subscription dedicated host.
	//
	// - false: Disables auto-renewal for the subscription dedicated host.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to enable auto-renewal for the dedicated host to follow the subscription ECS instances on the host.
	//
	// If your dedicated host (DDH) uses the subscription billing method and the subscription ECS instances on the DDH have auto-renewal enabled, you can use this parameter to configure the DDH to automatically renew along with the ECS instances. When an ECS instance on the DDH is automatically renewed, if the DDH expires earlier than the new expiration time of the ECS instance, the DDH is also automatically renewed. The principle of DDH auto-renewal following ECS instances is as follows:
	//
	// The DDH automatically determines the new expiration time of the corresponding ECS instance, and then selects the minimum renewal period that is greater than the ECS instance expiration time and meets the DDH renewal cycle. For details about the supported renewal cycles of DDHs, see the metric descriptions of the PeriodUnit and Duration parameters.
	//
	// Example: A subscription DDH expires on January 15 of the current year. After a subscription ECS instance on the DDH is automatically renewed, the ECS instance expiration is extended to November 15 of the current year. The DDH lifecycle is 10 months shorter than the ECS instance lifecycle. In this case, the DDH selects the minimum renewal period that is greater than 10 months and meets the DDH renewal cycle, which is 12 months (PeriodUnit=Month and Duration=12).
	//
	// Valid values:
	//
	// - AutoRenewWithEcs: Enables auto-renewal following the subscription ECS instances on the dedicated host.
	//
	// - StopRenewWithEcs: Disables auto-renewal following the subscription ECS instances on the dedicated host.
	//
	// - NoOperation: Does not change the current settings of the dedicated host.
	//
	// > If you set this parameter to AutoRenewWithEcs, make sure that auto-renewal is enabled for the dedicated host (AutoRenew=true). Otherwise, this parameter only changes the parameter value, and the actual auto-renewal feature following ECS instances does not take effect.
	//
	// Default value: NoOperation.
	//
	// example:
	//
	// StopRenewWithEcs
	AutoRenewWithEcs *string `json:"AutoRenewWithEcs,omitempty" xml:"AutoRenewWithEcs,omitempty"`
	// The IDs of dedicated hosts. You can specify up to 100 subscription dedicated host IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	// The renewal period. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	//
	//
	// <props="intl">
	//
	// - If PeriodUnit is set to Month: 1 and 12.
	//
	// - If PeriodUnit is set to Year: 1 and 12.
	//
	// example:
	//
	// 1
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal period. Valid values:
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
	// The region ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable auto-renewal for the subscription dedicated host. The RenewalStatus parameter takes precedence over the AutoRenew parameter. Valid values:
	//
	// - AutoRenewal: Enables auto-renewal.
	//
	// - Normal: Disables auto-renewal but the system still sends expiration notifications.
	//
	// - NotRenewal: Disables auto-renewal and the system does not send expiration notifications. Three days before expiration, the system automatically sends a non-renewal notification. You can change the value of this parameter to Normal for a dedicated host, and then manually renew the host by calling [RenewDedicatedHosts](https://help.aliyun.com/document_detail/134250.html) or set the value to AutoRenewal to enable auto-renewal.
	//
	// example:
	//
	// Normal
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetAutoRenewWithEcs() *string {
	return s.AutoRenewWithEcs
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetAutoRenew(v bool) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetAutoRenewWithEcs(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.AutoRenewWithEcs = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetDedicatedHostIds(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetDuration(v int32) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}

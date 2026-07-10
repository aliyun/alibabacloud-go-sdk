// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoProvisioningGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchConfiguration(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) *CreateAutoProvisioningGroupShrinkRequest
	GetLaunchConfiguration() *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration
	SetAutoProvisioningGroupName(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetAutoProvisioningGroupName() *string
	SetAutoProvisioningGroupType(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetAutoProvisioningGroupType() *string
	SetCandidateOptions(v *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) *CreateAutoProvisioningGroupShrinkRequest
	GetCandidateOptions() *CreateAutoProvisioningGroupShrinkRequestCandidateOptions
	SetClientToken(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetClientToken() *string
	SetDataDiskConfig(v []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) *CreateAutoProvisioningGroupShrinkRequest
	GetDataDiskConfig() []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig
	SetDefaultTargetCapacityType(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetDefaultTargetCapacityType() *string
	SetDescription(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetDescription() *string
	SetExcessCapacityTerminationPolicy(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetExcessCapacityTerminationPolicy() *string
	SetExecutionMode(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetExecutionMode() *string
	SetHibernationOptionsConfigured(v bool) *CreateAutoProvisioningGroupShrinkRequest
	GetHibernationOptionsConfigured() *bool
	SetLaunchTemplateConfig(v []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) *CreateAutoProvisioningGroupShrinkRequest
	GetLaunchTemplateConfig() []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig
	SetLaunchTemplateId(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersion(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetLaunchTemplateVersion() *string
	SetMaxSpotPrice(v float32) *CreateAutoProvisioningGroupShrinkRequest
	GetMaxSpotPrice() *float32
	SetMinTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetMinTargetCapacity() *string
	SetOwnerAccount(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAutoProvisioningGroupShrinkRequest
	GetOwnerId() *int64
	SetPayAsYouGoAllocationStrategy(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetPayAsYouGoAllocationStrategy() *string
	SetPayAsYouGoTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetPayAsYouGoTargetCapacity() *string
	SetPrePaidOptions(v *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) *CreateAutoProvisioningGroupShrinkRequest
	GetPrePaidOptions() *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions
	SetRegionId(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoProvisioningGroupShrinkRequest
	GetResourceOwnerId() *int64
	SetResourcePoolOptionsShrink(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetResourcePoolOptionsShrink() *string
	SetSpotAllocationStrategy(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetSpotAllocationStrategy() *string
	SetSpotInstanceInterruptionBehavior(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetSpotInstanceInterruptionBehavior() *string
	SetSpotInstancePoolsToUseCount(v int32) *CreateAutoProvisioningGroupShrinkRequest
	GetSpotInstancePoolsToUseCount() *int32
	SetSpotTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetSpotTargetCapacity() *string
	SetSystemDiskConfig(v []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) *CreateAutoProvisioningGroupShrinkRequest
	GetSystemDiskConfig() []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig
	SetTag(v []*CreateAutoProvisioningGroupShrinkRequestTag) *CreateAutoProvisioningGroupShrinkRequest
	GetTag() []*CreateAutoProvisioningGroupShrinkRequestTag
	SetTerminateInstances(v bool) *CreateAutoProvisioningGroupShrinkRequest
	GetTerminateInstances() *bool
	SetTerminateInstancesWithExpiration(v bool) *CreateAutoProvisioningGroupShrinkRequest
	GetTerminateInstancesWithExpiration() *bool
	SetTotalTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetTotalTargetCapacity() *string
	SetValidFrom(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetValidFrom() *string
	SetValidUntil(v string) *CreateAutoProvisioningGroupShrinkRequest
	GetValidUntil() *string
}

type CreateAutoProvisioningGroupShrinkRequest struct {
	LaunchConfiguration *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration `json:"LaunchConfiguration,omitempty" xml:"LaunchConfiguration,omitempty" type:"Struct"`
	// The name of the auto provisioning group. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// apg-test
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The delivery type of the auto provisioning group. Valid values:
	//
	// - request: one-time asynchronous delivery. The group delivers an instance cluster asynchronously only at startup. If scheduling fails, no retry is performed.
	//
	// - instant: one-time synchronous delivery. The group synchronously creates instances only at startup and returns the list of successfully created instances and the causes of creation failures in the response.
	//
	// - maintain: continuous delivery. The group attempts to deliver an instance cluster at startup and monitors real-time capacity. If the target capacity is not reached, the group continues to create ECS instances.
	//
	// Default value: maintain.
	//
	// example:
	//
	// maintain
	AutoProvisioningGroupType *string                                                   `json:"AutoProvisioningGroupType,omitempty" xml:"AutoProvisioningGroupType,omitempty"`
	CandidateOptions          *CreateAutoProvisioningGroupShrinkRequestCandidateOptions `json:"CandidateOptions,omitempty" xml:"CandidateOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of data disk configurations.
	DataDiskConfig []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig `json:"DataDiskConfig,omitempty" xml:"DataDiskConfig,omitempty" type:"Repeated"`
	// The billing method for the capacity difference when the sum of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is less than `TotalTargetCapacity`. Valid values:
	//
	// - PayAsYouGo: pay-as-you-go instances.
	//
	// - Spot: spot instances.
	//
	// Default value: Spot.
	//
	// example:
	//
	// Spot
	DefaultTargetCapacityType *string `json:"DefaultTargetCapacityType,omitempty" xml:"DefaultTargetCapacityType,omitempty"`
	// The description of the auto provisioning group.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to release instances when the real-time capacity of the auto provisioning group exceeds the target capacity and a scale-in event is triggered. Valid values:
	//
	// - termination: releases the scaled-in instances.
	//
	// - no-termination: only removes the scaled-in instances from the auto provisioning group.
	//
	// Default value: no-termination.
	//
	// example:
	//
	// termination
	ExcessCapacityTerminationPolicy *string `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	ExecutionMode                   *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	HibernationOptionsConfigured *bool `json:"HibernationOptionsConfigured,omitempty" xml:"HibernationOptionsConfigured,omitempty"`
	// The list of extended launch templates.
	LaunchTemplateConfig []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig `json:"LaunchTemplateConfig,omitempty" xml:"LaunchTemplateConfig,omitempty" type:"Repeated"`
	// The ID of the instance launch template associated with the auto provisioning group. You can invoke [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) to query active instance launch templates. If you specify both a launch template and launch configuration information (`LaunchConfiguration.*`), the launch template takes precedence.
	//
	// example:
	//
	// lt-bp1fgzds4bdogu03****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the instance launch template associated with the auto provisioning group. You can invoke [DescribeLaunchTemplateVersions](https://help.aliyun.com/document_detail/73761.html) to query active instance launch template versions.
	//
	// Default value: the default version of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum price for spot instances in the auto provisioning group.
	//
	// > If both `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` are specified, the lower value is used.
	//
	// example:
	//
	// 2
	MaxSpotPrice *float32 `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	// The minimum target capacity of the auto provisioning group. Valid values
	//
	// example:
	//
	// 20
	MinTargetCapacity *string `json:"MinTargetCapacity,omitempty" xml:"MinTargetCapacity,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for creating pay-as-you-go instances. Valid values:
	//
	// - lowest-price: cost optimization policy. Selects the instance type with the lowest price.
	//
	// - prioritized: priority-based policy. Creates instances based on the priority specified by `LaunchTemplateConfig.N.Priority`.
	//
	// Default value: lowest-price.
	//
	// example:
	//
	// prioritized
	PayAsYouGoAllocationStrategy *string `json:"PayAsYouGoAllocationStrategy,omitempty" xml:"PayAsYouGoAllocationStrategy,omitempty"`
	// The target capacity of pay-as-you-go instances in the auto provisioning group. Valid values: less than or equal to the parameter value of `TotalTargetCapacity`.
	//
	// example:
	//
	// 30
	PayAsYouGoTargetCapacity *string `json:"PayAsYouGoTargetCapacity,omitempty" xml:"PayAsYouGoTargetCapacity,omitempty"`
	// The detailed capacity configuration for subscription instances.
	PrePaidOptions *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions `json:"PrePaidOptions,omitempty" xml:"PrePaidOptions,omitempty" type:"Struct"`
	// The ID of the region in which the auto provisioning group resides. You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the auto provisioning group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource pool policy used to create instances. After you set this parameter, note the following items:
	//
	// - This parameter takes effect only when you create pay-as-you-go instances.
	//
	// - This parameter takes effect only when you create a one-time synchronization delivery auto provisioning group (`AutoProvisioningGroupType=instant`).
	ResourcePoolOptionsShrink *string `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty"`
	// The policy for creating spot instances. Valid values:
	//
	// - lowest-price: cost optimization policy. Selects the instance type with the lowest price.
	//
	// - diversified: balanced zone distribution policy. Creates instances in the zones specified in the extended launch template and evenly distributes them across zones.
	//
	// - capacity-optimized: capacity optimization distribution policy. Selects the optimal instance type and zone based on inventory availability.
	//
	// Default value: lowest-price.
	//
	// example:
	//
	// diversified
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The action to take when a spot instance is interrupted. Valid values:
	//
	// - stop: stops the instance.
	//
	// - terminate: releases the instance.
	//
	// Default value: terminate.
	//
	// example:
	//
	// terminate
	SpotInstanceInterruptionBehavior *string `json:"SpotInstanceInterruptionBehavior,omitempty" xml:"SpotInstanceInterruptionBehavior,omitempty"`
	// Takes effect when `SpotAllocationStrategy` is set to `lowest-price`. Specifies the number of instance types with the lowest prices from which the auto provisioning group creates instances.
	//
	// Valid values: less than the value of N in `LaunchTemplateConfig.N`.
	//
	// example:
	//
	// 2
	SpotInstancePoolsToUseCount *int32 `json:"SpotInstancePoolsToUseCount,omitempty" xml:"SpotInstancePoolsToUseCount,omitempty"`
	// The target capacity of spot instances in the auto provisioning group. Valid values: less than or equal to the parameter value of `TotalTargetCapacity`.
	//
	// example:
	//
	// 20
	SpotTargetCapacity *string `json:"SpotTargetCapacity,omitempty" xml:"SpotTargetCapacity,omitempty"`
	// The list of system disk configurations.
	SystemDiskConfig []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig `json:"SystemDiskConfig,omitempty" xml:"SystemDiskConfig,omitempty" type:"Repeated"`
	// The tags to attach to the auto provisioning group.
	Tag []*CreateAutoProvisioningGroupShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to release instances auto provisioning group when the auto-provisioning group is deleted. Valid values:
	//
	// - true: releases instances auto provisioning group.
	//
	// - false: retains instances auto provisioning group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstances *bool `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
	// Specifies whether to release instances auto provisioning group when the auto-provisioning group expires. Valid values:
	//
	// - true: releases instances auto provisioning group.
	//
	// - false: only removes instances from the auto-provisioning group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstancesWithExpiration *bool `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	// The total target capacity of the auto provisioning group. Valid values: positive integers.
	//
	// The total capacity must be greater than or equal to the sum of `PayAsYouGoTargetCapacity` (the target capacity of pay-as-you-go instances) and `SpotTargetCapacity` (the target capacity of spot instances).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	TotalTargetCapacity *string `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
	// The time when the auto provisioning group is started. Used together with `ValidUntil` to determine the valid period.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// Default value: the UNIX timestamp at which the request takes effect immediately.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	ValidFrom *string `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The time when the auto provisioning group expires. Used together with `ValidFrom` to determine the valid period.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// Default value: 2099-12-31T23:59:59Z.
	//
	// example:
	//
	// 2019-06-01T15:10:20Z
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetLaunchConfiguration() *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	return s.LaunchConfiguration
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetAutoProvisioningGroupName() *string {
	return s.AutoProvisioningGroupName
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetAutoProvisioningGroupType() *string {
	return s.AutoProvisioningGroupType
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetCandidateOptions() *CreateAutoProvisioningGroupShrinkRequestCandidateOptions {
	return s.CandidateOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetDataDiskConfig() []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig {
	return s.DataDiskConfig
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetDefaultTargetCapacityType() *string {
	return s.DefaultTargetCapacityType
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetExcessCapacityTerminationPolicy() *string {
	return s.ExcessCapacityTerminationPolicy
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetHibernationOptionsConfigured() *bool {
	return s.HibernationOptionsConfigured
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetLaunchTemplateConfig() []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	return s.LaunchTemplateConfig
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetMaxSpotPrice() *float32 {
	return s.MaxSpotPrice
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetMinTargetCapacity() *string {
	return s.MinTargetCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetPayAsYouGoAllocationStrategy() *string {
	return s.PayAsYouGoAllocationStrategy
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetPayAsYouGoTargetCapacity() *string {
	return s.PayAsYouGoTargetCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetPrePaidOptions() *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions {
	return s.PrePaidOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetResourcePoolOptionsShrink() *string {
	return s.ResourcePoolOptionsShrink
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetSpotInstanceInterruptionBehavior() *string {
	return s.SpotInstanceInterruptionBehavior
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetSpotInstancePoolsToUseCount() *int32 {
	return s.SpotInstancePoolsToUseCount
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetSpotTargetCapacity() *string {
	return s.SpotTargetCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetSystemDiskConfig() []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig {
	return s.SystemDiskConfig
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetTag() []*CreateAutoProvisioningGroupShrinkRequestTag {
	return s.Tag
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetTerminateInstances() *bool {
	return s.TerminateInstances
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetTerminateInstancesWithExpiration() *bool {
	return s.TerminateInstancesWithExpiration
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetTotalTargetCapacity() *string {
	return s.TotalTargetCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetValidFrom() *string {
	return s.ValidFrom
}

func (s *CreateAutoProvisioningGroupShrinkRequest) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetLaunchConfiguration(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) *CreateAutoProvisioningGroupShrinkRequest {
	s.LaunchConfiguration = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetAutoProvisioningGroupName(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.AutoProvisioningGroupName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetAutoProvisioningGroupType(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.AutoProvisioningGroupType = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetCandidateOptions(v *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) *CreateAutoProvisioningGroupShrinkRequest {
	s.CandidateOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetClientToken(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetDataDiskConfig(v []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) *CreateAutoProvisioningGroupShrinkRequest {
	s.DataDiskConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetDefaultTargetCapacityType(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.DefaultTargetCapacityType = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetDescription(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetExcessCapacityTerminationPolicy(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ExcessCapacityTerminationPolicy = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetExecutionMode(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetHibernationOptionsConfigured(v bool) *CreateAutoProvisioningGroupShrinkRequest {
	s.HibernationOptionsConfigured = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetLaunchTemplateConfig(v []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) *CreateAutoProvisioningGroupShrinkRequest {
	s.LaunchTemplateConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetLaunchTemplateId(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetLaunchTemplateVersion(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetMaxSpotPrice(v float32) *CreateAutoProvisioningGroupShrinkRequest {
	s.MaxSpotPrice = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetMinTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.MinTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetOwnerAccount(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetOwnerId(v int64) *CreateAutoProvisioningGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetPayAsYouGoAllocationStrategy(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.PayAsYouGoAllocationStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetPayAsYouGoTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.PayAsYouGoTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetPrePaidOptions(v *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) *CreateAutoProvisioningGroupShrinkRequest {
	s.PrePaidOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetRegionId(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetResourceGroupId(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetResourceOwnerAccount(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetResourceOwnerId(v int64) *CreateAutoProvisioningGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetResourcePoolOptionsShrink(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ResourcePoolOptionsShrink = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetSpotAllocationStrategy(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetSpotInstanceInterruptionBehavior(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.SpotInstanceInterruptionBehavior = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetSpotInstancePoolsToUseCount(v int32) *CreateAutoProvisioningGroupShrinkRequest {
	s.SpotInstancePoolsToUseCount = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetSpotTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.SpotTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetSystemDiskConfig(v []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) *CreateAutoProvisioningGroupShrinkRequest {
	s.SystemDiskConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetTag(v []*CreateAutoProvisioningGroupShrinkRequestTag) *CreateAutoProvisioningGroupShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetTerminateInstances(v bool) *CreateAutoProvisioningGroupShrinkRequest {
	s.TerminateInstances = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetTerminateInstancesWithExpiration(v bool) *CreateAutoProvisioningGroupShrinkRequest {
	s.TerminateInstancesWithExpiration = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetTotalTargetCapacity(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.TotalTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetValidFrom(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ValidFrom = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) SetValidUntil(v string) *CreateAutoProvisioningGroupShrinkRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequest) Validate() error {
	if s.LaunchConfiguration != nil {
		if err := s.LaunchConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.CandidateOptions != nil {
		if err := s.CandidateOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DataDiskConfig != nil {
		for _, item := range s.DataDiskConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LaunchTemplateConfig != nil {
		for _, item := range s.LaunchTemplateConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrePaidOptions != nil {
		if err := s.PrePaidOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDiskConfig != nil {
		for _, item := range s.SystemDiskConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration struct {
	// > This parameter is in invitational preview and is not available for use.
	Arn []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// The automatic release time of the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// - If the value of seconds (`ss`) is not `00`, the time is automatically rounded down to the start of the current minute (`mm`).
	//
	// - The earliest release time is 30 minutes after the current time.
	//
	// - The latest release time cannot be more than three years from the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The running mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information about instance performance, see the performance constrained mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information about instance performance, see the unlimited mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// Default value: none.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disk configurations in the launch configuration.
	DataDisk []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the instance. The following limits apply:
	//
	// - Periods (.) and hyphens (-) cannot be used as the first or last characters and cannot be used consecutively.
	//
	// - Windows instances: The hostname must be 2 to 15 characters in length and cannot contain periods (.) or consist entirely of digits. It can contain letters, digits, and hyphens (-).
	//
	// - Instances of other types (such as Linux): The hostname must be 2 to 64 characters in length and can contain multiple periods (.). Each segment between periods can contain letters, digits, and hyphens (-).
	//
	// - You cannot specify both `LaunchConfiguration.HostName` and `LaunchConfiguration.HostNames.N`. Otherwise, an error is returned.
	//
	// - If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// k8s-node-[1,4]-ecshost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The list of hostnames for one or more instances. The following limits apply:
	//
	// - This parameter takes effect only when you create a one-time synchronous delivery auto provisioning group (`AutoProvisioningGroupType=instant`).
	//
	// - N indicates the number of instances. Valid values of N: 1 to 1000. The value must be consistent with the TotalTargetCapacity parameter.
	//
	// - Periods (.) and hyphens (-) cannot be used as the first or last characters and cannot be used consecutively.
	//
	// - If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// ecs-host-01
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The name of the image family. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `aliyun` or `acs:`. The name cannot contain `http://` or `https://`. The name can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image used to launch instances. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// m-bp1g7004ksh0oeuc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// Instance_Description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// Default value: the `InstanceId` of the instance.
	//
	// When you create multiple ECS instances, you can batch configure sequential instance names. For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// > In pay-by-traffic mode, the peak inbound and outbound bandwidths are used as upper limits of bandwidths instead of guaranteed performance metrics. When resources are contended, the peak bandwidths may be limited. If you require guaranteed bandwidth, use pay-by-bandwidth.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// - If the maximum outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - If the maximum outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `LaunchConfiguration.InternetMaxBandwidthOut`. Default value: the value of `LaunchConfiguration.InternetMaxBandwidthOut`.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. Valid values:
	//
	// - none: non-I/O optimization.
	//
	// - optimized: I/O optimization.
	//
	// For retired instance types, the default value is none. For other instance types, the default value is optimized.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// -   For Windows instances, this parameter is ignored and is empty by default.
	//
	// -   For Linux instances, password-based logon is disabled during initialization.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// ``()`~!@#$%^&*-_+=|{}`[]`:;\\"<>,.?/``
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. Valid values:
	//
	// - true: uses the preset password.
	//
	// - false: does not use the preset password.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// true
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the instance RAM role. You can call the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// RAM_Name
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the instance belongs. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// -   Active: enables security hardening. This value is applicable only to public images.
	//
	// -   Deactive: disables security hardening. This value is applicable to all image types.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the instance belongs. If both a launch template and launch configuration information are specified, the launch template takes precedence.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The list of security groups to which the instance belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The system disk information of the instance. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	SystemDisk *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The category of the system disk. Valid values:
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD (ESSD).
	//
	// -   cloud: basic disk.
	//
	// For retired instance types that are non-I/O optimization instances, the default value is cloud. For other instance types, the default value is cloud_efficiency.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// SystemDisk_Description
	SystemDiskDescription *string `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// Default value: empty.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssdSystem
	SystemDiskName *string `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	// The performance level (PL) of the enterprise SSD used as the system disk. Valid values:
	//
	// - PL0 (default): a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// PL0
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values: 20 to 500. The value of this parameter must be greater than or equal to max{20, size of the image specified by LaunchConfiguration.ImageId}.
	//
	// Default value: max{40, size of the image specified by LaunchConfiguration.ImageId}.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The list of tags in the launch configuration.
	Tag []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Instance user data of the instance. Instance user data must be Base64-encoded. The maximum size of the raw data is 32 KB. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect when you create subscription instances. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false (default): does not enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	//
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The CPU-related configurations.
	CpuOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The image-related property information.
	//
	// After you set this parameter, note the following items:
	//
	// - This parameter takes effect only when you create a one-time synchronization delivery auto provisioning group (AutoProvisioningGroupType=instant).
	ImageOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The subscription duration of the resource. Unit: specified by `PeriodUnit`. This parameter is required when you create subscription instances. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week, valid values of Period: 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month, valid values of Period: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month, valid values of Period: 1, 2, 3, 6, and 12.
	//
	// <props="partner">If PeriodUnit is set to Month, valid values of Period: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription billable methods duration. Valid values:
	//
	// <props="china">
	//
	// - Week
	//
	// - Month (default)
	//
	//
	//
	// <props="intl">Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit       *string                                                                      `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SchedulerOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SecurityOptions  *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions  `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not subject to automatic release within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the marketplace price and checks the resource inventory to determine whether to retain or revoke the instance.
	//
	// - 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the marketplace price and checks the resource inventory to determine whether to retain or revoke the instance.
	//
	// Alibaba Cloud sends an ECS system event notification 5 minutes before the instance is released. Spot instances are billed by second. Select an appropriate protection period based on the expected task execution duration.
	//
	// After you set this parameter, note the following items:
	//
	// - This parameter takes effect only when you create a one-time synchronization delivery auto provisioning group (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The spot instance break mode. Valid values:
	//
	// - Terminate: directly releases the instance.
	//
	// - Stop: puts the instance into economical mode.
	//
	// For more information about economical mode, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// After you set this parameter, note the following items:
	//
	// - This parameter takes effect only when you create a one-time synchronization delivery auto provisioning group (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetArn() []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn {
	return s.Arn
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetDataDisk() []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	return s.DataDisk
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetHostName() *string {
	return s.HostName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetHostNames() []*string {
	return s.HostNames
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetPassword() *string {
	return s.Password
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDisk() *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	return s.SystemDisk
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDiskDescription() *string {
	return s.SystemDiskDescription
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDiskName() *string {
	return s.SystemDiskName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetTag() []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag {
	return s.Tag
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetUserData() *string {
	return s.UserData
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetCpuOptions() *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions {
	return s.CpuOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetImageOptions() *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions {
	return s.ImageOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSchedulerOptions() *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions {
	return s.SchedulerOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSecurityOptions() *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetArn(v []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.Arn = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetAutoReleaseTime(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetCreditSpecification(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.CreditSpecification = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetDataDisk(v []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.DataDisk = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetDeploymentSetId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetHostName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.HostName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetHostNames(v []*string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.HostNames = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetImageFamily(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.ImageFamily = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetImageId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.ImageId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetInstanceDescription(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.InstanceDescription = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetInstanceName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.InstanceName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetInternetChargeType(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.InternetChargeType = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetInternetMaxBandwidthIn(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetInternetMaxBandwidthOut(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetIoOptimized(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.IoOptimized = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetKeyPairName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.KeyPairName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetPassword(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.Password = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetPasswordInherit(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.PasswordInherit = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetRamRoleName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.RamRoleName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetResourceGroupId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSecurityEnhancementStrategy(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSecurityGroupId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSecurityGroupIds(v []*string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDisk(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDisk = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDiskCategory(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDiskDescription(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDiskDescription = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDiskName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDiskName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDiskPerformanceLevel(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSystemDiskSize(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetTag(v []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.Tag = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetUserData(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.UserData = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetAutoRenew(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.AutoRenew = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetAutoRenewPeriod(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetCpuOptions(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.CpuOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetImageOptions(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.ImageOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetPeriod(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.Period = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetPeriodUnit(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSchedulerOptions(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SchedulerOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSecurityOptions(v *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SecurityOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSpotDuration(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SpotDuration = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) SetSpotInterruptionBehavior(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfiguration) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
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
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn struct {
	// > This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// 123456789012****
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// 34458433936495****:alice
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) SetAssumeRoleFor(v int64) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) SetRoleType(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn {
	s.RoleType = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) SetRolearn(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn {
	s.Rolearn = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk struct {
	// The ID of the automatic snapshot policy applied to the data disk.
	//
	// Note:
	//
	// - This parameter takes effect only when you create a one-time synchronous delivery auto provisioning group (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of data disk N. Valid values of N: 1 to 16. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD (ESSD).
	//
	// - cloud: basic disk.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether the data disk is released when the instance is released. Valid values:
	//
	// - true: the data disk is released when the instance is released.
	//
	// - false: the data disk is not released when the instance is released.
	//
	// Default value: true.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// DataDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// /dev/vd1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// Default value: empty.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether data disk N is encrypted. Valid values:
	//
	// - true: encrypted.
	//
	// - false: not encrypted.
	//
	// Default value: false.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key for the data disk. If both a launch template and launch configuration are specified, the launch template takes precedence.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The performance level of the enterprise SSD used as a data disk. The value of N must be consistent with the N in `LaunchConfiguration.DataDisk.N.Category`. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: depends on the value of `LaunchConfiguration.DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 40 to 32768.
	//
	//     - PL1: 20 to 32768.
	//
	//     - PL2: 461 to 32768.
	//
	//     - PL3: 1261 to 32768.
	//
	// - cloud: 5 to 2000.
	//
	// > The value of this parameter must be greater than or equal to the size of the snapshot specified by `LaunchConfiguration.DataDisk.N.SnapshotId`.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create data disk N. Valid values of N: 1 to 16.
	//
	// After you specify this parameter, the `LaunchConfiguration.DataDisk.N.Size` parameter is ignored. The actual size of the created disk is the size of the specified snapshot. Snapshots created on or before July 15, 2013 cannot be used. Requests that use such snapshots are rejected.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetAutoSnapshotPolicyId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetBurstingEnabled(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetCategory(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.Category = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetDeleteWithInstance(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetDescription(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.Description = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetDevice(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.Device = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetDiskName(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetEncryptAlgorithm(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetEncrypted(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetKmsKeyId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.KmsKeyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetPerformanceLevel(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetProvisionedIops(v int64) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetSize(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.Size = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) SetSnapshotId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk struct {
	// The ID of the automatic snapshot policy to apply to the system disk.
	//
	// After you set this parameter, note the following items:
	//
	// - This parameter takes effect only when you create a one-time synchronization delivery auto provisioning group (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: does not enable the feature.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The encryption algorithm for the system disk. Valid values:
	//
	// - aes-256
	//
	// - sm4-128
	//
	// Default value: aes-256.
	//
	// If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// aes-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether the system disk is encrypted. Valid values:
	//
	// - true: encrypted.
	//
	// - false: not encrypted.
	//
	// Default value: false.
	//
	// If you specify both.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID of the system disk.
	//
	// When both a launch template and launch configuration information are specified, the launch template takes precedence.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// > This parameter is supported only when SystemDisk.Category is set to cloud_auto. For more information, see [ESSD AutoPL disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetBurstingEnabled(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetEncryptAlgorithm(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetEncrypted(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetKMSKeyId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) SetProvisionedIops(v int64) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag struct {
	// The tag key of the instance. Valid values of N: 1 to 20. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain `http://` or `https://`. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters in length and cannot start with acs:. It cannot contain `http://` or `https://`. If you specify both a launch template and launch configuration information, the launch template takes precedence.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) SetKey(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag {
	s.Key = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) SetValue(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag {
	s.Value = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions struct {
	// The number of CPU cores.
	//
	// Default value: see [Specify and view CPU options](https://www.alibabacloud.com/help/en/ecs/user-guide/specify-and-view-cpu-options).
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of threads per CPU core. The number of vCPUs of the ECS instance = CpuOptions.Core value × CpuOptions.ThreadsPerCore value.
	//
	// If CpuOptions.ThreadsPerCore is set to 1, CPU hyper-threading is disabled.
	//
	// Only specific instance types support custom CPU thread counts.
	//
	// For valid values and default values, see [Specify and view CPU options](https://www.alibabacloud.com/help/en/ecs/user-guide/specify-and-view-cpu-options).
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) SetCore(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions {
	s.Core = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) SetThreadsPerCore(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions struct {
	// Specifies whether the instance that uses this image supports logon as the ecs-user user. Valid values:
	//
	// - true: supported.
	//
	// - false: not supported.
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) SetLoginAsNonRoot(v bool) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions struct {
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	DedicatedHostId        *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) SetDedicatedHostClusterId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) SetDedicatedHostId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions struct {
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) SetTrustedSystemMode(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestCandidateOptions struct {
	Evaluate *bool `json:"Evaluate,omitempty" xml:"Evaluate,omitempty"`
	// example:
	//
	// 60
	TimeoutMinutes *int32 `json:"TimeoutMinutes,omitempty" xml:"TimeoutMinutes,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestCandidateOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) GetEvaluate() *bool {
	return s.Evaluate
}

func (s *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) GetTimeoutMinutes() *int32 {
	return s.TimeoutMinutes
}

func (s *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) SetEvaluate(v bool) *CreateAutoProvisioningGroupShrinkRequestCandidateOptions {
	s.Evaluate = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) SetTimeoutMinutes(v int32) *CreateAutoProvisioningGroupShrinkRequestCandidateOptions {
	s.TimeoutMinutes = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestDataDiskConfig struct {
	// The category of the data disk. You can specify multiple candidate disk categories. The specified order determines the priority of each disk category. When a disk category is unavailable, the system automatically switches to the next category. Valid values:
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD (ESSD).
	//
	// -   cloud: basic disk.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) SetDiskCategory(v string) *CreateAutoProvisioningGroupShrinkRequestDataDiskConfig {
	s.DiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestDataDiskConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig struct {
	// The list of architecture types for instance types.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// - Exclude: excludes burstable instance types.
	//
	// - Include: includes burstable instance types.
	//
	// - Required: includes only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The list of vCPU core counts for instance types.
	Cores []*int32 `json:"Cores,omitempty" xml:"Cores,omitempty" type:"Repeated"`
	// The list of instance types to exclude.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The image ID. You can use this parameter to set the image for the current resource pool. If not set, the image specified in `LaunchConfiguration.ImageId` or the launch template is used by default. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources.
	//
	// Note: This parameter is supported only when `AutoProvisioningGroupType = instant`.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20210425.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The level of the instance family, used to filter instance types that meet the requirements. Valid values:
	//
	// - EntryLevel: entry level, which refers to shared instance types. Lower cost but no guarantee of stable computing performance. Suitable for scenarios with low average CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise level. Stable performance with dedicated resources. Suitable for scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit-based entry level, which refers to burstable instances. Uses CPU credits to ensure computing performance. Suitable for scenarios with low average CPU utilization and occasional bursts. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance type in the extended launch template. Valid values of N: 1 to 20. For valid values, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum price for spot instances in the extended launch template.
	//
	// > After you set `LaunchTemplateConfig`, `LaunchTemplateConfig.N.MaxPrice` is required.
	//
	// example:
	//
	// 3
	MaxPrice *float64 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// > This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// false
	MaxQuantity *int32 `json:"MaxQuantity,omitempty" xml:"MaxQuantity,omitempty"`
	// The list of memory sizes for instance types.
	Memories []*float32 `json:"Memories,omitempty" xml:"Memories,omitempty" type:"Repeated"`
	// The priority of the extended launch template. A value of 0 indicates the highest priority. Valid values: 0 to +∞.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the vSwitch to which the ECS instance in the extended launch template is connected. The zone of the ECS instance created from the extended template is determined by the vSwitch.
	//
	// > If you specify LaunchTemplateConfig, LaunchTemplateConfig.N.VSwitchId is required.
	//
	// example:
	//
	// vsw-sn5bsitu4lfzgc5o7****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The weight of the instance type in the extended launch template. A higher value indicates that a single instance can meet more computing power requirements, which reduces the number of instances required. Valid values: greater than 0.
	//
	// You can calculate the weight based on the computing power of the specified instance type and the minimum computing power of a single node in the cluster. For example, if the minimum computing power of a single node is 8 vCPUs and 60 GiB:
	//
	// - The weight of an instance type with 8 vCPUs and 60 GiB can be set to 1.
	//
	// - The weight of an instance type with 16 vCPUs and 120 GiB can be set to 2.
	//
	// example:
	//
	// 2
	WeightedCapacity *float64 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetArchitectures() []*string {
	return s.Architectures
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetCores() []*int32 {
	return s.Cores
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetMaxPrice() *float64 {
	return s.MaxPrice
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetMaxQuantity() *int32 {
	return s.MaxQuantity
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetMemories() []*float32 {
	return s.Memories
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) GetWeightedCapacity() *float64 {
	return s.WeightedCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetArchitectures(v []*string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.Architectures = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetBurstablePerformance(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetCores(v []*int32) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.Cores = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetExcludedInstanceTypes(v []*string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetImageId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.ImageId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetInstanceFamilyLevel(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetInstanceType(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.InstanceType = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetMaxPrice(v float64) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.MaxPrice = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetMaxQuantity(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.MaxQuantity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetMemories(v []*float32) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.Memories = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetPriority(v int32) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.Priority = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetVSwitchId(v string) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.VSwitchId = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) SetWeightedCapacity(v float64) *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig {
	s.WeightedCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestPrePaidOptions struct {
	// The minimum capacity set for different instance types. This parameter is supported only when `AutoProvisioningGroupType = request`.
	SpecifyCapacityDistribution []*CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution `json:"SpecifyCapacityDistribution,omitempty" xml:"SpecifyCapacityDistribution,omitempty" type:"Repeated"`
}

func (s CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) GetSpecifyCapacityDistribution() []*CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution {
	return s.SpecifyCapacityDistribution
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) SetSpecifyCapacityDistribution(v []*CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions {
	s.SpecifyCapacityDistribution = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions) Validate() error {
	if s.SpecifyCapacityDistribution != nil {
		for _, item := range s.SpecifyCapacityDistribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution struct {
	// The set of instance types. Duplicate values are not allowed, and the instance types must be within the range of LaunchTemplateConfig.InstanceType.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The minimum number of instances to deliver within the `InstanceTypes` range.
	//
	// > The sum of all MinTargetCapacity values (`sum(MinTargetCapacity) <= TotalTargetCapacity`) cannot exceed TotalTargetCapacity. If any instance type set cannot meet the MinTargetCapacity requirement due to insufficient inventory or other reasons, the entire request fails.
	//
	// example:
	//
	// 5
	MinTargetCapacity *int32 `json:"MinTargetCapacity,omitempty" xml:"MinTargetCapacity,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) GetMinTargetCapacity() *int32 {
	return s.MinTargetCapacity
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) SetInstanceTypes(v []*string) *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution {
	s.InstanceTypes = v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) SetMinTargetCapacity(v int32) *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution {
	s.MinTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestPrePaidOptionsSpecifyCapacityDistribution) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig struct {
	// The category of the system disk. You can specify multiple candidate disk categories. The specified order determines the priority of each disk category. When a disk category is unavailable, the system automatically switches to the next category. Valid values:
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD (ESSD).
	//
	// -   cloud: basic disk.
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) SetDiskCategory(v string) *CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig {
	s.DiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupShrinkRequestTag struct {
	// The tag key of the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoProvisioningGroupShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoProvisioningGroupShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoProvisioningGroupShrinkRequestTag) SetKey(v string) *CreateAutoProvisioningGroupShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestTag) SetValue(v string) *CreateAutoProvisioningGroupShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAutoProvisioningGroupShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}

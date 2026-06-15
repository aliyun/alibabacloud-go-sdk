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
	// The name of the auto provisioning group. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// apg-test
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The delivery type of the auto provisioning group. Valid values:
	//
	// - request: One-time asynchronous delivery. The group delivers the instance cluster only at startup. If scheduling fails, no retry occurs.
	//
	// - instant: One-time synchronous delivery. The group creates instances synchronously at startup and returns the list of successfully created instances and reasons for failures in the response.
	//
	// - maintain: Continuous provisioning. The group attempts to deliver the instance cluster at startup and monitors real-time capacity. If the target capacity is not met, it continues creating ECS instances.
	//
	// Default value: maintain.
	//
	// example:
	//
	// maintain
	AutoProvisioningGroupType *string                                                   `json:"AutoProvisioningGroupType,omitempty" xml:"AutoProvisioningGroupType,omitempty"`
	CandidateOptions          *CreateAutoProvisioningGroupShrinkRequestCandidateOptions `json:"CandidateOptions,omitempty" xml:"CandidateOptions,omitempty" type:"Struct"`
	// Ensures request idempotence. Generate a unique value from your client for this parameter to ensure uniqueness across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of data disk configurations for instances.
	DataDiskConfig []*CreateAutoProvisioningGroupShrinkRequestDataDiskConfig `json:"DataDiskConfig,omitempty" xml:"DataDiskConfig,omitempty" type:"Repeated"`
	// Specifies the billing method for the capacity difference when the sum of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is less than `TotalTargetCapacity`. Valid values:
	//
	// - PayAsYouGo: Pay-as-you-go instances.
	//
	// - Spot: Spot instances.
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
	// Specifies whether to release instances when the real-time capacity of the auto provisioning group exceeds the target capacity and scale-in is triggered. Valid values:
	//
	// - termination: Releases scaled-in instances.
	//
	// - no-termination: Only removes scaled-in instances from the auto provisioning group.
	//
	// Default value: no-termination.
	//
	// example:
	//
	// termination
	ExcessCapacityTerminationPolicy *string `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	ExecutionMode                   *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// > This parameter is in invitational preview and is not yet available.
	//
	// example:
	//
	// false
	HibernationOptionsConfigured *bool `json:"HibernationOptionsConfigured,omitempty" xml:"HibernationOptionsConfigured,omitempty"`
	// The list of extended launch template configurations.
	LaunchTemplateConfig []*CreateAutoProvisioningGroupShrinkRequestLaunchTemplateConfig `json:"LaunchTemplateConfig,omitempty" xml:"LaunchTemplateConfig,omitempty" type:"Repeated"`
	// The ID of the launch template associated with the auto provisioning group. Call [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) to query available launch templates. When both a launch template and launch configuration parameters (`LaunchConfiguration.*`) are specified, the launch template takes precedence.
	//
	// example:
	//
	// lt-bp1fgzds4bdogu03****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the launch template associated with the auto provisioning group. Call [DescribeLaunchTemplateVersions](https://help.aliyun.com/document_detail/73761.html) to query available launch template versions.
	//
	// Default value: The default version of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum price for spot instances in the auto provisioning group.
	//
	// > If both `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` are set, the lower value takes effect.
	//
	// example:
	//
	// 2
	MaxSpotPrice *float32 `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	// The minimum target capacity of the auto provisioning group. Valid values: Positive integers.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with `AutoProvisioningGroupType=instant`.
	//
	// - If the instance inventory in the region is less than this value, the API call fails and no instances are created.
	//
	// - If the instance inventory in the region is greater than this value, instances are created based on other configured parameters.
	//
	// example:
	//
	// 20
	MinTargetCapacity *string `json:"MinTargetCapacity,omitempty" xml:"MinTargetCapacity,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The strategy for creating pay-as-you-go instances. Valid values:
	//
	// - lowest-price: Cost optimization strategy. Selects the instance type with the lowest price.
	//
	// - prioritized: Priority-based strategy. Creates instances based on the priority specified in `LaunchTemplateConfig.N.Priority`.
	//
	// Default value: lowest-price.
	//
	// example:
	//
	// prioritized
	PayAsYouGoAllocationStrategy *string `json:"PayAsYouGoAllocationStrategy,omitempty" xml:"PayAsYouGoAllocationStrategy,omitempty"`
	// The target capacity for pay-as-you-go instances in the auto provisioning group. Valid values: Integers less than or equal to the value of `TotalTargetCapacity`.
	//
	// example:
	//
	// 30
	PayAsYouGoTargetCapacity *string `json:"PayAsYouGoTargetCapacity,omitempty" xml:"PayAsYouGoTargetCapacity,omitempty"`
	// Detailed capacity configuration for subscription instances.
	PrePaidOptions *CreateAutoProvisioningGroupShrinkRequestPrePaidOptions `json:"PrePaidOptions,omitempty" xml:"PrePaidOptions,omitempty" type:"Struct"`
	// The region ID of the auto provisioning group. Call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest Alibaba Cloud region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the auto provisioning group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource pool strategy used when creating instances. Note:
	//
	// - This parameter takes effect only when creating pay-as-you-go instances.
	//
	// - This parameter takes effect only when creating an auto provisioning group with `AutoProvisioningGroupType=instant`.
	ResourcePoolOptionsShrink *string `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty"`
	// The strategy for creating spot instances. Valid values:
	//
	// - lowest-price: Cost optimization strategy. Selects the instance type with the lowest price.
	//
	// - diversified: Balanced zone distribution strategy. Creates instances across the zones specified in the launch template configurations and distributes them evenly.
	//
	// - capacity-optimized: Capacity optimization strategy. Selects the optimal instance type and zone based on inventory availability.
	//
	// Default value: lowest-price.
	//
	// example:
	//
	// diversified
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The behavior when a spot instance is interrupted. Valid values:
	//
	// - stop: Stops the instance.
	//
	// - terminate: Releases the instance.
	//
	// Default value: terminate.
	//
	// example:
	//
	// terminate
	SpotInstanceInterruptionBehavior *string `json:"SpotInstanceInterruptionBehavior,omitempty" xml:"SpotInstanceInterruptionBehavior,omitempty"`
	// Takes effect only when `SpotAllocationStrategy` is set to `lowest-price`. Specifies the number of lowest-priced instance types from which the auto provisioning group creates instances.
	//
	// Valid values: Less than the value of N in `LaunchTemplateConfig.N`.
	//
	// example:
	//
	// 2
	SpotInstancePoolsToUseCount *int32 `json:"SpotInstancePoolsToUseCount,omitempty" xml:"SpotInstancePoolsToUseCount,omitempty"`
	// The target capacity for spot instances in the auto provisioning group. Valid values: Integers less than or equal to the value of `TotalTargetCapacity`.
	//
	// example:
	//
	// 20
	SpotTargetCapacity *string `json:"SpotTargetCapacity,omitempty" xml:"SpotTargetCapacity,omitempty"`
	// The list of system disk configurations for instances.
	SystemDiskConfig []*CreateAutoProvisioningGroupShrinkRequestSystemDiskConfig `json:"SystemDiskConfig,omitempty" xml:"SystemDiskConfig,omitempty" type:"Repeated"`
	// The list of tags bound to the auto provisioning group.
	Tag []*CreateAutoProvisioningGroupShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to release instances in the group when you delete the auto provisioning group. Valid values:
	//
	// - true: Releases instances in the group.
	//
	// - false: Retains instances in the group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstances *bool `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
	// Specifies whether to release instances in the group when the auto provisioning group expires. Valid values:
	//
	// - true: Releases instances in the group.
	//
	// - false: Only removes instances from the auto provisioning group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstancesWithExpiration *bool `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	// The total target capacity of the auto provisioning group. Valid values: Positive integers.
	//
	// The total capacity must be greater than or equal to the sum of `PayAsYouGoTargetCapacity` (target capacity for pay-as-you-go instances) and `SpotTargetCapacity` (target capacity for spot instances).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	TotalTargetCapacity *string `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
	// The start time of the auto provisioning group. Used together with `ValidUntil` to define the validity period.
	//
	// Specify the time in [ISO 8601](https://help.aliyun.com/document_detail/25696.html) format using UTC+0 time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// Default value: The timestamp when the API call takes effect immediately.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	ValidFrom *string `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The expiration time of the auto provisioning group. Used together with `ValidFrom` to define the validity period.
	//
	// Specify the time in [ISO 8601](https://help.aliyun.com/document_detail/25696.html) format using UTC+0 time. Format: yyyy-MM-ddTHH:mm:ssZ.
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
	// > This parameter is in invitational preview and is not supported.
	Arn []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// The automatic release time for pay-as-you-go instances. Specify the time in [ISO 8601](https://help.aliyun.com/document_detail/25696.html) format using UTC+0 time. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// - If seconds (`ss`) are not `00`, the time is rounded down to the start of the current minute (`mm`).
	//
	// - The earliest release time is 30 minutes after the current time.
	//
	// - The latest release time cannot exceed three years from the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The running mode of burstable instances. Valid values:
	//
	// - Standard: Standard mode. For more information, see the "Performance-constrained mode" section in [What are burstable instances?](https://help.aliyun.com/document_detail/59977.html)
	//
	// - Unlimited: Unlimited mode. For more information, see the "Unlimited mode" section in [What are burstable instances?](https://help.aliyun.com/document_detail/59977.html)
	//
	// Default value: None.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disk configurations for the extended launch template.
	DataDisk []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the instance. Requirements:
	//
	// - Periods (.) and hyphens (-) cannot be the first or last character and cannot appear consecutively.
	//
	// - Windows instances: 2–15 characters. Periods (.) are not supported. Cannot consist of only digits. Can contain letters, digits, and hyphens (-).
	//
	// - Other instances (such as Linux): 2–64 characters. Multiple periods (.) are supported. Each segment between periods can contain letters, digits, and hyphens (-).
	//
	// - Do not set both `LaunchConfiguration.HostName` and `LaunchConfiguration.HostNames.N`. Otherwise, an error is returned.
	//
	// - When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// k8s-node-[1,4]-ecshost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The list of hostnames for one or more instances. Requirements:
	//
	// - This parameter takes effect only when creating an auto provisioning group with `AutoProvisioningGroupType=instant`.
	//
	// - N indicates the number of instances. Valid values: 1 to 1000. The value must match TotalTargetCapacity.
	//
	// - Periods (.) and hyphens (-) cannot be the first or last character and cannot appear consecutively.
	//
	// - When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// ecs-host-01
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The image family name. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image ID. This is the image used when launching instances. Call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available images. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// m-bp1g7004ksh0oeuc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// Instance_Description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// Default value: The instance `InstanceId`.
	//
	// To create multiple ECS instances, you can batch configure sequential instance names. For more information, see [Batch configure sequential instance names or hostnames](https://help.aliyun.com/document_detail/196048.html).
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network billing type. Valid values:
	//
	// - PayByBandwidth: Pay-by-bandwidth.
	//
	// - PayByTraffic: Pay-by-traffic.
	//
	// > For pay-by-traffic, inbound and outbound bandwidth peaks represent upper limits and are not service-level commitments. Bandwidth may be throttled during resource contention. Use pay-by-bandwidth if your workload requires guaranteed bandwidth.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth, in Mbit/s. Valid values:
	//
	// - When outbound public bandwidth is ≤ 10 Mbit/s: 1–10. Default: 10.
	//
	// - When outbound public bandwidth is > 10 Mbit/s: 1–`LaunchConfiguration.InternetMaxBandwidthOut`. Default: `LaunchConfiguration.InternetMaxBandwidthOut`.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth, in Mbit/s. Valid values: 0–100.
	//
	// Default value: 0.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// - none: Not I/O optimized.
	//
	// - optimized: I/O optimized.
	//
	// For retired instance types, the default value is none. For other instance types, the default value is optimized.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The key pair name.
	//
	// - For Windows instances, this parameter is ignored. Default value: empty.
	//
	// - For Linux instances, password logon is disabled after initialization.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The instance password. The password must be 8 to 30 characters in length and include at least three of the following: uppercase letters, lowercase letters, digits, and special characters. Valid special characters:
	//
	// ``()`~!@#$%^&*-_+=|{}`[]`:;\\"<>,.?/``
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. Valid values:
	//
	// - true: Uses the preset password.
	//
	// - false: Does not use the preset password.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// true
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The RAM role name of the instance. Use the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query your created RAM roles. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// RAM_Name
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The resource group ID of the instance. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// - Active: Enables security hardening. Applies only to public images.
	//
	// - Deactive: Disables security hardening. Applies to all image types.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The security group ID of the instance. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The list of security groups to which the instance belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// System disk information for the instance. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	SystemDisk *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The category of the system disk. Valid values:
	//
	// - cloud_efficiency: Ultra disk.
	//
	// - cloud_ssd: Standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: Basic disk.
	//
	// For retired instance types that are not I/O optimized, the default value is cloud. Otherwise, the default value is cloud_efficiency.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// SystemDisk_Description
	SystemDiskDescription *string `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssdSystem
	SystemDiskName *string `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	// The performance level of the ESSD used as the system disk. Valid values:
	//
	// - PL0 (default): Up to 10,000 random read/write IOPS per disk.
	//
	// - PL1: Up to 50,000 random read/write IOPS per disk.
	//
	// - PL2: Up to 100,000 random read/write IOPS per disk.
	//
	// - PL3: Up to 1,000,000 random read/write IOPS per disk.
	//
	// For more information about selecting ESSD performance levels, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// PL0
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The size of the system disk, in GiB. Valid values: 20–500. The value must be greater than or equal to max{20, size of the image specified by LaunchConfiguration.ImageId}.
	//
	// Default value: max{40, size of the image specified by LaunchConfiguration.ImageId}.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The list of tags for the extended launch template.
	Tag []*CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The instance user data. Encode the data in Base64. The raw data cannot exceed 32 KB. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether to enable auto-renewal. Takes effect when creating subscription instances. Valid values:
	//
	// - true: Enables auto-renewal.
	//
	// - false (default): Disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration per cycle. Valid values:
	//
	// <props="china">
	//
	// - When PeriodUnit=Week: 1, 2, 3.
	//
	// - When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, 60.
	//
	//
	//
	// <props="intl">
	//
	// When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, 60.
	//
	//
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// CPU configuration.
	CpuOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// Image-related properties.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with AutoProvisioningGroupType=instant.
	ImageOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Required when creating subscription instances. Valid values:
	//
	// <props="china">
	//
	// - When PeriodUnit=Week, Period values: 1, 2, 3, 4.
	//
	// - When PeriodUnit=Month, Period values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60.
	//
	//
	//
	// <props="intl">
	//
	// When PeriodUnit=Month, Period values: 1, 2, 3, 6, 12.
	//
	//
	//
	// <props="partner">
	//
	// When PeriodUnit=Month, Period values: 1, 2, 3, 6, 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for subscription billing. Valid values:
	//
	// <props="china">
	//
	// - Week.
	//
	// - Month (default).
	//
	//
	//
	// <props="intl">
	//
	// Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit       *string                                                                      `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SchedulerOptions *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SecurityOptions  *CreateAutoProvisioningGroupShrinkRequestLaunchConfigurationSecurityOptions  `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The reserved duration for spot instances, in hours. Default value: 1. Valid values:
	//
	// - 1: Alibaba Cloud guarantees that the instance runs for 1 hour without being automatically released. After 1 hour, the system compares your bid price with the market price and checks inventory to decide whether to retain or reclaim the instance.
	//
	// - 0: Alibaba Cloud does not guarantee that the instance runs for 1 hour. The system immediately compares your bid price with the market price and checks inventory to decide whether to retain or reclaim the instance.
	//
	// Alibaba Cloud sends an ECS system event notification 5 minutes before reclaiming a spot instance. Spot instances are billed per second. Choose the reserved duration based on your task execution time.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with AutoProvisioningGroupType=instant.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption behavior for spot instances. Valid values:
	//
	// - Terminate: Releases the instance immediately.
	//
	// - Stop: Puts the instance into economical mode.
	//
	// For more information about economical mode, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with AutoProvisioningGroupType=instant.
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
	// > This parameter is in invitational preview and is not supported.
	//
	// example:
	//
	// 123456789012****
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is in invitational preview and is not supported.
	//
	// example:
	//
	// 34458433936495****:alice
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is in invitational preview and is not supported.
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
	// The automatic snapshot policy ID applied to the data disk.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with AutoProvisioningGroupType=instant.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable performance burst. Valid values:
	//
	// - true: Enables performance burst.
	//
	// - false: Disables performance burst.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of data disk N. N ranges from 1 to 16. Valid values:
	//
	// - cloud_efficiency: Ultra disk.
	//
	// - cloud_ssd: Standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: Basic disk.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: Releases the data disk with the instance.
	//
	// - false: Does not release the data disk with the instance.
	//
	// Default value: true.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// DataDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// /dev/vd1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not yet available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
	//
	// - true: Encrypts the disk.
	//
	// - false: Does not encrypt the disk.
	//
	// Default value: false.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID for the data disk. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The performance level of the ESSD used as a data disk. The value of N must match the N in `LaunchConfiguration.DataDisk.N.Category`. Valid values:
	//
	// - PL0: Up to 10,000 random read/write IOPS per disk.
	//
	// - PL1 (default): Up to 50,000 random read/write IOPS per disk.
	//
	// - PL2: Up to 100,000 random read/write IOPS per disk.
	//
	// - PL3: Up to 1,000,000 random read/write IOPS per disk.
	//
	// For more information about selecting ESSD performance levels, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS for ESSD AutoPL disks. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N, in GiB. N ranges from 1 to 16. Valid values:
	//
	// - cloud_efficiency: 20–32768.
	//
	// - cloud_ssd: 20–32768.
	//
	// - cloud_essd: The valid range depends on the value of `LaunchConfiguration.DataDisk.N.PerformanceLevel`.
	//
	//   - PL0: 40–32768.
	//
	//   - PL1: 20–32768.
	//
	//   - PL2: 461–32768.
	//
	//   - PL3: 1261–32768
	//
	// - cloud: 5–2000.
	//
	// > The value must be greater than or equal to the size of the snapshot specified by `LaunchConfiguration.DataDisk.N.SnapshotId`.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot used to create data disk N. N ranges from 1 to 16.
	//
	// After this parameter is specified, `LaunchConfiguration.DataDisk.N.Size` is ignored. The actual disk size equals the size of the specified snapshot. Snapshots created on or before July 15, 2013 are not supported and will cause the request to fail.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
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
	// The automatic snapshot policy ID applied to the system disk.
	//
	// Note:
	//
	// - This parameter takes effect only when creating an auto provisioning group with AutoProvisioningGroupType=instant.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable performance burst. Valid values:
	//
	// - true: Enables performance burst.
	//
	// - false: Disables performance burst.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The encryption algorithm for the system disk. Valid values:
	//
	// - aes-256.
	//
	// - sm4-128.
	//
	// Default value: aes-256.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// aes-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt system disk N. Valid values:
	//
	// - true: Encrypts the disk.
	//
	// - false: Does not encrypt the disk.
	//
	// Default value: false.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID for the system disk.
	//
	// When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The provisioned read/write IOPS for ESSD AutoPL disks. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// > This parameter is supported only when SystemDisk.Category is set to cloud_auto. For more information, see [ESSD AutoPL](https://help.aliyun.com/document_detail/368372.html).
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
	// The tag key of the instance. N ranges from 1 to 20. If specified, the value cannot be an empty string. The key can be up to 128 characters in length and cannot start with aliyun or acs:. It also cannot contain `http://` or `https://`. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. N ranges from 1 to 20. If specified, the value can be an empty string. The value can be up to 128 characters in length and cannot start with acs:. It also cannot contain `http://` or `https://`. When both a launch template and launch configuration parameters are specified, the launch template takes precedence.
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
	// Default value: See [Customize CPU options](https://help.aliyun.com/zh/ecs/user-guide/specify-and-view-cpu-options?spm=a2c4g.11186623.0.0.734f769asTEobd).
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of CPU threads. The vCPU count of an ECS instance equals CpuOptions.Core × CpuOptions.ThreadsPerCore.
	//
	// CpuOptions.ThreadsPerCore=1 disables CPU hyper-threading.
	//
	// Only specific instance types support setting the number of CPU threads.
	//
	// Valid values and default values: See [Customize CPU options](https://help.aliyun.com/zh/ecs/user-guide/specify-and-view-cpu-options?spm=a2c4g.11186623.0.0.734f769aeIFsoj).
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
	// Specifies whether instances using this image support logging on as the ecs-user. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
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
	// The data disk type. You can specify multiple candidate disk types. The order specifies their priority. If one disk type is unavailable, the system automatically switches to the next type. Valid values:
	//
	// - cloud_efficiency: Ultra disk.
	//
	// - cloud_ssd: Standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: Basic disk.
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
	// Specifies whether the instance type supports performance bursts. Valid values:
	//
	// - Exclude: Excludes burstable instance types.
	//
	// - Include: Includes burstable instance types.
	//
	// - Required: Includes only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The list of vCPU counts for instance types.
	Cores []*int32 `json:"Cores,omitempty" xml:"Cores,omitempty" type:"Repeated"`
	// The list of instance types to exclude.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The image ID. Use this parameter to specify the image for the current resource pool. If not set, the image specified in `LaunchConfiguration.ImageId` or the launch template is used by default. Call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available images.
	//
	// Note: This parameter is supported only when `AutoProvisioningGroupType = instant`.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20210425.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance family level, used to filter eligible instance types. Valid values:
	//
	// - EntryLevel: Entry-level, or shared-resource instances. Lower cost but no guaranteed stable computing performance. Suitable for workloads with low average CPU usage. For more information, see [Shared-resource instances](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: Enterprise-level. Stable performance with dedicated resources. Suitable for workloads requiring high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: Credit entry-level, or burstable instances. Uses CPU credits to guarantee computing performance. Suitable for workloads with low average CPU usage and occasional bursts. For more information, see [Burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// N ranges from 1 to 10.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance type corresponding to the extended launch template. N ranges from 1 to 20. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price for spot instances in the extended launch template.
	//
	// > After `LaunchTemplateConfig` is set, `LaunchTemplateConfig.N.MaxPrice` is required.
	//
	// example:
	//
	// 3
	MaxPrice *float64 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// > This parameter is in invitational preview and is not supported.
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
	// The ID of the virtual switch to which the ECS instance belongs in the extended launch template. The zone of the ECS instance launched from this template is determined by the virtual switch.
	//
	// > After `LaunchTemplateConfig` is set, `LaunchTemplateConfig.N.VSwitchId` is required.
	//
	// example:
	//
	// vsw-sn5bsitu4lfzgc5o7****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The weight of the instance type in the extended launch template. A higher value indicates greater computing power per instance and fewer instances needed. Valid values: Greater than 0.
	//
	// You can calculate the weight based on the computing power of the specified instance type and the minimum computing power required per node in the cluster. For example, if the minimum computing power per node is 8 vCPUs and 60 GiB memory:
	//
	// - An instance type with 8 vCPUs and 60 GiB memory can have a weight of 1.
	//
	// - An instance type with 16 vCPUs and 120 GiB memory can have a weight of 2.
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
	// The set of instance types. Duplicates are not allowed, and the types must be within the range of LaunchTemplateConfig.InstanceType.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The minimum number of instances to deliver within the `InstanceTypes` range.
	//
	// > `sum(MinTargetCapacity)<= TotalTargetCapacity`. The sum of MinTargetCapacity across all instance type sets cannot exceed TotalTargetCapacity. If any instance type set fails to meet its MinTargetCapacity due to inventory issues, the entire request fails.
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
	// The system disk type. You can specify multiple candidate disk types. The order specifies their priority. If one disk type is unavailable, the system automatically switches to the next type. Valid values:
	//
	// - cloud_efficiency: Ultra disk.
	//
	// - cloud_ssd: Standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: Basic disk.
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
	// N ranges from 1 to 20. If specified, the value cannot be an empty string. The key can be up to 128 characters in length and cannot start with aliyun or acs:. It also cannot contain http\\:// or https\\://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the auto provisioning group.
	//
	// N ranges from 1 to 20. If specified, the value can be an empty string. The value can be up to 128 characters in length and cannot contain http\\:// or https\\://.
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

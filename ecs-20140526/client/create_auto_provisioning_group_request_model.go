// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoProvisioningGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchConfiguration(v *CreateAutoProvisioningGroupRequestLaunchConfiguration) *CreateAutoProvisioningGroupRequest
	GetLaunchConfiguration() *CreateAutoProvisioningGroupRequestLaunchConfiguration
	SetAutoProvisioningGroupName(v string) *CreateAutoProvisioningGroupRequest
	GetAutoProvisioningGroupName() *string
	SetAutoProvisioningGroupType(v string) *CreateAutoProvisioningGroupRequest
	GetAutoProvisioningGroupType() *string
	SetCandidateOptions(v *CreateAutoProvisioningGroupRequestCandidateOptions) *CreateAutoProvisioningGroupRequest
	GetCandidateOptions() *CreateAutoProvisioningGroupRequestCandidateOptions
	SetClientToken(v string) *CreateAutoProvisioningGroupRequest
	GetClientToken() *string
	SetDataDiskConfig(v []*CreateAutoProvisioningGroupRequestDataDiskConfig) *CreateAutoProvisioningGroupRequest
	GetDataDiskConfig() []*CreateAutoProvisioningGroupRequestDataDiskConfig
	SetDefaultTargetCapacityType(v string) *CreateAutoProvisioningGroupRequest
	GetDefaultTargetCapacityType() *string
	SetDescription(v string) *CreateAutoProvisioningGroupRequest
	GetDescription() *string
	SetExcessCapacityTerminationPolicy(v string) *CreateAutoProvisioningGroupRequest
	GetExcessCapacityTerminationPolicy() *string
	SetExecutionMode(v string) *CreateAutoProvisioningGroupRequest
	GetExecutionMode() *string
	SetHibernationOptionsConfigured(v bool) *CreateAutoProvisioningGroupRequest
	GetHibernationOptionsConfigured() *bool
	SetLaunchTemplateConfig(v []*CreateAutoProvisioningGroupRequestLaunchTemplateConfig) *CreateAutoProvisioningGroupRequest
	GetLaunchTemplateConfig() []*CreateAutoProvisioningGroupRequestLaunchTemplateConfig
	SetLaunchTemplateId(v string) *CreateAutoProvisioningGroupRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersion(v string) *CreateAutoProvisioningGroupRequest
	GetLaunchTemplateVersion() *string
	SetMaxSpotPrice(v float32) *CreateAutoProvisioningGroupRequest
	GetMaxSpotPrice() *float32
	SetMinTargetCapacity(v string) *CreateAutoProvisioningGroupRequest
	GetMinTargetCapacity() *string
	SetOwnerAccount(v string) *CreateAutoProvisioningGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAutoProvisioningGroupRequest
	GetOwnerId() *int64
	SetPayAsYouGoAllocationStrategy(v string) *CreateAutoProvisioningGroupRequest
	GetPayAsYouGoAllocationStrategy() *string
	SetPayAsYouGoTargetCapacity(v string) *CreateAutoProvisioningGroupRequest
	GetPayAsYouGoTargetCapacity() *string
	SetPrePaidOptions(v *CreateAutoProvisioningGroupRequestPrePaidOptions) *CreateAutoProvisioningGroupRequest
	GetPrePaidOptions() *CreateAutoProvisioningGroupRequestPrePaidOptions
	SetRegionId(v string) *CreateAutoProvisioningGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAutoProvisioningGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateAutoProvisioningGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoProvisioningGroupRequest
	GetResourceOwnerId() *int64
	SetResourcePoolOptions(v *CreateAutoProvisioningGroupRequestResourcePoolOptions) *CreateAutoProvisioningGroupRequest
	GetResourcePoolOptions() *CreateAutoProvisioningGroupRequestResourcePoolOptions
	SetSpotAllocationStrategy(v string) *CreateAutoProvisioningGroupRequest
	GetSpotAllocationStrategy() *string
	SetSpotInstanceInterruptionBehavior(v string) *CreateAutoProvisioningGroupRequest
	GetSpotInstanceInterruptionBehavior() *string
	SetSpotInstancePoolsToUseCount(v int32) *CreateAutoProvisioningGroupRequest
	GetSpotInstancePoolsToUseCount() *int32
	SetSpotTargetCapacity(v string) *CreateAutoProvisioningGroupRequest
	GetSpotTargetCapacity() *string
	SetSystemDiskConfig(v []*CreateAutoProvisioningGroupRequestSystemDiskConfig) *CreateAutoProvisioningGroupRequest
	GetSystemDiskConfig() []*CreateAutoProvisioningGroupRequestSystemDiskConfig
	SetTag(v []*CreateAutoProvisioningGroupRequestTag) *CreateAutoProvisioningGroupRequest
	GetTag() []*CreateAutoProvisioningGroupRequestTag
	SetTerminateInstances(v bool) *CreateAutoProvisioningGroupRequest
	GetTerminateInstances() *bool
	SetTerminateInstancesWithExpiration(v bool) *CreateAutoProvisioningGroupRequest
	GetTerminateInstancesWithExpiration() *bool
	SetTotalTargetCapacity(v string) *CreateAutoProvisioningGroupRequest
	GetTotalTargetCapacity() *string
	SetValidFrom(v string) *CreateAutoProvisioningGroupRequest
	GetValidFrom() *string
	SetValidUntil(v string) *CreateAutoProvisioningGroupRequest
	GetValidUntil() *string
}

type CreateAutoProvisioningGroupRequest struct {
	LaunchConfiguration *CreateAutoProvisioningGroupRequestLaunchConfiguration `json:"LaunchConfiguration,omitempty" xml:"LaunchConfiguration,omitempty" type:"Struct"`
	// The name of the auto provisioning group. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// apg-test
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The delivery type of the auto provisioning group. Valid values:
	//
	// - request: one-time asynchronous delivery. The auto provisioning group asynchronously delivers the instance cluster only at startup and does not retry if the scheduling failed.
	//
	// - instant: one-time synchronization delivery. The auto provisioning group creates instances in synchronization only at startup and returns the list of created instances and the reasons why creation failed in the response.
	//
	// - maintain: continuous delivery. The auto provisioning group attempts to deliver the instance cluster at startup, performs monitoring on the real-time capacity, and continues to create ECS instances if the target capacity is not reached.
	//
	// Default value: maintain.
	//
	// example:
	//
	// maintain
	AutoProvisioningGroupType *string `json:"AutoProvisioningGroupType,omitempty" xml:"AutoProvisioningGroupType,omitempty"`
	// The configuration options when AutoProvisioningGroupType is set to candidate.
	CandidateOptions *CreateAutoProvisioningGroupRequestCandidateOptions `json:"CandidateOptions,omitempty" xml:"CandidateOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of data disk configurations for the instance.
	DataDiskConfig []*CreateAutoProvisioningGroupRequestDataDiskConfig `json:"DataDiskConfig,omitempty" xml:"DataDiskConfig,omitempty" type:"Repeated"`
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
	// Specifies whether to release the instances that are removed during a scale-down event when the real-time capacity of the auto provisioning group exceeds the target capacity and triggers a scale-down. Valid values:
	//
	// - termination: releases the removed instances.
	//
	// - no-termination: only removes the instances from the auto provisioning group.
	//
	// Default value: no-termination.
	//
	// example:
	//
	// termination
	ExcessCapacityTerminationPolicy *string `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	ExecutionMode                   *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// > This parameter is in invitational preview and is not available.
	//
	// example:
	//
	// false
	HibernationOptionsConfigured *bool `json:"HibernationOptionsConfigured,omitempty" xml:"HibernationOptionsConfigured,omitempty"`
	// The extended launch templates.
	LaunchTemplateConfig []*CreateAutoProvisioningGroupRequestLaunchTemplateConfig `json:"LaunchTemplateConfig,omitempty" xml:"LaunchTemplateConfig,omitempty" type:"Repeated"`
	// The ID of the instance launch template associated with the auto provisioning group. You can invoke the [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) operation to query active instance launch templates. If you specify both a launch template and launch configurations (`LaunchConfiguration.*`), the system preferentially uses the launch template.
	//
	// example:
	//
	// lt-bp1fgzds4bdogu03****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the instance launch template associated with the auto provisioning group. You can invoke the [DescribeLaunchTemplateVersions](https://help.aliyun.com/document_detail/73761.html) operation to query active instance launch template versions.
	//
	// Default value: the default version of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum price of spot instances in the auto provisioning group.
	//
	// > If the settings include both `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice`, the lower value is used.
	//
	// example:
	//
	// 2
	MaxSpotPrice *float32 `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	// The minimum target capacity of the auto provisioning group. Valid values: positive integers.
	//
	// Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	//
	// - If the instance inventory in the current region is less than the value of this parameter, the API call fails and no instances are created.
	//
	// - If the instance inventory in the current region is greater than the value of this parameter, instances are created based on the values of other configured parameters.
	//
	// example:
	//
	// 20
	MinTargetCapacity *string `json:"MinTargetCapacity,omitempty" xml:"MinTargetCapacity,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for creating pay-as-you-go instances. Valid values:
	//
	// - lowest-price: the cost optimization policy. The instance type with the lowest price is selected.
	//
	// - prioritized: the priority-based policy. Instances are created based on the priorities specified by `LaunchTemplateConfig.N.Priority`.
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
	// The detailed capacity configurations for subscription instances.
	PrePaidOptions *CreateAutoProvisioningGroupRequestPrePaidOptions `json:"PrePaidOptions,omitempty" xml:"PrePaidOptions,omitempty" type:"Struct"`
	// The ID of the region where the auto provisioning group resides. You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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
	// The resource pool policy used when creating instances. Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create pay-as-you-go instances.
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	ResourcePoolOptions *CreateAutoProvisioningGroupRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The policy for creating spot instances. Valid values:
	//
	// - lowest-price: the cost optimization policy. The instance type with the lowest price is selected.
	//
	// - diversified: the zone distribution balancing policy. Instances are created in the zones specified by the extended launch templates and evenly distributed across the zones.
	//
	// - capacity-optimized: the capacity optimization policy. The optimal instance types and zones are selected based on the inventory.
	//
	// Default value: lowest-price.
	//
	// example:
	//
	// diversified
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The behavior of spot instances when they are interrupted. Valid values:
	//
	// - stop: stops the instances.
	//
	// - terminate: releases the instances.
	//
	// Default value: terminate.
	//
	// example:
	//
	// terminate
	SpotInstanceInterruptionBehavior *string `json:"SpotInstanceInterruptionBehavior,omitempty" xml:"SpotInstanceInterruptionBehavior,omitempty"`
	// The number of instance types with the lowest prices that are selected by the auto provisioning group to create instances. This parameter takes effect only when `SpotAllocationStrategy` is set to `lowest-price`.
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
	// The list of system disk configurations for the instance.
	SystemDiskConfig []*CreateAutoProvisioningGroupRequestSystemDiskConfig `json:"SystemDiskConfig,omitempty" xml:"SystemDiskConfig,omitempty" type:"Repeated"`
	// The list of tags bound to the auto provisioning group.
	Tag []*CreateAutoProvisioningGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to release the instances in the auto-provisioning group when the group is deleted. Valid values:
	//
	// - true: releases the instances auto provisioning group.
	//
	// - false: retains the instances auto provisioning group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstances *bool `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
	// Specifies whether to release the instances in the auto provisioning group when the group expires. Valid values:
	//
	// - true: releases the instances in the group.
	//
	// - false: only removes the instances from the group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TerminateInstancesWithExpiration *bool `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	// The total target capacity of the auto provisioning group. Valid values: positive integers.
	//
	// The total target capacity must be greater than or equal to the sum of `PayAsYouGoTargetCapacity` (the target capacity of pay-as-you-go instances) and `SpotTargetCapacity` (the target capacity of spot instances).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	TotalTargetCapacity *string `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
	// The time when the auto provisioning group starts. This parameter and `ValidUntil` determine the validity period.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// Default value: the UNIX timestamp when the API operation takes effect.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	ValidFrom *string `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The time when the auto provisioning group expires. This parameter and `ValidFrom` determine the validity period.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// Default value: 2099-12-31T23:59:59Z.
	//
	// example:
	//
	// 2019-06-01T15:10:20Z
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s CreateAutoProvisioningGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequest) GetLaunchConfiguration() *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	return s.LaunchConfiguration
}

func (s *CreateAutoProvisioningGroupRequest) GetAutoProvisioningGroupName() *string {
	return s.AutoProvisioningGroupName
}

func (s *CreateAutoProvisioningGroupRequest) GetAutoProvisioningGroupType() *string {
	return s.AutoProvisioningGroupType
}

func (s *CreateAutoProvisioningGroupRequest) GetCandidateOptions() *CreateAutoProvisioningGroupRequestCandidateOptions {
	return s.CandidateOptions
}

func (s *CreateAutoProvisioningGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAutoProvisioningGroupRequest) GetDataDiskConfig() []*CreateAutoProvisioningGroupRequestDataDiskConfig {
	return s.DataDiskConfig
}

func (s *CreateAutoProvisioningGroupRequest) GetDefaultTargetCapacityType() *string {
	return s.DefaultTargetCapacityType
}

func (s *CreateAutoProvisioningGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAutoProvisioningGroupRequest) GetExcessCapacityTerminationPolicy() *string {
	return s.ExcessCapacityTerminationPolicy
}

func (s *CreateAutoProvisioningGroupRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateAutoProvisioningGroupRequest) GetHibernationOptionsConfigured() *bool {
	return s.HibernationOptionsConfigured
}

func (s *CreateAutoProvisioningGroupRequest) GetLaunchTemplateConfig() []*CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	return s.LaunchTemplateConfig
}

func (s *CreateAutoProvisioningGroupRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateAutoProvisioningGroupRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *CreateAutoProvisioningGroupRequest) GetMaxSpotPrice() *float32 {
	return s.MaxSpotPrice
}

func (s *CreateAutoProvisioningGroupRequest) GetMinTargetCapacity() *string {
	return s.MinTargetCapacity
}

func (s *CreateAutoProvisioningGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAutoProvisioningGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoProvisioningGroupRequest) GetPayAsYouGoAllocationStrategy() *string {
	return s.PayAsYouGoAllocationStrategy
}

func (s *CreateAutoProvisioningGroupRequest) GetPayAsYouGoTargetCapacity() *string {
	return s.PayAsYouGoTargetCapacity
}

func (s *CreateAutoProvisioningGroupRequest) GetPrePaidOptions() *CreateAutoProvisioningGroupRequestPrePaidOptions {
	return s.PrePaidOptions
}

func (s *CreateAutoProvisioningGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAutoProvisioningGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoProvisioningGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoProvisioningGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoProvisioningGroupRequest) GetResourcePoolOptions() *CreateAutoProvisioningGroupRequestResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *CreateAutoProvisioningGroupRequest) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *CreateAutoProvisioningGroupRequest) GetSpotInstanceInterruptionBehavior() *string {
	return s.SpotInstanceInterruptionBehavior
}

func (s *CreateAutoProvisioningGroupRequest) GetSpotInstancePoolsToUseCount() *int32 {
	return s.SpotInstancePoolsToUseCount
}

func (s *CreateAutoProvisioningGroupRequest) GetSpotTargetCapacity() *string {
	return s.SpotTargetCapacity
}

func (s *CreateAutoProvisioningGroupRequest) GetSystemDiskConfig() []*CreateAutoProvisioningGroupRequestSystemDiskConfig {
	return s.SystemDiskConfig
}

func (s *CreateAutoProvisioningGroupRequest) GetTag() []*CreateAutoProvisioningGroupRequestTag {
	return s.Tag
}

func (s *CreateAutoProvisioningGroupRequest) GetTerminateInstances() *bool {
	return s.TerminateInstances
}

func (s *CreateAutoProvisioningGroupRequest) GetTerminateInstancesWithExpiration() *bool {
	return s.TerminateInstancesWithExpiration
}

func (s *CreateAutoProvisioningGroupRequest) GetTotalTargetCapacity() *string {
	return s.TotalTargetCapacity
}

func (s *CreateAutoProvisioningGroupRequest) GetValidFrom() *string {
	return s.ValidFrom
}

func (s *CreateAutoProvisioningGroupRequest) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *CreateAutoProvisioningGroupRequest) SetLaunchConfiguration(v *CreateAutoProvisioningGroupRequestLaunchConfiguration) *CreateAutoProvisioningGroupRequest {
	s.LaunchConfiguration = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetAutoProvisioningGroupName(v string) *CreateAutoProvisioningGroupRequest {
	s.AutoProvisioningGroupName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetAutoProvisioningGroupType(v string) *CreateAutoProvisioningGroupRequest {
	s.AutoProvisioningGroupType = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetCandidateOptions(v *CreateAutoProvisioningGroupRequestCandidateOptions) *CreateAutoProvisioningGroupRequest {
	s.CandidateOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetClientToken(v string) *CreateAutoProvisioningGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetDataDiskConfig(v []*CreateAutoProvisioningGroupRequestDataDiskConfig) *CreateAutoProvisioningGroupRequest {
	s.DataDiskConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetDefaultTargetCapacityType(v string) *CreateAutoProvisioningGroupRequest {
	s.DefaultTargetCapacityType = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetDescription(v string) *CreateAutoProvisioningGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetExcessCapacityTerminationPolicy(v string) *CreateAutoProvisioningGroupRequest {
	s.ExcessCapacityTerminationPolicy = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetExecutionMode(v string) *CreateAutoProvisioningGroupRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetHibernationOptionsConfigured(v bool) *CreateAutoProvisioningGroupRequest {
	s.HibernationOptionsConfigured = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetLaunchTemplateConfig(v []*CreateAutoProvisioningGroupRequestLaunchTemplateConfig) *CreateAutoProvisioningGroupRequest {
	s.LaunchTemplateConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetLaunchTemplateId(v string) *CreateAutoProvisioningGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetLaunchTemplateVersion(v string) *CreateAutoProvisioningGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetMaxSpotPrice(v float32) *CreateAutoProvisioningGroupRequest {
	s.MaxSpotPrice = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetMinTargetCapacity(v string) *CreateAutoProvisioningGroupRequest {
	s.MinTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetOwnerAccount(v string) *CreateAutoProvisioningGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetOwnerId(v int64) *CreateAutoProvisioningGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetPayAsYouGoAllocationStrategy(v string) *CreateAutoProvisioningGroupRequest {
	s.PayAsYouGoAllocationStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetPayAsYouGoTargetCapacity(v string) *CreateAutoProvisioningGroupRequest {
	s.PayAsYouGoTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetPrePaidOptions(v *CreateAutoProvisioningGroupRequestPrePaidOptions) *CreateAutoProvisioningGroupRequest {
	s.PrePaidOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetRegionId(v string) *CreateAutoProvisioningGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetResourceGroupId(v string) *CreateAutoProvisioningGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetResourceOwnerAccount(v string) *CreateAutoProvisioningGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetResourceOwnerId(v int64) *CreateAutoProvisioningGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetResourcePoolOptions(v *CreateAutoProvisioningGroupRequestResourcePoolOptions) *CreateAutoProvisioningGroupRequest {
	s.ResourcePoolOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetSpotAllocationStrategy(v string) *CreateAutoProvisioningGroupRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetSpotInstanceInterruptionBehavior(v string) *CreateAutoProvisioningGroupRequest {
	s.SpotInstanceInterruptionBehavior = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetSpotInstancePoolsToUseCount(v int32) *CreateAutoProvisioningGroupRequest {
	s.SpotInstancePoolsToUseCount = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetSpotTargetCapacity(v string) *CreateAutoProvisioningGroupRequest {
	s.SpotTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetSystemDiskConfig(v []*CreateAutoProvisioningGroupRequestSystemDiskConfig) *CreateAutoProvisioningGroupRequest {
	s.SystemDiskConfig = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetTag(v []*CreateAutoProvisioningGroupRequestTag) *CreateAutoProvisioningGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetTerminateInstances(v bool) *CreateAutoProvisioningGroupRequest {
	s.TerminateInstances = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetTerminateInstancesWithExpiration(v bool) *CreateAutoProvisioningGroupRequest {
	s.TerminateInstancesWithExpiration = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetTotalTargetCapacity(v string) *CreateAutoProvisioningGroupRequest {
	s.TotalTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetValidFrom(v string) *CreateAutoProvisioningGroupRequest {
	s.ValidFrom = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) SetValidUntil(v string) *CreateAutoProvisioningGroupRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequest) Validate() error {
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
	if s.ResourcePoolOptions != nil {
		if err := s.ResourcePoolOptions.Validate(); err != nil {
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

type CreateAutoProvisioningGroupRequestLaunchConfiguration struct {
	// > This parameter is in invitational preview and is not supported.
	Arn []*CreateAutoProvisioningGroupRequestLaunchConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// The automatic release time of the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// - If the value of seconds (ss) is not 00, the release time is automatically set to the start of the current minute (mm).
	//
	// - The earliest release time is half an hour after the current time.
	//
	// - The latest release time cannot be more than three years after the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information about the performance of the instance, see the "Standard mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information about the performance of the instance, see the "Unlimited mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// Default value: null.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The extended data disk configurations for the launch template.
	DataDisk []*CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the instance. Limits:
	//
	// - The hostname cannot start or end with a period (.) or a hyphen (-). It cannot contain consecutive periods (.) or hyphens (-).
	//
	// - For Windows instances, the hostname must be 2 to 15 characters in length. It cannot contain periods (.) and cannot contain only digits. It can contain letters, digits, and hyphens (-).
	//
	// - For other instances such as Linux instances, the hostname must be 2 to 64 characters in length. It can contain multiple periods (.). The characters between two consecutive periods are considered a segment. Each segment can contain letters, digits, and hyphens (-).
	//
	// - You cannot specify both `LaunchConfiguration.HostName` and `LaunchConfiguration.HostNames.N`. Otherwise, an error is returned.
	//
	// - If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// k8s-node-[1,4]-ecshost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The list of hostnames for one or more instances. Limits:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	//
	// - N indicates the number of instances. Valid values of N: 1 to 1000. The value of N must be the same as the value of the TotalTargetCapacity parameter.
	//
	// - Periods (.) and hyphens (-) cannot be used as the first or last characters, and cannot be used consecutively.
	//
	// - If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// ecs-host-01
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with aliyun or acs:. It cannot contain http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image. The image resource selected when the instance is started. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available image resources. If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// m-bp1g7004ksh0oeuc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// Instance_Description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// Default value: the `InstanceId` of the instance.
	//
	// When you create multiple ECS instances, you can batch configure sequential instance names. For more information, see [Configure sequential instance names or hostnames](https://help.aliyun.com/document_detail/196048.html).
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
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
	// > In pay-by-traffic mode, the peak inbound and outbound bandwidths are upper limits and are not guaranteed service level indicators. When resource contention occurs, the peak bandwidths may be limited. If your business requires guaranteed bandwidth, use the pay-by-bandwidth mode.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// - When the outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - When the outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `LaunchConfiguration.InternetMaxBandwidthOut`. Default value: the value of `LaunchConfiguration.InternetMaxBandwidthOut`.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
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
	// The default value is cloud for retired instance types and non-I/O optimized instances, and optimized for other instance types.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// - For Windows instances, this parameter is ignored. Default value: null.
	//
	// - For Linux instances, password logon is disabled by default.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	// ``()`~!@#$%^&*-_+=|{}`[]`:;\\"<>,.?/``
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
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
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// true
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the instance RAM role. You can use the RAM API operation [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created. If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// RAM_Name
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the instance belongs. If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// - Active: enables security hardening. This value is applicable only to public images.
	//
	// - Deactive: disables security hardening. This value is applicable to all image types.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
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
	// The system disk configuration of the instance. If both the launch template and the launch configurations are specified, the launch template takes precedence.
	SystemDisk *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The category of the system disk. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// The default value is cloud for retired instance types and non-I/O optimized instances, and cloud_efficiency for other instance types.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// SystemDisk_Description
	SystemDiskDescription *string `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// Default value: null.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// cloud_ssdSystem
	SystemDiskName *string `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	// The settings for the performance level of the enterprise SSD used as the system disk. Valid values:
	//
	// - PL0 (default): A single enterprise SSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single enterprise SSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single enterprise SSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single enterprise SSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// For more information about how to select enterprise SSD performance levels, see [enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// PL0
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values: 20 to 500. The value of this parameter must be greater than or equal to max{20, the size of the image specified by LaunchConfiguration.ImageId}.
	//
	// Default value: max{40, the size of the image specified by LaunchConfiguration.ImageId}.
	//
	// If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The extended list of tags for the launch template.
	Tag []*CreateAutoProvisioningGroupRequestLaunchConfigurationTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Instance user data of the instance. Instance user data must be encoded in Base64. The maximum size of the raw data is 32 KB. If you specify both a launch template and launch configurations, the system preferentially uses the launch template.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect when you create subscription instances. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false (default): Auto-renewal is disabled.
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
	// - When PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - When PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">When PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The CPU configurations.
	CpuOptions *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The image properties.
	//
	// Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	ImageOptions *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The subscription period of the resource. The unit is specified by PeriodUnit. This parameter is required and takes effect only when you create subscription instances. Valid values:
	//
	// <props="china">
	//
	// - When PeriodUnit is set to Week, valid values of Period are 1, 2, 3, and 4.
	//
	// - When PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">When PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 6, and 12.
	//
	// <props="partner">When PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// <props="china">
	//
	// - Week.
	//
	// - Month (default).
	//
	//
	//
	// <props="intl">Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The scheduling options.
	SchedulerOptions *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	// The security options of the instance.
	SecurityOptions *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: The instance is protected from automatic release for 1 hour after it is created. After the 1-hour protection period ends, the system compares your bid with the marketplace price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// - 0: The instance is not protected after it is created. The system compares your bid with the marketplace price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// Alibaba Cloud sends you a notification by using an ECS system event 5 minutes before the instance is revoked. Spot instances use billing by the second. We recommend that you select an appropriate protection period based on the execution duration of your specific tasks.
	//
	// Take note of the following items when you set this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the one-time synchronization delivery type (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The break mode of the spot instance. Valid values:
	//
	// - Terminate: The instance is released.
	//
	// - Stop: The instance enters the economical mode.
	//
	// For more information about the economical mode, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// Take note of the following items when you set this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the one-time synchronization delivery type (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetArn() []*CreateAutoProvisioningGroupRequestLaunchConfigurationArn {
	return s.Arn
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetDataDisk() []*CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	return s.DataDisk
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetHostName() *string {
	return s.HostName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetHostNames() []*string {
	return s.HostNames
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetPassword() *string {
	return s.Password
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDisk() *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	return s.SystemDisk
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDiskDescription() *string {
	return s.SystemDiskDescription
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDiskName() *string {
	return s.SystemDiskName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetTag() []*CreateAutoProvisioningGroupRequestLaunchConfigurationTag {
	return s.Tag
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetUserData() *string {
	return s.UserData
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetCpuOptions() *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions {
	return s.CpuOptions
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetImageOptions() *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions {
	return s.ImageOptions
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSchedulerOptions() *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions {
	return s.SchedulerOptions
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSecurityOptions() *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetArn(v []*CreateAutoProvisioningGroupRequestLaunchConfigurationArn) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.Arn = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetAutoReleaseTime(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetCreditSpecification(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.CreditSpecification = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetDataDisk(v []*CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.DataDisk = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetDeploymentSetId(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetHostName(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.HostName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetHostNames(v []*string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.HostNames = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetImageFamily(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.ImageFamily = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetImageId(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.ImageId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetInstanceDescription(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.InstanceDescription = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetInstanceName(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.InstanceName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetInternetChargeType(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.InternetChargeType = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetInternetMaxBandwidthIn(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetInternetMaxBandwidthOut(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetIoOptimized(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.IoOptimized = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetKeyPairName(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.KeyPairName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetPassword(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.Password = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetPasswordInherit(v bool) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.PasswordInherit = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetRamRoleName(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.RamRoleName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetResourceGroupId(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSecurityEnhancementStrategy(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSecurityGroupId(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSecurityGroupIds(v []*string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDisk(v *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDisk = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDiskCategory(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDiskDescription(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDiskDescription = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDiskName(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDiskName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDiskPerformanceLevel(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSystemDiskSize(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetTag(v []*CreateAutoProvisioningGroupRequestLaunchConfigurationTag) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.Tag = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetUserData(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.UserData = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetAutoRenew(v bool) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.AutoRenew = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetAutoRenewPeriod(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetCpuOptions(v *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.CpuOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetImageOptions(v *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.ImageOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetPeriod(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.Period = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetPeriodUnit(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSchedulerOptions(v *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SchedulerOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSecurityOptions(v *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SecurityOptions = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSpotDuration(v int32) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SpotDuration = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) SetSpotInterruptionBehavior(v string) *CreateAutoProvisioningGroupRequestLaunchConfiguration {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfiguration) Validate() error {
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

type CreateAutoProvisioningGroupRequestLaunchConfigurationArn struct {
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

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationArn) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationArn) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) SetAssumeRoleFor(v int64) *CreateAutoProvisioningGroupRequestLaunchConfigurationArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) SetRoleType(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationArn {
	s.RoleType = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) SetRolearn(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationArn {
	s.Rolearn = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationArn) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk struct {
	// The ID of the automatic snapshot policy applied to the data disk.
	//
	// Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	//
	// > This parameter is supported only when the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
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
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// If both the launch template and the launch configuration are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: The data disk is released when the instance is released.
	//
	// - false: The data disk is not released when the instance is released.
	//
	// Default value: true.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://. If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// DataDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// /dev/vd1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// Default value: null.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
	//
	// - true: The data disk is encrypted.
	//
	// - false: The data disk is not encrypted.
	//
	// Default value: false.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
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
	// When you create an enterprise SSD as a data disk, set the performance level of the disk. The value of N must be the same as the value of N in LaunchConfiguration.DataDisk.N.Category. Valid values:
	//
	// - PL0: A single enterprise SSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): A single enterprise SSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single enterprise SSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single enterprise SSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// For more information about how to select the performance level of an enterprise SSD, see [Enterprise SSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// If both the launch template and the launch configuration are specified, the launch template takes precedence.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	//
	// > This parameter is supported only when the DiskCategory parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Unit: GiB. Valid values of N: 1 to 16. Valid values:
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: The valid values vary based on the value of LaunchConfiguration.DataDisk.N.PerformanceLevel.
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
	// > The value of this parameter must be greater than or equal to the size of the snapshot specified by the LaunchConfiguration.DataDisk.N.SnapshotId parameter.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create data disk N. Valid values of N: 1 to 16.
	//
	// If this parameter is specified, the LaunchConfiguration.DataDisk.N.Size parameter is ignored, and the size of the created disk is the same as the size of the specified snapshot. You cannot use snapshots that were created on or before July 15, 2013. Otherwise, the request is rejected and an error is returned.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetAutoSnapshotPolicyId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetBurstingEnabled(v bool) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetCategory(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.Category = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetDeleteWithInstance(v bool) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetDescription(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.Description = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetDevice(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.Device = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetDiskName(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetEncryptAlgorithm(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetEncrypted(v bool) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetKmsKeyId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.KmsKeyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetPerformanceLevel(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetProvisionedIops(v int64) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetSize(v int32) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.Size = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) SetSnapshotId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk struct {
	// The ID of the automatic snapshot policy applied to the system disk.
	//
	// Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create an auto provisioning group of the instant delivery type (AutoProvisioningGroupType=instant).
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	//
	// > This parameter is supported only when the SystemDisk.Category parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The encryption algorithm of the system disk. Valid values:
	//
	// - aes-256.
	//
	// - sm4-128.
	//
	// Default value: aes-256.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// > This parameter is not available.
	//
	// example:
	//
	// aes-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt system disk N. Valid values:
	//
	// - true: The system disk is encrypted.
	//
	// - false: The system disk is not encrypted.
	//
	// Default value: false.
	//
	// If both the launch template and the launch configurations are specified, the launch template takes precedence.
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
	// > This parameter is supported only when the SystemDisk.Category parameter is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetBurstingEnabled(v bool) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetEncryptAlgorithm(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetEncrypted(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetKMSKeyId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) SetProvisionedIops(v int64) *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationTag struct {
	// The tag key of the instance. Valid values of N: 1 to 20. If this parameter is specified, it cannot be left empty. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://. If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values of N: 1 to 20. If this parameter is specified, it can be an empty string. The tag value can be up to 128 characters in length. It cannot start with acs:, and cannot contain http:// or https://. If both the launch template and the launch configurations are specified, the launch template takes precedence.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationTag) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationTag) SetKey(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationTag {
	s.Key = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationTag) SetValue(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationTag {
	s.Value = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationTag) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions struct {
	// The number of CPU cores.
	//
	// Default value: For more information, see [Customize CPU options](https://www.alibabacloud.com/help/en/ecs/user-guide/specify-and-view-cpu-options).
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of threads per core. The number of vCPUs of the ECS instance = the value of CpuOptions.Core × the value of CpuOptions.ThreadsPerCore.
	//
	// A value of 1 for CpuOptions.ThreadsPerCore indicates that CPU hyper-threading is disabled.
	//
	// Only specific instance types support the configuration of the number of threads per core.
	//
	// Valid values and default value: For more information, see [Customize CPU options](https://www.alibabacloud.com/help/en/ecs/user-guide/specify-and-view-cpu-options).
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) SetCore(v int32) *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions {
	s.Core = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) SetThreadsPerCore(v int32) *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationCpuOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions struct {
	// Specifies whether the instance that uses this image supports logons by using the ecs-user username. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) SetLoginAsNonRoot(v bool) *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions struct {
	// The ID of the dedicated host cluster.
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host.
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) SetDedicatedHostClusterId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) SetDedicatedHostId(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions struct {
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) SetTrustedSystemMode(v string) *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchConfigurationSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestCandidateOptions struct {
	// Specifies whether to only evaluate the success rate of delivering the target capacity for candidate instance types in the current request. If you set this parameter to true, no scale-out is performed. The system only evaluates the delivery success rate and returns the result in the LaunchResult.CandidateEvaluateLevel field of the response. Default value: false.
	Evaluate *bool `json:"Evaluate,omitempty" xml:"Evaluate,omitempty"`
	// The timeout period for the auto provisioning group to scale out to the target capacity. After the timeout period elapses, no further scale-out attempts are made. Unit: minutes.
	//
	// example:
	//
	// 60
	TimeoutMinutes *int32 `json:"TimeoutMinutes,omitempty" xml:"TimeoutMinutes,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestCandidateOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestCandidateOptions) GetEvaluate() *bool {
	return s.Evaluate
}

func (s *CreateAutoProvisioningGroupRequestCandidateOptions) GetTimeoutMinutes() *int32 {
	return s.TimeoutMinutes
}

func (s *CreateAutoProvisioningGroupRequestCandidateOptions) SetEvaluate(v bool) *CreateAutoProvisioningGroupRequestCandidateOptions {
	s.Evaluate = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestCandidateOptions) SetTimeoutMinutes(v int32) *CreateAutoProvisioningGroupRequestCandidateOptions {
	s.TimeoutMinutes = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestDataDiskConfig struct {
	// The category of the data disk for the instance. You can use this parameter to specify multiple candidate disk categories. The specified order determines the priority of the disk categories. If a disk category is not active, the system automatically switches to the next disk category. Valid values:
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD.
	//
	// -   cloud: basic disk.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestDataDiskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestDataDiskConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestDataDiskConfig) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateAutoProvisioningGroupRequestDataDiskConfig) SetDiskCategory(v string) *CreateAutoProvisioningGroupRequestDataDiskConfig {
	s.DiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestDataDiskConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestLaunchTemplateConfig struct {
	// The list of architecture types for the instance types.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether the instance type is a burstable instance type. Valid values:
	//
	// - Exclude: does not include burstable instance types.
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
	// The list of vCPU counts for the instance types.
	Cores []*int32 `json:"Cores,omitempty" xml:"Cores,omitempty" type:"Repeated"`
	// The list of instance types to be excluded.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The ID of the image. You can use this parameter to specify the image used by the current resource pool. If this parameter is not specified, the image specified by `LaunchConfiguration.ImageId` or the image configured in the launch template is used by default. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available image resources.
	//
	// Note: This parameter is supported only when `AutoProvisioningGroupType` is set to `instant`.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20210425.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The level of the instance family. This parameter is used to filter the instance types that meet the requirements. Valid values:
	//
	// - EntryLevel: entry-level, which refers to shared instance types. These instance types are more cost-effective but cannot guarantee stable computing performance. They are suitable for business scenarios with low CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level. These instance types provide stable performance and dedicated resources, and are suitable for business scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry-level, which refers to burstable instances. These instance types use CPU credits to guarantee computing performance and are suitable for business scenarios with low CPU utilization but occasional CPU bursts. For more information, see [Burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance type in the extended launch template. Valid values of N: 1 to 20. For more information about valid values, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum price of spot instances in the extended launch template.
	//
	// > After `LaunchTemplateConfig` is specified, `LaunchTemplateConfig.N.MaxPrice` is required.
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
	// The list of memory sizes for the instance types.
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
	// The weight of the instance type in the extended launch template. A higher weight indicates that a single instance has more computing power and fewer instances are required. Valid values: greater than 0.
	//
	// You can calculate the weight based on the computing power of the specified instance type and the minimum computing power of a single node in the cluster. For example, if the minimum computing power of a single node is 8 vCPUs and 60 GiB of memory:
	//
	// - You can set the weight of the instance type with 8 vCPUs and 60 GiB of memory to 1.
	//
	// - You can set the weight of the instance type with 16 vCPUs and 120 GiB of memory to 2.
	//
	// example:
	//
	// 2
	WeightedCapacity *float64 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestLaunchTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetArchitectures() []*string {
	return s.Architectures
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetCores() []*int32 {
	return s.Cores
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetMaxPrice() *float64 {
	return s.MaxPrice
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetMaxQuantity() *int32 {
	return s.MaxQuantity
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetMemories() []*float32 {
	return s.Memories
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) GetWeightedCapacity() *float64 {
	return s.WeightedCapacity
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetArchitectures(v []*string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.Architectures = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetBurstablePerformance(v string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetCores(v []*int32) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.Cores = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetExcludedInstanceTypes(v []*string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetImageId(v string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.ImageId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetInstanceFamilyLevel(v string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetInstanceType(v string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.InstanceType = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetMaxPrice(v float64) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.MaxPrice = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetMaxQuantity(v int32) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.MaxQuantity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetMemories(v []*float32) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.Memories = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetPriority(v int32) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.Priority = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetVSwitchId(v string) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.VSwitchId = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) SetWeightedCapacity(v float64) *CreateAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.WeightedCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestLaunchTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestPrePaidOptions struct {
	// The minimum capacity set for different instance types. This parameter is supported only when AutoProvisioningGroupType is set to request.
	SpecifyCapacityDistribution []*CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution `json:"SpecifyCapacityDistribution,omitempty" xml:"SpecifyCapacityDistribution,omitempty" type:"Repeated"`
}

func (s CreateAutoProvisioningGroupRequestPrePaidOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestPrePaidOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptions) GetSpecifyCapacityDistribution() []*CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution {
	return s.SpecifyCapacityDistribution
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptions) SetSpecifyCapacityDistribution(v []*CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) *CreateAutoProvisioningGroupRequestPrePaidOptions {
	s.SpecifyCapacityDistribution = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptions) Validate() error {
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

type CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution struct {
	// The set of instance types. Duplicate instance types are not allowed, and the instance types must be within the scope of LaunchTemplateConfig.InstanceType.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The minimum number of instances to be delivered within the scope of InstanceTypes.
	//
	// > sum(MinTargetCapacity) <= TotalTargetCapacity. The sum of MinTargetCapacity for all instance type sets cannot exceed TotalTargetCapacity. If the MinTargetCapacity requirement for any instance type set cannot be met due to inventory or other reasons, the entire request fails.
	//
	// example:
	//
	// 5
	MinTargetCapacity *int32 `json:"MinTargetCapacity,omitempty" xml:"MinTargetCapacity,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) GetMinTargetCapacity() *int32 {
	return s.MinTargetCapacity
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) SetInstanceTypes(v []*string) *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution {
	s.InstanceTypes = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) SetMinTargetCapacity(v int32) *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution {
	s.MinTargetCapacity = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestPrePaidOptionsSpecifyCapacityDistribution) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestResourcePoolOptions struct {
	// The list of private pool IDs. You can specify 1 to 20 IDs.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// The resource pool policy. Resource pools include private pools generated after the elasticity assurance or capacity reservation service takes effect, as well as the public pool. Valid values:
	//
	// - PrivatePoolFirst: Private pools take precedence. If ResourcePoolOptions.PrivatePoolIds is specified, the specified private pools are used first. If no private pools are specified or the specified private pools have insufficient capacity, the system automatically matches open private pools. If no eligible private pools are available, the public pool is used to create instances.
	//
	// - PrivatePoolOnly: Only private pools are used. If you select this policy, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pools have insufficient capacity, the instances fail to start.
	//
	//
	//
	// - PublicPoolOnly: The public pool is used to create instances.
	//
	// Default value: PublicPoolOnly.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *CreateAutoProvisioningGroupRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateAutoProvisioningGroupRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *CreateAutoProvisioningGroupRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *CreateAutoProvisioningGroupRequestResourcePoolOptions) SetStrategy(v string) *CreateAutoProvisioningGroupRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestSystemDiskConfig struct {
	// The category of the system disk for the instance. You can use this parameter to specify multiple candidate disk categories. The specified order determines the priority of the disk categories. If a disk category is not active, the system automatically switches to the next disk category. Valid values:
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD.
	//
	// -   cloud: basic disk.
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestSystemDiskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestSystemDiskConfig) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestSystemDiskConfig) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateAutoProvisioningGroupRequestSystemDiskConfig) SetDiskCategory(v string) *CreateAutoProvisioningGroupRequestSystemDiskConfig {
	s.DiskCategory = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestSystemDiskConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAutoProvisioningGroupRequestTag struct {
	// The tag key of the auto provisioning group.
	//
	// Valid values of N: 1 to 20. If this parameter is specified, it cannot be left empty. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the auto provisioning group.
	//
	// Valid values of N: 1 to 20. If this parameter is specified, it can be an empty string. The tag value can be up to 128 characters in length. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoProvisioningGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoProvisioningGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoProvisioningGroupRequestTag) SetKey(v string) *CreateAutoProvisioningGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestTag) SetValue(v string) *CreateAutoProvisioningGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAutoProvisioningGroupRequestTag) Validate() error {
	return dara.Validate(s)
}

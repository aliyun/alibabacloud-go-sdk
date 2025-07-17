// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableScalingGroupRequest interface {
  dara.Model
  String() string
  GoString() string
  SetActiveScalingConfigurationId(v string) *EnableScalingGroupRequest
  GetActiveScalingConfigurationId() *string 
  SetInstanceIds(v []*string) *EnableScalingGroupRequest
  GetInstanceIds() []*string 
  SetLaunchTemplateId(v string) *EnableScalingGroupRequest
  GetLaunchTemplateId() *string 
  SetLaunchTemplateOverrides(v []*EnableScalingGroupRequestLaunchTemplateOverrides) *EnableScalingGroupRequest
  GetLaunchTemplateOverrides() []*EnableScalingGroupRequestLaunchTemplateOverrides 
  SetLaunchTemplateVersion(v string) *EnableScalingGroupRequest
  GetLaunchTemplateVersion() *string 
  SetLoadBalancerWeights(v []*int32) *EnableScalingGroupRequest
  GetLoadBalancerWeights() []*int32 
  SetOwnerAccount(v string) *EnableScalingGroupRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableScalingGroupRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableScalingGroupRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableScalingGroupRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableScalingGroupRequest
  GetResourceOwnerId() *int64 
  SetScalingGroupId(v string) *EnableScalingGroupRequest
  GetScalingGroupId() *string 
}

type EnableScalingGroupRequest struct {
  // The ID of the scaling configuration that you want to enable in the scaling group.
  // 
  // example:
  // 
  // asc-bp1ffogfdauy0nu5****
  ActiveScalingConfigurationId *string `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
  // The IDs of the ECS instances that you want to add to the scaling group after the scaling group is enabled.
  // 
  // Before you add ECS instances to the scaling group, make sure that the instances meet the following requirements:
  // 
  // 	- The instances must reside in the same region as the scaling group.
  // 
  // 	- The instances must be in the Running state.
  // 
  // 	- The instances do not belong to another scaling group.
  // 
  // 	- The instances are billed on a subscription or pay-as-you-go basis, or the instances are preemptible instances.
  // 
  // 	- If you specify VswitchID for the scaling group, the instances must share the same VPC as the scaling group.
  // 
  // 	- If you do not specify VswitchID for the scaling group, the instances must use the classic network.
  InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
  // The ID of the launch template that is used by Auto Scaling to create ECS instances.
  // 
  // example:
  // 
  // lt-m5e3ofjr1zn1aw7****
  LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
  // The information about the instance types that you want to extend in the launch template.
  LaunchTemplateOverrides []*EnableScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
  // The version number of the launch template. Valid values:
  // 
  // 	- A fixed template version number.
  // 
  // 	- Default: The default template version is always used.
  // 
  // 	- Latest: The latest template version is always used.
  // 
  // example:
  // 
  // Default
  LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
  // The weights of ECS instances or elastic container instances as backend servers.
  // 
  // Default value: 50.
  LoadBalancerWeights []*int32 `json:"LoadBalancerWeights,omitempty" xml:"LoadBalancerWeights,omitempty" type:"Repeated"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the scaling group.
  // 
  // example:
  // 
  // cn-qingdao
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The ID of the scaling group.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // asg-bp14wlu85wrpchm0****
  ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s EnableScalingGroupRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableScalingGroupRequest) GoString() string {
  return s.String()
}

func (s *EnableScalingGroupRequest) GetActiveScalingConfigurationId() *string  {
  return s.ActiveScalingConfigurationId
}

func (s *EnableScalingGroupRequest) GetInstanceIds() []*string  {
  return s.InstanceIds
}

func (s *EnableScalingGroupRequest) GetLaunchTemplateId() *string  {
  return s.LaunchTemplateId
}

func (s *EnableScalingGroupRequest) GetLaunchTemplateOverrides() []*EnableScalingGroupRequestLaunchTemplateOverrides  {
  return s.LaunchTemplateOverrides
}

func (s *EnableScalingGroupRequest) GetLaunchTemplateVersion() *string  {
  return s.LaunchTemplateVersion
}

func (s *EnableScalingGroupRequest) GetLoadBalancerWeights() []*int32  {
  return s.LoadBalancerWeights
}

func (s *EnableScalingGroupRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableScalingGroupRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableScalingGroupRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableScalingGroupRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableScalingGroupRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableScalingGroupRequest) GetScalingGroupId() *string  {
  return s.ScalingGroupId
}

func (s *EnableScalingGroupRequest) SetActiveScalingConfigurationId(v string) *EnableScalingGroupRequest {
  s.ActiveScalingConfigurationId = &v
  return s
}

func (s *EnableScalingGroupRequest) SetInstanceIds(v []*string) *EnableScalingGroupRequest {
  s.InstanceIds = v
  return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateId(v string) *EnableScalingGroupRequest {
  s.LaunchTemplateId = &v
  return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateOverrides(v []*EnableScalingGroupRequestLaunchTemplateOverrides) *EnableScalingGroupRequest {
  s.LaunchTemplateOverrides = v
  return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateVersion(v string) *EnableScalingGroupRequest {
  s.LaunchTemplateVersion = &v
  return s
}

func (s *EnableScalingGroupRequest) SetLoadBalancerWeights(v []*int32) *EnableScalingGroupRequest {
  s.LoadBalancerWeights = v
  return s
}

func (s *EnableScalingGroupRequest) SetOwnerAccount(v string) *EnableScalingGroupRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableScalingGroupRequest) SetOwnerId(v int64) *EnableScalingGroupRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableScalingGroupRequest) SetRegionId(v string) *EnableScalingGroupRequest {
  s.RegionId = &v
  return s
}

func (s *EnableScalingGroupRequest) SetResourceOwnerAccount(v string) *EnableScalingGroupRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableScalingGroupRequest) SetResourceOwnerId(v int64) *EnableScalingGroupRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableScalingGroupRequest) SetScalingGroupId(v string) *EnableScalingGroupRequest {
  s.ScalingGroupId = &v
  return s
}

func (s *EnableScalingGroupRequest) Validate() error {
  return dara.Validate(s)
}

type EnableScalingGroupRequestLaunchTemplateOverrides struct {
  // The instance type. If you want to scale instances based on instance type weights in the scaling group, you must specify `LaunchTemplateOverrides.WeightedCapacity` after you specify this parameter.
  // 
  // The instance type specified by using this parameter overwrites the instance type of the launch template.
  // 
  // >  This parameter takes effect only if you specify LaunchTemplateId.
  // 
  // You can use this parameter to specify any instance types that are available for purchase.
  // 
  // example:
  // 
  // ecs.c5.xlarge
  InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
  // The weight of the instance type. If you want to scale instances based on instance type weights in the scaling group, you must specify this parameter after you specify `LaunchTemplateOverrides.InstanceType`.
  // 
  // The weight specifies the capacity of an instance of the specified instance type in the scaling group. A higher weight specifies that a smaller number of instances of the specified instance type are required to meet the expected capacity requirement.
  // 
  // Performance metrics such as the number of vCPUs and the memory size of each instance type may vary. You can specify different weights for different instance types based on your business requirements.
  // 
  // Example:
  // 
  // 	- Current capacity: 0
  // 
  // 	- Expected capacity: 6
  // 
  // 	- Capacity of ecs.c5.xlarge: 4
  // 
  // To reach the expected capacity, Auto Scaling must scale out two instances of ecs.c5.xlarge.
  // 
  // >  The total capacity of the scaling group is constrained and cannot surpass the combined total of the maximum group size defined by MaxSize and the highest weight assigned to any instance type.
  // 
  // Valid values of WeightedCapacity: 1 to 500.
  // 
  // example:
  // 
  // 4
  WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s EnableScalingGroupRequestLaunchTemplateOverrides) String() string {
  return dara.Prettify(s)
}

func (s EnableScalingGroupRequestLaunchTemplateOverrides) GoString() string {
  return s.String()
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) GetWeightedCapacity() *int32  {
  return s.WeightedCapacity
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *EnableScalingGroupRequestLaunchTemplateOverrides {
  s.InstanceType = &v
  return s
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *EnableScalingGroupRequestLaunchTemplateOverrides {
  s.WeightedCapacity = &v
  return s
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) Validate() error {
  return dara.Validate(s)
}


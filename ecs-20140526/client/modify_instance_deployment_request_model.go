// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAffinity(v string) *ModifyInstanceDeploymentRequest
	GetAffinity() *string
	SetDedicatedHostClusterId(v string) *ModifyInstanceDeploymentRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *ModifyInstanceDeploymentRequest
	GetDedicatedHostId() *string
	SetDeploymentSetGroupNo(v int32) *ModifyInstanceDeploymentRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *ModifyInstanceDeploymentRequest
	GetDeploymentSetId() *string
	SetForce(v bool) *ModifyInstanceDeploymentRequest
	GetForce() *bool
	SetInstanceId(v string) *ModifyInstanceDeploymentRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyInstanceDeploymentRequest
	GetInstanceType() *string
	SetMigrationType(v string) *ModifyInstanceDeploymentRequest
	GetMigrationType() *string
	SetOwnerAccount(v string) *ModifyInstanceDeploymentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceDeploymentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceDeploymentRequest
	GetRegionId() *string
	SetRemoveFromDeploymentSet(v bool) *ModifyInstanceDeploymentRequest
	GetRemoveFromDeploymentSet() *bool
	SetResourceOwnerAccount(v string) *ModifyInstanceDeploymentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceDeploymentRequest
	GetResourceOwnerId() *int64
	SetTenancy(v string) *ModifyInstanceDeploymentRequest
	GetTenancy() *string
}

type ModifyInstanceDeploymentRequest struct {
	// Specifies whether the instance is associated with the dedicated host. Valid values:
	//
	// - host: The instance is associated with the dedicated host. When an instance that has economical mode enabled is restarted after being stopped, the instance is still deployed on the original dedicated host.
	//
	// - default: The instance is not associated with the dedicated host. When an instance that has economical mode enabled is restarted after being stopped, if the resources of the original dedicated host are insufficient, the instance can be migrated to another dedicated host in the automatic deployment resource pool.
	//
	// Default value when migrating an instance from a shared host to a dedicated host: default.
	//
	// example:
	//
	// host
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-bp67acfmxazb4ph****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host. You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) to query available dedicated hosts.
	//
	// When you modify the host of an ECS instance (migrate the instance from a shared host to a dedicated host or between dedicated hosts):
	//
	// - To migrate the instance to a specified dedicated host, you must specify this parameter.
	//
	// - To migrate the instance to a dedicated host that is automatically selected by the system, you must set this parameter to empty and set the `Tenancy` parameter to host.
	//
	// For more information about the automatic deployment feature, see [Features of dedicated hosts](https://help.aliyun.com/document_detail/118938.html).
	//
	// example:
	//
	// dh-bp67acfmxazb4ph****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The group number of the instance in the deployment set when the deployment set uses the availability group strategy (AvailabilityGroup). Valid values: 1 to 7.
	//
	// > If you change the deployment set of an ECS instance and the deployment set uses the availability group strategy (`AvailablilityGroup`), the system automatically distributes ECS instances evenly across groups when this parameter is not specified. If you specify the same deployment set that the instance currently belongs to, the system also redistributes ECS instances evenly across groups.
	//
	// example:
	//
	// 3
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set.
	//
	// This parameter is required when you add an ECS instance to a deployment set or change the deployment set of an ECS instance.
	//
	// > When you modify dedicated host-related parameters (`Tenancy`, `Affinity`, and `DedicatedHostId`), you cannot modify the deployment set at the same time.
	//
	// example:
	//
	// ds-bp67acfmxazb4ph****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// Specifies whether to forcefully change the host when the instance is added to a deployment set. Valid values:
	//
	//
	//
	// - true: Allows the operation. Allows restarting ECS instances in the Running or Stopped state. Stopped instances do not include pay-as-you-go ECS instances that have economical mode enabled.
	//
	//     > If the specified ECS instance has local disks attached, the local disks are also forcefully replaced. This may cause data loss on the local disks during host replacement. Proceed with caution.
	//
	// - false: Does not allow the operation. The instance is added to the deployment set only on the current host. This may cause the deployment set change to fail.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target instance type of the ECS instance. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// When you modify the host of an ECS instance, you can also change ECS instance type. The target instance type must match the specifications of the specified dedicated host. For more information, see [Dedicated host types](https://help.aliyun.com/document_detail/68564.html).
	//
	// - To change ECS instance type, you must specify the dedicated host ID by setting the `DedicatedHostId` parameter.
	//
	// - You cannot change ECS instance type when using the automatic deployment feature to migrate an ECS instance.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to stop ECS instance before migrating it to the destination dedicated host. Valid values:
	//
	// - reboot: Stops ECS instance before migration.
	//
	// - live: Migrates ECS instance without stopping it. In this case, you must specify the DedicatedHostId parameter. This value does not support changing ECS instance type while migrating ECS instance.
	//
	// Default value: reboot.
	//
	// example:
	//
	// live
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to remove the selected instance from the selected deployment set. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// > When this parameter is set to true, you must specify the InstanceId and DeploymentSetId that have an ownership relationship.
	//
	// example:
	//
	// false
	RemoveFromDeploymentSet *bool   `json:"RemoveFromDeploymentSet,omitempty" xml:"RemoveFromDeploymentSet,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether the instance is deployed on a dedicated host. Valid values: host. The instance is deployed only on a dedicated host.
	//
	// example:
	//
	// host
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
}

func (s ModifyInstanceDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceDeploymentRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *ModifyInstanceDeploymentRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyInstanceDeploymentRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyInstanceDeploymentRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *ModifyInstanceDeploymentRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyInstanceDeploymentRequest) GetForce() *bool {
	return s.Force
}

func (s *ModifyInstanceDeploymentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceDeploymentRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyInstanceDeploymentRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ModifyInstanceDeploymentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceDeploymentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceDeploymentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceDeploymentRequest) GetRemoveFromDeploymentSet() *bool {
	return s.RemoveFromDeploymentSet
}

func (s *ModifyInstanceDeploymentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceDeploymentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceDeploymentRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *ModifyInstanceDeploymentRequest) SetAffinity(v string) *ModifyInstanceDeploymentRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDedicatedHostClusterId(v string) *ModifyInstanceDeploymentRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDedicatedHostId(v string) *ModifyInstanceDeploymentRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDeploymentSetGroupNo(v int32) *ModifyInstanceDeploymentRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDeploymentSetId(v string) *ModifyInstanceDeploymentRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetForce(v bool) *ModifyInstanceDeploymentRequest {
	s.Force = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetInstanceId(v string) *ModifyInstanceDeploymentRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetInstanceType(v string) *ModifyInstanceDeploymentRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetMigrationType(v string) *ModifyInstanceDeploymentRequest {
	s.MigrationType = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetOwnerAccount(v string) *ModifyInstanceDeploymentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetOwnerId(v int64) *ModifyInstanceDeploymentRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetRegionId(v string) *ModifyInstanceDeploymentRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetRemoveFromDeploymentSet(v bool) *ModifyInstanceDeploymentRequest {
	s.RemoveFromDeploymentSet = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetResourceOwnerAccount(v string) *ModifyInstanceDeploymentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetResourceOwnerId(v int64) *ModifyInstanceDeploymentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetTenancy(v string) *ModifyInstanceDeploymentRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) Validate() error {
	return dara.Validate(s)
}

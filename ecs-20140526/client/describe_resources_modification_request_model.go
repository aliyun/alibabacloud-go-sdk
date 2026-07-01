// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesModificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*string) *DescribeResourcesModificationRequest
	GetConditions() []*string
	SetCores(v int32) *DescribeResourcesModificationRequest
	GetCores() *int32
	SetDestinationResource(v string) *DescribeResourcesModificationRequest
	GetDestinationResource() *string
	SetInstanceType(v string) *DescribeResourcesModificationRequest
	GetInstanceType() *string
	SetMemory(v float32) *DescribeResourcesModificationRequest
	GetMemory() *float32
	SetMigrateAcrossZone(v bool) *DescribeResourcesModificationRequest
	GetMigrateAcrossZone() *bool
	SetOperationType(v string) *DescribeResourcesModificationRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *DescribeResourcesModificationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeResourcesModificationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeResourcesModificationRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeResourcesModificationRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeResourcesModificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResourcesModificationRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeResourcesModificationRequest
	GetZoneId() *string
}

type DescribeResourcesModificationRequest struct {
	// The list of conditions.
	Conditions []*string `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The number of vCPU kernels of the target instance type. For valid values, see [Instance family](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter takes effect only when DestinationResource is set to InstanceType.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The type of the resource to be changed. Valid values:
	//
	//
	//
	// - InstanceType: instance type.
	//
	// - SystemDisk: system disk type.
	//
	//   If you set this parameter to SystemDisk, you must also specify the InstanceType parameter to indicate the disk type required by the target instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	DestinationResource *string `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	// The target instance type. For more information, see [Instance family](https://help.aliyun.com/document_detail/25378.html). You can also call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// If DestinationResource is set to SystemDisk, you must also specify the InstanceType parameter to indicate the disk type required by the target instance type.
	//
	// example:
	//
	// ecs.g7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory size of the target instance type. Unit: GiB. For valid values, see [Instance family](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter takes effect only when DestinationResource is set to InstanceType.
	//
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Specifies whether cross-cluster instance type changes are supported. Valid values:
	//
	//
	//
	// - true: Supported.
	//
	// - false: Not supported.
	//
	// Default value: false.
	//
	// If the MigrateAcrossZone parameter is set to true and you upgrade or downgrade the Elastic Compute Service instance based on the returned information, take note of the following items:
	//
	//
	//
	// - VPC-type instances: For [retired instance types](https://help.aliyun.com/document_detail/55263.html), when a non-I/O optimized instance is changed to an I/O optimized instance, the disk device names and software authorization codes of the server are changed. For Linux instances, basic disks (cloud) are identified as xvda or xvdb. Ultra disks (cloud_efficiency) and standard SSDs (cloud_ssd) are identified as vda or vdb.
	//
	// example:
	//
	// true
	MigrateAcrossZone *bool `json:"MigrateAcrossZone,omitempty" xml:"MigrateAcrossZone,omitempty"`
	// The type of the Upgrade/Downgrade operation.
	//
	// - Valid values for subscription resources:
	//
	//     - Upgrade: upgrades resources.
	//
	//     - Downgrade: downgrades resources.
	//
	//     - RenewDowngrade: downgrades resources upon renewal.
	//
	//     - RenewModify: renewal with specification change for expired instances.
	//
	// - Valid value for pay-as-you-go resources: Upgrade.
	//
	// Default value: Upgrade.
	//
	// example:
	//
	// Upgrade
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance whose configuration you want to change. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID (InstanceId) of the instance whose instance type or system disk type you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the target zone.
	//
	// Specify this parameter when you want to change the instance type across zones.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourcesModificationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationRequest) GetConditions() []*string {
	return s.Conditions
}

func (s *DescribeResourcesModificationRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeResourcesModificationRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeResourcesModificationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeResourcesModificationRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeResourcesModificationRequest) GetMigrateAcrossZone() *bool {
	return s.MigrateAcrossZone
}

func (s *DescribeResourcesModificationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeResourcesModificationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeResourcesModificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResourcesModificationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourcesModificationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourcesModificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResourcesModificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourcesModificationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeResourcesModificationRequest) SetConditions(v []*string) *DescribeResourcesModificationRequest {
	s.Conditions = v
	return s
}

func (s *DescribeResourcesModificationRequest) SetCores(v int32) *DescribeResourcesModificationRequest {
	s.Cores = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetDestinationResource(v string) *DescribeResourcesModificationRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetInstanceType(v string) *DescribeResourcesModificationRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetMemory(v float32) *DescribeResourcesModificationRequest {
	s.Memory = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetMigrateAcrossZone(v bool) *DescribeResourcesModificationRequest {
	s.MigrateAcrossZone = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOperationType(v string) *DescribeResourcesModificationRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOwnerAccount(v string) *DescribeResourcesModificationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOwnerId(v int64) *DescribeResourcesModificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetRegionId(v string) *DescribeResourcesModificationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceId(v string) *DescribeResourcesModificationRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceOwnerAccount(v string) *DescribeResourcesModificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceOwnerId(v int64) *DescribeResourcesModificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetZoneId(v string) *DescribeResourcesModificationRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) Validate() error {
	return dara.Validate(s)
}

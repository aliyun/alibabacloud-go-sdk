// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBurstingEnabled(v bool) *ModifyDiskAttributeRequest
	GetBurstingEnabled() *bool
	SetDeleteAutoSnapshot(v bool) *ModifyDiskAttributeRequest
	GetDeleteAutoSnapshot() *bool
	SetDeleteWithInstance(v bool) *ModifyDiskAttributeRequest
	GetDeleteWithInstance() *bool
	SetDescription(v string) *ModifyDiskAttributeRequest
	GetDescription() *string
	SetDiskId(v string) *ModifyDiskAttributeRequest
	GetDiskId() *string
	SetDiskIds(v []*string) *ModifyDiskAttributeRequest
	GetDiskIds() []*string
	SetDiskName(v string) *ModifyDiskAttributeRequest
	GetDiskName() *string
	SetEnableAutoSnapshot(v bool) *ModifyDiskAttributeRequest
	GetEnableAutoSnapshot() *bool
	SetOwnerAccount(v string) *ModifyDiskAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDiskAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDiskAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDiskAttributeRequest struct {
	// Specifies whether to enable the performance burst feature for a cloud disk that supports this feature. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// > An error is returned if you specify this parameter for a cloud disk that does not support the performance burst feature.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Specifies whether to delete the automatic snapshots of the cloud disk when the cloud disk is deleted. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: null. A value of null indicates that the current value is not changed.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Specifies whether to release the cloud disk when the instance is released. Default value: null. A value of null indicates that the current value is not changed.
	//
	// <props="china">This parameter is not supported for cloud disks that have the multi-attach feature enabled.
	//
	// An error is returned if you set the `DeleteWithInstance` parameter to `false` in the following scenarios:
	//
	//
	//
	// - The category of the cloud disk is local disk (ephemeral).
	//
	// - The category of the cloud disk is basic cloud disk (cloud) and the cloud disk is not removable (Portable=false).
	//
	// 	Warning: If you set this parameter to false, the cloud disk is not released when the instance is released. However, if the ECS instance to which the cloud disk is attached is locked for security reasons and the "LockReason" : "security" state is added to OperationLocks, the DeleteWithInstance attribute of the cloud disk is ignored and the cloud disk is released along with the instance.</warning>
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud disk whose attributes you want to modify.
	//
	// > You cannot specify both the `DiskId` and `DiskIds.N` parameters. Specify one of them based on your business requirements.
	//
	// example:
	//
	// d-bp1famypsnar20bv****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The IDs of the cloud disks whose attributes you want to modify. Valid values of N: 0 to 100.
	//
	// > You cannot specify both the `DiskId` and `DiskIds.N` parameters. Specify one of them based on your business requirements.
	//
	// example:
	//
	// d-bp1famypsnar20bv****
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The name of the cloud disk. The name must be 2 to 128 characters in length and can contain letters and digits in Unicode. It can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// MyDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to enable the automatic snapshot policy feature for the cloud disk. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: null. A value of null indicates that the current value is not changed.
	//
	// > This parameter is deprecated. The automatic snapshot policy feature is enabled by default for newly created cloud disks. You only need to apply an automatic snapshot policy to the cloud disk to use the feature.
	//
	// example:
	//
	// true
	EnableAutoSnapshot *bool   `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDiskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskAttributeRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDiskAttributeRequest) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *ModifyDiskAttributeRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyDiskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDiskAttributeRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyDiskAttributeRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *ModifyDiskAttributeRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyDiskAttributeRequest) GetEnableAutoSnapshot() *bool {
	return s.EnableAutoSnapshot
}

func (s *ModifyDiskAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskAttributeRequest) SetBurstingEnabled(v bool) *ModifyDiskAttributeRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDeleteAutoSnapshot(v bool) *ModifyDiskAttributeRequest {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDeleteWithInstance(v bool) *ModifyDiskAttributeRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDescription(v string) *ModifyDiskAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskId(v string) *ModifyDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskIds(v []*string) *ModifyDiskAttributeRequest {
	s.DiskIds = v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskName(v string) *ModifyDiskAttributeRequest {
	s.DiskName = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetEnableAutoSnapshot(v bool) *ModifyDiskAttributeRequest {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetOwnerAccount(v string) *ModifyDiskAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetOwnerId(v int64) *ModifyDiskAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetRegionId(v string) *ModifyDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDiskAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetResourceOwnerId(v int64) *ModifyDiskAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) Validate() error {
	return dara.Validate(s)
}

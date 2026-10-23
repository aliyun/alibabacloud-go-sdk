// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v bool) *DetachDiskRequest
	GetDeleteWithInstance() *bool
	SetDiskId(v string) *DetachDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *DetachDiskRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DetachDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DetachDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachDiskRequest
	GetResourceOwnerId() *int64
}

type DetachDiskRequest struct {
	// Specifies whether to release the system cloud disk or data cloud disk when the ECS instance is released. This parameter sets the automatic release attribute when you unmount a system cloud disk or data cloud disk.
	//
	// - true: The disk is released.
	//
	// - false: The disk is not released. The cloud disk is converted to a pay-as-you-go data cloud disk and is retained.
	//
	// Default value: true.
	//
	// Take note of the following items:
	//
	// - This parameter cannot be set for cloud disks that have the multi-attach feature enabled.
	//
	// - If you unmount a data cloud disk, the default value is `false`.
	//
	// - If the disk to be unmounted is an elastic ephemeral disk, you must set the `DeleteWithInstance` parameter to `true`.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the disk to be unmounted.
	//
	// - The disk must be mounted to an instance and be in the In Use (`In_use`) state.
	//
	// - When you unmount a data cloud disk, the instance to which the disk is mounted must be in the Running (`Running`) or Stopped (`Stopped`) state.
	//
	// - When you unmount a system cloud disk, the instance to which the disk is mounted must be in the Stopped (`Stopped`) state.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The instance ID of the ECS instance to which the disk to be uninstalled is attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDiskRequest) GoString() string {
	return s.String()
}

func (s *DetachDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DetachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DetachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachDiskRequest) SetDeleteWithInstance(v bool) *DetachDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *DetachDiskRequest) SetDiskId(v string) *DetachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DetachDiskRequest) SetInstanceId(v string) *DetachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachDiskRequest) SetOwnerAccount(v string) *DetachDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachDiskRequest) SetOwnerId(v int64) *DetachDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachDiskRequest) SetResourceOwnerAccount(v string) *DetachDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachDiskRequest) SetResourceOwnerId(v int64) *DetachDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachDiskRequest) Validate() error {
	return dara.Validate(s)
}

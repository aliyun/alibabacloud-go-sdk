// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootable(v bool) *AttachDiskRequest
	GetBootable() *bool
	SetDeleteWithInstance(v bool) *AttachDiskRequest
	GetDeleteWithInstance() *bool
	SetDevice(v string) *AttachDiskRequest
	GetDevice() *string
	SetDiskId(v string) *AttachDiskRequest
	GetDiskId() *string
	SetForce(v bool) *AttachDiskRequest
	GetForce() *bool
	SetInstanceId(v string) *AttachDiskRequest
	GetInstanceId() *string
	SetKeyPairName(v string) *AttachDiskRequest
	GetKeyPairName() *string
	SetOwnerAccount(v string) *AttachDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachDiskRequest
	GetOwnerId() *int64
	SetPassword(v string) *AttachDiskRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *AttachDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachDiskRequest
	GetResourceOwnerId() *int64
}

type AttachDiskRequest struct {
	// Specifies whether to attach the disk as a system disk. Valid values:
	//
	// - true: Attach as a system disk.
	//
	// - false: Do not attach as a system disk.
	//
	// Default value: false.
	//
	// > If `Bootable` is set to `true`, the target ECS instance must have no system disk attached.
	//
	// example:
	//
	// false
	Bootable *bool `json:"Bootable,omitempty" xml:"Bootable,omitempty"`
	// Specifies whether to release the disk when the instance is released. Valid values:
	//
	// - true: The disk is released with the instance.
	//
	// - false: The disk is not released with the instance. The disk is retained as a pay-as-you-go data disk.
	//
	// Default value: false.
	//
	// Note the following when setting this parameter:
	//
	// - If `DeleteWithInstance` is set to `false` and the ECS instance is under security control (that is, `OperationLocks` contains `"LockReason" : "security"`), this attribute is ignored when the ECS instance is released, and the disk is released along with the instance.
	//
	// - If the disk to attach is an elastic ephemeral disk, you must set `DeleteWithInstance` to `true`.
	//
	// - This parameter is not supported for disks with the multi-attach feature enabled.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The device name of the disk.
	//
	// > This parameter is being deprecated. To improve compatibility, use other parameters to identify the disk.
	//
	// example:
	//
	// testDeviceName
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the disk to attach. The disk (`DiskId`) and the instance (`InstanceId`) must be in the same zone.
	//
	// > Both data disks and system disks are supported. For the relevant constraints, see the operation description above.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1j4l5axzdy6ftk****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether the request is a forced attach request. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	//
	// > Currently, only the ESSD regional disk type (cloud_regional_disk_auto) supports setting this field to true.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the ECS instance to which you want to attach the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1dq5lozx5f4pmd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the SSH key pair to bind to a Linux ECS instance when attaching a system disk.
	//
	// - Windows Server instances: SSH key pairs are not supported. Even if this parameter is specified, only the `Password` configuration takes effect.
	//
	// - Linux instances: Password-based logon is disabled after the key pair is bound.
	//
	// example:
	//
	// KeyPairTestName
	KeyPairName  *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password for the instance when attaching a system disk. This parameter applies only to the administrator and root usernames. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// ```
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// ```
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the `Password` parameter, use HTTPS to send the request to prevent password leakage.
	//
	// example:
	//
	// EcsV587!
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AttachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskRequest) GoString() string {
	return s.String()
}

func (s *AttachDiskRequest) GetBootable() *bool {
	return s.Bootable
}

func (s *AttachDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AttachDiskRequest) GetDevice() *string {
	return s.Device
}

func (s *AttachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *AttachDiskRequest) GetForce() *bool {
	return s.Force
}

func (s *AttachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachDiskRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachDiskRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachDiskRequest) SetBootable(v bool) *AttachDiskRequest {
	s.Bootable = &v
	return s
}

func (s *AttachDiskRequest) SetDeleteWithInstance(v bool) *AttachDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *AttachDiskRequest) SetDevice(v string) *AttachDiskRequest {
	s.Device = &v
	return s
}

func (s *AttachDiskRequest) SetDiskId(v string) *AttachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *AttachDiskRequest) SetForce(v bool) *AttachDiskRequest {
	s.Force = &v
	return s
}

func (s *AttachDiskRequest) SetInstanceId(v string) *AttachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachDiskRequest) SetKeyPairName(v string) *AttachDiskRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachDiskRequest) SetOwnerAccount(v string) *AttachDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachDiskRequest) SetOwnerId(v int64) *AttachDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachDiskRequest) SetPassword(v string) *AttachDiskRequest {
	s.Password = &v
	return s
}

func (s *AttachDiskRequest) SetResourceOwnerAccount(v string) *AttachDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachDiskRequest) SetResourceOwnerId(v int64) *AttachDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachDiskRequest) Validate() error {
	return dara.Validate(s)
}

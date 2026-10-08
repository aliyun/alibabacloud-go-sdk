// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v bool) *AttachRCDiskRequest
	GetDeleteWithInstance() *bool
	SetDiskId(v string) *AttachRCDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *AttachRCDiskRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachRCDiskRequest
	GetRegionId() *string
}

type AttachRCDiskRequest struct {
	// Specifies whether the cloud disk is released when the instance is released. Valid values:
	//
	// true: The cloud disk is released when the instance is released.
	//
	// false: The cloud disk is not released when the instance is released. The cloud disk is retained as a pay-as-you-go data cloud disk.
	//
	// Default value: false.
	//
	// When you configure this parameter, take note of the following items:
	//
	// If you set DeleteWithInstance to false and the instance is locked for security reasons, meaning that OperationLocks contains "LockReason" : "security", this parameter is ignored and the cloud disk is released along with the instance.
	//
	// If the cloud disk to be attached is an elastic ephemeral disk, you must set DeleteWithInstance to true.
	//
	// This parameter is not supported for cloud disks that have the multi-attach feature enabled.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the cloud disk to be attached. The cloud disk (DiskId) and the instance (InstanceId) must be in the same zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz98hnpj2sjo85zc7t2w
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the destination RDS Custom instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachRCDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachRCDiskRequest) GoString() string {
	return s.String()
}

func (s *AttachRCDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AttachRCDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *AttachRCDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachRCDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachRCDiskRequest) SetDeleteWithInstance(v bool) *AttachRCDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *AttachRCDiskRequest) SetDiskId(v string) *AttachRCDiskRequest {
	s.DiskId = &v
	return s
}

func (s *AttachRCDiskRequest) SetInstanceId(v string) *AttachRCDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachRCDiskRequest) SetRegionId(v string) *AttachRCDiskRequest {
	s.RegionId = &v
	return s
}

func (s *AttachRCDiskRequest) Validate() error {
	return dara.Validate(s)
}

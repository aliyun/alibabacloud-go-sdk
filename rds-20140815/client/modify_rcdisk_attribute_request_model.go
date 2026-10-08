// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBurstingEnabled(v bool) *ModifyRCDiskAttributeRequest
	GetBurstingEnabled() *bool
	SetDeleteWithInstance(v bool) *ModifyRCDiskAttributeRequest
	GetDeleteWithInstance() *bool
	SetDescription(v string) *ModifyRCDiskAttributeRequest
	GetDescription() *string
	SetDiskId(v string) *ModifyRCDiskAttributeRequest
	GetDiskId() *string
	SetDiskName(v string) *ModifyRCDiskAttributeRequest
	GetDiskName() *string
	SetRegionId(v string) *ModifyRCDiskAttributeRequest
	GetRegionId() *string
}

type ModifyRCDiskAttributeRequest struct {
	// Specifies whether to enable the performance burst feature for cloud disks that support burst. Valid values:
	//
	// true: Enabled.
	//
	// false: Disabled.
	//
	// Note
	//
	// An error is returned if you pass any value for cloud disks that do not support the burst feature.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Specifies whether to release the cloud disk when the associated instance is released. Default value: null, which indicates that the current value is not changed.
	//
	// Cloud disks that have the multi-attach feature enabled do not support this parameter.
	//
	// An error is returned if you set DeleteWithInstance to false in the following cases:
	//
	// The category of the cloud disk is local disk (ephemeral).
	//
	// The category of the cloud disk is basic cloud disk (cloud) and the cloud disk is not detachable (Portable=false).
	//
	// Warning
	//
	// If you set DeleteWithInstance to false and the ECS instance to which the cloud disk is attached is security-locked with "LockReason" : "security" in OperationLocks, the DeleteWithInstance attribute of the cloud disk is ignored and the cloud disk is released together with the instance.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the cloud disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud disk whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the cloud disk. The name must be 2 to 128 characters in length and can contain Unicode characters under the letter category (including letters from various languages, Chinese characters, and digits). The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// testDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The region ID. You can call DescribeRegions to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCDiskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskAttributeRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyRCDiskAttributeRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyRCDiskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRCDiskAttributeRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyRCDiskAttributeRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyRCDiskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCDiskAttributeRequest) SetBurstingEnabled(v bool) *ModifyRCDiskAttributeRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDeleteWithInstance(v bool) *ModifyRCDiskAttributeRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDescription(v string) *ModifyRCDiskAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDiskId(v string) *ModifyRCDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetDiskName(v string) *ModifyRCDiskAttributeRequest {
	s.DiskName = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) SetRegionId(v string) *ModifyRCDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCDiskAttributeRequest) Validate() error {
	return dara.Validate(s)
}

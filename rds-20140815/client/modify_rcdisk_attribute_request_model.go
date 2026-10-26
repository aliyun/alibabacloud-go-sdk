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
	// Specifies whether to enable the performance burst feature for disks that support this feature. Valid values:
	//
	// true: enables the feature.
	//
	// false: disables the feature.
	//
	// Note
	//
	// An error is returned if you specify any value for a disk that does not support the performance burst feature.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Specifies whether to release the disk when the instance is released. Default value: null. A value of null indicates that the current value remains unchanged.
	//
	// You cannot configure this parameter for cloud disks that have the multi-attach feature enabled.
	//
	// An error is returned if you set the DeleteWithInstance parameter to false in the following scenarios:
	//
	// The disk category is local disk (ephemeral).
	//
	// The disk category is basic cloud disk (cloud) and the disk is not detachable (Portable=false).
	//
	// Warning
	//
	// If you set this parameter to false and the ECS instance to which the disk is attached is locked for security reasons (the OperationLocks parameter indicates "LockReason" : "security"), the disk is released along with the instance regardless of the DeleteWithInstance setting when the instance is released.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the disk whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The cloud disk name. The name must be 2 to 128 characters in length and can contain characters in the Letter category of Unicode, such as English letters, Chinese characters, and digits. It can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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

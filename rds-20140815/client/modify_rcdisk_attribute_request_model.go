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
	// Specifies whether to enable the burst (performance burst) feature for disks that support it. Valid values:
	//
	// true: Enabled.
	//
	// false: No.
	//
	// Note
	//
	// If you specify any value for a disk that does not support the burst feature, an error is returned.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Specifies whether the disk is released when its associated instance is released. Default value: none, which means the current setting remains unchanged.
	//
	// This parameter cannot be set for disks with the multi-attach feature enabled.
	//
	// An error occurs if you set DeleteWithInstance to false in either of the following cases:
	//
	// - The disk category is local disk (ephemeral).
	//
	// - The disk category is basic disk (cloud) and the disk is non-portable (Portable=false).
	//
	// Warning
	//
	// If you set DeleteWithInstance to false, but the ECS instance to which the disk is attached is security locked (indicated by "LockReason": "security" in OperationLocks), the disk will be released together with the instance regardless of the DeleteWithInstance setting when the instance is released.
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
	// The ID of the disk whose property you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Disk name. The name must be 2 to 128 characters in length and can contain characters from the letter categorization in Unicode (including English letters, Chinese characters, and digits). It can also include colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// testDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The region ID. You can call DescribeRegions to obtain valid region IDs.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSystemDisk interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *SystemDisk
	GetCategory() *string
	SetCount(v int32) *SystemDisk
	GetCount() *int32
	SetPerformanceLevel(v string) *SystemDisk
	GetPerformanceLevel() *string
	SetSize(v int32) *SystemDisk
	GetSize() *int32
}

type SystemDisk struct {
	// The type of the system disk. Valid values:
	//
	// - `cloud_efficiency`: Ultra Disk.
	//
	// - `cloud_ssd`: SSD Cloud Disk.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud`: Basic Disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies the number of system disks on each node. Default value: 1.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The performance level of the ESSD. This parameter is valid only when `Category` is set to `cloud_essd`. Valid values:
	//
	// - `PL0`: Up to 10,000 random read/write IOPS per disk.
	//
	// - `PL1` (default): Up to 50,000 random read/write IOPS per disk.
	//
	// - `PL2`: Up to 100,000 random read/write IOPS per disk.
	//
	// - `PL3`: Up to 1,000,000 random read/write IOPS per disk.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk, in GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s SystemDisk) String() string {
	return dara.Prettify(s)
}

func (s SystemDisk) GoString() string {
	return s.String()
}

func (s *SystemDisk) GetCategory() *string {
	return s.Category
}

func (s *SystemDisk) GetCount() *int32 {
	return s.Count
}

func (s *SystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *SystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *SystemDisk) SetCategory(v string) *SystemDisk {
	s.Category = &v
	return s
}

func (s *SystemDisk) SetCount(v int32) *SystemDisk {
	s.Count = &v
	return s
}

func (s *SystemDisk) SetPerformanceLevel(v string) *SystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *SystemDisk) SetSize(v int32) *SystemDisk {
	s.Size = &v
	return s
}

func (s *SystemDisk) Validate() error {
	return dara.Validate(s)
}

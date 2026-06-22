// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataDisk interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DataDisk
	GetCategory() *string
	SetCount(v int32) *DataDisk
	GetCount() *int32
	SetPerformanceLevel(v string) *DataDisk
	GetPerformanceLevel() *string
	SetSize(v int32) *DataDisk
	GetSize() *int32
}

type DataDisk struct {
	// The disk type. Valid values:
	//
	// - `cloud_efficiency`: Ultra Disk.
	//
	// - `cloud_ssd`: Standard SSD.
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
	// The number of data disks.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The performance level of an ESSD. This parameter applies only when the `Category` parameter is set to `cloud_essd`. Valid values:
	//
	// - PL0: A maximum of 10,000 random read/write IOPS per disk.
	//
	// - PL1: A maximum of 50,000 random read/write IOPS per disk.
	//
	// - PL2: A maximum of 100,000 random read/write IOPS per disk.
	//
	// - PL3: A maximum of 1,000,000 random read/write IOPS per disk.
	//
	// Default value: PL1.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of each data disk, in GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DataDisk) String() string {
	return dara.Prettify(s)
}

func (s DataDisk) GoString() string {
	return s.String()
}

func (s *DataDisk) GetCategory() *string {
	return s.Category
}

func (s *DataDisk) GetCount() *int32 {
	return s.Count
}

func (s *DataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DataDisk) SetCategory(v string) *DataDisk {
	s.Category = &v
	return s
}

func (s *DataDisk) SetCount(v int32) *DataDisk {
	s.Count = &v
	return s
}

func (s *DataDisk) SetPerformanceLevel(v string) *DataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DataDisk) SetSize(v int32) *DataDisk {
	s.Size = &v
	return s
}

func (s *DataDisk) Validate() error {
	return dara.Validate(s)
}

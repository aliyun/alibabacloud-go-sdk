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
	SetPerformanceLevel(v string) *SystemDisk
	GetPerformanceLevel() *string
	SetSize(v int64) *SystemDisk
	GetSize() *int64
}

type SystemDisk struct {
	// The system cloud disk type, such as cloud_essd.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD cloud disk, such as PL0 and PL1.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system cloud disk, in GiB.
	//
	// example:
	//
	// 200
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
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

func (s *SystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *SystemDisk) GetSize() *int64 {
	return s.Size
}

func (s *SystemDisk) SetCategory(v string) *SystemDisk {
	s.Category = &v
	return s
}

func (s *SystemDisk) SetPerformanceLevel(v string) *SystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *SystemDisk) SetSize(v int64) *SystemDisk {
	s.Size = &v
	return s
}

func (s *SystemDisk) Validate() error {
	return dara.Validate(s)
}

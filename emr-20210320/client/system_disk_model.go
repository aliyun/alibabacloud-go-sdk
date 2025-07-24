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
	// 磁盘类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 每个节点系统盘数量，默认值为1。
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 创建ESSD云盘作为系统盘使用时，设置云盘的性能等级。取值范围：
	//
	// - PL0：单盘最高随机读写IOPS 1万。
	//
	// - PL1（默认）：单盘最高随机读写IOPS 5万。
	//
	// - PL2：单盘最高随机读写IOPS 10万。
	//
	// - PL3：单盘最高随机读写IOPS 100万。
	//
	// 默认值：PL1。
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// 单位GB。
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

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
	// 磁盘类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 每个节点磁盘数量。
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 创建ESSD云盘作为数据盘使用时，设置云盘的性能等级。取值范围：
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisk interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *Disk
	GetCategory() *string
	SetCount(v int32) *Disk
	GetCount() *int32
	SetPerformanceLevel(v string) *Disk
	GetPerformanceLevel() *string
	SetSize(v int32) *Disk
	GetSize() *int32
}

type Disk struct {
	// 磁盘类型。
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
	// 性能级别。
	//
	// example:
	//
	// S0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// 单位GB。
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s Disk) String() string {
	return dara.Prettify(s)
}

func (s Disk) GoString() string {
	return s.String()
}

func (s *Disk) GetCategory() *string {
	return s.Category
}

func (s *Disk) GetCount() *int32 {
	return s.Count
}

func (s *Disk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *Disk) GetSize() *int32 {
	return s.Size
}

func (s *Disk) SetCategory(v string) *Disk {
	s.Category = &v
	return s
}

func (s *Disk) SetCount(v int32) *Disk {
	s.Count = &v
	return s
}

func (s *Disk) SetPerformanceLevel(v string) *Disk {
	s.PerformanceLevel = &v
	return s
}

func (s *Disk) SetSize(v int32) *Disk {
	s.Size = &v
	return s
}

func (s *Disk) Validate() error {
	return dara.Validate(s)
}

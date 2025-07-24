// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiskInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DiskInfo
	GetCategory() *string
	SetCount(v int32) *DiskInfo
	GetCount() *int32
	SetPerformanceLevel(v string) *DiskInfo
	GetPerformanceLevel() *string
	SetSize(v int32) *DiskInfo
	GetSize() *int32
}

type DiskInfo struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DiskInfo) String() string {
	return dara.Prettify(s)
}

func (s DiskInfo) GoString() string {
	return s.String()
}

func (s *DiskInfo) GetCategory() *string {
	return s.Category
}

func (s *DiskInfo) GetCount() *int32 {
	return s.Count
}

func (s *DiskInfo) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DiskInfo) GetSize() *int32 {
	return s.Size
}

func (s *DiskInfo) SetCategory(v string) *DiskInfo {
	s.Category = &v
	return s
}

func (s *DiskInfo) SetCount(v int32) *DiskInfo {
	s.Count = &v
	return s
}

func (s *DiskInfo) SetPerformanceLevel(v string) *DiskInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *DiskInfo) SetSize(v int32) *DiskInfo {
	s.Size = &v
	return s
}

func (s *DiskInfo) Validate() error {
	return dara.Validate(s)
}

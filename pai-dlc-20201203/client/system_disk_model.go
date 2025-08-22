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
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSystemDiskParam interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *SystemDiskParam
	GetCategory() *string
	SetPerformanceLevel(v string) *SystemDiskParam
	GetPerformanceLevel() *string
	SetSize(v int32) *SystemDiskParam
	GetSize() *int32
}

type SystemDiskParam struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s SystemDiskParam) String() string {
	return dara.Prettify(s)
}

func (s SystemDiskParam) GoString() string {
	return s.String()
}

func (s *SystemDiskParam) GetCategory() *string {
	return s.Category
}

func (s *SystemDiskParam) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *SystemDiskParam) GetSize() *int32 {
	return s.Size
}

func (s *SystemDiskParam) SetCategory(v string) *SystemDiskParam {
	s.Category = &v
	return s
}

func (s *SystemDiskParam) SetPerformanceLevel(v string) *SystemDiskParam {
	s.PerformanceLevel = &v
	return s
}

func (s *SystemDiskParam) SetSize(v int32) *SystemDiskParam {
	s.Size = &v
	return s
}

func (s *SystemDiskParam) Validate() error {
	return dara.Validate(s)
}

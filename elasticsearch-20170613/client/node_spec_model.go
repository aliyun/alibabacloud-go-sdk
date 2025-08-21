// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeSpec interface {
	dara.Model
	String() string
	GoString() string
	SetDisk(v int32) *NodeSpec
	GetDisk() *int32
	SetDiskEncryption(v bool) *NodeSpec
	GetDiskEncryption() *bool
	SetDiskPreference(v string) *NodeSpec
	GetDiskPreference() *string
	SetDiskType(v string) *NodeSpec
	GetDiskType() *string
	SetPerformanceLevel(v string) *NodeSpec
	GetPerformanceLevel() *string
	SetSpec(v string) *NodeSpec
	GetSpec() *string
}

type NodeSpec struct {
	Disk             *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption   *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskPreference   *string `json:"diskPreference,omitempty" xml:"diskPreference,omitempty"`
	DiskType         *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// This parameter is required.
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s NodeSpec) String() string {
	return dara.Prettify(s)
}

func (s NodeSpec) GoString() string {
	return s.String()
}

func (s *NodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *NodeSpec) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *NodeSpec) GetDiskPreference() *string {
	return s.DiskPreference
}

func (s *NodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *NodeSpec) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *NodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *NodeSpec) SetDisk(v int32) *NodeSpec {
	s.Disk = &v
	return s
}

func (s *NodeSpec) SetDiskEncryption(v bool) *NodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *NodeSpec) SetDiskPreference(v string) *NodeSpec {
	s.DiskPreference = &v
	return s
}

func (s *NodeSpec) SetDiskType(v string) *NodeSpec {
	s.DiskType = &v
	return s
}

func (s *NodeSpec) SetPerformanceLevel(v string) *NodeSpec {
	s.PerformanceLevel = &v
	return s
}

func (s *NodeSpec) SetSpec(v string) *NodeSpec {
	s.Spec = &v
	return s
}

func (s *NodeSpec) Validate() error {
	return dara.Validate(s)
}

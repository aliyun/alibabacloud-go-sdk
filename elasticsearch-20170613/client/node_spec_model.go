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
	// Storage space size of data nodes, in GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// Whether to enable cloud disk encryption for data nodes:
	//
	// - true: Enabled
	//
	// - false: Disabled
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// Storage preference.
	DiskPreference *string `json:"diskPreference,omitempty" xml:"diskPreference,omitempty"`
	// Storage type of data nodes. Supported values:
	//
	// - cloud_ssd: SSD cloud disk
	//
	// - cloud_essd: ESSD cloud disk
	//
	// - cloud_efficiency: Ultra cloud disk
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// Performance level of ESSD cloud disks. Required when the disk type of data nodes is ESSD cloud disk. Supported values: PL1, PL2, PL3.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// Data node specification. Specification details can be viewed in [Product Specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// elasticsearch.sn2ne.large
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *ResourceSpec
	GetCu() *int32
	SetDiskNumber(v int32) *ResourceSpec
	GetDiskNumber() *int32
	SetLocalStorageInstanceType(v string) *ResourceSpec
	GetLocalStorageInstanceType() *string
	SetNodeNumber(v int32) *ResourceSpec
	GetNodeNumber() *int32
	SetSpecType(v string) *ResourceSpec
	GetSpecType() *string
	SetStoragePerformanceLevel(v string) *ResourceSpec
	GetStoragePerformanceLevel() *string
	SetStorageSize(v int32) *ResourceSpec
	GetStorageSize() *int32
}

type ResourceSpec struct {
	// The number of CUs. A compute unit (CU) is the basic metering unit of a service. 1 CU = 1 CPU core + 4 GiB of memory.
	//
	// example:
	//
	// 3
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of disk blocks.
	//
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// Local SSD Instance Specification for the node group. This parameter is applicable only when the node group is based on ECS instances and the SpecType is set to \\"Local SSD / Large-capacity Storage\\"
	//
	// example:
	//
	// local_ssd_4_4xlarge
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	NodeNumber *int32 `json:"nodeNumber,omitempty" xml:"nodeNumber,omitempty"`
	// The type of the node group. The following types are included:
	//
	// 	- standard, Standard Edition, ECS + cloud disk.
	//
	// 	- localSSD , local SSD.
	//
	// 	- bigData, which stores large specifications.
	//
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// The performance level of the disks. Valid values:
	//
	// 	- PL0: A single disk can achieve up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single disk can achieve up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single disk can achieve up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single disk can achieve up to 1 million random read/write IOPS.
	//
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s ResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) GetCu() *int32 {
	return s.Cu
}

func (s *ResourceSpec) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *ResourceSpec) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *ResourceSpec) GetNodeNumber() *int32 {
	return s.NodeNumber
}

func (s *ResourceSpec) GetSpecType() *string {
	return s.SpecType
}

func (s *ResourceSpec) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *ResourceSpec) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ResourceSpec) SetCu(v int32) *ResourceSpec {
	s.Cu = &v
	return s
}

func (s *ResourceSpec) SetDiskNumber(v int32) *ResourceSpec {
	s.DiskNumber = &v
	return s
}

func (s *ResourceSpec) SetLocalStorageInstanceType(v string) *ResourceSpec {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *ResourceSpec) SetNodeNumber(v int32) *ResourceSpec {
	s.NodeNumber = &v
	return s
}

func (s *ResourceSpec) SetSpecType(v string) *ResourceSpec {
	s.SpecType = &v
	return s
}

func (s *ResourceSpec) SetStoragePerformanceLevel(v string) *ResourceSpec {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *ResourceSpec) SetStorageSize(v int32) *ResourceSpec {
	s.StorageSize = &v
	return s
}

func (s *ResourceSpec) Validate() error {
	return dara.Validate(s)
}

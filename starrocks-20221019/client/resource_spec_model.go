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
	Cu                       *int32  `json:"cu,omitempty" xml:"cu,omitempty"`
	DiskNumber               *int32  `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	NodeNumber               *int32  `json:"nodeNumber,omitempty" xml:"nodeNumber,omitempty"`
	SpecType                 *string `json:"specType,omitempty" xml:"specType,omitempty"`
	StoragePerformanceLevel  *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	StorageSize              *int32  `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
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

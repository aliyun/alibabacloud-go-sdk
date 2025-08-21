// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultSpecInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCount(v string) *ResultSpecInfoMapValue
	GetCpuCount() *string
	SetMemorySize(v string) *ResultSpecInfoMapValue
	GetMemorySize() *string
	SetEnable(v string) *ResultSpecInfoMapValue
	GetEnable() *string
	SetSpec(v string) *ResultSpecInfoMapValue
	GetSpec() *string
	SetSpecGroupType(v string) *ResultSpecInfoMapValue
	GetSpecGroupType() *string
	SetDisk(v string) *ResultSpecInfoMapValue
	GetDisk() *string
	SetDiskType(v string) *ResultSpecInfoMapValue
	GetDiskType() *string
}

type ResultSpecInfoMapValue struct {
	CpuCount      *string `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
	MemorySize    *string `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	Enable        *string `json:"enable,omitempty" xml:"enable,omitempty"`
	Spec          *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecGroupType *string `json:"specGroupType,omitempty" xml:"specGroupType,omitempty"`
	Disk          *string `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType      *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
}

func (s ResultSpecInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ResultSpecInfoMapValue) GoString() string {
	return s.String()
}

func (s *ResultSpecInfoMapValue) GetCpuCount() *string {
	return s.CpuCount
}

func (s *ResultSpecInfoMapValue) GetMemorySize() *string {
	return s.MemorySize
}

func (s *ResultSpecInfoMapValue) GetEnable() *string {
	return s.Enable
}

func (s *ResultSpecInfoMapValue) GetSpec() *string {
	return s.Spec
}

func (s *ResultSpecInfoMapValue) GetSpecGroupType() *string {
	return s.SpecGroupType
}

func (s *ResultSpecInfoMapValue) GetDisk() *string {
	return s.Disk
}

func (s *ResultSpecInfoMapValue) GetDiskType() *string {
	return s.DiskType
}

func (s *ResultSpecInfoMapValue) SetCpuCount(v string) *ResultSpecInfoMapValue {
	s.CpuCount = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetMemorySize(v string) *ResultSpecInfoMapValue {
	s.MemorySize = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetEnable(v string) *ResultSpecInfoMapValue {
	s.Enable = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetSpec(v string) *ResultSpecInfoMapValue {
	s.Spec = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetSpecGroupType(v string) *ResultSpecInfoMapValue {
	s.SpecGroupType = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetDisk(v string) *ResultSpecInfoMapValue {
	s.Disk = &v
	return s
}

func (s *ResultSpecInfoMapValue) SetDiskType(v string) *ResultSpecInfoMapValue {
	s.DiskType = &v
	return s
}

func (s *ResultSpecInfoMapValue) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *IndexNodeConfiguration
	GetAmount() *int32
	SetDisk(v int32) *IndexNodeConfiguration
	GetDisk() *int32
	SetDiskType(v string) *IndexNodeConfiguration
	GetDiskType() *string
	SetPerformanceLevel(v string) *IndexNodeConfiguration
	GetPerformanceLevel() *string
	SetSpec(v string) *IndexNodeConfiguration
	GetSpec() *string
}

type IndexNodeConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The disk size.
	//
	// example:
	//
	// 500
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The disk type.
	//
	// example:
	//
	// cloud_essd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The disk performance level (PL) PL0.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// elasticsearch.turbo1.ci.xlarge
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s IndexNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s IndexNodeConfiguration) GoString() string {
	return s.String()
}

func (s *IndexNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *IndexNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *IndexNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *IndexNodeConfiguration) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *IndexNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *IndexNodeConfiguration) SetAmount(v int32) *IndexNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *IndexNodeConfiguration) SetDisk(v int32) *IndexNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *IndexNodeConfiguration) SetDiskType(v string) *IndexNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *IndexNodeConfiguration) SetPerformanceLevel(v string) *IndexNodeConfiguration {
	s.PerformanceLevel = &v
	return s
}

func (s *IndexNodeConfiguration) SetSpec(v string) *IndexNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *IndexNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

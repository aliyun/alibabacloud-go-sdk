// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *SearchNodeConfiguration
	GetAmount() *int32
	SetDisk(v int32) *SearchNodeConfiguration
	GetDisk() *int32
	SetDiskType(v string) *SearchNodeConfiguration
	GetDiskType() *string
	SetPerformanceLevel(v string) *SearchNodeConfiguration
	GetPerformanceLevel() *string
	SetSpec(v string) *SearchNodeConfiguration
	GetSpec() *string
}

type SearchNodeConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The cloud disk size.
	//
	// example:
	//
	// 500
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The cloud disk type.
	//
	// example:
	//
	// cloud_essd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The cloud disk performance level (PL).
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

func (s SearchNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SearchNodeConfiguration) GoString() string {
	return s.String()
}

func (s *SearchNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *SearchNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *SearchNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *SearchNodeConfiguration) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *SearchNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *SearchNodeConfiguration) SetAmount(v int32) *SearchNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *SearchNodeConfiguration) SetDisk(v int32) *SearchNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *SearchNodeConfiguration) SetDiskType(v string) *SearchNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *SearchNodeConfiguration) SetPerformanceLevel(v string) *SearchNodeConfiguration {
	s.PerformanceLevel = &v
	return s
}

func (s *SearchNodeConfiguration) SetSpec(v string) *SearchNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *SearchNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

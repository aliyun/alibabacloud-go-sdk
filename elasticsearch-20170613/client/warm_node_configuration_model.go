// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *WarmNodeConfiguration
	GetAmount() *int64
	SetDisk(v int64) *WarmNodeConfiguration
	GetDisk() *int64
	SetDiskEncryption(v bool) *WarmNodeConfiguration
	GetDiskEncryption() *bool
	SetDiskType(v string) *WarmNodeConfiguration
	GetDiskType() *string
	SetPerformanceLevel(v string) *WarmNodeConfiguration
	GetPerformanceLevel() *string
	SetSpec(v string) *WarmNodeConfiguration
	GetSpec() *string
}

type WarmNodeConfiguration struct {
	Amount           *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk             *int64  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption   *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType         *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// This parameter is required.
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s WarmNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s WarmNodeConfiguration) GoString() string {
	return s.String()
}

func (s *WarmNodeConfiguration) GetAmount() *int64 {
	return s.Amount
}

func (s *WarmNodeConfiguration) GetDisk() *int64 {
	return s.Disk
}

func (s *WarmNodeConfiguration) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *WarmNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *WarmNodeConfiguration) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *WarmNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *WarmNodeConfiguration) SetAmount(v int64) *WarmNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *WarmNodeConfiguration) SetDisk(v int64) *WarmNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *WarmNodeConfiguration) SetDiskEncryption(v bool) *WarmNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *WarmNodeConfiguration) SetDiskType(v string) *WarmNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *WarmNodeConfiguration) SetPerformanceLevel(v string) *WarmNodeConfiguration {
	s.PerformanceLevel = &v
	return s
}

func (s *WarmNodeConfiguration) SetSpec(v string) *WarmNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *WarmNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

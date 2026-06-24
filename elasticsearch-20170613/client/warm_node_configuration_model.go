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
	// The number of cold data nodes.
	//
	// example:
	//
	// 3
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage space size of cold data nodes, in GB.
	//
	// example:
	//
	// 500
	Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// Specifies whether to enable cloud disk encryption for cold data nodes. Valid values:
	//
	// - true: enabled.
	//
	// - false: not enabled.
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The storage type of cold data nodes. Only cloud_efficiency (ultra cloud disk) is supported.
	//
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. This parameter is required when the disk type of cold data nodes is a standard SSD. Valid values: PL1, PL2, and PL3.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// The node specifications of cold data nodes. For more information, see [Product specifications](https://help.aliyun.com/document_detail/271718.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// elasticsearch.sn2ne.large
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

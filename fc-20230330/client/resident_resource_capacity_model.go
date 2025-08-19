// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResidentResourceCapacity interface {
	dara.Model
	String() string
	GoString() string
	SetGpuType(v string) *ResidentResourceCapacity
	GetGpuType() *string
	SetTotalCpuCores(v int64) *ResidentResourceCapacity
	GetTotalCpuCores() *int64
	SetTotalDiskSize(v int64) *ResidentResourceCapacity
	GetTotalDiskSize() *int64
	SetTotalGpuCards(v int64) *ResidentResourceCapacity
	GetTotalGpuCards() *int64
	SetTotalGpuMemorySize(v int64) *ResidentResourceCapacity
	GetTotalGpuMemorySize() *int64
	SetTotalMemorySize(v int64) *ResidentResourceCapacity
	GetTotalMemorySize() *int64
}

type ResidentResourceCapacity struct {
	// GPU 卡型
	GpuType *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
	// CPU 总核数
	TotalCpuCores *int64 `json:"totalCpuCores,omitempty" xml:"totalCpuCores,omitempty"`
	// 总磁盘大小，单位 GB
	TotalDiskSize *int64 `json:"totalDiskSize,omitempty" xml:"totalDiskSize,omitempty"`
	// GPU总卡数
	TotalGpuCards *int64 `json:"totalGpuCards,omitempty" xml:"totalGpuCards,omitempty"`
	// 总显存大小，单位 GB
	TotalGpuMemorySize *int64 `json:"totalGpuMemorySize,omitempty" xml:"totalGpuMemorySize,omitempty"`
	// 总内存大小，单位 GB
	TotalMemorySize *int64 `json:"totalMemorySize,omitempty" xml:"totalMemorySize,omitempty"`
}

func (s ResidentResourceCapacity) String() string {
	return dara.Prettify(s)
}

func (s ResidentResourceCapacity) GoString() string {
	return s.String()
}

func (s *ResidentResourceCapacity) GetGpuType() *string {
	return s.GpuType
}

func (s *ResidentResourceCapacity) GetTotalCpuCores() *int64 {
	return s.TotalCpuCores
}

func (s *ResidentResourceCapacity) GetTotalDiskSize() *int64 {
	return s.TotalDiskSize
}

func (s *ResidentResourceCapacity) GetTotalGpuCards() *int64 {
	return s.TotalGpuCards
}

func (s *ResidentResourceCapacity) GetTotalGpuMemorySize() *int64 {
	return s.TotalGpuMemorySize
}

func (s *ResidentResourceCapacity) GetTotalMemorySize() *int64 {
	return s.TotalMemorySize
}

func (s *ResidentResourceCapacity) SetGpuType(v string) *ResidentResourceCapacity {
	s.GpuType = &v
	return s
}

func (s *ResidentResourceCapacity) SetTotalCpuCores(v int64) *ResidentResourceCapacity {
	s.TotalCpuCores = &v
	return s
}

func (s *ResidentResourceCapacity) SetTotalDiskSize(v int64) *ResidentResourceCapacity {
	s.TotalDiskSize = &v
	return s
}

func (s *ResidentResourceCapacity) SetTotalGpuCards(v int64) *ResidentResourceCapacity {
	s.TotalGpuCards = &v
	return s
}

func (s *ResidentResourceCapacity) SetTotalGpuMemorySize(v int64) *ResidentResourceCapacity {
	s.TotalGpuMemorySize = &v
	return s
}

func (s *ResidentResourceCapacity) SetTotalMemorySize(v int64) *ResidentResourceCapacity {
	s.TotalMemorySize = &v
	return s
}

func (s *ResidentResourceCapacity) Validate() error {
	return dara.Validate(s)
}

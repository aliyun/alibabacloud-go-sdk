// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsResources interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *StatisticsResources
	GetCPU() *string
	SetGPU(v string) *StatisticsResources
	GetGPU() *string
	SetHyperNodeNum(v int64) *StatisticsResources
	GetHyperNodeNum() *int64
	SetMemory(v string) *StatisticsResources
	GetMemory() *string
	SetNodeNum(v int64) *StatisticsResources
	GetNodeNum() *int64
}

type StatisticsResources struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	HyperNodeNum *int64  `json:"HyperNodeNum,omitempty" xml:"HyperNodeNum,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeNum      *int64  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s StatisticsResources) String() string {
	return dara.Prettify(s)
}

func (s StatisticsResources) GoString() string {
	return s.String()
}

func (s *StatisticsResources) GetCPU() *string {
	return s.CPU
}

func (s *StatisticsResources) GetGPU() *string {
	return s.GPU
}

func (s *StatisticsResources) GetHyperNodeNum() *int64 {
	return s.HyperNodeNum
}

func (s *StatisticsResources) GetMemory() *string {
	return s.Memory
}

func (s *StatisticsResources) GetNodeNum() *int64 {
	return s.NodeNum
}

func (s *StatisticsResources) SetCPU(v string) *StatisticsResources {
	s.CPU = &v
	return s
}

func (s *StatisticsResources) SetGPU(v string) *StatisticsResources {
	s.GPU = &v
	return s
}

func (s *StatisticsResources) SetHyperNodeNum(v int64) *StatisticsResources {
	s.HyperNodeNum = &v
	return s
}

func (s *StatisticsResources) SetMemory(v string) *StatisticsResources {
	s.Memory = &v
	return s
}

func (s *StatisticsResources) SetNodeNum(v int64) *StatisticsResources {
	s.NodeNum = &v
	return s
}

func (s *StatisticsResources) Validate() error {
	return dara.Validate(s)
}

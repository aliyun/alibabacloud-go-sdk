// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobMetric interface {
	dara.Model
	String() string
	GoString() string
	SetTotalCpu(v float64) *JobMetric
	GetTotalCpu() *float64
	SetTotalMemoryByte(v int64) *JobMetric
	GetTotalMemoryByte() *int64
}

type JobMetric struct {
	// example:
	//
	// 2
	TotalCpu *float64 `json:"totalCpu,omitempty" xml:"totalCpu,omitempty"`
	// example:
	//
	// 4096
	TotalMemoryByte *int64 `json:"totalMemoryByte,omitempty" xml:"totalMemoryByte,omitempty"`
}

func (s JobMetric) String() string {
	return dara.Prettify(s)
}

func (s JobMetric) GoString() string {
	return s.String()
}

func (s *JobMetric) GetTotalCpu() *float64 {
	return s.TotalCpu
}

func (s *JobMetric) GetTotalMemoryByte() *int64 {
	return s.TotalMemoryByte
}

func (s *JobMetric) SetTotalCpu(v float64) *JobMetric {
	s.TotalCpu = &v
	return s
}

func (s *JobMetric) SetTotalMemoryByte(v int64) *JobMetric {
	s.TotalMemoryByte = &v
	return s
}

func (s *JobMetric) Validate() error {
	return dara.Validate(s)
}

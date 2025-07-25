// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGPUMetric interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v int64) *GPUMetric
	GetIndex() *int64
	SetModel(v string) *GPUMetric
	GetModel() *string
	SetStatus(v int64) *GPUMetric
	GetStatus() *int64
	SetUsageRate(v float32) *GPUMetric
	GetUsageRate() *float32
}

type GPUMetric struct {
	Index *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 0：异常；1：正常
	Status    *int64   `json:"Status,omitempty" xml:"Status,omitempty"`
	UsageRate *float32 `json:"UsageRate,omitempty" xml:"UsageRate,omitempty"`
}

func (s GPUMetric) String() string {
	return dara.Prettify(s)
}

func (s GPUMetric) GoString() string {
	return s.String()
}

func (s *GPUMetric) GetIndex() *int64 {
	return s.Index
}

func (s *GPUMetric) GetModel() *string {
	return s.Model
}

func (s *GPUMetric) GetStatus() *int64 {
	return s.Status
}

func (s *GPUMetric) GetUsageRate() *float32 {
	return s.UsageRate
}

func (s *GPUMetric) SetIndex(v int64) *GPUMetric {
	s.Index = &v
	return s
}

func (s *GPUMetric) SetModel(v string) *GPUMetric {
	s.Model = &v
	return s
}

func (s *GPUMetric) SetStatus(v int64) *GPUMetric {
	s.Status = &v
	return s
}

func (s *GPUMetric) SetUsageRate(v float32) *GPUMetric {
	s.UsageRate = &v
	return s
}

func (s *GPUMetric) Validate() error {
	return dara.Validate(s)
}

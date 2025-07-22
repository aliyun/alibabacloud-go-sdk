// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdleInstanceCullerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest
	GetCpuPercentThreshold() *int32
	SetGpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest
	GetGpuPercentThreshold() *int32
	SetMaxIdleTimeInMinutes(v int32) *CreateIdleInstanceCullerRequest
	GetMaxIdleTimeInMinutes() *int32
}

type CreateIdleInstanceCullerRequest struct {
	// The CPU utilization threshold. Unit: percentage. Valid values: 1 to 100. If the CPU utilization of the instance is lower than this threshold, the instance is considered idle.
	//
	// example:
	//
	// 20
	CpuPercentThreshold *int32 `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	// The GPU utilization threshold. Unit: percentage. Valid values: 1 to 100. This parameter takes effect only if the instance is of the GPU instance type. If both CPU and GPU utilization is lower than the thresholds, the instance is considered idle.
	//
	// example:
	//
	// 10
	GpuPercentThreshold *int32 `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	// The maximum time duration for which the instance is idle. Unit: minutes. If the time duration for which the instance is idle exceeds this value, the system automatically stops the instance.
	//
	// example:
	//
	// 60
	MaxIdleTimeInMinutes *int32 `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
}

func (s CreateIdleInstanceCullerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdleInstanceCullerRequest) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerRequest) GetCpuPercentThreshold() *int32 {
	return s.CpuPercentThreshold
}

func (s *CreateIdleInstanceCullerRequest) GetGpuPercentThreshold() *int32 {
	return s.GpuPercentThreshold
}

func (s *CreateIdleInstanceCullerRequest) GetMaxIdleTimeInMinutes() *int32 {
	return s.MaxIdleTimeInMinutes
}

func (s *CreateIdleInstanceCullerRequest) SetCpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest {
	s.CpuPercentThreshold = &v
	return s
}

func (s *CreateIdleInstanceCullerRequest) SetGpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest {
	s.GpuPercentThreshold = &v
	return s
}

func (s *CreateIdleInstanceCullerRequest) SetMaxIdleTimeInMinutes(v int32) *CreateIdleInstanceCullerRequest {
	s.MaxIdleTimeInMinutes = &v
	return s
}

func (s *CreateIdleInstanceCullerRequest) Validate() error {
	return dara.Validate(s)
}

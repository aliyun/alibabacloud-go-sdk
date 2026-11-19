// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuota interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCores(v int32) *Quota
	GetCpuCores() *int32
	SetInstanceCount(v int32) *Quota
	GetInstanceCount() *int32
	SetMemoryGB(v int32) *Quota
	GetMemoryGB() *int32
	SetScalingCpuCoresPerMinute(v int64) *Quota
	GetScalingCpuCoresPerMinute() *int64
	SetTagValue(v string) *Quota
	GetTagValue() *string
}

type Quota struct {
	// The vCPU quota, in cores.
	//
	// example:
	//
	// 100
	CpuCores *int32 `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The quota for the number of active instances.
	//
	// example:
	//
	// 100
	InstanceCount *int32 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	// The memory quota, in GB.
	//
	// example:
	//
	// 200
	MemoryGB *int32 `json:"memoryGB,omitempty" xml:"memoryGB,omitempty"`
	// The team CPU creation rate in cores per minute. To retain the original configuration during an update, omit this parameter or set it to null. A positive integer sets an exclusive rate. A value of 0 prevents adding new CPU capacity. A value of -1 clears the exclusive rate and restores the remaining shared account rate. When queried, this parameter is omitted if it is not configured, an explicit 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 120
	ScalingCpuCoresPerMinute *int64 `json:"scalingCpuCoresPerMinute,omitempty" xml:"scalingCpuCoresPerMinute,omitempty"`
	// The quota tag value, which is set to the team ID. Quotas are independently calculated based on this dimension.
	//
	// example:
	//
	// 7696f4cf-****
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s Quota) String() string {
	return dara.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *Quota) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *Quota) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *Quota) GetScalingCpuCoresPerMinute() *int64 {
	return s.ScalingCpuCoresPerMinute
}

func (s *Quota) GetTagValue() *string {
	return s.TagValue
}

func (s *Quota) SetCpuCores(v int32) *Quota {
	s.CpuCores = &v
	return s
}

func (s *Quota) SetInstanceCount(v int32) *Quota {
	s.InstanceCount = &v
	return s
}

func (s *Quota) SetMemoryGB(v int32) *Quota {
	s.MemoryGB = &v
	return s
}

func (s *Quota) SetScalingCpuCoresPerMinute(v int64) *Quota {
	s.ScalingCpuCoresPerMinute = &v
	return s
}

func (s *Quota) SetTagValue(v string) *Quota {
	s.TagValue = &v
	return s
}

func (s *Quota) Validate() error {
	return dara.Validate(s)
}

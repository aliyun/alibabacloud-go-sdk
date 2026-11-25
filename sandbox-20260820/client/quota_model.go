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
	SetCreateSandboxWithSnapshotConcurrency(v int32) *Quota
	GetCreateSandboxWithSnapshotConcurrency() *int32
	SetCreateSnapshotConcurrency(v int32) *Quota
	GetCreateSnapshotConcurrency() *int32
	SetInstanceCount(v int32) *Quota
	GetInstanceCount() *int32
	SetMemoryGB(v int32) *Quota
	GetMemoryGB() *int32
	SetPauseConcurrency(v int32) *Quota
	GetPauseConcurrency() *int32
	SetResumeConcurrency(v int32) *Quota
	GetResumeConcurrency() *int32
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
	// The team snapshot-based sandbox creation concurrency quota, in requests. This value represents the maximum number of snapshot-based sandbox creation requests that can be executed at the same time. When updating, omit this parameter or set it to null to retain the original configuration. Set it to a positive integer to specify an exclusive quota. Set it to 0 to prohibit snapshot-based sandbox creation. Set it to -1 to clear the exclusive quota and restore the remaining quota of the shared account. In query responses, this parameter is omitted if it is not configured. An explicit value of 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 10
	CreateSandboxWithSnapshotConcurrency *int32 `json:"createSandboxWithSnapshotConcurrency,omitempty" xml:"createSandboxWithSnapshotConcurrency,omitempty"`
	// The team snapshot creation concurrency quota, in requests. This value represents the maximum number of snapshot creation requests that can be executed at the same time. When updating, omit this parameter or set it to null to retain the original configuration. Set it to a positive integer to specify an exclusive quota. Set it to 0 to prohibit snapshot creation. Set it to -1 to clear the exclusive quota and restore the remaining quota of the shared account. In query responses, this parameter is omitted if it is not configured. An explicit value of 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 10
	CreateSnapshotConcurrency *int32 `json:"createSnapshotConcurrency,omitempty" xml:"createSnapshotConcurrency,omitempty"`
	// The quota for active instances.
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
	// The team sandbox hibernation concurrency quota, in requests. This value represents the maximum number of sandbox hibernation requests that can be executed at the same time. When updating, omit this parameter or set it to null to retain the original configuration. Set it to a positive integer to specify an exclusive quota. Set it to 0 to prohibit sandbox hibernation. Set it to -1 to clear the exclusive quota and restore the remaining quota of the shared account. In query responses, this parameter is omitted if it is not configured. An explicit value of 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 10
	PauseConcurrency *int32 `json:"pauseConcurrency,omitempty" xml:"pauseConcurrency,omitempty"`
	// The team sandbox wake-up concurrency quota, in requests. This value represents the maximum number of sandbox wake-up requests that can be executed at the same time. When updating, omit this parameter or set it to null to retain the original configuration. Set it to a positive integer to specify an exclusive quota. Set it to 0 to prohibit sandbox wake-up. Set it to -1 to clear the exclusive quota and restore the remaining quota of the shared account. In query responses, this parameter is omitted if it is not configured. An explicit value of 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 10
	ResumeConcurrency *int32 `json:"resumeConcurrency,omitempty" xml:"resumeConcurrency,omitempty"`
	// The team CPU creation rate, in cores per minute. When updating, omit this parameter or set it to null to retain the original configuration. Set it to a positive integer to specify an exclusive rate. Set it to 0 to prohibit new CPU capacity. Set it to -1 to clear the exclusive rate and restore the remaining rate of the shared account. In query responses, this parameter is omitted if it is not configured. An explicit value of 0 is retained, and -1 is not returned.
	//
	// example:
	//
	// 120
	ScalingCpuCoresPerMinute *int64 `json:"scalingCpuCoresPerMinute,omitempty" xml:"scalingCpuCoresPerMinute,omitempty"`
	// The quota tag value, which represents the team ID. Quotas are independently calculated based on this dimension.
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

func (s *Quota) GetCreateSandboxWithSnapshotConcurrency() *int32 {
	return s.CreateSandboxWithSnapshotConcurrency
}

func (s *Quota) GetCreateSnapshotConcurrency() *int32 {
	return s.CreateSnapshotConcurrency
}

func (s *Quota) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *Quota) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *Quota) GetPauseConcurrency() *int32 {
	return s.PauseConcurrency
}

func (s *Quota) GetResumeConcurrency() *int32 {
	return s.ResumeConcurrency
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

func (s *Quota) SetCreateSandboxWithSnapshotConcurrency(v int32) *Quota {
	s.CreateSandboxWithSnapshotConcurrency = &v
	return s
}

func (s *Quota) SetCreateSnapshotConcurrency(v int32) *Quota {
	s.CreateSnapshotConcurrency = &v
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

func (s *Quota) SetPauseConcurrency(v int32) *Quota {
	s.PauseConcurrency = &v
	return s
}

func (s *Quota) SetResumeConcurrency(v int32) *Quota {
	s.ResumeConcurrency = &v
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

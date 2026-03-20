// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceInstanceWorker interface {
	dara.Model
	String() string
	GoString() string
	SetCpuLimit(v int32) *ResourceInstanceWorker
	GetCpuLimit() *int32
	SetCpuRequest(v int32) *ResourceInstanceWorker
	GetCpuRequest() *int32
	SetGpuLimit(v int32) *ResourceInstanceWorker
	GetGpuLimit() *int32
	SetGpuRequest(v int32) *ResourceInstanceWorker
	GetGpuRequest() *int32
	SetMemoryLimit(v int32) *ResourceInstanceWorker
	GetMemoryLimit() *int32
	SetMemoryRquest(v int32) *ResourceInstanceWorker
	GetMemoryRquest() *int32
	SetName(v string) *ResourceInstanceWorker
	GetName() *string
	SetReady(v bool) *ResourceInstanceWorker
	GetReady() *bool
	SetRestartCount(v int32) *ResourceInstanceWorker
	GetRestartCount() *int32
	SetServiceName(v string) *ResourceInstanceWorker
	GetServiceName() *string
	SetStartTime(v string) *ResourceInstanceWorker
	GetStartTime() *string
	SetStatus(v string) *ResourceInstanceWorker
	GetStatus() *string
}

type ResourceInstanceWorker struct {
	// The maximum number of CPU cores.
	//
	// example:
	//
	// 2
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// The number of CPU cores that you applied for.
	//
	// example:
	//
	// 1
	CpuRequest *int32 `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// The maximum number of GPUs.
	//
	// example:
	//
	// 0
	GpuLimit *int32 `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	// The number of GPUs that you applied for.
	//
	// example:
	//
	// 0
	GpuRequest *int32 `json:"GpuRequest,omitempty" xml:"GpuRequest,omitempty"`
	// The maximum memory size.
	//
	// example:
	//
	// 400
	MemoryLimit *int32 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The memory size that you applied for.
	//
	// example:
	//
	// 200
	MemoryRquest *int32 `json:"MemoryRquest,omitempty" xml:"MemoryRquest,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// test-india-5bbbfdbc98-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the instance worker is ready.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times the instance worker restarted.
	//
	// example:
	//
	// 3
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The service name.
	//
	// example:
	//
	// test-in****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The time when the instance worker started.
	//
	// example:
	//
	// 2021-05-09T09:56:19Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance state.
	//
	// Valid values:
	//
	// 	- Terminating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Succeeded
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Unknown
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Failed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResourceInstanceWorker) String() string {
	return dara.Prettify(s)
}

func (s ResourceInstanceWorker) GoString() string {
	return s.String()
}

func (s *ResourceInstanceWorker) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *ResourceInstanceWorker) GetCpuRequest() *int32 {
	return s.CpuRequest
}

func (s *ResourceInstanceWorker) GetGpuLimit() *int32 {
	return s.GpuLimit
}

func (s *ResourceInstanceWorker) GetGpuRequest() *int32 {
	return s.GpuRequest
}

func (s *ResourceInstanceWorker) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *ResourceInstanceWorker) GetMemoryRquest() *int32 {
	return s.MemoryRquest
}

func (s *ResourceInstanceWorker) GetName() *string {
	return s.Name
}

func (s *ResourceInstanceWorker) GetReady() *bool {
	return s.Ready
}

func (s *ResourceInstanceWorker) GetRestartCount() *int32 {
	return s.RestartCount
}

func (s *ResourceInstanceWorker) GetServiceName() *string {
	return s.ServiceName
}

func (s *ResourceInstanceWorker) GetStartTime() *string {
	return s.StartTime
}

func (s *ResourceInstanceWorker) GetStatus() *string {
	return s.Status
}

func (s *ResourceInstanceWorker) SetCpuLimit(v int32) *ResourceInstanceWorker {
	s.CpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetCpuRequest(v int32) *ResourceInstanceWorker {
	s.CpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuLimit(v int32) *ResourceInstanceWorker {
	s.GpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuRequest(v int32) *ResourceInstanceWorker {
	s.GpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryLimit(v int32) *ResourceInstanceWorker {
	s.MemoryLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryRquest(v int32) *ResourceInstanceWorker {
	s.MemoryRquest = &v
	return s
}

func (s *ResourceInstanceWorker) SetName(v string) *ResourceInstanceWorker {
	s.Name = &v
	return s
}

func (s *ResourceInstanceWorker) SetReady(v bool) *ResourceInstanceWorker {
	s.Ready = &v
	return s
}

func (s *ResourceInstanceWorker) SetRestartCount(v int32) *ResourceInstanceWorker {
	s.RestartCount = &v
	return s
}

func (s *ResourceInstanceWorker) SetServiceName(v string) *ResourceInstanceWorker {
	s.ServiceName = &v
	return s
}

func (s *ResourceInstanceWorker) SetStartTime(v string) *ResourceInstanceWorker {
	s.StartTime = &v
	return s
}

func (s *ResourceInstanceWorker) SetStatus(v string) *ResourceInstanceWorker {
	s.Status = &v
	return s
}

func (s *ResourceInstanceWorker) Validate() error {
	return dara.Validate(s)
}

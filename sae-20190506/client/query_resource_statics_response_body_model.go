// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceStaticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryResourceStaticsResponseBody
	GetCode() *string
	SetData(v *QueryResourceStaticsResponseBodyData) *QueryResourceStaticsResponseBody
	GetData() *QueryResourceStaticsResponseBodyData
	SetErrorCode(v string) *QueryResourceStaticsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryResourceStaticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryResourceStaticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryResourceStaticsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryResourceStaticsResponseBody
	GetTraceId() *string
}

type QueryResourceStaticsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The resource usage.
	Data *QueryResourceStaticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// - The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7CCF7092-72CA-4431-90D6-C7D98752****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the resource usage of an application was obtained. Valid values:
	//
	// 	- **true**: indicates that the resource usage was obtained.
	//
	// 	- **false**: indicates that the resource usage could not be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// ac1a08a015623098794277264e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s QueryResourceStaticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryResourceStaticsResponseBody) GetData() *QueryResourceStaticsResponseBodyData {
	return s.Data
}

func (s *QueryResourceStaticsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryResourceStaticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryResourceStaticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResourceStaticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryResourceStaticsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryResourceStaticsResponseBody) SetCode(v string) *QueryResourceStaticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetData(v *QueryResourceStaticsResponseBodyData) *QueryResourceStaticsResponseBody {
	s.Data = v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetErrorCode(v string) *QueryResourceStaticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetMessage(v string) *QueryResourceStaticsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetRequestId(v string) *QueryResourceStaticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetSuccess(v bool) *QueryResourceStaticsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetTraceId(v string) *QueryResourceStaticsResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourceStaticsResponseBodyData struct {
	// The real-time resource usage.
	RealTimeRes *QueryResourceStaticsResponseBodyDataRealTimeRes `json:"RealTimeRes,omitempty" xml:"RealTimeRes,omitempty" type:"Struct"`
	// The resource usage of the current month.
	Summary *QueryResourceStaticsResponseBodyDataSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s QueryResourceStaticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyData) GetRealTimeRes() *QueryResourceStaticsResponseBodyDataRealTimeRes {
	return s.RealTimeRes
}

func (s *QueryResourceStaticsResponseBodyData) GetSummary() *QueryResourceStaticsResponseBodyDataSummary {
	return s.Summary
}

func (s *QueryResourceStaticsResponseBodyData) SetRealTimeRes(v *QueryResourceStaticsResponseBodyDataRealTimeRes) *QueryResourceStaticsResponseBodyData {
	s.RealTimeRes = v
	return s
}

func (s *QueryResourceStaticsResponseBodyData) SetSummary(v *QueryResourceStaticsResponseBodyDataSummary) *QueryResourceStaticsResponseBodyData {
	s.Summary = v
	return s
}

func (s *QueryResourceStaticsResponseBodyData) Validate() error {
	if s.RealTimeRes != nil {
		if err := s.RealTimeRes.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourceStaticsResponseBodyDataRealTimeRes struct {
	// The CPU usage. Unit: core per minute.
	//
	// example:
	//
	// 13
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The storage size of the temporary storage space. Unit: GiB.
	//
	// example:
	//
	// 0
	EphemeralStorage *float32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The memory usage. Unit: GiB per minute.
	//
	// example:
	//
	// 26
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QueryResourceStaticsResponseBodyDataRealTimeRes) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyDataRealTimeRes) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) GetCpu() *float32 {
	return s.Cpu
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) GetEphemeralStorage() *float32 {
	return s.EphemeralStorage
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) GetMemory() *float32 {
	return s.Memory
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) SetCpu(v float32) *QueryResourceStaticsResponseBodyDataRealTimeRes {
	s.Cpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) SetEphemeralStorage(v float32) *QueryResourceStaticsResponseBodyDataRealTimeRes {
	s.EphemeralStorage = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) SetMemory(v float32) *QueryResourceStaticsResponseBodyDataRealTimeRes {
	s.Memory = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) Validate() error {
	return dara.Validate(s)
}

type QueryResourceStaticsResponseBodyDataSummary struct {
	// The usage of active vCPU. Unit: Core*min.
	//
	// example:
	//
	// 10
	ActiveCpu *float32 `json:"ActiveCpu,omitempty" xml:"ActiveCpu,omitempty"`
	// The CPU usage. Unit: core per minute.
	//
	// example:
	//
	// 3354
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The CU usage.
	//
	// example:
	//
	// 2312145
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The storage size of the temporary storage space. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *float32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// example:
	//
	// c8g1
	GpuA10     *float32 `json:"GpuA10,omitempty" xml:"GpuA10,omitempty"`
	GpuPpu810e *float32 `json:"GpuPpu810e,omitempty" xml:"GpuPpu810e,omitempty"`
	// The usage of idle CPU. Unit: Core*min.
	//
	// example:
	//
	// 10
	IdleCpu *float32 `json:"IdleCpu,omitempty" xml:"IdleCpu,omitempty"`
	// The memory usage. Unit: GiB per minute.
	//
	// example:
	//
	// 6708
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QueryResourceStaticsResponseBodyDataSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetActiveCpu() *float32 {
	return s.ActiveCpu
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetCpu() *float32 {
	return s.Cpu
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetCu() *float32 {
	return s.Cu
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetEphemeralStorage() *float32 {
	return s.EphemeralStorage
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetGpuA10() *float32 {
	return s.GpuA10
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetGpuPpu810e() *float32 {
	return s.GpuPpu810e
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetIdleCpu() *float32 {
	return s.IdleCpu
}

func (s *QueryResourceStaticsResponseBodyDataSummary) GetMemory() *float32 {
	return s.Memory
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetActiveCpu(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.ActiveCpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetCpu(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.Cpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetCu(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.Cu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetEphemeralStorage(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.EphemeralStorage = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetGpuA10(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.GpuA10 = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetGpuPpu810e(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.GpuPpu810e = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetIdleCpu(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.IdleCpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetMemory(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.Memory = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) Validate() error {
	return dara.Validate(s)
}

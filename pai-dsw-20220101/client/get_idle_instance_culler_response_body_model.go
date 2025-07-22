// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdleInstanceCullerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdleInstanceCullerResponseBody
	GetCode() *string
	SetCpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody
	GetCpuPercentThreshold() *int32
	SetGpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody
	GetGpuPercentThreshold() *int32
	SetIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody
	GetIdleTimeInMinutes() *int32
	SetInstanceId(v string) *GetIdleInstanceCullerResponseBody
	GetInstanceId() *string
	SetMaxIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody
	GetMaxIdleTimeInMinutes() *int32
	SetMessage(v string) *GetIdleInstanceCullerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdleInstanceCullerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdleInstanceCullerResponseBody
	GetSuccess() *bool
}

type GetIdleInstanceCullerResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// ValidationError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The time duration for which the instance is idle. Unit: minutes.
	//
	// example:
	//
	// 30
	IdleTimeInMinutes *int32 `json:"IdleTimeInMinutes,omitempty" xml:"IdleTimeInMinutes,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum time duration for which the instance is idle. Unit: minutes. If the time duration for which the instance is idle exceeds this value, the system automatically stops the instance.
	//
	// example:
	//
	// 60
	MaxIdleTimeInMinutes *int32 `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
	// The error message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIdleInstanceCullerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdleInstanceCullerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdleInstanceCullerResponseBody) GetCpuPercentThreshold() *int32 {
	return s.CpuPercentThreshold
}

func (s *GetIdleInstanceCullerResponseBody) GetGpuPercentThreshold() *int32 {
	return s.GpuPercentThreshold
}

func (s *GetIdleInstanceCullerResponseBody) GetIdleTimeInMinutes() *int32 {
	return s.IdleTimeInMinutes
}

func (s *GetIdleInstanceCullerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdleInstanceCullerResponseBody) GetMaxIdleTimeInMinutes() *int32 {
	return s.MaxIdleTimeInMinutes
}

func (s *GetIdleInstanceCullerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdleInstanceCullerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdleInstanceCullerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdleInstanceCullerResponseBody) SetCode(v string) *GetIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetCpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody {
	s.CpuPercentThreshold = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetGpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody {
	s.GpuPercentThreshold = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetInstanceId(v string) *GetIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetMaxIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody {
	s.MaxIdleTimeInMinutes = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetMessage(v string) *GetIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetRequestId(v string) *GetIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetSuccess(v bool) *GetIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUsedUtilizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeClusterUsedUtilizationResponseBody
	GetCode() *string
	SetCpuTotal(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetCpuTotal() *float32
	SetDedicatedClusterId(v string) *DescribeClusterUsedUtilizationResponseBody
	GetDedicatedClusterId() *string
	SetDiskTotal(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetDiskTotal() *float32
	SetDiskUsed(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetDiskUsed() *float32
	SetDuTotal(v int32) *DescribeClusterUsedUtilizationResponseBody
	GetDuTotal() *int32
	SetDuUsed(v int32) *DescribeClusterUsedUtilizationResponseBody
	GetDuUsed() *int32
	SetDynamicMessage(v string) *DescribeClusterUsedUtilizationResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeClusterUsedUtilizationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeClusterUsedUtilizationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeClusterUsedUtilizationResponseBody
	GetHttpStatusCode() *int32
	SetMemoryTotal(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetMemoryTotal() *float32
	SetMemoryUsed(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetMemoryUsed() *float32
	SetMemoryUsedPercentage(v float32) *DescribeClusterUsedUtilizationResponseBody
	GetMemoryUsedPercentage() *float32
	SetRequestId(v string) *DescribeClusterUsedUtilizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeClusterUsedUtilizationResponseBody
	GetSuccess() *bool
	SetTaskRunning(v int32) *DescribeClusterUsedUtilizationResponseBody
	GetTaskRunning() *int32
}

type DescribeClusterUsedUtilizationResponseBody struct {
	// The error code returned by the backend service. The number is incremented.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The CPU utilization of the cluster. Unit: percentage.
	//
	// example:
	//
	// 50
	CpuTotal *float32 `json:"CpuTotal,omitempty" xml:"CpuTotal,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// dtscluster_h3fl1cs217sx952
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The total disk size of the cluster. Unit: GB.
	//
	// example:
	//
	// 1024
	DiskTotal *float32 `json:"DiskTotal,omitempty" xml:"DiskTotal,omitempty"`
	// The disk usage of the cluster. Unit: GB.
	//
	// example:
	//
	// 96
	DiskUsed *float32 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The total number of DTS units (DUs).
	//
	// example:
	//
	// 30
	DuTotal *int32 `json:"DuTotal,omitempty" xml:"DuTotal,omitempty"`
	// The usage of DUs.
	//
	// example:
	//
	// 2
	DuUsed *int32 `json:"DuUsed,omitempty" xml:"DuUsed,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace %s in the ErrMessage parameter.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The total amount of memory. A value of 0 is temporarily returned.
	//
	// example:
	//
	// 0
	MemoryTotal *float32 `json:"MemoryTotal,omitempty" xml:"MemoryTotal,omitempty"`
	// The memory usage. A value of 0 is temporarily returned.
	//
	// example:
	//
	// 0
	MemoryUsed *float32 `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	// The memory usage.
	//
	// example:
	//
	// 1.0
	MemoryUsedPercentage *float32 `json:"MemoryUsedPercentage,omitempty" xml:"MemoryUsedPercentage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of tasks that are in progress.
	//
	// example:
	//
	// 3
	TaskRunning *int32 `json:"TaskRunning,omitempty" xml:"TaskRunning,omitempty"`
}

func (s DescribeClusterUsedUtilizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUsedUtilizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetCpuTotal() *float32 {
	return s.CpuTotal
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDiskTotal() *float32 {
	return s.DiskTotal
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDiskUsed() *float32 {
	return s.DiskUsed
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDuTotal() *int32 {
	return s.DuTotal
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDuUsed() *int32 {
	return s.DuUsed
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetMemoryTotal() *float32 {
	return s.MemoryTotal
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetMemoryUsed() *float32 {
	return s.MemoryUsed
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetMemoryUsedPercentage() *float32 {
	return s.MemoryUsedPercentage
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeClusterUsedUtilizationResponseBody) GetTaskRunning() *int32 {
	return s.TaskRunning
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetCode(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetCpuTotal(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.CpuTotal = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDedicatedClusterId(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDiskTotal(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.DiskTotal = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDiskUsed(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.DiskUsed = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDuTotal(v int32) *DescribeClusterUsedUtilizationResponseBody {
	s.DuTotal = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDuUsed(v int32) *DescribeClusterUsedUtilizationResponseBody {
	s.DuUsed = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetDynamicMessage(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetErrCode(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetErrMessage(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetHttpStatusCode(v int32) *DescribeClusterUsedUtilizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetMemoryTotal(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.MemoryTotal = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetMemoryUsed(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetMemoryUsedPercentage(v float32) *DescribeClusterUsedUtilizationResponseBody {
	s.MemoryUsedPercentage = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetRequestId(v string) *DescribeClusterUsedUtilizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetSuccess(v bool) *DescribeClusterUsedUtilizationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) SetTaskRunning(v int32) *DescribeClusterUsedUtilizationResponseBody {
	s.TaskRunning = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponseBody) Validate() error {
	return dara.Validate(s)
}

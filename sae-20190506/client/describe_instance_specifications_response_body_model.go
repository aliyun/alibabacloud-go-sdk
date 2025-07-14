// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceSpecificationsResponseBody
	GetCode() *string
	SetData(v []*DescribeInstanceSpecificationsResponseBodyData) *DescribeInstanceSpecificationsResponseBody
	GetData() []*DescribeInstanceSpecificationsResponseBodyData
	SetErrorCode(v string) *DescribeInstanceSpecificationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeInstanceSpecificationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceSpecificationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceSpecificationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeInstanceSpecificationsResponseBody
	GetTraceId() *string
}

type DescribeInstanceSpecificationsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the instance types.
	Data []*DescribeInstanceSpecificationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the instance types were queried. Valid values:
	//
	// 	- **true**: The instance types were queried.
	//
	// 	- **false**: The instance types failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeInstanceSpecificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceSpecificationsResponseBody) GetData() []*DescribeInstanceSpecificationsResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceSpecificationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInstanceSpecificationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceSpecificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSpecificationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceSpecificationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeInstanceSpecificationsResponseBody) SetCode(v string) *DescribeInstanceSpecificationsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetData(v []*DescribeInstanceSpecificationsResponseBodyData) *DescribeInstanceSpecificationsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetErrorCode(v string) *DescribeInstanceSpecificationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetMessage(v string) *DescribeInstanceSpecificationsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetRequestId(v string) *DescribeInstanceSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetSuccess(v bool) *DescribeInstanceSpecificationsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetTraceId(v string) *DescribeInstanceSpecificationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSpecificationsResponseBodyData struct {
	// The CPU specification of the instance type. Unit: millicore.
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Indicates whether the instance type is available. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the instance type.
	//
	// example:
	//
	// 4
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The memory size of the instance type. Unit: MB.
	//
	// example:
	//
	// 4096
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the instance type.
	SpecInfo *string `json:"SpecInfo,omitempty" xml:"SpecInfo,omitempty"`
	// The version number of the instance type.
	//
	// example:
	//
	// 0
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceSpecificationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceSpecificationsResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetCpu(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetEnable(v bool) *DescribeInstanceSpecificationsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetId(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetMemory(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetSpecInfo(v string) *DescribeInstanceSpecificationsResponseBodyData {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetVersion(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

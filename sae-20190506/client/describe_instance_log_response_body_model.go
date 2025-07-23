// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceLogResponseBody
	GetCode() *string
	SetData(v string) *DescribeInstanceLogResponseBody
	GetData() *string
	SetErrorCode(v string) *DescribeInstanceLogResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeInstanceLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceLogResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeInstanceLogResponseBody
	GetTraceId() *string
}

type DescribeInstanceLogResponseBody struct {
	// The interface state or POP error code. Valid values:
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
	// The information of instance logs.
	//
	// example:
	//
	// hello\\nsae\\n
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error code.
	//
	// - No error code returned if the request succeeded.
	//
	// - Error code returned if the request failed. Refer to error code list below for details.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// success is returned when the request succeeds.
	//
	// An error code is returned when the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the logs of the instance is obtained.
	//
	// - true: logs obtained.
	//
	// - false: failed to obtain logs.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Trace ID.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeInstanceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceLogResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeInstanceLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInstanceLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceLogResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeInstanceLogResponseBody) SetCode(v string) *DescribeInstanceLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetData(v string) *DescribeInstanceLogResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetErrorCode(v string) *DescribeInstanceLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetMessage(v string) *DescribeInstanceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetRequestId(v string) *DescribeInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetSuccess(v bool) *DescribeInstanceLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetTraceId(v string) *DescribeInstanceLogResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) Validate() error {
	return dara.Validate(s)
}

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
	// The HTTP status code.
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance log.
	//
	// example:
	//
	// hello\\nsae\\n
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// - **success*	- is returned if the request is successful.
	//
	// - An error code is returned if the request fails.
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
	// Indicates whether the instance log was obtained. Valid values:
	//
	// - **true**: The instance log was obtained.
	//
	// - **false**: The instance log failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID.
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

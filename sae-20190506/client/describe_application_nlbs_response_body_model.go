// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationNlbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationNlbsResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationNlbsResponseBodyData) *DescribeApplicationNlbsResponseBody
	GetData() *DescribeApplicationNlbsResponseBodyData
	SetErrorCode(v string) *DescribeApplicationNlbsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationNlbsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationNlbsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeApplicationNlbsResponseBody
	GetSuccess() *string
	SetTraceId(v string) *DescribeApplicationNlbsResponseBody
	GetTraceId() *string
}

type DescribeApplicationNlbsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: A client error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned in the response.
	Data *DescribeApplicationNlbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is returned only if the request fails.
	//
	// - For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. A value of `success` indicates a successful request; otherwise, an error code is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the details of the request.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationNlbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationNlbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationNlbsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationNlbsResponseBody) GetData() *DescribeApplicationNlbsResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationNlbsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationNlbsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationNlbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationNlbsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeApplicationNlbsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationNlbsResponseBody) SetCode(v string) *DescribeApplicationNlbsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetData(v *DescribeApplicationNlbsResponseBodyData) *DescribeApplicationNlbsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetErrorCode(v string) *DescribeApplicationNlbsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetMessage(v string) *DescribeApplicationNlbsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetRequestId(v string) *DescribeApplicationNlbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetSuccess(v string) *DescribeApplicationNlbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) SetTraceId(v string) *DescribeApplicationNlbsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationNlbsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationNlbsResponseBodyData struct {
	// A collection of instance details.
	Instances map[string]*DataInstancesValue `json:"Instances,omitempty" xml:"Instances,omitempty"`
}

func (s DescribeApplicationNlbsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationNlbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationNlbsResponseBodyData) GetInstances() map[string]*DataInstancesValue {
	return s.Instances
}

func (s *DescribeApplicationNlbsResponseBodyData) SetInstances(v map[string]*DataInstancesValue) *DescribeApplicationNlbsResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeApplicationNlbsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

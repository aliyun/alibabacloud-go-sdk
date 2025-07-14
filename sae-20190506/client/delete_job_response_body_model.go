// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteJobResponseBody
	GetCode() *string
	SetData(v string) *DeleteJobResponseBody
	GetData() *string
	SetErrorCode(v string) *DeleteJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteJobResponseBody
	GetTraceId() *string
}

type DeleteJobResponseBody struct {
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
	// The result returned.
	//
	// example:
	//
	// {msg: "", code: 200, success: true}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code that is returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request is successful.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application is deleted. Valid values:
	//
	// 	- **true**: The namespaces were obtained.
	//
	// 	- **false**: no
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

func (s DeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteJobResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteJobResponseBody) SetCode(v string) *DeleteJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobResponseBody) SetData(v string) *DeleteJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteJobResponseBody) SetErrorCode(v string) *DeleteJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetMessage(v string) *DeleteJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobResponseBody) SetTraceId(v string) *DeleteJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationDescriptionResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateApplicationDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateApplicationDescriptionResponseBody
	GetTraceId() *string
}

type UpdateApplicationDescriptionResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned if the call failed. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. The following limits are imposed on the ID:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- An error code is returned when a request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application was deployed. Valid values:
	//
	// 	- **true**: The application was deployed.
	//
	// 	- **false**: The application failed to be deployed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationDescriptionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateApplicationDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApplicationDescriptionResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateApplicationDescriptionResponseBody) SetCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetMessage(v string) *UpdateApplicationDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetTraceId(v string) *UpdateApplicationDescriptionResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}

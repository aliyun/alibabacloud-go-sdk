// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneEnableAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateSwimmingLaneEnableAttributeResponseBody
	GetTraceId() *string
}

type UpdateSwimmingLaneEnableAttributeResponseBody struct {
	// The interface status or POP error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: Redirection.
	//
	// 	- **4xx**: Request error.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error code. Valid values:
	//
	// 	- If the request is successful, no **ErrorCode*	- fields are returned.
	//
	// 	- Request failed: **ErrorCode*	- fields are returned. For more information, see **Error codes**.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// 	- The error message returned because the request is normal and **success*	- is returned.
	//
	// 	- If the request is abnormal, the specific exception error code is returned.
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
	// Whether the data is successful. Valid values:
	//
	// 	- **true**: The policy was deleted.
	//
	// 	- **false**: The policy failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. This parameter is used to query the exact call information.
	//
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateSwimmingLaneEnableAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneEnableAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetCode(v string) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetErrorCode(v string) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetMessage(v string) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetRequestId(v string) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetSuccess(v bool) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) SetTraceId(v string) *UpdateSwimmingLaneEnableAttributeResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

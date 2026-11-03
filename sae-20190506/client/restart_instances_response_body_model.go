// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartInstancesResponseBody
	GetCode() *string
	SetData(v *RestartInstancesResponseBodyData) *RestartInstancesResponseBody
	GetData() *RestartInstancesResponseBodyData
	SetErrorCode(v string) *RestartInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RestartInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RestartInstancesResponseBody
	GetTraceId() *string
}

type RestartInstancesResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application information.
	Data *RestartInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - Not returned: The request is successful.
	//
	// - Specific error code: The request fails. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidApplication.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: An exception occurred during the request.
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
	// Specifies whether the application instances were restarted successfully. Valid values:
	//
	// - **true**: The restart was successful.
	//
	// - **false**: The restart failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to precisely query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RestartInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartInstancesResponseBody) GetData() *RestartInstancesResponseBodyData {
	return s.Data
}

func (s *RestartInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RestartInstancesResponseBody) SetCode(v string) *RestartInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *RestartInstancesResponseBody) SetData(v *RestartInstancesResponseBodyData) *RestartInstancesResponseBody {
	s.Data = v
	return s
}

func (s *RestartInstancesResponseBody) SetErrorCode(v string) *RestartInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstancesResponseBody) SetMessage(v string) *RestartInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *RestartInstancesResponseBody) SetRequestId(v string) *RestartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstancesResponseBody) SetSuccess(v bool) *RestartInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartInstancesResponseBody) SetTraceId(v string) *RestartInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *RestartInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RestartInstancesResponseBodyData struct {
	// The change order ID.
	//
	// example:
	//
	// 5afa5b98-0c64-4637-983f-15eaa888****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RestartInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RestartInstancesResponseBodyData) SetChangeOrderId(v string) *RestartInstancesResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RestartInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

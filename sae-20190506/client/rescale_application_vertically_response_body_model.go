// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationVerticallyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RescaleApplicationVerticallyResponseBody
	GetCode() *string
	SetData(v *RescaleApplicationVerticallyResponseBodyData) *RescaleApplicationVerticallyResponseBody
	GetData() *RescaleApplicationVerticallyResponseBodyData
	SetErrorCode(v string) *RescaleApplicationVerticallyResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RescaleApplicationVerticallyResponseBody
	GetMessage() *string
	SetRequestId(v string) *RescaleApplicationVerticallyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RescaleApplicationVerticallyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RescaleApplicationVerticallyResponseBody
	GetTraceId() *string
}

type RescaleApplicationVerticallyResponseBody struct {
	// The HTTP status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *RescaleApplicationVerticallyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// None
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AB521DBB-FA78-42E6-803F-A862EA4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the instance specifications were modified. Valid values:
	//
	// - **true**: The modification was successful.
	//
	// - **false**: The modification failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to accurately query call information.
	//
	// example:
	//
	// 0bc3b6f315637273629117900d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RescaleApplicationVerticallyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationVerticallyResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponseBody) GetCode() *string {
	return s.Code
}

func (s *RescaleApplicationVerticallyResponseBody) GetData() *RescaleApplicationVerticallyResponseBodyData {
	return s.Data
}

func (s *RescaleApplicationVerticallyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RescaleApplicationVerticallyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RescaleApplicationVerticallyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RescaleApplicationVerticallyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RescaleApplicationVerticallyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RescaleApplicationVerticallyResponseBody) SetCode(v string) *RescaleApplicationVerticallyResponseBody {
	s.Code = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetData(v *RescaleApplicationVerticallyResponseBodyData) *RescaleApplicationVerticallyResponseBody {
	s.Data = v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetErrorCode(v string) *RescaleApplicationVerticallyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetMessage(v string) *RescaleApplicationVerticallyResponseBody {
	s.Message = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetRequestId(v string) *RescaleApplicationVerticallyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetSuccess(v bool) *RescaleApplicationVerticallyResponseBody {
	s.Success = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetTraceId(v string) *RescaleApplicationVerticallyResponseBody {
	s.TraceId = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RescaleApplicationVerticallyResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// ffd8cd45-2b5f-415d-b4d0-1003e80b****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RescaleApplicationVerticallyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationVerticallyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RescaleApplicationVerticallyResponseBodyData) SetChangeOrderId(v string) *RescaleApplicationVerticallyResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

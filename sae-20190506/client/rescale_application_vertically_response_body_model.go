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
	// The HTTP status code. Valid values:
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
	// Data returned.
	Data *RescaleApplicationVerticallyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Messages returned for additional information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AB521DBB-FA78-42E6-803F-A862EA4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the update of instance specifications was successful. Valid values:
	//
	// 	- **true**: Updated.
	//
	// 	- **false**: Failed to update.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Trace ID for request information.
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
	// The ticked ID of updates.
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

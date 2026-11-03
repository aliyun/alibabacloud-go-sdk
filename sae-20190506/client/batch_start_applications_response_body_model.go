// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchStartApplicationsResponseBody
	GetCode() *string
	SetData(v *BatchStartApplicationsResponseBodyData) *BatchStartApplicationsResponseBody
	GetData() *BatchStartApplicationsResponseBodyData
	SetErrorCode(v string) *BatchStartApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BatchStartApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchStartApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchStartApplicationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchStartApplicationsResponseBody
	GetTraceId() *string
}

type BatchStartApplicationsResponseBody struct {
	// The API status code or POP error code. Valid values:
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
	Data *BatchStartApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// System.Upgrading
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If an exception occurs, a specific error code is returned.
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
	// Indicates whether the applications were started successfully. Valid values:
	//
	// - **true**: The applications were started successfully.
	//
	// - **false**: The applications failed to start.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStartApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchStartApplicationsResponseBody) GetData() *BatchStartApplicationsResponseBodyData {
	return s.Data
}

func (s *BatchStartApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchStartApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchStartApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchStartApplicationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchStartApplicationsResponseBody) SetCode(v string) *BatchStartApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetData(v *BatchStartApplicationsResponseBodyData) *BatchStartApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetErrorCode(v string) *BatchStartApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetMessage(v string) *BatchStartApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetRequestId(v string) *BatchStartApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetSuccess(v bool) *BatchStartApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetTraceId(v string) *BatchStartApplicationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchStartApplicationsResponseBodyData struct {
	// The change order ID.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStartApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BatchStartApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStartApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BatchStartApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

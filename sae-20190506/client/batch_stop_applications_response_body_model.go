// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchStopApplicationsResponseBody
	GetCode() *string
	SetData(v *BatchStopApplicationsResponseBodyData) *BatchStopApplicationsResponseBody
	GetData() *BatchStopApplicationsResponseBodyData
	SetErrorCode(v string) *BatchStopApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BatchStopApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchStopApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchStopApplicationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchStopApplicationsResponseBody
	GetTraceId() *string
}

type BatchStopApplicationsResponseBody struct {
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
	// The returned result.
	Data *BatchStopApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// - If the request is abnormal, the specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BD8F4C7-D84C-4D46-9885-8212997E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the applications were stopped successfully. Valid values:
	//
	// - **true**: The applications were stopped successfully.
	//
	// - **false**: The applications failed to be stopped.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call information accurately.
	//
	// example:
	//
	// 0bc3b6e215637275918588187d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStopApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchStopApplicationsResponseBody) GetData() *BatchStopApplicationsResponseBodyData {
	return s.Data
}

func (s *BatchStopApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchStopApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchStopApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchStopApplicationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchStopApplicationsResponseBody) SetCode(v string) *BatchStopApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetData(v *BatchStopApplicationsResponseBodyData) *BatchStopApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetErrorCode(v string) *BatchStopApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetMessage(v string) *BatchStopApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetRequestId(v string) *BatchStopApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetSuccess(v bool) *BatchStopApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetTraceId(v string) *BatchStopApplicationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchStopApplicationsResponseBodyData struct {
	// The change order ID.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStopApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BatchStopApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStopApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BatchStopApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

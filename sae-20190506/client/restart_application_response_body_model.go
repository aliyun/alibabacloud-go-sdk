// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartApplicationResponseBody
	GetCode() *string
	SetData(v *RestartApplicationResponseBodyData) *RestartApplicationResponseBody
	GetData() *RestartApplicationResponseBodyData
	SetErrorCode(v string) *RestartApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RestartApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RestartApplicationResponseBody
	GetTraceId() *string
}

type RestartApplicationResponseBody struct {
	// The HTTP status code. Take note of the following rules:
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
	// The response.
	Data *RestartApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Take note of the following rules:
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
	// Indicates whether the instance is successfully restarted. Take note of the following rules:
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s RestartApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartApplicationResponseBody) GetData() *RestartApplicationResponseBodyData {
	return s.Data
}

func (s *RestartApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RestartApplicationResponseBody) SetCode(v string) *RestartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartApplicationResponseBody) SetData(v *RestartApplicationResponseBodyData) *RestartApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RestartApplicationResponseBody) SetErrorCode(v string) *RestartApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartApplicationResponseBody) SetMessage(v string) *RestartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartApplicationResponseBody) SetRequestId(v string) *RestartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartApplicationResponseBody) SetSuccess(v bool) *RestartApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RestartApplicationResponseBody) SetTraceId(v string) *RestartApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *RestartApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestartApplicationResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RestartApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RestartApplicationResponseBodyData) SetChangeOrderId(v string) *RestartApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RestartApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

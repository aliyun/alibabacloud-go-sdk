// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateJobResponseBody
	GetCode() *string
	SetData(v *UpdateJobResponseBodyData) *UpdateJobResponseBody
	GetData() *UpdateJobResponseBodyData
	SetErrorCode(v string) *UpdateJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateJobResponseBody
	GetTraceId() *string
}

type UpdateJobResponseBody struct {
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
	// The response.
	Data *UpdateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Valid values:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
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
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application deployment is successful. Valid values:
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
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateJobResponseBody) GetData() *UpdateJobResponseBodyData {
	return s.Data
}

func (s *UpdateJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateJobResponseBody) SetCode(v string) *UpdateJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobResponseBody) SetData(v *UpdateJobResponseBodyData) *UpdateJobResponseBody {
	s.Data = v
	return s
}

func (s *UpdateJobResponseBody) SetErrorCode(v string) *UpdateJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateJobResponseBody) SetMessage(v string) *UpdateJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobResponseBody) SetSuccess(v bool) *UpdateJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateJobResponseBody) SetTraceId(v string) *UpdateJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateJobResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the change order. It can be used to query the task status.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s UpdateJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateJobResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UpdateJobResponseBodyData) SetAppId(v string) *UpdateJobResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateJobResponseBodyData) SetChangeOrderId(v string) *UpdateJobResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateJobResponseBody
	GetCode() *string
	SetData(v *CreateJobResponseBodyData) *CreateJobResponseBody
	GetData() *CreateJobResponseBodyData
	SetErrorCode(v string) *CreateJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateJobResponseBody
	GetTraceId() *string
}

type CreateJobResponseBody struct {
	// The HTTP status code or a POP error code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - If the request fails, this parameter is returned. For more information, see the **error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job template was created successfully. Valid values:
	//
	// - **true**: The job template was created.
	//
	// - **false**: The job template was not created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The call trace ID. You can use this ID to query detailed information about the call.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateJobResponseBody) GetData() *CreateJobResponseBodyData {
	return s.Data
}

func (s *CreateJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateJobResponseBody) SetCode(v string) *CreateJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobResponseBody) SetData(v *CreateJobResponseBodyData) *CreateJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateJobResponseBody) SetErrorCode(v string) *CreateJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateJobResponseBody) SetMessage(v string) *CreateJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v bool) *CreateJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobResponseBody) SetTraceId(v string) *CreateJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobResponseBodyData struct {
	// The job template ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The change order ID. You can use this ID to check the execution status of the task.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s CreateJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateJobResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *CreateJobResponseBodyData) SetAppId(v string) *CreateJobResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateJobResponseBodyData) SetChangeOrderId(v string) *CreateJobResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *CreateJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

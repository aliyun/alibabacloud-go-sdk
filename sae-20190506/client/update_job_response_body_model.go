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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: success.
	//
	// - **3xx**: redirection.
	//
	// - **4xx**: client error.
	//
	// - **5xx**: server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *UpdateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// Application.MissingJdk
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
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
	// Indicates whether the job template was updated. Valid values:
	//
	// - **true**: The job template was updated.
	//
	// - **false**: The job template failed to be updated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to accurately query the call information.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateJobResponseBodyData struct {
	// The ID of the job template.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the returned change order, which is used to query the job execution status.
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

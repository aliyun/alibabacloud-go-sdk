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
	// The HTTP status code or POP error code. Valid values:
	//
	// - **2xx**: success.
	//
	// - **3xx**: redirection.
	//
	// - **4xx**: request error.
	//
	// - **5xx**: server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the job template was created. Valid values:
	//
	// - **true**: created.
	//
	// - **false**: failed to create.
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

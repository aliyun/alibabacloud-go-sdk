// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApplicationResponseBody
	GetCode() *string
	SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody
	GetData() *CreateApplicationResponseBodyData
	SetErrorCode(v string) *CreateApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateApplicationResponseBody
	GetTraceId() *string
}

type CreateApplicationResponseBody struct {
	// The API status or POP error code. Valid values:
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
	Data *CreateApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request fails, the **ErrorCode*	- field is returned. For more information, see the error code list in this topic.
	//
	// example:
	//
	// 空
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// - If the request is normal, **success*	- is returned.
	//
	// - If the request is abnormal, a specific error code is returned.
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
	// Whether the application is created successfully. Valid values:
	//
	// - **true**: Created successfully.
	//
	// - **false**: Failed to create.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, used for term query of call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApplicationResponseBody) GetData() *CreateApplicationResponseBodyData {
	return s.Data
}

func (s *CreateApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateApplicationResponseBody) SetCode(v string) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationResponseBody) SetErrorCode(v string) *CreateApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetSuccess(v bool) *CreateApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApplicationResponseBody) SetTraceId(v string) *CreateApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationResponseBodyData struct {
	// The ID of the successfully created application.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The returned release order ID, used to query the task execution status.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s CreateApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *CreateApplicationResponseBodyData) SetAppId(v string) *CreateApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetChangeOrderId(v string) *CreateApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

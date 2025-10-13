// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConfigMapResponseBody
	GetCode() *string
	SetData(v *CreateConfigMapResponseBodyData) *CreateConfigMapResponseBody
	GetData() *CreateConfigMapResponseBodyData
	SetErrorCode(v string) *CreateConfigMapResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConfigMapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateConfigMapResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateConfigMapResponseBody
	GetTraceId() *string
}

type CreateConfigMapResponseBody struct {
	// Empty
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The trace ID that is used to query the details of the request.
	Data *CreateConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the ConfigMap that was created.
	//
	// example:
	//
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The returned message.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConfigMapResponseBody) GetData() *CreateConfigMapResponseBodyData {
	return s.Data
}

func (s *CreateConfigMapResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateConfigMapResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateConfigMapResponseBody) SetCode(v string) *CreateConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetData(v *CreateConfigMapResponseBodyData) *CreateConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *CreateConfigMapResponseBody) SetErrorCode(v string) *CreateConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetMessage(v string) *CreateConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetRequestId(v string) *CreateConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetSuccess(v bool) *CreateConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetTraceId(v string) *CreateConfigMapResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateConfigMapResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConfigMapResponseBodyData struct {
	// The returned result.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s CreateConfigMapResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponseBodyData) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *CreateConfigMapResponseBodyData) SetConfigMapId(v int64) *CreateConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

func (s *CreateConfigMapResponseBodyData) Validate() error {
	return dara.Validate(s)
}

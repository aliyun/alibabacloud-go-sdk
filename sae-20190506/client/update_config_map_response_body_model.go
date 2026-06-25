// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConfigMapResponseBody
	GetCode() *string
	SetData(v *UpdateConfigMapResponseBodyData) *UpdateConfigMapResponseBody
	GetData() *UpdateConfigMapResponseBodyData
	SetErrorCode(v string) *UpdateConfigMapResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConfigMapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConfigMapResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateConfigMapResponseBody
	GetTraceId() *string
}

type UpdateConfigMapResponseBody struct {
	// The HTTP status code or the POP error code.
	//
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *UpdateConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information about the call.
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
	// Indicates whether the ConfigMap instance was updated.
	//
	// - **true**: The instance was updated.
	//
	// - **false**: The instance failed to be updated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the call.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConfigMapResponseBody) GetData() *UpdateConfigMapResponseBodyData {
	return s.Data
}

func (s *UpdateConfigMapResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConfigMapResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateConfigMapResponseBody) SetCode(v string) *UpdateConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetData(v *UpdateConfigMapResponseBodyData) *UpdateConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *UpdateConfigMapResponseBody) SetErrorCode(v string) *UpdateConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetMessage(v string) *UpdateConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetRequestId(v string) *UpdateConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetSuccess(v bool) *UpdateConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetTraceId(v string) *UpdateConfigMapResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateConfigMapResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConfigMapResponseBodyData struct {
	// The ID of the ConfigMap instance.
	//
	// example:
	//
	// 1
	ConfigMapId *string `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s UpdateConfigMapResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponseBodyData) GetConfigMapId() *string {
	return s.ConfigMapId
}

func (s *UpdateConfigMapResponseBodyData) SetConfigMapId(v string) *UpdateConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

func (s *UpdateConfigMapResponseBodyData) Validate() error {
	return dara.Validate(s)
}

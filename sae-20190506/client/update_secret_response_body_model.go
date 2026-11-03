// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecretResponseBody
	GetCode() *string
	SetData(v *UpdateSecretResponseBodyData) *UpdateSecretResponseBody
	GetData() *UpdateSecretResponseBodyData
	SetErrorCode(v string) *UpdateSecretResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecretResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateSecretResponseBody
	GetTraceId() *string
}

type UpdateSecretResponseBody struct {
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
	Data *UpdateSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Take note of the following items:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
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
	// Indicates whether the Secret instance was updated successfully. Valid values:
	//
	// - **true**: The update was successful.
	//
	// - **false**: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecretResponseBody) GetData() *UpdateSecretResponseBodyData {
	return s.Data
}

func (s *UpdateSecretResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecretResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateSecretResponseBody) SetCode(v string) *UpdateSecretResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecretResponseBody) SetData(v *UpdateSecretResponseBodyData) *UpdateSecretResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSecretResponseBody) SetErrorCode(v string) *UpdateSecretResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSecretResponseBody) SetMessage(v string) *UpdateSecretResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecretResponseBody) SetRequestId(v string) *UpdateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretResponseBody) SetSuccess(v bool) *UpdateSecretResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecretResponseBody) SetTraceId(v string) *UpdateSecretResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecretResponseBodyData struct {
	// The instance ID of the Secret.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s UpdateSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponseBodyData) GetSecretId() *int64 {
	return s.SecretId
}

func (s *UpdateSecretResponseBodyData) SetSecretId(v int64) *UpdateSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *UpdateSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}

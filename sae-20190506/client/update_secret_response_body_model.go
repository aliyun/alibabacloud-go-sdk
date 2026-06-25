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
	// The status of the API call or a POP error code. Valid values:
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
	// The returned result.
	Data *UpdateSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. The following list describes the values:
	//
	// - If the request is successful, this parameter is not returned.
	//
	// - If the request fails, this parameter is returned. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information about the call result.
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
	// Indicates whether the Secret instance was updated. Valid values:
	//
	// - **true**: The instance was updated.
	//
	// - **false**: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the call chain. You can use this ID to query the details of a call.
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
	// The ID of the Secret instance.
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

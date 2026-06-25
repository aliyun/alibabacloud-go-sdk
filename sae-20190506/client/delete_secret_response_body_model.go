// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecretResponseBody
	GetCode() *string
	SetData(v *DeleteSecretResponseBodyData) *DeleteSecretResponseBody
	GetData() *DeleteSecretResponseBodyData
	SetErrorCode(v string) *DeleteSecretResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecretResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteSecretResponseBody
	GetTraceId() *string
}

type DeleteSecretResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A request error.
	//
	// - **5xx**: A server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - The **ErrorCode*	- parameter is not returned if the request is successful.
	//
	// - The **ErrorCode*	- parameter is returned if the request fails. For a list of error codes, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message.
	//
	// - **success**: The request was successful.
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
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the secret was successfully deleted. Valid values:
	//
	// - **true**: The secret was deleted.
	//
	// - **false**: The secret was not deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID used to trace the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecretResponseBody) GetData() *DeleteSecretResponseBodyData {
	return s.Data
}

func (s *DeleteSecretResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecretResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteSecretResponseBody) SetCode(v string) *DeleteSecretResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecretResponseBody) SetData(v *DeleteSecretResponseBodyData) *DeleteSecretResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSecretResponseBody) SetErrorCode(v string) *DeleteSecretResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSecretResponseBody) SetMessage(v string) *DeleteSecretResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSuccess(v bool) *DeleteSecretResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecretResponseBody) SetTraceId(v string) *DeleteSecretResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecretResponseBodyData struct {
	// The ID of the deleted secret.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DeleteSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBodyData) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DeleteSecretResponseBodyData) SetSecretId(v int64) *DeleteSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *DeleteSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}

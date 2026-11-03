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
	// The HTTP status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DeleteSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: An exception occurred.
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
	// Indicates whether the secret instance was deleted. Valid values:
	//
	// - **true**: The instance was deleted.
	//
	// - **false**: The instance failed to be deleted.
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
	// The instance ID of the deleted secret instance.
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

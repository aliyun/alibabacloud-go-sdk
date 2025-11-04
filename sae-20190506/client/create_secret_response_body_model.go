// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecretResponseBody
	GetCode() *string
	SetData(v *CreateSecretResponseBodyData) *CreateSecretResponseBody
	GetData() *CreateSecretResponseBodyData
	SetErrorCode(v string) *CreateSecretResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecretResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateSecretResponseBody
	GetTraceId() *string
}

type CreateSecretResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CreateSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the operation.
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
	// Indicates whether the Secret was created. Valid values:
	//
	// 	- **true**: The ConfigMap was created.
	//
	// 	- **false**: The ConfigMap failed to be created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecretResponseBody) GetData() *CreateSecretResponseBodyData {
	return s.Data
}

func (s *CreateSecretResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecretResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateSecretResponseBody) SetCode(v string) *CreateSecretResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecretResponseBody) SetData(v *CreateSecretResponseBodyData) *CreateSecretResponseBody {
	s.Data = v
	return s
}

func (s *CreateSecretResponseBody) SetErrorCode(v string) *CreateSecretResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSecretResponseBody) SetMessage(v string) *CreateSecretResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecretResponseBody) SetRequestId(v string) *CreateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponseBody) SetSuccess(v bool) *CreateSecretResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecretResponseBody) SetTraceId(v string) *CreateSecretResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateSecretResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecretResponseBodyData struct {
	// The ID of the created Secret.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s CreateSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBodyData) GetSecretId() *int64 {
	return s.SecretId
}

func (s *CreateSecretResponseBodyData) SetSecretId(v int64) *CreateSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *CreateSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}

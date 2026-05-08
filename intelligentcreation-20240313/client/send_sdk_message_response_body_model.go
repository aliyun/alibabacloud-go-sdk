// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SendSdkMessageResponseBody
	GetData() *string
	SetErrorCode(v string) *SendSdkMessageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SendSdkMessageResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SendSdkMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendSdkMessageResponseBody
	GetSuccess() *bool
}

type SendSdkMessageResponseBody struct {
	// example:
	//
	// {}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// system-01
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendSdkMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSdkMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendSdkMessageResponseBody) GetData() *string {
	return s.Data
}

func (s *SendSdkMessageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SendSdkMessageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SendSdkMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSdkMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendSdkMessageResponseBody) SetData(v string) *SendSdkMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetErrorCode(v string) *SendSdkMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetErrorMessage(v string) *SendSdkMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetRequestId(v string) *SendSdkMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSdkMessageResponseBody) SetSuccess(v bool) *SendSdkMessageResponseBody {
	s.Success = &v
	return s
}

func (s *SendSdkMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateNacosConfigResponseBody
	GetCode() *string
	SetErrorCode(v string) *CreateNacosConfigResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *CreateNacosConfigResponseBody
	GetHttpCode() *string
	SetMessage(v string) *CreateNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNacosConfigResponseBody
	GetSuccess() *bool
}

type CreateNacosConfigResponseBody struct {
	// The code returned.
	//
	// example:
	//
	// 1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 673DBD43-569E-510F-B3DE-20BB8DFEB20A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateNacosConfigResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNacosConfigResponseBody) SetCode(v string) *CreateNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetErrorCode(v string) *CreateNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetHttpCode(v string) *CreateNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetMessage(v string) *CreateNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetRequestId(v string) *CreateNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetSuccess(v bool) *CreateNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

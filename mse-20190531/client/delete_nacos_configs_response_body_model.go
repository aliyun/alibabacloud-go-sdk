// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteNacosConfigsResponseBody
	GetCode() *int32
	SetErrorCode(v string) *DeleteNacosConfigsResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteNacosConfigsResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteNacosConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNacosConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNacosConfigsResponseBody
	GetSuccess() *bool
}

type DeleteNacosConfigsResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73EB5563-CBB3-5F48-999D-512F4EFB7377
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

func (s DeleteNacosConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteNacosConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteNacosConfigsResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteNacosConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNacosConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNacosConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNacosConfigsResponseBody) SetCode(v int32) *DeleteNacosConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetErrorCode(v string) *DeleteNacosConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetHttpCode(v string) *DeleteNacosConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetMessage(v string) *DeleteNacosConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetRequestId(v string) *DeleteNacosConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetSuccess(v bool) *DeleteNacosConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

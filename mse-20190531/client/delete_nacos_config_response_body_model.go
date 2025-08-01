// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteNacosConfigResponseBody
	GetCode() *string
	SetErrorCode(v string) *DeleteNacosConfigResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteNacosConfigResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNacosConfigResponseBody
	GetSuccess() *bool
}

type DeleteNacosConfigResponseBody struct {
	// Return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// Message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4FEFC13F-EB50-51E1-97D8-C5CBA8CD1B84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, with the following values: - `true`: The request was successful. - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteNacosConfigResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNacosConfigResponseBody) SetCode(v string) *DeleteNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetErrorCode(v string) *DeleteNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetHttpCode(v string) *DeleteNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetMessage(v string) *DeleteNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetRequestId(v string) *DeleteNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetSuccess(v bool) *DeleteNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateNacosConfigResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *UpdateNacosConfigResponseBody
	GetHttpCode() *string
	SetMessage(v string) *UpdateNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNacosConfigResponseBody
	GetSuccess() *bool
}

type UpdateNacosConfigResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// NoPermission
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
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
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

func (s UpdateNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateNacosConfigResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNacosConfigResponseBody) SetErrorCode(v string) *UpdateNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetHttpCode(v string) *UpdateNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetMessage(v string) *UpdateNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetRequestId(v string) *UpdateNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetSuccess(v bool) *UpdateNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

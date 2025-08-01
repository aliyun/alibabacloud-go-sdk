// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateImageResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageResponseBody
	GetSuccess() *bool
}

type UpdateImageResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3855D475-2B66-5CFF-9A51-3D698E52C472
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageResponseBody) SetErrorCode(v string) *UpdateImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateImageResponseBody) SetMessage(v string) *UpdateImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageResponseBody) Validate() error {
	return dara.Validate(s)
}

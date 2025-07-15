// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateInstanceConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateInstanceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceConfigResponseBody
	GetSuccess() *bool
}

type UpdateInstanceConfigResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B6D821D-7F67-4CAA-9E13-A5A997C35***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateInstanceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceConfigResponseBody) SetCode(v int32) *UpdateInstanceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetMessage(v string) *UpdateInstanceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetRequestId(v string) *UpdateInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetSuccess(v bool) *UpdateInstanceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

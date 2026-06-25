// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateImageResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageResponseBody
	GetSuccess() *bool
}

type UpdateImageResponseBody struct {
	// The error code.
	//
	// - 0: The request was successful.
	//
	// - Non-zero: The request failed.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) GetCode() *int32 {
	return s.Code
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

func (s *UpdateImageResponseBody) SetCode(v int32) *UpdateImageResponseBody {
	s.Code = &v
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

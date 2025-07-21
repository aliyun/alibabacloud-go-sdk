// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindDeviceSeatsResponseBody
	GetCode() *string
	SetMessage(v string) *UnbindDeviceSeatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindDeviceSeatsResponseBody
	GetRequestId() *string
}

type UnbindDeviceSeatsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDeviceSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindDeviceSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindDeviceSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDeviceSeatsResponseBody) SetCode(v string) *UnbindDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindDeviceSeatsResponseBody) SetMessage(v string) *UnbindDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDeviceSeatsResponseBody) SetRequestId(v string) *UnbindDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDeviceSeatsResponseBody) Validate() error {
	return dara.Validate(s)
}

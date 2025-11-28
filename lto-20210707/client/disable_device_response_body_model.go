// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableDeviceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisableDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableDeviceResponseBody
	GetSuccess() *bool
}

type DisableDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableDeviceResponseBody) SetCode(v string) *DisableDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDeviceResponseBody) SetHttpStatusCode(v int32) *DisableDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableDeviceResponseBody) SetMessage(v string) *DisableDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDeviceResponseBody) SetRequestId(v string) *DisableDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDeviceResponseBody) SetSuccess(v bool) *DisableDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DisableDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}

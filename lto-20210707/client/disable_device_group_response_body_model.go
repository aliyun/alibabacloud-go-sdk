// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableDeviceGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisableDeviceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableDeviceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableDeviceGroupResponseBody
	GetSuccess() *bool
}

type DisableDeviceGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableDeviceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableDeviceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableDeviceGroupResponseBody) SetCode(v string) *DisableDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDeviceGroupResponseBody) SetHttpStatusCode(v int32) *DisableDeviceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableDeviceGroupResponseBody) SetMessage(v string) *DisableDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDeviceGroupResponseBody) SetRequestId(v string) *DisableDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDeviceGroupResponseBody) SetSuccess(v bool) *DisableDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DisableDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

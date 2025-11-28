// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDeviceGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddDeviceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDeviceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDeviceGroupResponseBody
	GetSuccess() *bool
}

type AddDeviceGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDeviceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDeviceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDeviceGroupResponseBody) SetCode(v string) *AddDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddDeviceGroupResponseBody) SetHttpStatusCode(v int32) *AddDeviceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDeviceGroupResponseBody) SetMessage(v string) *AddDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceGroupResponseBody) SetRequestId(v string) *AddDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceGroupResponseBody) SetSuccess(v bool) *AddDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

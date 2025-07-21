// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceBindedEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDeviceBindedEndUserResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateDeviceBindedEndUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDeviceBindedEndUserResponseBody
	GetRequestId() *string
}

type UpdateDeviceBindedEndUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceBindedEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceBindedEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDeviceBindedEndUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDeviceBindedEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetCode(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetMessage(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponseBody) SetRequestId(v string) *UpdateDeviceBindedEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserResponseBody) Validate() error {
	return dara.Validate(s)
}

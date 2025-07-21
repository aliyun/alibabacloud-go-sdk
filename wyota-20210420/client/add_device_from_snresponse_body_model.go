// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceFromSNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDeviceFromSNResponseBody
	GetCode() *string
	SetMessage(v string) *AddDeviceFromSNResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDeviceFromSNResponseBody
	GetRequestId() *string
}

type AddDeviceFromSNResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceFromSNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceFromSNResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDeviceFromSNResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDeviceFromSNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDeviceFromSNResponseBody) SetCode(v string) *AddDeviceFromSNResponseBody {
	s.Code = &v
	return s
}

func (s *AddDeviceFromSNResponseBody) SetMessage(v string) *AddDeviceFromSNResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceFromSNResponseBody) SetRequestId(v string) *AddDeviceFromSNResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceFromSNResponseBody) Validate() error {
	return dara.Validate(s)
}

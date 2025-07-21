// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaAutoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetDeviceOtaAutoStatusResponseBody
	GetCode() *string
	SetMessage(v string) *SetDeviceOtaAutoStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetDeviceOtaAutoStatusResponseBody
	GetRequestId() *string
}

type SetDeviceOtaAutoStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeviceOtaAutoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaAutoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetDeviceOtaAutoStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetDeviceOtaAutoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetCode(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetMessage(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponseBody) SetRequestId(v string) *SetDeviceOtaAutoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

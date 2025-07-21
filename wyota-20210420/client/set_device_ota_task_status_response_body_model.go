// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetDeviceOtaTaskStatusResponseBody
	GetCode() *string
	SetMessage(v string) *SetDeviceOtaTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetDeviceOtaTaskStatusResponseBody
	GetRequestId() *string
}

type SetDeviceOtaTaskStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeviceOtaTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetDeviceOtaTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetDeviceOtaTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetCode(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetMessage(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponseBody) SetRequestId(v string) *SetDeviceOtaTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesFromCSVResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDevicesFromCSVResponseBody
	GetCode() *string
	SetMessage(v string) *AddDevicesFromCSVResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDevicesFromCSVResponseBody
	GetRequestId() *string
}

type AddDevicesFromCSVResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDevicesFromCSVResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesFromCSVResponseBody) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDevicesFromCSVResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDevicesFromCSVResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDevicesFromCSVResponseBody) SetCode(v string) *AddDevicesFromCSVResponseBody {
	s.Code = &v
	return s
}

func (s *AddDevicesFromCSVResponseBody) SetMessage(v string) *AddDevicesFromCSVResponseBody {
	s.Message = &v
	return s
}

func (s *AddDevicesFromCSVResponseBody) SetRequestId(v string) *AddDevicesFromCSVResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDevicesFromCSVResponseBody) Validate() error {
	return dara.Validate(s)
}

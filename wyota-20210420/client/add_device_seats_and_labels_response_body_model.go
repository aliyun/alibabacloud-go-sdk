// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceSeatsAndLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDeviceSeatsAndLabelsResponseBody
	GetCode() *string
	SetMessage(v string) *AddDeviceSeatsAndLabelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDeviceSeatsAndLabelsResponseBody
	GetRequestId() *string
}

type AddDeviceSeatsAndLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceSeatsAndLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDeviceSeatsAndLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDeviceSeatsAndLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetCode(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetMessage(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponseBody) SetRequestId(v string) *AddDeviceSeatsAndLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}

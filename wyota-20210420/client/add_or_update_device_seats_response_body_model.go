// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDeviceSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddOrUpdateDeviceSeatsResponseBody
	GetCode() *string
	SetMessage(v string) *AddOrUpdateDeviceSeatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrUpdateDeviceSeatsResponseBody
	GetRequestId() *string
}

type AddOrUpdateDeviceSeatsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrUpdateDeviceSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddOrUpdateDeviceSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrUpdateDeviceSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetCode(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetMessage(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponseBody) SetRequestId(v string) *AddOrUpdateDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponseBody) Validate() error {
	return dara.Validate(s)
}

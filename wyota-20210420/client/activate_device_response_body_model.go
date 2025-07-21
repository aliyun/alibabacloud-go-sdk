// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActivateDeviceResponseBody
	GetCode() *string
	SetMessage(v string) *ActivateDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActivateDeviceResponseBody
	GetRequestId() *string
}

type ActivateDeviceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActivateDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActivateDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateDeviceResponseBody) SetCode(v string) *ActivateDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateDeviceResponseBody) SetMessage(v string) *ActivateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateDeviceResponseBody) SetRequestId(v string) *ActivateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}

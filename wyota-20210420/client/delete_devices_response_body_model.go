// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDevicesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDevicesResponseBody
	GetRequestId() *string
}

type DeleteDevicesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDevicesResponseBody) SetCode(v string) *DeleteDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDevicesResponseBody) SetMessage(v string) *DeleteDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDevicesResponseBody) SetRequestId(v string) *DeleteDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

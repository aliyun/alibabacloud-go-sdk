// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeviceUpdateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeviceUpdateResponseBody
	GetSuccess() *bool
}

type DeviceUpdateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeviceUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeviceUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeviceUpdateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeviceUpdateResponseBody) SetRequestId(v string) *DeviceUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceUpdateResponseBody) SetSuccess(v bool) *DeviceUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *DeviceUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}

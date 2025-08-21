// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetDeviceStatusInfoShrinkRequest
	GetDeviceInfoShrink() *string
}

type GetDeviceStatusInfoShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceStatusInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetDeviceStatusInfoShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceStatusInfoShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceStatusInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

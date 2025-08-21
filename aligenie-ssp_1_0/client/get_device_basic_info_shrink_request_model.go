// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceBasicInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetDeviceBasicInfoShrinkRequest
	GetDeviceInfoShrink() *string
}

type GetDeviceBasicInfoShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceBasicInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetDeviceBasicInfoShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceBasicInfoShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceBasicInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

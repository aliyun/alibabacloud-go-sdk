// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetDeviceTagShrinkRequest
	GetDeviceInfoShrink() *string
}

type GetDeviceTagShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetDeviceTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceTagShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetDeviceTagShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceTagShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByDeviceIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetUserByDeviceIdShrinkRequest
	GetDeviceInfoShrink() *string
}

type GetUserByDeviceIdShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s GetUserByDeviceIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetUserByDeviceIdShrinkRequest) SetDeviceInfoShrink(v string) *GetUserByDeviceIdShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetUserByDeviceIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}

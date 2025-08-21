// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBasicInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfosShrink(v string) *ListDeviceBasicInfoShrinkRequest
	GetDeviceInfosShrink() *string
}

type ListDeviceBasicInfoShrinkRequest struct {
	DeviceInfosShrink *string `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty"`
}

func (s ListDeviceBasicInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoShrinkRequest) GetDeviceInfosShrink() *string {
	return s.DeviceInfosShrink
}

func (s *ListDeviceBasicInfoShrinkRequest) SetDeviceInfosShrink(v string) *ListDeviceBasicInfoShrinkRequest {
	s.DeviceInfosShrink = &v
	return s
}

func (s *ListDeviceBasicInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetDeviceStatusDetailShrinkRequest
	GetDeviceInfoShrink() *string
	SetKeysShrink(v string) *GetDeviceStatusDetailShrinkRequest
	GetKeysShrink() *string
}

type GetDeviceStatusDetailShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s GetDeviceStatusDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetDeviceStatusDetailShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *GetDeviceStatusDetailShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceStatusDetailShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceStatusDetailShrinkRequest) SetKeysShrink(v string) *GetDeviceStatusDetailShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *GetDeviceStatusDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}

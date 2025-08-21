// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetDeviceSettingShrinkRequest
	GetDeviceInfoShrink() *string
	SetKeysShrink(v string) *GetDeviceSettingShrinkRequest
	GetKeysShrink() *string
}

type GetDeviceSettingShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s GetDeviceSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetDeviceSettingShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *GetDeviceSettingShrinkRequest) SetDeviceInfoShrink(v string) *GetDeviceSettingShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetDeviceSettingShrinkRequest) SetKeysShrink(v string) *GetDeviceSettingShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *GetDeviceSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}

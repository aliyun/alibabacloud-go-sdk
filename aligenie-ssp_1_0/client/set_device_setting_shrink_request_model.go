// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *SetDeviceSettingShrinkRequest
	GetDeviceInfoShrink() *string
	SetKey(v string) *SetDeviceSettingShrinkRequest
	GetKey() *string
	SetValue(v interface{}) *SetDeviceSettingShrinkRequest
	GetValue() interface{}
}

type SetDeviceSettingShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nightMode
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// {"enable":true}
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDeviceSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *SetDeviceSettingShrinkRequest) GetKey() *string {
	return s.Key
}

func (s *SetDeviceSettingShrinkRequest) GetValue() interface{} {
	return s.Value
}

func (s *SetDeviceSettingShrinkRequest) SetDeviceInfoShrink(v string) *SetDeviceSettingShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SetDeviceSettingShrinkRequest) SetKey(v string) *SetDeviceSettingShrinkRequest {
	s.Key = &v
	return s
}

func (s *SetDeviceSettingShrinkRequest) SetValue(v interface{}) *SetDeviceSettingShrinkRequest {
	s.Value = v
	return s
}

func (s *SetDeviceSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}

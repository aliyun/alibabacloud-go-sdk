// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *UnbindPhoneRequest
	GetAppKey() *int64
	SetDeviceId(v string) *UnbindPhoneRequest
	GetDeviceId() *string
}

type UnbindPhoneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eb5f741d83d04d34807d229999eefa52
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s UnbindPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindPhoneRequest) GoString() string {
	return s.String()
}

func (s *UnbindPhoneRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *UnbindPhoneRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnbindPhoneRequest) SetAppKey(v int64) *UnbindPhoneRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindPhoneRequest) SetDeviceId(v string) *UnbindPhoneRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindPhoneRequest) Validate() error {
	return dara.Validate(s)
}

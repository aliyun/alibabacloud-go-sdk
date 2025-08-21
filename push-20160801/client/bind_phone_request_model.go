// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *BindPhoneRequest
	GetAppKey() *int64
	SetDeviceId(v string) *BindPhoneRequest
	GetDeviceId() *string
	SetPhoneNumber(v string) *BindPhoneRequest
	GetPhoneNumber() *string
}

type BindPhoneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 27725900
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eb5f741d83d04d34807d229999eefa52
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1381111****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s BindPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPhoneRequest) GoString() string {
	return s.String()
}

func (s *BindPhoneRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *BindPhoneRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindPhoneRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *BindPhoneRequest) SetAppKey(v int64) *BindPhoneRequest {
	s.AppKey = &v
	return s
}

func (s *BindPhoneRequest) SetDeviceId(v string) *BindPhoneRequest {
	s.DeviceId = &v
	return s
}

func (s *BindPhoneRequest) SetPhoneNumber(v string) *BindPhoneRequest {
	s.PhoneNumber = &v
	return s
}

func (s *BindPhoneRequest) Validate() error {
	return dara.Validate(s)
}

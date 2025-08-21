// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *CheckDeviceRequest
	GetAppKey() *int64
	SetDeviceId(v string) *CheckDeviceRequest
	GetDeviceId() *string
}

type CheckDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23419851
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae296f3b04a58a05b30c95f****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s CheckDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDeviceRequest) GoString() string {
	return s.String()
}

func (s *CheckDeviceRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CheckDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CheckDeviceRequest) SetAppKey(v int64) *CheckDeviceRequest {
	s.AppKey = &v
	return s
}

func (s *CheckDeviceRequest) SetDeviceId(v string) *CheckDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *CheckDeviceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetTlogDeviceInfoRequest
	GetAppKey() *int64
	SetDeviceId(v string) *GetTlogDeviceInfoRequest
	GetDeviceId() *string
	SetOs(v string) *GetTlogDeviceInfoRequest
	GetOs() *string
}

type GetTlogDeviceInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// ad-0008ane9g0qcyu90bpm1-829
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
}

func (s GetTlogDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceInfoRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetTlogDeviceInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetTlogDeviceInfoRequest) GetOs() *string {
	return s.Os
}

func (s *GetTlogDeviceInfoRequest) SetAppKey(v int64) *GetTlogDeviceInfoRequest {
	s.AppKey = &v
	return s
}

func (s *GetTlogDeviceInfoRequest) SetDeviceId(v string) *GetTlogDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetTlogDeviceInfoRequest) SetOs(v string) *GetTlogDeviceInfoRequest {
	s.Os = &v
	return s
}

func (s *GetTlogDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}

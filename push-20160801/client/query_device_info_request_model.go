// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryDeviceInfoRequest
	GetAppKey() *int64
	SetDeviceId(v string) *QueryDeviceInfoRequest
	GetDeviceId() *string
}

type QueryDeviceInfoRequest struct {
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
	// a64ae296f3b04a58a05b30c9****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s QueryDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryDeviceInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *QueryDeviceInfoRequest) SetAppKey(v int64) *QueryDeviceInfoRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDeviceInfoRequest) SetDeviceId(v string) *QueryDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *QueryDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}

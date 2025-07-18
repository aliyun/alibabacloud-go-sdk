// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceTag(v string) *GetUserDeviceRequest
	GetDeviceTag() *string
}

type GetUserDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
}

func (s GetUserDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetUserDeviceRequest) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *GetUserDeviceRequest) SetDeviceTag(v string) *GetUserDeviceRequest {
	s.DeviceTag = &v
	return s
}

func (s *GetUserDeviceRequest) Validate() error {
	return dara.Validate(s)
}

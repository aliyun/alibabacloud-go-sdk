// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceTags(v []*string) *DeleteUserDevicesRequest
	GetDeviceTags() []*string
}

type DeleteUserDevicesRequest struct {
	// The collection of endpoint device IDs. A maximum of 100 entries are supported.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
}

func (s DeleteUserDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesRequest) GetDeviceTags() []*string {
	return s.DeviceTags
}

func (s *DeleteUserDevicesRequest) SetDeviceTags(v []*string) *DeleteUserDevicesRequest {
	s.DeviceTags = v
	return s
}

func (s *DeleteUserDevicesRequest) Validate() error {
	return dara.Validate(s)
}

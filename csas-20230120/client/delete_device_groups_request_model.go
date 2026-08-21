// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupIds(v []*string) *DeleteDeviceGroupsRequest
	GetDeviceGroupIds() []*string
}

type DeleteDeviceGroupsRequest struct {
	// The collection of instance tag IDs to delete. Duplicate values are not allowed.
	DeviceGroupIds []*string `json:"DeviceGroupIds,omitempty" xml:"DeviceGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteDeviceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceGroupsRequest) GetDeviceGroupIds() []*string {
	return s.DeviceGroupIds
}

func (s *DeleteDeviceGroupsRequest) SetDeviceGroupIds(v []*string) *DeleteDeviceGroupsRequest {
	s.DeviceGroupIds = v
	return s
}

func (s *DeleteDeviceGroupsRequest) Validate() error {
	return dara.Validate(s)
}

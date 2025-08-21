// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindParentPlatformDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchBindParentPlatformDevicesRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *BatchBindParentPlatformDevicesRequest
	GetOwnerId() *int64
	SetParentPlatformId(v string) *BatchBindParentPlatformDevicesRequest
	GetParentPlatformId() *string
}

type BatchBindParentPlatformDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 361*****212-cn-qingdao
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchBindParentPlatformDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchBindParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindParentPlatformDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchBindParentPlatformDevicesRequest) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *BatchBindParentPlatformDevicesRequest) SetDeviceId(v string) *BatchBindParentPlatformDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesRequest) SetOwnerId(v int64) *BatchBindParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesRequest) SetParentPlatformId(v string) *BatchBindParentPlatformDevicesRequest {
	s.ParentPlatformId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesRequest) Validate() error {
	return dara.Validate(s)
}
